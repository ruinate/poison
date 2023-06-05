package utils

import (
	"PoisonFlow/src/conf"
	"errors"
	logger "github.com/sirupsen/logrus"
	"net"
	"os"
)

type CheckAPP struct{}

var (
	ErrorSendMode   = errors.New("please check format of send mode: e.g. TCP,UDP")
	ErrorAUTOMode   = errors.New("please check format of send mode: e.g. TCP,UDP,BLACK,ICS")
	ErrorServerMode = errors.New("please check format of Server mode: e.g. TCP, UDP")
	ErrorDDOSMode   = errors.New("Please check format of Server mode: e.g. TCP, UDP,ICMP, WinNuke, Smurf ")
	ErrorPort       = errors.New("Please check format of port : e.g. 1-65535 ")
	ErrorHost       = errors.New("please check format of host: e.g. 1.2.3.4")
	ErrorPath       = errors.New("Fatal error: lstat : no such file or directory ")
)

func (c *CheckAPP) CheckHost(host string) error {
	ip := net.ParseIP(host)
	if ip != nil || ip.To4() != nil {
		return nil
	} else {
		return ErrorHost
	}
}
func (c *CheckAPP) CheckDepth(depth int) error {
	return nil

}
func (c *CheckAPP) CheckSendMode(mode string) error {
	if mode == "TCP" || mode == "UDP" {
		return nil
	} else {
		return ErrorSendMode
	}
}
func (c *CheckAPP) CheckPort(port int) error {
	if port >= 1 && port <= 65535 {
		return nil
	} else {
		return ErrorPort
	}
}
func (c *CheckAPP) CheckPath(path string) error {
	if path == "" {
		return ErrorPath
	}
	if len(FindAllFiles(path)) == 0 {
		return ErrorPath
	}
	return nil
}
func (c *CheckAPP) CheckAutoMode(mode string) error {
	if mode == "TCP" || mode == "UDP" || mode == "ICMP" || mode == "ICS" {
		return nil
	} else {
		return ErrorAUTOMode
	}
}
func (c *CheckAPP) CheckServerMode(mode string) error {
	if mode == "TCP" || mode == "UDP" {
		return nil
	} else {
		return ErrorServerMode
	}
}

func (c *CheckAPP) CheckDDosMode(mode string) error {
	if mode == "TCP" || mode == "UDP" || mode == "ICMP" || mode == "WinNuke" || mode == "Smurf" {
		return nil
	} else {
		return ErrorDDOSMode
	}
}

func (c *CheckAPP) CheckSend(config *conf.FlowModel) error {
	err := c.CheckSendMode(config.Mode)
	if err != nil {
		return err
	}
	err = c.CheckHost(config.Host)
	if err != nil {
		return err
	}
	err = c.CheckPort(config.Port)
	if err != nil {
		return err
	}
	err = c.CheckDepth(config.Depth)
	if err != nil {
		return err
	}
	return nil
}
func (c *CheckAPP) CheckDDos(config *conf.FlowModel) error {
	err := c.CheckDDosMode(config.Mode)
	if err != nil {
		return err
	}
	err = c.CheckHost(config.Host)
	if err != nil {
		return err
	}
	err = c.CheckPort(config.Port)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CheckAPP) CheckReplay(config *conf.ReplayModel) error {
	err := c.CheckDepth(config.Depth)
	if err != nil {
		return err
	}
	err = c.CheckPath(config.FilePath)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (c *CheckAPP) CheckAuto(config *conf.FlowModel) error {
	err := c.CheckHost(config.Host)
	if err != nil {
		return err
	}
	err = c.CheckDepth(config.Depth)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CheckAPP) CheckSnmp(config *conf.FlowModel) error {
	err := c.CheckHost(config.Host)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CheckAPP) CheckServer(config *conf.FlowModel) error {
	err := c.CheckServerMode(config.Mode)
	if err != nil {
		return err
	}
	err = c.CheckHost(config.Host)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CheckAPP) CheckExit(err string) {
	logger.Errorf("Fatal error: %s\n ", err)
}
func (c *CheckAPP) CheckDebug(debug string) {
	logger.Errorf("debug:  %s\n ", debug)
}
func (c *CheckAPP) CheckError(err error) {
	if err != nil {
		logger.Errorf("Fatal error: %s ", err)
		os.Exit(0)
	}
}
func (c *CheckAPP) CheckTimeout(err error) string {
	logger.Fatalln(os.Stderr, "Fatal error: %s ", err)
	return ""
}
func (c *CheckAPP) CheckDepthSum(CounterDepth, depth, CounterPacket int) bool {
	if CounterDepth == depth {
		logger.Printf("stopped sending a total of %d packets", CounterPacket)
		return false
	}
	return true
}

var Check CheckAPP
