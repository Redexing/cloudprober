// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.27.5
// source: github.com/cloudprober/cloudprober/internal/oauth/proto/config.proto

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

type Config struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*Config_HttpRequest
	//	*Config_BearerToken
	//	*Config_GoogleCredentials
	Type isConfig_Type `protobuf_oneof:"type"`
	// How long before the expiry do we refresh. Default is 60 (1m). This applies
	// only to http_request and bearer_token types, and only if token presents
	// expiry in some way.
	// TODO(manugarg): Consider setting default based on probe interval.
	RefreshExpiryBufferSec *int32 `protobuf:"varint,20,opt,name=refresh_expiry_buffer_sec,json=refreshExpiryBufferSec,proto3,oneof" json:"refresh_expiry_buffer_sec,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetType() isConfig_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Config) GetHttpRequest() *HTTPRequest {
	if x != nil {
		if x, ok := x.Type.(*Config_HttpRequest); ok {
			return x.HttpRequest
		}
	}
	return nil
}

func (x *Config) GetBearerToken() *BearerToken {
	if x != nil {
		if x, ok := x.Type.(*Config_BearerToken); ok {
			return x.BearerToken
		}
	}
	return nil
}

func (x *Config) GetGoogleCredentials() *GoogleCredentials {
	if x != nil {
		if x, ok := x.Type.(*Config_GoogleCredentials); ok {
			return x.GoogleCredentials
		}
	}
	return nil
}

func (x *Config) GetRefreshExpiryBufferSec() int32 {
	if x != nil && x.RefreshExpiryBufferSec != nil {
		return *x.RefreshExpiryBufferSec
	}
	return 0
}

type isConfig_Type interface {
	isConfig_Type()
}

type Config_HttpRequest struct {
	HttpRequest *HTTPRequest `protobuf:"bytes,3,opt,name=http_request,json=httpRequest,proto3,oneof"`
}

type Config_BearerToken struct {
	BearerToken *BearerToken `protobuf:"bytes,1,opt,name=bearer_token,json=bearerToken,proto3,oneof"`
}

type Config_GoogleCredentials struct {
	GoogleCredentials *GoogleCredentials `protobuf:"bytes,2,opt,name=google_credentials,json=googleCredentials,proto3,oneof"`
}

func (*Config_HttpRequest) isConfig_Type() {}

func (*Config_BearerToken) isConfig_Type() {}

func (*Config_GoogleCredentials) isConfig_Type() {}

type HTTPRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	TokenUrl string                 `protobuf:"bytes,1,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	Method   string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// Data to be sent as request body. If there are multiple "data" fields, we combine
	// their values with a '&' in between. Note: 1) If data appears to be a valid json,
	// we automatically set the content-type header to "application/json", 2) If data
	// appears to be a query string we set content-type to
	// "application/x-www-form-urlencoded". Content type header can still be overridden
	// using the header field below.
	Data []string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	// HTTP request headers
	Header        map[string]string `protobuf:"bytes,8,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPRequest) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *HTTPRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRequest) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *HTTPRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

// Bearer token is added to the HTTP request through an HTTP header:
// "Authorization: Bearer <access_token>"
type BearerToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Source:
	//
	//	*BearerToken_File
	//	*BearerToken_Cmd
	//	*BearerToken_GceServiceAccount
	//	*BearerToken_K8SLocalToken
	Source isBearerToken_Source `protobuf_oneof:"source"`
	// If above sources return JSON tokens with an expiry, we use that info to
	// determine when to refresh tokens and refresh_interval_sec is completely
	// ignored. If above sources return a string, we refresh from the source
	// every 30s by default. To disable this behavior set refresh_interval_sec to
	// zero.
	RefreshIntervalSec *float32 `protobuf:"fixed32,90,opt,name=refresh_interval_sec,json=refreshIntervalSec,proto3,oneof" json:"refresh_interval_sec,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *BearerToken) Reset() {
	*x = BearerToken{}
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BearerToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BearerToken) ProtoMessage() {}

func (x *BearerToken) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BearerToken.ProtoReflect.Descriptor instead.
func (*BearerToken) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescGZIP(), []int{2}
}

func (x *BearerToken) GetSource() isBearerToken_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *BearerToken) GetFile() string {
	if x != nil {
		if x, ok := x.Source.(*BearerToken_File); ok {
			return x.File
		}
	}
	return ""
}

func (x *BearerToken) GetCmd() string {
	if x != nil {
		if x, ok := x.Source.(*BearerToken_Cmd); ok {
			return x.Cmd
		}
	}
	return ""
}

func (x *BearerToken) GetGceServiceAccount() string {
	if x != nil {
		if x, ok := x.Source.(*BearerToken_GceServiceAccount); ok {
			return x.GceServiceAccount
		}
	}
	return ""
}

func (x *BearerToken) GetK8SLocalToken() bool {
	if x != nil {
		if x, ok := x.Source.(*BearerToken_K8SLocalToken); ok {
			return x.K8SLocalToken
		}
	}
	return false
}

