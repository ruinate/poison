// Package strategy -----------------------------
// @file      : commands.go
// @author    : fzf
// @time      : 2023/5/9 上午9:55
// -------------------------------------------
package strategy

import (
	"PoisonFlow/src/common"
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ExecuteAPP ExecuteInterface = &Execute{}
	// Send 执行方法
	Send = &cobra.Command{
		Use:   "send ",
		Short: "发送数据包：TCP、UDP",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckSend(&common.Config)
			logrus.Infof("Starting  Send Mode %s ...\n", config.Mode)
			ExecuteAPP.Send(config)
		},
	}
	// Auto 执行命令
	Auto = &cobra.Command{
		Use:   "auto ",
		Short: "自动发送：TCP、UDP、BLACK、ICS",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckAuto(&common.Config)
			logrus.Infof("Starting Auto Mode %s ...\n", config.Mode)
			ExecuteAPP.Auto(config)
		},
	}
	// Snmp 执行方法
	Snmp = &cobra.Command{
		Use:   "snmp ",
		Short: "SNMP 客户端连接测试",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckSnmp(&common.Config)
			logrus.Infof("Starting  SNMP Host : %s ...\n", config.Host)
			ExecuteAPP.Snmp(config)
		},
	}
	// Server 执行方法
	Server = &cobra.Command{
		Use:   "server ",
		Short: "服务端：监听端口默认全部",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckServer(&common.Config)
			logrus.Infof("Starting server Host : %s  Mode : %s", config.Host, config.Mode)
			ExecuteAPP.Server(config)
		},
	}
	// DDOS 执行方法
	DDOS = &cobra.Command{
		Use:   "ddos ",
		Short: "安全防护",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckDDos(&common.Config)
			logrus.Println("Starting Target IP: " + config.Host)
			ExecuteAPP.Ddos(config)
		},
	}
)
