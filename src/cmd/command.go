// Package cmd -----------------------------
// @file      : command.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:58
// -------------------------------------------
package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"poison/src/model"
	"poison/src/service"
)

var (
	// PoisonCmd  总开关
	PoisonCmd = &cobra.Command{
		Use:       "poison",
		Short:     "Display one or many resources",
		Long:      ``,
		ValidArgs: []string{model.SEND, model.ETHER, model.AUTO, model.SNMP, model.SERVER, model.REPLAY, model.DDOS, model.PING, model.RPC},
		Args:      cobra.OnlyValidArgs,
		Version:   model.VERSION,
	}
	auto   service.AutoCmd
	send   service.SendCmd
	snmp   service.SnmpCmd
	ddos   service.DDOSCmd
	ping   service.PingCmd
	replay service.ReplayCmd
	rpc    service.RpcCmd
	ether  service.EtherCmd
	server service.ServerCmd
)

func init() {
	logger.SetLevel(logger.DebugLevel)
	if len(os.Args) >= 2 {
		model.Config.APPMode = os.Args[1]
	}
	//fmt.Println(
	//	`
	//	_ __   ___  (_)  ___  ___  _ __
	//	| '_ \ / _ \| | / __|/ _ \| '_ \
	//	| |_) | (_) | | \__ \ (_) | | | |
	//	| .__/ \___/|_| |___/\___/|_| |_|
	//	|_|
	//`,
	//)
	PoisonCmd.AddCommand(auto.InitCmd(), snmp.InitCmd(), send.InitCmd(), ddos.InitCmd(), ping.InitCmd(),
		replay.InitCmd(), rpc.InitCmd(), ether.InitCmd(), server.InitCmd())
}
