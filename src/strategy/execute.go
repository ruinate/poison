// Package strategy -----------------------------
// @file      : execute.go
// @author    : fzf
// @time      : 2023/5/9 上午9:25
// -------------------------------------------
package strategy

import (
	"poison/src/model"
	"poison/src/service"
)

type Factory interface {
	Execute(config *model.InterfaceModel)
}

func NewClient(config *model.InterfaceModel) Factory {
	switch config.APPMode {
	case model.SEND:
		return &service.SendStruct{}
	case model.AUTO:
		return &service.AutoStruct{}
	case model.DDOS:
		return &service.DDOSStruct{}
	case model.REPLAY:
		return &service.ReplayStruct{}
	case model.RPC:
		return &service.RPC{}
	case model.SERVER:
		return &service.ServerStruct{}
	case model.SNMP:
		return &service.SnmpStruct{}
	case model.PING:
		return &service.PingStruct{}
	default:
		return nil
	}
}
