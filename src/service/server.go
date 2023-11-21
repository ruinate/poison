// Package service  -----------------------------
// @file      : server.go
// @author    : fzf
// @time      : 2023/5/9 上午10:20
// -------------------------------------------
package service

import (
	logger "github.com/sirupsen/logrus"
	"github.com/syossan27/tebata"
	"net"
	"os"
	"os/signal"
	"poison/src/model"
	"poison/src/service/setting"
	"strconv"
	"syscall"
	"time"
)

type ServerStruct struct {
}

var buf = make([]byte, 1024)

// Execute 监听执行
func (s *ServerStruct) Execute(config *model.InterfaceModel) {
	t := tebata.New(syscall.SIGINT, syscall.SIGTERM)
	for port := config.StartPort; port < config.EndPort; {
		port++
		go func(port string) {
			if err := s.ExecuteListen(config.DstHost, port, config.Mode, t); err != nil {
			}
		}(strconv.Itoa(port))
	}
	if err := s.ExecuteListen(config.DstHost, strconv.Itoa(65535), config.Mode, t); err != nil {
	}
}

// ExecuteListen  监听主程序
func (s *ServerStruct) ExecuteListen(address, port, protocol string, t *tebata.Tebata) error {
	signal.Notify(setting.Signal, syscall.SIGINT, syscall.SIGTERM)
	if protocol == "TCP" {
		conn, err := net.Listen("tcp", address+":"+port)
		if err != nil {
			return err
		}
		_ = t.Reserve(conn.Close)

		for {
			select {
			case <-time.After(0 * time.Millisecond):
				// 监听端口
				TCPServer, err := conn.Accept()
				if err != nil {
					break
				}
				results, _ := TCPServer.Read(buf)
				_, _ = TCPServer.Write(buf[:results])
				if port == "135" {
					time.Sleep(time.Second * 5)
				}
				s.Close(TCPServer)
				logger.Printf("%s -> %s", TCPServer.RemoteAddr(), TCPServer.LocalAddr())
			case s := <-setting.Signal:
				logger.Errorf("received signal %s, exiting", s.String())
				os.Exit(0)
			}

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
				logger.Printf("%s -> %s", udpAddr, UDPServer.LocalAddr())
			case s := <-setting.Signal:
				logger.Errorf("received signal %s, exiting", s.String())
				os.Exit(0)
			}
		}
	}
}

// Close 服务端关闭连接
func (s *ServerStruct) Close(client net.Conn) {
	err := client.Close()
	if err != nil {
		return
	}
}
