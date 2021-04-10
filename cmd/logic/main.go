package main

import (
	"gim/config"
	"gim/internal/logic/api"
	"gim/pkg/db"
	"gim/pkg/logger"
	"gim/pkg/rpc"
)

func main() {
	logger.Init()
	db.InitMysql(config.Logic.MySQL)
	db.InitRedis(config.Logic.RedisIP, config.Logic.RedisPassword)

	// 初始化内部的连接客户端
	rpc.InitConnIntClient(config.Logic.ConnRPCAddrs)

	// 初始化内部的业务客户端
	rpc.InitBusinessIntClient(config.Logic.BusinessRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
