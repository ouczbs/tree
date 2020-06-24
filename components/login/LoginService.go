package login

import (
	"github.com/ouczbs/tree/engine/config"
	"github.com/ouczbs/tree/engine/gwutil"
	"github.com/ouczbs/tree/engine/netutil"
	"github.com/ouczbs/tree/engine/proto"
	"net"
)
// DispatcherService implements the dispatcher service
type loginService struct {
	messageQueue          chan proto.UMessage
	config                config.LoginConfig
}

func NewLoginService() *loginService {
	return &loginService{
		config: *(config.Login),
		messageQueue:          make(chan proto.UMessage, 10000),
	}
	//gameList = append(gameList, "ee")
}
func (service *loginService)handClientDisconnect(lcp * loginClientProxy){

}
func (service *loginService) Run() {
	service.initService()
	go gwutil.RepeatUntilPanicless(service.MessageLoop)
	netutil.ServeTCPForever(service.config.ListenAddr, service)
	service.ConnectToCenter()
}
func (service *loginService) ServeTCPConnection(conn net.Conn) {
	service.NewTcpConnection(conn)
}

func (service *loginService) NewTcpConnection(conn net.Conn) IClientProxy{
	client := newLoginClientProxy(service, conn)
	client.Serve()
	return client
}