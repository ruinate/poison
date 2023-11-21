// Package utils -----------------------------
// @file      : Inter.go
// @author    : fzf
// @time      : 2023/5/14 下午10:58
// -------------------------------------------
package utils

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
