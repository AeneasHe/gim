package rpc

import (
	"context"
	"fmt"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/pb"

	"google.golang.org/grpc"
)

var (
	LogicIntClient    pb.LogicIntClient
	ConnectIntClient  pb.ConnIntClient
	BusinessIntClient pb.BusinessIntClient
)

// 初始化内部的逻辑客户端
func InitLogicIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	LogicIntClient = pb.NewLogicIntClient(conn)
}

// 初始化内部的连接客户端
func InitConnIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpclib.Name)))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnectIntClient = pb.NewConnIntClient(conn)
}

// 初始化内部的业务客户端
func InitBusinessIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	BusinessIntClient = pb.NewBusinessIntClient(conn)
}
