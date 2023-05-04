package strategy

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/syossan27/tebata"
	"log"
	"net"
	"strconv"
	"syscall"
)

var (
	// Server 执行方法
	Server = &cobra.Command{
		Use:   "server [tab][tab]",
		Short: "服务端：监听端口默认全部",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckServer(&utils.Config)
			server.Execute(config)
		},
	}
)

type ServerApp struct {
}

var buf = make([]byte, 1024)

// Execute 监听执行
func (s *ServerApp) Execute(config *utils.ProtoAPP) {
	log.Println("server Host :", config.Host)
	log.Println("server Port  mode is ", config.Mode)
	t := tebata.New(syscall.SIGINT, syscall.SIGTERM)
	for port := 1; port < 65535; {
		port++
		go func(port string) {
			if err := s.ExecuteListen(config.Host, port, config.Mode, t); err != nil {
			}
		}(strconv.Itoa(port))
	}
	if err := s.ExecuteListen(config.Host, strconv.Itoa(65535), config.Mode, t); err != nil {
	}
}

// ExecuteListen  监听主程序
func (s *ServerApp) ExecuteListen(address, port, protocol string, t *tebata.Tebata) error {
	if protocol == "TCP" {
		conn, err := net.Listen("tcp", address+":"+port)
		if err != nil {
			return err
		}
		err = t.Reserve(conn.Close)
		utils.Check.CheckError(err)
		for {
			// 监听端口
			TCPServer, err := conn.Accept()
			utils.Check.CheckError(err)
			results, _ := TCPServer.Read(buf)
			_, _ = TCPServer.Write(buf[:results])
			s.Close(TCPServer)
			log.Printf("%s -> %s", TCPServer.RemoteAddr(), TCPServer.LocalAddr())
		}
	} else {
		port, _ := strconv.Atoi(port)
		UDPServer, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: port,
		})
		if err != nil {
			return err
		}
		for {
			// 读取消息，UDP不是面向连接的因此不需要等待连接
			length, udpAddr, err := UDPServer.ReadFromUDP(buf)
			if err != nil {
				log.Printf("Read from udp server:%s failed,err:%s", udpAddr, err)
				break
			}
			_, err = UDPServer.WriteToUDP(buf[:length], udpAddr)
			if err != nil {
				fmt.Println("write to udp server failed,err:", err)
			}
			s.Close(UDPServer)
			log.Println("[ server ]# UdpAddr: ", udpAddr, "Data: ", string(buf[:length]))
		}
		return err
	}

}

// Close 服务端关闭连接
func (s *ServerApp) Close(client net.Conn) {
	err := client.Close()
	if err != nil {
		return
	}
}

var server = ServerApp{}
