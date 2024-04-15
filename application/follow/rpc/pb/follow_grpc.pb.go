// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: follow.proto

package pb

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

const (
	Follow_Follow_FullMethodName     = "/pb.Follow/Follow"
	Follow_UnFollow_FullMethodName   = "/pb.Follow/UnFollow"
	Follow_FollowList_FullMethodName = "/pb.Follow/FollowList"
	Follow_FansList_FullMethodName   = "/pb.Follow/FansList"
)

// FollowClient is the client API for Follow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowClient interface {
	// 关注
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	// 取消关注
	UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error)
	// 关注列表
	FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
	// 粉丝列表
	FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error)
}

type followClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowClient(cc grpc.ClientConnInterface) FollowClient {
	return &followClient{cc}
}

func (c *followClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, Follow_Follow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) UnFollow(ctx context.Context, in *UnFollowRequest, opts ...grpc.CallOption) (*UnFollowResponse, error) {
	out := new(UnFollowResponse)
	err := c.cc.Invoke(ctx, Follow_UnFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	out := new(FollowListResponse)
	err := c.cc.Invoke(ctx, Follow_FollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) FansList(ctx context.Context, in *FansListRequest, opts ...grpc.CallOption) (*FansListResponse, error) {
	out := new(FansListResponse)
	err := c.cc.Invoke(ctx, Follow_FansList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServer is the server API for Follow service.
// All implementations must embed UnimplementedFollowServer
// for forward compatibility
type FollowServer interface {
	// 关注
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	// 取消关注
	UnFollow(context.Context, *UnFollowRequest) (*UnFollowResponse, error)
	// 关注列表
	FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error)
	// 粉丝列表
	FansList(context.Context, *FansListRequest) (*FansListResponse, error)
	mustEmbedUnimplementedFollowServer()
}

// UnimplementedFollowServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServer struct {
}

func (UnimplementedFollowServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowServer) UnFollow(context.Context, *UnFollowRequest) (*UnFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollow not implemented")
}
func (UnimplementedFollowServer) FollowList(context.Context, *FollowListRequest) (*FollowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowList not implemented")
}
func (UnimplementedFollowServer) FansList(context.Context, *FansListRequest) (*FansListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FansList not implemented")
}
func (UnimplementedFollowServer) mustEmbedUnimplementedFollowServer() {}

// UnsafeFollowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServer will
// result in compilation errors.
type UnsafeFollowServer interface {
	mustEmbedUnimplementedFollowServer()
}

func RegisterFollowServer(s grpc.ServiceRegistrar, srv FollowServer) {
	s.RegisterService(&Follow_ServiceDesc, srv)
}

func _Follow_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_UnFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).UnFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_UnFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).UnFollow(ctx, req.(*UnFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_FollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).FollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_FollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).FollowList(ctx, req.(*FollowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_FansList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FansListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).FansList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_FansList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).FansList(ctx, req.(*FansListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Follow_ServiceDesc is the grpc.ServiceDesc for Follow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Follow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Follow",
	HandlerType: (*FollowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _Follow_Follow_Handler,
		},
		{
			MethodName: "UnFollow",
			Handler:    _Follow_UnFollow_Handler,
		},
		{
			MethodName: "FollowList",
			Handler:    _Follow_FollowList_Handler,
		},
		{
			MethodName: "FansList",
			Handler:    _Follow_FansList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
