// Package conf -----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package conf

type PoisonConfig struct {
	Depth     int
	Mode      string
	Host      string
	Port      int
	Payload   string
	InterFace string
	Speed     int
	FilePath  string
}

var Config PoisonConfig
