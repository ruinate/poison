// Package setting  -----------------------------
// @file      : config.go
// @author    : fzf
// @time      : 2023/11/20 上午10:50
// -------------------------------------------
package settings

import (
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
)
