package strategy

import (
	"PoisonFlow/src/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/syossan27/tebata"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
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
	logrus.Infof("Starting server Host : %s  Mode : %s", config.Host, config.Mode)
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
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
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
			logrus.Printf("%s -> %s", TCPServer.RemoteAddr(), TCPServer.LocalAddr())
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
			select {
			case <-time.After(0 * time.Millisecond):
				// 读取消息，UDP不是面向连接的因此不需要等待连接
				length, udpAddr, err := UDPServer.ReadFromUDP(buf)
				if err != nil {
					break
				}
				_, err = UDPServer.WriteToUDP(buf[:length], udpAddr)
				s.Close(UDPServer)
				logrus.Printf("%s -> %s", udpAddr, UDPServer.LocalAddr())
			case s := <-sigs:
				logrus.Errorf("received signal %s, exiting", s.String())
				os.Exit(0)
			}
		}
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
