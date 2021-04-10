package tcp_conn

import (
	"gim/config"
	"gim/pkg/logger"
	"time"

	//"github.com/alberliu/gn"
	"github.com/cofepy/gn"
)

var server *gn.Server
var handler2 gn.Handler

var encoder = gn.NewHeaderLenEncoder(2, 1024)

// StartTCPServer 启动TCP长连接服务
func StartTCPServer() {
	// 设置日志
	//gn.SetLogger(logger.Sugar)

	// gn服务器实例，绑定处理方法handler
	var err error
	handler2 = &handler{}

	server, err = gn.NewServer(
		config.TCPConn.TCPListenAddr,
		handler2,
		gn.NewHeaderLenDecoder(2),
		gn.WithReadBufferLen(256),
		gn.WithTimeout(5*time.Minute, 11*time.Minute),
		gn.WithAcceptGNum(10),
		gn.WithIOGNum(100),
	)

	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	(*server).Run()
}
