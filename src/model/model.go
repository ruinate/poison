// Package model-----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package model

const (
	AUTO            = "auto"
	SEND            = "send"
	DDOS            = "ddos"
	PING            = "ping"
	REPLAY          = "replay"
	RPC             = "rpc"
	SERVER          = "server"
	SNMP            = "snmp"
	REPLAYINTERFACE = "interface"
	REPLAYFILE      = "file"
	MODE            = "mode"
	HOST            = "host"
	PAYLOAD         = "payload"
	PORT            = "port"
	DEPTH           = "depth"
	SPEED           = "speed"
	ICSMODE         = "icsmode"
	PROTOTCP        = "TCP"
	PROTOUDP        = "UDP"
	PROTOBLACK      = "BLACK"
	PROTOICS        = "ICS"
	ROUTE           = "ROUTE"
	MAC             = "MAC"
)

type Stream struct {
	SrcMAC     string
	DstMAC     string
	DstHost    string
	SrcHost    string
	DstPort    int
	SrcPort    int
	Payload    string
	Depth      int
	SendMode   string
	Mode       string
	TmpSrcPort int
	HexPayload []byte
	Message    string
}

type AutoModel struct {
	ICSMode string
}

type ReplayModel struct {
	InterFace string
	Speed     int
	FilePath  string
}

type ServerModel struct {
	StartPort int
	EndPort   int
}

type APP struct {
	APPMode string
}

type InterfaceModel struct {
	Stream
	AutoModel
	ReplayModel
	ServerModel
	APP
}

type TestModel struct {
	Test string
}

var Config InterfaceModel
