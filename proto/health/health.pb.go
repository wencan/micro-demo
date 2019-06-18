// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

package micro_health_v1

import (
	fmt "fmt"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
	3: "SERVICE_UNKNOWN",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":         0,
	"SERVING":         1,
	"NOT_SERVING":     2,
	"SERVICE_UNKNOWN": 3,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=micro.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbebe66dda7cb29, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("micro.health.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "micro.health.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "micro.health.v1.HealthCheckResponse")
}

func init() { proto.RegisterFile("health.proto", fileDescriptor_fdbebe66dda7cb29) }

var fileDescriptor_fdbebe66dda7cb29 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7,
	0x83, 0x8a, 0x95, 0x19, 0x2a, 0xe9, 0x71, 0x09, 0x79, 0x80, 0x39, 0xce, 0x19, 0xa9, 0xc9, 0xd9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0xd2, 0x26, 0x46, 0x2e, 0x61,
	0x14, 0x0d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x5e, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25,
	0xa5, 0xc5, 0x60, 0x0d, 0x7c, 0x46, 0x46, 0x7a, 0x68, 0x36, 0xe9, 0x61, 0xd1, 0xa5, 0x17, 0x0c,
	0x32, 0x35, 0x2f, 0x3d, 0x18, 0xac, 0x33, 0x08, 0x6a, 0x82, 0x92, 0x3f, 0x17, 0x2f, 0x8a, 0x84,
	0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x03, 0x88, 0x13, 0xec,
	0x1a, 0x14, 0xe6, 0xe9, 0xe7, 0x2e, 0xc0, 0x28, 0xc4, 0xcf, 0xc5, 0xed, 0xe7, 0x1f, 0x12, 0x0f,
	0x13, 0x60, 0x12, 0x12, 0xe6, 0xe2, 0x07, 0x73, 0x9c, 0x5d, 0xe3, 0x61, 0x5a, 0x98, 0x8d, 0x36,
	0x31, 0x72, 0xb1, 0x41, 0xac, 0x17, 0x0a, 0xe2, 0x62, 0x05, 0x3b, 0x41, 0x48, 0x19, 0xbf, 0x03,
	0xc1, 0xe1, 0x20, 0xa5, 0x42, 0x8c, 0x2f, 0x84, 0x42, 0xb8, 0x58, 0xc3, 0x13, 0x4b, 0x92, 0x33,
	0xa8, 0x68, 0xa6, 0x01, 0xa3, 0x93, 0x60, 0x14, 0x24, 0xb2, 0xe2, 0x21, 0x0a, 0xe3, 0xcb, 0x0c,
	0x93, 0xd8, 0xc0, 0x91, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xee, 0x19, 0xc6, 0x82, 0xd4,
	0x01, 0x00, 0x00,
}