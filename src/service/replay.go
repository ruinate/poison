// Package service  -----------------------------
// @file      : replay.go
// @author    : fzf
// @time      : 2023/5/10 下午4:49
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"context"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type ReplayInterFace interface {
	Execute(config *conf.ReplayModel)
}

type Replay struct {
}

func (r *Replay) Execute(config *conf.ReplayModel) {
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	go r.ReplaySpeed()
	if config.Speed == 0 {
		numCPU := runtime.NumCPU()
		logger.Printf("Limit Send mode---CPU：%d", numCPU)
		for i := 0; i < numCPU; i++ {
			go r.R(config)
		}
	}
	for {
		select {
		// 捕获ctrl + c
		case _ = <-Signal:
			r.PcapResults(TotalPacket, TotalBytes)
		default:
			err := r.SendPacket(config.FilePath, config.InterFace, config.Speed, config.Depth)
			if err != nil {
				utils.Check.CheckError(err)
			}
		}

	}
}
func (r *Replay) SendPacket(path, inter string, speed, depth int) error {
	files := utils.FindAllFiles(path)
	limiter := rate.NewLimiter(rate.Limit(speed), 10000)
	for _, file := range files {
		// 打开 pcap 文件
		pcapFile, err := pcap.OpenOffline(file)
		if err != nil {
			return err
		}
		// 获取网络接口
		InterFace, err := net.InterfaceByName(inter)
		if err != nil {
			return err
		}
		// 打开网络接口
		handle, err := pcap.OpenLive(InterFace.Name, 65536, true, time.Second)
		if err != nil {
			return err
		}
		// 循环读取 pcap 文件中的数据包
		packetSource := gopacket.NewPacketSource(pcapFile, pcapFile.LinkType())

		for packet := range packetSource.Packets() {
			select {
			// 捕获ctrl + c
			case _ = <-Signal:
				r.PcapResults(TotalPacket, TotalBytes)
			default:
				TemporaryPacket += 1
				TotalBytes += int64(len(packet.Data()))
				if err := limiter.Wait(context.Background()); err != nil {
					// 将数据包发送到本地网卡
					if err := handle.WritePacketData(packet.Data()); err != nil {
						break
					}
				}
				// 将数据包发送到本地网卡
				if err := handle.WritePacketData(packet.Data()); err != nil {
					break
				}
			}
		}
		pcapFile.Close()
		handle.Close()
	}
	TotalPacket += TemporaryPacket
	TotalDepth++
	if depth == TotalDepth {
		r.PcapResults(TotalPacket, TotalBytes)
	}
	return nil
}

func (r *Replay) PcapResults(packet int, bytes int64) {
	elapsed := time.Now().Sub(
		StartTime)
	logger.Printf("stopped sending a total of %d packet", packet)
	logger.Printf("Total bytes: %d\n", bytes)
	logger.Printf("Elapsed time: %v\n", elapsed)
	logger.Printf("Mbps: %.2f\n", float64(bytes)/elapsed.Seconds()*8/1000000)
	os.Exit(0)
}

// ReplaySpeed 专用
func (r *Replay) ReplaySpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logger.Infof("Sended packet : %d  pps: %d \n", TemporaryPacket, TemporaryPacket/3)
		TemporaryPacket = 0
	}
}

func (r *Replay) R(config *conf.ReplayModel) {
	for {
		select {
		// 捕获ctrl + c
		case _ = <-Signal:
			r.PcapResults(TotalPacket, TotalBytes)
		default:
			err := r.SendPacket(config.FilePath, config.InterFace, config.Speed, config.Depth)
			if err != nil {
				utils.Check.CheckError(err)
			}
		}

	}
}
