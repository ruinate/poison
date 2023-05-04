package flow

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	// Send 执行方法
	Send = &cobra.Command{
		Use:   "send [tab][tab]",
		Short: "发送数据包：TCP、UDP",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckSend(&utils.Config)
			fmt.Println("Send is running...")
			SEND.Execute(config)
		},
	}
)

// SendAPP 结构体
type SendAPP struct{}

// Execute Send执行方法
func (s *SendAPP) Execute(config *utils.ProtoAPP) *SendAPP {
	time.Sleep(time.Millisecond * 300)
	if config.Depth != 0 {
		p, err := config.Execute(config)
		utils.LogDebug(p, err)
		status := utils.Check.CheckDepthSum(config)
		if status {
			return s.Execute(config)
		}
		return nil
	} else {
		p, err := config.Execute(config)
		utils.LogDebug(p, err)
		return s.Execute(config)
	}

}

// SEND 实例化SendAPP
var SEND SendAPP
