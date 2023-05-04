package strategy

import (
	"PoisonFlow/src/utils"
	"github.com/gosnmp/gosnmp"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	// Snmp 执行方法
	Snmp = &cobra.Command{
		Use:   "snmp [tab][tab]",
		Short: "SNMP 客户端连接测试",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			config := utils.Check.CheckSnmp(&utils.Config)
			SNMPVersion := [...]string{"v1", "v2", "v3"}
			for _, version := range SNMPVersion {
				// 获取客户端
				client := SNMP.SNMPClient(version, config)
				if client != nil {
					// 执行方法
					s, err := SNMP.SNMPExecute(client)
					if err != nil {
						log.Printf("snmp: %s   result: %s            \n", version, err.Error())
					} else {
						log.Printf("snmp: %s   result: %s            \n", version, s.result)
					}
				}
				time.Sleep(time.Millisecond * 300)
			}
		},
	}
)

type snmp struct {
	result string
}

// SNMPExecute SNMP执行程序
func (s *snmp) SNMPExecute(client *gosnmp.GoSNMP) (*snmp, error) {
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
	s.result = string(result.Variables[0].Value.([]byte))
	return s, nil
}

// SNMPClient SNMP客户端
func (s *snmp) SNMPClient(version string, config *utils.ProtoAPP) *gosnmp.GoSNMP {
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

// Close SNMP关闭连接
func (s *snmp) Close(client *gosnmp.GoSNMP) {
	err := client.Conn.Close()
	if err != nil {
		return
	}

}

// SNMP 实例化snmp
var SNMP snmp
