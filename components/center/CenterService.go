package center

import (
	"github.com/ouczbs/tree/components/base"
	"github.com/ouczbs/tree/engine/config"
	"github.com/ouczbs/tree/engine/gwutil"
	"github.com/ouczbs/tree/engine/netutil"
	"github.com/ouczbs/tree/engine/proto"
	"net"
)
type Component = string
type ComponentId = uint32
var (
	loginList = make(map[ComponentId]Component)
	gateList = make(map[ComponentId]Component)
	gameList = make(map[ComponentId]Component)
	dispatcherList = make(map[ComponentId]Component)
	sequence ComponentId = 0
)

func InscSequence()ComponentId{
	sequence++
	return sequence
}
// DispatcherService implements the dispatcher service
type UCenterService struct {
	messageQueue          chan proto.UMessage
	config                config.CenterConfig
}

func NewCenterService() *UCenterService {
	return &UCenterService{
		config: *(config.Center),
		messageQueue:          make(chan proto.UMessage, 10000),
	}
	//gameList = append(gameList, "ee")
}
func (service *UCenterService)handClientDisconnect(ccp * UCenterClientProxy){

}
func (service *UCenterService) Run() {
	service.initService()
	go gwutil.RepeatUntilPanicless(service.MessageLoop)
	netutil.ServeTCPForever(service.config.ListenAddr, service)
}

// ServeTCPConnection handles dispatcher client connections to dispatcher
func (service *UCenterService) ServeTCPConnection(conn net.Conn) {
	service.NewTcpConnection(conn)
}
func (service *UCenterService) NewTcpConnection(conn net.Conn) base.IClientProxy{
	client := newCenterClientProxy(service, conn)
	client.Serve()
	return client
}