// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3alpha/http_uri.proto

package envoy_config_core_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f28f389e61fb10f, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.config.core.v3alpha.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3alpha/http_uri.proto", fileDescriptor_1f28f389e61fb10f)
}

var fileDescriptor_1f28f389e61fb10f = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0x02, 0x31,
	0x18, 0xc7, 0x29, 0x10, 0x4f, 0xab, 0xd3, 0x2d, 0x72, 0x98, 0x90, 0x53, 0x07, 0x99, 0xda, 0x04,
	0x26, 0x5d, 0x4c, 0x1a, 0x07, 0x46, 0x42, 0xe2, 0x4c, 0x0a, 0x94, 0xa3, 0xc9, 0xd1, 0xaf, 0xe9,
	0x7d, 0xbd, 0xc8, 0x1b, 0xf8, 0x0c, 0x3e, 0x82, 0xef, 0x61, 0x7c, 0x25, 0xc3, 0x64, 0xb8, 0xf6,
	0x06, 0x07, 0xb7, 0xe6, 0xfb, 0xff, 0xfe, 0xcd, 0x2f, 0x7f, 0x3a, 0x56, 0xa6, 0x86, 0x03, 0x5f,
	0x83, 0xd9, 0xea, 0x82, 0xaf, 0xc1, 0x29, 0x5e, 0x4f, 0x65, 0x69, 0x77, 0x92, 0xef, 0x10, 0xed,
	0xd2, 0x3b, 0xcd, 0xac, 0x03, 0x84, 0x34, 0x6b, 0x48, 0x16, 0x48, 0x76, 0x22, 0x59, 0x24, 0x87,
	0xa3, 0x02, 0xa0, 0x28, 0x15, 0x6f, 0xc0, 0x95, 0xdf, 0xf2, 0x8d, 0x77, 0x12, 0x35, 0x98, 0x50,
	0x1d, 0xde, 0xfa, 0x8d, 0x95, 0x5c, 0x1a, 0x03, 0xd8, 0x9c, 0x2b, 0x5e, 0x2b, 0x57, 0x69, 0x30,
	0xda, 0x14, 0x11, 0xb9, 0xae, 0x65, 0xa9, 0x37, 0x12, 0x15, 0x6f, 0x1f, 0x21, 0xb8, 0xfb, 0x26,
	0x34, 0x99, 0x21, 0xda, 0x57, 0xa7, 0xd3, 0x8c, 0xf6, 0xbc, 0xd3, 0x03, 0x92, 0x93, 0xf1, 0x85,
	0x48, 0x8e, 0xa2, 0xef, 0xba, 0x39, 0x59, 0x9c, 0x6e, 0xe9, 0x3d, 0x4d, 0xd6, 0xa5, 0xaf, 0x50,
	0xb9, 0x41, 0xf7, 0x4f, 0x3c, 0xeb, 0x2c, 0xda, 0x24, 0x7d, 0xa6, 0x09, 0xea, 0xbd, 0x02, 0x8f,
	0x83, 0x5e, 0x4e, 0xc6, 0x97, 0x93, 0x8c, 0x05, 0x73, 0xd6, 0x9a, 0xb3, 0x97, 0x68, 0x2e, 0xe8,
	0x51, 0x24, 0x9f, 0xa4, 0x7f, 0x4e, 0x26, 0x9d, 0x45, 0xdb, 0x7a, 0xca, 0x3f, 0xbe, 0xde, 0x47,
	0x37, 0x34, 0x4e, 0x21, 0xad, 0x66, 0xf5, 0x24, 0x4c, 0x11, 0x15, 0x45, 0x46, 0xd3, 0xb0, 0x9b,
	0xad, 0xd0, 0x29, 0xb9, 0x5f, 0xe2, 0xc1, 0xaa, 0xb4, 0xf7, 0x23, 0x88, 0x78, 0xa4, 0x0f, 0x1a,
	0x58, 0x53, 0xb5, 0x0e, 0xde, 0x0e, 0xec, 0xdf, 0x41, 0xc5, 0x55, 0xfc, 0x6e, 0x7e, 0xd2, 0x9a,
	0x93, 0xd5, 0x59, 0xe3, 0x37, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xcf, 0x65, 0x88, 0xae,
	0x01, 0x00, 0x00,
}
