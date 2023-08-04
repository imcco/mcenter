// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/notify/pb/sms.proto

package notify

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

type SmsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信服务商
	// @gotags: bson:"provider" json:"provider" env:"SMS_PROVIDER"
	Provider PROVIDER `protobuf:"varint,1,opt,name=provider,proto3,enum=infraboard.mcenter.notify.PROVIDER" json:"provider" bson:"provider" env:"SMS_PROVIDER"`
	// 腾讯短信服务配置
	// @gotags: bson:"tencent" json:"tencent"
	Tencent *TencentSmsConfig `protobuf:"bytes,2,opt,name=tencent,proto3" json:"tencent" bson:"tencent"`
}

func (x *SmsConfig) Reset() {
	*x = SmsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_notify_pb_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsConfig) ProtoMessage() {}

func (x *SmsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsConfig.ProtoReflect.Descriptor instead.
func (*SmsConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SmsConfig) GetProvider() PROVIDER {
	if x != nil {
		return x.Provider
	}
	return PROVIDER_TENCENT
}

func (x *SmsConfig) GetTencent() *TencentSmsConfig {
	if x != nil {
		return x.Tencent
	}
	return nil
}

// 接口和相关文档请参考https://console.cloud.tencent.com/api/explorer?Product=sms&Version=2019-07-11&Action=SendSms&SignVersion=
type TencentSmsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务地址
	// @gotags: bson:"endpoint" json:"endpoint" env:"SMS_TENCENT_ENDPOINT"
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint" bson:"endpoint" env:"SMS_TENCENT_ENDPOINT"`
	// id
	// @gotags: bson:"secret_id" json:"secret_id" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_ID"
	SecretId string `protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3" json:"secret_id" bson:"secret_id" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_ID"`
	// secret
	// @gotags: bson:"secret_key" json:"secret_key" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_KEY"
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key" bson:"secret_key" validate:"required,lte=64" env:"SMS_TENCENT_SECRET_KEY"`
	// app id
	// @gotags: bson:"app_id" json:"app_id" validate:"required,lte=64" env:"SMS_TENCENT_APPID"
	AppId string `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id" bson:"app_id" validate:"required,lte=64" env:"SMS_TENCENT_APPID"`
	// sign
	// @gotags: bson:"sign" json:"sign" validate:"required,lte=128" env:"SMS_TENCENT_SIGN"
	Sign string `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign" bson:"sign" validate:"required,lte=128" env:"SMS_TENCENT_SIGN"`
}

func (x *TencentSmsConfig) Reset() {
	*x = TencentSmsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_notify_pb_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TencentSmsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TencentSmsConfig) ProtoMessage() {}

func (x *TencentSmsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TencentSmsConfig.ProtoReflect.Descriptor instead.
func (*TencentSmsConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_sms_proto_rawDescGZIP(), []int{1}
}

func (x *TencentSmsConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TencentSmsConfig) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *TencentSmsConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *TencentSmsConfig) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *TencentSmsConfig) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

var File_mcenter_apps_notify_pb_sms_proto protoreflect.FileDescriptor

var file_mcenter_apps_notify_pb_sms_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x23, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x54,
	0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x67, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_notify_pb_sms_proto_rawDescOnce sync.Once
	file_mcenter_apps_notify_pb_sms_proto_rawDescData = file_mcenter_apps_notify_pb_sms_proto_rawDesc
)

func file_mcenter_apps_notify_pb_sms_proto_rawDescGZIP() []byte {
	file_mcenter_apps_notify_pb_sms_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_notify_pb_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_notify_pb_sms_proto_rawDescData)
	})
	return file_mcenter_apps_notify_pb_sms_proto_rawDescData
}

var file_mcenter_apps_notify_pb_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mcenter_apps_notify_pb_sms_proto_goTypes = []interface{}{
	(*SmsConfig)(nil),        // 0: infraboard.mcenter.notify.SmsConfig
	(*TencentSmsConfig)(nil), // 1: infraboard.mcenter.notify.TencentSmsConfig
	(PROVIDER)(0),            // 2: infraboard.mcenter.notify.PROVIDER
}
var file_mcenter_apps_notify_pb_sms_proto_depIdxs = []int32{
	2, // 0: infraboard.mcenter.notify.SmsConfig.provider:type_name -> infraboard.mcenter.notify.PROVIDER
	1, // 1: infraboard.mcenter.notify.SmsConfig.tencent:type_name -> infraboard.mcenter.notify.TencentSmsConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_notify_pb_sms_proto_init() }
func file_mcenter_apps_notify_pb_sms_proto_init() {
	if File_mcenter_apps_notify_pb_sms_proto != nil {
		return
	}
	file_mcenter_apps_notify_pb_notify_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_notify_pb_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsConfig); i {
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
		file_mcenter_apps_notify_pb_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TencentSmsConfig); i {
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
			RawDescriptor: file_mcenter_apps_notify_pb_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_notify_pb_sms_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_notify_pb_sms_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_notify_pb_sms_proto_msgTypes,
	}.Build()
	File_mcenter_apps_notify_pb_sms_proto = out.File
	file_mcenter_apps_notify_pb_sms_proto_rawDesc = nil
	file_mcenter_apps_notify_pb_sms_proto_goTypes = nil
	file_mcenter_apps_notify_pb_sms_proto_depIdxs = nil
}
