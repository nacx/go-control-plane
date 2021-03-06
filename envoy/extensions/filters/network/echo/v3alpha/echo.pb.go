// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/echo/v3alpha/echo.proto

package envoy_extensions_filters_network_echo_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Echo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Echo) Reset()         { *m = Echo{} }
func (m *Echo) String() string { return proto.CompactTextString(m) }
func (*Echo) ProtoMessage()    {}
func (*Echo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e94fe45086e956, []int{0}
}

func (m *Echo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Echo.Unmarshal(m, b)
}
func (m *Echo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Echo.Marshal(b, m, deterministic)
}
func (m *Echo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Echo.Merge(m, src)
}
func (m *Echo) XXX_Size() int {
	return xxx_messageInfo_Echo.Size(m)
}
func (m *Echo) XXX_DiscardUnknown() {
	xxx_messageInfo_Echo.DiscardUnknown(m)
}

var xxx_messageInfo_Echo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Echo)(nil), "envoy.extensions.filters.network.echo.v3alpha.Echo")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/echo/v3alpha/echo.proto", fileDescriptor_e5e94fe45086e956)
}

var fileDescriptor_e5e94fe45086e956 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0x4d, 0xce,
	0xc8, 0xd7, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x04, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x74, 0xc1, 0x3a, 0xf5, 0x10, 0x3a, 0xf5, 0xa0, 0x3a, 0xf5, 0xa0, 0x3a, 0xf5, 0xc0,
	0x8a, 0xa1, 0x3a, 0xa5, 0x14, 0x4b, 0x53, 0x0a, 0x12, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0xc0, 0x16, 0x95, 0xa5, 0x16, 0x81, 0xf4, 0x65, 0xe6, 0xa5, 0x43, 0x4c, 0x54, 0x32, 0xe7,
	0x62, 0x71, 0x4d, 0xce, 0xc8, 0xb7, 0xd2, 0x9f, 0x75, 0xb4, 0x43, 0x4e, 0x8b, 0x4b, 0x03, 0x62,
	0x41, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0x3a, 0xd4, 0x70, 0x34, 0xb3, 0x8d, 0xf4, 0x40, 0x1a, 0x9c,
	0xbc, 0xb8, 0xac, 0x33, 0xf3, 0xf5, 0xc0, 0xca, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x48, 0x72,
	0x9a, 0x13, 0x27, 0xc8, 0x90, 0x00, 0x90, 0x13, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x6e, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x85, 0x30, 0x1b, 0xd7, 0x19, 0x01, 0x00, 0x00,
}
