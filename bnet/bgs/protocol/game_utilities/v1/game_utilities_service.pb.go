// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: game_utilities_service.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
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

type ClientRequest struct {
	Attribute            []*protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Host                 *protocol.ProcessId   `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"` // Deprecated: Do not use.
	AccountId            *protocol.EntityId    `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	GameAccountId        *protocol.EntityId    `protobuf:"bytes,4,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	Program              *uint32               `protobuf:"fixed32,5,opt,name=program" json:"program,omitempty"`
	ClientInfo           *ClientInfo           `protobuf:"bytes,6,opt,name=client_info,json=clientInfo" json:"client_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClientRequest) Reset()         { *m = ClientRequest{} }
func (m *ClientRequest) String() string { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()    {}
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{0}
}

func (m *ClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientRequest.Unmarshal(m, b)
}
func (m *ClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientRequest.Marshal(b, m, deterministic)
}
func (m *ClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRequest.Merge(m, src)
}
func (m *ClientRequest) XXX_Size() int {
	return xxx_messageInfo_ClientRequest.Size(m)
}
func (m *ClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRequest proto.InternalMessageInfo

func (m *ClientRequest) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

// Deprecated: Do not use.
func (m *ClientRequest) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ClientRequest) GetAccountId() *protocol.EntityId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *ClientRequest) GetGameAccountId() *protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *ClientRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ClientRequest) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

type ClientResponse struct {
	Attribute            []*protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClientResponse) Reset()         { *m = ClientResponse{} }
func (m *ClientResponse) String() string { return proto.CompactTextString(m) }
func (*ClientResponse) ProtoMessage()    {}
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{1}
}

func (m *ClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientResponse.Unmarshal(m, b)
}
func (m *ClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientResponse.Marshal(b, m, deterministic)
}
func (m *ClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientResponse.Merge(m, src)
}
func (m *ClientResponse) XXX_Size() int {
	return xxx_messageInfo_ClientResponse.Size(m)
}
func (m *ClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientResponse proto.InternalMessageInfo

func (m *ClientResponse) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type ServerRequest struct {
	Attribute            []*protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Program              *uint32               `protobuf:"fixed32,2,req,name=program" json:"program,omitempty"`
	Host                 *protocol.ProcessId   `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServerRequest) Reset()         { *m = ServerRequest{} }
func (m *ServerRequest) String() string { return proto.CompactTextString(m) }
func (*ServerRequest) ProtoMessage()    {}
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{2}
}

func (m *ServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerRequest.Unmarshal(m, b)
}
func (m *ServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerRequest.Marshal(b, m, deterministic)
}
func (m *ServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerRequest.Merge(m, src)
}
func (m *ServerRequest) XXX_Size() int {
	return xxx_messageInfo_ServerRequest.Size(m)
}
func (m *ServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerRequest proto.InternalMessageInfo

func (m *ServerRequest) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ServerRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerRequest) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type ServerResponse struct {
	Attribute            []*protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServerResponse) Reset()         { *m = ServerResponse{} }
func (m *ServerResponse) String() string { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()    {}
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{3}
}

