echo "======== clean ========"
sysbench  oltp_common.lua --mysql-db=sbtest --mysql-user=root --mysql-password=123456 --mysql-host=127.0.0.1 --mysql-port=3308 --db-driver=mysql --time=5 cleanup
