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

#sysbench ../RadonDB_test_lua/oltp_read_write.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --db-ps-mode=disable --tables=16 --threads=512 --report-interval=5 --events=2000 --delete-inserts=2000 --auto-inc=off --time=20 run > oltp_insert.log 2>&1 

sysbench ../../RadonDB_Mysql_test_lua/oltp_read_write.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --db-ps-mode=disable --tables=16 --threads=512 --report-interval=5 --auto-inc=off --events=100000000 --time=86400 --delete-inserts=2 --skip-trx=on run > oltp_read_write.log 2>&1 &
