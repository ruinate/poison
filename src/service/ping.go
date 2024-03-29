// Package service -----------------------------
// @file      : ping.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:23
// -------------------------------------------
package service

import (
	probing "github.com/prometheus-community/pro-bing"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"poison/src/common/settings"
	"poison/src/model"
	"poison/src/utils"
	"time"
)

type PingCmd struct {
	cmd *cobra.Command
}

func (p PingCmd) InitCmd() *cobra.Command {
	p.cmd = &cobra.Command{
		Use:   model.PING,
		Short: "ping",
		Long:  ``,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.CheckFlag(&model.Config); err != nil {
				logger.Fatalln(err)
			}
			logger.Println("Starting Ping...")
			p.Execute(&model.Config)
		},
	}
	p.cmd.Flags().StringVarP(&model.Config.DstHost, "host", "H", "127.0.0.1", "Host载体")
	p.cmd.Flags().IntVarP(&model.Config.Depth, "depth", "d", 4, "循环载体")
	return p.cmd
}

func (p PingCmd) Execute(config *model.Stream) {
	if config.Depth == 0 {
		logger.Fatalln("depth must be greater than 0")
	}
	pinger, err := probing.NewPinger(config.DstHost)
	signal.Notify(settings.Signal, os.Interrupt)
	go func() {
		for _ = range settings.Signal {
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
