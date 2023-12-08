// Package model -----------------------------
// @file      : server.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午12:52
// -------------------------------------------
package model

const (
	AUTO            string = "auto"
	ETHER           string = "ether"
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
)

const (
	PROTOTCP   string = "TCP"
	PROTOUDP   string = "UDP"
	PROTOBLACK string = "BLACK"
	PROTOICS   string = "ICS"
	PROTOICMP  string = "ICMP"
	ROUTE      string = "ROUTE"
	MAC        string = "MAC"
)

const (
	ConnectionUSEERROR = "address already in use"
)

type Stream struct {
	SrcMAC  string
	DstMAC  string
	DstHost string
	SrcHost string
	DstPort int
	SrcPort int
	// 临时源端口
	TmpSrcPort int

	Payload string
	Depth   int
	// MAC/ROUTE
	SendMode string
	// SERVER
	StartPort int
	EndPort   int

	// ICS
	ICSMode string
	Mode    string
	// Ether/Replay
	InterFace string

	// Replay
	Speed    int
	FilePath string
	// 判断APP
	APPMode string

	EtherFlag uint16
}

var (
	Config       Stream
	PROTOMODE    = []string{"TCP", "UDP", "ICS", "ICMP", "WinNuke", "Smurf", "Land", "TearDrop", "MAXICMP", "BLACK"}
	PROTOICSMODE = []string{"Modbus", "BACnet", "DNP3", "FINS", "OpcUA", "OpcDA",
		"OpcAE", "S7COMM", "ADS/AMS", "Umas", "ENIP",
		"Hart/IP", "S7COMM_PLUS", "IEC104", "CIP", "GE_SRTP", "EGD",
		"H1", "FF", "MELSOFT", "Ovation",
		"CoAP", "MQTT", "DLT645", "MELSOFT(1E)", "DeltaV", "Foxboro", "EtherCAT"}
)
