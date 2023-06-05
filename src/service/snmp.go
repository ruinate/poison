// Package service -----------------------------
// @file      : snmp.go
// @author    : fzf
// @time      : 2023/5/9 上午10:06
// -------------------------------------------
package service

import (
	"PoisonFlow/src/conf"
	"PoisonFlow/src/utils"
	"github.com/gosnmp/gosnmp"
	logger "github.com/sirupsen/logrus"
	"time"
)

type SnmpAPP struct {
	Result string
}

// Execute SNMP执行程序
func (s *SnmpAPP) Execute(config *conf.FlowModel) {
	SNMPVersion := [...]string{"v1", "v2", "v3"}
	for _, version := range SNMPVersion {
		// 获取客户端
		client := s.SNMPClient(version, config)
		s.RUN(version, client)
		time.Sleep(time.Millisecond * 300)
	}
	logger.Infof("Stoped  SNMP Host : %s ...\n", config.Host)
}

// SNMPClient SNMP客户端
func (s *SnmpAPP) SNMPClient(version string, config *conf.FlowModel) *gosnmp.GoSNMP {
	switch version {
	case "v1":
		{
			return &gosnmp.GoSNMP{
				Target:    config.Host,
				Port:      161,
				Community: "public",
				Version:   gosnmp.Version1,
				Timeout:   time.Millisecond * 1000,
			}
		}
	case "v2":
		{
			return &gosnmp.GoSNMP{
				Target:    config.Host,
				Port:      161,
				Community: "public",
				Version:   gosnmp.Version2c,
				Timeout:   time.Millisecond * 1000,
			}
		}
	case "v3":
		{
			return &gosnmp.GoSNMP{
				Target:        config.Host,
				Port:          161,
				Community:     "public",
				Version:       gosnmp.Version3,
				Timeout:       time.Millisecond * 1000,
				SecurityModel: gosnmp.UserSecurityModel,
				MsgFlags:      gosnmp.AuthPriv, //Authentication and no privacy
				SecurityParameters: &gosnmp.UsmSecurityParameters{
					UserName:                 "bolean",   //输入你设置的snmp用户名
					AuthenticationProtocol:   gosnmp.SHA, //经过身份验证的SnmpV3连接正在使用的身份验证协议。
					AuthenticationPassphrase: "admin123", //输入你公司的密码
					PrivacyProtocol:          gosnmp.AES,
					PrivacyPassphrase:        "admin123",
				},
			}
		}

	default:
		utils.Check.CheckExit("snmp version is not found")
		return nil
	}

}

// SNMPGetOID  SNMP执行程序
func (s *SnmpAPP) SNMPGetOID(client *gosnmp.GoSNMP) (*SnmpAPP, error) {
	err := client.Connect()
	defer s.Close(client)
	if err != nil {
		return nil, err
	}
	oid := []string{".1.3.6.1.2.1.1.1.0"}
	result, err := client.Get(oid)
	if err != nil {
		return nil, err
	}
	// 输出结果
	if len(result.Variables) == 0 {
		return nil, err
	}
	s.Result = string(result.Variables[0].Value.([]byte))
	return s, nil
}
func (s *SnmpAPP) RUN(version string, client *gosnmp.GoSNMP) {
	if client != nil {
		// 执行方法
		snmp, err := s.SNMPGetOID(client)
		if err != nil {
			logger.Infof("snmp: %s   result: %s            \n", version, err.Error())
		} else {
			logger.Infof("snmp: %s   result: %s            \n", version, snmp.Result)
		}
	}
}

// Close SNMP关闭连接
func (s *SnmpAPP) Close(client *gosnmp.GoSNMP) {
	err := client.Conn.Close()
	if err != nil {
		return
	}
}
