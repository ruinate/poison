package config

import (
	"PoisonFlow/src/common"
	"PoisonFlow/src/strategy"
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// 命令行提示
	validArgs = []string{"send", "auto", "ddos"}
	// Poison 总开关
	Poison = &cobra.Command{
		Use:       "poison [command] [tab][tab]",
		Short:     "Display one or many resources",
		Long:      ``,
		ValidArgs: validArgs,
	}
	n string
)

func init() {
	fmt.Println(
		`
				 _
	_ __   ___ (_)  ___  ___  _ __
	| '_ \ / _ \| | / __|/ _ \| '_ \
	| |_) | (_) | | \__ \ (_) | | | |
	| .__/ \___/|_| |___/\___/|_| |_|
	|_|
	`,
	)
	Poison.AddCommand(strategy.CompletionCmd, strategy.Snmp, strategy.Server, strategy.Auto, strategy.Send, strategy.DDOS)
	Poison.PersistentFlags().StringVarP(&n, "none", "n", "text", "send: 基础发送	auto: 自动发送	hping: 安全防护流量 \n"+
		"snmp：snmp客户端	server: 服务端")
	// Send flags
	strategy.Send.Flags().StringVarP(&common.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP")
	strategy.Send.Flags().StringVarP(&common.Config.Host, "host", "H", "0.0.0.0", "Host载体")
	strategy.Send.Flags().StringVarP(&common.Config.Payload, "payload", "p", utils.RandStr(10), "数据载体")
	strategy.Send.Flags().IntVarP(&common.Config.Port, "port", "P", 22, "端口载体")
	strategy.Send.Flags().IntVarP(&common.Config.Depth, "depth", "d", 1, "循环载体")

	// Auto flags
	strategy.Auto.Flags().StringVarP(&common.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICS、BLACK")
	strategy.Auto.Flags().StringVarP(&common.Config.Host, "host", "H", "0.0.0.0", "Host载体")
	strategy.Auto.Flags().IntVarP(&common.Config.Depth, "depth", "d", 1, "循环载体")
	// DDos flags
	strategy.DDOS.Flags().StringVarP(&common.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICMP、WinNuke、Smurf:广播攻击\n"+
		"'Land、TearDrop、MAXICMP ，默认：TCP'")
	strategy.DDOS.Flags().StringVarP(&common.Config.Host, "host", "H", "0.0.0.0", "Host载体")
	strategy.DDOS.Flags().IntVarP(&common.Config.Port, "port", "P", 10086, "端口载体")
	// Server flags
	strategy.Server.Flags().StringVarP(&common.Config.Host, "host", "H", "0.0.0.0", "Host载体")
	strategy.Server.Flags().StringVarP(&common.Config.Mode, "mode", "m", "TCP", "模式载体")
	// Snmp flags
	strategy.Snmp.Flags().StringVarP(&common.Config.Host, "host", "H", "0.0.0.0", "Host载体")
}
