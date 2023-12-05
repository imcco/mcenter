// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/namespace/pb/namespace.proto

package namespace

import (
	resource "github.com/infraboard/mcube/v2/pb/resource"
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

type VISIBLE int32

const (
	// 默认空间是私有的, 只对加入该空间的用户可见
	VISIBLE_PRIVATE VISIBLE = 0
	// PUBLIC  公开的空间, 对所有人可见, 所有人均可访问该空间
	VISIBLE_PUBLIC VISIBLE = 1
)

// Enum value maps for VISIBLE.
var (
	VISIBLE_name = map[int32]string{
		0: "PRIVATE",
		1: "PUBLIC",
	}
	VISIBLE_value = map[string]int32{
		"PRIVATE": 0,
		"PUBLIC":  1,
	}
)

func (x VISIBLE) Enum() *VISIBLE {
	p := new(VISIBLE)
	*p = x
	return p
}

func (x VISIBLE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VISIBLE) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_namespace_pb_namespace_proto_enumTypes[0].Descriptor()
}

func (VISIBLE) Type() protoreflect.EnumType {
	return &file_mcenter_apps_namespace_pb_namespace_proto_enumTypes[0]
}

func (x VISIBLE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VISIBLE.Descriptor instead.
func (VISIBLE) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

// Namespace tenant resource container
type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *resource.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 空间定义
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateNamespaceRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetMeta() *resource.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Namespace) GetSpec() *CreateNamespaceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域ID
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 父Namespace Id
	// @gotags: bson:"parent_id" json:"parent_id"
	ParentId string `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id" bson:"parent_id"`
	// 空间名称, 不允许修改
	// @gotags: bson:"name" json:"name" validate:"required"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" bson:"name" validate:"required"`
	// 空间负责人
	// @gotags: bson:"owner" json:"owner"
	Owner string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner" bson:"owner"`
	// 空间负责人助理
	// @gotags: bson:"assistants" json:"assistants"
	Assistants []string `protobuf:"bytes,9,rep,name=assistants,proto3" json:"assistants" bson:"assistants"`
	// 禁用项目, 该项目所有人暂时都无法访问
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 项目描述图片
	// @gotags: bson:"picture" json:"picture"
	Picture string `protobuf:"bytes,6,opt,name=picture,proto3" json:"picture" bson:"picture"`
	// 项目描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description" bson:"description"`
	// 空间可见性, 默认是私有空间
	// @gotags: bson:"visible" json:"visible"
	Visible VISIBLE `protobuf:"varint,8,opt,name=visible,proto3,enum=infraboard.mcenter.namespace.VISIBLE" json:"visible" bson:"visible"`
	// 空间的资源配额限制
	// @gotags: bson:"resource_quota" json:"resource_quota"
	ResourceQuota *ResourceQuotaSet `protobuf:"bytes,13,opt,name=resource_quota,json=resourceQuota,proto3" json:"resource_quota" bson:"resource_quota"`
	// 扩展信息
	// @gotags: bson:"extension" json:"extension"
	Extension map[string]string `protobuf:"bytes,14,rep,name=extension,proto3" json:"extension" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extension"`
	// 空间标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNamespaceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateNamespaceRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateNamespaceRequest) GetAssistants() []string {
	if x != nil {
		return x.Assistants
	}
	return nil
}

func (x *CreateNamespaceRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateNamespaceRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateNamespaceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateNamespaceRequest) GetVisible() VISIBLE {
	if x != nil {
		return x.Visible
	}
	return VISIBLE_PRIVATE
}

func (x *CreateNamespaceRequest) GetResourceQuota() *ResourceQuotaSet {
	if x != nil {
		return x.ResourceQuota
	}
	return nil
}

func (x *CreateNamespaceRequest) GetExtension() map[string]string {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *CreateNamespaceRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ResourceQuotaSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源列表
	// @gotags: bson:"items" json:"items"
	Items []*ResourceQuota `protobuf:"bytes,1,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ResourceQuotaSet) Reset() {
	*x = ResourceQuotaSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceQuotaSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceQuotaSet) ProtoMessage() {}

func (x *ResourceQuotaSet) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceQuotaSet.ProtoReflect.Descriptor instead.
func (*ResourceQuotaSet) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceQuotaSet) GetItems() []*ResourceQuota {
	if x != nil {
		return x.Items
	}
	return nil
}

// 资源配额限制策略
type ResourceQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源名称
	// @gotags: bson:"resource_name" json:"resource_name"
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name" bson:"resource_name"`
	// 软限制(水位线)
	// @gotags: bson:"soft" json:"soft"
	Soft int64 `protobuf:"varint,2,opt,name=soft,proto3" json:"soft" bson:"soft"`
	// 硬限制
	// @gotags: bson:"hard" json:"hard"
	Hard int64 `protobuf:"varint,3,opt,name=hard,proto3" json:"hard" bson:"hard"`
}

func (x *ResourceQuota) Reset() {
	*x = ResourceQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceQuota) ProtoMessage() {}

