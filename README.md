# PoisonFlow

#### 介绍
    flow


#### 软件架构
    软件架构说明
    go net
#### 安装教程

    / server
    Because there are too many open sockets, you need to set
    uname -SHn  70000
    
    insatll-
    bash build.sh
    

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
    
    Flags:
    -h, --help          help for PoisonFlow
    -n, --none string   send: 基础发送    auto: 自动发送  hping: 安全防护流量
    snmp：snmp客户端  server: 服务端 (default "text")
    
    Use "PoisonFlow [command] --help" for more information about a command.



#### 参与贡献
    1.  Fork 本仓库
    2.  新建 Feat_xxx 分支
    3.  提交代码
    4.  新建 Pull Request


