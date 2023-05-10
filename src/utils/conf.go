// Package strategy -----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package utils

type PoisonConfig struct {
	Depth   int
	Mode    string
	Host    string
	Port    int
	Payload string
}

var Config PoisonConfig
