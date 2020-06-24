package login

import (
	"github.com/ouczbs/tree/components/base"
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/proto"
	proto2 "github.com/ouczbs/tree/engine/proto"
	"github.com/ouczbs/tree/engine/proto/pb"
)

func (service *ULoginService) MessageLoop() {
	for {
		select {
		case msg := <-service.messageQueue:
			proxy := msg.Proxy
			messageType := msg.MessageType
			packet := msg.Packet
			if messageType > proto.MT_TO_GAME_START && messageType<proto.MT_TO_GAME_END {
				service.ForwardToGame(proxy , packet)
				break
			}
			proxy.HandleMessage(packet)
			break
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *ULoginService) AddEngineComponentAck(_ proto.IRequestProxy,request * proto.URequest){
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		gwlog.Debugf("AddEngineComponentAck parse data error: %s ", ok)
	}
	if request.Code == proto.CodeError{
		return
	}
	dispatcherProxyList = base.MakeComponentProxyList(message.ComponentList , service)
	gwlog.Debugf("%s", message)
}
func (service *ULoginService) initConfig() {
	//config := service.config
	//debug.SetGCPercent(1000)
	//binutil.SetupGWLog("UloginService", config.LogLevel, config.LogFile, config.LogStderr)
	//binutil.SetupHTTPServer(config.HTTPAddr, nil)
}
func (service *ULoginService) initDownHandles() {
	proto.RegisterRequestHandle(proto2.TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK), service.AddEngineComponentAck)
}
func (service *ULoginService) initService() {
	service.initConfig()
	service.initDownHandles()
}
func (service *ULoginService) ConnectToCenter(){
	centerProxy = base.MakeCenterProxy("" , service)
	request := proto.RequestPool.Pop()
	if request == nil {
	}
	message := &pb.ADD_ENGINE_COMPONENT{}
	message.Type = pb.COMPONENT_TYPE_LOGIN
	message.ListenAddr = service.config.ListenAddr
	centerProxy.SendPbMessage(message, request)
	proto.RequestPool.Push(request)
}
func (service * ULoginService) ForwardToGame(proxy proto.IRequestProxy,packet * proto.UPacket){
	nums := len(dispatcherProxyList)
	if nums == 0 {
		gwlog.Debugf("ForwardToGame: dispatcherProxyList's num is zero , please register dispatcher")
		return
	}
	login := proxy.(* ULoginClientProxy)
	packet.AppendUint32(login.entityId)
	id := int(login.entityId) % nums
	dispatcherProxyList[id].ForwardPacket(packet)
}