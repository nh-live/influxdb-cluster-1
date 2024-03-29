// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meta.proto

package internal

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

type Data struct {
	Term                 uint64          `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	Index                uint64          `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	ClusterID            uint64          `protobuf:"varint,3,opt,name=ClusterID,proto3" json:"ClusterID,omitempty"`
	DataBases            []*DatabaseInfo `protobuf:"bytes,4,rep,name=DataBases,proto3" json:"DataBases,omitempty"`
	MetaNode             []*NodeInfo     `protobuf:"bytes,5,rep,name=MetaNode,proto3" json:"MetaNode,omitempty"`
	DataNode             []*NodeInfo     `protobuf:"bytes,6,rep,name=DataNode,proto3" json:"DataNode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Data) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Data) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *Data) GetDataBases() []*DatabaseInfo {
	if m != nil {
		return m.DataBases
	}
	return nil
}

func (m *Data) GetMetaNode() []*NodeInfo {
	if m != nil {
		return m.MetaNode
	}
	return nil
}

func (m *Data) GetDataNode() []*NodeInfo {
	if m != nil {
		return m.DataNode
	}
	return nil
}

type DatabaseInfo struct {
	Name                   string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	DefaultRetentionPolicy string                 `protobuf:"bytes,2,opt,name=DefaultRetentionPolicy,proto3" json:"DefaultRetentionPolicy,omitempty"`
	RetentionPolicies      []*RetentionPolicyInfo `protobuf:"bytes,3,rep,name=RetentionPolicies,proto3" json:"RetentionPolicies,omitempty"`
	ContinuousQueries      []*ContinuousQueryInfo `protobuf:"bytes,4,rep,name=ContinuousQueries,proto3" json:"ContinuousQueries,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *DatabaseInfo) Reset()         { *m = DatabaseInfo{} }
func (m *DatabaseInfo) String() string { return proto.CompactTextString(m) }
func (*DatabaseInfo) ProtoMessage()    {}
func (*DatabaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{1}
}

func (m *DatabaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseInfo.Unmarshal(m, b)
}
func (m *DatabaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseInfo.Marshal(b, m, deterministic)
}
func (m *DatabaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseInfo.Merge(m, src)
}
func (m *DatabaseInfo) XXX_Size() int {
	return xxx_messageInfo_DatabaseInfo.Size(m)
}
func (m *DatabaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseInfo proto.InternalMessageInfo

func (m *DatabaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatabaseInfo) GetDefaultRetentionPolicy() string {
	if m != nil {
		return m.DefaultRetentionPolicy
	}
	return ""
}

func (m *DatabaseInfo) GetRetentionPolicies() []*RetentionPolicyInfo {
	if m != nil {
		return m.RetentionPolicies
	}
	return nil
}

func (m *DatabaseInfo) GetContinuousQueries() []*ContinuousQueryInfo {
	if m != nil {
		return m.ContinuousQueries
	}
	return nil
}

type NodeInfo struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	TCPHost              string   `protobuf:"bytes,3,opt,name=TCPHost,proto3" json:"TCPHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeInfo) Reset()         { *m = NodeInfo{} }
func (m *NodeInfo) String() string { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()    {}
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{2}
}

func (m *NodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeInfo.Unmarshal(m, b)
}
func (m *NodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeInfo.Marshal(b, m, deterministic)
}
func (m *NodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeInfo.Merge(m, src)
}
func (m *NodeInfo) XXX_Size() int {
	return xxx_messageInfo_NodeInfo.Size(m)
}
func (m *NodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeInfo proto.InternalMessageInfo

func (m *NodeInfo) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *NodeInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *NodeInfo) GetTCPHost() string {
	if m != nil {
		return m.TCPHost
	}
	return ""
}

type RetentionPolicyInfo struct {
	Name                 string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ReplicaN             uint32            `protobuf:"varint,2,opt,name=ReplicaN,proto3" json:"ReplicaN,omitempty"`
	Duration             uint64            `protobuf:"varint,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
	ShardGroupDuration   uint64            `protobuf:"varint,4,opt,name=ShardGroupDuration,proto3" json:"ShardGroupDuration,omitempty"`
	ShardGroups          []*ShardGroupInfo `protobuf:"bytes,5,rep,name=ShardGroups,proto3" json:"ShardGroups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RetentionPolicyInfo) Reset()         { *m = RetentionPolicyInfo{} }
func (m *RetentionPolicyInfo) String() string { return proto.CompactTextString(m) }
func (*RetentionPolicyInfo) ProtoMessage()    {}
func (*RetentionPolicyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{3}
}

func (m *RetentionPolicyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetentionPolicyInfo.Unmarshal(m, b)
}
func (m *RetentionPolicyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetentionPolicyInfo.Marshal(b, m, deterministic)
}
func (m *RetentionPolicyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetentionPolicyInfo.Merge(m, src)
}
func (m *RetentionPolicyInfo) XXX_Size() int {
	return xxx_messageInfo_RetentionPolicyInfo.Size(m)
}
func (m *RetentionPolicyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RetentionPolicyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RetentionPolicyInfo proto.InternalMessageInfo

func (m *RetentionPolicyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RetentionPolicyInfo) GetReplicaN() uint32 {
	if m != nil {
		return m.ReplicaN
	}
	return 0
}

func (m *RetentionPolicyInfo) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *RetentionPolicyInfo) GetShardGroupDuration() uint64 {
	if m != nil {
		return m.ShardGroupDuration
	}
	return 0
}

func (m *RetentionPolicyInfo) GetShardGroups() []*ShardGroupInfo {
	if m != nil {
		return m.ShardGroups
	}
	return nil
}

type ShardGroupInfo struct {
	ID                   uint64       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StartTime            uint64       `protobuf:"varint,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              uint64       `protobuf:"varint,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	DeletedAt            uint64       `protobuf:"varint,4,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	Shards               []*ShardInfo `protobuf:"bytes,5,rep,name=Shards,proto3" json:"Shards,omitempty"`
	TruncatedAt          uint64       `protobuf:"varint,6,opt,name=TruncatedAt,proto3" json:"TruncatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ShardGroupInfo) Reset()         { *m = ShardGroupInfo{} }
func (m *ShardGroupInfo) String() string { return proto.CompactTextString(m) }
func (*ShardGroupInfo) ProtoMessage()    {}
func (*ShardGroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{4}
}

func (m *ShardGroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardGroupInfo.Unmarshal(m, b)
}
func (m *ShardGroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardGroupInfo.Marshal(b, m, deterministic)
}
func (m *ShardGroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardGroupInfo.Merge(m, src)
}
func (m *ShardGroupInfo) XXX_Size() int {
	return xxx_messageInfo_ShardGroupInfo.Size(m)
}
func (m *ShardGroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardGroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ShardGroupInfo proto.InternalMessageInfo

func (m *ShardGroupInfo) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ShardGroupInfo) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ShardGroupInfo) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *ShardGroupInfo) GetDeletedAt() uint64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func (m *ShardGroupInfo) GetShards() []*ShardInfo {
	if m != nil {
		return m.Shards
	}
	return nil
}

