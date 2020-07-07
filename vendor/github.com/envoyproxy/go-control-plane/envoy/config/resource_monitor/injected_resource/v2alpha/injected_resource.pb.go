// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/resource_monitor/injected_resource/v2alpha/injected_resource.proto

package envoy_config_resource_monitor_injected_resource_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type InjectedResourceConfig struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjectedResourceConfig) Reset()         { *m = InjectedResourceConfig{} }
func (m *InjectedResourceConfig) String() string { return proto.CompactTextString(m) }
func (*InjectedResourceConfig) ProtoMessage()    {}
func (*InjectedResourceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_de2fb4e1cfb2f415, []int{0}
}

func (m *InjectedResourceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjectedResourceConfig.Unmarshal(m, b)
}
func (m *InjectedResourceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjectedResourceConfig.Marshal(b, m, deterministic)
}
func (m *InjectedResourceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjectedResourceConfig.Merge(m, src)
}
func (m *InjectedResourceConfig) XXX_Size() int {
	return xxx_messageInfo_InjectedResourceConfig.Size(m)
}
func (m *InjectedResourceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InjectedResourceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InjectedResourceConfig proto.InternalMessageInfo

func (m *InjectedResourceConfig) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func init() {
	proto.RegisterType((*InjectedResourceConfig)(nil), "envoy.config.resource_monitor.injected_resource.v2alpha.InjectedResourceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/resource_monitor/injected_resource/v2alpha/injected_resource.proto", fileDescriptor_de2fb4e1cfb2f415)
}

var fileDescriptor_de2fb4e1cfb2f415 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0x8d, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xd2, 0xcf, 0xcc, 0xcb, 0x4a, 0x4d, 0x2e,
	0x49, 0x4d, 0x89, 0x87, 0xc9, 0xe8, 0x97, 0x19, 0x25, 0xe6, 0x14, 0x64, 0x24, 0x62, 0xca, 0xe8,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x99, 0x83, 0x0d, 0xd4, 0x83, 0x18, 0xa8, 0x87, 0x6e, 0xa0,
	0x1e, 0xa6, 0x36, 0xa8, 0x81, 0x52, 0xb2, 0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9,
	0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x10, 0x73,
	0xa5, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x84, 0x92,
	0x2d, 0x97, 0x98, 0x27, 0xd4, 0xd0, 0x20, 0xa8, 0x99, 0xce, 0x60, 0xcb, 0x85, 0x94, 0xb9, 0x38,
	0xd2, 0x32, 0x73, 0x52, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xd8,
	0x7f, 0x39, 0xb1, 0x14, 0x31, 0x29, 0x30, 0x06, 0xc1, 0x25, 0x9c, 0x32, 0x76, 0x35, 0x9c, 0xb8,
	0xc8, 0xc6, 0x24, 0xc0, 0xc8, 0xe5, 0x9a, 0x99, 0xaf, 0x07, 0x76, 0x7c, 0x41, 0x51, 0x7e, 0x45,
	0xa5, 0x1e, 0x99, 0xfe, 0x70, 0x12, 0x45, 0x77, 0x4d, 0x00, 0xc8, 0x99, 0x01, 0x8c, 0x49, 0x6c,
	0x60, 0xf7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0x29, 0x27, 0xc8, 0x73, 0x01, 0x00,
	0x00,
}
