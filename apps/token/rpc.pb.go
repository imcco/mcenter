// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/token/pb/rpc.proto

package token

import (
	user "github.com/infraboard/mcenter/apps/user"
	request "github.com/infraboard/mcube/v2/http/request"
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

type DESCRIBY_BY int32

const (
	// 通过access token查看
	DESCRIBY_BY_ACCESS_TOKEN DESCRIBY_BY = 0
	// 通过刷新token查询
	DESCRIBY_BY_REFRESH_TOKEN DESCRIBY_BY = 1
)

// Enum value maps for DESCRIBY_BY.
var (
	DESCRIBY_BY_name = map[int32]string{
		0: "ACCESS_TOKEN",
		1: "REFRESH_TOKEN",
	}
	DESCRIBY_BY_value = map[string]int32{
		"ACCESS_TOKEN":  0,
		"REFRESH_TOKEN": 1,
	}
)

func (x DESCRIBY_BY) Enum() *DESCRIBY_BY {
	p := new(DESCRIBY_BY)
	*p = x
	return p
}

func (x DESCRIBY_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBY_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_token_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBY_BY) Type() protoreflect.EnumType {
	return &file_mcenter_apps_token_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBY_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBY_BY.Descriptor instead.
func (DESCRIBY_BY) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RevolkTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	// 刷新令牌
	// @gotags: json:"refresh_token"
	RefreshToken string `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
}

func (x *RevolkTokenRequest) Reset() {
	*x = RevolkTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevolkTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevolkTokenRequest) ProtoMessage() {}

func (x *RevolkTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevolkTokenRequest.ProtoReflect.Descriptor instead.
func (*RevolkTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RevolkTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RevolkTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ChangeNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要切换空间令牌
	// @gotags: json:"token" validate:"required"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" validate:"required"`
	// 空间名称
	// @gotags: json:"namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" validate:"required"`
}

func (x *ChangeNamespaceRequest) Reset() {
	*x = ChangeNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNamespaceRequest) ProtoMessage() {}

func (x *ChangeNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNamespaceRequest.ProtoReflect.Descriptor instead.
func (*ChangeNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeNamespaceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ChangeNamespaceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type QueryTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 颁发平台, 根据授权方式判断
	// @gotags: bson:"platform" json:"platform"
	Platform *PLATFORM `protobuf:"varint,2,opt,name=platform,proto3,enum=infraboard.mcenter.token.PLATFORM,oneof" json:"platform" bson:"platform"`
	// 访问令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
	// 刷新令牌, 当访问令牌过期时, 可以刷新令牌进行刷新
	// @gotags: json:"refresh_token"
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token"`
	// 用户类型
	// @gotags: json:"user_type"
	UserType *user.TYPE `protobuf:"varint,5,opt,name=user_type,json=userType,proto3,enum=infraboard.mcenter.user.TYPE,oneof" json:"user_type"`
	// 用户当前所处域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain"`
	// 用户名
	// @gotags: json:"username"
	Username string `protobuf:"bytes,7,opt,name=username,proto3" json:"username"`
	// 用户Id
	// @gotags: json:"user_id"
	UserId string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id"`
	// 授权类型
	// @gotags: json:"grant_type"
	GrantType *GRANT_TYPE `protobuf:"varint,9,opt,name=grant_type,json=grantType,proto3,enum=infraboard.mcenter.token.GRANT_TYPE,oneof" json:"grant_type"`
	// 令牌类型
	// @gotags: json:"type"
	Type *TOKEN_TYPE `protobuf:"varint,10,opt,name=type,proto3,enum=infraboard.mcenter.token.TOKEN_TYPE,oneof" json:"type"`
	// 当前空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,11,opt,name=namespace,proto3" json:"namespace"`
	// 令牌描述信息, 当授权类型为Private Token时, 做描述使用
	// @gotags: json:"description"
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description"`
	// 是否冻结
	// @gotags: json:"is_block"
	IsBlock *bool `protobuf:"varint,13,opt,name=is_block,json=isBlock,proto3,oneof" json:"is_block"`
	// 冻结类型
	// @gotags: json:"block_type"
	BlockType *BLOCK_TYPE `protobuf:"varint,14,opt,name=block_type,json=blockType,proto3,enum=infraboard.mcenter.token.BLOCK_TYPE,oneof" json:"block_type"`
}

func (x *QueryTokenRequest) Reset() {
	*x = QueryTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTokenRequest) ProtoMessage() {}

func (x *QueryTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTokenRequest.ProtoReflect.Descriptor instead.
func (*QueryTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *QueryTokenRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryTokenRequest) GetPlatform() PLATFORM {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return PLATFORM_WEB
}

func (x *QueryTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *QueryTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *QueryTokenRequest) GetUserType() user.TYPE {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return user.TYPE(0)
}

func (x *QueryTokenRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryTokenRequest) GetGrantType() GRANT_TYPE {
	if x != nil && x.GrantType != nil {
		return *x.GrantType
	}
	return GRANT_TYPE_PASSWORD
}

func (x *QueryTokenRequest) GetType() TOKEN_TYPE {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TOKEN_TYPE_BEARER
}

func (x *QueryTokenRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryTokenRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QueryTokenRequest) GetIsBlock() bool {
	if x != nil && x.IsBlock != nil {
		return *x.IsBlock
	}
	return false
}

func (x *QueryTokenRequest) GetBlockType() BLOCK_TYPE {
	if x != nil && x.BlockType != nil {
		return *x.BlockType
	}
	return BLOCK_TYPE_REFRESH_TOKEN_EXPIRED
}

type DescribeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 参数类型
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBY_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=infraboard.mcenter.token.DESCRIBY_BY" json:"describe_by"`
	// 参数值
	// @gotags: json:"describe_value" validate:"required"
	DescribeValue string `protobuf:"bytes,2,opt,name=describe_value,json=describeValue,proto3" json:"describe_value" validate:"required"`
}

func (x *DescribeTokenRequest) Reset() {
	*x = DescribeTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTokenRequest) ProtoMessage() {}

func (x *DescribeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTokenRequest.ProtoReflect.Descriptor instead.
func (*DescribeTokenRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTokenRequest) GetDescribeBy() DESCRIBY_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBY_BY_ACCESS_TOKEN
}

func (x *DescribeTokenRequest) GetDescribeValue() string {
	if x != nil {
		return x.DescribeValue
	}
	return ""
}

// IssueCodeResponse todo
type IssueCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 颁发后返回的消息, 比如以发送到xxx手机
	// @gotags: json:"message"
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
}

func (x *IssueCodeResponse) Reset() {
	*x = IssueCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCodeResponse) ProtoMessage() {}

func (x *IssueCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCodeResponse.ProtoReflect.Descriptor instead.
func (*IssueCodeResponse) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *IssueCodeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// VerifyCodeRequest 验证码校验请求
type VerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"username" validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" validate:"required"`
	// 验证码
	// @gotags: json:"code" validate:"required"
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code" validate:"required"`
}

func (x *VerifyCodeRequest) Reset() {
	*x = VerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyCodeRequest) ProtoMessage() {}

func (x *VerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_token_pb_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VerifyCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_mcenter_apps_token_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_token_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x18, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x14, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5c, 0x0a, 0x12, 0x52, 0x65, 0x76, 0x6f, 0x6c, 0x6b,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x22, 0xe8, 0x05, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x43, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x48,
	0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0a, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x52, 0x41, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x48, 0x02, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x48, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x48, 0x05,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x67, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x85, 0x01,
	0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x59, 0x5f,
	0x42, 0x59, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x25,
	0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x32, 0x0a, 0x0b, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x59, 0x5f, 0x42, 0x59, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45,
	0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x32, 0xc4, 0x01,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x62, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_token_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_token_pb_rpc_proto_rawDescData = file_mcenter_apps_token_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_token_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_token_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_token_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_token_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_token_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_token_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_token_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mcenter_apps_token_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBY_BY)(0),               // 0: infraboard.mcenter.token.DESCRIBY_BY
	(*ValidateTokenRequest)(nil),   // 1: infraboard.mcenter.token.ValidateTokenRequest
	(*RevolkTokenRequest)(nil),     // 2: infraboard.mcenter.token.RevolkTokenRequest
	(*ChangeNamespaceRequest)(nil), // 3: infraboard.mcenter.token.ChangeNamespaceRequest
	(*QueryTokenRequest)(nil),      // 4: infraboard.mcenter.token.QueryTokenRequest
	(*DescribeTokenRequest)(nil),   // 5: infraboard.mcenter.token.DescribeTokenRequest
	(*IssueCodeResponse)(nil),      // 6: infraboard.mcenter.token.IssueCodeResponse
	(*VerifyCodeRequest)(nil),      // 7: infraboard.mcenter.token.VerifyCodeRequest
	(*request.PageRequest)(nil),    // 8: infraboard.mcube.page.PageRequest
	(PLATFORM)(0),                  // 9: infraboard.mcenter.token.PLATFORM
	(user.TYPE)(0),                 // 10: infraboard.mcenter.user.TYPE
	(GRANT_TYPE)(0),                // 11: infraboard.mcenter.token.GRANT_TYPE
	(TOKEN_TYPE)(0),                // 12: infraboard.mcenter.token.TOKEN_TYPE
	(BLOCK_TYPE)(0),                // 13: infraboard.mcenter.token.BLOCK_TYPE
	(*Token)(nil),                  // 14: infraboard.mcenter.token.Token
	(*Code)(nil),                   // 15: infraboard.mcenter.token.Code
}
var file_mcenter_apps_token_pb_rpc_proto_depIdxs = []int32{
	8,  // 0: infraboard.mcenter.token.QueryTokenRequest.page:type_name -> infraboard.mcube.page.PageRequest
	9,  // 1: infraboard.mcenter.token.QueryTokenRequest.platform:type_name -> infraboard.mcenter.token.PLATFORM
	10, // 2: infraboard.mcenter.token.QueryTokenRequest.user_type:type_name -> infraboard.mcenter.user.TYPE
	11, // 3: infraboard.mcenter.token.QueryTokenRequest.grant_type:type_name -> infraboard.mcenter.token.GRANT_TYPE
	12, // 4: infraboard.mcenter.token.QueryTokenRequest.type:type_name -> infraboard.mcenter.token.TOKEN_TYPE
	13, // 5: infraboard.mcenter.token.QueryTokenRequest.block_type:type_name -> infraboard.mcenter.token.BLOCK_TYPE
	0,  // 6: infraboard.mcenter.token.DescribeTokenRequest.describe_by:type_name -> infraboard.mcenter.token.DESCRIBY_BY
	1,  // 7: infraboard.mcenter.token.RPC.ValidateToken:input_type -> infraboard.mcenter.token.ValidateTokenRequest
	7,  // 8: infraboard.mcenter.token.RPC.VerifyCode:input_type -> infraboard.mcenter.token.VerifyCodeRequest
	14, // 9: infraboard.mcenter.token.RPC.ValidateToken:output_type -> infraboard.mcenter.token.Token
	15, // 10: infraboard.mcenter.token.RPC.VerifyCode:output_type -> infraboard.mcenter.token.Code
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_mcenter_apps_token_pb_rpc_proto_init() }
func file_mcenter_apps_token_pb_rpc_proto_init() {
	if File_mcenter_apps_token_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_token_pb_token_proto_init()
	file_mcenter_apps_token_pb_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevolkTokenRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNamespaceRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTokenRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTokenRequest); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCodeResponse); i {
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
		file_mcenter_apps_token_pb_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyCodeRequest); i {
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
	file_mcenter_apps_token_pb_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mcenter_apps_token_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_token_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_token_pb_rpc_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_token_pb_rpc_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_token_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_token_pb_rpc_proto = out.File
	file_mcenter_apps_token_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_token_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_token_pb_rpc_proto_depIdxs = nil
}
