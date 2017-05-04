#!/bin/bash

# 准备镜像
tag=$(docker images | grep redis | grep kytt | wc -l)
if [ $tag = 1 ]; then
	echo "Image Exist"
else
	echo "Image Not Exist"
	echo "Build redis:kytt Image"
	cd ../conf
	docker build -t "redis:kytt" -f Dockerfile.Redis.kytt .
fi

# 清除已经启动的redis_server容器
docker ps -a | grep redis_server
if [ $? = 0 ]; then
	echo "Force Remove redis_server Container"
	docker stop redis_server
	docker rm redis_server
fi

# 启动redis_server容器
cd ../
echo $(pwd)
echo "Start Redis Server Container"
#docker run --name redis_server -v $(pwd)/data/redis_db:/data -v $(pwd)/conf/redis.conf:/usr/local/etc/redis/redis.conf redis:3.2.8 redis-server /usr/local/etc/redis/redis.conf
##docker run --name redis_server -d -v $(pwd)/data/redis_db:/var/lib/reids -v $(pwd)/log:/var/log/redis -v $(pwd)/conf/redis.conf:/usr/local/etc/redis/redis.conf redis:3.0 redis-server /usr/local/etc/redis/redis.conf
