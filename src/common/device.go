// Package common -----------------------------
// @file      : device.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:51
// -------------------------------------------
package common

import (
	"net"
	"strings"
)

func TotalDevice() []string {
	inter := make([]string, 100)
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if strings.IndexRune(iface.Name, 'e') == 0 || strings.IndexRune(iface.Name, 'w') == 0 || strings.IndexRune(iface.Name, 'l') == 0 {
			inter = append(inter, iface.Name)
		}
	}
	return inter
}
