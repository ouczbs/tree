package proto

import "github.com/ouczbs/tree/engine/proto/pb"

type UPackageData struct {
	reqHandleMaps      map[TCmd]FRequestHandle
	pbMessageTypes map[TCmd]IReflectMessageType
}
func NewPackageData()* UPackageData{
	data := &UPackageData{
		reqHandleMaps: make(map[TCmd]FRequestHandle),
		pbMessageTypes:make(map[TCmd]IReflectMessageType),
	}
	return data
}
func (This * UPackageData) GetCommandMap()map[TEnum]string{
	return pb.CommandList_name
}
func (This * UPackageData) GetHandleMap()map[TCmd]FRequestHandle{
	return This.reqHandleMaps
}
func (This * UPackageData) GetIReflectTypeMap()map[TCmd]IReflectMessageType{
	return This.pbMessageTypes
}