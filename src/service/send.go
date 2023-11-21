// Package service -----------------------------
// @file      : send.go
// @author    : fzf
// @time      : 2023/11/20 上午10:36
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"os/signal"
	"poison/src/model"
	"poison/src/service/setting"
	"poison/src/utils"
	"syscall"
	"time"
)

type SendStruct struct{}

func (s *SendStruct) Execute(config *model.InterfaceModel) {
	signal.Notify(setting.Signal, syscall.SIGINT, syscall.SIGTERM)
	for {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-setting.Signal:
			logger.Fatalln("stopped sending a total of %d packets", setting.TotalPacket)
		case <-time.After(0 * time.Millisecond):
			setting.TotalPacket += 1
			setting.TotalDepth += 1
			if err := utils.Client.Execute(config); err != nil {
				logger.Errorln(err)
			}
			if setting.TotalDepth == config.Depth {
				logger.Fatalln("stopped sending a total of %d packets", setting.TotalPacket)
			}
		}
	}
}
