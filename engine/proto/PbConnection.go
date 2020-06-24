package proto

import (
	"github.com/ouczbs/tree/engine/consts"
	"github.com/ouczbs/tree/engine/gwlog"
	"github.com/ouczbs/tree/engine/netutil"
	"time"
)

type PbConnection struct {
	IPackageData IPackageData

	packetConn   *netutil.PacketConnection
	autoFlushing bool
	closed       bool
	execFlag     bool
	reqHandleMaps map[TCallId]FRequestHandle
	request		 TCallId

	Packet       * UPacket
	MessageType  TMessageType
	Cmd          TCmd

	Code         TCode
}

func NewPbConnection(conn UConnection) *PbConnection {
	pbc := &PbConnection{
		packetConn: netutil.NewPacketConnection(conn),
		execFlag: false,
	}
	pbc.IPackageData = packageData
	pbc.reqHandleMaps = make(map[TCallId]FRequestHandle)
	return pbc
}
func (pbc * PbConnection) Recv(messageType *TMessageType)(*UPacket, error){
	pkt, err := pbc.packetConn.RecvPacket()
	if err != nil {
		return nil, err
	}

	*messageType = pkt.ReadUint16()
	if consts.DEBUG_PACKETS {
		gwlog.Infof("%s: Recv msgtype=%v, payload size=%d", pbc, *messageType, pkt.GetPayloadLen())
	}
	return pkt, nil
}
func (pbc *PbConnection) Close() error {
	return pbc.packetConn.Close()
}
func (pbc *PbConnection) Then(handle FRequestHandle , request * URequest) *PbConnection{
	if handle == nil {
		return pbc
	}
	if request.Next {
		request.Next = false
		handle(pbc, request)
	}
	return pbc
}
func (pbc *PbConnection) HandleMessage(packet * UPacket)  {
	wrapBytes := packet.MessagePayload()
	wrapMessage := RecvPbWrapMessage(wrapBytes)
	cmd := wrapMessage.Cmd
	handle := pbc.reqHandleMaps[wrapMessage.Response]
	if handle != nil {
		delete(pbc.reqHandleMaps , wrapMessage.Response)
	}
	globalHandle := GetRequestHandle(cmd,pbc.IPackageData)
	if cmd == 0 || (handle == nil && globalHandle == nil){
		return
	}
	pbMessage,_ := newPbMessage(cmd , pbc.IPackageData)
	request := GetRequestMessage(wrapMessage ,pbMessage)
	pbc.Then(handle , request).Then(globalHandle , request)
}
func (pbc *PbConnection) RequestMessage(message IReflectMessage,request * URequest , handle FRequestHandle)error{
	sequence = InscSequence()
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
		Request: sequence,
		Code:request.Code,
	}
	pbc.reqHandleMaps[sequence] = handle
	return pbc.sendPbMessage(message , wrap, request.MessageType)
}
func (pbc *PbConnection) ResponseMessage(message IReflectMessage,request * URequest)error{
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
		Response: request.request,
		Code:request.Code,
	}
	return pbc.sendPbMessage(message , wrap , request.MessageType)
}
func (pbc *PbConnection) sendPbMessage(message IReflectMessage, wrap *UWrapMessage , messageType TMessageType)error{
	packet := pbc.packetConn.NewPacket()
	packet.AppendUint16(messageType)
	out , err := Marshal(message)
	if err != nil{
		return err
	}
	wrap.Content = out
	buf , err := Marshal(wrap)
	if err != nil{
		return err
	}
	packet.AppendBytes(buf)
	err = pbc.packetConn.SendPacket(packet)
	packet.Release()
	return err
}
func (pbc *PbConnection) ForwardPacket(packet * UPacket)error{
	err := pbc.packetConn.SendPacket(packet)
	packet.Release()
	return err
}
func (pbc *PbConnection) SendPbMessage(message IReflectMessage,request * URequest)error{
	wrap := &UWrapMessage{
		Cmd: request.Cmd,
	}
	return pbc.sendPbMessage(message , wrap, request.MessageType)
}
// Flush connection writes
func (pbc *PbConnection) Flush(reason string) error {
	return pbc.packetConn.Flush(reason)
}
// IsClosed returns if the connection is closed
func (pbc *PbConnection) IsClosed() bool {
	return pbc.closed
}
// SetAutoFlush starts a goroutine to flush connection writes at some specified interval
func (pbc *PbConnection) SetAutoFlush(interval time.Duration) {
	if pbc.autoFlushing {
		gwlog.Panicf("%s.SetAutoFlush: already auto flushing!", pbc)
	}
	pbc.autoFlushing = true
	go func() {
		//defer gwlog.Debugf("%s: auto flush routine quited", gwc)
		for !pbc.IsClosed() {
			time.Sleep(interval)
			err := pbc.Flush("AutoFlush")
			if err != nil {
				break
			}
		}
	}()
}