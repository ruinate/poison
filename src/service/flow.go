// Package service -----------------------------
// @file      : flow.go
// @author    : fzf
// @time      : 2023/5/9 下午1:01
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Flow interface {
	Execute(mode string, config *conf.PoisonConfig) *FlowAPP
	AutoExecute(payload [][2]interface{})
}
type FlowAPP struct {
}

func (f *FlowAPP) Execute(mode string, config *conf.PoisonConfig) *FlowAPP {
	// 捕获ctrl+c
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	switch mode {
	case "Send":
		client := new(utils.ProtoConfig)
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case _ = <-Signal:
				logrus.Printf("stopped sending a total of %d packets", CounterPacket)
				os.Exit(0)
			case <-time.After(0 * time.Millisecond):
				p, err := client.Execute(config)
				CounterPacket += 1
				CounterDepth += 1
				utils.LogDebug(p, err)
				status := utils.Check.CheckDepthSum(CounterDepth, config.Depth)
				if status != true {
					logrus.Printf("stopped sending a total of %d packets", CounterPacket)
					return nil
				}
			}

		}
	case "Auto":
		payload := utils.Check.CheckAutoMode(config.Mode)
		for {
			f.AutoExecute(payload)
			CounterDepth += 1
			status := utils.Check.CheckDepthSum(CounterDepth, config.Depth)
			if status != true {
				logrus.Printf("stopped sending a total of %d packets", CounterPacket)
				return nil
			}
		}
	default:
		utils.Check.CheckExit("")
		return nil
	}
}

func (f *FlowAPP) AutoExecute(payload [][2]interface{}) {
	client := new(utils.ProtoConfig)
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-Signal:
			logrus.Printf("stopped sending a total of %d packets", CounterPacket)
			os.Exit(0)
		case <-time.After(0 * time.Millisecond):
			conf.Config.Port = P[0].(int)
			conf.Config.Payload = P[1].(string)
			p, err := client.Execute(&conf.Config)
			CounterPacket += 2
			utils.LogDebug(p, err)
		}

	}
}
