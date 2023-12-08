// Package service -----------------------------
// @file      : server.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:23
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"poison/src/core/server"
	"poison/src/model"
	"poison/src/utils"
	"strconv"
)

type ServerCmd struct {
	cmd *cobra.Command
}

func (s *ServerCmd) InitCmd() *cobra.Command {
	s.cmd = &cobra.Command{
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
			s.Execute(&model.Config)
		},
	}
	s.cmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "0.0.0.0", "Host载体")
	s.cmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体")
	s.cmd.Flags().IntVarP(&model.Config.StartPort, "srcport", "s", 1, "监听开始端口")
	s.cmd.Flags().IntVarP(&model.Config.EndPort, "dstport", "e", 65535, "监听结束端口")
	return s.cmd
}

// Execute 监听执行
func (s *ServerCmd) Execute(config *model.Stream) {
	for port := config.StartPort; port < config.EndPort; {
		port++
		switch config.Mode {
		case "TCP":
			go server.ListenerTCP(config.DstHost, strconv.Itoa(port))
		case "UDP":
			go server.ListenerUDP(port)
		}
	}
	ERRORNUMBER := len(server.SERVERPORTERROR)
	if ERRORNUMBER > 15 {
		logger.Errorf("监听失败端口数量：%d个端口", ERRORNUMBER)
	} else {
		logger.Errorf("监听失败端口：%s", server.SERVERPORTERROR)
	}
	select {}
}
