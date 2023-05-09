// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/service/pb/rpc.proto

package service

import (
	request "github.com/infraboard/mcube/http/request"
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

type DescribeBy int32

const (
	// 通过service id查询应用详情
	DescribeBy_SERVICE_ID DescribeBy = 0
	// 通过service name查询应用详情
	DescribeBy_SERVICE_NAME DescribeBy = 1
	// 通过service client_id查询应用详情
	DescribeBy_SERVICE_CLIENT_ID DescribeBy = 2
)

// Enum value maps for DescribeBy.
var (
	DescribeBy_name = map[int32]string{
		0: "SERVICE_ID",
		1: "SERVICE_NAME",
		2: "SERVICE_CLIENT_ID",
	}
	DescribeBy_value = map[string]int32{
		"SERVICE_ID":        0,
		"SERVICE_NAME":      1,
		"SERVICE_CLIENT_ID": 2,
	}
)

func (x DescribeBy) Enum() *DescribeBy {
	p := new(DescribeBy)
	*p = x
	return p
}

func (x DescribeBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeBy) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_service_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DescribeBy) Type() protoreflect.EnumType {
	return &file_mcenter_apps_service_pb_rpc_proto_enumTypes[0]
}

func (x DescribeBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeBy.Descriptor instead.
func (DescribeBy) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type QueryGitlabProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gitlab服务地址
	// @gotags: json:"address"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address"`
	// 访问Token
	// @gotags: json:"token" validate:"required"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token" validate:"required"`
}

func (x *QueryGitlabProjectRequest) Reset() {
	*x = QueryGitlabProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGitlabProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGitlabProjectRequest) ProtoMessage() {}

