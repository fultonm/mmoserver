// Code generated by protoc-gen-gcraft. DO NOT EDIT.
// source: global_extensions/field_options.proto

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

type LogOption int32

const (
	LogOption_HIDDEN LogOption = 1
	LogOption_HEX    LogOption = 2
)

var LogOption_name = map[int32]string{
	1: "HIDDEN",
	2: "HEX",
}

var LogOption_value = map[string]int32{
	"HIDDEN": 1,
	"HEX":    2,
}

func (x LogOption) Enum() *LogOption {
	p := new(LogOption)
	*p = x
	return p
}

func (x LogOption) String() string {
	return proto.EnumName(LogOption_name, int32(x))
}

func (x *LogOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LogOption_value, data, "LogOption")
	if err != nil {
		return err
	}
	*x = LogOption(value)
	return nil
}

func (LogOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{0}
}

type EntityIdRestriction_Kind int32

const (
	EntityIdRestriction_ANY                     EntityIdRestriction_Kind = 0
	EntityIdRestriction_ACCOUNT                 EntityIdRestriction_Kind = 1
	EntityIdRestriction_GAME_ACCOUNT            EntityIdRestriction_Kind = 2
	EntityIdRestriction_ACCOUNT_OR_GAME_ACCOUNT EntityIdRestriction_Kind = 3
	EntityIdRestriction_SERVICE                 EntityIdRestriction_Kind = 4
	EntityIdRestriction_CHANNEL                 EntityIdRestriction_Kind = 5
)

var EntityIdRestriction_Kind_name = map[int32]string{
	0: "ANY",
	1: "ACCOUNT",
	2: "GAME_ACCOUNT",
	3: "ACCOUNT_OR_GAME_ACCOUNT",
	4: "SERVICE",
	5: "CHANNEL",
}

var EntityIdRestriction_Kind_value = map[string]int32{
	"ANY":                     0,
	"ACCOUNT":                 1,
	"GAME_ACCOUNT":            2,
	"ACCOUNT_OR_GAME_ACCOUNT": 3,
	"SERVICE":                 4,
	"CHANNEL":                 5,
}

func (x EntityIdRestriction_Kind) Enum() *EntityIdRestriction_Kind {
	p := new(EntityIdRestriction_Kind)
	*p = x
	return p
}

func (x EntityIdRestriction_Kind) String() string {
	return proto.EnumName(EntityIdRestriction_Kind_name, int32(x))
}

func (x *EntityIdRestriction_Kind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntityIdRestriction_Kind_value, data, "EntityIdRestriction_Kind")
	if err != nil {
		return err
	}
	*x = EntityIdRestriction_Kind(value)
	return nil
}

func (EntityIdRestriction_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{7, 0}
}

