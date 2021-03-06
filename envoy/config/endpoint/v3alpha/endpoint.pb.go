// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/endpoint/v3alpha/endpoint.proto

package envoy_config_endpoint_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ClusterLoadAssignment struct {
	ClusterName          string                        `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*LocalityLbEndpoints        `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*Endpoint          `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_4667cf48650033d2, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads           []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor  *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter      *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	DisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                     `json:"-"`
	XXX_unrecognized        []byte                                       `json:"-"`
	XXX_sizecache           int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_4667cf48650033d2, []int{0, 0}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                     `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *v3alpha.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_4667cf48650033d2, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.config.endpoint.v3alpha.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*Endpoint)(nil), "envoy.config.endpoint.v3alpha.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.config.endpoint.v3alpha.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.config.endpoint.v3alpha.ClusterLoadAssignment.Policy.DropOverload")
}

func init() {
	proto.RegisterFile("envoy/config/endpoint/v3alpha/endpoint.proto", fileDescriptor_4667cf48650033d2)
}

var fileDescriptor_4667cf48650033d2 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdf, 0x4e, 0xd4, 0x4e,
	0x18, 0xa5, 0x65, 0x97, 0xdf, 0x32, 0xfc, 0xcd, 0xfc, 0x54, 0xea, 0x06, 0x49, 0x45, 0x8d, 0x0b,
	0x98, 0xd6, 0x2c, 0x89, 0xe2, 0x1a, 0x4c, 0x58, 0x81, 0xa0, 0x21, 0xd0, 0x94, 0xe0, 0xa5, 0xcd,
	0x6c, 0x3b, 0x5b, 0x47, 0xbb, 0x33, 0xe3, 0x74, 0xb6, 0xda, 0x37, 0xf0, 0x19, 0x7c, 0x04, 0x1f,
	0xc3, 0xc4, 0xe7, 0xf0, 0xd6, 0x67, 0xe0, 0xca, 0xb4, 0x9d, 0x16, 0x5c, 0x11, 0x88, 0xde, 0xb5,
	0x73, 0xbe, 0x73, 0xce, 0x7c, 0xdf, 0x77, 0x06, 0x3c, 0xc0, 0x34, 0x61, 0xa9, 0xed, 0x33, 0xda,
	0x27, 0xa1, 0x8d, 0x69, 0xc0, 0x19, 0xa1, 0xd2, 0x4e, 0xd6, 0x51, 0xc4, 0xdf, 0xa0, 0xea, 0xc0,
	0xe2, 0x82, 0x49, 0x06, 0x6f, 0xe5, 0xd5, 0x56, 0x51, 0x6d, 0x55, 0xa0, 0xaa, 0x6e, 0x3e, 0xbe,
	0x9a, 0x98, 0xe7, 0xb3, 0x01, 0x67, 0x14, 0x53, 0x19, 0x17, 0xba, 0x4d, 0xb3, 0x20, 0xca, 0x94,
	0xe3, 0xaa, 0x9a, 0x63, 0xe1, 0xe3, 0xd2, 0xb9, 0xb9, 0x18, 0x32, 0x16, 0x46, 0xd8, 0x46, 0x9c,
	0xd8, 0x88, 0x52, 0x26, 0x91, 0x24, 0x8c, 0x96, 0xfc, 0x25, 0x85, 0xe6, 0x7f, 0xbd, 0x61, 0xdf,
	0x0e, 0x86, 0x22, 0x2f, 0xf8, 0x13, 0xfe, 0x41, 0x20, 0xce, 0xb1, 0x28, 0xf9, 0xb7, 0x87, 0x01,
	0x47, 0x67, 0x75, 0xed, 0x04, 0x8b, 0x98, 0x30, 0x4a, 0x68, 0xa8, 0x4a, 0x16, 0x12, 0x14, 0x91,
	0x00, 0x49, 0x6c, 0x97, 0x1f, 0x05, 0xb0, 0xfc, 0xbd, 0x01, 0xae, 0x3f, 0x8f, 0x86, 0xb1, 0xc4,
	0x62, 0x9f, 0xa1, 0x60, 0x2b, 0x8e, 0x49, 0x48, 0x07, 0x98, 0x4a, 0xb8, 0x0a, 0xa6, 0xfd, 0x02,
	0xf0, 0x28, 0x1a, 0x60, 0x43, 0x33, 0xb5, 0xd6, 0x64, 0xf7, 0xbf, 0x93, 0x6e, 0x4d, 0xe8, 0xa6,
	0xe6, 0x4e, 0x29, 0xf0, 0x00, 0x0d, 0x30, 0x74, 0xc0, 0x64, 0x39, 0x9e, 0xd8, 0xd0, 0xcd, 0xf1,
	0xd6, 0x54, 0xbb, 0x6d, 0x5d, 0x38, 0x6d, 0x6b, 0x9f, 0xf9, 0x28, 0x22, 0x32, 0xdd, 0xef, 0xed,
	0x94, 0x4c, 0xf7, 0x54, 0x04, 0xbe, 0x07, 0x73, 0x99, 0x6b, 0xe0, 0x9d, 0xea, 0xd6, 0x73, 0xdd,
	0xbd, 0x4b, 0x74, 0xcf, 0x6d, 0xc6, 0xca, 0x6e, 0x19, 0x54, 0x46, 0x3b, 0x54, 0x8a, 0xd4, 0x9d,
	0xa5, 0xbf, 0x1c, 0xc2, 0x23, 0x30, 0xc1, 0x59, 0x44, 0xfc, 0xd4, 0xa8, 0x99, 0x5a, 0x6b, 0xaa,
	0xfd, 0xf4, 0xaf, 0x9c, 0x9c, 0x5c, 0xc2, 0x55, 0x52, 0xcd, 0x1f, 0x35, 0x30, 0x51, 0x1c, 0x41,
	0x06, 0x66, 0x03, 0xc1, 0xb8, 0xc7, 0x12, 0x2c, 0x22, 0x86, 0x82, 0x72, 0x52, 0x7b, 0xff, 0xe0,
	0x63, 0x6d, 0x0b, 0xc6, 0x0f, 0x95, 0xa0, 0x3b, 0x13, 0x9c, 0xf9, 0x8b, 0xe1, 0x6b, 0xb0, 0x90,
	0x79, 0x71, 0xc1, 0x12, 0xa2, 0xe2, 0xe0, 0xf5, 0x91, 0x2f, 0x99, 0x30, 0xc6, 0xf3, 0x0e, 0x17,
	0xad, 0x22, 0x59, 0x56, 0x99, 0x2c, 0xeb, 0xf8, 0x05, 0x95, 0xeb, 0xed, 0x57, 0x28, 0x1a, 0xe2,
	0x7c, 0xd5, 0xab, 0xba, 0x39, 0xe6, 0xde, 0x18, 0x55, 0xd9, 0xcd, 0x45, 0xe0, 0x31, 0xb8, 0x56,
	0x3d, 0x8a, 0x58, 0xa2, 0x08, 0x7b, 0xa8, 0x2f, 0xb1, 0x50, 0xe3, 0xbb, 0xf9, 0x9b, 0xf8, 0xb6,
	0x8a, 0x75, 0xb7, 0x71, 0xd2, 0xad, 0x7f, 0xd1, 0xf4, 0xd5, 0x31, 0x17, 0x96, 0x02, 0x47, 0x19,
	0x7f, 0x2b, 0xa3, 0xc3, 0x27, 0xc0, 0x08, 0x48, 0x8c, 0x7a, 0x11, 0xf6, 0x46, 0x8d, 0x8d, 0xba,
	0xa9, 0xb5, 0x1a, 0xee, 0x82, 0xc2, 0x0f, 0x47, 0xe0, 0xe6, 0x57, 0x0d, 0x4c, 0x9f, 0x9d, 0x08,
	0xbc, 0x03, 0x1a, 0x3e, 0x92, 0x38, 0x64, 0x22, 0x1d, 0x0d, 0x70, 0x05, 0xc0, 0x03, 0x30, 0x97,
	0x2f, 0x46, 0xbd, 0x59, 0x14, 0x62, 0x43, 0xcf, 0x5b, 0xb8, 0xa7, 0x36, 0x93, 0xbd, 0xec, 0x6a,
	0x1d, 0xbb, 0x02, 0xf9, 0x59, 0x17, 0x28, 0x72, 0x8a, 0x7a, 0x37, 0x5f, 0xab, 0x53, 0x91, 0x3b,
	0x9b, 0x9f, 0xbf, 0x7d, 0x5a, 0xda, 0x00, 0x8f, 0x0a, 0x32, 0xe2, 0xc4, 0x4a, 0xda, 0x57, 0xdf,
	0x62, 0xe7, 0x61, 0x46, 0x5f, 0x03, 0x2b, 0x57, 0xa6, 0xbf, 0xac, 0x35, 0xb4, 0x79, 0xbd, 0xf9,
	0x16, 0xfc, 0x7f, 0x4e, 0xcc, 0xe1, 0x3c, 0x18, 0x7f, 0x87, 0x55, 0xf7, 0x6e, 0xf6, 0x09, 0x37,
	0x41, 0x3d, 0xc9, 0x36, 0xac, 0xba, 0xbc, 0x7f, 0x49, 0xfe, 0x4a, 0x3d, 0xb7, 0x60, 0x75, 0xf4,
	0x0d, 0xad, 0xb3, 0x92, 0xdd, 0xf1, 0x2e, 0x58, 0xbe, 0xfc, 0x8e, 0xdd, 0x67, 0x60, 0x8d, 0xb0,
	0xc2, 0x82, 0x0b, 0xf6, 0x31, 0xbd, 0xd8, 0xad, 0x3b, 0x53, 0xda, 0x39, 0x59, 0x6c, 0x1c, 0xad,
	0x37, 0x91, 0xe7, 0x67, 0xfd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x39, 0x64, 0x63, 0xec,
	0x05, 0x00, 0x00,
}
