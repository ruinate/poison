// Package utils -----------------------------
// @file      : client.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:17
// -------------------------------------------
package utils

import (
	logger "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"poison/src/common"
	"poison/src/core/conn"
	"poison/src/model"
	"strings"
)

type ClientModel struct{}

var (
	Client ClientModel
)

func NewClient(config *model.Stream) conn.LayerModel {
	payload := common.SwitchHex(config.Payload)
	if config.SrcPort == 0 {
		config.TmpSrcPort = rand.Intn(65535-10254) + 1024
	} else {
		config.TmpSrcPort = config.SrcPort
	}
	if config.SrcHost == "127.0.0.1" {
		config.SrcHost = "0.0.0.0"
	}

	if config.SendMode == model.MAC {
		return conn.MacModel{
			SrcMac:    config.SrcMAC,
			DstMac:    config.DstMAC,
			EtherFlag: config.EtherFlag,
			Payload:   payload,
			InterFace: config.InterFace,
		}
	}
	switch config.Mode {
	case model.PROTOTCP:
		return conn.TCPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}
	case model.PROTOUDP:
		return conn.UDPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}
	case model.PROTOICMP:
		return conn.ICMPModel{
			DstHost: config.DstHost,
			Depth:   config.Depth,
		}
	case model.PROTOICS:
		conn_udp := conn.UDPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}
		go conn_udp.Send()
		return conn.TCPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}

	case model.PROTOBLACK:
		connUdp := conn.UDPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}
		go connUdp.Send()
		return conn.TCPModel{
			DstHost: config.DstHost,
			SrcHost: config.SrcHost,
			DstPort: config.DstPort,
			SrcPort: config.TmpSrcPort,
			Payload: payload,
			Common:  conn.Common{},
		}
	}
	return nil
}

func (c ClientModel) Execute(config *model.Stream) error {
	client := NewClient(config)
	result := client.Send()
	if config.SendMode == model.MAC {
		if byteResult, ok := result.([]byte); ok {
			logger.Infof("%s connected to the %s  payload: %#v", config.SrcMAC, config.DstMAC, byteResult)
		}
		return nil
	}
	// 判断返回类型
	switch result.(type) {
	// 服务器端口未监听
	case error:
		if opErr, ok := result.(*net.OpError); ok {
			if strings.Contains(opErr.Error(), model.ConnectionUSEERROR) {
				logger.Errorf("%s connected to the %s  port: %d payload: %s", config.Mode, config.DstHost, config.DstPort, opErr)
				return c.Execute(config)
			}

			logger.Errorf("%s connected to the %s  port: %d payload: %s", config.Mode, config.DstHost, config.DstPort, opErr)
		}
	// 默认打印
	default:
		logger.Infof("%s connected to the %s  port: %d payload: %#v", config.Mode, config.DstHost, config.DstPort, result)
	}
	return nil
}
