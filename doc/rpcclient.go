// Package main -----------------------------
// @file      : rpcclient.go
// @author    : fzf
// @time      : 2023/6/5 上午9:37
// -------------------------------------------
package main

import (
	"PoisonFlow/src/conf"
	logger "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var (
	config = conf.FlowModel{
		Depth:   1,
		Mode:    "TCP",
		Host:    "10.30.5.103",
		Port:    10086,
		Payload: "aqwert",
	}
	result *string = new(string)
)

func main() {
	// 这里不能再用rpc做连接，因为rpc内部会用Gob协议
	conn, err := net.Dial("tcp", "10.30.1.127:1234")
	if err != nil {
		panic("connection failed")
	}
	// 这里指定序列化协议为JSON
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("Flow.RPC", &config, result)
	if err != nil {
		panic("调用失败")
	}
	logger.Printf("RPC函数 调用成功")
	logger.Println(*result)
}
