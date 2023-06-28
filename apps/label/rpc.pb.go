// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/label/pb/rpc.proto

package label

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type QueryLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
}

func (x *QueryLabelRequest) Reset() {
	*x = QueryLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLabelRequest) ProtoMessage() {}

func (x *QueryLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLabelRequest.ProtoReflect.Descriptor instead.
func (*QueryLabelRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryLabelRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpdateLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 标签Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 标签信息
	// @gotags: json:"spec"
	Spec *CreateLabelRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec"`
}

func (x *UpdateLabelRequest) Reset() {
	*x = UpdateLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLabelRequest) ProtoMessage() {}

func (x *UpdateLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLabelRequest.ProtoReflect.Descriptor instead.
func (*UpdateLabelRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateLabelRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateLabelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateLabelRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateLabelRequest) GetSpec() *CreateLabelRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 构建配置Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteLabelRequest) Reset() {
	*x = DeleteLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLabelRequest) ProtoMessage() {}

func (x *DeleteLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_label_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLabelRequest.ProtoReflect.Descriptor instead.
func (*DeleteLabelRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_label_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteLabelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mcenter_apps_label_pb_rpc_proto protoreflect.FileDescriptor

var file_mcenter_apps_label_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x18, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x12, 0x40, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x64, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x5d, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2b, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mcenter_apps_label_pb_rpc_proto_rawDescOnce sync.Once
	file_mcenter_apps_label_pb_rpc_proto_rawDescData = file_mcenter_apps_label_pb_rpc_proto_rawDesc
)

func file_mcenter_apps_label_pb_rpc_proto_rawDescGZIP() []byte {
	file_mcenter_apps_label_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_label_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_label_pb_rpc_proto_rawDescData)
	})
	return file_mcenter_apps_label_pb_rpc_proto_rawDescData
}

var file_mcenter_apps_label_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mcenter_apps_label_pb_rpc_proto_goTypes = []interface{}{
	(*QueryLabelRequest)(nil),   // 0: infraboard.mcenter.label.QueryLabelRequest
	(*UpdateLabelRequest)(nil),  // 1: infraboard.mcenter.label.UpdateLabelRequest
	(*DeleteLabelRequest)(nil),  // 2: infraboard.mcenter.label.DeleteLabelRequest
	(*request.PageRequest)(nil), // 3: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),    // 4: infraboard.mcube.request.UpdateMode
	(*CreateLabelRequest)(nil),  // 5: infraboard.mcenter.label.CreateLabelRequest
	(*LabelSet)(nil),            // 6: infraboard.mcenter.label.LabelSet
}
var file_mcenter_apps_label_pb_rpc_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.label.QueryLabelRequest.page:type_name -> infraboard.mcube.page.PageRequest
	4, // 1: infraboard.mcenter.label.UpdateLabelRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	5, // 2: infraboard.mcenter.label.UpdateLabelRequest.spec:type_name -> infraboard.mcenter.label.CreateLabelRequest
	0, // 3: infraboard.mcenter.label.RPC.QueryLabel:input_type -> infraboard.mcenter.label.QueryLabelRequest
	6, // 4: infraboard.mcenter.label.RPC.QueryLabel:output_type -> infraboard.mcenter.label.LabelSet
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mcenter_apps_label_pb_rpc_proto_init() }
func file_mcenter_apps_label_pb_rpc_proto_init() {
	if File_mcenter_apps_label_pb_rpc_proto != nil {
		return
	}
	file_mcenter_apps_label_pb_label_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_label_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLabelRequest); i {
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
		file_mcenter_apps_label_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLabelRequest); i {
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
		file_mcenter_apps_label_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLabelRequest); i {
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
			RawDescriptor: file_mcenter_apps_label_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcenter_apps_label_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_label_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_label_pb_rpc_proto_msgTypes,
	}.Build()
	File_mcenter_apps_label_pb_rpc_proto = out.File
	file_mcenter_apps_label_pb_rpc_proto_rawDesc = nil
	file_mcenter_apps_label_pb_rpc_proto_goTypes = nil
	file_mcenter_apps_label_pb_rpc_proto_depIdxs = nil
}
