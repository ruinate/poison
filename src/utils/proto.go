// Package utils -----------------------------
// @file      : proto.go
// @author    : fzf
// @time      : 2023/11/21 上午9:46
// -------------------------------------------
package utils

import (
	"encoding/hex"
	logger "github.com/sirupsen/logrus"
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
	case model.PROTOTCP:
		if err != nil {
			logger.Errorf("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost, c.config.DstPort, err.Error())
		} else {
			logger.Infof("%s connected to the %s  port: %d payload: %#v", c.config.Mode, c.config.DstHost, c.config.DstPort, message)
		}
	case model.PROTOUDP:
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
	if err := client.Close(); err != nil {

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

func (c ClientModel) Execute(config *model.InterfaceModel) error {
	c.config = config
	c.init()
	switch strings.ToUpper(c.config.Mode) {

	case model.PROTOTCP:
		return c.TCP()
	case model.PROTOUDP:
		return c.UDP()
	case model.PROTOICS:
		go func() {
			wg.Add(1)
			c.UDP()
			c.TCP()
		}()
		wg.Wait()
		return nil
	case model.PROTOBLACK:
		return c.TCP()
	default:
		logger.Errorln("模式输入错误")
		return nil
	}
}
