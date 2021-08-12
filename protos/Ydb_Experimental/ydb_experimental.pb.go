// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/ydb_experimental.proto

package Ydb_Experimental

import (
	Ydb "github.com/YandexDatabase/ydb-go-genproto/protos/Ydb"
	Ydb_Issue "github.com/YandexDatabase/ydb-go-genproto/protos/Ydb_Issue"
	Ydb_Operations "github.com/YandexDatabase/ydb-go-genproto/protos/Ydb_Operations"
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

type ExecuteStreamQueryRequest_ProfileMode int32

const (
	ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED ExecuteStreamQueryRequest_ProfileMode = 0
	ExecuteStreamQueryRequest_NONE                     ExecuteStreamQueryRequest_ProfileMode = 1
	ExecuteStreamQueryRequest_BASIC                    ExecuteStreamQueryRequest_ProfileMode = 2
	ExecuteStreamQueryRequest_FULL                     ExecuteStreamQueryRequest_ProfileMode = 3
)

// Enum value maps for ExecuteStreamQueryRequest_ProfileMode.
var (
	ExecuteStreamQueryRequest_ProfileMode_name = map[int32]string{
		0: "PROFILE_MODE_UNSPECIFIED",
		1: "NONE",
		2: "BASIC",
		3: "FULL",
	}
	ExecuteStreamQueryRequest_ProfileMode_value = map[string]int32{
		"PROFILE_MODE_UNSPECIFIED": 0,
		"NONE":                     1,
		"BASIC":                    2,
		"FULL":                     3,
	}
)

func (x ExecuteStreamQueryRequest_ProfileMode) Enum() *ExecuteStreamQueryRequest_ProfileMode {
	p := new(ExecuteStreamQueryRequest_ProfileMode)
	*p = x
	return p
}

func (x ExecuteStreamQueryRequest_ProfileMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecuteStreamQueryRequest_ProfileMode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ydb_experimental_proto_enumTypes[0].Descriptor()
}

func (ExecuteStreamQueryRequest_ProfileMode) Type() protoreflect.EnumType {
	return &file_protos_ydb_experimental_proto_enumTypes[0]
}

func (x ExecuteStreamQueryRequest_ProfileMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecuteStreamQueryRequest_ProfileMode.Descriptor instead.
func (ExecuteStreamQueryRequest_ProfileMode) EnumDescriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{3, 0}
}

type UploadRowsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table           string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Rows            *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"` // Must be List of Structs
	OperationParams *Ydb_Operations.OperationParams `protobuf:"bytes,3,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
}

func (x *UploadRowsRequest) Reset() {
	*x = UploadRowsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRowsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRowsRequest) ProtoMessage() {}

func (x *UploadRowsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRowsRequest.ProtoReflect.Descriptor instead.
func (*UploadRowsRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRowsRequest) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *UploadRowsRequest) GetRows() *Ydb.TypedValue {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *UploadRowsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if x != nil {
		return x.OperationParams
	}
	return nil
}

type UploadRowsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *UploadRowsResponse) Reset() {
	*x = UploadRowsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRowsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRowsResponse) ProtoMessage() {}

func (x *UploadRowsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRowsResponse.ProtoReflect.Descriptor instead.
func (*UploadRowsResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{1}
}

func (x *UploadRowsResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type UploadRowsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadRowsResult) Reset() {
	*x = UploadRowsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRowsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRowsResult) ProtoMessage() {}

func (x *UploadRowsResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRowsResult.ProtoReflect.Descriptor instead.
func (*UploadRowsResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{2}
}

type ExecuteStreamQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YqlText     string                                `protobuf:"bytes,1,opt,name=yql_text,json=yqlText,proto3" json:"yql_text,omitempty"`
	Parameters  map[string]*Ydb.TypedValue            `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProfileMode ExecuteStreamQueryRequest_ProfileMode `protobuf:"varint,3,opt,name=profile_mode,json=profileMode,proto3,enum=Ydb.Experimental.ExecuteStreamQueryRequest_ProfileMode" json:"profile_mode,omitempty"`
	Explain     bool                                  `protobuf:"varint,4,opt,name=explain,proto3" json:"explain,omitempty"`
}

func (x *ExecuteStreamQueryRequest) Reset() {
	*x = ExecuteStreamQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteStreamQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteStreamQueryRequest) ProtoMessage() {}

