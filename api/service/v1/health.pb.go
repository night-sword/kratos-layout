// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: service/v1/health.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Status int32

const (
	Status_UP   Status = 0
	Status_DOWN Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UP",
		1: "DOWN",
	}
	Status_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_service_v1_health_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_service_v1_health_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_service_v1_health_proto_rawDescGZIP(), []int{0}
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_service_v1_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.service.v1.Status" json:"status,omitempty"`
}

func (x *HealthReply) Reset() {
	*x = HealthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthReply) ProtoMessage() {}

func (x *HealthReply) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthReply.ProtoReflect.Descriptor instead.
func (*HealthReply) Descriptor() ([]byte, []int) {
	return file_service_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthReply) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UP
}

var File_service_v1_health_proto protoreflect.FileDescriptor

var file_service_v1_health_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x3d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x55,
	0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x32, 0x64, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x5a, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x42, 0x4a, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2f,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_v1_health_proto_rawDescOnce sync.Once
	file_service_v1_health_proto_rawDescData = file_service_v1_health_proto_rawDesc
)

func file_service_v1_health_proto_rawDescGZIP() []byte {
	file_service_v1_health_proto_rawDescOnce.Do(func() {
		file_service_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_v1_health_proto_rawDescData)
	})
	return file_service_v1_health_proto_rawDescData
}

var file_service_v1_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_v1_health_proto_goTypes = []interface{}{
	(Status)(0),           // 0: api.service.v1.Status
	(*HealthRequest)(nil), // 1: api.service.v1.HealthRequest
	(*HealthReply)(nil),   // 2: api.service.v1.HealthReply
}
var file_service_v1_health_proto_depIdxs = []int32{
	0, // 0: api.service.v1.HealthReply.status:type_name -> api.service.v1.Status
	1, // 1: api.service.v1.Health.HealthCheck:input_type -> api.service.v1.HealthRequest
	2, // 2: api.service.v1.Health.HealthCheck:output_type -> api.service.v1.HealthReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_v1_health_proto_init() }
func file_service_v1_health_proto_init() {
	if File_service_v1_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_v1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_service_v1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthReply); i {
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
			RawDescriptor: file_service_v1_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_v1_health_proto_goTypes,
		DependencyIndexes: file_service_v1_health_proto_depIdxs,
		EnumInfos:         file_service_v1_health_proto_enumTypes,
		MessageInfos:      file_service_v1_health_proto_msgTypes,
	}.Build()
	File_service_v1_health_proto = out.File
	file_service_v1_health_proto_rawDesc = nil
	file_service_v1_health_proto_goTypes = nil
	file_service_v1_health_proto_depIdxs = nil
}
