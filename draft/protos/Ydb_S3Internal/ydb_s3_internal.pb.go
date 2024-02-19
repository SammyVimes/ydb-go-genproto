// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: draft/protos/ydb_s3_internal.proto

package Ydb_S3Internal

import (
	Ydb "github.com/ydb-platform/ydb-go-genproto/protos/Ydb"
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
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

type S3ListingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName           string                          `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	KeyPrefix           *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"` // A tuple representing all key columns that preceed path column
	PathColumnPrefix    string                          `protobuf:"bytes,3,opt,name=path_column_prefix,json=pathColumnPrefix,proto3" json:"path_column_prefix,omitempty"`
	PathColumnDelimiter string                          `protobuf:"bytes,4,opt,name=path_column_delimiter,json=pathColumnDelimiter,proto3" json:"path_column_delimiter,omitempty"`
	StartAfterKeySuffix *Ydb.TypedValue                 `protobuf:"bytes,5,opt,name=start_after_key_suffix,json=startAfterKeySuffix,proto3" json:"start_after_key_suffix,omitempty"` // A tuple representing key columns that succeed path column
	MaxKeys             uint32                          `protobuf:"varint,6,opt,name=max_keys,json=maxKeys,proto3" json:"max_keys,omitempty"`
	ColumnsToReturn     []string                        `protobuf:"bytes,7,rep,name=columns_to_return,json=columnsToReturn,proto3" json:"columns_to_return,omitempty"`
	OperationParams     *Ydb_Operations.OperationParams `protobuf:"bytes,8,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
}

func (x *S3ListingRequest) Reset() {
	*x = S3ListingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ListingRequest) ProtoMessage() {}

func (x *S3ListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ListingRequest.ProtoReflect.Descriptor instead.
func (*S3ListingRequest) Descriptor() ([]byte, []int) {
	return file_draft_protos_ydb_s3_internal_proto_rawDescGZIP(), []int{0}
}

func (x *S3ListingRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *S3ListingRequest) GetKeyPrefix() *Ydb.TypedValue {
	if x != nil {
		return x.KeyPrefix
	}
	return nil
}

func (x *S3ListingRequest) GetPathColumnPrefix() string {
	if x != nil {
		return x.PathColumnPrefix
	}
	return ""
}

func (x *S3ListingRequest) GetPathColumnDelimiter() string {
	if x != nil {
		return x.PathColumnDelimiter
	}
	return ""
}

func (x *S3ListingRequest) GetStartAfterKeySuffix() *Ydb.TypedValue {
	if x != nil {
		return x.StartAfterKeySuffix
	}
	return nil
}

func (x *S3ListingRequest) GetMaxKeys() uint32 {
	if x != nil {
		return x.MaxKeys
	}
	return 0
}

func (x *S3ListingRequest) GetColumnsToReturn() []string {
	if x != nil {
		return x.ColumnsToReturn
	}
	return nil
}

func (x *S3ListingRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if x != nil {
		return x.OperationParams
	}
	return nil
}

type S3ListingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *S3ListingResponse) Reset() {
	*x = S3ListingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ListingResponse) ProtoMessage() {}

func (x *S3ListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ListingResponse.ProtoReflect.Descriptor instead.
func (*S3ListingResponse) Descriptor() ([]byte, []int) {
	return file_draft_protos_ydb_s3_internal_proto_rawDescGZIP(), []int{1}
}

func (x *S3ListingResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type S3ListingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonPrefixes *Ydb.ResultSet `protobuf:"bytes,1,opt,name=common_prefixes,json=commonPrefixes,proto3" json:"common_prefixes,omitempty"` // Every Contents row starts with key suffix with KeySuffixSize columns
	Contents       *Ydb.ResultSet `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`                                   // Every Contents row starts with key suffix with KeySuffixSize columns
}

func (x *S3ListingResult) Reset() {
	*x = S3ListingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ListingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ListingResult) ProtoMessage() {}

func (x *S3ListingResult) ProtoReflect() protoreflect.Message {
	mi := &file_draft_protos_ydb_s3_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ListingResult.ProtoReflect.Descriptor instead.
func (*S3ListingResult) Descriptor() ([]byte, []int) {
	return file_draft_protos_ydb_s3_internal_proto_rawDescGZIP(), []int{2}
}

func (x *S3ListingResult) GetCommonPrefixes() *Ydb.ResultSet {
	if x != nil {
		return x.CommonPrefixes
	}
	return nil
}

func (x *S3ListingResult) GetContents() *Ydb.ResultSet {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_draft_protos_ydb_s3_internal_proto protoreflect.FileDescriptor

var file_draft_protos_ydb_s3_internal_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79,
	0x64, 0x62, 0x5f, 0x73, 0x33, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x33, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x03, 0x0a, 0x10, 0x53, 0x33, 0x4c,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0a,
	0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x74, 0x68, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x61,
	0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x74, 0x68, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x13, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x53, 0x75,
	0x66, 0x66, 0x69, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x4b, 0x65, 0x79, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x4a, 0x0a, 0x10, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x11, 0x53, 0x33, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x0f, 0x53, 0x33, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65,
	0x74, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65,
	0x73, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x6a, 0x0a,
	0x20, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x64, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x73, 0x33, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64,
	0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x53, 0x33, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_draft_protos_ydb_s3_internal_proto_rawDescOnce sync.Once
	file_draft_protos_ydb_s3_internal_proto_rawDescData = file_draft_protos_ydb_s3_internal_proto_rawDesc
)

func file_draft_protos_ydb_s3_internal_proto_rawDescGZIP() []byte {
	file_draft_protos_ydb_s3_internal_proto_rawDescOnce.Do(func() {
		file_draft_protos_ydb_s3_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_draft_protos_ydb_s3_internal_proto_rawDescData)
	})
	return file_draft_protos_ydb_s3_internal_proto_rawDescData
}

var file_draft_protos_ydb_s3_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_draft_protos_ydb_s3_internal_proto_goTypes = []interface{}{
	(*S3ListingRequest)(nil),               // 0: Ydb.S3Internal.S3ListingRequest
	(*S3ListingResponse)(nil),              // 1: Ydb.S3Internal.S3ListingResponse
	(*S3ListingResult)(nil),                // 2: Ydb.S3Internal.S3ListingResult
	(*Ydb.TypedValue)(nil),                 // 3: Ydb.TypedValue
	(*Ydb_Operations.OperationParams)(nil), // 4: Ydb.Operations.OperationParams
	(*Ydb_Operations.Operation)(nil),       // 5: Ydb.Operations.Operation
	(*Ydb.ResultSet)(nil),                  // 6: Ydb.ResultSet
}
var file_draft_protos_ydb_s3_internal_proto_depIdxs = []int32{
	3, // 0: Ydb.S3Internal.S3ListingRequest.key_prefix:type_name -> Ydb.TypedValue
	3, // 1: Ydb.S3Internal.S3ListingRequest.start_after_key_suffix:type_name -> Ydb.TypedValue
	4, // 2: Ydb.S3Internal.S3ListingRequest.operation_params:type_name -> Ydb.Operations.OperationParams
	5, // 3: Ydb.S3Internal.S3ListingResponse.operation:type_name -> Ydb.Operations.Operation
	6, // 4: Ydb.S3Internal.S3ListingResult.common_prefixes:type_name -> Ydb.ResultSet
	6, // 5: Ydb.S3Internal.S3ListingResult.contents:type_name -> Ydb.ResultSet
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_draft_protos_ydb_s3_internal_proto_init() }
func file_draft_protos_ydb_s3_internal_proto_init() {
	if File_draft_protos_ydb_s3_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_draft_protos_ydb_s3_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ListingRequest); i {
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
		file_draft_protos_ydb_s3_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ListingResponse); i {
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
		file_draft_protos_ydb_s3_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ListingResult); i {
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
			RawDescriptor: file_draft_protos_ydb_s3_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_draft_protos_ydb_s3_internal_proto_goTypes,
		DependencyIndexes: file_draft_protos_ydb_s3_internal_proto_depIdxs,
		MessageInfos:      file_draft_protos_ydb_s3_internal_proto_msgTypes,
	}.Build()
	File_draft_protos_ydb_s3_internal_proto = out.File
	file_draft_protos_ydb_s3_internal_proto_rawDesc = nil
	file_draft_protos_ydb_s3_internal_proto_goTypes = nil
	file_draft_protos_ydb_s3_internal_proto_depIdxs = nil
}
