package api

import (
	"gim/config"
	"gim/pkg/pb"
	"gim/pkg/util"
	"net"

	"google.golang.org/grpc"
)

// StartRpcServer 启动rpc服务
func StartRpcServer() {
	// 先启动内部的逻辑服务器
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.Logic.RPCIntListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicIntInterceptor))
		pb.RegisterLogicIntServer(intServer, &LogicIntServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

	// 然后启动外部的逻辑服务器
	go func() {
		defer util.RecoverPanic()

		extListen, err := net.Listen("tcp", config.Logic.RPCExtListenAddr)
		if err != nil {
			panic(err)
		}
		extServer := grpc.NewServer(grpc.UnaryInterceptor(LogicExtInterceptor))
		pb.RegisterLogicExtServer(extServer, &LogicExtServer{})
		err = extServer.Serve(extListen)
		if err != nil {
			panic(err)
		}
	}()

}
