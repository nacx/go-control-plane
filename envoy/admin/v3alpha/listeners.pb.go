// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3alpha/listeners.proto

package envoy_admin_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
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

type Listeners struct {
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_088657d81081dfcb, []int{0}
}

func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listeners.Unmarshal(m, b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return xxx_messageInfo_Listeners.Size(m)
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

type ListenerStatus struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LocalAddress         *v3alpha.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_088657d81081dfcb, []int{1}
}

func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerStatus.Unmarshal(m, b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return xxx_messageInfo_ListenerStatus.Size(m)
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() *v3alpha.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Listeners)(nil), "envoy.admin.v3alpha.Listeners")
	proto.RegisterType((*ListenerStatus)(nil), "envoy.admin.v3alpha.ListenerStatus")
}

func init() {
	proto.RegisterFile("envoy/admin/v3alpha/listeners.proto", fileDescriptor_088657d81081dfcb)
}

var fileDescriptor_088657d81081dfcb = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x99, 0x7b, 0x45, 0xe8, 0x54, 0x45, 0xe3, 0xa6, 0x14, 0xd4, 0x36, 0x2d, 0x18, 0x37,
	0x33, 0x90, 0xe2, 0xa6, 0x3b, 0xbb, 0x71, 0xe3, 0x22, 0xc4, 0x07, 0x28, 0x63, 0x32, 0xd6, 0x81,
	0x74, 0x4e, 0x98, 0x33, 0x0d, 0x76, 0xed, 0xc6, 0x47, 0x10, 0xdf, 0xc7, 0xf7, 0x92, 0xcc, 0x24,
	0x81, 0x40, 0x71, 0x37, 0x07, 0xfe, 0x6f, 0xce, 0x77, 0x7e, 0x3a, 0x93, 0xba, 0x82, 0x3d, 0x17,
	0xf9, 0x56, 0x69, 0x5e, 0x2d, 0x44, 0x51, 0xbe, 0x09, 0x5e, 0x28, 0xb4, 0x52, 0x4b, 0x83, 0xac,
	0x34, 0x60, 0x21, 0xb8, 0x74, 0x21, 0xe6, 0x42, 0xac, 0x09, 0x8d, 0x6f, 0x3d, 0x99, 0x81, 0x7e,
	0x55, 0x1b, 0x9e, 0x81, 0x91, 0x1d, 0x2f, 0xf2, 0xdc, 0x48, 0x6c, 0xe8, 0xf1, 0x74, 0x97, 0x97,
	0x82, 0x0b, 0xad, 0xc1, 0x0a, 0xab, 0x40, 0x23, 0xaf, 0xa4, 0x41, 0x05, 0x5a, 0xe9, 0x8d, 0x8f,
	0x84, 0x1f, 0x84, 0x0e, 0x9e, 0xda, 0xa5, 0x41, 0x42, 0x2f, 0x5a, 0x83, 0x35, 0x5a, 0x61, 0x77,
	0x28, 0x71, 0x44, 0x26, 0xff, 0xa3, 0x61, 0x3c, 0x63, 0x07, 0x54, 0x58, 0x8b, 0x3e, 0xbb, 0x70,
	0x7a, 0x5e, 0xf4, 0x66, 0x89, 0xcb, 0xf9, 0xf7, 0xcf, 0xe7, 0xf5, 0x0d, 0xbd, 0xea, 0xc1, 0x71,
	0x1f, 0xc6, 0xf0, 0x8b, 0xd0, 0xb3, 0xfe, 0x57, 0x41, 0x40, 0x8f, 0xb4, 0xd8, 0xca, 0x11, 0x99,
	0x90, 0x68, 0x90, 0xba, 0x77, 0xf0, 0x48, 0x4f, 0x0b, 0xc8, 0x44, 0xb1, 0x6e, 0xce, 0x1c, 0xfd,
	0x9b, 0x90, 0x68, 0x18, 0x87, 0x8d, 0x9a, 0x2f, 0x84, 0xd5, 0x85, 0x74, 0x82, 0x0f, 0x3e, 0x99,
	0x9e, 0x38, 0xb0, 0x99, 0x96, 0x77, 0xb5, 0xd5, 0x9c, 0x86, 0x7f, 0x59, 0x79, 0x8f, 0xd5, 0x3d,
	0x9d, 0x2a, 0xf0, 0x0b, 0x4a, 0x03, 0xef, 0xfb, 0x43, 0x35, 0xac, 0x3a, 0x79, 0x4c, 0xea, 0x56,
	0x13, 0xf2, 0x72, 0xec, 0xea, 0x5d, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xba, 0x47, 0x77, 0x45,
	0xe6, 0x01, 0x00, 0x00,
}
