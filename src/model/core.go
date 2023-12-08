// Package model -----------------------------
// @file      : core.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:10
// -------------------------------------------
package model

import "net"

var (
	MaxLength = 20
	Error     error
	Data      = make([]byte, 1024)
	OpcConn   net.Conn
)

type Messages interface{}
