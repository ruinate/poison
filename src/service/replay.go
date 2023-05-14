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
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type ReplayInterFace interface {
	FindAllFiles(path string) []string
	Execute(config *conf.ReplayModel)
}

type Replay struct {
}

func (r *Replay) Execute(config *conf.ReplayModel) {
	signal.Notify(Signal, syscall.SIGINT, syscall.SIGTERM)
	go r.ReplaySpeed()
	for {
		if depth := r.SendPacket(config.FilePath, config.InterFace, config.Speed); config.Depth == depth {
			r.PcapResults(TotalPacket, TotalBytes)
		}
	}
}
func (r *Replay) SendPacket(path, inter string, speed int) int {
	files := r.FindAllFiles(path)
	limiter := rate.NewLimiter(rate.Limit(speed), 10000)
	for _, file := range files {
		// 打开 pcap 文件
		pcapFile, err := pcap.OpenOffline(file)
		utils.Check.CheckError(err)
		// 获取网络接口
		InterFace, err := net.InterfaceByName(inter)
		utils.Check.CheckError(err)
		// 打开网络接口
		handle, err := pcap.OpenLive(InterFace.Name, 65536, true, time.Second)
		utils.Check.CheckError(err)
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
	return TotalDepth
}
func (r *Replay) FindAllFiles(path string) []string {
	file := make([]string, 0)
	if filepath.Ext(path) == ".pcap" || filepath.Ext(path) == ".pcapng" {
		file = append(file, path)
		return file
	}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是 pcap文件  或者pcapng 则添加
		if !info.IsDir() && filepath.Ext(path) == ".pcap" || !info.IsDir() && filepath.Ext(path) == ".pcapng" {
			file = append(file, path)
		}
		return nil
	})
	utils.Check.CheckError(err)
	return file
}

func (r *Replay) PcapResults(packet int, bytes int64) {
	elapsed := time.Now().Sub(StartTime)
	logrus.Printf("stopped sending a total of %d packet", packet)
	logrus.Printf("Total bytes: %d\n", bytes)
	logrus.Printf("Elapsed time: %v\n", elapsed)
	logrus.Printf("Mbps: %.2f\n", float64(bytes)/elapsed.Seconds()*8/1000000)
	os.Exit(0)
}

// ReplaySpeed 专用
func (r *Replay) ReplaySpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logrus.Infof("Sended packet : %d  pps: %d \n", TemporaryPacket, TemporaryPacket/3)
		TemporaryPacket = 0
	}
}
