// Package model-----------------------------
// @file      : conf.go
// @author    : fzf
// @time      : 2023/5/10 下午12:48
// -------------------------------------------
package model

const (
	AUTO            string = "auto"
	SEND            string = "send"
	DDOS            string = "ddos"
	PING            string = "ping"
	REPLAY          string = "replay"
	RPC             string = "rpc"
	SERVER          string = "server"
	SNMP            string = "snmp"
	REPLAYINTERFACE string = "interface"
	REPLAYFILE      string = "file"
	MODE            string = "mode"
	HOST            string = "host"
	PAYLOAD         string = "payload"
	PORT            string = "port"
	DEPTH           string = "depth"
	SPEED           string = "speed"
	ICSMODE         string = "icsmode"
	PROTOTCP        string = "TCP"
	PROTOUDP        string = "UDP"
	PROTOBLACK      string = "BLACK"
	PROTOICS        string = "ICS"
	PROTOICMP       string = "ICMP"
	ROUTE           string = "ROUTE"
	MAC             string = "MAC"
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

var (
	Config       InterfaceModel
	PROTOMODE    = []string{"TCP", "UDP", "ICS", "ICMP", "WinNuke", "Smurf", "Land", "TearDrop", "MAXICMP"}
	PROTOICSMODE = []string{"Modbus", "BACnet", "DNP3", "FINS", "OpcUA", "OpcDA",
		"OpcAE", "S7COMM", "ADS/AMS", "Umas", "ENIP",
		"Hart/IP", "S7COMM_PLUS", "IEC104", "CIP", "GE_SRTP", "EGD",
		"H1", "FF", "MELSOFT", "Ovation",
		"CoAP", "MQTT", "DLT645", "MELSOFT(1E)", "DeltaV", "Foxboro", "EtherCAT"}
)
