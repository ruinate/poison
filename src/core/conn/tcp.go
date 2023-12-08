// Package conn -----------------------------
// @file      : tcp.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:10
// -------------------------------------------
package conn

import (
	"fmt"
	"net"
	"poison/src/common"
	"poison/src/model"
	"poison/src/payload"
	"time"
)

type TCPModel struct {
	DstHost string
	SrcHost string
	DstPort int
	SrcPort int
	Payload []byte
	Common
}

var (
	buf = make([]byte, 1024)
)

func (t TCPModel) record(conn net.Conn) string {
	// 接收响应
	n, err := conn.Read(buf)
	if err != nil {
		//fmt.Println("接收响应失败:", err)
	}
	response := string(buf[:n])
	return response
}

func (t TCPModel) init() model.Messages {
	// 如果是目的端口49168的话，确认该次发包为OPC协议
	if t.DstPort == payload.OPCPORT {
		model.OpcConn, model.Error = net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.DstHost, 135), time.Millisecond*300)
		// 连接出错则打印错误消息并退出程序
		if model.Error != nil {
			return model.Error
		}
		model.OpcConn.Write(common.SwitchHex(payload.DCERPCREST))
		t.record(model.OpcConn)
		model.OpcConn.Write(common.SwitchHex(payload.ISYSTEMREST))
		t.record(model.OpcConn)
	}
	return nil
}

func (t TCPModel) Send() model.Messages {
	t.init()
	lAddr := &net.TCPAddr{IP: net.ParseIP(t.SrcHost), Port: t.SrcPort}
	//rAddr := &net.TCPAddr{IP: net.ParseIP(t.DstHost), Port: t.DstPort}
	d := net.Dialer{Timeout: time.Second * 1,
		LocalAddr: lAddr}
	conn, err := d.Dial("tcp", fmt.Sprintf("%s:%d", t.DstHost, t.DstPort))
	if err != nil {
		return err
	}
	defer t.Close(conn)
	if t.DstPort == payload.OPCPORT {
		conn.Write(common.SwitchHex(payload.OPCDAREST))
		t.record(conn)
		conn.Write(common.SwitchHex(payload.OPCAEREST))
		t.record(conn)
	}
	_, err = conn.Write(t.Payload)
	return t.Record(conn)
}
