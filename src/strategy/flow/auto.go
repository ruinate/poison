package flow

import (
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var (
	// Auto 执行命令
	Auto = &cobra.Command{
		Use:   "auto [tab][tab]",
		Short: "自动发送：TCP、UDP、BLACK、ICS",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			payload := utils.Check.CheckAuto(&utils.Config)
			logrus.Infof("Auto  Mode %s is running...\n", utils.Config.Mode)
			AUTO.AutoExecute(&utils.Config, payload)
		},
	}
)

type auto struct {
}

func (a *auto) AutoExecute(config *utils.ProtoAPP, payload [][2]interface{}) *auto {
	time.Sleep(time.Millisecond * 200)
	if config.Depth != 0 {
		a.Execute(payload)
		status := utils.Check.CheckDepthSum(config)
		if status {
			return a.AutoExecute(config, payload)
		}
		return nil
	} else {
		a.Execute(payload)
		return a.AutoExecute(config, payload)
	}
}

// Execute 执行方法
func (a *auto) Execute(payload [][2]interface{}) auto {
	config := &utils.ProtoAPP{
		Host:  utils.Config.Host,
		Depth: utils.Config.Depth,
		Mode:  utils.Config.Mode,
	}
	for _, P := range payload {
		time.Sleep(time.Millisecond * 100)
		config.Port = P[0].(int)
		config.Payload = P[1].(string)
		p, err := config.Execute(config)
		utils.LogDebug(p, err)
	}
	return AUTO
}

// AUTO 实例化auto
var AUTO auto
