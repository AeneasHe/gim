package gnet

import (
	"gim/pkg/pb"
	"net"

	"github.com/golang/protobuf/proto"
)

type Conn interface {
	GetFd() int32
	GetData() interface{}
	GetAddr() string
	SetData(data interface{})
}

type Server interface {
	Run()
	GetConn(fd int32) (*Conn, bool)
	Serve(lis net.Listener) error
}

type Handler interface {
	OnConnect(c *Conn)
	OnMessage(c *Conn, bytes []byte)
	OnClose(c *Conn, err error)
	Send(c *Conn, pt pb.PackageType, requestId int64, err error, message proto.Message)
	SignIn(c *Conn, input pb.Input)
	Sync(c *Conn, input pb.Input)
	Heartbeat(c *Conn, input pb.Input)
	MessageACK(c *Conn, input pb.Input)
}
