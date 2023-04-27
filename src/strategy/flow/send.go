package flow

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// Send 执行方法
var Send = &cobra.Command{
	Use:   "send firewall",
	Short: "发送数据包：TCP、UDP",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.Check.CheckSend(&utils.Config)
		fmt.Println("Send is running...")
		SEND.Execute(config)
	},
}

// SendAPP 结构体
type SendAPP struct{}

// Execute Send执行方法
func (s *SendAPP) Execute(config *utils.ProtoAPP) SendAPP {
	time.Sleep(time.Millisecond * 300)
	if config.Depth != 0 {
		p, err := config.RUN(config)
		if err != nil {
			utils.LogError(err)
		} else {
			log.Println(p.Result)
		}
		config.Depth -= 1
		if config.Depth == 0 {
			utils.Check.CheckDebug(config.Mode + " Task execution completed......")
		}
		return s.Execute(config)
	} else {
		p, err := config.RUN(config)
		if err != nil {
			utils.LogError(err)
		} else {
			log.Println(p.Result)
		}
		return s.Execute(config)
	}

}

// SEND 实例化SendAPP
var SEND SendAPP
