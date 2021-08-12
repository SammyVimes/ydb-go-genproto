// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/ydb_common.proto

package Ydb

import (
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

type FeatureFlag_Status int32

const (
	FeatureFlag_STATUS_UNSPECIFIED FeatureFlag_Status = 0
	FeatureFlag_ENABLED            FeatureFlag_Status = 1
	FeatureFlag_DISABLED           FeatureFlag_Status = 2
)

// Enum value maps for FeatureFlag_Status.
var (
	FeatureFlag_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
	}
	FeatureFlag_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ENABLED":            1,
		"DISABLED":           2,
	}
)

func (x FeatureFlag_Status) Enum() *FeatureFlag_Status {
	p := new(FeatureFlag_Status)
	*p = x
	return p
}

func (x FeatureFlag_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureFlag_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ydb_common_proto_enumTypes[0].Descriptor()
}

func (FeatureFlag_Status) Type() protoreflect.EnumType {
	return &file_protos_ydb_common_proto_enumTypes[0]
}

func (x FeatureFlag_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureFlag_Status.Descriptor instead.
func (FeatureFlag_Status) EnumDescriptor() ([]byte, []int) {
	return file_protos_ydb_common_proto_rawDescGZIP(), []int{0, 0}
}

type FeatureFlag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeatureFlag) Reset() {
	*x = FeatureFlag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureFlag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureFlag) ProtoMessage() {}

func (x *FeatureFlag) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureFlag.ProtoReflect.Descriptor instead.
func (*FeatureFlag) Descriptor() ([]byte, []int) {
	return file_protos_ydb_common_proto_rawDescGZIP(), []int{0}
}

type CostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total amount of request units (RU), consumed by the operation.
	ConsumedUnits float64 `protobuf:"fixed64,1,opt,name=consumed_units,json=consumedUnits,proto3" json:"consumed_units,omitempty"`
}

func (x *CostInfo) Reset() {
	*x = CostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostInfo) ProtoMessage() {}

func (x *CostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostInfo.ProtoReflect.Descriptor instead.
func (*CostInfo) Descriptor() ([]byte, []int) {
	return file_protos_ydb_common_proto_rawDescGZIP(), []int{1}
}

func (x *CostInfo) GetConsumedUnits() float64 {
	if x != nil {
		return x.ConsumedUnits
	}
	return 0
}

var File_protos_ydb_common_proto protoreflect.FileDescriptor

var file_protos_ydb_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x59, 0x64, 0x62, 0x22, 0x4a,
	0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x3b, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x31, 0x0a, 0x08, 0x43, 0x6f,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x42, 0x5e, 0x0a,
	0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_ydb_common_proto_rawDescOnce sync.Once
	file_protos_ydb_common_proto_rawDescData = file_protos_ydb_common_proto_rawDesc
)

func file_protos_ydb_common_proto_rawDescGZIP() []byte {
	file_protos_ydb_common_proto_rawDescOnce.Do(func() {
		file_protos_ydb_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_common_proto_rawDescData)
	})
	return file_protos_ydb_common_proto_rawDescData
}

var file_protos_ydb_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ydb_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_ydb_common_proto_goTypes = []interface{}{
	(FeatureFlag_Status)(0), // 0: Ydb.FeatureFlag.Status
	(*FeatureFlag)(nil),     // 1: Ydb.FeatureFlag
	(*CostInfo)(nil),        // 2: Ydb.CostInfo
}
var file_protos_ydb_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_ydb_common_proto_init() }
func file_protos_ydb_common_proto_init() {
	if File_protos_ydb_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureFlag); i {
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
		file_protos_ydb_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostInfo); i {
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
			RawDescriptor: file_protos_ydb_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_common_proto_goTypes,
		DependencyIndexes: file_protos_ydb_common_proto_depIdxs,
		EnumInfos:         file_protos_ydb_common_proto_enumTypes,
		MessageInfos:      file_protos_ydb_common_proto_msgTypes,
	}.Build()
	File_protos_ydb_common_proto = out.File
	file_protos_ydb_common_proto_rawDesc = nil
	file_protos_ydb_common_proto_goTypes = nil
	file_protos_ydb_common_proto_depIdxs = nil
}