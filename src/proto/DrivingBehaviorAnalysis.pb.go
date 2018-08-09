// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DrivingBehaviorAnalysis.proto

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

type DrivingBehaviorAnalysis struct {
	TurningAngle         *int32   `protobuf:"varint,1,opt,name=turningAngle" json:"turningAngle,omitempty"`
	LowOilDrivingValue   *int32   `protobuf:"varint,2,opt,name=lowOilDrivingValue" json:"lowOilDrivingValue,omitempty"`
	VehicleSpeedFromEcu  *int32   `protobuf:"varint,3,opt,name=vehicleSpeedFromEcu" json:"vehicleSpeedFromEcu,omitempty"`
	RpmWhenAlarming      *int32   `protobuf:"varint,4,opt,name=rpmWhenAlarming" json:"rpmWhenAlarming,omitempty"`
	VelocityChangeValue  *int32   `protobuf:"varint,5,opt,name=velocityChangeValue" json:"velocityChangeValue,omitempty"`
	CurrentGearshift     *int32   `protobuf:"varint,6,opt,name=currentGearshift" json:"currentGearshift,omitempty"`
	CurrentRPM           *int32   `protobuf:"varint,7,opt,name=currentRPM" json:"currentRPM,omitempty"`
	DeviceIdentity       *string  `protobuf:"bytes,8,opt,name=deviceIdentity" json:"deviceIdentity,omitempty"`
	BrakeTimes           *int32   `protobuf:"varint,9,opt,name=brakeTimes" json:"brakeTimes,omitempty"`
	ClutchTimes          *int32   `protobuf:"varint,10,opt,name=clutchTimes" json:"clutchTimes,omitempty"`
	RetarderTimes        *int32   `protobuf:"varint,11,opt,name=retarderTimes" json:"retarderTimes,omitempty"`
	AbsTimes             *int32   `protobuf:"varint,12,opt,name=absTimes" json:"absTimes,omitempty"`
	ReverseTimes         *int32   `protobuf:"varint,13,opt,name=reverseTimes" json:"reverseTimes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrivingBehaviorAnalysis) Reset()         { *m = DrivingBehaviorAnalysis{} }
func (m *DrivingBehaviorAnalysis) String() string { return proto.CompactTextString(m) }
func (*DrivingBehaviorAnalysis) ProtoMessage()    {}
func (*DrivingBehaviorAnalysis) Descriptor() ([]byte, []int) {
	return fileDescriptor_DrivingBehaviorAnalysis_c4e7ce55d1a3a8d9, []int{0}
}
func (m *DrivingBehaviorAnalysis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrivingBehaviorAnalysis.Unmarshal(m, b)
}
func (m *DrivingBehaviorAnalysis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrivingBehaviorAnalysis.Marshal(b, m, deterministic)
}
func (dst *DrivingBehaviorAnalysis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrivingBehaviorAnalysis.Merge(dst, src)
}
func (m *DrivingBehaviorAnalysis) XXX_Size() int {
	return xxx_messageInfo_DrivingBehaviorAnalysis.Size(m)
}
func (m *DrivingBehaviorAnalysis) XXX_DiscardUnknown() {
	xxx_messageInfo_DrivingBehaviorAnalysis.DiscardUnknown(m)
}

var xxx_messageInfo_DrivingBehaviorAnalysis proto.InternalMessageInfo

func (m *DrivingBehaviorAnalysis) GetTurningAngle() int32 {
	if m != nil && m.TurningAngle != nil {
		return *m.TurningAngle
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetLowOilDrivingValue() int32 {
	if m != nil && m.LowOilDrivingValue != nil {
		return *m.LowOilDrivingValue
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetVehicleSpeedFromEcu() int32 {
	if m != nil && m.VehicleSpeedFromEcu != nil {
		return *m.VehicleSpeedFromEcu
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetRpmWhenAlarming() int32 {
	if m != nil && m.RpmWhenAlarming != nil {
		return *m.RpmWhenAlarming
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetVelocityChangeValue() int32 {
	if m != nil && m.VelocityChangeValue != nil {
		return *m.VelocityChangeValue
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetCurrentGearshift() int32 {
	if m != nil && m.CurrentGearshift != nil {
		return *m.CurrentGearshift
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetCurrentRPM() int32 {
	if m != nil && m.CurrentRPM != nil {
		return *m.CurrentRPM
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetDeviceIdentity() string {
	if m != nil && m.DeviceIdentity != nil {
		return *m.DeviceIdentity
	}
	return ""
}

func (m *DrivingBehaviorAnalysis) GetBrakeTimes() int32 {
	if m != nil && m.BrakeTimes != nil {
		return *m.BrakeTimes
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetClutchTimes() int32 {
	if m != nil && m.ClutchTimes != nil {
		return *m.ClutchTimes
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetRetarderTimes() int32 {
	if m != nil && m.RetarderTimes != nil {
		return *m.RetarderTimes
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetAbsTimes() int32 {
	if m != nil && m.AbsTimes != nil {
		return *m.AbsTimes
	}
	return 0
}

func (m *DrivingBehaviorAnalysis) GetReverseTimes() int32 {
	if m != nil && m.ReverseTimes != nil {
		return *m.ReverseTimes
	}
	return 0
}

func init() {
	proto.RegisterType((*DrivingBehaviorAnalysis)(nil), "proto_dic.DrivingBehaviorAnalysis")
}

func init() {
	proto.RegisterFile("DrivingBehaviorAnalysis.proto", fileDescriptor_DrivingBehaviorAnalysis_c4e7ce55d1a3a8d9)
}

var fileDescriptor_DrivingBehaviorAnalysis_c4e7ce55d1a3a8d9 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0xa9, 0xb6, 0xda, 0x4e, 0x5b, 0x95, 0x78, 0x30, 0x08, 0x4a, 0x29, 0x22, 0xc5, 0x43,
	0xf1, 0x2f, 0xd4, 0x4f, 0x3c, 0x88, 0x52, 0x45, 0x8f, 0x92, 0x66, 0xc7, 0xdd, 0xc1, 0x6c, 0xb6,
	0xcc, 0x66, 0x57, 0xfa, 0x4b, 0xfc, 0xbb, 0xd2, 0xa4, 0x6a, 0xbf, 0x3c, 0x85, 0x3c, 0xef, 0x33,
	0x2f, 0x03, 0x03, 0x47, 0x57, 0x4c, 0x25, 0xd9, 0xf8, 0x02, 0x13, 0x55, 0x52, 0xc6, 0x03, 0xab,
	0xcc, 0x24, 0xa7, 0xbc, 0x3f, 0xe6, 0xcc, 0x65, 0xa2, 0xe1, 0x9f, 0xb7, 0x88, 0x74, 0xf7, 0xab,
	0x0a, 0x07, 0xff, 0xc8, 0xa2, 0x0b, 0x2d, 0x57, 0xb0, 0x25, 0x1b, 0x0f, 0x6c, 0x6c, 0x50, 0x56,
	0x3a, 0x95, 0x5e, 0x6d, 0xb8, 0xc0, 0x44, 0x1f, 0x84, 0xc9, 0x3e, 0x1f, 0xc8, 0xcc, 0x4a, 0x5e,
	0x94, 0x29, 0x50, 0x6e, 0x78, 0x73, 0x4d, 0x22, 0xce, 0x61, 0xbf, 0xc4, 0x84, 0xb4, 0xc1, 0xa7,
	0x31, 0x62, 0x74, 0xc3, 0x59, 0x7a, 0xad, 0x0b, 0xb9, 0xe9, 0x07, 0xd6, 0x45, 0xa2, 0x07, 0xbb,
	0x3c, 0x4e, 0x5f, 0x13, 0xb4, 0x03, 0xa3, 0x38, 0x25, 0x1b, 0xcb, 0xaa, 0xb7, 0x97, 0x71, 0xe8,
	0x36, 0x99, 0x26, 0x37, 0xb9, 0x4c, 0x94, 0x8d, 0x31, 0x2c, 0x53, 0xfb, 0xe9, 0x5e, 0x89, 0xc4,
	0x19, 0xec, 0xe9, 0x82, 0x19, 0xad, 0xbb, 0x45, 0xc5, 0x79, 0x42, 0xef, 0x4e, 0x6e, 0x79, 0x7d,
	0x85, 0x8b, 0x63, 0x80, 0x19, 0x1b, 0x3e, 0xde, 0xcb, 0x6d, 0x6f, 0xcd, 0x11, 0x71, 0x0a, 0x3b,
	0x11, 0x96, 0xa4, 0xf1, 0x2e, 0x42, 0xeb, 0xc8, 0x4d, 0x64, 0xbd, 0x53, 0xe9, 0x35, 0x86, 0x4b,
	0x74, 0xda, 0x33, 0x62, 0xf5, 0x81, 0xcf, 0x94, 0x62, 0x2e, 0x1b, 0xa1, 0xe7, 0x8f, 0x88, 0x0e,
	0x34, 0xb5, 0x29, 0x9c, 0x4e, 0x82, 0x00, 0x5e, 0x98, 0x47, 0xe2, 0x04, 0xda, 0x8c, 0x4e, 0x71,
	0x84, 0x1c, 0x9c, 0xa6, 0x77, 0x16, 0xa1, 0x38, 0x84, 0xba, 0x1a, 0xe5, 0x41, 0x68, 0x79, 0xe1,
	0xf7, 0x3f, 0xbd, 0x2c, 0x63, 0x89, 0x9c, 0xcf, 0xb6, 0x68, 0x87, 0xcb, 0xce, 0xb3, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x2b, 0xdf, 0x81, 0x44, 0x02, 0x00, 0x00,
}
