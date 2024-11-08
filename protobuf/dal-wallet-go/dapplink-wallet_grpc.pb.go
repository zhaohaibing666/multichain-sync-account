// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: protobuf/dapplink-wallet.proto

package dal_wallet_go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BusinessMiddleWireServices_BusinessRegister_FullMethodName = "/syncs.BusinessMiddleWireServices/businessRegister"
)

// BusinessMiddleWireServicesClient is the client API for BusinessMiddleWireServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessMiddleWireServicesClient interface {
	BusinessRegister(ctx context.Context, in *BusinessRegisterRequest, opts ...grpc.CallOption) (*BusinessRegisterResponse, error)
}

type businessMiddleWireServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessMiddleWireServicesClient(cc grpc.ClientConnInterface) BusinessMiddleWireServicesClient {
	return &businessMiddleWireServicesClient{cc}
}

func (c *businessMiddleWireServicesClient) BusinessRegister(ctx context.Context, in *BusinessRegisterRequest, opts ...grpc.CallOption) (*BusinessRegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BusinessRegisterResponse)
	err := c.cc.Invoke(ctx, BusinessMiddleWireServices_BusinessRegister_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessMiddleWireServicesServer is the server API for BusinessMiddleWireServices service.
// All implementations should embed UnimplementedBusinessMiddleWireServicesServer
// for forward compatibility.
type BusinessMiddleWireServicesServer interface {
	BusinessRegister(context.Context, *BusinessRegisterRequest) (*BusinessRegisterResponse, error)
}

// UnimplementedBusinessMiddleWireServicesServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBusinessMiddleWireServicesServer struct{}

func (UnimplementedBusinessMiddleWireServicesServer) BusinessRegister(context.Context, *BusinessRegisterRequest) (*BusinessRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BusinessRegister not implemented")
}
func (UnimplementedBusinessMiddleWireServicesServer) testEmbeddedByValue() {}

// UnsafeBusinessMiddleWireServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessMiddleWireServicesServer will
// result in compilation errors.
type UnsafeBusinessMiddleWireServicesServer interface {
	mustEmbedUnimplementedBusinessMiddleWireServicesServer()
}

func RegisterBusinessMiddleWireServicesServer(s grpc.ServiceRegistrar, srv BusinessMiddleWireServicesServer) {
	// If the following call pancis, it indicates UnimplementedBusinessMiddleWireServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BusinessMiddleWireServices_ServiceDesc, srv)
}

func _BusinessMiddleWireServices_BusinessRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessMiddleWireServicesServer).BusinessRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessMiddleWireServices_BusinessRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessMiddleWireServicesServer).BusinessRegister(ctx, req.(*BusinessRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusinessMiddleWireServices_ServiceDesc is the grpc.ServiceDesc for BusinessMiddleWireServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusinessMiddleWireServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "syncs.BusinessMiddleWireServices",
	HandlerType: (*BusinessMiddleWireServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "businessRegister",
			Handler:    _BusinessMiddleWireServices_BusinessRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/dapplink-wallet.proto",
}
