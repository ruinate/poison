// Package strategy -----------------------------
// @file      : flow.go
// @author    : fzf
// @time      : 2023/5/9 下午1:01
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/common"
	"PoisonFlow/src/utils"
	"time"
)

type Flow interface {
	Execute(mode string, config *common.ConfigType) *FlowAPP
	AutoExecute(payload [][2]interface{})
}
type FlowAPP struct {
}

func (f *FlowAPP) Execute(mode string, config *common.ConfigType) *FlowAPP {
	switch mode {
	case "Send":
		client := new(utils.ProtoAPP)
		time.Sleep(time.Millisecond * 300)
		if config.Depth != 0 {
			p, err := client.Execute(config)
			utils.LogDebug(p, err)
			status := utils.Check.CheckDepthSum(config)
			if status {
				return f.Execute(mode, config)
			}
		} else {
			p, err := client.Execute(config)
			utils.LogDebug(p, err)
			return f.Execute(mode, config)
		}
	case "Auto":
		payload := utils.Check.CheckAutoMode(config.Mode)
		if config.Depth != 0 {
			f.AutoExecute(payload)
			status := utils.Check.CheckDepthSum(config)
			if status {
				return f.Execute(mode, config)
			}

		} else {
			f.AutoExecute(payload)
			return f.Execute(mode, config)
		}
	default:
		utils.Check.CheckExit("")
		return nil
	}
	return nil
}

func (f *FlowAPP) AutoExecute(payload [][2]interface{}) {
	client := new(utils.ProtoAPP)
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		common.Config.Port = P[0].(int)
		common.Config.Payload = P[1].(string)
		p, err := client.Execute(&common.Config)
		utils.LogDebug(p, err)
	}
}
