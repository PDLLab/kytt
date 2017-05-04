#!/bin/bash

tag=$(docker images | grep redis | awk '{print $2}')
if [ $tag = "3.0" ]; then
	echo "Image Exist"
else
	echo "Image Not Exist"
	docker pull redis:3.0
fi
