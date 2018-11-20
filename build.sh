#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 godep go build

docker build -m 64m -f ./Dockerfile -t bee-web .
