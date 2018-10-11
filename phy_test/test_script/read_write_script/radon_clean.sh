echo "======== clean tables ========"
#sysbench ../mysql_test_lua/oltp_common.lua --mysql-db=sbtest --mysql-user=root --mysql-password=123456 --mysql-port=3306 --db-driver=mysql prepare
#sysbench ../RadonDB_test_lua/oltp_common.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --table_size=500000 --tables=32 --threads=64 --report-interval=60 cleanup
sysbench ../../RadonDB_Mysql_test_lua/oltp_common.lua --mysql-db=sbtest --mysql-user=usr --mysql-password=123456 --mysql-host=172.31.102.94 --mysql-port=3306 --db-driver=mysql --tables=16 --threads=64 cleanup
