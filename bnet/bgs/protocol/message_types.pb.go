// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: message_types.proto

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

type TypingIndicator int32

const (
	TypingIndicator_TYPING_START TypingIndicator = 0
	TypingIndicator_TYPING_STOP  TypingIndicator = 1
)

var TypingIndicator_name = map[int32]string{
	0: "TYPING_START",
	1: "TYPING_STOP",
}

var TypingIndicator_value = map[string]int32{
	"TYPING_START": 0,
	"TYPING_STOP":  1,
}

func (x TypingIndicator) Enum() *TypingIndicator {
	p := new(TypingIndicator)
	*p = x
	return p
}

func (x TypingIndicator) String() string {
	return proto.EnumName(TypingIndicator_name, int32(x))
}

func (x *TypingIndicator) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TypingIndicator_value, data, "TypingIndicator")
	if err != nil {
		return err
	}
	*x = TypingIndicator(value)
	return nil
}

func (TypingIndicator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_804025860b4453ba, []int{0}
}

type MessageId struct {
	Epoch                *uint64  `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	Position             *uint64  `protobuf:"varint,2,opt,name=position" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageId) Reset()         { *m = MessageId{} }
func (m *MessageId) String() string { return proto.CompactTextString(m) }
func (*MessageId) ProtoMessage()    {}
func (*MessageId) Descriptor() ([]byte, []int) {
	return fileDescriptor_804025860b4453ba, []int{0}
}

func (m *MessageId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageId.Unmarshal(m, b)
}
func (m *MessageId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageId.Marshal(b, m, deterministic)
}
func (m *MessageId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageId.Merge(m, src)
}
func (m *MessageId) XXX_Size() int {
	return xxx_messageInfo_MessageId.Size(m)
}
func (m *MessageId) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageId.DiscardUnknown(m)
}

var xxx_messageInfo_MessageId proto.InternalMessageInfo

func (m *MessageId) GetEpoch() uint64 {
	if m != nil && m.Epoch != nil {
		return *m.Epoch
	}
	return 0
}

func (m *MessageId) GetPosition() uint64 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func init() {
	proto.RegisterEnum("bgs.protocol.TypingIndicator", TypingIndicator_name, TypingIndicator_value)
	proto.RegisterType((*MessageId)(nil), "bgs.protocol.MessageId")
}

func init() { proto.RegisterFile("message_types.proto", fileDescriptor_804025860b4453ba) }

var fileDescriptor_804025860b4453ba = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x49, 0x4a, 0x87, 0x32, 0x93, 0xf3, 0x73, 0x94, 0x6c, 0xb9, 0x38, 0x7d, 0x21, 0x8a, 0x3c,
	0x53, 0x84, 0x44, 0xb8, 0x58, 0x53, 0x0b, 0xf2, 0x93, 0x33, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x82, 0x20, 0x1c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xfc, 0xe2, 0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x09,
	0x26, 0xb0, 0x04, 0x9c, 0xaf, 0x65, 0xc2, 0xc5, 0x1f, 0x52, 0x59, 0x90, 0x99, 0x97, 0xee, 0x99,
	0x97, 0x92, 0x99, 0x9c, 0x58, 0x92, 0x5f, 0x24, 0x24, 0xc0, 0xc5, 0x13, 0x12, 0x19, 0xe0, 0xe9,
	0xe7, 0x1e, 0x1f, 0x1c, 0xe2, 0x18, 0x14, 0x22, 0xc0, 0x20, 0xc4, 0xcf, 0xc5, 0x0d, 0x17, 0xf1,
	0x0f, 0x10, 0x60, 0x74, 0x32, 0x89, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x2f, 0x2e, 0x2d, 0x48, 0x2d, 0x2a, 0x30, 0x30, 0x28, 0xd1, 0x4f, 0xcf, 0x2f, 0xc8,
	0x48, 0x2d, 0x4a, 0x2e, 0x4a, 0x4c, 0x2b, 0xd1, 0x4f, 0xca, 0x4b, 0x2d, 0xd1, 0x4f, 0x4a, 0x2f,
	0xd6, 0x87, 0x39, 0x15, 0x10, 0x00, 0x00, 0xff, 0xff, 0x49, 0xa3, 0x4f, 0xc6, 0xce, 0x00, 0x00,
	0x00,
}
