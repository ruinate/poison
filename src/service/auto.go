// Package service -----------------------------
// @file      : auto.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:02
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"poison/src/common/settings"
	"poison/src/model"
	"poison/src/payload"
	"poison/src/utils"
	"time"
)

type AutoCmd struct {
	cmd *cobra.Command
}

func (a *AutoCmd) InitCmd() *cobra.Command {
	a.cmd = &cobra.Command{
		Use:       model.AUTO,
		Short:     "自动发送：TCP、UDP、BLACK、ICS",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-m", "-H", "-d"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			model.Config.SendMode = "ROUTE"
			logger.Infof("Starting Auto Mode %s ...\n", model.Config.Mode)
			a.Execute(&model.Config)
		},
	}
	a.cmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICS、BLACK")
	a.cmd.Flags().StringVarP(&model.Config.DstHost, "dsthost", "H", "127.0.0.1", "Host载体")
	a.cmd.Flags().StringVarP(&model.Config.SrcHost, "srchost", "S", "127.0.0.1", "Host载体")
	a.cmd.Flags().IntVarP(&model.Config.SrcPort, "sport", "s", 0, "源端口")
	a.cmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	a.cmd.Flags().StringVarP(&model.Config.ICSMode, "icsmode", "i", "all", "ICS模式选择")

	// auto
	err := a.cmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"TCP", "UDP", "ICS", "BLACK"}, cobra.ShellCompDirectiveDefault
	})
	err = a.cmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = a.cmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = a.cmd.RegisterFlagCompletionFunc(model.ICSMODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return model.PROTOICSMODE, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
	}
	return a.cmd
}

func (a *AutoCmd) Execute(config *model.Stream) {
	dataList := payload.Output.Execute(config.Mode, config.ICSMode)
	if dataList != nil {
		for {
			for _, P := range dataList {
				time.Sleep(time.Millisecond * 300)
				select {
				case _ = <-settings.Signal:
					logger.Printf("stopped sending a total of %d packets", settings.TotalPacket)
					os.Exit(0)
				case <-time.After(0 * time.Millisecond):
					settings.TotalPacket += 1
					config.DstPort = P[0].(int)
					config.Payload = P[1].(string)
					config.SrcPort = 0
					if err := utils.Client.Execute(config); err != nil {
						logger.Errorln(err)
					}
				}
			}
			if settings.TotalDepth++; settings.TotalDepth == config.Depth {
				logger.Debugf("stopped sending a total of %d packets", settings.TotalPacket)
				return
			}
		}
	}
}
