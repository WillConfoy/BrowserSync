// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: node.proto

package node

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
	SyncService_SendClickInternal_FullMethodName   = "/nodeproto.SyncService/SendClickInternal"
	SyncService_SendKeyDownInternal_FullMethodName = "/nodeproto.SyncService/SendKeyDownInternal"
	SyncService_SendCommandInternal_FullMethodName = "/nodeproto.SyncService/SendCommandInternal"
	SyncService_SendKeyUpInternal_FullMethodName   = "/nodeproto.SyncService/SendKeyUpInternal"
	SyncService_HeartbeatInternal_FullMethodName   = "/nodeproto.SyncService/HeartbeatInternal"
	SyncService_SendScrollInternal_FullMethodName  = "/nodeproto.SyncService/SendScrollInternal"
	SyncService_UpdateLeader_FullMethodName        = "/nodeproto.SyncService/UpdateLeader"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	SendClickInternal(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ClickResponse, error)
	SendKeyDownInternal(ctx context.Context, in *KeyDownRequest, opts ...grpc.CallOption) (*KeyDownResponse, error)
	SendCommandInternal(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	SendKeyUpInternal(ctx context.Context, in *KeyUpRequest, opts ...grpc.CallOption) (*KeyUpResponse, error)
	HeartbeatInternal(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	SendScrollInternal(ctx context.Context, in *ScrollRequest, opts ...grpc.CallOption) (*ScrollResponse, error)
	UpdateLeader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SendClickInternal(ctx context.Context, in *ClickRequest, opts ...grpc.CallOption) (*ClickResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClickResponse)
	err := c.cc.Invoke(ctx, SyncService_SendClickInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SendKeyDownInternal(ctx context.Context, in *KeyDownRequest, opts ...grpc.CallOption) (*KeyDownResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KeyDownResponse)
	err := c.cc.Invoke(ctx, SyncService_SendKeyDownInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SendCommandInternal(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, SyncService_SendCommandInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SendKeyUpInternal(ctx context.Context, in *KeyUpRequest, opts ...grpc.CallOption) (*KeyUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KeyUpResponse)
	err := c.cc.Invoke(ctx, SyncService_SendKeyUpInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) HeartbeatInternal(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, SyncService_HeartbeatInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SendScrollInternal(ctx context.Context, in *ScrollRequest, opts ...grpc.CallOption) (*ScrollResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScrollResponse)
	err := c.cc.Invoke(ctx, SyncService_SendScrollInternal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdateLeader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaderResponse)
	err := c.cc.Invoke(ctx, SyncService_UpdateLeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility.
type SyncServiceServer interface {
	SendClickInternal(context.Context, *ClickRequest) (*ClickResponse, error)
	SendKeyDownInternal(context.Context, *KeyDownRequest) (*KeyDownResponse, error)
	SendCommandInternal(context.Context, *CommandRequest) (*CommandResponse, error)
	SendKeyUpInternal(context.Context, *KeyUpRequest) (*KeyUpResponse, error)
	HeartbeatInternal(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	SendScrollInternal(context.Context, *ScrollRequest) (*ScrollResponse, error)
	UpdateLeader(context.Context, *LeaderRequest) (*LeaderResponse, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSyncServiceServer struct{}

func (UnimplementedSyncServiceServer) SendClickInternal(context.Context, *ClickRequest) (*ClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClickInternal not implemented")
}
func (UnimplementedSyncServiceServer) SendKeyDownInternal(context.Context, *KeyDownRequest) (*KeyDownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKeyDownInternal not implemented")
}
func (UnimplementedSyncServiceServer) SendCommandInternal(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandInternal not implemented")
}
func (UnimplementedSyncServiceServer) SendKeyUpInternal(context.Context, *KeyUpRequest) (*KeyUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKeyUpInternal not implemented")
}
func (UnimplementedSyncServiceServer) HeartbeatInternal(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartbeatInternal not implemented")
}
func (UnimplementedSyncServiceServer) SendScrollInternal(context.Context, *ScrollRequest) (*ScrollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendScrollInternal not implemented")
}
func (UnimplementedSyncServiceServer) UpdateLeader(context.Context, *LeaderRequest) (*LeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeader not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}
func (UnimplementedSyncServiceServer) testEmbeddedByValue()                     {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	// If the following call pancis, it indicates UnimplementedSyncServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_SendClickInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendClickInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendClickInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendClickInternal(ctx, req.(*ClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SendKeyDownInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendKeyDownInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendKeyDownInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendKeyDownInternal(ctx, req.(*KeyDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SendCommandInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendCommandInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendCommandInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendCommandInternal(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SendKeyUpInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendKeyUpInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendKeyUpInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendKeyUpInternal(ctx, req.(*KeyUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_HeartbeatInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).HeartbeatInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_HeartbeatInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).HeartbeatInternal(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SendScrollInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendScrollInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendScrollInternal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendScrollInternal(ctx, req.(*ScrollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdateLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdateLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_UpdateLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdateLeader(ctx, req.(*LeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nodeproto.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClickInternal",
			Handler:    _SyncService_SendClickInternal_Handler,
		},
		{
			MethodName: "SendKeyDownInternal",
			Handler:    _SyncService_SendKeyDownInternal_Handler,
		},
		{
			MethodName: "SendCommandInternal",
			Handler:    _SyncService_SendCommandInternal_Handler,
		},
		{
			MethodName: "SendKeyUpInternal",
			Handler:    _SyncService_SendKeyUpInternal_Handler,
		},
		{
			MethodName: "HeartbeatInternal",
			Handler:    _SyncService_HeartbeatInternal_Handler,
		},
		{
			MethodName: "SendScrollInternal",
			Handler:    _SyncService_SendScrollInternal_Handler,
		},
		{
			MethodName: "UpdateLeader",
			Handler:    _SyncService_UpdateLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
