#!/usr/bin/env bash
#go build -o p-vip-gateway ../main.go
echo "构建 go 应用..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o p-vip-gateway ../main.go
# 登录 docker hub

echo "构建 docker 本地镜像..."
# 构建 docker 镜像并上传到 docker hub
docker build -t p-vip-gateway .

export VERSION=1.0
docker tag p-vip-gateway dauglastang/p-vip-gateway:${VERSION}
docker push dauglastang/p-vip-gateway:${VERSION}

echo "构建完成"