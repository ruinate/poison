// Package snmp -----------------------------
// @file      : snmp.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:24
// -------------------------------------------
package snmp

import (
	"github.com/gosnmp/gosnmp"
	logger "github.com/sirupsen/logrus"
	"time"
)

type SNMP struct {
	DstHost    string
	snmpclient *gosnmp.GoSNMP
	Result     string
}

// SNMPClient SNMP客户端
func (s *SNMP) init(version string) {
	switch version {
	case "v1":
		{
			s.snmpclient = &gosnmp.GoSNMP{
				Target:    s.DstHost,
				Port:      161,
				Community: "public",
				Version:   gosnmp.Version1,
				Timeout:   time.Millisecond * 1000,
			}
		}
	case "v2":
		{
			s.snmpclient = &gosnmp.GoSNMP{
				Target:    s.DstHost,
				Port:      161,
				Community: "public",
				Version:   gosnmp.Version2c,
				Timeout:   time.Millisecond * 1000,
			}
		}
	case "v3":
		{
			s.snmpclient = &gosnmp.GoSNMP{
				Target:        s.DstHost,
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
		logger.Fatalln("snmp version is not found")
	}
}

// SNMPGetOID  SNMP执行程序
func (s *SNMP) SNMPGetOID() (*SNMP, error) {
	err := s.snmpclient.Connect()
	defer s.Close(s.snmpclient)
	if err != nil {
		return nil, err
	}
	oid := []string{".1.3.6.1.2.1.1.1.0"}
	result, err := s.snmpclient.Get(oid)
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

// Close SNMP关闭连接
func (s *SNMP) Close(client *gosnmp.GoSNMP) {
	err := client.Conn.Close()
	if err != nil {
		return
	}
}

func (s *SNMP) Execute(version string) {
	s.init(version)
	if s.snmpclient != nil {
		// 执行方法
		snmp, err := s.SNMPGetOID()
		if err != nil {
			logger.Infof("snmp: %s   result: %s            \n", version, err.Error())
		} else {
			logger.Infof("snmp: %s   result: %s            \n", version, snmp.Result)
		}
	}
}

func NewSnmpClient(host string) SNMP {
	return SNMP{
		DstHost: host,
	}
}
