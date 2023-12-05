// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/role/pb/permission.proto

package role

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

// EffectType 授权效力包括两种：允许（Allow）和拒绝（Deny）
type EffectType int32

const (
	// 允许访问
	EffectType_ALLOW EffectType = 0
	// 拒绝访问
	EffectType_DENY EffectType = 1
)

// Enum value maps for EffectType.
var (
	EffectType_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	EffectType_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x EffectType) Enum() *EffectType {
	p := new(EffectType)
	*p = x
	return p
}

func (x EffectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EffectType) Descriptor() protoreflect.EnumDescriptor {
	return file_mcenter_apps_role_pb_permission_proto_enumTypes[0].Descriptor()
}

func (EffectType) Type() protoreflect.EnumType {
	return &file_mcenter_apps_role_pb_permission_proto_enumTypes[0]
}

func (x EffectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EffectType.Descriptor instead.
func (EffectType) EnumDescriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{0}
}

// Permission 权限
type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 角色Id
	// @gotags: bson:"role_id" json:"role_id"
	RoleId string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id" bson:"role_id"`
	// 权限具体定义
	// @gotags: bson:",inline" json:"spec"
	Spec *PermissionSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 范围限制, 由策略引擎动态补充
	// @gotags: bson:"-" json:"scope"
	Scope []*resource.LabelRequirement `protobuf:"bytes,6,rep,name=scope,proto3" json:"scope" bson:"-"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Permission) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *Permission) GetSpec() *PermissionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Permission) GetScope() []*resource.LabelRequirement {
	if x != nil {
		return x.Scope
	}
	return nil
}

type PermissionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建人
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 权限描述
	// @gotags: bson:"desc" json:"desc"
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc" bson:"desc"`
	// 效力
	// @gotags: bson:"effect" json:"effect"
	Effect EffectType `protobuf:"varint,4,opt,name=effect,proto3,enum=infraboard.mcenter.role.EffectType" json:"effect" bson:"effect"`
	// 服务ID
	// @gotags: bson:"service_id" json:"service_id"
	ServiceId string `protobuf:"bytes,5,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// 资源列表
	// @gotags: bson:"resource_name" json:"resource_name"
	ResourceName string `protobuf:"bytes,6,opt,name=resource_name,json=resourceName,proto3" json:"resource_name" bson:"resource_name"`
	// 维度
	// @gotags: bson:"label_key" json:"label_key"
	LabelKey string `protobuf:"bytes,7,opt,name=label_key,json=labelKey,proto3" json:"label_key" bson:"label_key"`
	// 适配所有值
	// @gotags: bson:"match_all" json:"match_all"
	MatchAll bool `protobuf:"varint,8,opt,name=match_all,json=matchAll,proto3" json:"match_all" bson:"match_all"`
	// 标识值
	// @gotags: bson:"label_values" json:"label_values"
	LabelValues []string `protobuf:"bytes,9,rep,name=label_values,json=labelValues,proto3" json:"label_values" bson:"label_values"`
}

func (x *PermissionSpec) Reset() {
	*x = PermissionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionSpec) ProtoMessage() {}

func (x *PermissionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionSpec.ProtoReflect.Descriptor instead.
func (*PermissionSpec) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionSpec) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *PermissionSpec) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PermissionSpec) GetEffect() EffectType {
	if x != nil {
		return x.Effect
	}
	return EffectType_ALLOW
}

func (x *PermissionSpec) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *PermissionSpec) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *PermissionSpec) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *PermissionSpec) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *PermissionSpec) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

// PermissionSet 用户列表
type PermissionSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Permission `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PermissionSet) Reset() {
	*x = PermissionSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionSet) ProtoMessage() {}

func (x *PermissionSet) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionSet.ProtoReflect.Descriptor instead.
func (*PermissionSet) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PermissionSet) GetItems() []*Permission {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddPermissionToRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建者ID
	// @gotags: json:"create_by" validate:"required"
	CreateBy string `protobuf:"bytes,2,opt,name=create_by,json=createBy,proto3" json:"create_by" validate:"required"`
	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// 角色Id
	// @gotags: json:"role_id" validate:"required,lte=64"
	RoleId string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id" validate:"required,lte=64"`
	// 添加的权限列表
	// @gotags: json:"permissions" validate:"required"
	Permissions []*PermissionSpec `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions" validate:"required"`
}

func (x *AddPermissionToRoleRequest) Reset() {
	*x = AddPermissionToRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionToRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionToRoleRequest) ProtoMessage() {}

func (x *AddPermissionToRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionToRoleRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionToRoleRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{3}
}

func (x *AddPermissionToRoleRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *AddPermissionToRoleRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *AddPermissionToRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *AddPermissionToRoleRequest) GetPermissions() []*PermissionSpec {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RemovePermissionFromRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// @gotags: json:"role_id" validate:"required,lte=64"
	// 角色Id
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id" validate:"required,lte=64"`
	// @gotags: json:"remove_all"
	// 移除所有权限
	RemoveAll bool `protobuf:"varint,3,opt,name=remove_all,json=removeAll,proto3" json:"remove_all"`
	// 需要移除的权限
	// @gotags: json:"permission_id"
	PermissionId []string `protobuf:"bytes,4,rep,name=permission_id,json=permissionId,proto3" json:"permission_id"`
}

func (x *RemovePermissionFromRoleRequest) Reset() {
	*x = RemovePermissionFromRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionFromRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionFromRoleRequest) ProtoMessage() {}

func (x *RemovePermissionFromRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionFromRoleRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionFromRoleRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{4}
}

func (x *RemovePermissionFromRoleRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *RemovePermissionFromRoleRequest) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RemovePermissionFromRoleRequest) GetRemoveAll() bool {
	if x != nil {
		return x.RemoveAll
	}
	return false
}

