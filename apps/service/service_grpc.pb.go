// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: apps/service/pb/service.proto

package service

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

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaServiceClient interface {
	ValidateCredential(ctx context.Context, in *ValidateCredentialRequest, opts ...grpc.CallOption) (*Service, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error)
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error)
	RefreshCredential(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error)
}

type metaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaServiceClient(cc grpc.ClientConnInterface) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) ValidateCredential(ctx context.Context, in *ValidateCredentialRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/ValidateCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/UpdateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error) {
	out := new(ServiceSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/QueryService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/DescribeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) RefreshCredential(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.service.MetaService/RefreshCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServiceServer is the server API for MetaService service.
// All implementations must embed UnimplementedMetaServiceServer
// for forward compatibility
type MetaServiceServer interface {
	ValidateCredential(context.Context, *ValidateCredentialRequest) (*Service, error)
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*Service, error)
	QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error)
	DescribeService(context.Context, *DescribeServiceRequest) (*Service, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*Service, error)
	RefreshCredential(context.Context, *DescribeServiceRequest) (*Service, error)
	mustEmbedUnimplementedMetaServiceServer()
}

// UnimplementedMetaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (UnimplementedMetaServiceServer) ValidateCredential(context.Context, *ValidateCredentialRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredential not implemented")
}
func (UnimplementedMetaServiceServer) CreateService(context.Context, *CreateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedMetaServiceServer) UpdateService(context.Context, *UpdateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}
func (UnimplementedMetaServiceServer) QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}
func (UnimplementedMetaServiceServer) DescribeService(context.Context, *DescribeServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedMetaServiceServer) DeleteService(context.Context, *DeleteServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedMetaServiceServer) RefreshCredential(context.Context, *DescribeServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCredential not implemented")
}
func (UnimplementedMetaServiceServer) mustEmbedUnimplementedMetaServiceServer() {}

// UnsafeMetaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaServiceServer will
// result in compilation errors.
type UnsafeMetaServiceServer interface {
	mustEmbedUnimplementedMetaServiceServer()
}

func RegisterMetaServiceServer(s grpc.ServiceRegistrar, srv MetaServiceServer) {
	s.RegisterService(&MetaService_ServiceDesc, srv)
}

func _MetaService_ValidateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).ValidateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/ValidateCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).ValidateCredential(ctx, req.(*ValidateCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_QueryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).QueryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/QueryService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).QueryService(ctx, req.(*QueryServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/DescribeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_RefreshCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).RefreshCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.service.MetaService/RefreshCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).RefreshCredential(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaService_ServiceDesc is the grpc.ServiceDesc for MetaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.service.MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCredential",
			Handler:    _MetaService_ValidateCredential_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _MetaService_CreateService_Handler,
		},
		{
			MethodName: "UpdateService",
			Handler:    _MetaService_UpdateService_Handler,
		},
		{
			MethodName: "QueryService",
			Handler:    _MetaService_QueryService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _MetaService_DescribeService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _MetaService_DeleteService_Handler,
		},
		{
			MethodName: "RefreshCredential",
			Handler:    _MetaService_RefreshCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/service/pb/service.proto",
}
