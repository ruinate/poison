// Package conf -----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package conf

type FlowModel struct {
	Depth   int
	Mode    string
	ICSMode string
	Host    string
	Port    int
	Payload string
	Scan    int
}

type ReplayModel struct {
	Depth     int
	InterFace string
	Speed     int
	FilePath  string
}
