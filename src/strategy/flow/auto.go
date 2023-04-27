package flow

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"sync"
	"time"
)

// wg  协程锁
var wg sync.WaitGroup

// Auto 执行命令
var Auto = &cobra.Command{
	Use:   "auto firewall",
	Short: "自动发送：TCP、UDP、BLACK、ICS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		payload := utils.Check.CheckAuto(&utils.Config)
		fmt.Printf("Auto  Mode %s is running...\n", utils.Config.Mode)
		wg.Add(1)
		go AUTO.Execute(payload)
		wg.Wait()
	},
}

type auto struct {
}

func (a *auto) AutoExecute(config *utils.ProtoAPP, payload [][2]interface{}) auto {
	time.Sleep(time.Millisecond * 300)
	if config.Depth != 0 {
		a.Execute(payload)
		if config.Depth == 0 {
			utils.Check.CheckDebug(config.Mode + " Task execution completed......")
		}
		return a.Execute(payload)
	} else {
		a.Execute(payload)
		return a.AutoExecute(config, payload)

	}
}

// Execute 执行方法
func (a *auto) Execute(payload [][2]interface{}) auto {
	defer wg.Done()
	config := utils.ProtoAPP{
		Host:  utils.Config.Host,
		Depth: utils.Config.Depth,
		Mode:  utils.Config.Mode,
	}
	for _, P := range payload {
		time.Sleep(time.Millisecond * 300)
		config.Port = P[0].(int)
		config.Payload = P[1].(string)
		p, err := config.RUN(&config)
		if err != nil {
			utils.LogError(err)
		} else {
			log.Println(p.Result)
		}

	}
	return AUTO

}

// AUTO 实例化auto
var AUTO auto
