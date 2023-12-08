// Package service -----------------------------
// @file      : send.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:22
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/signal"
	"poison/src/common"
	"poison/src/common/settings"
	"poison/src/model"
	"poison/src/utils"
	"syscall"
	"time"
)

type SendCmd struct {
	cmd *cobra.Command
}

func (s *SendCmd) InitCmd() *cobra.Command {
	s.cmd = &cobra.Command{
		Use:       model.SEND,
		Short:     "发送数据包：TCP、UDP",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-M", "-m", "-H", "-S", "s", "-P", "-p", "-d"},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("Starting  Send Mode %s ...\n", model.Config.APPMode)
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			model.Config.SendMode = "ROUTE"
			s.Execute(&model.Config)
		},
	}
	s.cmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP")
	s.cmd.Flags().StringVarP(&model.Config.DstHost, "dsthost", "H", "127.0.0.1", "目的地址")
	s.cmd.Flags().StringVarP(&model.Config.SrcHost, "srchost", "S", "127.0.0.1", "源目的地址")
	s.cmd.Flags().StringVarP(&model.Config.Payload, "payload", "p", common.RandStr(20), "数据载体")
	s.cmd.Flags().IntVarP(&model.Config.SrcPort, "sport", "s", 0, "源端口")
	s.cmd.Flags().IntVarP(&model.Config.DstPort, "port", "P", 22, "目的端口")
	s.cmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	s.cmd.Flags().StringVarP(&model.Config.SendMode, "sendmode", "M", "ROUTE", "发送模式：MAC和ROUTE")
	err := s.cmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP"}, cobra.ShellCompDirectiveDefault
	})
	err = s.cmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = s.cmd.RegisterFlagCompletionFunc(model.PORT, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = s.cmd.RegisterFlagCompletionFunc(model.PAYLOAD, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = s.cmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
	}
	return s.cmd
}

func (s *SendCmd) Execute(config *model.Stream) {
	signal.Notify(settings.Signal, syscall.SIGINT, syscall.SIGTERM)
	for {
		time.Sleep(time.Millisecond * 300)
		select {
		case _ = <-settings.Signal:
			logger.Fatalln("stopped sending a total of %d packets", settings.TotalPacket)
		case <-time.After(0 * time.Millisecond):
			settings.TotalPacket += 1
			settings.TotalDepth += 1
			if err := utils.Client.Execute(config); err != nil {
				logger.Errorln(err)
			}
			if settings.TotalDepth == config.Depth {
				logger.Fatalln("stopped sending a total of %d packets", settings.TotalPacket)
			}
		}
	}
}
