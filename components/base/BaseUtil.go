package base

import (
	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/proto/pb"
	"net"
)
type FConstructor func()IClientProxy
func NewConnect(addr string) net.Conn{
	conn, _ := net.Dial("tcp", addr)
	return conn
}
func MakeComponentProxyList(componentList []*pb.ADD_ENGINE_COMPONENT , IService IService) [] IClientProxy{
	l := len(componentList)
	proxyList := make([]IClientProxy , l)
	for i,component := range componentList{
		conn := NewConnect(component.ListenAddr)
		proxyList[i] = IService.NewTcpConnection(conn)
	}
	return proxyList
}
func MakeCenterProxy(addr string , IService IService)IClientProxy{
	if addr == "" {
		addr = consts.CenterAddr
	}
	conn := NewConnect(addr)
	proxy := IService.NewTcpConnection(conn)
	return proxy
}