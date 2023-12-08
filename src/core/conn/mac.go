// Package conn -----------------------------
// @file      : mac.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:10
// -------------------------------------------
package conn

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	logger "github.com/sirupsen/logrus"
	"log"
	"net"
	"poison/src/model"
)

type MacModel struct {
	SrcMac    string
	DstMac    string
	EtherFlag uint16
	Payload   []byte
	InterFace string
}

func (m MacModel) init() model.Messages { // 构建以太网帧
	if m.EtherFlag == 0 {
		return false
	}
	return true

}

func (m MacModel) Send() model.Messages {
	m.init()
	conn, err := pcap.OpenLive(m.InterFace, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// 构建packet
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	switch m.init() {
	case false:
		err = gopacket.SerializeLayers(buffer, options, gopacket.Payload(m.Payload))
		if err != nil {
			logger.Fatal(err)
		}
	case true:
		srcMAC, _ := net.ParseMAC(m.SrcMac) // 发送者MAC地址
		dstMAC, _ := net.ParseMAC(m.DstMac) // 接收者MAC地址
		ethLayer := &layers.Ethernet{
			SrcMAC:       srcMAC,
			DstMAC:       dstMAC,
			EthernetType: layers.EthernetType(m.EtherFlag), // 假设发送的是IPv4数据包
		}
		err = gopacket.SerializeLayers(buffer, options, ethLayer, gopacket.Payload(m.Payload))
		if err != nil {
			logger.Fatal(err)
		}
	}
	// 发送数据包
	err = conn.WritePacketData(buffer.Bytes())
	return buffer.Bytes()
}
