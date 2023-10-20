// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/internal/tlsconfig/proto/config.proto

package proto

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

type TLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CA certificate file to verify certificates provided by the other party.
	CaCertFile *string `protobuf:"bytes,1,opt,name=ca_cert_file,json=caCertFile" json:"ca_cert_file,omitempty"`
	// Local certificate file.
	TlsCertFile *string `protobuf:"bytes,2,opt,name=tls_cert_file,json=tlsCertFile" json:"tls_cert_file,omitempty"`
	// Private key file corresponding to the certificate above.
	TlsKeyFile *string `protobuf:"bytes,3,opt,name=tls_key_file,json=tlsKeyFile" json:"tls_key_file,omitempty"`
	// Whether to ignore the cert validation.
	DisableCertValidation *bool `protobuf:"varint,4,opt,name=disable_cert_validation,json=disableCertValidation" json:"disable_cert_validation,omitempty"`
	// ServerName override
	ServerName *string `protobuf:"bytes,5,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
	// Certificate reload interval in seconds. If configured, the TLS cert will
	// be reloaded every reload_interval_sec seconds. This is useful when
	// certificates are generated and refreshed dynamically.
	ReloadIntervalSec *int32 `protobuf:"varint,6,opt,name=reload_interval_sec,json=reloadIntervalSec" json:"reload_interval_sec,omitempty"`
}

func (x *TLSConfig) Reset() {
	*x = TLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSConfig) ProtoMessage() {}

func (x *TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSConfig.ProtoReflect.Descriptor instead.
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *TLSConfig) GetCaCertFile() string {
	if x != nil && x.CaCertFile != nil {
		return *x.CaCertFile
	}
	return ""
}

func (x *TLSConfig) GetTlsCertFile() string {
	if x != nil && x.TlsCertFile != nil {
		return *x.TlsCertFile
	}
	return ""
}

func (x *TLSConfig) GetTlsKeyFile() string {
	if x != nil && x.TlsKeyFile != nil {
		return *x.TlsKeyFile
	}
	return ""
}

func (x *TLSConfig) GetDisableCertValidation() bool {
	if x != nil && x.DisableCertValidation != nil {
		return *x.DisableCertValidation
	}
	return false
}

func (x *TLSConfig) GetServerName() string {
	if x != nil && x.ServerName != nil {
		return *x.ServerName
	}
	return ""
}

func (x *TLSConfig) GetReloadIntervalSec() int32 {
	if x != nil && x.ReloadIntervalSec != nil {
		return *x.ReloadIntervalSec
	}
	return 0
}

var File_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x6c,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xfc, 0x01, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x20, 0x0a, 0x0c, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6c, 0x73, 0x43, 0x65, 0x72,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x6c, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6c, 0x73,
	0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72,
	0x65, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_goTypes = []interface{}{
	(*TLSConfig)(nil), // 0: cloudprober.tlsconfig.TLSConfig
}
var file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSConfig); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_internal_tlsconfig_proto_config_proto_depIdxs = nil
}
