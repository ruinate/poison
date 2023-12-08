// Package service -----------------------------
// @file      : rpc.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:23
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpcapp "poison/src/core/rpc"
	"poison/src/model"
	"poison/src/utils"
)

type RpcCmd struct {
	cmd *cobra.Command
}

func (r *RpcCmd) InitCmd() *cobra.Command {
	r.cmd = &cobra.Command{
		Use:   model.RPC,
		Short: "rpc服务器",
		Long:  ``,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Println("Starting RPC_SERVER...")
			r.Execute()
		},
	}
	r.cmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "127.0.0.1", "Host载体")
	return r.cmd
}

func (r *RpcCmd) Execute() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		logger.Fatalln(err)
	}
	_ = rpc.RegisterName("Flow", new(rpcapp.RPC))
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持高并发
	}
}
