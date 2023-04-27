# GO文件编译优化

## 1、编译命令
    -s：省略符号表和调试信息。 大多数情况下，在生产环境中不需要它们。
    -w: 省略 DWARF 消息。
    go build -ldflags '-w -s'

## 2、UDP优化
    upx 是一个二进制压缩工具。它可用于压缩二进制文件和进一步减少文件大小。
    upx  build_file（文件名称）
## 3、build  windows 包
    GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o mian.exe main.go
