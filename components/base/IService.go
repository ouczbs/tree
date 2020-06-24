package base

import "net"

type IService interface {
	Run()
	NewTcpConnection(conn net.Conn)IClientProxy
	MessageLoop()
}