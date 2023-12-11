// Package service -----------------------------
// @file      : ether.go
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

type EtherCmd struct {
	cmd *cobra.Command
}

func (e *EtherCmd) InitCmd() *cobra.Command {
	e.cmd = &cobra.Command{
		Use:       model.ETHER,
		Short:     "发送Ether数据包",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-I", "-S", "-D", "-p", "-d"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Infof("Starting  Host : %s ...\n", model.Config.DstHost)
			model.Config.SendMode = "MAC"
			e.Execute(&model.Config)
		},
	}
	e.cmd.Flags().StringVarP(&model.Config.InterFace, "interface", "i", "lo", "接口载体")
	e.cmd.Flags().StringVarP(&model.Config.DstMAC, "dstmac", "D", "00:11:22:33:44:55", "目的MAC地址")
	e.cmd.Flags().StringVarP(&model.Config.SrcMAC, "srcmac", "S", "66:77:88:99:aa:bb", "源目的MAC地址")
	e.cmd.Flags().StringVarP(&model.Config.Payload, "payload", "p", common.RandStr(20), "数据载体")
	e.cmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	e.cmd.Flags().Uint16VarP(&model.Config.EtherFlag, "flag", "f", 0x0800, "Ethernet协议头标识")
	inter := common.TotalDevice()
	err := e.cmd.RegisterFlagCompletionFunc(model.REPLAYINTERFACE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inter, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
	}
	return e.cmd
}

func (e *EtherCmd) Execute(config *model.Stream) {
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
				logger.Debugf("stopped sending a total of %d packets", settings.TotalPacket)
				return
			}
		}
	}
}
