// Package service -----------------------------
// @file      : rpc.go
// @author    : fzf
// @time      : 2023/6/7 下午2:59
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RPCModel struct {
	stopChan chan bool
}

func (r *RPCModel) Execute() error {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	_ = rpc.RegisterName("Flow", new(RPCModel))
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持高并发
	}
}
func (r *RPCModel) Send(config *conf.FlowModel) error {
	err := utils.Check.CheckSend(config)
	if err != nil {
		return err
	}
	p, err := client.Execute(config)
	utils.LogDebug(p, err)
	if err != nil {
		return err
	}
	return nil
}
func (r *RPCModel) Start(config *conf.FlowModel, result *error) error {
	if config.APPMode == "Send" {
		err := r.Send(config)
		if err != nil {
			*result = err
			return err
		}
	}
	return nil
}
