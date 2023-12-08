// Package conn -----------------------------
// @file      : udp.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:10
// -------------------------------------------
package conn

import (
	"net"
	"poison/src/model"
)

type UDPModel struct {
	DstHost string
	SrcHost string
	DstPort int
	SrcPort int
	Payload []byte
	Common
}

func (u UDPModel) init() model.Messages {
	return nil
}

func (u UDPModel) Send() model.Messages {
	lAddr := &net.UDPAddr{IP: net.ParseIP(u.SrcHost), Port: u.SrcPort}
	rAddr := &net.UDPAddr{IP: net.ParseIP(u.DstHost), Port: u.DstPort}
	conn, err := net.DialUDP("udp", lAddr, rAddr)
	if err != nil {
		return nil
	}
	defer u.Close(conn)
	_, err = conn.Write(u.Payload)
	return string(u.Payload[:])
}
