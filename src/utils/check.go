// Package utils -----------------------------
// @file      : check.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:51
// -------------------------------------------
package utils

import (
	"errors"
	"net"
	"poison/src/model"
)

var (
	ERRORPORT  = errors.New("Please check format of port: e.g. 1-65535 ")
	ERRORHOST  = errors.New("please check format of host: e.g. 192.168.1.1")
	ERRORPATH  = errors.New("please check format of file: no such file or directory")
	ERRORDEPTH = errors.New("please check format of depth: depth <= 0")
	ERRORMODE  = errors.New("please check format of mode: \"TCP\", \"UDP\", \"ICS\", \"ICMP\", \"WinNuke\", \"Smurf\", \"Land\", \"TearDrop\", \"MAXICMP\"")
)

type Check struct {
	Config *model.Stream
}

func (c Check) host() error {
	dsthost := net.ParseIP(c.Config.DstHost)
	srchost := net.ParseIP(c.Config.DstHost)
	if dsthost != nil || dsthost.To4() != nil || srchost != nil || srchost.To4() != nil {
		return nil
	} else {
		return ERRORHOST
	}
}

func (c Check) port() error {
	if c.Config.DstPort >= 1 && c.Config.DstPort <= 65535 {
		if c.Config.SrcPort >= 1 && c.Config.SrcPort <= 65535 || c.Config.SrcPort == 0 {
			return nil
		}
	}
	return ERRORPORT
}

func (c Check) depth() error {
	if c.Config.Depth < 0 {
		return ERRORDEPTH
	}
	return nil

}

//func (c Check) filepath() error {
//	if len(FindAllFiles(c.Config.FilePath)) == 0 {
//		return ERRORPATH
//	}
//	return nil
//}

func (c Check) mode() error {
	for _, mode := range model.PROTOMODE {
		if mode == c.Config.Mode {
			return nil
		}
	}
	return ERRORMODE
}

func CheckFlag(config *model.Stream) error {
	check := &Check{
		Config: config,
	}
	if err := check.port(); err != nil {
		return err
	}
	if err := check.host(); err != nil {
		return err
	}
	if err := check.depth(); err != nil {
		return err
	}
	if err := check.mode(); err != nil {
		return err
	}
	return nil
}
