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

	// 先启动内部的业务服务器，内部指不对app提供服务，只对服务器端其他微服务提供服务
	go func() {
		defer util.RecoverPanic()
		intListen, err := net.Listen("tcp", config.Business.RPCIntListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(UserIntInterceptor))
		pb.RegisterBusinessIntServer(intServer, &BusinessIntServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

	// 然后启动外部的业务服务器
	go func() {
		defer util.RecoverPanic()

		extListen, err := net.Listen("tcp", config.Business.RPCExtListenAddr)
		if err != nil {
			panic(err)
		}
		extServer := grpc.NewServer(grpc.UnaryInterceptor(UserExtInterceptor))
		pb.RegisterBusinessExtServer(extServer, &BusinessExtServer{})
		err = extServer.Serve(extListen)
		if err != nil {
			panic(err)
		}
	}()

}
