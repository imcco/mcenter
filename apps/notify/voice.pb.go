// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/notify/pb/voice.proto

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

type VoiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 短信服务商
	// @gotags: bson:"provider" json:"provider" env:"SMS_PROVIDER"
	Provider PROVIDER `protobuf:"varint,1,opt,name=provider,proto3,enum=infraboard.mcenter.notify.PROVIDER" json:"provider" bson:"provider" env:"SMS_PROVIDER"`
	// 腾讯云配置
	// @gotags: bson:"tencent" json:"tencent"
	Tencent *TencentVoiceConfig `protobuf:"bytes,2,opt,name=tencent,proto3" json:"tencent" bson:"tencent"`
}

func (x *VoiceConfig) Reset() {
	*x = VoiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_notify_pb_voice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceConfig) ProtoMessage() {}

func (x *VoiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_voice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceConfig.ProtoReflect.Descriptor instead.
func (*VoiceConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_voice_proto_rawDescGZIP(), []int{0}
}

func (x *VoiceConfig) GetProvider() PROVIDER {
	if x != nil {
		return x.Provider
	}
	return PROVIDER_TENCENT
}

func (x *VoiceConfig) GetTencent() *TencentVoiceConfig {
	if x != nil {
		return x.Tencent
	}
	return nil
}

type TencentVoiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SDK 会自动指定域名，通常无需指定域名，但访问金融区的服务时必须手动指定域名
	// 例如 VMS 的上海金融区域名为 vms.ap-shanghai-fsi.tencentcloudapi.com
	// @gotags: bson:"endpoint" json:"endpoint" env:"VOICE_TENCENT_ENDPOINT"
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint" bson:"endpoint" env:"VOICE_TENCENT_ENDPOINT"`
	// 腾讯云凭证Id
	// @gotags: bson:"secret_id" json:"secret_id" validate:"required" env:"VOICE_TENCENT_SECRET_ID"
	SecretId string `protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3" json:"secret_id" bson:"secret_id" validate:"required" env:"VOICE_TENCENT_SECRET_ID"`
	// 腾讯云凭证Key
	// @gotags: bson:"secret_key" json:"secret_key" validate:"required" env:"VOICE_TENCENT_SECRET_KEY"
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key" bson:"secret_key" validate:"required" env:"VOICE_TENCENT_SECRET_KEY"`
	// 腾讯云控制台: https://console.cloud.tencent.com/vms/app 查看appId
	// @gotags: bson:"app_id" json:"app_id" validate:"required" env:"VOICE_TENCENT_APPID"
	AppId string `protobuf:"bytes,4,opt,name=app_id,json=appId,proto3" json:"app_id" bson:"app_id" validate:"required" env:"VOICE_TENCENT_APPID"`
	// 实例化 VMS 的 client 对象
	// 第二个参数是地域信息，可以直接填写字符串 ap-guangzhou，或者引用预设的常量
	// @gotags: bson:"region" json:"region" env:"VOICE_TENCENT_REGION"
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region" bson:"region" env:"VOICE_TENCENT_REGION"`
	// SDK 默认用 TC3-HMAC-SHA256 进行签名，非必要请不要修改该字段
	// @gotags: bson:"sign_method" json:"sign_method" env:"VOICE_TENCENT_SIGN_METHOD"
	SignMethod string `protobuf:"bytes,6,opt,name=sign_method,json=signMethod,proto3" json:"sign_method" bson:"sign_method" env:"VOICE_TENCENT_SIGN_METHOD"`
	// SDK 默认使用 POST 方法
	// 如需使用 GET 方法，可以在此处设置，但 GET 方法无法处理较大的请求
	// @gotags: bson:"req_method" json:"req_method" env:"VOICE_TENCENT_REQ_METHOD"
	ReqMethod string `protobuf:"bytes,7,opt,name=req_method,json=reqMethod,proto3" json:"req_method" bson:"req_method" env:"VOICE_TENCENT_REQ_METHOD"`
}

func (x *TencentVoiceConfig) Reset() {
	*x = TencentVoiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_notify_pb_voice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TencentVoiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TencentVoiceConfig) ProtoMessage() {}

func (x *TencentVoiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_voice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TencentVoiceConfig.ProtoReflect.Descriptor instead.
func (*TencentVoiceConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_voice_proto_rawDescGZIP(), []int{1}
}

func (x *TencentVoiceConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TencentVoiceConfig) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *TencentVoiceConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *TencentVoiceConfig) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *TencentVoiceConfig) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *TencentVoiceConfig) GetSignMethod() string {
	if x != nil {
		return x.SignMethod
	}
	return ""
}

func (x *TencentVoiceConfig) GetReqMethod() string {
	if x != nil {
		return x.ReqMethod
	}
	return ""
}

var File_mcenter_apps_notify_pb_voice_proto protoreflect.FileDescriptor

var file_mcenter_apps_notify_pb_voice_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a,
	0x23, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x22, 0xdb,
	0x01, 0x0a, 0x12, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mcenter_apps_notify_pb_voice_proto_rawDescOnce sync.Once
	file_mcenter_apps_notify_pb_voice_proto_rawDescData = file_mcenter_apps_notify_pb_voice_proto_rawDesc
)

func file_mcenter_apps_notify_pb_voice_proto_rawDescGZIP() []byte {
	file_mcenter_apps_notify_pb_voice_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_notify_pb_voice_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_notify_pb_voice_proto_rawDescData)
	})
	return file_mcenter_apps_notify_pb_voice_proto_rawDescData
}

var file_mcenter_apps_notify_pb_voice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mcenter_apps_notify_pb_voice_proto_goTypes = []interface{}{
	(*VoiceConfig)(nil),        // 0: infraboard.mcenter.notify.VoiceConfig
	(*TencentVoiceConfig)(nil), // 1: infraboard.mcenter.notify.TencentVoiceConfig
	(PROVIDER)(0),              // 2: infraboard.mcenter.notify.PROVIDER
}
var file_mcenter_apps_notify_pb_voice_proto_depIdxs = []int32{
	2, // 0: infraboard.mcenter.notify.VoiceConfig.provider:type_name -> infraboard.mcenter.notify.PROVIDER
	1, // 1: infraboard.mcenter.notify.VoiceConfig.tencent:type_name -> infraboard.mcenter.notify.TencentVoiceConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mcenter_apps_notify_pb_voice_proto_init() }
func file_mcenter_apps_notify_pb_voice_proto_init() {
	if File_mcenter_apps_notify_pb_voice_proto != nil {
		return
	}
	file_mcenter_apps_notify_pb_notify_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_notify_pb_voice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceConfig); i {
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
		file_mcenter_apps_notify_pb_voice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TencentVoiceConfig); i {
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
			RawDescriptor: file_mcenter_apps_notify_pb_voice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_notify_pb_voice_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_notify_pb_voice_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_notify_pb_voice_proto_msgTypes,
	}.Build()
	File_mcenter_apps_notify_pb_voice_proto = out.File
	file_mcenter_apps_notify_pb_voice_proto_rawDesc = nil
	file_mcenter_apps_notify_pb_voice_proto_goTypes = nil
	file_mcenter_apps_notify_pb_voice_proto_depIdxs = nil
}
