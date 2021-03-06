// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/rbac/v3alpha/rbac.proto

package envoy_extensions_filters_http_rbac_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3alpha"
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

type RBAC struct {
	Rules                *v3alpha.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v3alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f7401660038e7c, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f7401660038e7c, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.extensions.filters.http.rbac.v3alpha.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.extensions.filters.http.rbac.v3alpha.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/rbac/v3alpha/rbac.proto", fileDescriptor_89f7401660038e7c)
}

var fileDescriptor_89f7401660038e7c = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xc9, 0xa8, 0xa2, 0xd9, 0x0e, 0xd2, 0x8b, 0xb2, 0x83, 0x7f, 0x86, 0xa2, 0x88, 0x26,
	0xb2, 0x39, 0x90, 0xdd, 0xac, 0x82, 0xe0, 0xa9, 0xf4, 0x0b, 0x48, 0xda, 0x66, 0x6b, 0xa0, 0x24,
	0x25, 0x49, 0xeb, 0xf6, 0x0d, 0x3c, 0x7a, 0xf6, 0x4b, 0xf8, 0x29, 0xfc, 0x5e, 0x92, 0xbc, 0x1d,
	0xba, 0x83, 0x50, 0x6f, 0x29, 0x7d, 0x7e, 0xef, 0xf3, 0xe3, 0x7d, 0xf1, 0x94, 0xcb, 0x46, 0xad,
	0x28, 0x5f, 0x5a, 0x2e, 0x8d, 0x50, 0xd2, 0xd0, 0xb9, 0x28, 0x2d, 0xd7, 0x86, 0x16, 0xd6, 0x56,
	0x54, 0xa7, 0x2c, 0xa3, 0xcd, 0x84, 0x95, 0x55, 0xc1, 0xfc, 0x07, 0xa9, 0xb4, 0xb2, 0x2a, 0xbc,
	0xf4, 0x18, 0xf9, 0xc1, 0x48, 0x8b, 0x11, 0x87, 0x11, 0x9f, 0x6c, 0xb1, 0xe1, 0x29, 0x54, 0x64,
	0x4a, 0xce, 0xc5, 0xe2, 0xaf, 0x89, 0xc3, 0x93, 0x3a, 0xaf, 0x18, 0x65, 0x52, 0x2a, 0xcb, 0xac,
	0x17, 0x69, 0xb8, 0x76, 0xa3, 0x85, 0x5c, 0xb4, 0x91, 0xfd, 0x86, 0x95, 0x22, 0x67, 0x96, 0xd3,
	0xf5, 0x03, 0x7e, 0x8c, 0x3e, 0x11, 0x0e, 0x92, 0xe8, 0xfe, 0x21, 0x9c, 0xe2, 0x2d, 0x5d, 0x97,
	0xdc, 0x1c, 0xa0, 0x63, 0x74, 0xd1, 0x1f, 0x1f, 0x11, 0xd0, 0x84, 0xea, 0x0d, 0x2b, 0xe2, 0xf2,
	0x09, 0xa4, 0xc3, 0x08, 0x0f, 0x4c, 0xc1, 0x72, 0xf5, 0xfa, 0x02, 0x74, 0xaf, 0x1b, 0xdd, 0x07,
	0x28, 0x71, 0xcc, 0xec, 0xea, 0xe3, 0xeb, 0xed, 0xf0, 0x1c, 0x9f, 0x6d, 0x30, 0xb0, 0x94, 0xdf,
	0x3b, 0x19, 0x7b, 0x74, 0xf4, 0x8e, 0xf0, 0xc0, 0x3d, 0x62, 0xae, 0x13, 0x55, 0x5b, 0x1e, 0x3e,
	0xe2, 0xc0, 0x05, 0xda, 0xea, 0x1b, 0xd2, 0x7d, 0xbf, 0xe0, 0xe2, 0xe9, 0xd9, 0xad, 0x93, 0xa0,
	0xf8, 0xba, 0x93, 0xc4, 0xba, 0xfb, 0x39, 0xd8, 0x41, 0x7b, 0xbd, 0xe8, 0x09, 0xdf, 0x09, 0x05,
	0xbd, 0x95, 0x56, 0xcb, 0xd5, 0x3f, 0x14, 0xa2, 0xdd, 0x24, 0x65, 0x59, 0xec, 0x6e, 0x11, 0xa3,
	0x74, 0xdb, 0x1f, 0x65, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa0, 0x61, 0x37, 0x5b, 0x02,
	0x00, 0x00,
}
