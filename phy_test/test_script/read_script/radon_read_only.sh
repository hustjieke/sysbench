#!/bin/bash
#sysbench --test=oltp --mysql-db=testdb --mysql-user=test --mysql-password=test --mysql-host=192.168.1.240 --mysql-port=3306 --oltp-table-size=100000  prepare
# sysbench 0.4.12
#echo "======== create database ========"
#mysql -utest -ptest -h192.168.1.248 -e"create database testdb"
#echo "======== run ========"
#sysbench --test=oltp --mysql-db=testdb --mysql-user=test --mysql-password=test --mysql-host=192.168.1.248 --mysql-port=3306 --oltp-table-size=1000 --max-time=20 run
#echo "======== cleanup ========"
#sysbench --test=oltp --mysql-db=testdb --mysql-user=test --mysql-password=test --mysql-host=192.168.1.248 --mysql-port=3306 --oltp-table-size=1000 --max-time=20 cleanup
#
echo "======== read only ========"
sysbench ../../RadonDB_Mysql_test_lua/oltp_read_only.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --db-ps-mode=disable --tables=16 --threads=512 --report-interval=3 --events=200000000 --time=7200 run > oltp_read_only.log 2>&1 &
