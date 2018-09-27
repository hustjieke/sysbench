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
echo "======== read write ========"
sysbench oltp_read_only.lua --mysql-db=sbtest --mysql-user=root --mysql-password=123456 --mysql-host=127.0.0.1 --mysql-port=3308 --db-driver=mysql --db-ps-mode=disable --time=5 run > oltp_read_only.log 2>&1 &
