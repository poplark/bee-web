#!/bin/bash

IMAGE=bee-web
APP=bee-web

docker stop $APP
docker rm $APP
docker run --name $APP -p 127.0.0.1:3001:3001 -d $IMAGE
