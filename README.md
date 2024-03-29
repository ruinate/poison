# PoisonFlow

#### 介绍
    poison


#### 软件架构
```text
软件架构说明
.
├── Makefile
├── README.md
├── main.go
├── src
│   ├── model
│   │   ├── core.go
│   │   └── server.go
│   ├── common
│   │   ├── code.go
│   │   ├── device.go
│   │   ├── random.go
│   │   ├── hping.go
│   │   └── settings
│   │       ├── hping
│   │       └── setting.go
│   ├── core
│   │   ├── snmp
│   │   │   └── snmp.go
│   │   ├── rpc
│   │   │   └── rpc.go
│   │   ├── replay
│   │   │   └── replay.go
│   │   ├── conn
│   │   │   ├── common.go
│   │   │   ├── mac.go
│   │   │   ├── model.go
│   │   │   ├── udp.go
│   │   │   ├── icmp.go
│   │   │   └── tcp.go
│   │   └── server
│   │       └── server.go
│   ├── service
│   │   ├── ping.go
│   │   ├── rpc.go
│   │   ├── send.go
│   │   ├── server.go
│   │   ├── snmp.go
│   │   ├── auto.go
│   │   ├── ddos.go
│   │   ├── ether.go
│   │   └── replay.go
│   ├── cmd
│   │   ├── completion.go
│   │   └── command.go
│   ├── payload
│   │   └── payload.go
│   └── utils
│       ├── search.go
│       ├── check.go
│       └── client.go
├── go.mod
├── go.sum
├── doc
│   ├── compre.md
│   └── poison.conf
└── poison
```
#### 安装教程
   
```shell
insatll-
    make 
    
    命令行 tab联想
    mv poison /usr/local/bin/
    poison completion bash > poison_completion
    mv poison_completion /etc/bash_completion.d/
    source /etc/bash_completion.d/poison_completion
```




#### 使用说明

    Usage:
    PoisonFlow [command]
    
    Available Commands:
    auto        自动发送：TCP、UDP、BLACK、ICS
    completion  Generate the autocompletion script for the specified shell
    ddos        安全防护
    help        Help about any command
    send        发送数据包：TCP、UDP
    server      服务端：监听端口默认全部
    snmp        SNMP 客户端连接测试
    rpc         RPC发送流量服务
    ping        发送ICMP
    
    Flags:
    -h, --help          help for PoisonFlow
    -n, --none string   send: 基础发送    auto: 自动发送  hping: 安全防护流量
    snmp：snmp客户端  server: 服务端 (default "text")
    
    Use "PoisonFlow [command] --help" for more information about a command.

#### 
    RPC示例
    var (
	config = conf.FlowModel{
		Depth:   1,
		Mode:    "TCP",
		Host:    "10.30.5.103",
		Port:    10086,
		Payload: "aqwert",
	}
	result *string = new(string)
    )
    
    func main() {
        // 这里不能再用rpc做连接，因为rpc内部会用Gob协议
        conn, err := net.Dial("tcp", "10.30.1.127:1234")
        if err != nil {
        panic("connection failed")
    }
        // 这里指定序列化协议为JSON
        client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
        err = client.Call("Flow.RPC", &config, result)
        if err != nil {
        panic("调用失败")
    }
        logger.Printf("RPC函数 调用成功")
        logger.Println(*result)
    }



#### 参与贡献
    1.  Fork 本仓库
    2.  新建 Feat_xxx 分支
    3.  提交代码
    4.  新建 Pull Request



#### 报错处理

    / server config
    Because there are too many open sockets, you need to set
    ulimit -SHn  70000
    / ping  config
    sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
    / supervisorexited: poison_server (exit status 0; expected)
    句柄数不够,修改/etc/supervisor/supervisord.conf 
    [supervisord]
    ...
    minfds=102400
    ...

    
    
