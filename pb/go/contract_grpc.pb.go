// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: contract.proto

package _go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldClient interface {
	CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*CreateAlbumResponse, error)
	GetAlbum(ctx context.Context, in *GetAlbumRequest, opts ...grpc.CallOption) (*GetAlbumResponse, error)
	GetAlbumById(ctx context.Context, in *GetAlbumByIdRequest, opts ...grpc.CallOption) (*GetAlbumByIdResponse, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*CreateAlbumResponse, error) {
	out := new(CreateAlbumResponse)
	err := c.cc.Invoke(ctx, "/hello.world.v1.HelloWorld/CreateAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GetAlbum(ctx context.Context, in *GetAlbumRequest, opts ...grpc.CallOption) (*GetAlbumResponse, error) {
	out := new(GetAlbumResponse)
	err := c.cc.Invoke(ctx, "/hello.world.v1.HelloWorld/GetAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GetAlbumById(ctx context.Context, in *GetAlbumByIdRequest, opts ...grpc.CallOption) (*GetAlbumByIdResponse, error) {
	out := new(GetAlbumByIdResponse)
	err := c.cc.Invoke(ctx, "/hello.world.v1.HelloWorld/GetAlbumById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloWorldServer is the server API for HelloWorld service.
// All implementations must embed UnimplementedHelloWorldServer
// for forward compatibility
type HelloWorldServer interface {
	CreateAlbum(context.Context, *CreateAlbumRequest) (*CreateAlbumResponse, error)
	GetAlbum(context.Context, *GetAlbumRequest) (*GetAlbumResponse, error)
	GetAlbumById(context.Context, *GetAlbumByIdRequest) (*GetAlbumByIdResponse, error)
	mustEmbedUnimplementedHelloWorldServer()
}

// UnimplementedHelloWorldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (UnimplementedHelloWorldServer) CreateAlbum(context.Context, *CreateAlbumRequest) (*CreateAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlbum not implemented")
}
func (UnimplementedHelloWorldServer) GetAlbum(context.Context, *GetAlbumRequest) (*GetAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbum not implemented")
}
func (UnimplementedHelloWorldServer) GetAlbumById(context.Context, *GetAlbumByIdRequest) (*GetAlbumByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumById not implemented")
}
func (UnimplementedHelloWorldServer) mustEmbedUnimplementedHelloWorldServer() {}

// UnsafeHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServer will
// result in compilation errors.
type UnsafeHelloWorldServer interface {
	mustEmbedUnimplementedHelloWorldServer()
}

func RegisterHelloWorldServer(s grpc.ServiceRegistrar, srv HelloWorldServer) {
	s.RegisterService(&HelloWorld_ServiceDesc, srv)
}

func _HelloWorld_CreateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).CreateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.world.v1.HelloWorld/CreateAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).CreateAlbum(ctx, req.(*CreateAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GetAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GetAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.world.v1.HelloWorld/GetAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GetAlbum(ctx, req.(*GetAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GetAlbumById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).GetAlbumById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.world.v1.HelloWorld/GetAlbumById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).GetAlbumById(ctx, req.(*GetAlbumByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloWorld_ServiceDesc is the grpc.ServiceDesc for HelloWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.world.v1.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlbum",
			Handler:    _HelloWorld_CreateAlbum_Handler,
		},
		{
			MethodName: "GetAlbum",
			Handler:    _HelloWorld_GetAlbum_Handler,
		},
		{
			MethodName: "GetAlbumById",
			Handler:    _HelloWorld_GetAlbumById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract.proto",
}
