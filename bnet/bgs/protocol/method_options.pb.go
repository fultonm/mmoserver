// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: global_extensions/method_options.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type BGSMethodOptions struct {
	Id                   *uint32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BGSMethodOptions) Reset()         { *m = BGSMethodOptions{} }
func (m *BGSMethodOptions) String() string { return proto.CompactTextString(m) }
func (*BGSMethodOptions) ProtoMessage()    {}
func (*BGSMethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cba562a9dddbbfa, []int{0}
}

func (m *BGSMethodOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BGSMethodOptions.Unmarshal(m, b)
}
func (m *BGSMethodOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BGSMethodOptions.Marshal(b, m, deterministic)
}
func (m *BGSMethodOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BGSMethodOptions.Merge(m, src)
}
func (m *BGSMethodOptions) XXX_Size() int {
	return xxx_messageInfo_BGSMethodOptions.Size(m)
}
func (m *BGSMethodOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_BGSMethodOptions.DiscardUnknown(m)
}

var xxx_messageInfo_BGSMethodOptions proto.InternalMessageInfo

func (m *BGSMethodOptions) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

var E_MethodOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*BGSMethodOptions)(nil),
	Field:         90000,
	Name:          "bgs.protocol.method_options",
	Tag:           "bytes,90000,opt,name=method_options",
	Filename:      "global_extensions/method_options.proto",
}

func init() {
	proto.RegisterType((*BGSMethodOptions)(nil), "bgs.protocol.BGSMethodOptions")
	proto.RegisterExtension(E_MethodOptions)
}

func init() {
	proto.RegisterFile("global_extensions/method_options.proto", fileDescriptor_3cba562a9dddbbfa)
}

var fileDescriptor_3cba562a9dddbbfa = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xce, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x1c, 0xe8, 0xa2, 0xd6, 0xa1, 0x78, 0x0a, 0x1d, 0x82, 0xc9, 0x50, 0x32, 0xe9,
	0x42, 0xc6, 0x8e, 0x5e, 0x3a, 0x95, 0x96, 0x74, 0x28, 0x74, 0x09, 0x91, 0x75, 0x39, 0x0b, 0x64,
	0x9f, 0x90, 0xce, 0xd0, 0xc7, 0xe8, 0x1b, 0xf5, 0xd5, 0x8a, 0xe3, 0xa6, 0xc4, 0xd9, 0x74, 0xc7,
	0xaf, 0xef, 0x3f, 0xf5, 0x48, 0x9e, 0xcd, 0xc1, 0xef, 0xf1, 0x4b, 0xb0, 0x4b, 0x8e, 0xbb, 0x04,
	0x2d, 0x4a, 0xc3, 0x76, 0xcf, 0x41, 0x86, 0x51, 0x87, 0xc8, 0xc2, 0xc5, 0x9d, 0xa1, 0xbf, 0x67,
	0xcd, 0xfe, 0xa1, 0x24, 0x66, 0xf2, 0x08, 0xa7, 0x85, 0xe9, 0x8f, 0x60, 0x31, 0xd5, 0xd1, 0x05,
	0xe1, 0x38, 0x86, 0x56, 0x2b, 0x75, 0x5f, 0x3d, 0xbf, 0xbf, 0x9c, 0xa8, 0xd7, 0x51, 0x2a, 0xe6,
	0x6a, 0xe6, 0xec, 0x22, 0x2b, 0xb3, 0x75, 0xbe, 0x9b, 0x39, 0xfb, 0x44, 0x6a, 0x3e, 0xed, 0x2a,
	0x96, 0x7a, 0x84, 0xf5, 0x19, 0xd6, 0x13, 0x61, 0xf1, 0xfd, 0x73, 0x53, 0x66, 0xeb, 0xdb, 0xed,
	0x52, 0x5f, 0x9e, 0xa3, 0xaf, 0x9b, 0x76, 0x79, 0x7b, 0x39, 0x56, 0x1f, 0x2a, 0x37, 0x1d, 0xca,
	0xff, 0x87, 0xaa, 0x98, 0xc4, 0xdf, 0x86, 0xf5, 0xe7, 0x96, 0x9c, 0x34, 0xbd, 0xd1, 0x35, 0xb7,
	0x90, 0xfa, 0x80, 0x31, 0x6c, 0x36, 0x02, 0xc4, 0xa1, 0xc1, 0x58, 0xc7, 0xc3, 0x51, 0x60, 0x30,
	0xc0, 0x50, 0x82, 0xb3, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x8e, 0x15, 0x47, 0x3e, 0x01,
	0x00, 0x00,
}
