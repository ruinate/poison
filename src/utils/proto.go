package utils

import (
	"PoisonFlow/src/conf"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
	"time"
)

type ProtoConfig struct {
	Result     string
	HexPayload []byte
}

var (
	maxLength = 20 // 设置result最大长度为20个字符

)

// TCP 客户端
func (p *ProtoConfig) TCP(address string, config *conf.FlowModel) (*ProtoConfig, error) {
	client, err := net.DialTimeout("tcp", address, time.Millisecond*300)
	// 连接出错则打印错误消息并退出程序
	if err != nil {
		return nil, err
	}
	_, err = client.Write(p.HexPayload)
	defer p.Close(client)
	p.ProcessResult(config, p.Receive(client))
	return p, nil
}

// UDP 客户端
func (p *ProtoConfig) UDP(address string, config *conf.FlowModel) (*ProtoConfig, error) {
	client, _ := net.DialTimeout("udp", address, time.Millisecond*500)
	defer p.Close(client)
	_, _ = client.Write(p.HexPayload)
	if len(config.Payload) > maxLength {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, config.Payload)
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, config.Payload)
	}

	return p, nil
}

// Receive 读取服务端返回数据
func (p *ProtoConfig) Receive(conn net.Conn) error {
	for {
		//返回数据，
		var data = make([]byte, 1024)
		err := conn.SetDeadline(time.Now().Add(time.Millisecond * 500))
		n, err := conn.Read(data)
		if err != nil && n == 0 {
			return err
		} else {
			message := string(data[:n])
			if len(message) > maxLength {
				message = message[:maxLength]
			}
			p.Result = message
		}
		return nil
	}

}

// Close 关闭客户端连接
func (p *ProtoConfig) Close(conn net.Conn) {
	err := conn.Close()
	if err != nil {
	}
}

//SwitchHex  转换为16进制
func (p *ProtoConfig) SwitchHex(payload string) []byte {
	var HexData []byte
	s := strings.Split(payload, "|")
	for _, s := range s {
		HexPayload, err := hex.DecodeString(s)
		if err != nil {
			HexData = append(HexData, []byte(payload)...)
		} else {
			HexData = append(HexData, HexPayload...)
		}
	}
	return HexData

}

func (p *ProtoConfig) ProcessResult(config *conf.FlowModel, err error) {
	if err != nil {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, err.Error())
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, p.Result)
	}
}

// Execute  运行方法
func (p *ProtoConfig) Execute(config *conf.FlowModel) (*ProtoConfig, error) {
	p.HexPayload = p.SwitchHex(config.Payload)
	var address = fmt.Sprintf("%s:%d", config.Host, config.Port)
	switch config.Mode {
	case "TCP":
		{
			return p.TCP(address, config)
		}
	case "UDP":
		{
			return p.UDP(address, config)
		}
	case "ICS":
		{
			// 协程UDP
			go func() {
				p, err := p.UDP(address, config)
				LogDebug(p, err)
			}()

			return p.TCP(address, config)
		}
	case "BLACK":
		{
			return p.TCP(address, config)
		}
	default:
		return p, nil
	}
}

func LogDebug(p *ProtoConfig, err error) {
	if err != nil {
		logrus.Errorf(err.Error())
	} else {
		logrus.Infof(p.Result)
	}
}
