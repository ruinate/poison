// Package utils -----------------------------
// @file      : proto.go
// @author    : fzf
// @time      : 2023/11/21 上午9:46
// -------------------------------------------
package utils

import (
	"encoding/hex"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	logger "github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"net"
	"poison/src/model"
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
func (c ClientView) Receive(conn net.Conn) (string, error) {

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

func (c ClientView) ProcessResult(message string, err error) {
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

func (c ClientView) Close(client net.Conn) {
	if client != nil {
		if err := client.Close(); err != nil {

		}
	}

}

// SwitchHex  转换为16进制
func (c ClientView) SwitchHex(payload string) []byte {
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

func (c ClientModel) init() {
	c.config.HexPayload = c.SwitchHex(c.config.Payload)
	if c.config.DstPort == 0 {
		c.config.DstPort = rand.Intn(65535-10254) + 1024
	}

	if c.config.SrcPort == 0 {
		c.config.TmpSrcPort = rand.Intn(65535-10254) + 1024
	}
	if c.config.SrcHost == "127.0.0.1" {
		c.config.SrcHost = "0.0.0.0"
	}
}

func (c ClientModel) TCP() error {
	lAddr := &net.TCPAddr{IP: net.ParseIP(c.config.SrcHost), Port: c.config.TmpSrcPort}
	rAddr := &net.TCPAddr{IP: net.ParseIP(c.config.DstHost), Port: c.config.DstPort}
	client, err := net.DialTCP("tcp", lAddr, rAddr)
	if err != nil {
		return err
	}
	defer c.Close(client)
	_, err = client.Write(c.config.HexPayload)
	c.ProcessResult(c.Receive(client))
	return nil
}

func (c ClientModel) UDP() error {
	lAddr := &net.UDPAddr{IP: net.ParseIP(c.config.SrcHost), Port: c.config.TmpSrcPort}
	rAddr := &net.UDPAddr{IP: net.ParseIP(c.config.DstHost), Port: c.config.DstPort}
	client, err := net.DialUDP("udp", lAddr, rAddr)
	defer c.Close(client)
	if err != nil {
		return nil
	}
	_, err = client.Write(c.config.HexPayload)
	c.ProcessResult(c.Receive(client))

	return nil
}

func (c ClientModel) MAC() error {
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

func (c ClientModel) Execute(config *model.InterfaceModel) error {
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
				c.TCP()
				c.UDP()
			}()
			return nil
		case model.PROTOBLACK:
			return c.TCP()
		default:
			logger.Errorln("模式输入错误")
			return nil
		}
	case model.MAC:
		return c.MAC()
	}
	return nil
}
