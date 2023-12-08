// Package replay -----------------------------
// @file      : replay.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午2:01
// -------------------------------------------
package replay

import (
	"context"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net"
	"os"
	"path/filepath"
	"poison/src/common/settings"
	"poison/src/model"
	"time"
)

func SendPacket(path, inter string, speed, depth int) error {
	files := FindAllFiles(path)
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
			case _ = <-settings.Signal:
				PcapResults(settings.TotalPacket, settings.TotalBytes)
			default:
				settings.TemporaryPacket += 1
				settings.TotalBytes += int64(len(packet.Data()))
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
	settings.TotalPacket += settings.TemporaryPacket
	settings.TotalDepth++
	if depth == settings.TotalDepth {
		PcapResults(settings.TotalPacket, settings.TotalBytes)
	}
	return nil
}

func PcapResults(packet int, bytes int64) {
	elapsed := time.Now().Sub(settings.StartTime)
	logger.Printf("stopped sending a total of %d packet", packet)
	logger.Printf("Total bytes: %d\n", bytes)
	logger.Printf("Elapsed time: %v\n", elapsed)
	logger.Printf("Mbps: %.2f\n", float64(bytes)/elapsed.Seconds()*8/1000000)
	os.Exit(0)
}

// ReplaySpeed 专用
func ReplaySpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logger.Infof("Sended packet : %d  pps: %d \n", settings.TemporaryPacket, settings.TemporaryPacket/3)
		settings.TemporaryPacket = 0
	}
}

func R(config *model.Stream) {
	for {
		select {
		// 捕获ctrl + c
		case _ = <-settings.Signal:
			PcapResults(settings.TotalPacket, settings.TotalBytes)
		default:
			err := SendPacket(config.FilePath, config.InterFace, config.Speed, config.Depth)
			if err != nil {
				logger.Fatalln(err)
			}
		}

	}
}

func FindAllFiles(path string) []string {
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
	if err != nil {
		logger.Fatalln(err)
	}
	return file
}
