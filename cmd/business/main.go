package main

import (
	"fmt"
	"gim/config"
	"gim/internal/business/api"
	"gim/pkg/db"
	"gim/pkg/logger"
	"gim/pkg/rpc"
)

func main() {
	logger.Init()
	fmt.Println(config.Business.MySQL)
	db.InitMysql(config.Business.MySQL)
	db.InitRedis(config.Business.RedisIP, config.Logic.RedisPassword)

	// 初始化内部的逻辑客户端
	rpc.InitLogicIntClient(config.Business.LogicRPCAddrs)

	api.StartRpcServer()
	logger.Logger.Info("user server start")
	select {}
}
