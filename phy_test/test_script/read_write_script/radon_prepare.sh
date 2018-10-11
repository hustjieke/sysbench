#!/bin/bash
#sysbench --test=oltp --mysql-db=testdb --mysql-user=test --mysql-password=test --mysql-host=192.168.1.240 --mysql-port=3306 --oltp-table-size=100000  prepare
# sysbench 0.4.12
#echo "======== create database ========"
#mysql -utest -ptest -h192.168.1.248 -e"create database testdb"
echo "======== prepare tables ========"
sysbench ../../RadonDB_Mysql_test_lua/oltp_common.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --table_size=0 --tables=16 --threads=64 prepare