func (x *QueryGitlabProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGitlabProjectRequest.ProtoReflect.Descriptor instead.
func (*QueryGitlabProjectRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryGitlabProjectRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryGitlabProjectRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// ValidateCredentialRequest 校验服务凭证
type ValidateCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务客户端ID
	// @gotags: json:"client_id" yaml:"client_id" validate:"required,lte=100"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id" yaml:"client_id" validate:"required,lte=100"`
	// 服务客户端凭证
	// @gotags: json:"client_secret" yaml:"client_secret" validate:"required,lte=100"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret" yaml:"client_secret" validate:"required,lte=100"`
}

func (x *ValidateCredentialRequest) Reset() {
	*x = ValidateCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCredentialRequest) ProtoMessage() {}

func (x *ValidateCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCredentialRequest.ProtoReflect.Descriptor instead.
func (*ValidateCredentialRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateCredentialRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ValidateCredentialRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

// QueryMicroRequest 查询应用列表
type QueryServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 服务类型
	// @gotags: json:"type" yaml:"type"
	Type Type `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.mcenter.service.Type" json:"type" yaml:"type"`
	// 服务代码SSH仓库地址
	// @gotags: json:"repository_ssh_urls" yaml:"repository_ssh_urls"
	RepositorySshUrls []string `protobuf:"bytes,3,rep,name=repository_ssh_urls,json=repositorySshUrls,proto3" json:"repository_ssh_urls" yaml:"repository_ssh_urls"`
	// 名称关键字搜索
	// @gotags: json:"keywords" yaml:"keywords"
	Keywords string `protobuf:"bytes,15,opt,name=keywords,proto3" json:"keywords" yaml:"keywords"`
}

func (x *QueryServiceRequest) Reset() {
	*x = QueryServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryServiceRequest) ProtoMessage() {}

func (x *QueryServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryServiceRequest.ProtoReflect.Descriptor instead.
func (*QueryServiceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *QueryServiceRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryServiceRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_SOURCE_CODE
}

func (x *QueryServiceRequest) GetRepositorySshUrls() []string {
	if x != nil {
		return x.RepositorySshUrls
	}
	return nil
}

func (x *QueryServiceRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

// DescribeMicroRequest 查询应用详情
type DescribeServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询详情的方式
	// @gotags: json:"describe_by" yaml:"describe_by"
	DescribeBy DescribeBy `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=infraboard.mcenter.service.DescribeBy" json:"describe_by" yaml:"describe_by"`
	// 服务客户端Id
	// @gotags: json:"client_id" yaml:"client_id"
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id" yaml:"client_id"`
	// 服务名称
	// @gotags: json:"name" yaml:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" yaml:"name"`
	// 服务Id
	// @gotags: json:"id" yaml:"id"
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (x *DescribeServiceRequest) Reset() {
	*x = DescribeServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeServiceRequest) ProtoMessage() {}

func (x *DescribeServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeServiceRequest.ProtoReflect.Descriptor instead.
func (*DescribeServiceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeServiceRequest) GetDescribeBy() DescribeBy {
	if x != nil {
		return x.DescribeBy
	}
	return DescribeBy_SERVICE_ID
}

func (x *DescribeServiceRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *DescribeServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeServiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteMicroRequest todo
type DeleteServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务Id
	// @gotags: json:"id" yaml:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
}

func (x *DeleteServiceRequest) Reset() {
	*x = DeleteServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceRequest) ProtoMessage() {}

func (x *DeleteServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_service_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceRequest.ProtoReflect.Descriptor instead.
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteServiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mcenter_apps_service_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_service_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x18, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d, 0x0a,
	0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xcf, 0x01, 0x0a,
	0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x73, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x73, 0x68, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xa2,
	0x01, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x45, 0x0a, 0x0a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x44,
	0x10, 0x02, 0x32, 0xc1, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x70, 0x0a, 0x12, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x12, 0x35, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0c,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x6a, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x73, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_service_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_service_pb_rpc_proto_rawDescData = file_mcenter_apps_service_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_service_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_service_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_service_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_service_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_service_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_service_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_service_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mcenter_apps_service_pb_rpc_proto_goTypes = []interface{}{
	(DescribeBy)(0),                   // 0: infraboard.mcenter.service.DescribeBy
	(*QueryGitlabProjectRequest)(nil), // 1: infraboard.mcenter.service.QueryGitlabProjectRequest
	(*ValidateCredentialRequest)(nil), // 2: infraboard.mcenter.service.ValidateCredentialRequest
	(*QueryServiceRequest)(nil),       // 3: infraboard.mcenter.service.QueryServiceRequest
	(*DescribeServiceRequest)(nil),    // 4: infraboard.mcenter.service.DescribeServiceRequest
	(*DeleteServiceRequest)(nil),      // 5: infraboard.mcenter.service.DeleteServiceRequest
	(*request.PageRequest)(nil),       // 6: infraboard.mcube.page.PageRequest
	(Type)(0),                         // 7: infraboard.mcenter.service.Type
	(*Service)(nil),                   // 8: infraboard.mcenter.service.Service
	(*ServiceSet)(nil),                // 9: infraboard.mcenter.service.ServiceSet
}
var file_mcenter_apps_service_pb_rpc_proto_depIdxs = []int32{
	6, // 0: infraboard.mcenter.service.QueryServiceRequest.page:type_name -> infraboard.mcube.page.PageRequest
	7, // 1: infraboard.mcenter.service.QueryServiceRequest.type:type_name -> infraboard.mcenter.service.Type
	0, // 2: infraboard.mcenter.service.DescribeServiceRequest.describe_by:type_name -> infraboard.mcenter.service.DescribeBy
	2, // 3: infraboard.mcenter.service.RPC.ValidateCredential:input_type -> infraboard.mcenter.service.ValidateCredentialRequest
	3, // 4: infraboard.mcenter.service.RPC.QueryService:input_type -> infraboard.mcenter.service.QueryServiceRequest
	4, // 5: infraboard.mcenter.service.RPC.DescribeService:input_type -> infraboard.mcenter.service.DescribeServiceRequest
	1, // 6: infraboard.mcenter.service.RPC.QueryGitlabProject:input_type -> infraboard.mcenter.service.QueryGitlabProjectRequest
	8, // 7: infraboard.mcenter.service.RPC.ValidateCredential:output_type -> infraboard.mcenter.service.Service
	9, // 8: infraboard.mcenter.service.RPC.QueryService:output_type -> infraboard.mcenter.service.ServiceSet
	8, // 9: infraboard.mcenter.service.RPC.DescribeService:output_type -> infraboard.mcenter.service.Service
	9, // 10: infraboard.mcenter.service.RPC.QueryGitlabProject:output_type -> infraboard.mcenter.service.ServiceSet
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mcenter_apps_service_pb_rpc_proto_init() }
func file_mcenter_apps_service_pb_rpc_proto_init() {
	if File_mcenter_apps_service_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_service_pb_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_service_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGitlabProjectRequest); i {
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
		file_mcenter_apps_service_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCredentialRequest); i {
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
		file_mcenter_apps_service_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryServiceRequest); i {
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
		file_mcenter_apps_service_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeServiceRequest); i {
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
		file_mcenter_apps_service_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceRequest); i {
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
			RawDescriptor: file_mcenter_apps_service_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_service_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_service_pb_rpc_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_service_pb_rpc_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_service_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_service_pb_rpc_proto = out.File
	file_mcenter_apps_service_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_service_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_service_pb_rpc_proto_depIdxs = nil
}
