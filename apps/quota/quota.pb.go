// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/quota/pb/quota.proto

package quota

import (
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

// 资源配额限制策略
type Quota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 配额定义
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateQuotaRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Quota) Reset() {
	*x = Quota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_quota_pb_quota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_apps_quota_pb_quota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_apps_quota_pb_quota_proto_rawDescGZIP(), []int{0}
}

func (x *Quota) GetSpec() *CreateQuotaRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

// 资源配额限制策略
type CreateQuotaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 所属空间
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 所属组
	// @gotags: bson:"group" json:"group"
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group" bson:"group"`
}

func (x *CreateQuotaRequest) Reset() {
	*x = CreateQuotaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_quota_pb_quota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuotaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuotaRequest) ProtoMessage() {}

func (x *CreateQuotaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_quota_pb_quota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuotaRequest.ProtoReflect.Descriptor instead.
func (*CreateQuotaRequest) Descriptor() ([]byte, []int) {
	return file_apps_quota_pb_quota_proto_rawDescGZIP(), []int{1}
}

func (x *CreateQuotaRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateQuotaRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateQuotaRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

var File_apps_quota_pb_quota_proto protoreflect.FileDescriptor

var file_apps_quota_pb_quota_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x70, 0x62, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x40,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x22, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_quota_pb_quota_proto_rawDescOnce sync.Once
	file_apps_quota_pb_quota_proto_rawDescData = file_apps_quota_pb_quota_proto_rawDesc
)

func file_apps_quota_pb_quota_proto_rawDescGZIP() []byte {
	file_apps_quota_pb_quota_proto_rawDescOnce.Do(func() {
		file_apps_quota_pb_quota_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_quota_pb_quota_proto_rawDescData)
	})
	return file_apps_quota_pb_quota_proto_rawDescData
}

var file_apps_quota_pb_quota_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_quota_pb_quota_proto_goTypes = []interface{}{
	(*Quota)(nil),              // 0: infraboard.mcenter.quota.Quota
	(*CreateQuotaRequest)(nil), // 1: infraboard.mcenter.quota.CreateQuotaRequest
}
var file_apps_quota_pb_quota_proto_depIdxs = []int32{
	1, // 0: infraboard.mcenter.quota.Quota.spec:type_name -> infraboard.mcenter.quota.CreateQuotaRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_quota_pb_quota_proto_init() }
func file_apps_quota_pb_quota_proto_init() {
	if File_apps_quota_pb_quota_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_quota_pb_quota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quota); i {
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
		file_apps_quota_pb_quota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuotaRequest); i {
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
			RawDescriptor: file_apps_quota_pb_quota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_quota_pb_quota_proto_goTypes,
		DependencyIndexes: file_apps_quota_pb_quota_proto_depIdxs,
		MessageInfos:      file_apps_quota_pb_quota_proto_msgTypes,
	}.Build()
	File_apps_quota_pb_quota_proto = out.File
	file_apps_quota_pb_quota_proto_rawDesc = nil
	file_apps_quota_pb_quota_proto_goTypes = nil
	file_apps_quota_pb_quota_proto_depIdxs = nil
}
