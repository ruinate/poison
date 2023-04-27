package utils

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"
)

type ProtoAPP struct {
	Depth      int
	Mode       string
	Host       string
	Port       int
	Payload    string
	Result     string
	HexPayload []byte
}

var (
	data      = make([]byte, 1024)
	maxLength = 20 // 设置result最大长度为20个字符
)

// TCP 客户端
func (p *ProtoAPP) TCP(config *ProtoAPP) (*ProtoAPP, error) {
	var addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := net.DialTimeout("tcp", addr, time.Millisecond*300)
	// 连接出错则打印错误消息并退出程序
	if err != nil {
		return nil, err
	}
	_, err = client.Write(config.HexPayload)
	defer p.Close(client)
	_, err = p.Receive(client)
	if err != nil {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, err.Error())
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, p.Payload)
	}
	return p, nil
}

// UDP 客户端
func (p *ProtoAPP) UDP(config *ProtoAPP) (*ProtoAPP, error) {
	var addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	client, err := net.DialTimeout("udp", addr, time.Millisecond*300)
	if err != nil {
		return nil, err
	}
	defer p.Close(client)
	_, err = client.Write(config.HexPayload)
	_, err = p.Receive(client)
	if err != nil {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, err.Error())
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", config.Mode, config.Host, config.Port, p.Payload)
	}
	return p, nil

}

// RUN 运行方法
func (p *ProtoAPP) RUN(config *ProtoAPP) (*ProtoAPP, error) {
	config.HexPayload = p.SwitchHex(config.Payload)
	switch p.Mode {
	case "TCP":
		{
			return p.TCP(config)
		}
	case "UDP":
		{
			return p.UDP(config)
		}
	case "ICS":
		{
			// 协程UDPe
			go p.UDP(config)
			return p.TCP(config)
		}
	case "BLACK":
		{
			return p.TCP(config)
		}
	default:
		return p, nil
	}
}

// Receive 读取服务端返回数据
func (p *ProtoAPP) Receive(conn net.Conn) (*ProtoAPP, error) {
	for {
		err := conn.SetDeadline(time.Now().Add(time.Millisecond * 500))
		n, err := conn.Read(data)
		if err != nil && n == 0 {
			return nil, err
		} else {
			message := string(data[:n])
			if len(message) > maxLength {
				message = message[:maxLength]
			}
			p.Result = message
		}
		if len(p.Payload) > maxLength {
			p.Payload = p.Payload[:maxLength]
		}
		return p, nil
	}

}

// Close 关闭客户端连接
func (p *ProtoAPP) Close(conn net.Conn) {
	err := conn.Close()
	if err != nil {
	}
}

//SwitchHex  转换为16进制
func (p *ProtoAPP) SwitchHex(payload string) []byte {
	s := strings.Split(payload, "|")
	for _, s := range s {
		HexPayload, err := hex.DecodeString(s)
		if err != nil {
			p.HexPayload = append(p.HexPayload, []byte(payload)...)
		} else {
			p.HexPayload = append(p.HexPayload, HexPayload...)
		}
	}
	return p.HexPayload

}

var Config ProtoAPP
