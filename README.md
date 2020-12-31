#### GO-Fping
> 用于产线扫描IP/端口的CLI小工具，支持高并发。


#### 安装
```shell
go build -v -o  main.go
```

#### 使用方法
```shell
批量IP网络探测工具，支持TCP,UDP,ICMP，支持高并发以及文件读取。

Usage:
  go-fping [command]

Available Commands:
  help        Help about any command
  icmp        icmp fping network
  tcp         tcp fping network
  udp         udp fping network

Flags:
      --config string     config file (default is $HOME/.go-fping.yaml)
  -c, --count int         ICMP 探测IP的Packet 数量 (default 2)
  -f, --file string       基于文件内IP探测
  -h, --help              help for go-fping
  -o, --output string     探测结果输出位置
  -r, --routinepool int   批量探测的并发池, 默认300个goroutine (default 300)
  -i, --singleip string   探测单个IP，EP: 192.168.1.1
  -g, --subnet string     探测整个网段, EP: 192.168.1.1/16
  -T, --timeout int       探测超时时间, 单位MS，默认1000ms (default 1000)
  -t, --toggle            Help message for toggle

Use "go-fping [command] --help" for more information about a command.
```

eg:
```shell
go-fping icmp -g 172.16.95.114/27
Reachable IP: [172.16.95.125 172.16.95.123 172.16.95.111 172.16.95.124 172.16.95.107 172.16.95.126 172.16.95.101 172.16.95.114 172.16.95.117]
UnReachable IP: [172.16.95.103 172.16.95.102 172.16.95.115 172.16.95.108 172.16.95.109 172.16.95.116 172.16.95.105 172.16.95.97 172.16.95.104 172.16.95.121 172.16.95.113 172.16.95.106 172.16.95.119 172.16.95.112 172.16.95.100 172.16.95.120 172.16.95.118 172.16.95.99 172.16.95.98 172.16.95.122 172.16.95.110]
```