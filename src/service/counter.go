// Package service -----------------------------
// @file      : counter.go
// @author    : fzf
// @time      : 2023/5/11 上午9:27
// -------------------------------------------
package service

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	CounterPacket   int
	TemporaryPacket int
	CounterDepth    int
	Signal          = make(chan os.Signal, 1)
)

// DDosSpeed 专用
func DDosSpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logrus.Infof("Sended packet : %d  pps: %d \n", TemporaryPacket, TemporaryPacket/3)
		CounterPacket += TemporaryPacket
		TemporaryPacket = 0
	}
}

// ReplaySpeed 专用
func ReplaySpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logrus.Infof("Sended packet : %d  pps: %d \n", TemporaryPacket, TemporaryPacket/3)
		TemporaryPacket = 0
	}
}

type Limiter struct {
	limit int
	burst int
}
