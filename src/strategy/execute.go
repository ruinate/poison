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
	Send(config *conf.PoisonConfig)
	Auto(config *conf.PoisonConfig)
	Ddos(config *conf.PoisonConfig)
	Server(config *conf.PoisonConfig)
	Snmp(config *conf.PoisonConfig)
	Replay(config *conf.PoisonConfig)
}
type Execute struct {
}

func (e *Execute) Send(config *conf.PoisonConfig) {
	FlowClient.Execute("Send", config)
}
func (e *Execute) Auto(config *conf.PoisonConfig) {
	FlowClient.Execute("Auto", config)
}
func (e *Execute) Ddos(config *conf.PoisonConfig) {
	client := new(service.DdosAPP)
	client.Execute(config)
}
func (e *Execute) Server(config *conf.PoisonConfig) {
	client := new(service.ServerApp)
	client.Execute(config)
}
func (e *Execute) Snmp(config *conf.PoisonConfig) {
	client := new(service.SnmpAPP)
	client.Execute(config)
}
func (e *Execute) Replay(config *conf.PoisonConfig) {
	ReplayClient.Execute(config)
}
