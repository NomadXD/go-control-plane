// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/ext_authz/v3/ext_authz.proto

package envoy_extensions_filters_http_ext_authz_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExtAuthz struct {
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services                      isExtAuthz_Services          `protobuf_oneof:"services"`
	FailureModeAllow              bool                         `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	HiddenEnvoyDeprecatedUseAlpha bool                         `protobuf:"varint,4,opt,name=hidden_envoy_deprecated_use_alpha,json=hiddenEnvoyDeprecatedUseAlpha,proto3" json:"hidden_envoy_deprecated_use_alpha,omitempty"` // Deprecated: Do not use.
	WithRequestBody               *BufferSettings              `protobuf:"bytes,5,opt,name=with_request_body,json=withRequestBody,proto3" json:"with_request_body,omitempty"`
	ClearRouteCache               bool                         `protobuf:"varint,6,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	StatusOnError                 *v31.HttpStatus              `protobuf:"bytes,7,opt,name=status_on_error,json=statusOnError,proto3" json:"status_on_error,omitempty"`
	MetadataContextNamespaces     []string                     `protobuf:"bytes,8,rep,name=metadata_context_namespaces,json=metadataContextNamespaces,proto3" json:"metadata_context_namespaces,omitempty"`
	FilterEnabled                 *v3.RuntimeFractionalPercent `protobuf:"bytes,9,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	IncludePeerCertificate        bool                         `protobuf:"varint,10,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                     `json:"-"`
	XXX_unrecognized              []byte                       `json:"-"`
	XXX_sizecache                 int32                        `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *v3.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// Deprecated: Do not use.
func (m *ExtAuthz) GetHiddenEnvoyDeprecatedUseAlpha() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseAlpha
	}
	return false
}

func (m *ExtAuthz) GetWithRequestBody() *BufferSettings {
	if m != nil {
		return m.WithRequestBody
	}
	return nil
}

func (m *ExtAuthz) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *ExtAuthz) GetStatusOnError() *v31.HttpStatus {
	if m != nil {
		return m.StatusOnError
	}
	return nil
}

func (m *ExtAuthz) GetMetadataContextNamespaces() []string {
	if m != nil {
		return m.MetadataContextNamespaces
	}
	return nil
}

func (m *ExtAuthz) GetFilterEnabled() *v3.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *ExtAuthz) GetIncludePeerCertificate() bool {
	if m != nil {
		return m.IncludePeerCertificate
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

type BufferSettings struct {
	MaxRequestBytes      uint32   `protobuf:"varint,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	AllowPartialMessage  bool     `protobuf:"varint,2,opt,name=allow_partial_message,json=allowPartialMessage,proto3" json:"allow_partial_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{1}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

func (m *BufferSettings) GetAllowPartialMessage() bool {
	if m != nil {
		return m.AllowPartialMessage
	}
	return false
}

type HttpService struct {
	ServerUri             *v3.HttpUri            `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	PathPrefix            string                 `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	AuthorizationRequest  *AuthorizationRequest  `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{2}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *v3.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	AllowedHeaders       *v32.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	HeadersToAdd         []*v3.HeaderValue      `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{3}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*v3.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	AllowedUpstreamHeaders *v32.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	AllowedClientHeaders   *v32.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{4}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *v32.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{5}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

type CheckSettings struct {
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a074478ef6deb0a, []int{6}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.extensions.filters.http.ext_authz.v3.ExtAuthz")
	proto.RegisterType((*BufferSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3.BufferSettings")
	proto.RegisterType((*HttpService)(nil), "envoy.extensions.filters.http.ext_authz.v3.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.extensions.filters.http.ext_authz.v3.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.extensions.filters.http.ext_authz.v3.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.extensions.filters.http.ext_authz.v3.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.extensions.filters.http.ext_authz.v3.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ext_authz/v3/ext_authz.proto", fileDescriptor_9a074478ef6deb0a)
}

var fileDescriptor_9a074478ef6deb0a = []byte{
	// 1200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xef, 0xda, 0x4e, 0xba, 0x99, 0x34, 0x89, 0x33, 0x24, 0x61, 0x5b, 0x54, 0x9a, 0x1a, 0x21,
	0xa2, 0x0a, 0xad, 0xa1, 0xa6, 0x25, 0xb8, 0x2d, 0xc2, 0x4e, 0xd3, 0x46, 0x81, 0x82, 0xd9, 0x12,
	0x4e, 0x88, 0xd1, 0x64, 0xf7, 0x39, 0x1e, 0xba, 0xde, 0x5d, 0x66, 0x66, 0x5d, 0xbb, 0x17, 0xae,
	0x88, 0x23, 0x47, 0x24, 0xc4, 0x47, 0xe0, 0x86, 0xf8, 0x02, 0x7c, 0x01, 0x3e, 0x08, 0x07, 0x24,
	0x04, 0xea, 0x09, 0xcd, 0x9f, 0xb5, 0x9d, 0x62, 0x44, 0x5c, 0x6e, 0xde, 0x79, 0xef, 0xfd, 0xde,
	0xfb, 0xfd, 0xde, 0x9b, 0x37, 0x46, 0x4d, 0x48, 0x06, 0xe9, 0xa8, 0x0e, 0x43, 0x09, 0x89, 0x60,
	0x69, 0x22, 0xea, 0x5d, 0x16, 0x4b, 0xe0, 0xa2, 0xde, 0x93, 0x32, 0x53, 0xe7, 0x84, 0xe6, 0xb2,
	0xf7, 0xa4, 0x3e, 0x68, 0x4c, 0x3e, 0xfc, 0x8c, 0xa7, 0x32, 0xc5, 0xd7, 0x74, 0xac, 0x3f, 0x89,
	0xf5, 0x6d, 0xac, 0xaf, 0x62, 0xfd, 0x89, 0xfb, 0xa0, 0x71, 0xe9, 0x8a, 0xc9, 0x13, 0xa6, 0x49,
	0x97, 0x9d, 0xd4, 0xc3, 0x94, 0x83, 0x42, 0x3c, 0xa6, 0x02, 0x0c, 0xd8, 0xa5, 0xd7, 0x66, 0x3a,
	0x9c, 0xf0, 0x2c, 0x24, 0x02, 0xf8, 0x80, 0x85, 0x85, 0xe3, 0x2b, 0x33, 0x1d, 0x55, 0x46, 0x92,
	0x73, 0x66, 0x9d, 0x6a, 0xc6, 0x49, 0x8e, 0x32, 0xa8, 0xf7, 0xa9, 0x0c, 0x7b, 0xc0, 0x95, 0x97,
	0x90, 0x9c, 0x25, 0x27, 0xd6, 0xe7, 0xca, 0x94, 0x4f, 0x81, 0x20, 0x24, 0x95, 0xb9, 0xb0, 0x0e,
	0x57, 0xf3, 0x28, 0xa3, 0x75, 0x9a, 0x24, 0xa9, 0xa4, 0x52, 0x6b, 0x33, 0x00, 0xae, 0x88, 0x4e,
	0x30, 0x6c, 0x31, 0xd3, 0x3e, 0x11, 0x64, 0x1c, 0x42, 0xfd, 0x61, 0x9d, 0x5e, 0x1c, 0xd0, 0x98,
	0x45, 0x54, 0x42, 0xbd, 0xf8, 0x61, 0x0c, 0xb5, 0xef, 0x17, 0x91, 0xbb, 0x3f, 0x94, 0x2d, 0x25,
	0x12, 0xbe, 0x87, 0x2e, 0x4c, 0xb3, 0xf5, 0x9c, 0x6d, 0x67, 0x67, 0xf9, 0xfa, 0x55, 0xdf, 0x88,
	0x6c, 0xe8, 0xfa, 0x8a, 0xae, 0x3f, 0x68, 0xf8, 0xf7, 0x79, 0x16, 0x3e, 0x34, 0x8e, 0x07, 0xe7,
	0x82, 0xe5, 0x93, 0xc9, 0x27, 0xfe, 0x0c, 0x5d, 0x30, 0x54, 0x2c, 0x4e, 0x59, 0xe3, 0xbc, 0xed,
	0x9f, 0xbd, 0x59, 0xfe, 0x81, 0x94, 0xd9, 0x14, 0x7a, 0x6f, 0xf2, 0x89, 0x5f, 0x47, 0xb8, 0x4b,
	0x59, 0x9c, 0x73, 0x20, 0xfd, 0x34, 0x02, 0x42, 0xe3, 0x38, 0x7d, 0xec, 0x95, 0xb6, 0x9d, 0x1d,
	0x37, 0xa8, 0x5a, 0xcb, 0x83, 0x34, 0x82, 0x96, 0x3a, 0xc7, 0x01, 0xba, 0xda, 0x63, 0x51, 0x04,
	0x09, 0xd1, 0xd9, 0x49, 0xa1, 0x0d, 0x44, 0x24, 0x17, 0x2a, 0x38, 0xeb, 0x51, 0xaf, 0xa2, 0x82,
	0xdb, 0xee, 0xcf, 0xbf, 0xfd, 0xfe, 0xeb, 0x82, 0xe3, 0x39, 0xc1, 0x65, 0x13, 0xb2, 0xaf, 0x22,
	0xee, 0x8e, 0x03, 0x8e, 0x04, 0xb4, 0x94, 0x3b, 0xee, 0xa2, 0xf5, 0xc7, 0x4c, 0xf6, 0x08, 0x87,
	0x2f, 0x73, 0x10, 0x92, 0x1c, 0xa7, 0xd1, 0xc8, 0x5b, 0xd0, 0x24, 0x9b, 0xf3, 0x90, 0x6c, 0xe7,
	0xdd, 0x2e, 0xf0, 0x87, 0x20, 0x25, 0x4b, 0x4e, 0x44, 0xb0, 0xa6, 0x40, 0x03, 0x83, 0xd9, 0x4e,
	0xa3, 0x11, 0xbe, 0x86, 0xd6, 0xc3, 0x18, 0x28, 0x27, 0x3c, 0xcd, 0x25, 0x90, 0x90, 0x86, 0x3d,
	0xf0, 0x16, 0x35, 0xd1, 0x35, 0x6d, 0x08, 0xd4, 0xf9, 0x9e, 0x3a, 0xc6, 0x2d, 0xb4, 0x66, 0x26,
	0x87, 0xa4, 0x09, 0x01, 0xce, 0x53, 0xee, 0x9d, 0xd7, 0x15, 0x5d, 0xb4, 0x15, 0xa9, 0x21, 0x1b,
	0x2b, 0xab, 0x3d, 0x83, 0x15, 0x13, 0xf1, 0x51, 0xb2, 0xaf, 0xfc, 0xf1, 0xbb, 0xe8, 0xa5, 0x3e,
	0x48, 0x1a, 0x51, 0x49, 0x49, 0x98, 0x26, 0x52, 0xd5, 0x9a, 0xd0, 0x3e, 0x88, 0x8c, 0x86, 0x20,
	0x3c, 0x77, 0xbb, 0xbc, 0xb3, 0x14, 0x5c, 0x2c, 0x5c, 0xf6, 0x8c, 0xc7, 0x87, 0x63, 0x07, 0x7c,
	0x84, 0x56, 0x0d, 0x57, 0x02, 0x09, 0x3d, 0x8e, 0x21, 0xf2, 0x96, 0x74, 0x05, 0xfe, 0xec, 0x01,
	0x0a, 0xf2, 0x44, 0xb2, 0x3e, 0xdc, 0xe3, 0x34, 0x54, 0x93, 0x4a, 0xe3, 0x0e, 0xf0, 0x10, 0x12,
	0x19, 0xac, 0x18, 0x94, 0x7d, 0x03, 0x82, 0x77, 0x91, 0xc7, 0x92, 0x30, 0xce, 0x23, 0x20, 0x19,
	0x00, 0x27, 0x21, 0x70, 0xc9, 0xba, 0x4c, 0xb5, 0xc4, 0x43, 0x5a, 0x8c, 0x2d, 0x6b, 0xef, 0x00,
	0xf0, 0xbd, 0x89, 0xb5, 0x79, 0xe3, 0xbb, 0x5f, 0xbe, 0x7e, 0xf9, 0x0d, 0x74, 0x3a, 0xbd, 0x01,
	0xff, 0x47, 0x37, 0xae, 0xfb, 0xc5, 0x35, 0x68, 0x23, 0xe4, 0xda, 0xc9, 0x15, 0xb5, 0x9f, 0x1c,
	0xb4, 0x7a, 0xba, 0x4d, 0xb8, 0x81, 0xd6, 0xfb, 0x74, 0x38, 0x69, 0xfe, 0x48, 0x82, 0xd0, 0x57,
	0x65, 0xa5, 0x7d, 0xfe, 0x69, 0xbb, 0x72, 0xad, 0xb4, 0x7d, 0x2e, 0x58, 0xeb, 0xd3, 0x61, 0xd1,
	0x49, 0x65, 0xc7, 0xd7, 0xd1, 0xa6, 0x9e, 0x53, 0x92, 0x51, 0x2e, 0x19, 0x8d, 0x49, 0x1f, 0x84,
	0xa0, 0x27, 0x60, 0xe7, 0xf6, 0x05, 0x6d, 0xec, 0x18, 0xdb, 0x03, 0x63, 0x6a, 0xde, 0x52, 0xe5,
	0xdf, 0x44, 0x6f, 0x9d, 0xad, 0xfc, 0xd3, 0x55, 0xd6, 0x7e, 0x2c, 0xa3, 0xe5, 0xa9, 0x4b, 0x84,
	0x6f, 0x23, 0xa4, 0x48, 0x01, 0x57, 0x2b, 0xca, 0xde, 0xec, 0xcb, 0xb3, 0x1b, 0xa3, 0xc2, 0x8e,
	0x38, 0x0b, 0x96, 0x4c, 0xc0, 0x11, 0x67, 0xf8, 0x0a, 0x5a, 0xce, 0xa8, 0xec, 0x91, 0x8c, 0x43,
	0x97, 0x0d, 0x75, 0xd1, 0x4b, 0x01, 0x52, 0x47, 0x1d, 0x7d, 0x82, 0x73, 0xb4, 0xa9, 0x2a, 0x49,
	0x39, 0x7b, 0xa2, 0xf7, 0x4e, 0x21, 0x8f, 0x1d, 0xc2, 0xf7, 0xe6, 0xb9, 0x16, 0xad, 0x69, 0x20,
	0xab, 0x62, 0xb0, 0x41, 0x67, 0x9c, 0xe2, 0x21, 0xda, 0x7a, 0x36, 0xad, 0xc8, 0xd2, 0x44, 0x80,
	0xe7, 0xea, 0xbc, 0xad, 0xff, 0x91, 0xd7, 0x00, 0x05, 0x9b, 0x74, 0xd6, 0x71, 0x73, 0x57, 0x35,
	0xa7, 0x81, 0xde, 0x3c, 0x5b, 0x73, 0xa6, 0x3a, 0x71, 0x58, 0x71, 0xcb, 0xd5, 0xca, 0x61, 0xc5,
	0xad, 0x54, 0x17, 0x0e, 0x2b, 0xee, 0x42, 0x75, 0xf1, 0xb0, 0xe2, 0x2e, 0x56, 0xcf, 0xd7, 0xfe,
	0x70, 0xd0, 0xc6, 0x2c, 0xea, 0xf8, 0x63, 0xb4, 0xa6, 0xc7, 0x03, 0x22, 0xd2, 0x03, 0x1a, 0x01,
	0x17, 0xb6, 0x7f, 0x3b, 0xd3, 0x57, 0xdb, 0xbe, 0x31, 0x8a, 0xc8, 0x07, 0x4c, 0xc8, 0x87, 0xfa,
	0x9d, 0x79, 0x60, 0x0e, 0x83, 0x55, 0x0b, 0x70, 0x60, 0xe2, 0xf1, 0x7d, 0xb4, 0x6a, 0xa1, 0x88,
	0x4c, 0x09, 0x8d, 0x22, 0xaf, 0xb4, 0x5d, 0xfe, 0xf7, 0x5d, 0x6f, 0xc2, 0x3e, 0xa5, 0x71, 0x0e,
	0xc1, 0x05, 0x1b, 0xf8, 0x49, 0xda, 0x8a, 0xa2, 0x66, 0x4b, 0xc9, 0x70, 0xdb, 0xbe, 0xe1, 0xff,
	0x29, 0xc3, 0x2c, 0x7a, 0xb5, 0x1f, 0x4a, 0x68, 0x73, 0xa6, 0xf4, 0xf8, 0x18, 0x79, 0x05, 0xf1,
	0x3c, 0x13, 0x92, 0x03, 0xed, 0x3f, 0xb7, 0x02, 0x5b, 0x16, 0xe9, 0xc8, 0x02, 0x15, 0x4a, 0x7c,
	0x8e, 0x0a, 0x0b, 0x09, 0x63, 0x06, 0x89, 0x1c, 0x67, 0x28, 0xcd, 0x99, 0x61, 0xc3, 0xe2, 0xec,
	0x69, 0x18, 0x8b, 0xdf, 0x6c, 0x2b, 0x81, 0xee, 0xa0, 0x5b, 0xcf, 0x25, 0x90, 0xd1, 0xa1, 0xf6,
	0xa7, 0x83, 0xaa, 0xc5, 0x76, 0xea, 0x80, 0xd9, 0xfa, 0xf8, 0x55, 0xe4, 0x46, 0x4c, 0x98, 0x3d,
	0xeb, 0xe8, 0xf7, 0x4b, 0x6d, 0x9f, 0x2f, 0x4a, 0xae, 0x73, 0x70, 0x2e, 0x18, 0x9b, 0x70, 0x8c,
	0x56, 0xc3, 0x1e, 0x84, 0x8f, 0x88, 0xb0, 0x9b, 0xc1, 0xf2, 0x7a, 0x67, 0x9e, 0x9b, 0xb1, 0xa7,
	0x10, 0x8a, 0xd5, 0xd2, 0x76, 0x9f, 0xb6, 0x17, 0xbe, 0x71, 0x4a, 0x55, 0x95, 0x68, 0x25, 0x9c,
	0x36, 0x35, 0xef, 0x28, 0xb6, 0xbb, 0xe8, 0xe6, 0x7c, 0x1b, 0xb7, 0xe0, 0xd4, 0x5e, 0x43, 0x6e,
	0x3a, 0x00, 0xce, 0x59, 0x04, 0xb8, 0xfc, 0x57, 0xdb, 0xa9, 0x7d, 0x5b, 0x42, 0x2b, 0xa7, 0x92,
	0xe3, 0xaf, 0x10, 0x2e, 0xde, 0xa6, 0x49, 0xe9, 0x9e, 0xa3, 0xa7, 0xb7, 0xf3, 0xdc, 0x9c, 0x7c,
	0xfb, 0x9a, 0xed, 0x8f, 0x23, 0xf7, 0x13, 0xc9, 0x47, 0xc1, 0x7a, 0xf8, 0xec, 0xf9, 0xa5, 0xbb,
	0x68, 0x6b, 0xb6, 0x33, 0xae, 0xa2, 0xf2, 0x23, 0x18, 0xe9, 0x66, 0x2c, 0x05, 0xea, 0x27, 0xde,
	0x40, 0x0b, 0x03, 0x75, 0x69, 0xec, 0xc2, 0x34, 0x1f, 0xcd, 0xd2, 0xae, 0xd3, 0x6c, 0x2a, 0xa1,
	0x6e, 0xa0, 0xc6, 0xd9, 0x84, 0x3a, 0xad, 0xff, 0xfb, 0x68, 0x97, 0xa5, 0x86, 0x6a, 0xc6, 0xd3,
	0xe1, 0x68, 0x0e, 0xd6, 0xed, 0x95, 0xb1, 0xe6, 0xea, 0xef, 0x5f, 0xc7, 0x39, 0x5e, 0xd4, 0xff,
	0x03, 0x1b, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xeb, 0x21, 0xa6, 0x86, 0x0b, 0x00, 0x00,
}