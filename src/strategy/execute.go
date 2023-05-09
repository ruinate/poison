// Package strategy -----------------------------
// @file      : execute.go
// @author    : fzf
// @time      : 2023/5/9 上午9:25
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/common"
	"PoisonFlow/src/utils"
)

var FlowClient Flow = &FlowAPP{}

type ExecuteInterface interface {
	Send(config *common.ConfigType)
	Auto(config *common.ConfigType)
	Ddos(config *common.ConfigType)
	Server(config *common.ConfigType)
	Snmp(config *common.ConfigType)
}
type Execute struct {
}

func (e *Execute) Send(config *common.ConfigType) {
	FlowClient.Execute("Send", config)
}
func (e *Execute) Auto(config *common.ConfigType) {
	FlowClient.Execute("Auto", config)
}
func (e *Execute) Ddos(config *common.ConfigType) {
	client := new(utils.DdosAPP)
	client.Execute(config)
}
func (e *Execute) Server(config *common.ConfigType) {
	client := new(utils.ServerApp)
	client.Execute(config)
}
func (e *Execute) Snmp(config *common.ConfigType) {
	client := new(utils.SnmpAPP)
	client.Execute(config)
}