func (x *BearerToken) GetRefreshIntervalSec() float32 {
	if x != nil && x.RefreshIntervalSec != nil {
		return *x.RefreshIntervalSec
	}
	return 0
}

type isBearerToken_Source interface {
	isBearerToken_Source()
}

type BearerToken_File struct {
	// Path to token file.
	File string `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type BearerToken_Cmd struct {
	// Run a comand to obtain the token, e.g.
	// cat /var/lib/myapp/token, or
	// /var/lib/run/get_token.sh
	Cmd string `protobuf:"bytes,2,opt,name=cmd,proto3,oneof"`
}

type BearerToken_GceServiceAccount struct {
	// GCE metadata token
	GceServiceAccount string `protobuf:"bytes,3,opt,name=gce_service_account,json=gceServiceAccount,proto3,oneof"`
}

type BearerToken_K8SLocalToken struct {
	// K8s service account token file:
	// /var/run/secrets/kubernetes.io/serviceaccount/token
	K8SLocalToken bool `protobuf:"varint,4,opt,name=k8s_local_token,json=k8sLocalToken,proto3,oneof"`
}

func (*BearerToken_File) isBearerToken_Source() {}

func (*BearerToken_Cmd) isBearerToken_Source() {}

func (*BearerToken_GceServiceAccount) isBearerToken_Source() {}

func (*BearerToken_K8SLocalToken) isBearerToken_Source() {}

// Google credentials in JSON format. We simply use oauth2/google package to
// use these credentials.
type GoogleCredentials struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	JsonFile string                 `protobuf:"bytes,1,opt,name=json_file,json=jsonFile,proto3" json:"json_file,omitempty"`
	Scope    []string               `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	// Use encoded JWT directly as access token, instead of implementing the whole
	// OAuth2.0 flow.
	JwtAsAccessToken bool `protobuf:"varint,4,opt,name=jwt_as_access_token,json=jwtAsAccessToken,proto3" json:"jwt_as_access_token,omitempty"`
	// Audience works only if jwt_as_access_token is true.
	Audience      string `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GoogleCredentials) Reset() {
	*x = GoogleCredentials{}
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GoogleCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleCredentials) ProtoMessage() {}

func (x *GoogleCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleCredentials.ProtoReflect.Descriptor instead.
func (*GoogleCredentials) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescGZIP(), []int{3}
}

func (x *GoogleCredentials) GetJsonFile() string {
	if x != nil {
		return x.JsonFile
	}
	return ""
}

func (x *GoogleCredentials) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *GoogleCredentials) GetJwtAsAccessToken() bool {
	if x != nil {
		return x.JwtAsAccessToken
	}
	return false
}

func (x *GoogleCredentials) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

var File_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDesc = []byte{
	0x0a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x22, 0xcf, 0x02, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48,
	0x00, 0x52, 0x0b, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x55,
	0x0a, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x48, 0x00, 0x52, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x19, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x16, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x1c, 0x0a,
	0x1a, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x22, 0xd5, 0x01, 0x0a, 0x0b,
	0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xed, 0x01, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x30, 0x0a,
	0x13, 0x67, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x67, 0x63,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x6b, 0x38, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x6b, 0x38, 0x73, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x14, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x12, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73,
	0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x13,
	0x6a, 0x77, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6a, 0x77, 0x74, 0x41, 0x73,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_goTypes = []any{
	(*Config)(nil),            // 0: cloudprober.oauth.Config
	(*HTTPRequest)(nil),       // 1: cloudprober.oauth.HTTPRequest
	(*BearerToken)(nil),       // 2: cloudprober.oauth.BearerToken
	(*GoogleCredentials)(nil), // 3: cloudprober.oauth.GoogleCredentials
	nil,                       // 4: cloudprober.oauth.HTTPRequest.HeaderEntry
}
var file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_depIdxs = []int32{
	1, // 0: cloudprober.oauth.Config.http_request:type_name -> cloudprober.oauth.HTTPRequest
	2, // 1: cloudprober.oauth.Config.bearer_token:type_name -> cloudprober.oauth.BearerToken
	3, // 2: cloudprober.oauth.Config.google_credentials:type_name -> cloudprober.oauth.GoogleCredentials
	4, // 3: cloudprober.oauth.HTTPRequest.header:type_name -> cloudprober.oauth.HTTPRequest.HeaderEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto != nil {
		return
	}
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[0].OneofWrappers = []any{
		(*Config_HttpRequest)(nil),
		(*Config_BearerToken)(nil),
		(*Config_GoogleCredentials)(nil),
	}
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes[2].OneofWrappers = []any{
		(*BearerToken_File)(nil),
		(*BearerToken_Cmd)(nil),
		(*BearerToken_GceServiceAccount)(nil),
		(*BearerToken_K8SLocalToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_internal_oauth_proto_config_proto_depIdxs = nil
}
