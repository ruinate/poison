// Package utils -----------------------------
// @file      : proto.go
// @author    : fzf
// @time      : 2023/11/21 上午9:46
// -------------------------------------------
package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	probing "github.com/prometheus-community/pro-bing"
	logger "github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"poison/src/model"
	"poison/src/service/setting"
	"strings"
	"sync"
	"time"
)

var (
	wg        sync.WaitGroup
	Client    ClientModel
	maxLength = 20
)

type ClientView struct {
	config *model.InterfaceModel
}

// Receive 读取服务端返回数据
func (c *ClientView) Receive(conn net.Conn) (string, error) {

	for {
		//返回数据，
		var data = make([]byte, 1024)
		err := conn.SetDeadline(time.Now().Add(time.Millisecond * 500))
		n, err := conn.Read(data)
		if err != nil && n == 0 {
			return "", err
		} else {
			message := string(data[:n])
			if len(message) > maxLength {
				message = message[:maxLength]
			}
			return message, nil
		}
	}
}

func (c *ClientView) ProcessResult(message string, err error) {
	switch c.config.Mode {
	case model.PROTOTCP, model.PROTOBLACK:
		if err != nil {
			logger.Errorf("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost, c.config.DstPort, err.Error())
		} else {
			logger.Infof("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost, c.config.DstPort, message)
		}
	case model.PROTOUDP, model.PROTOICS:
		if len(c.config.Payload) > maxLength {
			logger.Infof("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost,
				c.config.DstPort, c.config.Payload[:maxLength])
		} else {
			logger.Infof("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost,
				c.config.DstPort, c.config.Payload)
		}
	}
}

func (c *ClientView) Close(client net.Conn) {
	if client != nil {
		if err := client.Close(); err != nil {

		}
	}

}

// SwitchHex  转换为16进制
func (c *ClientView) SwitchHex(payload string) []byte {
	var HexData []byte
	PayloadSplit := strings.Split(payload, "|")
	for _, split := range PayloadSplit {
		HexPayload, err := hex.DecodeString(split)
		if err != nil {

			HexData = append(HexData, []byte(split)...)
		} else {
			HexData = append(HexData, HexPayload...)
		}
	}
	return HexData
}

type ClientModel struct {
	ClientView
}

func (c *ClientModel) init() {
	c.config.HexPayload = c.SwitchHex(c.config.Payload)
	if c.config.DstPort == 0 {
		c.config.DstPort = rand.Intn(65535-10254) + 1024
		c.MAC()
	}

	if c.config.SrcPort == 0 {
		c.config.TmpSrcPort = rand.Intn(65535-10254) + 1024
	}
	if c.config.SrcHost == "127.0.0.1" {
		c.config.SrcHost = "0.0.0.0"
	}
	// 如果是目的端口50012的话，确认该次发包为OPC协议
	if c.config.DstPort == 49168 {
		client, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", c.config.DstHost, 135), time.Millisecond*300)
		// 连接出错则打印错误消息并退出程序
		if err != nil {
			logger.Println(err)
		}
		_, err = client.Write(c.SwitchHex(DRPRPCPayload))
	}
}

func (c *ClientModel) TCP() error {
	lAddr := &net.TCPAddr{IP: net.ParseIP(c.config.SrcHost), Port: c.config.TmpSrcPort}
	//rAddr := &net.TCPAddr{IP: net.ParseIP(c.config.DstHost), Port: c.config.DstPort}
	d := net.Dialer{Timeout: time.Second * 1,
		LocalAddr: lAddr}
	client, err := d.Dial("tcp", fmt.Sprintf("%s:%d", c.config.DstHost, c.config.DstPort))
	if err != nil {
		return err
	}
	defer c.Close(client)
	_, err = client.Write(c.config.HexPayload)
	c.ProcessResult(c.Receive(client))
	return nil
}

func (c *ClientModel) UDP() error {
	lAddr := &net.UDPAddr{IP: net.ParseIP(c.config.SrcHost), Port: c.config.TmpSrcPort}
	rAddr := &net.UDPAddr{IP: net.ParseIP(c.config.DstHost), Port: c.config.DstPort}
	client, err := net.DialUDP("udp", lAddr, rAddr)
	if err != nil {
		return nil
	}
	//defer c.Close(client)
	_, err = client.Write(c.config.HexPayload)
	c.ProcessResult(c.Receive(client))

	return nil
}

func (c *ClientModel) MAC() error {
	// 获取设备的句柄
	handle, err := pcap.OpenLive(c.config.InterFace, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// 构建以太网帧
	srcMAC, _ := net.ParseMAC(c.config.SrcMAC) // 发送者MAC地址
	dstMAC, _ := net.ParseMAC(c.config.DstMAC) // 接收者MAC地址
	ethLayer := &layers.Ethernet{
		SrcMAC:       srcMAC,
		DstMAC:       dstMAC,
		EthernetType: 0x88a4, // 假设发送的是IPv4数据包
	}

	// 构建packet
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buffer, options, ethLayer, gopacket.Payload(c.config.HexPayload))
	if err != nil {
		logger.Fatal(err)
	}
	// 发送数据包
	err = handle.WritePacketData(buffer.Bytes())
	logger.Infof("%s connected  %s  to the  payload: %#v", c.config.DstMAC,
		c.config.SrcMAC, c.config.Payload)
	return nil
}

func (c *ClientModel) ICMP() error {
	if c.config.Depth == 0 {
		logger.Fatalln("depth must be greater than 0")
	}
	pinger, err := probing.NewPinger(c.config.DstHost)
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
	pinger.Timeout = time.Second * time.Duration(c.config.Depth)
	pinger.Size = 56
	pinger.TTL = 128
	logger.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		logger.Errorln(err)
	}
	return nil
}

func (c *ClientModel) Execute(config *model.InterfaceModel) error {
	c.config = config
	c.init()
	switch c.config.SendMode {
	case model.ROUTE:
		switch strings.ToUpper(c.config.Mode) {
		case model.PROTOTCP:
			return c.TCP()
		case model.PROTOUDP:
			return c.UDP()
		case model.PROTOICS:
			go func() {
				c.UDP()
			}()
			return c.TCP()
		case model.PROTOBLACK:
			return c.TCP()
		case model.PROTOICMP:
			return c.ICMP()
		default:
			logger.Errorln("模式输入错误")
			return nil
		}
	case model.MAC:
		return c.MAC()
	}
	return nil
}
