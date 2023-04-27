#!/bin/bash
echo build go
go build -ldflags '-w -s' -o poison src/main.go

echo upx 压缩build文件

upx  poison

echo  success posion
