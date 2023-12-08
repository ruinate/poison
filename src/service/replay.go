// Package service -----------------------------
// @file      : replay.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:16
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/signal"
	"poison/src/common"
	"poison/src/common/settings"
	"poison/src/core/replay"
	"poison/src/model"
	"poison/src/utils"
	"runtime"
	"syscall"
)

type ReplayCmd struct {
	cmd *cobra.Command
}

func (r *ReplayCmd) InitCmd() *cobra.Command {
	r.cmd = &cobra.Command{
		Use:       model.REPLAY,
		Short:     "流量重放",
		Long:      ``,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"-i", "-f", "-s"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Debugln(err)
				return
			}
			logger.Printf("Starting Interface :%s   path :%s...\n", model.Config.InterFace, model.Config.FilePath)
			r.Execute(&model.Config)
		},
	}
	r.cmd.Flags().StringVarP(&model.Config.InterFace, "interface", "i", "lo", "接口载体")
	r.cmd.Flags().IntVarP(&model.Config.Speed, "speed", "s", 100000, "速度载体")
	r.cmd.Flags().StringVarP(&model.Config.FilePath, "file", "f", "", "路径载体")
	r.cmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 1, "循环载体")
	inter := common.TotalDevice()
	err := r.cmd.RegisterFlagCompletionFunc(model.REPLAYINTERFACE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inter, cobra.ShellCompDirectiveDefault
	})
	err = r.cmd.RegisterFlagCompletionFunc(model.REPLAYFILE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = r.cmd.RegisterFlagCompletionFunc(model.SPEED, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = r.cmd.RegisterFlagCompletionFunc(model.DEPTH, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
	}

	return r.cmd
}

func (r *ReplayCmd) Execute(config *model.Stream) {
	signal.Notify(settings.Signal, syscall.SIGINT, syscall.SIGTERM)
	if len(replay.FindAllFiles(config.FilePath)) == 0 {
		logger.Fatalln("please check format of file: no such file or directory")
	}
	go replay.ReplaySpeed()
	if config.Speed == 0 {
		numCPU := runtime.NumCPU()
		logger.Printf("Limit Send mode---CPU：%d", numCPU)
		for i := 0; i < numCPU; i++ {
			go replay.R(config)
		}
	}
	for {
		select {
		// 捕获ctrl + c
		case _ = <-settings.Signal:
			replay.PcapResults(settings.TotalPacket, settings.TotalBytes)
		default:
			err := replay.SendPacket(config.FilePath, config.InterFace, config.Speed, config.Depth)
			if err != nil {
				logger.Fatalln(err)
			}
		}

	}
}
