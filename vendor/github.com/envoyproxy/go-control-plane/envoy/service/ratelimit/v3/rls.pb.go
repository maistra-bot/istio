// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v3/rls.proto

package envoy_service_ratelimit_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{1, 0, 0}
}

type RateLimitRequest struct {
	Domain               string                    `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*v3.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	HitsAddend           uint32                    `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*v3.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

type RateLimitResponse struct {
	OverallCode          RateLimitResponse_Code                `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v3.RateLimitResponse_Code" json:"overall_code,omitempty"`
	Statuses             []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	ResponseHeadersToAdd []*v31.HeaderValue                    `protobuf:"bytes,3,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	RequestHeadersToAdd  []*v31.HeaderValue                    `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetResponseHeadersToAdd() []*v31.HeaderValue {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RateLimitResponse) GetRequestHeadersToAdd() []*v31.HeaderValue {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

type RateLimitResponse_RateLimit struct {
	RequestsPerUnit      uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v3.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	Code                 RateLimitResponse_Code       `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v3.RateLimitResponse_Code" json:"code,omitempty"`
	CurrentLimit         *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	LimitRemaining       uint32                       `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_42d8f587b632794e, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v3.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v3.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v3.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v3.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v3.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v3.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v3/rls.proto", fileDescriptor_42d8f587b632794e)
}

var fileDescriptor_42d8f587b632794e = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x4f, 0xdb, 0x48,
	0x14, 0xc7, 0xb1, 0x13, 0x85, 0xf0, 0xc2, 0x0f, 0x33, 0xbb, 0x82, 0x28, 0xd2, 0x2e, 0x10, 0xad,
	0xb4, 0xd1, 0x02, 0xb6, 0x94, 0xac, 0xf6, 0x47, 0xb4, 0xb0, 0x0a, 0x84, 0x0a, 0x04, 0x24, 0x91,
	0x43, 0xe8, 0x4f, 0xc9, 0x32, 0xf1, 0x14, 0x2c, 0x39, 0x33, 0xe9, 0xcc, 0x24, 0x82, 0x5b, 0x0f,
	0x3d, 0x70, 0xeb, 0xbd, 0xff, 0x40, 0xff, 0x87, 0xde, 0x2b, 0xf5, 0xda, 0x7b, 0xff, 0x98, 0x6a,
	0x66, 0x9c, 0x10, 0x82, 0xa8, 0x48, 0x7b, 0xb3, 0xdf, 0xbc, 0xef, 0xe7, 0xf9, 0x7d, 0xdf, 0xf3,
	0xc0, 0x6f, 0x98, 0xf4, 0xe9, 0x95, 0xc3, 0x31, 0xeb, 0x87, 0x6d, 0xec, 0x30, 0x5f, 0xe0, 0x28,
	0xec, 0x84, 0xc2, 0xe9, 0x97, 0x1c, 0x16, 0x71, 0xbb, 0xcb, 0xa8, 0xa0, 0x28, 0xa7, 0xb2, 0xec,
	0x38, 0xcb, 0x1e, 0x66, 0xd9, 0xfd, 0x52, 0x6e, 0x45, 0x13, 0xda, 0x94, 0xbc, 0x0c, 0xcf, 0x9d,
	0x36, 0x65, 0x58, 0x6a, 0xcf, 0x7c, 0x8e, 0xb5, 0x38, 0xf7, 0xa7, 0x4e, 0xc0, 0x97, 0x02, 0x13,
	0x1e, 0x52, 0xc2, 0x9d, 0x36, 0xed, 0x74, 0x28, 0x19, 0x2b, 0x36, 0x64, 0x6a, 0xd5, 0x2f, 0xbd,
	0xa0, 0xeb, 0x3b, 0x3e, 0x21, 0x54, 0xf8, 0x42, 0xa9, 0xb8, 0xf0, 0x45, 0x2f, 0xfe, 0xa2, 0xdc,
	0xda, 0x9d, 0xe3, 0x3e, 0x66, 0x92, 0x1e, 0x92, 0xf3, 0x38, 0x65, 0xb9, 0xef, 0x47, 0x61, 0xe0,
	0x0b, 0xec, 0x0c, 0x1e, 0xf4, 0x41, 0xfe, 0x8b, 0x01, 0x96, 0xeb, 0x0b, 0x7c, 0x24, 0xcb, 0xb9,
	0xf8, 0x55, 0x0f, 0x73, 0x81, 0x96, 0x20, 0x15, 0xd0, 0x8e, 0x1f, 0x92, 0xac, 0xb1, 0x6a, 0x14,
	0x66, 0xdc, 0xf8, 0x0d, 0x3d, 0x87, 0x4c, 0x80, 0x79, 0x9b, 0x85, 0x5d, 0x41, 0x19, 0xcf, 0x9a,
	0xab, 0x89, 0x42, 0xa6, 0xf8, 0xaf, 0xad, 0x0d, 0xb9, 0xe9, 0xc9, 0xd6, 0x3d, 0xdd, 0xb2, 0xc6,
	0x1e, 0x16, 0xa9, 0x0e, 0x09, 0xee, 0x28, 0x0d, 0xad, 0x40, 0xe6, 0x22, 0x14, 0xdc, 0xf3, 0x83,
	0x00, 0x93, 0x20, 0x9b, 0x58, 0x35, 0x0a, 0x73, 0x2e, 0xc8, 0x50, 0x45, 0x45, 0xca, 0xc5, 0x77,
	0x1f, 0xaf, 0x7f, 0xdd, 0x84, 0xf5, 0x7b, 0xfd, 0x2f, 0xda, 0xe3, 0x9d, 0xe4, 0xdf, 0xa6, 0x61,
	0x71, 0x24, 0xc8, 0xbb, 0x94, 0x70, 0x8c, 0x5a, 0x30, 0x4b, 0xfb, 0x98, 0xf9, 0x51, 0xe4, 0xb5,
	0x69, 0x80, 0x55, 0x97, 0xf3, 0xc5, 0xa2, 0x7d, 0xff, 0x64, 0xed, 0x3b, 0x10, 0x7b, 0x97, 0x06,
	0xd8, 0xcd, 0xc4, 0x1c, 0xf9, 0x82, 0x9e, 0x41, 0x5a, 0xcf, 0x05, 0x0f, 0xbc, 0xd9, 0x9e, 0x0c,
	0x79, 0x63, 0x4d, 0x53, 0x71, 0xdc, 0x21, 0x0f, 0x3d, 0x81, 0x65, 0x16, 0xa7, 0x79, 0x17, 0xd8,
	0x0f, 0x30, 0xe3, 0x9e, 0xa0, 0xd2, 0xac, 0x6c, 0x42, 0x95, 0x5a, 0x8b, 0x4b, 0xe9, 0xdd, 0xb3,
	0xe5, 0xee, 0xc9, 0x22, 0xfb, 0x2a, 0xf7, 0xd4, 0x8f, 0x7a, 0xd8, 0xfd, 0x79, 0x40, 0xd0, 0x41,
	0x7e, 0x42, 0x2b, 0x41, 0x80, 0x4e, 0x61, 0x89, 0x69, 0xb7, 0xc6, 0xc1, 0xc9, 0x87, 0x82, 0x7f,
	0x8a, 0x01, 0xa3, 0xdc, 0xdc, 0xb5, 0x09, 0x33, 0xc3, 0x16, 0xd1, 0x1f, 0xb0, 0x18, 0x27, 0x71,
	0xaf, 0x8b, 0x99, 0xd7, 0x23, 0xa1, 0x50, 0xbe, 0xcf, 0xb9, 0x0b, 0x83, 0x83, 0x06, 0x66, 0x2d,
	0x12, 0x0a, 0xd4, 0x80, 0xa4, 0x3a, 0x36, 0xd5, 0x58, 0xfe, 0x9b, 0xcc, 0xc3, 0x61, 0xc4, 0x96,
	0x2c, 0x57, 0x91, 0xf2, 0xdb, 0x90, 0x54, 0xe4, 0x0c, 0x4c, 0xb7, 0x6a, 0x87, 0xb5, 0xfa, 0xe3,
	0x9a, 0x35, 0x85, 0x00, 0x52, 0xcd, 0xbd, 0xdd, 0x7a, 0xad, 0x6a, 0x19, 0xf2, 0xf9, 0xf8, 0xa0,
	0xd6, 0x3a, 0xd9, 0xb3, 0x4c, 0x94, 0x86, 0xe4, 0x7e, 0xbd, 0xe5, 0x5a, 0x09, 0x34, 0x0d, 0x89,
	0x6a, 0xe5, 0xa9, 0x95, 0x2c, 0x6f, 0xc9, 0xd5, 0xfb, 0x07, 0xfe, 0x7a, 0xd8, 0xea, 0x8d, 0x7f,
	0x49, 0xee, 0xbd, 0x09, 0xd6, 0xf8, 0x6c, 0xd1, 0x23, 0x48, 0xfe, 0xe0, 0xf2, 0x29, 0x3d, 0x7a,
	0x01, 0x73, 0xed, 0x1e, 0x63, 0x98, 0x08, 0x4f, 0x09, 0x94, 0x6d, 0x99, 0xe2, 0xdf, 0xdf, 0x69,
	0x9b, 0x3b, 0x1b, 0xd3, 0xf4, 0xdc, 0x7e, 0x87, 0x05, 0xa5, 0xf2, 0x18, 0x96, 0x57, 0x40, 0x48,
	0xce, 0xe3, 0x3f, 0x73, 0x3e, 0xd2, 0xfa, 0x38, 0x5a, 0xae, 0x4a, 0x8b, 0xfe, 0x87, 0xad, 0x89,
	0x2c, 0x1a, 0x37, 0x25, 0xbf, 0x0e, 0x49, 0xf5, 0x2b, 0xdd, 0x1a, 0x54, 0x0a, 0xcc, 0xfa, 0xa1,
	0x65, 0xa0, 0x79, 0x80, 0xfa, 0xe9, 0x9e, 0xeb, 0x1d, 0x1d, 0x1c, 0x1f, 0x9c, 0x58, 0x66, 0xb9,
	0x24, 0x4b, 0xda, 0xb0, 0x31, 0x49, 0xc9, 0xe2, 0x9b, 0xd1, 0x0b, 0xaf, 0xa9, 0x35, 0xa8, 0x0b,
	0x0b, 0xcd, 0x0b, 0xda, 0x8b, 0x82, 0x9b, 0x85, 0xdd, 0x78, 0xa0, 0x7f, 0x6a, 0x75, 0x73, 0x9b,
	0x13, 0xb9, 0x9d, 0x9f, 0xda, 0xa9, 0x7c, 0x78, 0xfd, 0xe9, 0x73, 0xca, 0xb4, 0x4c, 0x28, 0x84,
	0x54, 0x8b, 0xbb, 0x8c, 0x5e, 0x5e, 0x7d, 0x83, 0xb3, 0x93, 0x76, 0x23, 0xde, 0x90, 0xb7, 0x76,
	0xc3, 0xb8, 0x36, 0x8c, 0xb3, 0x94, 0xba, 0xc1, 0x4b, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xad,
	0x1f, 0x76, 0xed, 0xb7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v3.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(ctx context.Context, req *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v3.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v3.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v3/rls.proto",
}