func (m *ShardGroupInfo) GetTruncatedAt() uint64 {
	if m != nil {
		return m.TruncatedAt
	}
	return 0
}

type ShardInfo struct {
	ID                   uint64        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Owners               []*ShardOwner `protobuf:"bytes,2,rep,name=Owners,proto3" json:"Owners,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShardInfo) Reset()         { *m = ShardInfo{} }
func (m *ShardInfo) String() string { return proto.CompactTextString(m) }
func (*ShardInfo) ProtoMessage()    {}
func (*ShardInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{5}
}

func (m *ShardInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardInfo.Unmarshal(m, b)
}
func (m *ShardInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardInfo.Marshal(b, m, deterministic)
}
func (m *ShardInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardInfo.Merge(m, src)
}
func (m *ShardInfo) XXX_Size() int {
	return xxx_messageInfo_ShardInfo.Size(m)
}
func (m *ShardInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ShardInfo proto.InternalMessageInfo

func (m *ShardInfo) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ShardInfo) GetOwners() []*ShardOwner {
	if m != nil {
		return m.Owners
	}
	return nil
}

type ShardOwner struct {
	NodeID               uint64   `protobuf:"varint,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardOwner) Reset()         { *m = ShardOwner{} }
func (m *ShardOwner) String() string { return proto.CompactTextString(m) }
func (*ShardOwner) ProtoMessage()    {}
func (*ShardOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{6}
}

func (m *ShardOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardOwner.Unmarshal(m, b)
}
func (m *ShardOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardOwner.Marshal(b, m, deterministic)
}
func (m *ShardOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardOwner.Merge(m, src)
}
func (m *ShardOwner) XXX_Size() int {
	return xxx_messageInfo_ShardOwner.Size(m)
}
func (m *ShardOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardOwner.DiscardUnknown(m)
}

var xxx_messageInfo_ShardOwner proto.InternalMessageInfo

func (m *ShardOwner) GetNodeID() uint64 {
	if m != nil {
		return m.NodeID
	}
	return 0
}

type ContinuousQueryInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=Query,proto3" json:"Query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContinuousQueryInfo) Reset()         { *m = ContinuousQueryInfo{} }
func (m *ContinuousQueryInfo) String() string { return proto.CompactTextString(m) }
func (*ContinuousQueryInfo) ProtoMessage()    {}
func (*ContinuousQueryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{7}
}

