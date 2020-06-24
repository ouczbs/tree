package center

import (
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/proto"
	proto2 "github.com/ouczbs/tree/engine/proto"
	"github.com/ouczbs/tree/engine/proto/pb"
)
func (service *UCenterService) MessageLoop(){
	for {
		select {
		case msg := <-service.messageQueue:
			pbc := msg.Proxy
			messageType := msg.MessageType
			switch messageType {
			case proto.MT_TO_CENTER:
				pkt := msg.Packet
				pbc.HandleMessage(pkt)
				break
			default:
				gwlog.Debugf("recive %s " , messageType)
			}
			break
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:

		}
	}
}
func (service *UCenterService) UtilAddEngineComponentList(proxy proto.IRequestProxy,request * proto.URequest, list map[ComponentId]Component) {
	result := &pb.ADD_ENGINE_COMPONENT_ACK{}
	if list != nil {
		for id,comp := range list{
			component := &pb.ADD_ENGINE_COMPONENT{ComponentId: id , ListenAddr: comp}
			result.ComponentList = append(result.ComponentList, component)
		}
	}
	result.ComponentId = InscSequence()
	proxy.ResponseMessage(result,request)
}
// 中心服务器执行一对一服务 type FRequestHandle = func(IRequestProxy,* RequestMessage)
func (service *UCenterService) AddEngineComponent(proxy proto.IRequestProxy,request * proto.URequest){
	message ,ok := request.ProtoMessage.(* pb.ADD_ENGINE_COMPONENT)
	if !ok {
		gwlog.Debugf("AddEngineComponent parse data error: %s "  , ok)
	}
	request.Cmd = proto2.TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)
	request.MessageType = proto.MT_FROM_CENTER
	switch message.Type{
	case pb.COMPONENT_TYPE_DISPATCHER:
		service.UtilAddEngineComponentList(proxy,request, nil)
		break
	case pb.COMPONENT_TYPE_GAME:
		service.UtilAddEngineComponentList(proxy,request, dispatcherList)
		break
	case pb.COMPONENT_TYPE_GATE:
		service.UtilAddEngineComponentList(proxy,request , gateList)
		break
	case pb.COMPONENT_TYPE_LOGIN:
		service.UtilAddEngineComponentList(proxy ,request , dispatcherList)
		break
	default:
		gwlog.Debugf("accept unKnow type %s" , message.Type)
	}
	gwlog.Debugf("%s" , message)
}
func (service *UCenterService) initConfig(){
	//config := service.config
	//debug.SetGCPercent(1000)
	//binutil.SetupGWLog("CenterService", config.LogLevel, config.LogFile, config.LogStderr)
	//binutil.SetupHTTPServer(config.HTTPAddr, nil)
}
func (service *UCenterService) initDownHandles(){
	proto.RegisterRequestHandle(proto2.TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT), service.AddEngineComponent)
}
func (service *UCenterService) initService(){
	service.initConfig()
	service.initDownHandles()
}
