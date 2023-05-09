// Package common -----------------------------
// @file      : common.go
// @author    : fzf
// @time      : 2023/5/9 上午11:00
// -------------------------------------------
package common

type ConfigType struct {
	Depth   int
	Mode    string
	Host    string
	Port    int
	Payload string
}

var Config ConfigType
