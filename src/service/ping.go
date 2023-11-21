// Package service----------------------------
// @file      : ping.go
// @author    : fzf
// @time      : 2023/11/8 下午3:27
// -------------------------------------------
package service

import (
	probing "github.com/prometheus-community/pro-bing"
	logger "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"poison/src/model"
	"poison/src/service/setting"
	"time"
)

type PingStruct struct{}

func (p PingStruct) Execute(config *model.InterfaceModel) {
	if config.Depth == 0 {
		logger.Fatalln("depth must be greater than 0")
	}
	pinger, err := probing.NewPinger(config.DstHost)
	signal.Notify(setting.Signal, os.Interrupt)
	go func() {
		for _ = range setting.Signal {
			pinger.Stop()
		}
	}()
	pinger.OnRecv = func(pkt *probing.Packet) {
		logger.Printf("time: %s------%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			time.Now().Format("2006-01-02 15:04:05"), pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		logger.Printf("time: %s------%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			time.Now().Format("2006-01-02 15:04:05"), pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnFinish = func(stats *probing.Statistics) {
		logger.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		logger.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		logger.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}
	pinger.Timeout = time.Second * time.Duration(config.Depth)
	pinger.Size = 56
	pinger.TTL = 128
	logger.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		logger.Errorln(err)
	}
}
