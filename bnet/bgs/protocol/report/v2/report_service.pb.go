// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: api/client/v2/report_service.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/account/v1"
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

type SubmitReportRequest struct {
	AgentId         *v1.AccountId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	UserDescription *string       `protobuf:"bytes,2,opt,name=user_description,json=userDescription" json:"user_description,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*SubmitReportRequest_UserOptions
	//	*SubmitReportRequest_ClubOptions
	//	*SubmitReportRequest_EntityOptions
	Type                 isSubmitReportRequest_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SubmitReportRequest) Reset()         { *m = SubmitReportRequest{} }
func (m *SubmitReportRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitReportRequest) ProtoMessage()    {}
func (*SubmitReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4fd3cd45fc6f72a, []int{0}
}

func (m *SubmitReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitReportRequest.Unmarshal(m, b)
}
func (m *SubmitReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitReportRequest.Marshal(b, m, deterministic)
}
func (m *SubmitReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitReportRequest.Merge(m, src)
}
func (m *SubmitReportRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitReportRequest.Size(m)
}
func (m *SubmitReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitReportRequest proto.InternalMessageInfo

func (m *SubmitReportRequest) GetAgentId() *v1.AccountId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubmitReportRequest) GetUserDescription() string {
	if m != nil && m.UserDescription != nil {
		return *m.UserDescription
	}
	return ""
}

type isSubmitReportRequest_Type interface {
	isSubmitReportRequest_Type()
}

type SubmitReportRequest_UserOptions struct {
	UserOptions *UserOptions `protobuf:"bytes,10,opt,name=user_options,json=userOptions,oneof"`
}

type SubmitReportRequest_ClubOptions struct {
	ClubOptions *ClubOptions `protobuf:"bytes,11,opt,name=club_options,json=clubOptions,oneof"`
}

type SubmitReportRequest_EntityOptions struct {
	EntityOptions *EntityOptions `protobuf:"bytes,20,opt,name=entity_options,json=entityOptions,oneof"`
}

func (*SubmitReportRequest_UserOptions) isSubmitReportRequest_Type() {}

func (*SubmitReportRequest_ClubOptions) isSubmitReportRequest_Type() {}

func (*SubmitReportRequest_EntityOptions) isSubmitReportRequest_Type() {}

func (m *SubmitReportRequest) GetType() isSubmitReportRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SubmitReportRequest) GetUserOptions() *UserOptions {
	if x, ok := m.GetType().(*SubmitReportRequest_UserOptions); ok {
		return x.UserOptions
	}
	return nil
}

func (m *SubmitReportRequest) GetClubOptions() *ClubOptions {
	if x, ok := m.GetType().(*SubmitReportRequest_ClubOptions); ok {
		return x.ClubOptions
	}
	return nil
}

func (m *SubmitReportRequest) GetEntityOptions() *EntityOptions {
	if x, ok := m.GetType().(*SubmitReportRequest_EntityOptions); ok {
		return x.EntityOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubmitReportRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubmitReportRequest_UserOptions)(nil),
		(*SubmitReportRequest_ClubOptions)(nil),
		(*SubmitReportRequest_EntityOptions)(nil),
	}
}

func init() {
	proto.RegisterType((*SubmitReportRequest)(nil), "bgs.protocol.report.v2.SubmitReportRequest")
}

func init() {
	proto.RegisterFile("api/client/v2/report_service.proto", fileDescriptor_d4fd3cd45fc6f72a)
}

var fileDescriptor_d4fd3cd45fc6f72a = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0x20, 0x6b, 0x9d, 0x6d, 0x6d, 0x49, 0x8b, 0x2c, 0x39, 0x2d, 0x91, 0x42, 0x35,
	0x30, 0x53, 0x73, 0x13, 0xf4, 0xd0, 0x5a, 0xa1, 0xbd, 0x54, 0x48, 0xf1, 0xa0, 0x97, 0x90, 0x4c,
	0x9e, 0xe9, 0x40, 0x9a, 0x19, 0x67, 0xde, 0x04, 0xf6, 0x26, 0x7b, 0xf4, 0x7b, 0xf8, 0xe9, 0xfc,
	0x04, 0x7b, 0x93, 0x24, 0xeb, 0x26, 0x0b, 0xd9, 0xde, 0xe6, 0xfd, 0xf3, 0xcb, 0xef, 0x4d, 0xde,
	0x0b, 0x09, 0x52, 0x25, 0x18, 0x2f, 0x05, 0x54, 0xc8, 0xea, 0x88, 0x69, 0x50, 0x52, 0x63, 0x62,
	0x40, 0xd7, 0x82, 0x03, 0x55, 0x5a, 0xa2, 0xf4, 0x5e, 0x65, 0x85, 0xe9, 0x8e, 0x5c, 0x96, 0xb4,
	0x43, 0x68, 0x1d, 0xf9, 0x27, 0x29, 0xe7, 0xd2, 0x56, 0x98, 0xe0, 0x42, 0xc1, 0x9a, 0xf0, 0xe7,
	0xa3, 0xc2, 0x21, 0x71, 0xa4, 0x15, 0x1f, 0x06, 0xc1, 0x5f, 0x97, 0x9c, 0xdc, 0xdb, 0xec, 0x51,
	0x60, 0xdc, 0xd2, 0x31, 0xfc, 0xb4, 0x60, 0xd0, 0xfb, 0x48, 0xf6, 0xd3, 0x02, 0x2a, 0x4c, 0x44,
	0x3e, 0x73, 0xe6, 0xce, 0xf9, 0x34, 0x0a, 0xe8, 0xd6, 0x55, 0xd6, 0xfd, 0x69, 0xfd, 0x8e, 0x5e,
	0x76, 0xc7, 0xdb, 0x3c, 0x7e, 0xde, 0xbe, 0x73, 0x9b, 0x7b, 0x6f, 0xc8, 0xb1, 0x35, 0xa0, 0x93,
	0x1c, 0x0c, 0xd7, 0x42, 0xa1, 0x90, 0xd5, 0xcc, 0x9d, 0x3b, 0xe7, 0x2f, 0xe2, 0xa3, 0x26, 0xbf,
	0xee, 0x63, 0xef, 0x86, 0x1c, 0xb4, 0xa8, 0x6c, 0x4b, 0x33, 0x23, 0x6d, 0xb7, 0xd7, 0x74, 0xfc,
	0xc3, 0xe9, 0x57, 0x03, 0xfa, 0x4b, 0x87, 0xde, 0xec, 0xc5, 0x53, 0xdb, 0x97, 0x8d, 0x89, 0x97,
	0x36, 0xdb, 0x98, 0xa6, 0x4f, 0x9b, 0x3e, 0x95, 0x36, 0x1b, 0x98, 0x78, 0x5f, 0x7a, 0x77, 0xe4,
	0x25, 0x54, 0x28, 0x70, 0xb1, 0x71, 0x9d, 0xb6, 0xae, 0xb3, 0x5d, 0xae, 0xcf, 0x2d, 0xdd, 0xdb,
	0x0e, 0x61, 0x18, 0x5c, 0x4d, 0xc8, 0xb3, 0x66, 0xe8, 0xd1, 0x1f, 0x87, 0x1c, 0x76, 0x73, 0xbe,
	0xef, 0xb6, 0xec, 0x7d, 0x23, 0x07, 0xc3, 0xf1, 0x7b, 0xe1, 0xae, 0x0e, 0x23, 0x4b, 0xf2, 0x4f,
	0xb7, 0xe1, 0x3b, 0x79, 0x9d, 0x62, 0x1a, 0x4c, 0x96, 0xab, 0xd0, 0xdd, 0x77, 0xfc, 0xf7, 0xcb,
	0x55, 0xc8, 0xc8, 0x59, 0x56, 0x01, 0x8e, 0x09, 0xb7, 0xee, 0xf1, 0x76, 0xd2, 0x3d, 0xf8, 0xbd,
	0x0a, 0xdd, 0x63, 0xe7, 0xea, 0xf2, 0xfb, 0x87, 0x42, 0xe0, 0x83, 0xcd, 0x28, 0x97, 0x8f, 0xcc,
	0x58, 0x05, 0x5a, 0x5d, 0x5c, 0x20, 0x2b, 0xa4, 0x7a, 0x00, 0xcd, 0x75, 0xfa, 0x03, 0x59, 0xe3,
	0x65, 0x59, 0x61, 0xd8, 0x7f, 0xf7, 0xfa, 0x7f, 0x63, 0x75, 0xf4, 0xcb, 0xd9, 0xfb, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x56, 0x08, 0x31, 0xdd, 0x02, 0x00, 0x00,
}
