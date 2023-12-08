// Package common -----------------------------
// @file      : hping.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午2:05
// -------------------------------------------
package common

import (
	_ "embed"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"strconv"
)

//go:embed settings/hping
var content []byte

func Generate() string {
	fd, err := unix.MemfdCreate("hping", 0)
	if err != nil {
		log.Fatal(err)
	}
	hping := "/proc/" + strconv.Itoa(os.Getpid()) + "/fd/" + strconv.Itoa(int(fd))
	err = os.WriteFile(hping, content, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return hping
}
