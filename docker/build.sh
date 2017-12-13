#!/usr/bin/env bash
#go build -o p-vip-gateway ../main.go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o p-vip-gateway ../main.go