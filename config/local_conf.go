package config

import (
	"gim/pkg/logger"

	"go.uber.org/zap"
)

func initLocalConf() {
	// 逻辑层
	Logic = LogicConf{
		MySQL:            "root:root123@tcp(localhost:3306)/gim?charset=utf8&parseTime=true&loc=Local",
		NSQIP:            "127.0.0.1:4150",
		RedisIP:          "127.0.0.1:6379",
		RedisPassword:    "",
		RPCIntListenAddr: ":50000",
		RPCExtListenAddr: ":50001",
		ConnRPCAddrs:     "addrs:///127.0.0.1:50100,127.0.0.1:50200",
		BusinessRPCAddrs: "addrs:///127.0.0.1:50300",
	}

	// TCP通讯层
	TCPConn = TCPConnConf{
		TCPListenAddr: 8080,
		RPCListenAddr: ":50100",
		LocalAddr:     "127.0.0.1:50100",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	// WS通讯层
	WSConn = WSConnConf{
		WSListenAddr:  ":8081",
		RPCListenAddr: ":50200",
		LocalAddr:     "127.0.0.1:50200",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	// 业务层
	Business = BusinessConf{
		MySQL:            "root:root123@tcp(localhost:3306)/gim?charset=utf8&parseTime=true",
		NSQIP:            "127.0.0.1:4150",
		RedisIP:          "127.0.0.1:6379",
		RPCIntListenAddr: ":50300",
		RPCExtListenAddr: ":50301",
		LogicRPCAddrs:    "addrs:///127.0.0.1:50000",
	}

	logger.Leavel = zap.DebugLevel
	logger.Target = logger.Console
}
