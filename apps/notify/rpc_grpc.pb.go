// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: apps/notify/pb/rpc.proto

package notify

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 短信通知
	SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
	// 语音通知
	SendVoice(ctx context.Context, in *SendVoiceRequest, opts ...grpc.CallOption) (*SendVoiceResponse, error)
	// 邮件通知
	SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error)
	// IM通知
	SendIM(ctx context.Context, in *SendIMRequest, opts ...grpc.CallOption) (*SendImResponse, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	out := new(SendSmsResponse)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.notify.RPC/SendSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) SendVoice(ctx context.Context, in *SendVoiceRequest, opts ...grpc.CallOption) (*SendVoiceResponse, error) {
	out := new(SendVoiceResponse)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.notify.RPC/SendVoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) SendMail(ctx context.Context, in *SendMailRequest, opts ...grpc.CallOption) (*SendMailResponse, error) {
	out := new(SendMailResponse)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.notify.RPC/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) SendIM(ctx context.Context, in *SendIMRequest, opts ...grpc.CallOption) (*SendImResponse, error) {
	out := new(SendImResponse)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.notify.RPC/SendIM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 短信通知
	SendSMS(context.Context, *SendSMSRequest) (*SendSmsResponse, error)
	// 语音通知
	SendVoice(context.Context, *SendVoiceRequest) (*SendVoiceResponse, error)
	// 邮件通知
	SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error)
	// IM通知
	SendIM(context.Context, *SendIMRequest) (*SendImResponse, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) SendSMS(context.Context, *SendSMSRequest) (*SendSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMS not implemented")
}
func (UnimplementedRPCServer) SendVoice(context.Context, *SendVoiceRequest) (*SendVoiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVoice not implemented")
}
func (UnimplementedRPCServer) SendMail(context.Context, *SendMailRequest) (*SendMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedRPCServer) SendIM(context.Context, *SendIMRequest) (*SendImResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendIM not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.notify.RPC/SendSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).SendSMS(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_SendVoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendVoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).SendVoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.notify.RPC/SendVoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).SendVoice(ctx, req.(*SendVoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.notify.RPC/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).SendMail(ctx, req.(*SendMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_SendIM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendIMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).SendIM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.notify.RPC/SendIM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).SendIM(ctx, req.(*SendIMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.notify.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSMS",
			Handler:    _RPC_SendSMS_Handler,
		},
		{
			MethodName: "SendVoice",
			Handler:    _RPC_SendVoice_Handler,
		},
		{
			MethodName: "SendMail",
			Handler:    _RPC_SendMail_Handler,
		},
		{
			MethodName: "SendIM",
			Handler:    _RPC_SendIM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/notify/pb/rpc.proto",
}
