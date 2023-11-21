// Package service -----------------------------
// @file      : flow.go
// @author    : fzf
// @time      : 2023/5/9 下午1:01
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"poison/src/model"
	"poison/src/service/setting"
	"poison/src/utils"
	"time"
)

type AutoStruct struct {
}

func (f *AutoStruct) Execute(config *model.InterfaceModel) {
	payload := utils.Output.Execute(config.Mode, config.ICSMode)
	for {
		for _, P := range payload {
			time.Sleep(time.Millisecond * 300)
			select {
			case _ = <-setting.Signal:
				logger.Printf("stopped sending a total of %d packets", setting.TotalPacket)
				os.Exit(0)
			case <-time.After(0 * time.Millisecond):
				setting.TotalPacket += 1
				config.DstPort = P[0].(int)
				config.Payload = P[1].(string)
				config.SrcPort = 0
				if err := utils.Client.Execute(config); err != nil {
					logger.Errorln(err)
				}
			}
		}
		if setting.TotalDepth++; setting.TotalDepth == config.Depth {
			logger.Fatalln("stopped sending a total of %d packets", setting.TotalPacket)
		}
	}

}
