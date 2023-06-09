// Package service -----------------------------
// @file      : counter.go
// @author    : fzf
// @time      : 2023/5/11 上午9:27
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	TotalPacket     int
	TotalBytes      int64
	TotalDepth      int
	TemporaryPacket int
	Signal          = make(chan os.Signal, 1)
	StartTime       = time.Now()
	ScanPort        int
)

// DDosSpeed 专用
func DDosSpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logger.Infof("Sended packet : %d  pps: %d \n", TemporaryPacket, TemporaryPacket/3)
		TotalPacket += TemporaryPacket
		TemporaryPacket = 0
	}
}

type Limiter struct {
	limit int
	burst int
}
