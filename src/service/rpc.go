// Package service -----------------------------
// @file      : rpc.go
// @author    : fzf
// @time      : 2023/6/7 下午2:59
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"poison/src/model"
	"poison/src/utils"
)

type RPC struct {
}

func (r *RPC) Execute(config *model.InterfaceModel) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		logger.Fatalln(err)
	}
	_ = rpc.RegisterName("Flow", new(RPC))
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持高并发
	}
}
func (r *RPC) Send(config *model.InterfaceModel) error {
	if err := utils.Client.Execute(config); err != nil {
		return err
	}
	return nil
}
func (r *RPC) Start(config *model.InterfaceModel, result *error) error {
	if config.APPMode == "Send" {
		if err := r.Send(config); err != nil {
			*result = err
			return err
		}
	}
	return nil
}