func (x *ExecuteStreamQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteStreamQueryRequest.ProtoReflect.Descriptor instead.
func (*ExecuteStreamQueryRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteStreamQueryRequest) GetYqlText() string {
	if x != nil {
		return x.YqlText
	}
	return ""
}

func (x *ExecuteStreamQueryRequest) GetParameters() map[string]*Ydb.TypedValue {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *ExecuteStreamQueryRequest) GetProfileMode() ExecuteStreamQueryRequest_ProfileMode {
	if x != nil {
		return x.ProfileMode
	}
	return ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED
}

func (x *ExecuteStreamQueryRequest) GetExplain() bool {
	if x != nil {
		return x.Explain
	}
	return false
}

type ExecuteStreamQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Ydb.StatusIds_StatusCode  `protobuf:"varint,1,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues []*Ydb_Issue.IssueMessage `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	Result *ExecuteStreamQueryResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExecuteStreamQueryResponse) Reset() {
	*x = ExecuteStreamQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteStreamQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteStreamQueryResponse) ProtoMessage() {}

func (x *ExecuteStreamQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteStreamQueryResponse.ProtoReflect.Descriptor instead.
func (*ExecuteStreamQueryResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteStreamQueryResponse) GetStatus() Ydb.StatusIds_StatusCode {
	if x != nil {
		return x.Status
	}
	return Ydb.StatusIds_StatusCode(0)
}

func (x *ExecuteStreamQueryResponse) GetIssues() []*Ydb_Issue.IssueMessage {
	if x != nil {
		return x.Issues
	}
	return nil
}

func (x *ExecuteStreamQueryResponse) GetResult() *ExecuteStreamQueryResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type StreamQueryProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamQueryProgress) Reset() {
	*x = StreamQueryProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamQueryProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamQueryProgress) ProtoMessage() {}

func (x *StreamQueryProgress) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamQueryProgress.ProtoReflect.Descriptor instead.
func (*StreamQueryProgress) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{5}
}

type ExecuteStreamQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*ExecuteStreamQueryResult_ResultSet
	//	*ExecuteStreamQueryResult_Profile
	//	*ExecuteStreamQueryResult_Progress
	//	*ExecuteStreamQueryResult_QueryPlan
	Result isExecuteStreamQueryResult_Result `protobuf_oneof:"result"`
}

func (x *ExecuteStreamQueryResult) Reset() {
	*x = ExecuteStreamQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteStreamQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteStreamQueryResult) ProtoMessage() {}

func (x *ExecuteStreamQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteStreamQueryResult.ProtoReflect.Descriptor instead.
func (*ExecuteStreamQueryResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{6}
}

func (m *ExecuteStreamQueryResult) GetResult() isExecuteStreamQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ExecuteStreamQueryResult) GetResultSet() *Ydb.ResultSet {
	if x, ok := x.GetResult().(*ExecuteStreamQueryResult_ResultSet); ok {
		return x.ResultSet
	}
	return nil
}

func (x *ExecuteStreamQueryResult) GetProfile() string {
	if x, ok := x.GetResult().(*ExecuteStreamQueryResult_Profile); ok {
		return x.Profile
	}
	return ""
}

func (x *ExecuteStreamQueryResult) GetProgress() *StreamQueryProgress {
	if x, ok := x.GetResult().(*ExecuteStreamQueryResult_Progress); ok {
		return x.Progress
	}
	return nil
}

func (x *ExecuteStreamQueryResult) GetQueryPlan() string {
	if x, ok := x.GetResult().(*ExecuteStreamQueryResult_QueryPlan); ok {
		return x.QueryPlan
	}
	return ""
}

type isExecuteStreamQueryResult_Result interface {
	isExecuteStreamQueryResult_Result()
}

type ExecuteStreamQueryResult_ResultSet struct {
	ResultSet *Ydb.ResultSet `protobuf:"bytes,1,opt,name=result_set,json=resultSet,proto3,oneof"`
}

type ExecuteStreamQueryResult_Profile struct {
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3,oneof"`
}

type ExecuteStreamQueryResult_Progress struct {
	Progress *StreamQueryProgress `protobuf:"bytes,3,opt,name=progress,proto3,oneof"`
}

type ExecuteStreamQueryResult_QueryPlan struct {
	QueryPlan string `protobuf:"bytes,4,opt,name=query_plan,json=queryPlan,proto3,oneof"`
}

func (*ExecuteStreamQueryResult_ResultSet) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Profile) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Progress) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_QueryPlan) isExecuteStreamQueryResult_Result() {}

type GetDiskSpaceUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationParams *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Database        string                          `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *GetDiskSpaceUsageRequest) Reset() {
	*x = GetDiskSpaceUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiskSpaceUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiskSpaceUsageRequest) ProtoMessage() {}

func (x *GetDiskSpaceUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiskSpaceUsageRequest.ProtoReflect.Descriptor instead.
func (*GetDiskSpaceUsageRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{7}
}

func (x *GetDiskSpaceUsageRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if x != nil {
		return x.OperationParams
	}
	return nil
}

func (x *GetDiskSpaceUsageRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type GetDiskSpaceUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *GetDiskSpaceUsageResponse) Reset() {
	*x = GetDiskSpaceUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiskSpaceUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiskSpaceUsageResponse) ProtoMessage() {}

func (x *GetDiskSpaceUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiskSpaceUsageResponse.ProtoReflect.Descriptor instead.
func (*GetDiskSpaceUsageResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{8}
}

func (x *GetDiskSpaceUsageResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type GetDiskSpaceUsageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId    string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	FolderId   string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	DatabaseId string `protobuf:"bytes,3,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	// in bytes
	TotalSize uint64 `protobuf:"varint,4,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	DataSize  uint64 `protobuf:"varint,5,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	IndexSize uint64 `protobuf:"varint,6,opt,name=index_size,json=indexSize,proto3" json:"index_size,omitempty"`
}

func (x *GetDiskSpaceUsageResult) Reset() {
	*x = GetDiskSpaceUsageResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_experimental_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiskSpaceUsageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiskSpaceUsageResult) ProtoMessage() {}

func (x *GetDiskSpaceUsageResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_experimental_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiskSpaceUsageResult.ProtoReflect.Descriptor instead.
func (*GetDiskSpaceUsageResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_experimental_proto_rawDescGZIP(), []int{9}
}

func (x *GetDiskSpaceUsageResult) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *GetDiskSpaceUsageResult) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *GetDiskSpaceUsageResult) GetDatabaseId() string {
	if x != nil {
		return x.DatabaseId
	}
	return ""
}

func (x *GetDiskSpaceUsageResult) GetTotalSize() uint64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *GetDiskSpaceUsageResult) GetDataSize() uint64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *GetDiskSpaceUsageResult) GetIndexSize() uint64 {
	if x != nil {
		return x.IndexSize
	}
	return 0
}

var File_protos_ydb_experimental_proto protoreflect.FileDescriptor

var file_protos_ydb_experimental_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x1a, 0x17, 0x79, 0x64, 0x62, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x79, 0x64, 0x62, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x79, 0x64, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x79, 0x64, 0x62, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x4a, 0x0a, 0x10, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4d, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x6f,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa5, 0x03, 0x0a, 0x19, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x79, 0x71, 0x6c, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x79, 0x71, 0x6c, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x5b, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5a,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x1a, 0x4e, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x03,
	0x22, 0xc4, 0x01, 0x0a, 0x1a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x64, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0xd7,
	0x01, 0x0a, 0x18, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x54, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xcd, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x53,
	0x69, 0x7a, 0x65, 0x42, 0x77, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x42, 0x12, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0x5f, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_ydb_experimental_proto_rawDescOnce sync.Once
	file_protos_ydb_experimental_proto_rawDescData = file_protos_ydb_experimental_proto_rawDesc
)

func file_protos_ydb_experimental_proto_rawDescGZIP() []byte {
	file_protos_ydb_experimental_proto_rawDescOnce.Do(func() {
		file_protos_ydb_experimental_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_experimental_proto_rawDescData)
	})
	return file_protos_ydb_experimental_proto_rawDescData
}

