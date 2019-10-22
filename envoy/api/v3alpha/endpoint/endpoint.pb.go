// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/endpoint/endpoint.proto

package envoy_api_v3alpha_endpoint

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Endpoint struct {
	Address              *core.Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

type Endpoint_HealthCheckConfig struct {
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{0, 0}
}

func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(m, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

type LbEndpoint struct {
	// Types that are valid to be assigned to HostIdentifier:
	//	*LbEndpoint_Endpoint
	//	*LbEndpoint_EndpointName
	HostIdentifier       isLbEndpoint_HostIdentifier `protobuf_oneof:"host_identifier"`
	HealthStatus         core.HealthStatus           `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v3alpha.core.HealthStatus" json:"health_status,omitempty"`
	Metadata             *core.Metadata              `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value       `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{1}
}

func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

type isLbEndpoint_HostIdentifier interface {
	isLbEndpoint_HostIdentifier()
}

type LbEndpoint_Endpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3,oneof"`
}

type LbEndpoint_EndpointName struct {
	EndpointName string `protobuf:"bytes,5,opt,name=endpoint_name,json=endpointName,proto3,oneof"`
}

func (*LbEndpoint_Endpoint) isLbEndpoint_HostIdentifier() {}

func (*LbEndpoint_EndpointName) isLbEndpoint_HostIdentifier() {}

func (m *LbEndpoint) GetHostIdentifier() isLbEndpoint_HostIdentifier {
	if m != nil {
		return m.HostIdentifier
	}
	return nil
}

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_Endpoint); ok {
		return x.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetEndpointName() string {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_EndpointName); ok {
		return x.EndpointName
	}
	return ""
}

func (m *LbEndpoint) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LbEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LbEndpoint_Endpoint)(nil),
		(*LbEndpoint_EndpointName)(nil),
	}
}

type LocalityLbEndpoints struct {
	Locality             *core.Locality        `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	LbEndpoints          []*LbEndpoint         `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	Priority             uint32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Proximity            *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=proximity,proto3" json:"proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{2}
}

