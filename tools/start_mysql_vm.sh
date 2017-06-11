#!/bin/bash

# 准备镜像
tag=$(docker images | grep mysql | grep 5.7 | wc -l)
if [ $tag = 1 ]; then
    echo "Image Exist"
else
    echo "Image Not Exist"
    echo "Download mysql:5.7 Image"
    docker pull mysql:5.7
fi

# 清除已经启动的mysql_server容器
docker ps -a | grep mysql_server
if [ $? = 0 ]; then
    echo "Force Remove mysql_server Container"
    docker stop mysql_server
    docker rm mysql_server
fi

# 启动mysql_server容器
cd ../
echo "Start Mysql Server Container"
# for production environment
# docker run --name mysql_server -v $(pwd)/data/mysql_db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

# for testing environment
docker run --name mysql_server -p 3306:3306 -v $(pwd)/volumn/mysql_db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=kytt207 -d mysql:5.7

