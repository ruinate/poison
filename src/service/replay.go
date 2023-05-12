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
	"fmt"
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
	go ReplaySpeed()
	for {
		r.SendPacket(config.FilePath, config.InterFace, config.Speed)
		if config.Depth == CounterDepth {
			logrus.Printf("stopped sending a total of %d packet", CounterPacket)
			os.Exit(0)
		}
	}
}
func (r *Replay) SendPacket(path, inter string, speed int) {
	files := r.FindAllFiles(path)
	limiter := rate.NewLimiter(rate.Limit(speed), 1)
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
				fmt.Println(CounterPacket)
				logrus.Printf("stopped sending a total of %d packet", CounterPacket)
				os.Exit(0)
			default:
				TemporaryPacket += 1
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
	CounterPacket += TemporaryPacket
	CounterDepth++
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
