#!/usr/bin/env bash
#go build -o p-vip-gateway ../main.go
CGO_ENABLED=0
GOOS=linux
go build -a -installsuffix cgo -o p-vip-gateway ../main.go

# 登录 docker hub
# docker login -u dauglastang --password taNg##85718

# 构建 docker 镜像并上传到 docker hub
docker build -t p-vip-gateway .
docker tag p-vip-gateway:latest dauglastang/p-vip-gateway:latest
docker push dauglastang/p-vip-gateway:latest