package proto

type IPoolObject interface {
	Release()
	//Default()
}
type IPackageData interface {
	GetCommandMap()map[TEnum]string
	GetHandleMap()map[TCmd]FRequestHandle
	GetIReflectTypeMap()map[TCmd]IReflectMessageType
}

type IRequestProxy interface {
	HandleMessage( * UPacket)
	RequestMessage( IReflectMessage, * URequest ,  FRequestHandle)error
	ResponseMessage( IReflectMessage, * URequest)error
	SendPbMessage( IReflectMessage, * URequest)error
	ForwardPacket( * UPacket)error
}
type UMessage struct{
	Proxy 		 IRequestProxy
	MessageType  TMessageType
	Packet      *UPacket
}