// Package strategy -----------------------------
// @file      : execute.go
// @author    : fzf
// @time      : 2023/5/9 上午9:25
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/service"
)

var FlowClient service.Flow = &service.FlowAPP{}
var ReplayClient service.ReplayInterFace = &service.Replay{}

type ExecuteInterface interface {
	Send(config *conf.FlowModel)
	Auto(config *conf.FlowModel)
	Ddos(config *conf.FlowModel)
	Server(config *conf.FlowModel)
	Snmp(config *conf.FlowModel)
	Replay(config *conf.ReplayModel)
	RPC() error
}
type Execute struct {
}

func (e *Execute) Send(config *conf.FlowModel) {
	FlowClient.Execute("Send", config)
}
func (e *Execute) Auto(config *conf.FlowModel) {
	FlowClient.Execute("Auto", config)
}
func (e *Execute) Ddos(config *conf.FlowModel) {
	client := new(service.DdosAPP)
	go service.DDosSpeed()
	client.Execute(config)
}
func (e *Execute) Server(config *conf.FlowModel) {
	client := new(service.ServerApp)
	client.Execute(config)
}
func (e *Execute) Snmp(config *conf.FlowModel) {
	client := new(service.SnmpAPP)
	client.Execute(config)
}
func (e *Execute) Replay(config *conf.ReplayModel) {
	ReplayClient.Execute(config)
}

func (e *Execute) RPC() error {
	client := new(service.RPCModel)
	err := client.Execute()
	return err
}
