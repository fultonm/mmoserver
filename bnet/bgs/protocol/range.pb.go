// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: global_extensions/range.proto

package protocol

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

type UnsignedIntRange struct {
	Min                  *uint64  `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max                  *uint64  `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnsignedIntRange) Reset()         { *m = UnsignedIntRange{} }
func (m *UnsignedIntRange) String() string { return proto.CompactTextString(m) }
func (*UnsignedIntRange) ProtoMessage()    {}
func (*UnsignedIntRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8991336042d8ec59, []int{0}
}

func (m *UnsignedIntRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsignedIntRange.Unmarshal(m, b)
}
func (m *UnsignedIntRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsignedIntRange.Marshal(b, m, deterministic)
}
func (m *UnsignedIntRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedIntRange.Merge(m, src)
}
func (m *UnsignedIntRange) XXX_Size() int {
	return xxx_messageInfo_UnsignedIntRange.Size(m)
}
func (m *UnsignedIntRange) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedIntRange.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedIntRange proto.InternalMessageInfo

func (m *UnsignedIntRange) GetMin() uint64 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *UnsignedIntRange) GetMax() uint64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

type SignedIntRange struct {
	Min                  *int64   `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max                  *int64   `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedIntRange) Reset()         { *m = SignedIntRange{} }
func (m *SignedIntRange) String() string { return proto.CompactTextString(m) }
func (*SignedIntRange) ProtoMessage()    {}
func (*SignedIntRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8991336042d8ec59, []int{1}
}

func (m *SignedIntRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedIntRange.Unmarshal(m, b)
}
func (m *SignedIntRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedIntRange.Marshal(b, m, deterministic)
}
func (m *SignedIntRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedIntRange.Merge(m, src)
}
func (m *SignedIntRange) XXX_Size() int {
	return xxx_messageInfo_SignedIntRange.Size(m)
}
func (m *SignedIntRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedIntRange.DiscardUnknown(m)
}

var xxx_messageInfo_SignedIntRange proto.InternalMessageInfo

func (m *SignedIntRange) GetMin() int64 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *SignedIntRange) GetMax() int64 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

type FloatRange struct {
	Min                  *float32 `protobuf:"fixed32,1,opt,name=min" json:"min,omitempty"`
	Max                  *float32 `protobuf:"fixed32,2,opt,name=max" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatRange) Reset()         { *m = FloatRange{} }
func (m *FloatRange) String() string { return proto.CompactTextString(m) }
func (*FloatRange) ProtoMessage()    {}
func (*FloatRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8991336042d8ec59, []int{2}
}

func (m *FloatRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatRange.Unmarshal(m, b)
}
func (m *FloatRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatRange.Marshal(b, m, deterministic)
}
func (m *FloatRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatRange.Merge(m, src)
}
func (m *FloatRange) XXX_Size() int {
	return xxx_messageInfo_FloatRange.Size(m)
}
func (m *FloatRange) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatRange.DiscardUnknown(m)
}

var xxx_messageInfo_FloatRange proto.InternalMessageInfo

func (m *FloatRange) GetMin() float32 {
	if m != nil && m.Min != nil {
		return *m.Min
	}
	return 0
}

func (m *FloatRange) GetMax() float32 {
	if m != nil && m.Max != nil {
		return *m.Max
	}
	return 0
}

func init() {
	proto.RegisterType((*UnsignedIntRange)(nil), "bgs.protocol.UnsignedIntRange")
	proto.RegisterType((*SignedIntRange)(nil), "bgs.protocol.SignedIntRange")
	proto.RegisterType((*FloatRange)(nil), "bgs.protocol.FloatRange")
}

func init() { proto.RegisterFile("global_extensions/range.proto", fileDescriptor_8991336042d8ec59) }

var fileDescriptor_8991336042d8ec59 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0x31, 0xcb, 0x83, 0x30,
	0x10, 0xc6, 0x71, 0xd4, 0x77, 0x0a, 0x2f, 0x45, 0x9c, 0x5c, 0x0a, 0xc5, 0xa9, 0x93, 0x91, 0x22,
	0xfd, 0x00, 0x1d, 0x0a, 0x5d, 0x2d, 0x5d, 0xba, 0x94, 0xc4, 0xa6, 0x31, 0x10, 0xef, 0x42, 0x72,
	0x82, 0x1f, 0xbf, 0xa8, 0xed, 0x24, 0xdd, 0x1e, 0xfe, 0xdc, 0x6f, 0x38, 0xb6, 0xd5, 0x16, 0xa5,
	0xb0, 0x0f, 0x35, 0x92, 0x82, 0x60, 0x10, 0x02, 0xf7, 0x02, 0xb4, 0x2a, 0x9d, 0x47, 0xc2, 0xec,
	0x5f, 0xea, 0xb0, 0xcc, 0x16, 0x6d, 0x71, 0x64, 0xe9, 0x0d, 0x82, 0xd1, 0xa0, 0x9e, 0x17, 0xa0,
	0x66, 0xba, 0xcb, 0x52, 0x96, 0xf4, 0x06, 0xf2, 0x68, 0x17, 0xed, 0xff, 0x9a, 0x69, 0xce, 0x45,
	0x8c, 0x79, 0xfc, 0x29, 0x62, 0x2c, 0x6a, 0xb6, 0xb9, 0xfe, 0x54, 0xc9, 0x4a, 0x25, 0x8b, 0xaa,
	0x18, 0x3b, 0x5b, 0x14, 0x6b, 0x11, 0xaf, 0x44, 0x3c, 0x8b, 0x53, 0x7d, 0x3f, 0x68, 0x43, 0xdd,
	0x20, 0xcb, 0x16, 0x7b, 0x1e, 0x06, 0xa7, 0xbc, 0xab, 0x2a, 0xe2, 0x1a, 0x5d, 0xa7, 0x7c, 0xeb,
	0xc5, 0x8b, 0xb8, 0x04, 0x45, 0x5c, 0xea, 0xc0, 0xbf, 0x5f, 0xbd, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x61, 0x0b, 0x49, 0x92, 0x03, 0x01, 0x00, 0x00,
}
