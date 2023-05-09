// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/workload/pb/workload.proto

package workload

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

type WorkLoad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户
	// @gotags: json:"user" validate:"required"
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user" validate:"required"`
	// 空间
	// @gotags: json:"namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" validate:"required"`
	// 年
	// @gotags: json:"year" validate:"required"
	Year int64 `protobuf:"varint,3,opt,name=year,proto3" json:"year" validate:"required"`
	// 月
	// @gotags: json:"month" validate:"required"
	Month int32 `protobuf:"varint,4,opt,name=month,proto3" json:"month" validate:"required"`
	// 日
	// @gotags: json:"day" validate:"required"
	Day int32 `protobuf:"varint,5,opt,name=day,proto3" json:"day" validate:"required"`
	// 开始时间
	// @gotags: json:"start_houre" validate:"required"
	StartHoure int32 `protobuf:"varint,6,opt,name=start_houre,json=startHoure,proto3" json:"start_houre" validate:"required"`
	// 结束时间
	// @gotags: json:"end_houre" validate:"required"
	EndHoure int32 `protobuf:"varint,7,opt,name=end_houre,json=endHoure,proto3" json:"end_houre" validate:"required"`
}

func (x *WorkLoad) Reset() {
	*x = WorkLoad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_workload_pb_workload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkLoad) ProtoMessage() {}

func (x *WorkLoad) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_workload_pb_workload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkLoad.ProtoReflect.Descriptor instead.
func (*WorkLoad) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_workload_pb_workload_proto_rawDescGZIP(), []int{0}
}

func (x *WorkLoad) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *WorkLoad) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkLoad) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *WorkLoad) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *WorkLoad) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *WorkLoad) GetStartHoure() int32 {
	if x != nil {
		return x.StartHoure
	}
	return 0
}

func (x *WorkLoad) GetEndHoure() int32 {
	if x != nil {
		return x.EndHoure
	}
	return 0
}

var File_mcenter_apps_workload_pb_workload_proto protoreflect.FileDescriptor

var file_mcenter_apps_workload_pb_workload_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x4c,
	0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x6f, 0x75,
	0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x65, 0x42,
	0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_workload_pb_workload_proto_rawDescOnce sync.Once
	file_mcenter_apps_workload_pb_workload_proto_rawDescData = file_mcenter_apps_workload_pb_workload_proto_rawDesc
)

func file_mcenter_apps_workload_pb_workload_proto_rawDescGZIP() []byte {
	file_mcenter_apps_workload_pb_workload_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_workload_pb_workload_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_workload_pb_workload_proto_rawDescData)
	})
	return file_mcenter_apps_workload_pb_workload_proto_rawDescData
}

var file_mcenter_apps_workload_pb_workload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcenter_apps_workload_pb_workload_proto_goTypes = []interface{}{
	(*WorkLoad)(nil), // 0: infraboard.mcenter.workload.WorkLoad
}
var file_mcenter_apps_workload_pb_workload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_workload_pb_workload_proto_init() }
func file_mcenter_apps_workload_pb_workload_proto_init() {
	if File_mcenter_apps_workload_pb_workload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_workload_pb_workload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkLoad); i {
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
			RawDescriptor: file_mcenter_apps_workload_pb_workload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_workload_pb_workload_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_workload_pb_workload_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_workload_pb_workload_proto_msgTypes,
	}.Build()
	File_mcenter_apps_workload_pb_workload_proto = out.File
	file_mcenter_apps_workload_pb_workload_proto_rawDesc = nil
	file_mcenter_apps_workload_pb_workload_proto_goTypes = nil
	file_mcenter_apps_workload_pb_workload_proto_depIdxs = nil
}
