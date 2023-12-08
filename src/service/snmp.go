// Package service -----------------------------
// @file      : snmp.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:23
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"poison/src/core/snmp"
	"poison/src/model"
	"poison/src/utils"
	"time"
)

type SnmpCmd struct {
	cmd *cobra.Command
}

var SNMPVersion = [...]string{"v1", "v2", "v3"}

func (s *SnmpCmd) InitCmd() *cobra.Command {
	s.cmd = &cobra.Command{
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
			s.Execute(&model.Config)
		},
	}
	s.cmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "127.0.0.1", "Host载体")
	return s.cmd
}

func (c *SnmpCmd) Execute(config *model.Stream) {
	client := snmp.NewSnmpClient(config.DstHost)
	for _, version := range SNMPVersion {
		// 获取客户端
		client.Execute(version)
		time.Sleep(time.Millisecond * 300)
	}
	logger.Infof("Stoped  SNMP Host : %s ...\n", config.DstHost)
}
