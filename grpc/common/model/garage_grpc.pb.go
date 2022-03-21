// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: garage.proto

package model 

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GaragesClient is the client API for Garages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GaragesClient interface {
	List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error)
	Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type garagesClient struct {
	cc grpc.ClientConnInterface
}

func NewGaragesClient(cc grpc.ClientConnInterface) GaragesClient {
	return &garagesClient{cc}
}

func (c *garagesClient) List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error) {
	out := new(GarageList)
	err := c.cc.Invoke(ctx, "/model.Garages/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *garagesClient) Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/model.Garages/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaragesServer is the server API for Garages service.
// All implementations must embed UnimplementedGaragesServer
// for forward compatibility
type GaragesServer interface {
	List(context.Context, *GarageUserId) (*GarageList, error)
	Add(context.Context, *GarageAndUserId) (*emptypb.Empty, error)
	// mustEmbedUnimplementedGaragesServer()
}

// UnimplementedGaragesServer must be embedded to have forward compatible implementations.
type UnimplementedGaragesServer struct {
}

func (UnimplementedGaragesServer) List(context.Context, *GarageUserId) (*GarageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGaragesServer) Add(context.Context, *GarageAndUserId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGaragesServer) mustEmbedUnimplementedGaragesServer() {}

// UnsafeGaragesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GaragesServer will
// result in compilation errors.
type UnsafeGaragesServer interface {
	mustEmbedUnimplementedGaragesServer()
}

func RegisterGaragesServer(s grpc.ServiceRegistrar, srv GaragesServer) {
	s.RegisterService(&Garages_ServiceDesc, srv)
}

func _Garages_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).List(ctx, req.(*GarageUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Garages_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageAndUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).Add(ctx, req.(*GarageAndUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// Garages_ServiceDesc is the grpc.ServiceDesc for Garages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Garages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Garages",
	HandlerType: (*GaragesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Garages_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Garages_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garage.proto",
}
