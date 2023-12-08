// Package conn -----------------------------
// @file      : common.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:10
// -------------------------------------------
package conn

import (
	"net"
	"poison/src/model"
	"time"
)

type Common struct{}

func (c Common) Close(conn net.Conn) model.Messages {
	if conn != nil {
		if model.OpcConn != nil {
			if model.Error = model.OpcConn.Close(); model.Error != nil {
			}
		}
		if model.Error = conn.Close(); model.Error != nil {
			return true
		}
	}
	return false
}

func (c Common) Record(conn net.Conn) model.Messages {
	for {
		//返回数据
		model.Error = conn.SetDeadline(time.Now().Add(time.Millisecond * 500))
		n, err := conn.Read(model.Data)
		if err != nil && n == 0 {
			return nil
		} else {
			message := string(model.Data[:n])
			if len(message) > model.MaxLength {
				message = message[:model.MaxLength]
			}
			return model.Messages(message)
		}
	}
}
