// Package strategy -----------------------------
// @file      : flow.go
// @author    : fzf
// @time      : 2023/5/9 下午1:01
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Flow interface {
	Execute(mode string, config *utils.PoisonConfig) *FlowAPP
	AutoExecute(payload [][2]interface{})
}
type FlowAPP struct {
}

var (
	CountPacket = 0
	c           = make(chan os.Signal, 1)
)

func (f *FlowAPP) Execute(mode string, config *utils.PoisonConfig) *FlowAPP {
	// 捕获ctrl+c
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	switch mode {
	case "Send":
		client := new(utils.ProtoConfig)
		if config.Depth != 0 {
			for {
				time.Sleep(time.Millisecond * 300)
				select {
				case _ = <-c:
					logrus.Printf("stopped sending a total of %d packets", CountPacket)
					os.Exit(0)
				default:
					CountPacket++
					p, err := client.Execute(config)
					utils.LogDebug(p, err)
					status := utils.Check.CheckDepthSum(config)
					if status == false {
						logrus.Printf("stopped sending a total of %d packets", CountPacket)
						return nil
					}
				}

			}

		} else {

			for {
				time.Sleep(time.Millisecond * 300)
				select {
				case _ = <-c:
					logrus.Printf("stopped sending a total of %d packets", CountPacket)
					os.Exit(0)
				default:
					CountPacket++
					p, err := client.Execute(config)
					utils.LogDebug(p, err)
				}
			}
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
	client := new(utils.ProtoConfig)
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-c:
			logrus.Printf("stopped sending a total of %d packets", CountPacket)
			os.Exit(0)
		default:
			utils.Config.Port = P[0].(int)
			utils.Config.Payload = P[1].(string)
			p, err := client.Execute(&utils.Config)
			CountPacket += 2
			utils.LogDebug(p, err)
		}

	}
}
