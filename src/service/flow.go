// Package service -----------------------------
// @file      : flow.go
// @author    : fzf
// @time      : 2023/5/9 下午1:01
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	logger "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	client utils.ProtoConfig
)

type Flow interface {
	RPCExecute() error
	RPC(config *conf.FlowModel, result *string) error
	Execute(mode string, config *conf.FlowModel) *FlowAPP
	AutoExecute(config *conf.FlowModel, payload [][2]interface{}) int
}
type FlowAPP struct {
}

func (f *FlowAPP) RPCExecute() error {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		return err
	}
	_ = rpc.RegisterName("Flow", new(FlowAPP))
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持高并发
	}
}

func (f *FlowAPP) RPC(config *conf.FlowModel, result *string) error {
	p, err := client.Execute(config)
	if err != nil {
		*result = err.Error()
	}
	utils.LogDebug(p, err)
	return nil
}

func (f *FlowAPP) Execute(mode string, config *conf.FlowModel) *FlowAPP {
	// 捕获ctrl+c
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	switch mode {
	case "Send":
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case _ = <-Signal:
				logger.Printf("stopped sending a total of %d packets", TotalPacket)
				os.Exit(0)
			case <-time.After(0 * time.Millisecond):
				p, err := client.Execute(config)
				TotalPacket += 1
				TotalDepth += 1
				utils.LogDebug(p, err)
				utils.Check.CheckDepthSum(TotalDepth, config.Depth, TotalPacket)
			}
		}
	case "Auto":
		payload := utils.Check.CheckAutoMode(config.Mode, config.ICSMode)
		for {
			TotalPacket = f.AutoExecute(config, payload)
			TotalDepth += 1
			utils.Check.CheckDepthSum(TotalDepth, config.Depth, TotalPacket)
		}
	default:
		utils.Check.CheckExit("")
		return nil
	}
}

func (f *FlowAPP) AutoExecute(config *conf.FlowModel, payload [][2]interface{}) int {
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-Signal:
			logger.Printf("stopped sending a total of %d packets", TotalPacket)
			os.Exit(0)
		case <-time.After(0 * time.Millisecond):
			TotalPacket += 1
			config.Port = P[0].(int)
			config.Payload = P[1].(string)
			p, err := client.Execute(config)
			utils.LogDebug(p, err)
		}
	}
	return TotalPacket
}
