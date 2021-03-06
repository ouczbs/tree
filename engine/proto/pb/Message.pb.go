// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: Message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type COMPONENT_TYPE int32

const (
	COMPONENT_TYPE_CENTER     COMPONENT_TYPE = 0
	COMPONENT_TYPE_DISPATCHER COMPONENT_TYPE = 1
	COMPONENT_TYPE_LOGIN      COMPONENT_TYPE = 2
	COMPONENT_TYPE_GATE       COMPONENT_TYPE = 3
	COMPONENT_TYPE_GAME       COMPONENT_TYPE = 4
)

// Enum value maps for COMPONENT_TYPE.
var (
	COMPONENT_TYPE_name = map[int32]string{
		0: "CENTER",
		1: "DISPATCHER",
		2: "LOGIN",
		3: "GATE",
		4: "GAME",
	}
	COMPONENT_TYPE_value = map[string]int32{
		"CENTER":     0,
		"DISPATCHER": 1,
		"LOGIN":      2,
		"GATE":       3,
		"GAME":       4,
	}
)

func (x COMPONENT_TYPE) Enum() *COMPONENT_TYPE {
	p := new(COMPONENT_TYPE)
	*p = x
	return p
}

func (x COMPONENT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (COMPONENT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_Message_proto_enumTypes[0].Descriptor()
}

func (COMPONENT_TYPE) Type() protoreflect.EnumType {
	return &file_Message_proto_enumTypes[0]
}

func (x COMPONENT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use COMPONENT_TYPE.Descriptor instead.
func (COMPONENT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

type ADD_ENGINE_COMPONENT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenAddr  string         `protobuf:"bytes,1,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	ComponentId uint32         `protobuf:"varint,2,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
	Type        COMPONENT_TYPE `protobuf:"varint,3,opt,name=type,proto3,enum=pb.COMPONENT_TYPE" json:"type,omitempty"`
}

func (x *ADD_ENGINE_COMPONENT) Reset() {
	*x = ADD_ENGINE_COMPONENT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADD_ENGINE_COMPONENT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADD_ENGINE_COMPONENT) ProtoMessage() {}

func (x *ADD_ENGINE_COMPONENT) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADD_ENGINE_COMPONENT.ProtoReflect.Descriptor instead.
func (*ADD_ENGINE_COMPONENT) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *ADD_ENGINE_COMPONENT) GetListenAddr() string {
	if x != nil {
		return x.ListenAddr
	}
	return ""
}

func (x *ADD_ENGINE_COMPONENT) GetComponentId() uint32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

func (x *ADD_ENGINE_COMPONENT) GetType() COMPONENT_TYPE {
	if x != nil {
		return x.Type
	}
	return COMPONENT_TYPE_CENTER
}

type ADD_ENGINE_COMPONENT_ACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentId   uint32                  `protobuf:"varint,1,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
	ComponentList []*ADD_ENGINE_COMPONENT `protobuf:"bytes,2,rep,name=component_list,json=componentList,proto3" json:"component_list,omitempty"`
}

func (x *ADD_ENGINE_COMPONENT_ACK) Reset() {
	*x = ADD_ENGINE_COMPONENT_ACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADD_ENGINE_COMPONENT_ACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADD_ENGINE_COMPONENT_ACK) ProtoMessage() {}

func (x *ADD_ENGINE_COMPONENT_ACK) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADD_ENGINE_COMPONENT_ACK.ProtoReflect.Descriptor instead.
func (*ADD_ENGINE_COMPONENT_ACK) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{1}
}

func (x *ADD_ENGINE_COMPONENT_ACK) GetComponentId() uint32 {
	if x != nil {
		return x.ComponentId
	}
	return 0
}

func (x *ADD_ENGINE_COMPONENT_ACK) GetComponentList() []*ADD_ENGINE_COMPONENT {
	if x != nil {
		return x.ComponentList
	}
	return nil
}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x4e, 0x47, 0x49,
	0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x18, 0x41, 0x44, 0x44, 0x5f,
	0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x43, 0x4b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2a, 0x4b, 0x0a, 0x0e, 0x43, 0x4f, 0x4d, 0x50,
	0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x45,
	0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x50, 0x41, 0x54,
	0x43, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x41, 0x4d, 0x45, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Message_proto_goTypes = []interface{}{
	(COMPONENT_TYPE)(0),              // 0: pb.COMPONENT_TYPE
	(*ADD_ENGINE_COMPONENT)(nil),     // 1: pb.ADD_ENGINE_COMPONENT
	(*ADD_ENGINE_COMPONENT_ACK)(nil), // 2: pb.ADD_ENGINE_COMPONENT_ACK
}
var file_Message_proto_depIdxs = []int32{
	0, // 0: pb.ADD_ENGINE_COMPONENT.type:type_name -> pb.COMPONENT_TYPE
	1, // 1: pb.ADD_ENGINE_COMPONENT_ACK.component_list:type_name -> pb.ADD_ENGINE_COMPONENT
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADD_ENGINE_COMPONENT); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ADD_ENGINE_COMPONENT_ACK); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		EnumInfos:         file_Message_proto_enumTypes,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}
