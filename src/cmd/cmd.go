// Package cmd  -----------------------------
// @file      : commands.go
// @author    : fzf
// @time      : 2023/5/9 上午9:55
// -------------------------------------------
package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"poison/src/model"
	"poison/src/strategy"
	"poison/src/utils"
)

var (

	// PoisonCmd  总开关
	PoisonCmd = &cobra.Command{
		Use:       "poison",
		Short:     "Display one or many resources",
		Long:      ``,
		ValidArgs: []string{model.SEND, model.AUTO, model.SNMP, model.SERVER, model.REPLAY, model.DDOS, model.PING, model.RPC},
		Args:      cobra.OnlyValidArgs,
	}
	SendCmd = &cobra.Command{
		Use:       model.SEND,
		Short:     "发送数据包：TCP、UDP",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-m", "-H", "-P", "-p", "-d"},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("Starting  Send Mode %s ...\n", model.Config.APPMode)
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
	AutoCmd = &cobra.Command{
		Use:       model.AUTO,
		Short:     "自动发送：TCP、UDP、BLACK、ICS",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-m", "-H", "-d"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Infof("Starting Auto Mode %s ...\n", model.Config.Mode)
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
	SnmpCmd = &cobra.Command{
		Use:       model.SNMP,
		Short:     "SNMP 客户端连接测试",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-H"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Infof("Starting  Host : %s ...\n", model.Config.DstHost)
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}

	ServerCmd = &cobra.Command{
		Use:       model.SERVER,
		Short:     "服务端：监听端口默认全部",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-m", "-H"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Infof("Starting server Host : %s  Mode : %s...\n", model.Config.DstHost, model.Config.Mode)
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
	DDOSCmd = &cobra.Command{
		Use:       model.DDOS,
		Short:     "安全防护",
		Long:      ``,
		ValidArgs: []string{"-m", "-H", "-P"},
		Args:      cobra.OnlyValidArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Printf("Starting  Host:%s  Mode:%s ...\n", model.Config.DstHost, model.Config.Mode)
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}

	ReplayCmd = &cobra.Command{
		Use:       model.REPLAY,
		Short:     "流量重放",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-i", "-f", "-s"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Printf("Starting Interface :%s   path :%s...\n", model.Config.InterFace, model.Config.FilePath)
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
	RPCCmd = &cobra.Command{
		Use:   model.RPC,
		Short: "rpc服务器",
		Long:  ``,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Println("Starting RPC_SERVER...")
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
	PingCmd = &cobra.Command{
		Use:   model.PING,
		Short: "ping",
		Long:  ``,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Println("Starting Ping...")
			client := strategy.NewClient(&model.Config)
			client.Execute(&model.Config)
		},
	}
)

func init() {

	if len(os.Args) >= 2 {
		model.Config.APPMode = os.Args[1]
	}
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
	PoisonCmd.AddCommand(ServerCmd, SnmpCmd, SendCmd, PingCmd, AutoCmd, ReplayCmd, RPCCmd, DDOSCmd, CompletionCmd)
	SendCmd.Flags().StringVarP(&model.Config.DstMAC, "dstmac", "c", "00:11:22:33:44:55", "目的MAC地址")
	SendCmd.Flags().StringVarP(&model.Config.SrcMAC, "srcmac", "C", "66:77:88:99:aa:bb", "源目的MAC地址")
	SendCmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP")
	SendCmd.Flags().StringVarP(&model.Config.DstHost, "dsthost", "H", "127.0.0.1", "目的地址")
	SendCmd.Flags().StringVarP(&model.Config.SrcHost, "srchost", "S", "0.0.0.0", "源目的地址")
	SendCmd.Flags().StringVarP(&model.Config.Payload, "payload", "p", utils.RandStr(20), "数据载体")
	SendCmd.Flags().IntVarP(&model.Config.SrcPort, "sport", "s", 0, "源端口")
	SendCmd.Flags().IntVarP(&model.Config.DstPort, "port", "P", 22, "目的端口")
	SendCmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	SendCmd.Flags().StringVarP(&model.Config.SendMode, "sendmode", "M", "ROUTE", "发送模式：MAC和ROUTE")
	SendCmd.Flags().StringVarP(&model.Config.InterFace, "interface", "i", "enp2s0", "网卡名称")
	// Auto flags
	AutoCmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICS、BLACK")
	AutoCmd.Flags().StringVarP(&model.Config.DstHost, "dsthost", "H", "127.0.0.1", "Host载体")
	AutoCmd.Flags().StringVarP(&model.Config.SrcHost, "srchost", "S", "0.0.0.0", "Host载体")
	AutoCmd.Flags().IntVarP(&model.Config.SrcPort, "sport", "s", 0, "源端口")
	AutoCmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	AutoCmd.Flags().StringVarP(&model.Config.ICSMode, "icsmode", "i", "all", "ICS模式选择")
	// DDos flags
	DDOSCmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICMP、WinNuke、Smurf"+
		"、Land、TearDrop、MAXICMP")
	DDOSCmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "0.0.0.0", "Host载体")
	DDOSCmd.Flags().IntVarP(&model.Config.DstPort, "port", "P", 10086, "端口载体")
	// Server flags
	ServerCmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "0.0.0.0", "Host载体")
	ServerCmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体")
	ServerCmd.Flags().IntVarP(&model.Config.StartPort, "srcport", "S", 1, "监听开始端口")
	ServerCmd.Flags().IntVarP(&model.Config.EndPort, "dstport", "E", 65535, "监听结束端口")
	// Snmp flags
	SnmpCmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "0.0.0.0", "Host载体")

	// Snmp flags
	PingCmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "0.0.0.0", "Host载体")
	PingCmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 4, "循环载体")
	// Replay flags
	ReplayCmd.Flags().StringVarP(&model.Config.InterFace, "interface", "i", "lo", "接口载体")
	ReplayCmd.Flags().IntVarP(&model.Config.Speed, "speed", "s", 100000, "速度载体")
	ReplayCmd.Flags().StringVarP(&model.Config.FilePath, "file", "f", "", "路径载体")
	ReplayCmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	// Flag TAB
	inter := utils.TotalDevice()
	// send
	err := SendCmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP"}, cobra.ShellCompDirectiveDefault
	})
	err = SendCmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = SendCmd.RegisterFlagCompletionFunc(model.PORT, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = SendCmd.RegisterFlagCompletionFunc(model.PAYLOAD, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = SendCmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})

	err = SendCmd.RegisterFlagCompletionFunc(model.REPLAYINTERFACE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inter, cobra.ShellCompDirectiveDefault
	})

	err = SendCmd.RegisterFlagCompletionFunc(model.ROUTE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inter, cobra.ShellCompDirectiveDefault
	})

	// auto
	err = AutoCmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP", "ICS", "BLACK"}, cobra.ShellCompDirectiveDefault
	})
	err = AutoCmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = AutoCmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = AutoCmd.RegisterFlagCompletionFunc(model.ICSMODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return model.PROTOICSMODE, cobra.ShellCompDirectiveDefault
	})
	// ddos
	err = DDOSCmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return model.PROTOMODE, cobra.ShellCompDirectiveDefault
	})
	err = DDOSCmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = DDOSCmd.RegisterFlagCompletionFunc(model.PORT, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	// server
	err = ServerCmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP"}, cobra.ShellCompDirectiveDefault
	})
	err = ServerCmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})

	// snmp
	err = SnmpCmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	// replay

	err = ReplayCmd.RegisterFlagCompletionFunc(model.REPLAYINTERFACE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inter, cobra.ShellCompDirectiveDefault
	})
	err = ReplayCmd.RegisterFlagCompletionFunc(model.REPLAYFILE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = ReplayCmd.RegisterFlagCompletionFunc(model.SPEED, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = ReplayCmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		logger.Fatalln(err)
	}
}
