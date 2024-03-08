// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: draft/ydb_object_storage_v1.proto

package Ydb_ObjectStorage_V1

import (
	context "context"
	Ydb_ObjectStorage "github.com/ydb-platform/ydb-go-genproto/draft/protos/Ydb_ObjectStorage"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ObjectStorageService_List_FullMethodName = "/Ydb.ObjectStorage.V1.ObjectStorageService/List"
)

// ObjectStorageServiceClient is the client API for ObjectStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectStorageServiceClient interface {
	List(ctx context.Context, in *Ydb_ObjectStorage.ListingRequest, opts ...grpc.CallOption) (*Ydb_ObjectStorage.ListingResponse, error)
}

type objectStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectStorageServiceClient(cc grpc.ClientConnInterface) ObjectStorageServiceClient {
	return &objectStorageServiceClient{cc}
}

func (c *objectStorageServiceClient) List(ctx context.Context, in *Ydb_ObjectStorage.ListingRequest, opts ...grpc.CallOption) (*Ydb_ObjectStorage.ListingResponse, error) {
	out := new(Ydb_ObjectStorage.ListingResponse)
	err := c.cc.Invoke(ctx, ObjectStorageService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStorageServiceServer is the server API for ObjectStorageService service.
// All implementations must embed UnimplementedObjectStorageServiceServer
// for forward compatibility
type ObjectStorageServiceServer interface {
	List(context.Context, *Ydb_ObjectStorage.ListingRequest) (*Ydb_ObjectStorage.ListingResponse, error)
	mustEmbedUnimplementedObjectStorageServiceServer()
}

// UnimplementedObjectStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectStorageServiceServer struct {
}

func (UnimplementedObjectStorageServiceServer) List(context.Context, *Ydb_ObjectStorage.ListingRequest) (*Ydb_ObjectStorage.ListingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedObjectStorageServiceServer) mustEmbedUnimplementedObjectStorageServiceServer() {}

// UnsafeObjectStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectStorageServiceServer will
// result in compilation errors.
type UnsafeObjectStorageServiceServer interface {
	mustEmbedUnimplementedObjectStorageServiceServer()
}

func RegisterObjectStorageServiceServer(s grpc.ServiceRegistrar, srv ObjectStorageServiceServer) {
	s.RegisterService(&ObjectStorageService_ServiceDesc, srv)
}

func _ObjectStorageService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_ObjectStorage.ListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStorageServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectStorageService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStorageServiceServer).List(ctx, req.(*Ydb_ObjectStorage.ListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectStorageService_ServiceDesc is the grpc.ServiceDesc for ObjectStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.ObjectStorage.V1.ObjectStorageService",
	HandlerType: (*ObjectStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ObjectStorageService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "draft/ydb_object_storage_v1.proto",
}