func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LocalityLbEndpoints) GetProximity() *wrappers.UInt32Value {
	if m != nil {
		return m.Proximity
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v3alpha.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v3alpha.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v3alpha.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v3alpha.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/endpoint/endpoint.proto", fileDescriptor_9e420efb9b41886f)
}

var fileDescriptor_9e420efb9b41886f = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xdf, 0x8a, 0xd3, 0x4c,
	0x1c, 0xdd, 0xb4, 0xdb, 0x6e, 0x3b, 0x6d, 0xbf, 0x8f, 0xa6, 0x88, 0xa1, 0x88, 0xd6, 0x52, 0x25,
	0xdb, 0x8b, 0x09, 0xb4, 0x20, 0xf8, 0xef, 0xc2, 0x2c, 0x42, 0x0b, 0xab, 0x2c, 0x23, 0x2a, 0x5e,
	0x85, 0x49, 0x32, 0x6d, 0x06, 0xa7, 0x99, 0x90, 0x4c, 0xbb, 0xf6, 0xce, 0xf7, 0xf4, 0x15, 0xc4,
	0x07, 0xe8, 0xcd, 0x4a, 0x26, 0x99, 0xb4, 0x58, 0xb3, 0x0a, 0xde, 0x4d, 0xe6, 0x77, 0xce, 0xf9,
	0x9d, 0x73, 0x32, 0xe0, 0x9c, 0x84, 0x1b, 0xbe, 0xb5, 0x70, 0x44, 0xad, 0xcd, 0x14, 0xb3, 0x28,
	0xc0, 0x16, 0x09, 0xfd, 0x88, 0xd3, 0x50, 0x14, 0x07, 0x18, 0xc5, 0x5c, 0x70, 0xbd, 0x2f, 0xa1,
	0x10, 0x47, 0x14, 0xe6, 0x50, 0xa8, 0x10, 0xfd, 0xd1, 0xb1, 0x8c, 0xc7, 0x63, 0x62, 0x61, 0xdf,
	0x8f, 0x49, 0x92, 0x64, 0x0a, 0xfd, 0x87, 0x25, 0x28, 0x17, 0x27, 0x24, 0x87, 0x9c, 0x97, 0x40,
	0x02, 0x82, 0x99, 0x08, 0x1c, 0x2f, 0x20, 0xde, 0xe7, 0x1c, 0x7a, 0x7f, 0xc9, 0xf9, 0x92, 0x11,
	0x4b, 0x7e, 0xb9, 0xeb, 0x85, 0x75, 0x1d, 0xe3, 0x28, 0x22, 0xb1, 0xda, 0x76, 0x77, 0x83, 0x19,
	0xf5, 0xb1, 0x20, 0x96, 0x3a, 0x64, 0x83, 0xe1, 0x77, 0x0d, 0x34, 0x5e, 0xe7, 0xce, 0xf5, 0xa7,
	0xe0, 0x2c, 0x37, 0x69, 0x68, 0x03, 0xcd, 0x6c, 0x4d, 0x1e, 0xc0, 0xe3, 0x9c, 0xa9, 0x05, 0xf8,
	0x2a, 0x83, 0x21, 0x85, 0xd7, 0x17, 0xa0, 0x77, 0x68, 0xcb, 0xf1, 0x78, 0xb8, 0xa0, 0x4b, 0xa3,
	0x22, 0x65, 0x9e, 0xc0, 0xf2, 0xba, 0xa0, 0xda, 0x0e, 0x67, 0x92, 0x7f, 0x91, 0xd2, 0x2f, 0x24,
	0x1b, 0x75, 0x83, 0x5f, 0xaf, 0xfa, 0x2f, 0x41, 0xf7, 0x08, 0xa7, 0x9b, 0x00, 0x44, 0x3c, 0x16,
	0xce, 0x06, 0xb3, 0x35, 0x91, 0xd6, 0x3b, 0x76, 0x73, 0x67, 0xd7, 0xc7, 0xa7, 0xc6, 0xcd, 0x4d,
	0x15, 0x35, 0xd3, 0xe1, 0x87, 0x74, 0x36, 0xfc, 0x51, 0x01, 0xe0, 0xd2, 0x2d, 0x02, 0xdb, 0xa0,
	0xa1, 0x7c, 0xe4, 0x89, 0x47, 0x7f, 0x63, 0x75, 0x76, 0x82, 0x0a, 0x9e, 0xfe, 0x08, 0x74, 0xd4,
	0xd9, 0x09, 0xf1, 0x8a, 0x18, 0xb5, 0x81, 0x66, 0x36, 0x67, 0x27, 0xa8, 0xad, 0xae, 0xdf, 0xe2,
	0x15, 0xd1, 0xe7, 0xa0, 0x93, 0x17, 0x94, 0x08, 0x2c, 0xd6, 0x89, 0xac, 0xe6, 0xbf, 0xdf, 0xee,
	0x93, 0x0d, 0x67, 0x29, 0xdf, 0x49, 0x2c, 0x6a, 0x07, 0x07, 0x5f, 0xfa, 0x0b, 0xd0, 0x58, 0x11,
	0x81, 0x7d, 0x2c, 0xb0, 0x51, 0x95, 0xae, 0x07, 0x65, 0x2a, 0x6f, 0x72, 0x1c, 0x2a, 0x18, 0xfa,
	0x27, 0x70, 0x87, 0x71, 0xec, 0x3b, 0x2e, 0x66, 0x38, 0xf4, 0x68, 0xb8, 0x74, 0xae, 0x09, 0x5d,
	0x06, 0xc2, 0x38, 0x95, 0x52, 0xf7, 0x60, 0xf6, 0x94, 0xa0, 0x7a, 0x4a, 0xf0, 0xfd, 0x3c, 0x14,
	0xd3, 0x89, 0xec, 0xcf, 0x3e, 0xdb, 0xd9, 0xa7, 0xe3, 0x8a, 0xa9, 0xa1, 0x5e, 0xaa, 0x61, 0x2b,
	0x89, 0x8f, 0x52, 0xc1, 0xee, 0x82, 0xff, 0x03, 0x9e, 0x08, 0x87, 0xfa, 0x24, 0x14, 0x74, 0x41,
	0x49, 0x3c, 0xfc, 0x56, 0x01, 0xbd, 0x4b, 0xee, 0x61, 0x46, 0xc5, 0x76, 0x5f, 0xbc, 0xcc, 0xc0,
	0xf2, 0xeb, 0xbc, 0xf9, 0xd2, 0x0c, 0x8a, 0x8e, 0x0a, 0x86, 0x3e, 0x07, 0x6d, 0xe6, 0x3a, 0xaa,
	0xdf, 0xb4, 0xcb, 0xaa, 0xd9, 0x9a, 0x3c, 0xbe, 0xed, 0xdf, 0xed, 0x97, 0xa3, 0x16, 0x3b, 0x30,
	0x52, 0x5a, 0x47, 0xf5, 0x5f, 0xeb, 0xd0, 0x47, 0xa0, 0x11, 0xc5, 0x94, 0xc7, 0x69, 0xc6, 0x9a,
	0x7c, 0x94, 0x8d, 0x9d, 0x5d, 0x1b, 0x57, 0x8d, 0xaf, 0x1a, 0x2a, 0x26, 0xfa, 0x33, 0xd0, 0x8c,
	0x62, 0xfe, 0x85, 0xae, 0x52, 0x58, 0xfd, 0xcf, 0x4b, 0xd1, 0x1e, 0x6e, 0x3f, 0x07, 0x26, 0xe5,
	0x59, 0xea, 0xf4, 0x72, 0x7b, 0x4b, 0x01, 0x76, 0x47, 0x65, 0xbe, 0x4a, 0x45, 0xaf, 0x34, 0xb7,
	0x2e, 0xd5, 0xa7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x04, 0x8a, 0xc5, 0x47, 0xf7, 0x04, 0x00,
	0x00,
}