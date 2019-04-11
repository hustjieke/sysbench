#!/usr/bin/env python
# -*- coding: utf-8 -*-
import base64
import collections
import json
import yaml
import sys
import os
import time
import traceback
import logging.handlers
import ruamel.yaml
import sqlite3
from configobj import ConfigObj
from ruamel.yaml.scalarstring import SingleQuotedScalarString
from filelock import FileLock
from utils import (
    init_logger,
    get_json_params,
    exec_cmd,
    check_local_service,
    read_file,
    wait_conf_file_ready,
    http_request,
    write_json_file,
    read_json_file
)

SUPPORTED_ACTIONS = {
    "start": "start datatravel",
    "updateparam": "update datatravel parameters",
    "startalertmanager": "start alertmanager",
    "stopalertmanager": "stop alertmanager",
    "travelprogress": "show alertmanager"
    #"travelprogress": "show the progress rate of datatravel"
}
DATATRAVEL_VERSION_FILE= "/etc/datatravel/version"
DATATRAVEL_CNF = "/etc/datatravel/datatravel.conf"
IP_FILE = "/etc/datatravel/ip"
LOCK_TIMEOUT = 300

def load_params():
    logger.info("load params start:")
    ret = wait_conf_file_ready(DATATRAVEL_CNF)
    if ret != 0:
        return None
    
    params = {}
    f = open(DATATRAVEL_CNF, 'r')
    for line in open(DATATRAVEL_CNF):
        line = f.readline()
        param = line.split("=")
        logger.info("param_split:%s",param)
        if len(param) != 2:
            continue
        
        key = param[0].strip()
        val = param[1].strip('\n')
        val = param[1].strip()
	logger.info("val:%s", val)

        if key[0] == '#':
            continue

        if key and val:
            params[key] = val

    rets = wait_conf_file_ready(IP_FILE)
    if rets != 0:
        return None
    ip = read_file(IP_FILE)
    if not ip:
        logger.error("load_params get ip failed")
        return None
    params["local_ip"] = ip
    
    logger.info("load params:%s",params)
    return params


from_database=""
from_host_mysqlPort="127.0.0.1:3306"
from_password=""
from_table=""
from_user=""
to_database=""
to_host_mysqlPort=""
to_password=""
to_table=""
to_user=""

def start_datatravel(params, flock_path):
    logger.info("start datatravel")

    generate_cnf(params, flock_path)

    cmd= "/opt/datatravel/bin/datatravel --mysqldump=/usr/local/bin/mysqldump --from=%s --from-user=%s --from-password=%s --to=%s --to-user=%s --to-password=%s  --to-flavor=%s --set-global-read-lock=%s --checksum=%s >> /data/log/datatravel.log 2>&1 &" % (params["from_host_mysqlPort"], params["from_user"], params["from_password"], params["to_host_mysqlPort"] ,params["to_user"], params["to_password"], params["to_flavor"], params["set_global_read_lock"], params["checksum"])

    logger.info("datatravel cmd:%s", cmd)
    ret_code, output = exec_cmd(cmd)
    if ret_code != 0:
        logger.error('datatravel start failed, reason %s' % output)
        return -1

    time.sleep(3)

    logger.info("datatravel start succeeded")

    return 0

def update_params(params, flock_path):
    logger.info("update_params")

    with FileLock(flock_path, LOCK_TIMEOUT, stealing=True) as locked:
        if not locked.is_locked:
            logger.error("update_params get lock failed")
            return -1
    if not os.path.exists(DATATRAVEL_CNF):
    	logger.info("datatravel config path does not exist")

    logger.info("update datatravel config params")
    generate_cnf(params, flock_path)
    logger.info("update datatravel config params succeeded")
    return 0

def generate_cnf(params, flock_path):
    logger.info("generate_cnf")
    with FileLock(flock_path, LOCK_TIMEOUT, stealing=True) as locked:
        if not locked.is_locked:
            logger.error("generate_cnf get lock failed")
            return -1
        if not os.path.exists(DATATRAVEL_CNF):
            logger.error("generate_cnf get grafana.ini failed")
            return -1
        config = ConfigObj(DATATRAVEL_CNF)
        config.write()
    logger.info("generate_cnf succeeded")
    return 0

def start_alertmanager():
    logger.info("start_alertmanager")

    if check_local_service("alertmanager", 9093):
        logger.warning("start_alertmanager skipped")
    else:
        ret_code, output = exec_cmd('service alertmanager start')
        if ret_code != 0:
            logger.error('start_alertmanager failed, reason %s' % output)
            return -1
    logger.info("start_alertmanager succeeded")
    return 0

def stop_alertmanager():
    logger.info("stop_alertmanager")

    if check_local_service("alertmanager", 9093):
        ret_code, output = exec_cmd('service alertmanager stop')
        if ret_code != 0:
            logger.error('stop_alertmanager failed, reason %s' % output)
            return -1
    logger.info("stop_alertmanager succeeded")
    return 0

def travel_progress():
    logger.info("show travel progress")

    ret_code, output = exec_cmd("/opt/datatravel/bin/datatravelcli datatravel progressrate")
    if ret_code != 0:
        print("show travel progress fail")
        logger.error('execute datatravel failed, reason: %s' % output)
        return -1
    logger.info("test:", output)
    status = json.loads(output)
    DumpProgress = status["dump-progress"].split(":")[0]
    DumpRemainTime = status["remain-time"].split(":")[0]
    IncreBehinds = status["position-behinds"].split(":")[0]

    travelProgress = {
       "labels": ["全量迁移进度", "全量迁移剩余时间估计", "增量迁移落后源集群位点pos"],
       "data": [[DumpProgress, DumpRemainTime, IncreBehinds]]
    }

    print(json.dumps(travelProgress))


def print_usage():
    print "usage:\n"
    #new_actions = sorted(SUPPORTED_ACTIONS.items(), lambda x, y: cmp(x[0], y[0]))
    #for key, val in new_actions:
    for key, val in SUPPORTED_ACTIONS.iteritems():
        print "    %-20s    %s" % (key, val)
    print "\n"


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print_usage()
        exit(-1)

    action = sys.argv[1]
    if action not in SUPPORTED_ACTIONS.keys():
        print_usage()
        exit(-1)

    logger = init_logger(action, "/data/log")
    logger.info(sys.argv)

    #version = read_file(DATATRAVEL_VERSION_FILE)
    version = "datatravel_1"
    if not version:
        logger.error("read_version failed")
        exit(-1)
    flock_path = "/tmp/%s" %version
    params = load_params()

    ret = 0
    try:
        if action == "start":
            ret = start_datatravel(params, flock_path)
        elif action == "updateparam":
            ret = update_params(params, flock_path)
        elif action == "startalertmanager":
            ret = start_alertmanager()
        elif action == "stopalertmanager":
            ret = stop_alertmanager()
        elif action == "travelprogress":
            ret = travel_progress()
        exit(ret)
    except Exception, e:
        logger.error("%s" % traceback.format_exc())
        exit(-1)
