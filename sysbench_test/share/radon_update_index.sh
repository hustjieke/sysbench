#!/bin/bash
#sysbench --test=oltp --mysql-db=testdb --mysql-user=test --mysql-password=test --mysql-host=192.168.1.240 --mysql-port=3306 --oltp-table-size=100000  prepare
# sysbench 0.4.12
#echo "======== create database ========"
#mysql -utest -ptest -h192.168.1.248 -e"create database testdb"
#echo "======== prepare ========"
#sysbench oltp_common.lua --mysql-db=sbtest --mysql-user=root --mysql-password=123456 --mysql-host=127.0.0.1 --mysql-port=3309 --db-driver=mysql prepare
echo "======== update index ========"
sysbench oltp_update_index.lua --mysql-db=sbtest --mysql-user=root --mysql-password=123456 --mysql-host=127.0.0.1 --mysql-port=3308 --db-driver=mysql --time=5 --db-ps-mode=disable run > oltp_update_index.log 2>&1 &
