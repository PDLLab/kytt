#!/bin/bash

# 判断mysql_server是否已经启动
docker ps -a | grep mysql_server
if [ $? != 0 ]; then
    echo "Please Start mysql_server Container First"
    exit 1
fi

# 启动mysql client容器
echo "Start Mysql Client Container"
docker run -it --link mysql_server:mysql --rm mysql:5.7 sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
