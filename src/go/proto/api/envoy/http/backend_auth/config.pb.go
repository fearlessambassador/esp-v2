// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/envoy/http/backend_auth/config.proto

package google_api_envoy_http_backend_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

type BackendAuthRule struct {
	Operation            string   `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	JwtAudience          string   `protobuf:"bytes,2,opt,name=jwt_audience,json=jwtAudience,proto3" json:"jwt_audience,omitempty"`
	TokenCluster         string   `protobuf:"bytes,3,opt,name=token_cluster,json=tokenCluster,proto3" json:"token_cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendAuthRule) Reset()         { *m = BackendAuthRule{} }
func (m *BackendAuthRule) String() string { return proto.CompactTextString(m) }
func (*BackendAuthRule) ProtoMessage()    {}
func (*BackendAuthRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_f545d61a25950416, []int{0}
}

func (m *BackendAuthRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendAuthRule.Unmarshal(m, b)
}
func (m *BackendAuthRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendAuthRule.Marshal(b, m, deterministic)
}
func (m *BackendAuthRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendAuthRule.Merge(m, src)
}
func (m *BackendAuthRule) XXX_Size() int {
	return xxx_messageInfo_BackendAuthRule.Size(m)
}
func (m *BackendAuthRule) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendAuthRule.DiscardUnknown(m)
}

var xxx_messageInfo_BackendAuthRule proto.InternalMessageInfo

func (m *BackendAuthRule) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *BackendAuthRule) GetJwtAudience() string {
	if m != nil {
		return m.JwtAudience
	}
	return ""
}

func (m *BackendAuthRule) GetTokenCluster() string {
	if m != nil {
		return m.TokenCluster
	}
	return ""
}

type FilterConfig struct {
	Rules                []*BackendAuthRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f545d61a25950416, []int{1}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetRules() []*BackendAuthRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*BackendAuthRule)(nil), "google.api.envoy.http.backend_auth.BackendAuthRule")
	proto.RegisterType((*FilterConfig)(nil), "google.api.envoy.http.backend_auth.FilterConfig")
}

func init() {
	proto.RegisterFile("api/envoy/http/backend_auth/config.proto", fileDescriptor_f545d61a25950416)
}

var fileDescriptor_f545d61a25950416 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x99, 0x56, 0x85, 0x4c, 0x23, 0xc2, 0x6c, 0x0c, 0xae, 0x42, 0x36, 0x66, 0x21, 0x33,
	0x60, 0x4f, 0xd0, 0x16, 0x04, 0xb7, 0xd9, 0xb9, 0x0a, 0xd3, 0xe4, 0x99, 0x4c, 0x3b, 0xcc, 0x1b,
	0xa6, 0x6f, 0x5a, 0xbc, 0x89, 0x67, 0x71, 0xe5, 0x75, 0xbc, 0x85, 0x98, 0x28, 0x4a, 0x36, 0xee,
	0x1e, 0x7c, 0xff, 0xf7, 0xe0, 0xe3, 0xa5, 0xf6, 0x46, 0x81, 0x3b, 0xe2, 0x8b, 0xea, 0x89, 0xbc,
	0xda, 0xea, 0x66, 0x0f, 0xae, 0xad, 0x75, 0xa4, 0x5e, 0x35, 0xe8, 0x9e, 0x4d, 0x27, 0x7d, 0x40,
	0x42, 0x51, 0x74, 0x88, 0x9d, 0x05, 0xa9, 0xbd, 0x91, 0x83, 0x20, 0xbf, 0x04, 0xf9, 0x57, 0xb8,
	0xb9, 0x3e, 0x6a, 0x6b, 0x5a, 0x4d, 0xa0, 0x7e, 0x8e, 0x51, 0x2e, 0x5e, 0x19, 0xbf, 0x5a, 0x8f,
	0xcb, 0x55, 0xa4, 0xbe, 0x8a, 0x16, 0xc4, 0x2d, 0x4f, 0xd0, 0x43, 0xd0, 0x64, 0xd0, 0x65, 0x2c,
	0x67, 0x65, 0xb2, 0x4e, 0xde, 0x3e, 0xde, 0xe7, 0x67, 0x61, 0x96, 0xb3, 0xea, 0x97, 0x89, 0x3b,
	0x9e, 0xee, 0x4e, 0x54, 0xeb, 0xd8, 0x1a, 0x70, 0x0d, 0x64, 0xb3, 0xe9, 0x76, 0xb1, 0x3b, 0xd1,
	0xea, 0x9b, 0x0a, 0xc9, 0x2f, 0x09, 0xf7, 0xe0, 0xea, 0xc6, 0xc6, 0x03, 0x41, 0xc8, 0xe6, 0xd3,
	0x79, 0x3a, 0xf0, 0xcd, 0x88, 0x8b, 0x27, 0x9e, 0x3e, 0x18, 0x4b, 0x10, 0x36, 0x43, 0xad, 0x78,
	0xe4, 0xe7, 0x21, 0x5a, 0x38, 0x64, 0x2c, 0x9f, 0x97, 0x8b, 0xfb, 0xa5, 0xfc, 0xbf, 0x5b, 0x4e,
	0xd2, 0xaa, 0xf1, 0xc3, 0xf6, 0x62, 0x88, 0x5f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x31, 0xcb,
	0x6a, 0x89, 0x65, 0x01, 0x00, 0x00,
}