func (x *ResourceQuota) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceQuota.ProtoReflect.Descriptor instead.
func (*ResourceQuota) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceQuota) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ResourceQuota) GetSoft() int64 {
	if x != nil {
		return x.Soft
	}
	return 0
}

func (x *ResourceQuota) GetHard() int64 {
	if x != nil {
		return x.Hard
	}
	return 0
}

type NamespaceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表
	// @gotags: json:"items"
	Items []*Namespace `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *NamespaceSet) Reset() {
	*x = NamespaceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceSet) ProtoMessage() {}

func (x *NamespaceSet) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceSet.ProtoReflect.Descriptor instead.
func (*NamespaceSet) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *NamespaceSet) GetItems() []*Namespace {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_mcenter_apps_namespace_pb_namespace_proto protoreflect.FileDescriptor

var file_mcenter_apps_namespace_pb_namespace_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x48, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0xbb, 0x05, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x61, 0x0a, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x43, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x58, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x55, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x6f, 0x66, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x68, 0x61, 0x72, 0x64, 0x22, 0x63, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3d, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x22, 0x0a, 0x07,
	0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_namespace_pb_namespace_proto_rawDescOnce sync.Once
	file_mcenter_apps_namespace_pb_namespace_proto_rawDescData = file_mcenter_apps_namespace_pb_namespace_proto_rawDesc
)

func file_mcenter_apps_namespace_pb_namespace_proto_rawDescGZIP() []byte {
	file_mcenter_apps_namespace_pb_namespace_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_namespace_pb_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_namespace_pb_namespace_proto_rawDescData)
	})
	return file_mcenter_apps_namespace_pb_namespace_proto_rawDescData
}

var file_mcenter_apps_namespace_pb_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_namespace_pb_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mcenter_apps_namespace_pb_namespace_proto_goTypes = []interface{}{
	(VISIBLE)(0),                   // 0: infraboard.mcenter.namespace.VISIBLE
	(*Namespace)(nil),              // 1: infraboard.mcenter.namespace.Namespace
	(*CreateNamespaceRequest)(nil), // 2: infraboard.mcenter.namespace.CreateNamespaceRequest
	(*ResourceQuotaSet)(nil),       // 3: infraboard.mcenter.namespace.ResourceQuotaSet
	(*ResourceQuota)(nil),          // 4: infraboard.mcenter.namespace.ResourceQuota
	(*NamespaceSet)(nil),           // 5: infraboard.mcenter.namespace.NamespaceSet
	nil,                            // 6: infraboard.mcenter.namespace.CreateNamespaceRequest.ExtensionEntry
	nil,                            // 7: infraboard.mcenter.namespace.CreateNamespaceRequest.LabelsEntry
	(*resource.Meta)(nil),          // 8: infraboard.mcube.resource.Meta
}
var file_mcenter_apps_namespace_pb_namespace_proto_depIdxs = []int32{
	8, // 0: infraboard.mcenter.namespace.Namespace.meta:type_name -> infraboard.mcube.resource.Meta
	2, // 1: infraboard.mcenter.namespace.Namespace.spec:type_name -> infraboard.mcenter.namespace.CreateNamespaceRequest
	0, // 2: infraboard.mcenter.namespace.CreateNamespaceRequest.visible:type_name -> infraboard.mcenter.namespace.VISIBLE
	3, // 3: infraboard.mcenter.namespace.CreateNamespaceRequest.resource_quota:type_name -> infraboard.mcenter.namespace.ResourceQuotaSet
	6, // 4: infraboard.mcenter.namespace.CreateNamespaceRequest.extension:type_name -> infraboard.mcenter.namespace.CreateNamespaceRequest.ExtensionEntry
	7, // 5: infraboard.mcenter.namespace.CreateNamespaceRequest.labels:type_name -> infraboard.mcenter.namespace.CreateNamespaceRequest.LabelsEntry
	4, // 6: infraboard.mcenter.namespace.ResourceQuotaSet.items:type_name -> infraboard.mcenter.namespace.ResourceQuota
	1, // 7: infraboard.mcenter.namespace.NamespaceSet.items:type_name -> infraboard.mcenter.namespace.Namespace
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mcenter_apps_namespace_pb_namespace_proto_init() }
func file_mcenter_apps_namespace_pb_namespace_proto_init() {
	if File_mcenter_apps_namespace_pb_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceQuotaSet); i {
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
		file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceQuota); i {
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
		file_mcenter_apps_namespace_pb_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceSet); i {
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
			RawDescriptor: file_mcenter_apps_namespace_pb_namespace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_namespace_pb_namespace_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_namespace_pb_namespace_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_namespace_pb_namespace_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_namespace_pb_namespace_proto_msgTypes,
	}.Build()
	File_mcenter_apps_namespace_pb_namespace_proto = out.File
	file_mcenter_apps_namespace_pb_namespace_proto_rawDesc = nil
	file_mcenter_apps_namespace_pb_namespace_proto_goTypes = nil
	file_mcenter_apps_namespace_pb_namespace_proto_depIdxs = nil
}
