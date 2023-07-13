// Package conf -----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package conf

type FlowModel struct {
	APPMode string
	Depth   int
	Mode    string
	ICSMode string
	Host    string
	Port    int
	Sport   int
	Payload string
	Scan    int
	Ports   PortRange
}
type PortRange struct {
	StartPort int
	EndPort   int
}

type ReplayModel struct {
	Depth     int
	InterFace string
	Speed     int
	FilePath  string
}