func (m *ServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerResponse.Unmarshal(m, b)
}
func (m *ServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerResponse.Marshal(b, m, deterministic)
}
func (m *ServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerResponse.Merge(m, src)
}
func (m *ServerResponse) XXX_Size() int {
	return xxx_messageInfo_ServerResponse.Size(m)
}
func (m *ServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerResponse proto.InternalMessageInfo

func (m *ServerResponse) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type PresenceChannelCreatedRequest struct {
	Id                   *protocol.EntityId  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	GameAccountId        *protocol.EntityId  `protobuf:"bytes,3,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	AccountId            *protocol.EntityId  `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Host                 *protocol.ProcessId `protobuf:"bytes,5,opt,name=host" json:"host,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PresenceChannelCreatedRequest) Reset()         { *m = PresenceChannelCreatedRequest{} }
func (m *PresenceChannelCreatedRequest) String() string { return proto.CompactTextString(m) }
func (*PresenceChannelCreatedRequest) ProtoMessage()    {}
func (*PresenceChannelCreatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{4}
}

func (m *PresenceChannelCreatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceChannelCreatedRequest.Unmarshal(m, b)
}
func (m *PresenceChannelCreatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceChannelCreatedRequest.Marshal(b, m, deterministic)
}
func (m *PresenceChannelCreatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceChannelCreatedRequest.Merge(m, src)
}
func (m *PresenceChannelCreatedRequest) XXX_Size() int {
	return xxx_messageInfo_PresenceChannelCreatedRequest.Size(m)
}
func (m *PresenceChannelCreatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceChannelCreatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceChannelCreatedRequest proto.InternalMessageInfo

func (m *PresenceChannelCreatedRequest) GetId() *protocol.EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PresenceChannelCreatedRequest) GetGameAccountId() *protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *PresenceChannelCreatedRequest) GetAccountId() *protocol.EntityId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

// Deprecated: Do not use.
func (m *PresenceChannelCreatedRequest) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type GameAccountOnlineNotification struct {
	GameAccountId        *protocol.EntityId  `protobuf:"bytes,1,req,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	Host                 *protocol.ProcessId `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"` // Deprecated: Do not use.
	SessionId            *string             `protobuf:"bytes,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GameAccountOnlineNotification) Reset()         { *m = GameAccountOnlineNotification{} }
func (m *GameAccountOnlineNotification) String() string { return proto.CompactTextString(m) }
func (*GameAccountOnlineNotification) ProtoMessage()    {}
func (*GameAccountOnlineNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{5}
}

func (m *GameAccountOnlineNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameAccountOnlineNotification.Unmarshal(m, b)
}
func (m *GameAccountOnlineNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameAccountOnlineNotification.Marshal(b, m, deterministic)
}
func (m *GameAccountOnlineNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameAccountOnlineNotification.Merge(m, src)
}
func (m *GameAccountOnlineNotification) XXX_Size() int {
	return xxx_messageInfo_GameAccountOnlineNotification.Size(m)
}
func (m *GameAccountOnlineNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_GameAccountOnlineNotification.DiscardUnknown(m)
}

var xxx_messageInfo_GameAccountOnlineNotification proto.InternalMessageInfo

func (m *GameAccountOnlineNotification) GetGameAccountId() *protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

// Deprecated: Do not use.
func (m *GameAccountOnlineNotification) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *GameAccountOnlineNotification) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

type GameAccountOfflineNotification struct {
	GameAccountId        *protocol.EntityId  `protobuf:"bytes,1,req,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	Host                 *protocol.ProcessId `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"` // Deprecated: Do not use.
	SessionId            *string             `protobuf:"bytes,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GameAccountOfflineNotification) Reset()         { *m = GameAccountOfflineNotification{} }
func (m *GameAccountOfflineNotification) String() string { return proto.CompactTextString(m) }
func (*GameAccountOfflineNotification) ProtoMessage()    {}
func (*GameAccountOfflineNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{6}
}

func (m *GameAccountOfflineNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameAccountOfflineNotification.Unmarshal(m, b)
}
func (m *GameAccountOfflineNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameAccountOfflineNotification.Marshal(b, m, deterministic)
}
func (m *GameAccountOfflineNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameAccountOfflineNotification.Merge(m, src)
}
func (m *GameAccountOfflineNotification) XXX_Size() int {
	return xxx_messageInfo_GameAccountOfflineNotification.Size(m)
}
func (m *GameAccountOfflineNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_GameAccountOfflineNotification.DiscardUnknown(m)
}

var xxx_messageInfo_GameAccountOfflineNotification proto.InternalMessageInfo

func (m *GameAccountOfflineNotification) GetGameAccountId() *protocol.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

// Deprecated: Do not use.
func (m *GameAccountOfflineNotification) GetHost() *protocol.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *GameAccountOfflineNotification) GetSessionId() string {
	if m != nil && m.SessionId != nil {
		return *m.SessionId
	}
	return ""
}

type GetAllValuesForAttributeRequest struct {
	AttributeKey         *string            `protobuf:"bytes,1,opt,name=attribute_key,json=attributeKey" json:"attribute_key,omitempty"`
	AgentId              *protocol.EntityId `protobuf:"bytes,2,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Program              *uint32            `protobuf:"fixed32,5,opt,name=program" json:"program,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetAllValuesForAttributeRequest) Reset()         { *m = GetAllValuesForAttributeRequest{} }
func (m *GetAllValuesForAttributeRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllValuesForAttributeRequest) ProtoMessage()    {}
func (*GetAllValuesForAttributeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{7}
}

func (m *GetAllValuesForAttributeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllValuesForAttributeRequest.Unmarshal(m, b)
}
func (m *GetAllValuesForAttributeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllValuesForAttributeRequest.Marshal(b, m, deterministic)
}
func (m *GetAllValuesForAttributeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllValuesForAttributeRequest.Merge(m, src)
}
func (m *GetAllValuesForAttributeRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllValuesForAttributeRequest.Size(m)
}
func (m *GetAllValuesForAttributeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllValuesForAttributeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllValuesForAttributeRequest proto.InternalMessageInfo

func (m *GetAllValuesForAttributeRequest) GetAttributeKey() string {
	if m != nil && m.AttributeKey != nil {
		return *m.AttributeKey
	}
	return ""
}

func (m *GetAllValuesForAttributeRequest) GetAgentId() *protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GetAllValuesForAttributeRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

type GetAllValuesForAttributeResponse struct {
	AttributeValue       []*protocol.Variant `protobuf:"bytes,1,rep,name=attribute_value,json=attributeValue" json:"attribute_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetAllValuesForAttributeResponse) Reset()         { *m = GetAllValuesForAttributeResponse{} }
func (m *GetAllValuesForAttributeResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllValuesForAttributeResponse) ProtoMessage()    {}
func (*GetAllValuesForAttributeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{8}
}

func (m *GetAllValuesForAttributeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllValuesForAttributeResponse.Unmarshal(m, b)
}
func (m *GetAllValuesForAttributeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllValuesForAttributeResponse.Marshal(b, m, deterministic)
}
func (m *GetAllValuesForAttributeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllValuesForAttributeResponse.Merge(m, src)
}
func (m *GetAllValuesForAttributeResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllValuesForAttributeResponse.Size(m)
}
func (m *GetAllValuesForAttributeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllValuesForAttributeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllValuesForAttributeResponse proto.InternalMessageInfo

func (m *GetAllValuesForAttributeResponse) GetAttributeValue() []*protocol.Variant {
	if m != nil {
		return m.AttributeValue
	}
	return nil
}

type RegisterUtilitiesRequest struct {
	Attribute            []*protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Program              *uint32               `protobuf:"fixed32,2,opt,name=program" json:"program,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RegisterUtilitiesRequest) Reset()         { *m = RegisterUtilitiesRequest{} }
func (m *RegisterUtilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUtilitiesRequest) ProtoMessage()    {}
func (*RegisterUtilitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{9}
}

func (m *RegisterUtilitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUtilitiesRequest.Unmarshal(m, b)
}
func (m *RegisterUtilitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUtilitiesRequest.Marshal(b, m, deterministic)
}
func (m *RegisterUtilitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUtilitiesRequest.Merge(m, src)
}
func (m *RegisterUtilitiesRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUtilitiesRequest.Size(m)
}
func (m *RegisterUtilitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUtilitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUtilitiesRequest proto.InternalMessageInfo

func (m *RegisterUtilitiesRequest) GetAttribute() []*protocol.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *RegisterUtilitiesRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

type RegisterUtilitiesResponse struct {
	ClientId             *string  `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUtilitiesResponse) Reset()         { *m = RegisterUtilitiesResponse{} }
func (m *RegisterUtilitiesResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterUtilitiesResponse) ProtoMessage()    {}
func (*RegisterUtilitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{10}
}

func (m *RegisterUtilitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUtilitiesResponse.Unmarshal(m, b)
}
func (m *RegisterUtilitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUtilitiesResponse.Marshal(b, m, deterministic)
}
func (m *RegisterUtilitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUtilitiesResponse.Merge(m, src)
}
func (m *RegisterUtilitiesResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterUtilitiesResponse.Size(m)
}
func (m *RegisterUtilitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUtilitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUtilitiesResponse proto.InternalMessageInfo

func (m *RegisterUtilitiesResponse) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

type UnregisterUtilitiesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnregisterUtilitiesRequest) Reset()         { *m = UnregisterUtilitiesRequest{} }
func (m *UnregisterUtilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterUtilitiesRequest) ProtoMessage()    {}
func (*UnregisterUtilitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19892fdd2d88828c, []int{11}
}

func (m *UnregisterUtilitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnregisterUtilitiesRequest.Unmarshal(m, b)
}
func (m *UnregisterUtilitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnregisterUtilitiesRequest.Marshal(b, m, deterministic)
}
func (m *UnregisterUtilitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisterUtilitiesRequest.Merge(m, src)
}
func (m *UnregisterUtilitiesRequest) XXX_Size() int {
	return xxx_messageInfo_UnregisterUtilitiesRequest.Size(m)
}
func (m *UnregisterUtilitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisterUtilitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisterUtilitiesRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientRequest)(nil), "bgs.protocol.game_utilities.v1.ClientRequest")
	proto.RegisterType((*ClientResponse)(nil), "bgs.protocol.game_utilities.v1.ClientResponse")
	proto.RegisterType((*ServerRequest)(nil), "bgs.protocol.game_utilities.v1.ServerRequest")
	proto.RegisterType((*ServerResponse)(nil), "bgs.protocol.game_utilities.v1.ServerResponse")
	proto.RegisterType((*PresenceChannelCreatedRequest)(nil), "bgs.protocol.game_utilities.v1.PresenceChannelCreatedRequest")
	proto.RegisterType((*GameAccountOnlineNotification)(nil), "bgs.protocol.game_utilities.v1.GameAccountOnlineNotification")
	proto.RegisterType((*GameAccountOfflineNotification)(nil), "bgs.protocol.game_utilities.v1.GameAccountOfflineNotification")
	proto.RegisterType((*GetAllValuesForAttributeRequest)(nil), "bgs.protocol.game_utilities.v1.GetAllValuesForAttributeRequest")
	proto.RegisterType((*GetAllValuesForAttributeResponse)(nil), "bgs.protocol.game_utilities.v1.GetAllValuesForAttributeResponse")
	proto.RegisterType((*RegisterUtilitiesRequest)(nil), "bgs.protocol.game_utilities.v1.RegisterUtilitiesRequest")
	proto.RegisterType((*RegisterUtilitiesResponse)(nil), "bgs.protocol.game_utilities.v1.RegisterUtilitiesResponse")
	proto.RegisterType((*UnregisterUtilitiesRequest)(nil), "bgs.protocol.game_utilities.v1.UnregisterUtilitiesRequest")
}

func init() { proto.RegisterFile("game_utilities_service.proto", fileDescriptor_19892fdd2d88828c) }

var fileDescriptor_19892fdd2d88828c = []byte{
	// 860 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0x66, 0x36, 0x69, 0x62, 0xbf, 0x34, 0x09, 0x4c, 0xd3, 0xb2, 0x59, 0x9a, 0xd6, 0x32, 0x12,
	0xaa, 0x12, 0x61, 0x37, 0x95, 0x2a, 0x15, 0x04, 0x81, 0x24, 0x0d, 0xc1, 0xaa, 0x94, 0x44, 0x1b,
	0xb5, 0x07, 0x2e, 0xd6, 0x7a, 0xf7, 0xd9, 0x19, 0x75, 0x33, 0xb3, 0xcc, 0xcc, 0x1a, 0xf9, 0x80,
	0x84, 0x7a, 0xaa, 0x38, 0x81, 0xf8, 0x0d, 0x9c, 0xe1, 0xcf, 0xf0, 0x3b, 0x38, 0x87, 0x13, 0xf2,
	0x7a, 0xd7, 0xf6, 0x24, 0xf6, 0xae, 0x93, 0x72, 0xe0, 0x66, 0xcf, 0x7c, 0xef, 0x7d, 0xef, 0x7b,
	0xef, 0xcd, 0xb7, 0x70, 0xbf, 0xe3, 0x9d, 0x63, 0x33, 0xd6, 0x2c, 0x64, 0x9a, 0xa1, 0x6a, 0x2a,
	0x94, 0x5d, 0xe6, 0x63, 0x2d, 0x92, 0x42, 0x0b, 0xfa, 0xa0, 0xd5, 0x51, 0x83, 0x9f, 0xbe, 0x08,
	0x6b, 0x26, 0xb4, 0xd6, 0xdd, 0x76, 0xee, 0x7a, 0x5a, 0x4b, 0xd6, 0x8a, 0x35, 0x36, 0x75, 0x2f,
	0xc2, 0x14, 0xeb, 0x50, 0xe4, 0x9a, 0xe9, 0x9e, 0x71, 0xe6, 0x5c, 0x22, 0x1a, 0xbf, 0x5b, 0x95,
	0x91, 0x3f, 0x7e, 0x50, 0xfd, 0xcb, 0x82, 0xe5, 0xfd, 0x90, 0x21, 0xd7, 0x2e, 0x7e, 0x1f, 0xa3,
	0xd2, 0xf4, 0x29, 0x94, 0x87, 0x5c, 0x36, 0xa9, 0xcc, 0x3d, 0x5a, 0x7a, 0xf2, 0x61, 0xcd, 0xa8,
	0x6e, 0x37, 0xbb, 0x76, 0x47, 0x48, 0x5a, 0x87, 0xf9, 0x33, 0xa1, 0xb4, 0x6d, 0x55, 0xc8, 0xd5,
	0x88, 0x13, 0x29, 0x7c, 0x54, 0xaa, 0x11, 0xec, 0x59, 0x36, 0x71, 0x13, 0x20, 0x7d, 0x0a, 0xe0,
	0xf9, 0xbe, 0x88, 0xb9, 0x6e, 0xb2, 0xc0, 0x9e, 0x4b, 0xc2, 0xee, 0x99, 0x61, 0x07, 0x89, 0xb8,
	0x46, 0xe0, 0x96, 0x53, 0x64, 0x23, 0xa0, 0x3b, 0xb0, 0x9a, 0xe8, 0x1b, 0x8b, 0x9d, 0xcf, 0x8d,
	0x5d, 0xee, 0xc3, 0x77, 0x87, 0xf1, 0x36, 0x2c, 0x46, 0x52, 0x74, 0xa4, 0x77, 0x6e, 0xdf, 0xaa,
	0x90, 0x47, 0x8b, 0x6e, 0xf6, 0x97, 0xbe, 0x80, 0x25, 0x3f, 0xe9, 0x44, 0x93, 0xf1, 0xb6, 0xb0,
	0x17, 0x92, 0xac, 0x9b, 0xb5, 0xfc, 0xc1, 0xd4, 0x06, 0xcd, 0x6b, 0xf0, 0xb6, 0x70, 0xc1, 0x1f,
	0xfe, 0xae, 0x1e, 0xc2, 0x4a, 0xd6, 0x56, 0x15, 0x09, 0xae, 0xf0, 0x86, 0x7d, 0xad, 0xfe, 0x4a,
	0x60, 0xf9, 0x14, 0x65, 0x17, 0xe5, 0x3b, 0x0e, 0x68, 0x4c, 0xb8, 0x55, 0xb1, 0xc6, 0x85, 0x67,
	0xa3, 0x9b, 0x9b, 0x71, 0x74, 0x7d, 0x71, 0x59, 0x49, 0xef, 0x26, 0xee, 0x6f, 0x02, 0x1b, 0x27,
	0x12, 0x15, 0x72, 0x1f, 0xf7, 0xcf, 0x3c, 0xce, 0x31, 0xdc, 0x97, 0xe8, 0x69, 0x0c, 0x32, 0xb1,
	0x9f, 0x80, 0xc5, 0x02, 0x9b, 0x54, 0xac, 0x9c, 0x09, 0x5b, 0x6c, 0xe2, 0x5a, 0xcc, 0x5d, 0x67,
	0x2d, 0xcc, 0x6d, 0x9c, 0x9f, 0x75, 0x1b, 0xb3, 0xd6, 0xdd, 0x9a, 0xb5, 0x75, 0x7f, 0x10, 0xd8,
	0x38, 0x1c, 0x31, 0x1f, 0xf3, 0x90, 0x71, 0x3c, 0x12, 0x9a, 0xb5, 0x99, 0xef, 0x69, 0x26, 0xf8,
	0x24, 0x25, 0xf9, 0xf2, 0x2f, 0x29, 0xb9, 0xf6, 0x43, 0xdc, 0x00, 0x50, 0xa8, 0x14, 0x13, 0x3c,
	0xeb, 0x5a, 0xd9, 0x2d, 0xa7, 0x27, 0x8d, 0xa0, 0xfa, 0x27, 0x81, 0x07, 0xe3, 0x15, 0xb7, 0xdb,
	0xff, 0xfb, 0x92, 0x7f, 0x23, 0xf0, 0xf0, 0x10, 0xf5, 0x6e, 0x18, 0xbe, 0xf2, 0xc2, 0x18, 0xd5,
	0x37, 0x42, 0x8e, 0xd6, 0x2f, 0x5d, 0xac, 0x8f, 0x61, 0x79, 0x64, 0xa9, 0xaf, 0xb1, 0x67, 0x93,
	0x24, 0xcb, 0xed, 0xe1, 0xe1, 0x0b, 0xec, 0xd1, 0x6d, 0x28, 0x79, 0x1d, 0x1c, 0x28, 0xb2, 0x72,
	0x77, 0x62, 0x31, 0xc1, 0xe5, 0xf9, 0x4b, 0xb5, 0x05, 0x95, 0xe9, 0x45, 0xa5, 0xef, 0x68, 0x07,
	0x56, 0x47, 0x55, 0x75, 0xfb, 0xb8, 0xf4, 0x35, 0xdd, 0x35, 0x79, 0x5f, 0x79, 0x92, 0x79, 0x5c,
	0xbb, 0x2b, 0x43, 0x74, 0x92, 0xb4, 0xfa, 0x1a, 0x6c, 0x17, 0x3b, 0x4c, 0x69, 0x94, 0x2f, 0x33,
	0x97, 0xfa, 0x2f, 0x7d, 0xc3, 0x10, 0xf4, 0x0c, 0xd6, 0x27, 0x90, 0xa5, 0x4a, 0x3e, 0x82, 0x72,
	0xe6, 0xa6, 0x41, 0xda, 0xdb, 0x52, 0xea, 0x8f, 0x41, 0xf5, 0x3e, 0x38, 0x2f, 0xb9, 0x9c, 0x52,
	0xe8, 0x93, 0x7f, 0x4a, 0xb0, 0xd6, 0xdf, 0xb8, 0xe1, 0xc5, 0xe9, 0xe0, 0x53, 0x49, 0x7f, 0x84,
	0xb5, 0x74, 0x1b, 0xcc, 0x4f, 0xd6, 0xa7, 0xb3, 0x99, 0x74, 0x0a, 0x77, 0x6a, 0xb3, 0xc2, 0x07,
	0x52, 0xaa, 0x0b, 0x6f, 0x2e, 0xb6, 0xac, 0x12, 0xa1, 0x31, 0xdc, 0x9b, 0x6c, 0x56, 0xf4, 0xcb,
	0xa2, 0x8c, 0xb9, 0x26, 0xe7, 0xac, 0x99, 0xe1, 0x47, 0xe2, 0xb9, 0xa7, 0xbd, 0x94, 0xd6, 0x1a,
	0x53, 0x6d, 0x7e, 0x07, 0x0a, 0x55, 0x1b, 0xf0, 0x62, 0xd5, 0xa6, 0xa5, 0xa7, 0xf4, 0x0b, 0xf4,
	0x07, 0xb8, 0x73, 0xcc, 0xaf, 0x58, 0x56, 0xb1, 0xe4, 0x5c, 0x97, 0x73, 0xd6, 0x2f, 0x49, 0x3e,
	0x6e, 0xba, 0x07, 0xa7, 0x27, 0xc7, 0x47, 0xa7, 0x07, 0x29, 0xf1, 0x22, 0xed, 0xc1, 0x9a, 0x49,
	0x3c, 0x70, 0x1e, 0xba, 0x73, 0x1d, 0xe6, 0xab, 0x6e, 0x55, 0x4c, 0x5d, 0xa2, 0xbf, 0x13, 0xb0,
	0xa7, 0xbd, 0x55, 0xfa, 0x55, 0x21, 0x7f, 0xbe, 0xf5, 0x38, 0x5f, 0xdf, 0x3c, 0x81, 0x31, 0x1b,
	0xa0, 0xbf, 0x10, 0xf8, 0xe0, 0xca, 0x13, 0xa4, 0xcf, 0x8a, 0xf2, 0x4f, 0xb3, 0x08, 0xe7, 0xb3,
	0x1b, 0x44, 0x1a, 0x25, 0x2d, 0x51, 0x0d, 0x77, 0x26, 0x3c, 0x6d, 0xfa, 0x79, 0x51, 0xe6, 0xe9,
	0x7e, 0x50, 0x3c, 0xb0, 0xdb, 0xce, 0xb7, 0x6f, 0x2e, 0xb6, 0xbe, 0x80, 0xcd, 0x16, 0x47, 0x3d,
	0x95, 0xc4, 0xb0, 0x94, 0xcd, 0x15, 0xf3, 0xf6, 0xe7, 0x8b, 0xad, 0xf9, 0x12, 0x79, 0x9f, 0xec,
	0xbd, 0x25, 0xf0, 0x30, 0x37, 0x4d, 0x77, 0x7b, 0x6f, 0x7d, 0x92, 0x3b, 0x9d, 0xf4, 0xf1, 0xdf,
	0x3d, 0xef, 0x30, 0x7d, 0x16, 0xb7, 0x6a, 0xbe, 0x38, 0xaf, 0xab, 0x38, 0x42, 0x19, 0x3d, 0x7e,
	0xac, 0xeb, 0x1d, 0x11, 0x9d, 0xa1, 0xf4, 0xa5, 0xd7, 0xd6, 0xf5, 0x7e, 0xf2, 0x7a, 0xab, 0xa3,
	0xea, 0x19, 0x41, 0xdd, 0x24, 0xa8, 0x77, 0xb7, 0x7f, 0x22, 0xef, 0xbd, 0x25, 0xe4, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x3d, 0x2d, 0xd3, 0x32, 0x0c, 0x00, 0x00,
}
