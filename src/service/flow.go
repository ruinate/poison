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
	Execute(mode string, config *conf.FlowModel) *FlowAPP
	AutoExecute(config *conf.FlowModel, payload [][2]interface{}) int
}
type FlowAPP struct {
}

func (f *FlowAPP) Execute(mode string, config *conf.FlowModel) *FlowAPP {
	// 捕获ctrl+c
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	switch mode {
	case "Send":
		client := new(utils.ProtoConfig)
		for {
			time.Sleep(time.Millisecond * 300)
			select {
			case _ = <-Signal:
				logrus.Printf("stopped sending a total of %d packets", TotalPacket)
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
	client := new(utils.ProtoConfig)
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-Signal:
			logrus.Printf("stopped sending a total of %d packets", TotalPacket)
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
