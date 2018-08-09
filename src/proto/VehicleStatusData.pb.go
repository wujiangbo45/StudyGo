// Code generated by protoc-gen-go. DO NOT EDIT.
// source: VehicleStatusData.proto

package proto_dic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VehicleStatusData struct {
	Types                *StatusType `protobuf:"varint,1,req,name=types,enum=proto_dic.StatusType" json:"types,omitempty"`
	StatusValue          *int64      `protobuf:"varint,2,req,name=statusValue" json:"statusValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VehicleStatusData) Reset()         { *m = VehicleStatusData{} }
func (m *VehicleStatusData) String() string { return proto.CompactTextString(m) }
func (*VehicleStatusData) ProtoMessage()    {}
func (*VehicleStatusData) Descriptor() ([]byte, []int) {
	return fileDescriptor_VehicleStatusData_4ff1f08681dbb66b, []int{0}
}
func (m *VehicleStatusData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VehicleStatusData.Unmarshal(m, b)
}
func (m *VehicleStatusData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VehicleStatusData.Marshal(b, m, deterministic)
}
func (dst *VehicleStatusData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VehicleStatusData.Merge(dst, src)
}
func (m *VehicleStatusData) XXX_Size() int {
	return xxx_messageInfo_VehicleStatusData.Size(m)
}
func (m *VehicleStatusData) XXX_DiscardUnknown() {
	xxx_messageInfo_VehicleStatusData.DiscardUnknown(m)
}

var xxx_messageInfo_VehicleStatusData proto.InternalMessageInfo

func (m *VehicleStatusData) GetTypes() StatusType {
	if m != nil && m.Types != nil {
		return *m.Types
	}
	return StatusType_dpfPressure
}

func (m *VehicleStatusData) GetStatusValue() int64 {
	if m != nil && m.StatusValue != nil {
		return *m.StatusValue
	}
	return 0
}

func init() {
	proto.RegisterType((*VehicleStatusData)(nil), "proto_dic.VehicleStatusData")
}

func init() {
	proto.RegisterFile("VehicleStatusData.proto", fileDescriptor_VehicleStatusData_4ff1f08681dbb66b)
}

var fileDescriptor_VehicleStatusData_4ff1f08681dbb66b = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0f, 0x4b, 0xcd, 0xc8,
	0x4c, 0xce, 0x49, 0x0d, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x76, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0xf1, 0x29, 0x99, 0xc9, 0x52, 0x02, 0x10, 0xc9, 0x90,
	0xca, 0x82, 0x54, 0x88, 0xa4, 0x52, 0x12, 0x97, 0x20, 0x86, 0x3e, 0x21, 0x6d, 0x2e, 0xd6, 0x92,
	0xca, 0x82, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x3e, 0x23, 0x51, 0x3d, 0xb8, 0x09, 0x7a,
	0x08, 0x03, 0x82, 0x20, 0x6a, 0x84, 0x14, 0xb8, 0xb8, 0x8b, 0xc1, 0x82, 0x61, 0x89, 0x39, 0xa5,
	0xa9, 0x12, 0x4c, 0x0a, 0x4c, 0x1a, 0xcc, 0x41, 0xc8, 0x42, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x8d, 0xf4, 0x11, 0x80, 0x9a, 0x00, 0x00, 0x00,
}