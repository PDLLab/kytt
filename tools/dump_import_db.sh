#!/bin/bash

# 判断mysql_server是否已经启动 
docker ps -a | grep mysql_server
if [ $? != 0 ]; then
    echo "Please Start mysql_server Container First"
    exit 0
fi

if [ $# != 1 ]; then
    echo "Please select your action, dump or import ?"
    exit 0
fi

if [ $1 = "dump" ];then
    cd ../
    echo "Start dump data"
    docker exec mysql_server sh -c 'exec mysqldump -uroot -p"$MYSQL_ROOT_PASSWORD" kaoyantoutiao' > ./backup/mysql_db_backup/kaoyantoutiao.sql
fi

if [ $1 = "import" ];then
    cd ../
    echo "Start import data"
    docker exec mysql_server sh -c 'exec mysqladmin -uroot -p"$MYSQL_ROOT_PASSWORD" create kaoyantoutiao'
    docker exec -i mysql_server sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" kaoyantoutiao' < ./backup/mysql_db_backup/kaoyantoutiao.sql
fi