type BGSFieldOptions struct {
	Log                  *LogOption `protobuf:"varint,1,opt,name=log,enum=bgs.protocol.LogOption" json:"log,omitempty"`
	ShardKey             *bool      `protobuf:"varint,2,opt,name=shard_key,json=shardKey" json:"shard_key,omitempty"`
	FanoutKey            *bool      `protobuf:"varint,3,opt,name=fanout_key,json=fanoutKey" json:"fanout_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BGSFieldOptions) Reset()         { *m = BGSFieldOptions{} }
func (m *BGSFieldOptions) String() string { return proto.CompactTextString(m) }
func (*BGSFieldOptions) ProtoMessage()    {}
func (*BGSFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{0}
}

func (m *BGSFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BGSFieldOptions.Unmarshal(m, b)
}
func (m *BGSFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BGSFieldOptions.Marshal(b, m, deterministic)
}
func (m *BGSFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BGSFieldOptions.Merge(m, src)
}
func (m *BGSFieldOptions) XXX_Size() int {
	return xxx_messageInfo_BGSFieldOptions.Size(m)
}
func (m *BGSFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_BGSFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_BGSFieldOptions proto.InternalMessageInfo

func (m *BGSFieldOptions) GetLog() LogOption {
	if m != nil && m.Log != nil {
		return *m.Log
	}
	return LogOption_HIDDEN
}

func (m *BGSFieldOptions) GetShardKey() bool {
	if m != nil && m.ShardKey != nil {
		return *m.ShardKey
	}
	return false
}

func (m *BGSFieldOptions) GetFanoutKey() bool {
	if m != nil && m.FanoutKey != nil {
		return *m.FanoutKey
	}
	return false
}

type FieldRestriction struct {
	// Types that are valid to be assigned to Type:
	//	*FieldRestriction_Signed
	//	*FieldRestriction_Unsigned
	//	*FieldRestriction_Float
	//	*FieldRestriction_String_
	//	*FieldRestriction_Repeated
	//	*FieldRestriction_Message
	//	*FieldRestriction_EntityId
	Type                 isFieldRestriction_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FieldRestriction) Reset()         { *m = FieldRestriction{} }
func (m *FieldRestriction) String() string { return proto.CompactTextString(m) }
func (*FieldRestriction) ProtoMessage()    {}
func (*FieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{1}
}

func (m *FieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldRestriction.Unmarshal(m, b)
}
func (m *FieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldRestriction.Marshal(b, m, deterministic)
}
func (m *FieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldRestriction.Merge(m, src)
}
func (m *FieldRestriction) XXX_Size() int {
	return xxx_messageInfo_FieldRestriction.Size(m)
}
func (m *FieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_FieldRestriction proto.InternalMessageInfo

type isFieldRestriction_Type interface {
	isFieldRestriction_Type()
}

type FieldRestriction_Signed struct {
	Signed *SignedFieldRestriction `protobuf:"bytes,1,opt,name=signed,oneof"`
}

type FieldRestriction_Unsigned struct {
	Unsigned *UnsignedFieldRestriction `protobuf:"bytes,2,opt,name=unsigned,oneof"`
}

type FieldRestriction_Float struct {
	Float *FloatFieldRestriction `protobuf:"bytes,3,opt,name=float,oneof"`
}

type FieldRestriction_String_ struct {
	String_ *StringFieldRestriction `protobuf:"bytes,4,opt,name=string,oneof"`
}

type FieldRestriction_Repeated struct {
	Repeated *RepeatedFieldRestriction `protobuf:"bytes,5,opt,name=repeated,oneof"`
}

type FieldRestriction_Message struct {
	Message *MessageFieldRestriction `protobuf:"bytes,6,opt,name=message,oneof"`
}

type FieldRestriction_EntityId struct {
	EntityId *EntityIdRestriction `protobuf:"bytes,7,opt,name=entity_id,json=entityId,oneof"`
}

func (*FieldRestriction_Signed) isFieldRestriction_Type() {}

func (*FieldRestriction_Unsigned) isFieldRestriction_Type() {}

func (*FieldRestriction_Float) isFieldRestriction_Type() {}

func (*FieldRestriction_String_) isFieldRestriction_Type() {}

func (*FieldRestriction_Repeated) isFieldRestriction_Type() {}

func (*FieldRestriction_Message) isFieldRestriction_Type() {}

func (*FieldRestriction_EntityId) isFieldRestriction_Type() {}

func (m *FieldRestriction) GetType() isFieldRestriction_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *FieldRestriction) GetSigned() *SignedFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_Signed); ok {
		return x.Signed
	}
	return nil
}

func (m *FieldRestriction) GetUnsigned() *UnsignedFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_Unsigned); ok {
		return x.Unsigned
	}
	return nil
}

func (m *FieldRestriction) GetFloat() *FloatFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_Float); ok {
		return x.Float
	}
	return nil
}

func (m *FieldRestriction) GetString_() *StringFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_String_); ok {
		return x.String_
	}
	return nil
}

func (m *FieldRestriction) GetRepeated() *RepeatedFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_Repeated); ok {
		return x.Repeated
	}
	return nil
}

func (m *FieldRestriction) GetMessage() *MessageFieldRestriction {
	if x, ok := m.GetType().(*FieldRestriction_Message); ok {
		return x.Message
	}
	return nil
}

func (m *FieldRestriction) GetEntityId() *EntityIdRestriction {
	if x, ok := m.GetType().(*FieldRestriction_EntityId); ok {
		return x.EntityId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldRestriction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldRestriction_Signed)(nil),
		(*FieldRestriction_Unsigned)(nil),
		(*FieldRestriction_Float)(nil),
		(*FieldRestriction_String_)(nil),
		(*FieldRestriction_Repeated)(nil),
		(*FieldRestriction_Message)(nil),
		(*FieldRestriction_EntityId)(nil),
	}
}

type RepeatedFieldRestriction struct {
	Size   *UnsignedIntRange `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
	Unique *bool             `protobuf:"varint,2,opt,name=unique" json:"unique,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*RepeatedFieldRestriction_Signed
	//	*RepeatedFieldRestriction_Unsigned
	//	*RepeatedFieldRestriction_Float
	//	*RepeatedFieldRestriction_String_
	//	*RepeatedFieldRestriction_EntityId
	Type                 isRepeatedFieldRestriction_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RepeatedFieldRestriction) Reset()         { *m = RepeatedFieldRestriction{} }
func (m *RepeatedFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*RepeatedFieldRestriction) ProtoMessage()    {}
func (*RepeatedFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{2}
}

func (m *RepeatedFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatedFieldRestriction.Unmarshal(m, b)
}
func (m *RepeatedFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatedFieldRestriction.Marshal(b, m, deterministic)
}
func (m *RepeatedFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatedFieldRestriction.Merge(m, src)
}
func (m *RepeatedFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_RepeatedFieldRestriction.Size(m)
}
func (m *RepeatedFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatedFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatedFieldRestriction proto.InternalMessageInfo

func (m *RepeatedFieldRestriction) GetSize() *UnsignedIntRange {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetUnique() bool {
	if m != nil && m.Unique != nil {
		return *m.Unique
	}
	return false
}

type isRepeatedFieldRestriction_Type interface {
	isRepeatedFieldRestriction_Type()
}

type RepeatedFieldRestriction_Signed struct {
	Signed *SignedFieldRestriction `protobuf:"bytes,3,opt,name=signed,oneof"`
}

type RepeatedFieldRestriction_Unsigned struct {
	Unsigned *UnsignedFieldRestriction `protobuf:"bytes,4,opt,name=unsigned,oneof"`
}

type RepeatedFieldRestriction_Float struct {
	Float *FloatFieldRestriction `protobuf:"bytes,5,opt,name=float,oneof"`
}

type RepeatedFieldRestriction_String_ struct {
	String_ *StringFieldRestriction `protobuf:"bytes,6,opt,name=string,oneof"`
}

type RepeatedFieldRestriction_EntityId struct {
	EntityId *EntityIdRestriction `protobuf:"bytes,7,opt,name=entity_id,json=entityId,oneof"`
}

func (*RepeatedFieldRestriction_Signed) isRepeatedFieldRestriction_Type() {}

func (*RepeatedFieldRestriction_Unsigned) isRepeatedFieldRestriction_Type() {}

func (*RepeatedFieldRestriction_Float) isRepeatedFieldRestriction_Type() {}

func (*RepeatedFieldRestriction_String_) isRepeatedFieldRestriction_Type() {}

func (*RepeatedFieldRestriction_EntityId) isRepeatedFieldRestriction_Type() {}

func (m *RepeatedFieldRestriction) GetType() isRepeatedFieldRestriction_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetSigned() *SignedFieldRestriction {
	if x, ok := m.GetType().(*RepeatedFieldRestriction_Signed); ok {
		return x.Signed
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetUnsigned() *UnsignedFieldRestriction {
	if x, ok := m.GetType().(*RepeatedFieldRestriction_Unsigned); ok {
		return x.Unsigned
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetFloat() *FloatFieldRestriction {
	if x, ok := m.GetType().(*RepeatedFieldRestriction_Float); ok {
		return x.Float
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetString_() *StringFieldRestriction {
	if x, ok := m.GetType().(*RepeatedFieldRestriction_String_); ok {
		return x.String_
	}
	return nil
}

func (m *RepeatedFieldRestriction) GetEntityId() *EntityIdRestriction {
	if x, ok := m.GetType().(*RepeatedFieldRestriction_EntityId); ok {
		return x.EntityId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RepeatedFieldRestriction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RepeatedFieldRestriction_Signed)(nil),
		(*RepeatedFieldRestriction_Unsigned)(nil),
		(*RepeatedFieldRestriction_Float)(nil),
		(*RepeatedFieldRestriction_String_)(nil),
		(*RepeatedFieldRestriction_EntityId)(nil),
	}
}

type SignedFieldRestriction struct {
	Limits               *SignedIntRange `protobuf:"bytes,1,opt,name=limits" json:"limits,omitempty"`
	Exclude              []int64         `protobuf:"zigzag64,2,rep,name=exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SignedFieldRestriction) Reset()         { *m = SignedFieldRestriction{} }
func (m *SignedFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*SignedFieldRestriction) ProtoMessage()    {}
func (*SignedFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{3}
}

func (m *SignedFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedFieldRestriction.Unmarshal(m, b)
}
func (m *SignedFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedFieldRestriction.Marshal(b, m, deterministic)
}
func (m *SignedFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedFieldRestriction.Merge(m, src)
}
func (m *SignedFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_SignedFieldRestriction.Size(m)
}
func (m *SignedFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedFieldRestriction proto.InternalMessageInfo

func (m *SignedFieldRestriction) GetLimits() *SignedIntRange {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *SignedFieldRestriction) GetExclude() []int64 {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type UnsignedFieldRestriction struct {
	Limits               *UnsignedIntRange `protobuf:"bytes,1,opt,name=limits" json:"limits,omitempty"`
	Exclude              []uint64          `protobuf:"varint,2,rep,name=exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UnsignedFieldRestriction) Reset()         { *m = UnsignedFieldRestriction{} }
func (m *UnsignedFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*UnsignedFieldRestriction) ProtoMessage()    {}
func (*UnsignedFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{4}
}

func (m *UnsignedFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsignedFieldRestriction.Unmarshal(m, b)
}
func (m *UnsignedFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsignedFieldRestriction.Marshal(b, m, deterministic)
}
func (m *UnsignedFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedFieldRestriction.Merge(m, src)
}
func (m *UnsignedFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_UnsignedFieldRestriction.Size(m)
}
func (m *UnsignedFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedFieldRestriction proto.InternalMessageInfo

func (m *UnsignedFieldRestriction) GetLimits() *UnsignedIntRange {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *UnsignedFieldRestriction) GetExclude() []uint64 {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type FloatFieldRestriction struct {
	Limits               *FloatRange `protobuf:"bytes,1,opt,name=limits" json:"limits,omitempty"`
	Exclude              []float32   `protobuf:"fixed32,2,rep,name=exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FloatFieldRestriction) Reset()         { *m = FloatFieldRestriction{} }
func (m *FloatFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*FloatFieldRestriction) ProtoMessage()    {}
func (*FloatFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{5}
}

func (m *FloatFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatFieldRestriction.Unmarshal(m, b)
}
func (m *FloatFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatFieldRestriction.Marshal(b, m, deterministic)
}
func (m *FloatFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatFieldRestriction.Merge(m, src)
}
func (m *FloatFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_FloatFieldRestriction.Size(m)
}
func (m *FloatFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_FloatFieldRestriction proto.InternalMessageInfo

func (m *FloatFieldRestriction) GetLimits() *FloatRange {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *FloatFieldRestriction) GetExclude() []float32 {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type StringFieldRestriction struct {
	Size                 *UnsignedIntRange `protobuf:"bytes,1,opt,name=size" json:"size,omitempty"`
	Exclude              []string          `protobuf:"bytes,2,rep,name=exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StringFieldRestriction) Reset()         { *m = StringFieldRestriction{} }
func (m *StringFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*StringFieldRestriction) ProtoMessage()    {}
func (*StringFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{6}
}

func (m *StringFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringFieldRestriction.Unmarshal(m, b)
}
func (m *StringFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringFieldRestriction.Marshal(b, m, deterministic)
}
func (m *StringFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringFieldRestriction.Merge(m, src)
}
func (m *StringFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_StringFieldRestriction.Size(m)
}
func (m *StringFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_StringFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_StringFieldRestriction proto.InternalMessageInfo

func (m *StringFieldRestriction) GetSize() *UnsignedIntRange {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *StringFieldRestriction) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type EntityIdRestriction struct {
	Needed               *bool                     `protobuf:"varint,1,opt,name=needed" json:"needed,omitempty"`
	Kind                 *EntityIdRestriction_Kind `protobuf:"varint,2,opt,name=kind,enum=bgs.protocol.EntityIdRestriction_Kind" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EntityIdRestriction) Reset()         { *m = EntityIdRestriction{} }
func (m *EntityIdRestriction) String() string { return proto.CompactTextString(m) }
func (*EntityIdRestriction) ProtoMessage()    {}
func (*EntityIdRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{7}
}

func (m *EntityIdRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityIdRestriction.Unmarshal(m, b)
}
func (m *EntityIdRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityIdRestriction.Marshal(b, m, deterministic)
}
func (m *EntityIdRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityIdRestriction.Merge(m, src)
}
func (m *EntityIdRestriction) XXX_Size() int {
	return xxx_messageInfo_EntityIdRestriction.Size(m)
}
func (m *EntityIdRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityIdRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_EntityIdRestriction proto.InternalMessageInfo

func (m *EntityIdRestriction) GetNeeded() bool {
	if m != nil && m.Needed != nil {
		return *m.Needed
	}
	return false
}

func (m *EntityIdRestriction) GetKind() EntityIdRestriction_Kind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return EntityIdRestriction_ANY
}

type MessageFieldRestriction struct {
	Needed               *bool    `protobuf:"varint,1,opt,name=needed" json:"needed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageFieldRestriction) Reset()         { *m = MessageFieldRestriction{} }
func (m *MessageFieldRestriction) String() string { return proto.CompactTextString(m) }
func (*MessageFieldRestriction) ProtoMessage()    {}
func (*MessageFieldRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaf0fd32d10c0881, []int{8}
}

func (m *MessageFieldRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageFieldRestriction.Unmarshal(m, b)
}
func (m *MessageFieldRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageFieldRestriction.Marshal(b, m, deterministic)
}
func (m *MessageFieldRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageFieldRestriction.Merge(m, src)
}
func (m *MessageFieldRestriction) XXX_Size() int {
	return xxx_messageInfo_MessageFieldRestriction.Size(m)
}
func (m *MessageFieldRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageFieldRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_MessageFieldRestriction proto.InternalMessageInfo

func (m *MessageFieldRestriction) GetNeeded() bool {
	if m != nil && m.Needed != nil {
		return *m.Needed
	}
	return false
}

var E_FieldOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*BGSFieldOptions)(nil),
	Field:         90000,
	Name:          "bgs.protocol.field_options",
	Tag:           "bytes,90000,opt,name=field_options",
	Filename:      "global_extensions/field_options.proto",
}

var E_Valid = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldRestriction)(nil),
	Field:         90001,
	Name:          "bgs.protocol.valid",
	Tag:           "bytes,90001,opt,name=valid",
	Filename:      "global_extensions/field_options.proto",
}

func init() {
	proto.RegisterEnum("bgs.protocol.LogOption", LogOption_name, LogOption_value)
	proto.RegisterEnum("bgs.protocol.EntityIdRestriction_Kind", EntityIdRestriction_Kind_name, EntityIdRestriction_Kind_value)
	proto.RegisterType((*BGSFieldOptions)(nil), "bgs.protocol.BGSFieldOptions")
	proto.RegisterType((*FieldRestriction)(nil), "bgs.protocol.FieldRestriction")
	proto.RegisterType((*RepeatedFieldRestriction)(nil), "bgs.protocol.RepeatedFieldRestriction")
	proto.RegisterType((*SignedFieldRestriction)(nil), "bgs.protocol.SignedFieldRestriction")
	proto.RegisterType((*UnsignedFieldRestriction)(nil), "bgs.protocol.UnsignedFieldRestriction")
	proto.RegisterType((*FloatFieldRestriction)(nil), "bgs.protocol.FloatFieldRestriction")
	proto.RegisterType((*StringFieldRestriction)(nil), "bgs.protocol.StringFieldRestriction")
	proto.RegisterType((*EntityIdRestriction)(nil), "bgs.protocol.EntityIdRestriction")
	proto.RegisterType((*MessageFieldRestriction)(nil), "bgs.protocol.MessageFieldRestriction")
	proto.RegisterExtension(E_FieldOptions)
	proto.RegisterExtension(E_Valid)
}

func init() {
	proto.RegisterFile("global_extensions/field_options.proto", fileDescriptor_eaf0fd32d10c0881)
}

var fileDescriptor_eaf0fd32d10c0881 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x6e, 0x62, 0xe7, 0x6f, 0xfa, 0x73, 0x7c, 0xf6, 0xe8, 0xb4, 0x56, 0x7b, 0x72, 0x94, 0x93,
	0x43, 0x51, 0xe1, 0xc2, 0x2e, 0x11, 0xe2, 0xa2, 0x48, 0x88, 0x24, 0x4d, 0x9b, 0xa8, 0x6d, 0x8a,
	0x5c, 0x5a, 0x01, 0x37, 0x91, 0x63, 0x6f, 0x9c, 0x55, 0x5d, 0xdb, 0xd8, 0x6b, 0xd4, 0xf0, 0x14,
	0xf0, 0x2c, 0x3c, 0x00, 0xaf, 0xc0, 0x63, 0xf0, 0x18, 0xc8, 0x6b, 0x27, 0x8d, 0xcd, 0x86, 0x4a,
	0x2d, 0x77, 0x9e, 0xf9, 0x66, 0xbe, 0x6f, 0x67, 0x67, 0xbc, 0x03, 0xdb, 0x96, 0xed, 0x0e, 0x75,
	0x7b, 0x80, 0xaf, 0x29, 0x76, 0x02, 0xe2, 0x3a, 0x81, 0x3a, 0x22, 0xd8, 0x36, 0x07, 0xae, 0x47,
	0x23, 0x4b, 0xf1, 0x7c, 0x97, 0xba, 0x68, 0x65, 0x68, 0x25, 0x9f, 0x86, 0x6b, 0x6f, 0xd6, 0x2c,
	0xd7, 0xb5, 0x6c, 0xac, 0x32, 0xc7, 0x30, 0x1c, 0xa9, 0x26, 0x0e, 0x0c, 0x9f, 0x78, 0xd4, 0xf5,
	0xe3, 0xa0, 0xcd, 0xea, 0xcf, 0xb4, 0xbe, 0xee, 0x58, 0x38, 0x86, 0xeb, 0xd7, 0xf0, 0x47, 0xeb,
	0xf0, 0xec, 0x20, 0x12, 0x3a, 0x8d, 0x75, 0xd0, 0x23, 0x10, 0x6c, 0xd7, 0x92, 0x73, 0xb5, 0xdc,
	0xce, 0x5a, 0x63, 0x43, 0x99, 0xd7, 0x53, 0x8e, 0x5d, 0x2b, 0x0e, 0xd3, 0xa2, 0x18, 0xb4, 0x05,
	0x95, 0x60, 0xac, 0xfb, 0xe6, 0xe0, 0x12, 0x4f, 0xe4, 0x7c, 0x2d, 0xb7, 0x53, 0xd6, 0xca, 0xcc,
	0x71, 0x84, 0x27, 0xa8, 0x0a, 0x30, 0xd2, 0x1d, 0x37, 0xa4, 0x0c, 0x15, 0x18, 0x5a, 0x89, 0x3d,
	0x47, 0x78, 0x52, 0xff, 0x2e, 0x80, 0xc4, 0x74, 0x35, 0x1c, 0x50, 0x9f, 0x18, 0x11, 0x2b, 0x7a,
	0x01, 0xc5, 0x80, 0x58, 0x0e, 0x36, 0x99, 0xfc, 0x72, 0xe3, 0x41, 0x5a, 0xfe, 0x8c, 0x61, 0xd9,
	0xac, 0xee, 0x92, 0x96, 0x64, 0xa1, 0x7d, 0x28, 0x87, 0x4e, 0xc2, 0x90, 0x67, 0x0c, 0x0f, 0xd3,
	0x0c, 0xe7, 0x09, 0xca, 0xe1, 0x98, 0x65, 0xa2, 0xe7, 0x50, 0x18, 0xd9, 0xae, 0x4e, 0xd9, 0xa1,
	0x97, 0x1b, 0xff, 0xa7, 0x29, 0x0e, 0x22, 0x88, 0x93, 0x1f, 0xe7, 0xb0, 0x12, 0xa8, 0x4f, 0x1c,
	0x4b, 0x16, 0xb9, 0x25, 0x30, 0x8c, 0x5b, 0x02, 0x43, 0xa2, 0x12, 0x7c, 0xec, 0x61, 0x9d, 0x62,
	0x53, 0x2e, 0xf0, 0x4a, 0xd0, 0x12, 0x94, 0x57, 0xc2, 0x34, 0x13, 0x35, 0xa1, 0x74, 0x85, 0x83,
	0x40, 0xb7, 0xb0, 0x5c, 0x64, 0x24, 0xdb, 0x69, 0x92, 0x93, 0x18, 0xe4, 0x70, 0x4c, 0xf3, 0xd0,
	0x4b, 0xa8, 0x60, 0x87, 0x12, 0x3a, 0x19, 0x10, 0x53, 0x2e, 0x31, 0x92, 0xff, 0xd2, 0x24, 0x1d,
	0x06, 0xf7, 0xb2, 0x87, 0xc0, 0x89, 0xbb, 0x55, 0x04, 0x91, 0x4e, 0x3c, 0x5c, 0xff, 0x22, 0x80,
	0xbc, 0xe8, 0xd4, 0xa8, 0x01, 0x62, 0x40, 0x3e, 0xe2, 0xa4, 0xe1, 0xff, 0xf2, 0xdb, 0xd5, 0x73,
	0xa8, 0x16, 0x4d, 0xad, 0xc6, 0x62, 0xd1, 0x3a, 0x14, 0x43, 0x87, 0xbc, 0x0f, 0x71, 0x32, 0x74,
	0x89, 0x35, 0x37, 0x3e, 0xc2, 0xbd, 0xc7, 0x47, 0xbc, 0xff, 0xf8, 0x14, 0xee, 0x35, 0x3e, 0xc5,
	0x3b, 0x8d, 0xcf, 0xef, 0xeb, 0xda, 0x18, 0xd6, 0xf9, 0x17, 0x86, 0x9e, 0x42, 0xd1, 0x26, 0x57,
	0x84, 0x06, 0x49, 0xd3, 0xfe, 0xe1, 0x5d, 0xf3, 0xac, 0x65, 0x49, 0x2c, 0x92, 0xa1, 0x84, 0xaf,
	0x0d, 0x3b, 0x34, 0xa3, 0xae, 0x09, 0x3b, 0x48, 0x9b, 0x9a, 0x75, 0x1b, 0xe4, 0x45, 0x17, 0x8b,
	0x9e, 0x65, 0xb4, 0x6e, 0x1b, 0x90, 0x05, 0x6a, 0xe2, 0x8d, 0x9a, 0x01, 0x7f, 0x73, 0x7b, 0x80,
	0x76, 0x33, 0x52, 0x32, 0xa7, 0x71, 0xbf, 0x14, 0xc9, 0xdf, 0x88, 0x8c, 0x60, 0x9d, 0xdf, 0xaa,
	0x3b, 0xcd, 0x7b, 0x46, 0xa7, 0x72, 0xa3, 0xf3, 0x2d, 0x07, 0x7f, 0x71, 0x1a, 0x1a, 0xfd, 0x21,
	0x0e, 0xc6, 0x66, 0xf2, 0x90, 0x96, 0xb5, 0xc4, 0x42, 0x7b, 0x20, 0x5e, 0x12, 0x27, 0x7e, 0x1c,
	0xd7, 0xb2, 0xd3, 0xcd, 0x21, 0x52, 0x8e, 0x88, 0x63, 0x6a, 0x2c, 0xa7, 0x8e, 0x41, 0x8c, 0x2c,
	0x54, 0x02, 0xa1, 0xd9, 0x7f, 0x2b, 0x2d, 0xa1, 0x65, 0x28, 0x35, 0xdb, 0xed, 0xd3, 0xf3, 0xfe,
	0x6b, 0x29, 0x87, 0x24, 0x58, 0x39, 0x6c, 0x9e, 0x74, 0x06, 0x53, 0x4f, 0x1e, 0x6d, 0xc1, 0x46,
	0x62, 0x0c, 0x4e, 0xb5, 0x41, 0x0a, 0x14, 0xa2, 0xdc, 0xb3, 0x8e, 0x76, 0xd1, 0x6b, 0x77, 0x24,
	0x31, 0x32, 0xda, 0xdd, 0x66, 0xbf, 0xdf, 0x39, 0x96, 0x0a, 0xf5, 0x27, 0xb0, 0xb1, 0xe0, 0x75,
	0x5a, 0x54, 0xd5, 0xe3, 0x1a, 0x54, 0x66, 0x9b, 0x09, 0x01, 0x14, 0xbb, 0xbd, 0xfd, 0xfd, 0x4e,
	0x5f, 0xca, 0x45, 0x47, 0xed, 0x76, 0xde, 0x48, 0xf9, 0x3d, 0x03, 0x56, 0x53, 0xdb, 0x14, 0x55,
	0x95, 0x78, 0x75, 0x2a, 0xd3, 0xd5, 0xa9, 0xcc, 0x2f, 0x41, 0xf9, 0xd3, 0xd7, 0xf8, 0xe7, 0xad,
	0xa6, 0x6f, 0x28, 0xb3, 0x2b, 0xb5, 0x95, 0xd1, 0x9c, 0xb5, 0x77, 0x0e, 0x85, 0x0f, 0xba, 0x4d,
	0xcc, 0xdb, 0xc8, 0x3f, 0x27, 0xe4, 0x99, 0xe6, 0x67, 0xcb, 0xd5, 0x62, 0xb6, 0xd6, 0x05, 0xac,
	0x0e, 0x1d, 0x4c, 0x67, 0x81, 0xad, 0x3f, 0xe7, 0xd9, 0x5e, 0x45, 0xde, 0x77, 0x0d, 0x8b, 0xd0,
	0x71, 0x38, 0x54, 0x0c, 0xf7, 0x4a, 0x0d, 0x42, 0x0f, 0xfb, 0xde, 0xee, 0x2e, 0x55, 0x2d, 0xd7,
	0x1b, 0x63, 0xdf, 0xf0, 0xf5, 0x11, 0x55, 0x23, 0x0a, 0x75, 0x68, 0x05, 0xea, 0x94, 0xe6, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xc5, 0x48, 0x38, 0x72, 0x08, 0x00, 0x00,
}
