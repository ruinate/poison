// Package utils -----------------------------
// @file      : Icmp.go
// @author    : fzf
// @time      : 2023/6/6 上午9:42
// -------------------------------------------
package utils

import (
	probing "github.com/prometheus-community/pro-bing"
	logger "github.com/sirupsen/logrus"
	"time"
)

func PING(host string) error {
	pinger, err := probing.NewPinger(host)
	pinger.OnRecv = func(pkt *probing.Packet) {
		logger.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		logger.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.TTL)
	}
	pinger.OnFinish = func(stats *probing.Statistics) {
		logger.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		logger.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		logger.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}
	pinger.Timeout = time.Second * 3
	pinger.Count = 4
	pinger.Size = 64
	pinger.TTL = 128
	logger.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		return err
	}
	return nil
}