func (m *ContinuousQueryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousQueryInfo.Unmarshal(m, b)
}
func (m *ContinuousQueryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousQueryInfo.Marshal(b, m, deterministic)
}
func (m *ContinuousQueryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousQueryInfo.Merge(m, src)
}
func (m *ContinuousQueryInfo) XXX_Size() int {
	return xxx_messageInfo_ContinuousQueryInfo.Size(m)
}
func (m *ContinuousQueryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousQueryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousQueryInfo proto.InternalMessageInfo

func (m *ContinuousQueryInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContinuousQueryInfo) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "internal.Data")
	proto.RegisterType((*DatabaseInfo)(nil), "internal.DatabaseInfo")
	proto.RegisterType((*NodeInfo)(nil), "internal.NodeInfo")
	proto.RegisterType((*RetentionPolicyInfo)(nil), "internal.RetentionPolicyInfo")
	proto.RegisterType((*ShardGroupInfo)(nil), "internal.ShardGroupInfo")
	proto.RegisterType((*ShardInfo)(nil), "internal.ShardInfo")
	proto.RegisterType((*ShardOwner)(nil), "internal.ShardOwner")
	proto.RegisterType((*ContinuousQueryInfo)(nil), "internal.ContinuousQueryInfo")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0x6d, 0x36, 0x36, 0xa7, 0xba, 0xe0, 0xb4, 0x94, 0x20, 0x2b, 0x94, 0xe0, 0xc5, 0x82,
	0x92, 0x0b, 0x15, 0x2f, 0xbc, 0x11, 0x6d, 0xc4, 0x0d, 0x62, 0x5d, 0xb3, 0x7d, 0x81, 0xd9, 0xe6,
	0x2c, 0x06, 0xd2, 0x99, 0x32, 0x99, 0xa0, 0x3e, 0xa0, 0x97, 0xbe, 0x85, 0x8f, 0xe0, 0x03, 0xc8,
	0x9c, 0x4e, 0x32, 0xfd, 0xc9, 0xde, 0xcd, 0xf9, 0xce, 0xf7, 0x7d, 0xe7, 0xa7, 0xa7, 0x01, 0xd8,
	0xa0, 0xe6, 0xc9, 0x56, 0x49, 0x2d, 0xd9, 0xa8, 0x14, 0x1a, 0x95, 0xe0, 0x55, 0xfc, 0xd7, 0x03,
	0x3f, 0xe5, 0x9a, 0x33, 0x06, 0xfe, 0x0a, 0xd5, 0x26, 0xf2, 0xe6, 0xde, 0xa5, 0x9f, 0xd3, 0x9b,
	0x4d, 0xe1, 0x2c, 0x13, 0x05, 0xfe, 0x8c, 0x06, 0x04, 0xee, 0x02, 0x76, 0x01, 0xe1, 0xa2, 0x6a,
	0x6a, 0x8d, 0x2a, 0x4b, 0xa3, 0x21, 0x65, 0x1c, 0xc0, 0x5e, 0x43, 0x68, 0xfc, 0x3e, 0xf0, 0x1a,
	0xeb, 0xc8, 0x9f, 0x0f, 0x2f, 0xc7, 0x2f, 0x67, 0x49, 0x5b, 0x2e, 0x31, 0xa9, 0x5b, 0x5e, 0x63,
	0x26, 0xee, 0x64, 0xee, 0x88, 0x2c, 0x81, 0xd1, 0x17, 0xd4, 0x7c, 0x29, 0x0b, 0x8c, 0xce, 0x48,
	0xc4, 0x9c, 0xc8, 0xa0, 0x24, 0xe8, 0x38, 0x86, 0x6f, 0xc4, 0xc4, 0x0f, 0xee, 0xe7, 0xb7, 0x9c,
	0xf8, 0x9f, 0x07, 0x0f, 0xf7, 0x6b, 0x9b, 0x71, 0x97, 0x7c, 0x83, 0x34, 0x6e, 0x98, 0xd3, 0x9b,
	0xbd, 0x81, 0x59, 0x8a, 0x77, 0xbc, 0xa9, 0x74, 0x8e, 0x1a, 0x85, 0x2e, 0xa5, 0xb8, 0x96, 0x55,
	0xb9, 0xfe, 0x45, 0xf3, 0x87, 0xf9, 0x3d, 0x59, 0xf6, 0x19, 0x1e, 0x1f, 0x42, 0x25, 0xd6, 0xd1,
	0x90, 0xba, 0x7a, 0xea, 0xba, 0x3a, 0x52, 0x51, 0x83, 0xa7, 0x3a, 0x63, 0xb6, 0x90, 0x42, 0x97,
	0xa2, 0x91, 0x4d, 0xfd, 0xad, 0x41, 0x55, 0x76, 0x7b, 0xdc, 0x33, 0x3b, 0xa4, 0x58, 0xb3, 0x13,
	0x5d, 0x7c, 0x05, 0xa3, 0x76, 0x19, 0xec, 0x1c, 0x06, 0x59, 0x6a, 0x7f, 0xde, 0x41, 0x96, 0x9a,
	0x0d, 0x5c, 0xc9, 0x5a, 0xdb, 0xd9, 0xe8, 0xcd, 0x22, 0x78, 0xb0, 0x5a, 0x5c, 0x13, 0x3c, 0x24,
	0xb8, 0x0d, 0xe3, 0x3f, 0x1e, 0x4c, 0x7a, 0x26, 0xe8, 0xdd, 0xe3, 0x13, 0x18, 0xe5, 0xb8, 0xad,
	0xca, 0x35, 0x5f, 0x92, 0xfb, 0xa3, 0xbc, 0x8b, 0x4d, 0x2e, 0x6d, 0x14, 0x37, 0x2e, 0xf6, 0x76,
	0xba, 0x98, 0x25, 0xc0, 0x6e, 0xbe, 0x73, 0x55, 0x7c, 0x52, 0xb2, 0xd9, 0x76, 0x2c, 0x9f, 0x58,
	0x3d, 0x19, 0xf6, 0x16, 0xc6, 0x0e, 0xad, 0xed, 0xdd, 0x44, 0x6e, 0x49, 0x2e, 0x49, 0xfb, 0xd9,
	0x27, 0xc7, 0xbf, 0x3d, 0x38, 0x3f, 0xcc, 0x9f, 0x2c, 0xe8, 0x02, 0xc2, 0x1b, 0xcd, 0x95, 0x5e,
	0x95, 0x1b, 0xb4, 0xff, 0x00, 0x07, 0x98, 0x55, 0x7d, 0x14, 0x05, 0xe5, 0x76, 0x73, 0xb4, 0xa1,
	0xd1, 0xa5, 0x58, 0xa1, 0xc6, 0xe2, 0xbd, 0xb6, 0xdd, 0x3b, 0x80, 0x3d, 0x87, 0x80, 0xea, 0xb6,
	0xfd, 0x4e, 0x8e, 0xfa, 0xa5, 0x56, 0x2d, 0x85, 0xcd, 0x61, 0xbc, 0x52, 0x8d, 0x58, 0xf3, 0x9d,
	0x59, 0x40, 0x66, 0xfb, 0x50, 0x9c, 0x41, 0xd8, 0xc9, 0x4e, 0x26, 0x78, 0x01, 0xc1, 0xd7, 0x1f,
	0x02, 0x55, 0x1d, 0x0d, 0xa8, 0xd6, 0xf4, 0xa8, 0x16, 0x25, 0x73, 0xcb, 0x89, 0x9f, 0x01, 0x38,
	0x94, 0xcd, 0x20, 0xa0, 0xd3, 0x69, 0xfd, 0x6c, 0x14, 0xbf, 0x83, 0x49, 0xcf, 0xf1, 0xf5, 0xde,
	0xc1, 0x14, 0xce, 0x88, 0x60, 0x4f, 0x6c, 0x17, 0xdc, 0x06, 0xf4, 0x09, 0x7a, 0xf5, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x8f, 0x6f, 0x2d, 0x90, 0x04, 0x00, 0x00,
}
