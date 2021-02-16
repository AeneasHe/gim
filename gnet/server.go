package gnet

import "net"

type GServer struct {
}

func (s *GServer) Run() {

}
func (s *GServer) GetConn(fd int32) (*Conn, bool) {
	return nil, true
}
func (s *GServer) Serve(lis net.Listener) error {
	return nil
}

func NewServer(port int, handler *Handler, others ...interface{}) (*Server, error) {
	// var s *GServer
	// s = &GServer{}

	return nil, nil
}
