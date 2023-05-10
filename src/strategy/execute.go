// Package strategy -----------------------------
// @file      : execute.go
// @author    : fzf
// @time      : 2023/5/9 上午9:25
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/utils"
)

var FlowClient Flow = &FlowAPP{}

type ExecuteInterface interface {
	Send(config *utils.PoisonConfig)
	Auto(config *utils.PoisonConfig)
	Ddos(config *utils.PoisonConfig)
	Server(config *utils.PoisonConfig)
	Snmp(config *utils.PoisonConfig)
}
type Execute struct {
}

func (e *Execute) Send(config *utils.PoisonConfig) {
	FlowClient.Execute("Send", config)
}
func (e *Execute) Auto(config *utils.PoisonConfig) {
	FlowClient.Execute("Auto", config)
}
func (e *Execute) Ddos(config *utils.PoisonConfig) {
	client := new(utils.DdosAPP)
	client.Execute(config)
}
func (e *Execute) Server(config *utils.PoisonConfig) {
	client := new(utils.ServerApp)
	client.Execute(config)
}
func (e *Execute) Snmp(config *utils.PoisonConfig) {
	client := new(utils.SnmpAPP)
	client.Execute(config)
}
