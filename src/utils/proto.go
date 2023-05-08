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
	maxLength = 20 // 设置result最大长度为20个字符

)

// TCP 客户端
func (p *ProtoAPP) TCP(address string, config *ProtoAPP) (*ProtoAPP, error) {
	client, err := net.DialTimeout("tcp", address, time.Millisecond*300)
	// 连接出错则打印错误消息并退出程序
	if err != nil {
		return nil, err
	}
	_, err = client.Write(config.HexPayload)
	defer p.Close(client)
	p.ProcessResult(p.Receive(client))
	return p, nil
}

// UDP 客户端
func (p *ProtoAPP) UDP(address string, config *ProtoAPP) (*ProtoAPP, error) {
	client, _ := net.DialTimeout("udp", address, time.Millisecond*500)
	defer p.Close(client)
	_, _ = client.Write(config.HexPayload)
	if len(p.Payload) > maxLength {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", p.Mode, p.Host, p.Port, p.Payload)
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", p.Mode, p.Host, p.Port, p.Payload)
	}

	return p, nil
}

// Receive 读取服务端返回数据
func (p *ProtoAPP) Receive(conn net.Conn) error {
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
func (p *ProtoAPP) Close(conn net.Conn) {
	err := conn.Close()
	if err != nil {
	}
}

//SwitchHex  转换为16进制
func (p *ProtoAPP) SwitchHex(payload string) []byte {
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

func (p *ProtoAPP) ProcessResult(err error) {
	if err != nil {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", p.Mode, p.Host, p.Port, err.Error())
	} else {
		p.Result = fmt.Sprintf("%s connected to the %s  port: %d payload: %#v", p.Mode, p.Host, p.Port, p.Result)
	}
}

// Execute  运行方法
func (p *ProtoAPP) Execute(config *ProtoAPP) (*ProtoAPP, error) {
	p.HexPayload = p.SwitchHex(config.Payload)
	var address = fmt.Sprintf("%s:%d", config.Host, config.Port)
	switch p.Mode {
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

var Config ProtoAPP
