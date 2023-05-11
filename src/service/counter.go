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

func PacketSpeed() {
	// 3秒中
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	// 协程输出发送pps
	for range ticker.C {
		logrus.Infof("Sended packet : %d  pps: %d \n", CounterPacket, CounterPacket/3)
	}
}

type Limiter struct {
	limit int
	burst int
}