var file_protos_ydb_experimental_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ydb_experimental_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protos_ydb_experimental_proto_goTypes = []interface{}{
	(ExecuteStreamQueryRequest_ProfileMode)(0), // 0: Ydb.Experimental.ExecuteStreamQueryRequest.ProfileMode
	(*UploadRowsRequest)(nil),                  // 1: Ydb.Experimental.UploadRowsRequest
	(*UploadRowsResponse)(nil),                 // 2: Ydb.Experimental.UploadRowsResponse
	(*UploadRowsResult)(nil),                   // 3: Ydb.Experimental.UploadRowsResult
	(*ExecuteStreamQueryRequest)(nil),          // 4: Ydb.Experimental.ExecuteStreamQueryRequest
	(*ExecuteStreamQueryResponse)(nil),         // 5: Ydb.Experimental.ExecuteStreamQueryResponse
	(*StreamQueryProgress)(nil),                // 6: Ydb.Experimental.StreamQueryProgress
	(*ExecuteStreamQueryResult)(nil),           // 7: Ydb.Experimental.ExecuteStreamQueryResult
	(*GetDiskSpaceUsageRequest)(nil),           // 8: Ydb.Experimental.GetDiskSpaceUsageRequest
	(*GetDiskSpaceUsageResponse)(nil),          // 9: Ydb.Experimental.GetDiskSpaceUsageResponse
	(*GetDiskSpaceUsageResult)(nil),            // 10: Ydb.Experimental.GetDiskSpaceUsageResult
	nil,                                        // 11: Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry
	(*Ydb.TypedValue)(nil),                     // 12: Ydb.TypedValue
	(*Ydb_Operations.OperationParams)(nil),     // 13: Ydb.Operations.OperationParams
	(*Ydb_Operations.Operation)(nil),           // 14: Ydb.Operations.Operation
	(Ydb.StatusIds_StatusCode)(0),              // 15: Ydb.StatusIds.StatusCode
	(*Ydb_Issue.IssueMessage)(nil),             // 16: Ydb.Issue.IssueMessage
	(*Ydb.ResultSet)(nil),                      // 17: Ydb.ResultSet
}
var file_protos_ydb_experimental_proto_depIdxs = []int32{
	12, // 0: Ydb.Experimental.UploadRowsRequest.rows:type_name -> Ydb.TypedValue
	13, // 1: Ydb.Experimental.UploadRowsRequest.operation_params:type_name -> Ydb.Operations.OperationParams
	14, // 2: Ydb.Experimental.UploadRowsResponse.operation:type_name -> Ydb.Operations.Operation
	11, // 3: Ydb.Experimental.ExecuteStreamQueryRequest.parameters:type_name -> Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry
	0,  // 4: Ydb.Experimental.ExecuteStreamQueryRequest.profile_mode:type_name -> Ydb.Experimental.ExecuteStreamQueryRequest.ProfileMode
	15, // 5: Ydb.Experimental.ExecuteStreamQueryResponse.status:type_name -> Ydb.StatusIds.StatusCode
	16, // 6: Ydb.Experimental.ExecuteStreamQueryResponse.issues:type_name -> Ydb.Issue.IssueMessage
	7,  // 7: Ydb.Experimental.ExecuteStreamQueryResponse.result:type_name -> Ydb.Experimental.ExecuteStreamQueryResult
	17, // 8: Ydb.Experimental.ExecuteStreamQueryResult.result_set:type_name -> Ydb.ResultSet
	6,  // 9: Ydb.Experimental.ExecuteStreamQueryResult.progress:type_name -> Ydb.Experimental.StreamQueryProgress
	13, // 10: Ydb.Experimental.GetDiskSpaceUsageRequest.operation_params:type_name -> Ydb.Operations.OperationParams
	14, // 11: Ydb.Experimental.GetDiskSpaceUsageResponse.operation:type_name -> Ydb.Operations.Operation
	12, // 12: Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry.value:type_name -> Ydb.TypedValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_protos_ydb_experimental_proto_init() }
func file_protos_ydb_experimental_proto_init() {
	if File_protos_ydb_experimental_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_experimental_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRowsRequest); i {
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
		file_protos_ydb_experimental_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRowsResponse); i {
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
		file_protos_ydb_experimental_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRowsResult); i {
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
		file_protos_ydb_experimental_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteStreamQueryRequest); i {
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
		file_protos_ydb_experimental_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteStreamQueryResponse); i {
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
		file_protos_ydb_experimental_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamQueryProgress); i {
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
		file_protos_ydb_experimental_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteStreamQueryResult); i {
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
		file_protos_ydb_experimental_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiskSpaceUsageRequest); i {
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
		file_protos_ydb_experimental_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiskSpaceUsageResponse); i {
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
		file_protos_ydb_experimental_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiskSpaceUsageResult); i {
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
	file_protos_ydb_experimental_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ExecuteStreamQueryResult_ResultSet)(nil),
		(*ExecuteStreamQueryResult_Profile)(nil),
		(*ExecuteStreamQueryResult_Progress)(nil),
		(*ExecuteStreamQueryResult_QueryPlan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_ydb_experimental_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_experimental_proto_goTypes,
		DependencyIndexes: file_protos_ydb_experimental_proto_depIdxs,
		EnumInfos:         file_protos_ydb_experimental_proto_enumTypes,
		MessageInfos:      file_protos_ydb_experimental_proto_msgTypes,
	}.Build()
	File_protos_ydb_experimental_proto = out.File
	file_protos_ydb_experimental_proto_rawDesc = nil
	file_protos_ydb_experimental_proto_goTypes = nil
	file_protos_ydb_experimental_proto_depIdxs = nil
}