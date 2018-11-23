#!/bin/bash

#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 godep go build
# 不需要在外面编译好，再将可执行文件拷贝到镜像里了
# /usr/local/go/pkg/tool/linux_amd64/link: signal: killed 的问题解决了，是由于 -m 64m 导致的，分配的内存太小
# 现在使用256m完美解决，若再出现这个问题，可以将分配的内存调大一些，一般建议512m以上

TAG=`date "+%Y%m%d-%H%M%S"`
IMAGE=bee-web

docker build -f ./Dockerfile -t $IMAGE:$TAG .
docker tag $IMAGE:$TAG $IMAGE