func (x *RemovePermissionFromRoleRequest) GetPermissionId() []string {
	if x != nil {
		return x.PermissionId
	}
	return nil
}

type UpdatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"scope"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope"`
	// permission id
	// @gotags: json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" validate:"required,lte=64"`
	// 维度
	// @gotags: json:"label_key"
	LabelKey string `protobuf:"bytes,3,opt,name=label_key,json=labelKey,proto3" json:"label_key"`
	// 适配所有值
	// @gotags: json:"match_all"
	MatchAll bool `protobuf:"varint,4,opt,name=match_all,json=matchAll,proto3" json:"match_all"`
	// 标识值
	// @gotags: json:"label_values"
	LabelValues []string `protobuf:"bytes,5,rep,name=label_values,json=labelValues,proto3" json:"label_values"`
}

func (x *UpdatePermissionRequest) Reset() {
	*x = UpdatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePermissionRequest) ProtoMessage() {}

func (x *UpdatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_role_pb_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePermissionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_role_pb_permission_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePermissionRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *UpdatePermissionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePermissionRequest) GetLabelKey() string {
	if x != nil {
		return x.LabelKey
	}
	return ""
}

func (x *UpdatePermissionRequest) GetMatchAll() bool {
	if x != nil {
		return x.MatchAll
	}
	return false
}

func (x *UpdatePermissionRequest) GetLabelValues() []string {
	if x != nil {
		return x.LabelValues
	}
	return nil
}

var File_mcenter_apps_role_pb_permission_proto protoreflect.FileDescriptor

var file_mcenter_apps_role_pb_permission_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x1a, 0x1d, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x41, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x6c,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb6,
	0x01, 0x0a, 0x1f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41,
	0x6c, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x21, 0x0a, 0x0a, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_role_pb_permission_proto_rawDescOnce sync.Once
	file_mcenter_apps_role_pb_permission_proto_rawDescData = file_mcenter_apps_role_pb_permission_proto_rawDesc
)

func file_mcenter_apps_role_pb_permission_proto_rawDescGZIP() []byte {
	file_mcenter_apps_role_pb_permission_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_role_pb_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_role_pb_permission_proto_rawDescData)
	})
	return file_mcenter_apps_role_pb_permission_proto_rawDescData
}

var file_mcenter_apps_role_pb_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcenter_apps_role_pb_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mcenter_apps_role_pb_permission_proto_goTypes = []interface{}{
	(EffectType)(0),                         // 0: infraboard.mcenter.role.EffectType
	(*Permission)(nil),                      // 1: infraboard.mcenter.role.Permission
	(*PermissionSpec)(nil),                  // 2: infraboard.mcenter.role.PermissionSpec
	(*PermissionSet)(nil),                   // 3: infraboard.mcenter.role.PermissionSet
	(*AddPermissionToRoleRequest)(nil),      // 4: infraboard.mcenter.role.AddPermissionToRoleRequest
	(*RemovePermissionFromRoleRequest)(nil), // 5: infraboard.mcenter.role.RemovePermissionFromRoleRequest
	(*UpdatePermissionRequest)(nil),         // 6: infraboard.mcenter.role.UpdatePermissionRequest
	(*resource.LabelRequirement)(nil),       // 7: infraboard.mcube.resource.LabelRequirement
	(*resource.Scope)(nil),                  // 8: infraboard.mcube.resource.Scope
}
var file_mcenter_apps_role_pb_permission_proto_depIdxs = []int32{
	2, // 0: infraboard.mcenter.role.Permission.spec:type_name -> infraboard.mcenter.role.PermissionSpec
	7, // 1: infraboard.mcenter.role.Permission.scope:type_name -> infraboard.mcube.resource.LabelRequirement
	0, // 2: infraboard.mcenter.role.PermissionSpec.effect:type_name -> infraboard.mcenter.role.EffectType
	1, // 3: infraboard.mcenter.role.PermissionSet.items:type_name -> infraboard.mcenter.role.Permission
	8, // 4: infraboard.mcenter.role.AddPermissionToRoleRequest.scope:type_name -> infraboard.mcube.resource.Scope
	2, // 5: infraboard.mcenter.role.AddPermissionToRoleRequest.permissions:type_name -> infraboard.mcenter.role.PermissionSpec
	8, // 6: infraboard.mcenter.role.RemovePermissionFromRoleRequest.scope:type_name -> infraboard.mcube.resource.Scope
	8, // 7: infraboard.mcenter.role.UpdatePermissionRequest.scope:type_name -> infraboard.mcube.resource.Scope
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_mcenter_apps_role_pb_permission_proto_init() }
func file_mcenter_apps_role_pb_permission_proto_init() {
	if File_mcenter_apps_role_pb_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_role_pb_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_mcenter_apps_role_pb_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionSpec); i {
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
		file_mcenter_apps_role_pb_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionSet); i {
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
		file_mcenter_apps_role_pb_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermissionToRoleRequest); i {
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
		file_mcenter_apps_role_pb_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionFromRoleRequest); i {
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
		file_mcenter_apps_role_pb_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePermissionRequest); i {
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
			RawDescriptor: file_mcenter_apps_role_pb_permission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_role_pb_permission_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_role_pb_permission_proto_depIdxs,
		EnumInfos:         file_mcenter_apps_role_pb_permission_proto_enumTypes,
		MessageInfos:      file_mcenter_apps_role_pb_permission_proto_msgTypes,
	}.Build()
	File_mcenter_apps_role_pb_permission_proto = out.File
	file_mcenter_apps_role_pb_permission_proto_rawDesc = nil
	file_mcenter_apps_role_pb_permission_proto_goTypes = nil
	file_mcenter_apps_role_pb_permission_proto_depIdxs = nil
}
