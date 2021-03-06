// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TireTemperatureAddition.proto

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

type TireTemperatureAddition struct {
	VehType              *int32               `protobuf:"varint,1,opt,name=vehType" json:"vehType,omitempty"`
	TyreCount            *int32               `protobuf:"varint,2,opt,name=tyreCount" json:"tyreCount,omitempty"`
	TireConditionItem    []*TireConditionItem `protobuf:"bytes,3,rep,name=tireConditionItem" json:"tireConditionItem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TireTemperatureAddition) Reset()         { *m = TireTemperatureAddition{} }
func (m *TireTemperatureAddition) String() string { return proto.CompactTextString(m) }
func (*TireTemperatureAddition) ProtoMessage()    {}
func (*TireTemperatureAddition) Descriptor() ([]byte, []int) {
	return fileDescriptor_TireTemperatureAddition_c5df332303b7e397, []int{0}
}
func (m *TireTemperatureAddition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TireTemperatureAddition.Unmarshal(m, b)
}
func (m *TireTemperatureAddition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TireTemperatureAddition.Marshal(b, m, deterministic)
}
func (dst *TireTemperatureAddition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TireTemperatureAddition.Merge(dst, src)
}
func (m *TireTemperatureAddition) XXX_Size() int {
	return xxx_messageInfo_TireTemperatureAddition.Size(m)
}
func (m *TireTemperatureAddition) XXX_DiscardUnknown() {
	xxx_messageInfo_TireTemperatureAddition.DiscardUnknown(m)
}

var xxx_messageInfo_TireTemperatureAddition proto.InternalMessageInfo

func (m *TireTemperatureAddition) GetVehType() int32 {
	if m != nil && m.VehType != nil {
		return *m.VehType
	}
	return 0
}

func (m *TireTemperatureAddition) GetTyreCount() int32 {
	if m != nil && m.TyreCount != nil {
		return *m.TyreCount
	}
	return 0
}

func (m *TireTemperatureAddition) GetTireConditionItem() []*TireConditionItem {
	if m != nil {
		return m.TireConditionItem
	}
	return nil
}

type TireConditionItem struct {
	TyreHorizontalPosition *int32   `protobuf:"varint,1,opt,name=tyreHorizontalPosition" json:"tyreHorizontalPosition,omitempty"`
	TyreVerticalPosition   *int32   `protobuf:"varint,2,opt,name=tyreVerticalPosition" json:"tyreVerticalPosition,omitempty"`
	BatteryStatus          *int32   `protobuf:"varint,3,opt,name=batteryStatus" json:"batteryStatus,omitempty"`
	UnBalance              *int32   `protobuf:"varint,4,opt,name=unBalance" json:"unBalance,omitempty"`
	AirLeakage             *int32   `protobuf:"varint,5,opt,name=airLeakage" json:"airLeakage,omitempty"`
	HighTemperature        *int32   `protobuf:"varint,6,opt,name=highTemperature" json:"highTemperature,omitempty"`
	NoRFSignal             *int32   `protobuf:"varint,7,opt,name=noRFSignal" json:"noRFSignal,omitempty"`
	NoMatch                *int32   `protobuf:"varint,8,opt,name=noMatch" json:"noMatch,omitempty"`
	TirePressureAlarm      *int32   `protobuf:"varint,9,opt,name=tirePressureAlarm" json:"tirePressureAlarm,omitempty"`
	TyrePressure           *int32   `protobuf:"varint,10,opt,name=tyrePressure" json:"tyrePressure,omitempty"`
	TyreTemperature        *int32   `protobuf:"varint,11,opt,name=tyreTemperature" json:"tyreTemperature,omitempty"`
	TemAlarmThreshold      *int32   `protobuf:"varint,12,opt,name=temAlarmThreshold" json:"temAlarmThreshold,omitempty"`
	TyreHalarmThreshold    *int32   `protobuf:"varint,13,opt,name=tyreHalarmThreshold" json:"tyreHalarmThreshold,omitempty"`
	TyreLalarmThreshold    *int32   `protobuf:"varint,14,opt,name=tyreLalarmThreshold" json:"tyreLalarmThreshold,omitempty"`
	TyreNomimalValue       *int32   `protobuf:"varint,15,opt,name=tyreNomimalValue" json:"tyreNomimalValue,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TireConditionItem) Reset()         { *m = TireConditionItem{} }
func (m *TireConditionItem) String() string { return proto.CompactTextString(m) }
func (*TireConditionItem) ProtoMessage()    {}
func (*TireConditionItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_TireTemperatureAddition_c5df332303b7e397, []int{1}
}
func (m *TireConditionItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TireConditionItem.Unmarshal(m, b)
}
func (m *TireConditionItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TireConditionItem.Marshal(b, m, deterministic)
}
func (dst *TireConditionItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TireConditionItem.Merge(dst, src)
}
func (m *TireConditionItem) XXX_Size() int {
	return xxx_messageInfo_TireConditionItem.Size(m)
}
func (m *TireConditionItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TireConditionItem.DiscardUnknown(m)
}

var xxx_messageInfo_TireConditionItem proto.InternalMessageInfo

func (m *TireConditionItem) GetTyreHorizontalPosition() int32 {
	if m != nil && m.TyreHorizontalPosition != nil {
		return *m.TyreHorizontalPosition
	}
	return 0
}

func (m *TireConditionItem) GetTyreVerticalPosition() int32 {
	if m != nil && m.TyreVerticalPosition != nil {
		return *m.TyreVerticalPosition
	}
	return 0
}

func (m *TireConditionItem) GetBatteryStatus() int32 {
	if m != nil && m.BatteryStatus != nil {
		return *m.BatteryStatus
	}
	return 0
}

func (m *TireConditionItem) GetUnBalance() int32 {
	if m != nil && m.UnBalance != nil {
		return *m.UnBalance
	}
	return 0
}

func (m *TireConditionItem) GetAirLeakage() int32 {
	if m != nil && m.AirLeakage != nil {
		return *m.AirLeakage
	}
	return 0
}

func (m *TireConditionItem) GetHighTemperature() int32 {
	if m != nil && m.HighTemperature != nil {
		return *m.HighTemperature
	}
	return 0
}

func (m *TireConditionItem) GetNoRFSignal() int32 {
	if m != nil && m.NoRFSignal != nil {
		return *m.NoRFSignal
	}
	return 0
}

func (m *TireConditionItem) GetNoMatch() int32 {
	if m != nil && m.NoMatch != nil {
		return *m.NoMatch
	}
	return 0
}

func (m *TireConditionItem) GetTirePressureAlarm() int32 {
	if m != nil && m.TirePressureAlarm != nil {
		return *m.TirePressureAlarm
	}
	return 0
}

func (m *TireConditionItem) GetTyrePressure() int32 {
	if m != nil && m.TyrePressure != nil {
		return *m.TyrePressure
	}
	return 0
}

func (m *TireConditionItem) GetTyreTemperature() int32 {
	if m != nil && m.TyreTemperature != nil {
		return *m.TyreTemperature
	}
	return 0
}

func (m *TireConditionItem) GetTemAlarmThreshold() int32 {
	if m != nil && m.TemAlarmThreshold != nil {
		return *m.TemAlarmThreshold
	}
	return 0
}

func (m *TireConditionItem) GetTyreHalarmThreshold() int32 {
	if m != nil && m.TyreHalarmThreshold != nil {
		return *m.TyreHalarmThreshold
	}
	return 0
}

func (m *TireConditionItem) GetTyreLalarmThreshold() int32 {
	if m != nil && m.TyreLalarmThreshold != nil {
		return *m.TyreLalarmThreshold
	}
	return 0
}

func (m *TireConditionItem) GetTyreNomimalValue() int32 {
	if m != nil && m.TyreNomimalValue != nil {
		return *m.TyreNomimalValue
	}
	return 0
}

func init() {
	proto.RegisterType((*TireTemperatureAddition)(nil), "proto_dic.TireTemperatureAddition")
	proto.RegisterType((*TireConditionItem)(nil), "proto_dic.TireConditionItem")
}

func init() {
	proto.RegisterFile("TireTemperatureAddition.proto", fileDescriptor_TireTemperatureAddition_c5df332303b7e397)
}

var fileDescriptor_TireTemperatureAddition_c5df332303b7e397 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x83, 0x88, 0xc8, 0x05, 0x44, 0x46, 0xa3, 0xb3, 0x40, 0x43, 0x88, 0x0b, 0x62, 0x0c,
	0x31, 0x2c, 0xdc, 0xab, 0x89, 0x51, 0x83, 0x86, 0x40, 0xc3, 0xd6, 0x8c, 0xed, 0x0d, 0x9d, 0xd8,
	0xce, 0x90, 0xe9, 0xd4, 0x04, 0x77, 0x3e, 0x88, 0xef, 0x6a, 0x66, 0xca, 0x4f, 0x69, 0x61, 0xd5,
	0xdc, 0xef, 0x9c, 0xfb, 0xd3, 0x39, 0x70, 0xe1, 0x70, 0x85, 0x0e, 0x86, 0x33, 0x54, 0x4c, 0xc7,
	0x0a, 0xef, 0x3d, 0x8f, 0x6b, 0x2e, 0x45, 0x6f, 0xa6, 0xa4, 0x96, 0xa4, 0x62, 0x3f, 0x1f, 0x1e,
	0x77, 0x3b, 0x7f, 0x05, 0x38, 0xdf, 0x61, 0x26, 0x14, 0xca, 0xdf, 0xe8, 0x3b, 0xf3, 0x19, 0xd2,
	0x42, 0xbb, 0xd0, 0x2d, 0x8d, 0x96, 0x25, 0x69, 0x41, 0x45, 0xcf, 0x15, 0x3e, 0xca, 0x58, 0x68,
	0xba, 0x67, 0xb5, 0x35, 0x20, 0xaf, 0xd0, 0xd4, 0xdc, 0x14, 0x22, 0x19, 0xf4, 0xa2, 0x31, 0xa4,
	0xc5, 0x76, 0xb1, 0x5b, 0xed, 0xb7, 0x7a, 0xab, 0xd5, 0x3d, 0x27, 0xeb, 0x19, 0xe5, 0xdb, 0x3a,
	0xbf, 0x25, 0x68, 0xe6, 0x8c, 0xe4, 0x0e, 0xce, 0xcc, 0xba, 0x67, 0xa9, 0xf8, 0x8f, 0x14, 0x9a,
	0x05, 0x43, 0x19, 0x59, 0x75, 0x71, 0xe8, 0x0e, 0x95, 0xf4, 0xe1, 0xd4, 0x28, 0x13, 0x54, 0x9a,
	0xbb, 0xa9, 0xae, 0xe4, 0x17, 0xb6, 0x6a, 0xe4, 0x0a, 0xea, 0x9f, 0x4c, 0x6b, 0x54, 0xf3, 0xb1,
	0x66, 0x3a, 0x8e, 0x68, 0xd1, 0x9a, 0x37, 0xa1, 0x79, 0x91, 0x58, 0x3c, 0xb0, 0x80, 0x09, 0x17,
	0xe9, 0x7e, 0xf2, 0x22, 0x2b, 0x40, 0x2e, 0x01, 0x18, 0x57, 0x03, 0x64, 0x5f, 0x6c, 0x8a, 0xb4,
	0x64, 0xe5, 0x14, 0x21, 0x5d, 0x68, 0xf8, 0x7c, 0xea, 0xa7, 0x42, 0xa0, 0x07, 0xd6, 0x94, 0xc5,
	0x66, 0x92, 0x90, 0xa3, 0xa7, 0x31, 0x9f, 0x0a, 0x16, 0xd0, 0x72, 0x32, 0x69, 0x4d, 0x4c, 0x66,
	0x42, 0xbe, 0x31, 0xed, 0xfa, 0xf4, 0x30, 0xc9, 0x6c, 0x51, 0x92, 0x9b, 0x24, 0x95, 0xa1, 0xc2,
	0x28, 0x32, 0x29, 0x07, 0x4c, 0x85, 0xb4, 0x62, 0x3d, 0x79, 0x81, 0x74, 0xa0, 0x66, 0x5e, 0x63,
	0x09, 0x29, 0x58, 0xe3, 0x06, 0x33, 0x57, 0x9b, 0x3a, 0x7d, 0x75, 0x35, 0xb9, 0x3a, 0x83, 0xed,
	0x6e, 0x0c, 0xed, 0x64, 0xc7, 0x57, 0x18, 0xf9, 0x32, 0xf0, 0x68, 0x6d, 0xb1, 0x3b, 0x2b, 0x90,
	0x5b, 0x38, 0xb1, 0xf9, 0xb1, 0x4d, 0x7f, 0xdd, 0xfa, 0xb7, 0x49, 0xcb, 0x8e, 0x41, 0xa6, 0xe3,
	0x68, 0xdd, 0x91, 0x91, 0xc8, 0x35, 0x1c, 0x1b, 0xfc, 0x2e, 0x43, 0x1e, 0xb2, 0x60, 0xc2, 0x82,
	0x18, 0x69, 0xc3, 0xda, 0x73, 0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x29, 0x26, 0x88, 0xc4, 0x4e,
	0x03, 0x00, 0x00,
}
