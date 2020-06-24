package proto

import (
	"github.com/ouczbs/tree/engine/netutil"
	"github.com/ouczbs/tree/engine/proto/pb"
	protoLib "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"strings"
)

type (
	TMessageType = uint16
	TCallId = uint32
	TCode = uint32
	TFlag = bool
	TCmd = uint32
	TEnum = int32

	IReflectMessage = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage = protoreflect.ProtoMessage

	UWrapMessage = pb.WrapMessage
	UConnection = netutil.Connection
	UPacket = netutil.Packet
)
type FRequestHandle func(IRequestProxy,* URequest)

const (
	MT_INVALID = iota
	MT_TO_CLIENT
	MT_TO_CENTER
	MT_FROM_CENTER
	MT_TO_LOGIN
	MT_TO_GAME_START = 100 + iota

	MT_TO_GAME_END = 1000 + iota

	CodeOk = iota
	CodeError
)
var (
	sequence TCallId = 0
	pbMessageMaps map[TCmd]map[TCallId]func(* PbConnection,IReflectMessage)
	Marshal = protoLib.Marshal
	Unmarshal = protoLib.Unmarshal
	reqMessageTypeMaps map[TCmd]IReflectMessageType
	packageData * UPackageData

	RequestPool * URequestPool
)
func RecvPbWrapMessage(bytes []byte) * UWrapMessage{
	message := &UWrapMessage{}
	Unmarshal(bytes , message)
	return message
}
func RecvPbMessage(bytes []byte , pb IReflectMessage){
	Unmarshal(bytes , pb)
}
func newPbMessage(cmd TCmd, IPackageData IPackageData) (IReflectMessage , error){
	pbMessageTypes := IPackageData.GetIReflectTypeMap()
	messageType := pbMessageTypes[cmd]
	if messageType != nil {
		return messageType.New().Interface() ,nil
	}
	CommandMap := IPackageData.GetCommandMap()
	pbName := CommandMap[TEnum(cmd)]
	pbName = strings.Replace(pbName, "MT_" , "pb." , 1)
	messageType,err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pbName))
	if err != nil {
		return nil , err
	}
	pbMessageTypes[cmd] = messageType
	return messageType.New().Interface() ,err
}
func RegisterDownHandle(cmd TCmd,callId  TCallId,handle func(* PbConnection,IReflectMessage)){
	pbMessageHandles := pbMessageMaps[cmd]
	if pbMessageHandles == nil {
		pbMessageHandles = make(map[TCallId]func(* PbConnection,IReflectMessage))
		pbMessageMaps[cmd] = pbMessageHandles
	}
	pbMessageHandles[callId] = handle
}
func RegisterRequestHandle(cmd TCmd ,handle FRequestHandle){
	reqHandleMaps := packageData.GetHandleMap()
	reqHandleMaps[cmd] = handle
}
func unRegisterRequestHandle(cmd TCmd){
	reqHandleMaps := packageData.GetHandleMap()
	delete(reqHandleMaps , cmd)
}

func GetRequestHandle(cmd TCmd, IPackageData IPackageData) FRequestHandle {
	reqHandleMaps := IPackageData.GetHandleMap()
	handle := reqHandleMaps[cmd]
	return handle
}
func GetRequestMessage(wrap * UWrapMessage ,pbMessage IReflectMessage)* URequest{
	RecvPbMessage(wrap.Content, pbMessage)
	request := RequestPool.Pop()
	request.ProtoMessage = pbMessage
	request.Code = wrap.Code
	request.Next = true
	request.request = wrap.Request
	return request
}
func init(){
	pbMessageMaps = make(map[TCmd]map[TCallId]func(* PbConnection,IReflectMessage))
	RequestPool = NewRequestPool(16)
	packageData = NewPackageData()

}
func InscSequence()TCallId{
	sequence++
	return sequence
}