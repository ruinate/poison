// Package strategy -----------------------------
// @file      : commands.go
// @author    : fzf
// @time      : 2023/5/9 上午9:55
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	FlowConfig   conf.FlowModel
	ReplayConfig conf.ReplayModel
	// 命令行提示
	validArgs = []string{"send", "auto", "ddos"}
	// Poison 总开关
	Poison = &cobra.Command{
		Use:       "poison [command] [tab][tab]",
		Short:     "Display one or many resources",
		Long:      ``,
		ValidArgs: validArgs,
	}
	n          string
	ExecuteAPP ExecuteInterface = &Execute{}
	// Send 执行方法
	Send = &cobra.Command{
		Use:   "send ",
		Short: "发送数据包：TCP、UDP",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckSend(&FlowConfig)
			logrus.Infof("Starting  Send Mode %s ...\n", Config.Mode)
			ExecuteAPP.Send(Config)
		},
	}
	// Auto 执行命令
	Auto = &cobra.Command{
		Use:   "auto ",
		Short: "自动发送：TCP、UDP、BLACK、ICS",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckAuto(&FlowConfig)
			logrus.Infof("Starting Auto Mode %s ...\n", Config.Mode)
			ExecuteAPP.Auto(Config)
		},
	}
	// Snmp 执行方法
	Snmp = &cobra.Command{
		Use:   "snmp ",
		Short: "SNMP 客户端连接测试",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckSnmp(&FlowConfig)
			logrus.Infof("Starting  Host : %s ...\n", Config.Host)
			ExecuteAPP.Snmp(Config)
		},
	}
	// Server 执行方法
	Server = &cobra.Command{
		Use:   "server ",
		Short: "服务端：监听端口默认全部",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckServer(&FlowConfig)
			logrus.Infof("Starting server Host : %s  Mode : %s...\n", Config.Host, Config.Mode)
			ExecuteAPP.Server(Config)
		},
	}
	// DDOS 执行方法
	DDOS = &cobra.Command{
		Use:   "ddos ",
		Short: "安全防护",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckDDos(&FlowConfig)
			logrus.Printf("Starting  Host:%s  Mode:%s ...\n", Config.Host, Config.Mode)
			ExecuteAPP.Ddos(Config)
		},
	}
	// Replay 执行方法
	Replay = &cobra.Command{
		Use:   "replay ",
		Short: "流量重放",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			Config := utils.Check.CheckReplay(&ReplayConfig)
			logrus.Printf("Starting Interface :%s   path :%s...\n", Config.InterFace, Config.FilePath)
			ExecuteAPP.Replay(&ReplayConfig)
		},
	}
)

func init() {
	//fmt.Println(
	//	`
	//			 _
	//_ __   ___ (_)  ___  ___  _ __
	//| '_ \ / _ \| | / __|/ _ \| '_ \
	//| |_) | (_) | | \__ \ (_) | | | |
	//| .__/ \___/|_| |___/\___/|_| |_|
	//|_|
	//`,
	//)
	Poison.AddCommand(CompletionCmd, Snmp, Server, Auto, Send, DDOS, Replay)
	Poison.PersistentFlags().StringVarP(&n, "none", "n", "text", "send: 基础发送	auto: 自动发送	hping: 安全防护流量 \n"+
		"snmp：snmp客户端	server: 服务端")
	// Send flags
	Send.Flags().StringVarP(&FlowConfig.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP")
	Send.Flags().StringVarP(&FlowConfig.Host, "host", "H", "0.0.0.0", "Host载体")
	Send.Flags().StringVarP(&FlowConfig.Payload, "payload", "p", utils.RandStr(20), "数据载体")
	Send.Flags().IntVarP(&FlowConfig.Port, "port", "P", 22, "端口载体")
	Send.Flags().IntVarP(&FlowConfig.Depth, "depth", "d", 1, "循环载体")
	// Auto flags
	Auto.Flags().StringVarP(&FlowConfig.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICS、BLACK")
	Auto.Flags().StringVarP(&FlowConfig.Host, "host", "H", "0.0.0.0", "Host载体")
	Auto.Flags().IntVarP(&FlowConfig.Depth, "depth", "d", 1, "循环载体")
	// DDos flags
	DDOS.Flags().StringVarP(&FlowConfig.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICMP、WinNuke、Smurf:广播攻击\n"+
		"'Land、TearDrop、MAXICMP ，默认：TCP'")
	DDOS.Flags().StringVarP(&FlowConfig.Host, "host", "H", "0.0.0.0", "Host载体")
	DDOS.Flags().IntVarP(&FlowConfig.Port, "port", "P", 10086, "端口载体")
	// Server flags
	Server.Flags().StringVarP(&FlowConfig.Host, "host", "H", "0.0.0.0", "Host载体")
	Server.Flags().StringVarP(&FlowConfig.Mode, "mode", "m", "TCP", "模式载体")
	// Snmp flags
	Snmp.Flags().StringVarP(&FlowConfig.Host, "host", "H", "0.0.0.0", "Host载体")
	// Replay flags
	Replay.Flags().StringVarP(&ReplayConfig.InterFace, "interface", "i", "lo", "接口载体")
	Replay.Flags().IntVarP(&ReplayConfig.Speed, "speed", "s", 100000, "速度载体")
	Replay.Flags().StringVarP(&ReplayConfig.FilePath, "file", "f", "", "路径载体")
	Replay.Flags().IntVarP(&ReplayConfig.Depth, "depth", "d", 1, "循环载体")
	// Flag TAB
	flagName := "mode"
	err := Send.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP"}, cobra.ShellCompDirectiveDefault
	})
	err = Auto.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP", "ICS", "BLACK"}, cobra.ShellCompDirectiveDefault
	})
	err = DDOS.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP", "ICMP", "WinNuke", "Smurf"}, cobra.ShellCompDirectiveDefault
	})
	err = Server.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP"}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		return
	}
}
