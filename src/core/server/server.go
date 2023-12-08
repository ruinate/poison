// Package server -----------------------------
// @file      : server.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午2:12
// -------------------------------------------
package server

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io"
	"net"
	"poison/src/common"
	"poison/src/payload"
	"strconv"
	"strings"
)

var (
	buf             = make([]byte, 65542)
	SERVERPORTERROR = make([]string, 0)
	rpcstatic       int
	opcstatic       int
)

func ListenerTCP(host, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		SERVERPORTERROR = append(SERVERPORTERROR, port)
		return
	}
	for {
		// 监听端口
		conn, err := listener.Accept()
		if err != nil {
			SERVERPORTERROR = append(SERVERPORTERROR, port)
			return
		}
		go HandleTCP(conn)
	}
}

func ListenerUDP(port int) {
	UDPServer, err := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: port,
		},
	)
	if err != nil {
		SERVERPORTERROR = append(SERVERPORTERROR, strconv.Itoa(port))
	}
	for {
		// 读取消息，UDP不是面向连接的因此不需要等待连接
		length, udpAddr, err := UDPServer.ReadFromUDP(buf)
		if err != nil {
			break
		}
		_, err = UDPServer.WriteToUDP(buf[:length], udpAddr)
		Close(UDPServer)
		logger.Printf("%s -> %s", udpAddr, UDPServer.LocalAddr())
	}
}

func HandleTCP(conn net.Conn) {

	localport := strings.Split(conn.LocalAddr().String(), ":")
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		logger.Printf("%s -> %s", conn.RemoteAddr(), conn.LocalAddr())
		switch localport[1] {
		case "135":
			switch rpcstatic {
			case 0:
				conn.Write(common.SwitchHex(payload.DCERPCRESP))
				rpcstatic++
			case 1:
				conn.Write(common.SwitchHex(payload.ISYSTEMRESP))
				rpcstatic = 0
				return
			}
		case "49168":
			switch opcstatic {
			case 0:
				conn.Write(common.SwitchHex(payload.OPCDARESP))
				opcstatic++
			case 1:
				conn.Write(common.SwitchHex(payload.OPCAERESP))
				opcstatic++
			case 2:
				conn.Write(buf[:n])
				opcstatic = 0
				Close(conn)
				return
			}
		default:
			conn.Write(buf[:n])
			conn.Close()
		}
	}
}

// Close 服务端关闭连接
func Close(client net.Conn) {
	err := client.Close()
	if err != nil {
		return
	}
}
