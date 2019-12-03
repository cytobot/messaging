// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transport/discord/discord.proto

package msg_cyto_discord

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HealthCheckStatus struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	InstanceID           string               `protobuf:"bytes,2,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	ShardID              int32                `protobuf:"varint,3,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Uptime               int64                `protobuf:"varint,4,opt,name=uptime,proto3" json:"uptime,omitempty"`
	MemAllocated         int64                `protobuf:"varint,5,opt,name=memAllocated,proto3" json:"memAllocated,omitempty"`
	MemSystem            int64                `protobuf:"varint,6,opt,name=memSystem,proto3" json:"memSystem,omitempty"`
	MemCumulative        int64                `protobuf:"varint,7,opt,name=memCumulative,proto3" json:"memCumulative,omitempty"`
	TaskCount            int32                `protobuf:"varint,8,opt,name=taskCount,proto3" json:"taskCount,omitempty"`
	ConnectedServers     int32                `protobuf:"varint,9,opt,name=connectedServers,proto3" json:"connectedServers,omitempty"`
	ConnectedUsers       int32                `protobuf:"varint,10,opt,name=connectedUsers,proto3" json:"connectedUsers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HealthCheckStatus) Reset()         { *m = HealthCheckStatus{} }
func (m *HealthCheckStatus) String() string { return proto.CompactTextString(m) }
func (*HealthCheckStatus) ProtoMessage()    {}
func (*HealthCheckStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0531cb5b87159325, []int{0}
}

func (m *HealthCheckStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckStatus.Unmarshal(m, b)
}
func (m *HealthCheckStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckStatus.Marshal(b, m, deterministic)
}
func (m *HealthCheckStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckStatus.Merge(m, src)
}
func (m *HealthCheckStatus) XXX_Size() int {
	return xxx_messageInfo_HealthCheckStatus.Size(m)
}
func (m *HealthCheckStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckStatus proto.InternalMessageInfo

func (m *HealthCheckStatus) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *HealthCheckStatus) GetInstanceID() string {
	if m != nil {
		return m.InstanceID
	}
	return ""
}

func (m *HealthCheckStatus) GetShardID() int32 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *HealthCheckStatus) GetUptime() int64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *HealthCheckStatus) GetMemAllocated() int64 {
	if m != nil {
		return m.MemAllocated
	}
	return 0
}

func (m *HealthCheckStatus) GetMemSystem() int64 {
	if m != nil {
		return m.MemSystem
	}
	return 0
}

func (m *HealthCheckStatus) GetMemCumulative() int64 {
	if m != nil {
		return m.MemCumulative
	}
	return 0
}

func (m *HealthCheckStatus) GetTaskCount() int32 {
	if m != nil {
		return m.TaskCount
	}
	return 0
}

func (m *HealthCheckStatus) GetConnectedServers() int32 {
	if m != nil {
		return m.ConnectedServers
	}
	return 0
}

func (m *HealthCheckStatus) GetConnectedUsers() int32 {
	if m != nil {
		return m.ConnectedUsers
	}
	return 0
}

type DiscordInformationResponse struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	InstanceID           string               `protobuf:"bytes,2,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	Payload              map[string]string    `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DiscordInformationResponse) Reset()         { *m = DiscordInformationResponse{} }
func (m *DiscordInformationResponse) String() string { return proto.CompactTextString(m) }
func (*DiscordInformationResponse) ProtoMessage()    {}
func (*DiscordInformationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0531cb5b87159325, []int{1}
}

func (m *DiscordInformationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscordInformationResponse.Unmarshal(m, b)
}
func (m *DiscordInformationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscordInformationResponse.Marshal(b, m, deterministic)
}
func (m *DiscordInformationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscordInformationResponse.Merge(m, src)
}
func (m *DiscordInformationResponse) XXX_Size() int {
	return xxx_messageInfo_DiscordInformationResponse.Size(m)
}
func (m *DiscordInformationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscordInformationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscordInformationResponse proto.InternalMessageInfo

func (m *DiscordInformationResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DiscordInformationResponse) GetInstanceID() string {
	if m != nil {
		return m.InstanceID
	}
	return ""
}

func (m *DiscordInformationResponse) GetPayload() map[string]string {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheckStatus)(nil), "msg.cyto.discord.HealthCheckStatus")
	proto.RegisterType((*DiscordInformationResponse)(nil), "msg.cyto.discord.DiscordInformationResponse")
	proto.RegisterMapType((map[string]string)(nil), "msg.cyto.discord.DiscordInformationResponse.PayloadEntry")
}

func init() { proto.RegisterFile("transport/discord/discord.proto", fileDescriptor_0531cb5b87159325) }

var fileDescriptor_0531cb5b87159325 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0x63, 0x5b, 0xf3, 0x76, 0x95, 0x3a, 0x88, 0x0c, 0x45, 0xdc, 0x50, 0x44, 0x82,
	0x87, 0x29, 0xac, 0x97, 0x75, 0x6f, 0xd2, 0x0a, 0xf6, 0x26, 0x89, 0x7e, 0x80, 0xd9, 0xe4, 0x6d,
	0x1b, 0x9a, 0x99, 0x09, 0x33, 0x2f, 0x85, 0x7c, 0x02, 0x3f, 0xb3, 0x37, 0xc9, 0xa4, 0xd9, 0xee,
	0x56, 0x3c, 0x7a, 0x4a, 0xde, 0xef, 0xfd, 0xe6, 0xcf, 0x3f, 0x4c, 0xe0, 0x8a, 0xac, 0xd4, 0xae,
	0x36, 0x96, 0x96, 0x45, 0xe9, 0x72, 0x63, 0x8b, 0xe1, 0x29, 0x6a, 0x6b, 0xc8, 0xb0, 0x99, 0x72,
	0x5b, 0x91, 0xb7, 0x64, 0xc4, 0x91, 0xcf, 0xaf, 0xb6, 0xc6, 0x6c, 0x2b, 0x5c, 0xfa, 0xfd, 0x5d,
	0x73, 0xbf, 0xa4, 0x52, 0xa1, 0x23, 0xa9, 0xea, 0xfe, 0xc8, 0xe2, 0xf7, 0x08, 0x5e, 0x7d, 0x43,
	0x59, 0xd1, 0x6e, 0xb5, 0xc3, 0x7c, 0x9f, 0x91, 0xa4, 0xc6, 0xb1, 0x1b, 0x88, 0x1e, 0x44, 0x1e,
	0xc4, 0x41, 0x72, 0x71, 0x3d, 0x17, 0x7d, 0x94, 0x18, 0xa2, 0xc4, 0x8f, 0xc1, 0x48, 0x4f, 0x32,
	0x7b, 0x07, 0x50, 0x6a, 0x47, 0x52, 0xe7, 0xb8, 0x59, 0xf3, 0x51, 0x1c, 0x24, 0x51, 0xfa, 0x88,
	0x30, 0x0e, 0x53, 0xb7, 0x93, 0xb6, 0xd8, 0xac, 0x79, 0x18, 0x07, 0xc9, 0x38, 0x1d, 0x46, 0xf6,
	0x06, 0x26, 0x4d, 0xdd, 0x05, 0xf1, 0x67, 0x71, 0x90, 0x84, 0xe9, 0x71, 0x62, 0x0b, 0xb8, 0x54,
	0xa8, 0xbe, 0x54, 0x95, 0xc9, 0x25, 0x61, 0xc1, 0xc7, 0x7e, 0xfb, 0x84, 0xb1, 0xb7, 0x10, 0x29,
	0x54, 0x59, 0xeb, 0x08, 0x15, 0x9f, 0x78, 0xe1, 0x04, 0xd8, 0x7b, 0x78, 0xa1, 0x50, 0xad, 0x1a,
	0xd5, 0x54, 0x92, 0xca, 0x03, 0xf2, 0xa9, 0x37, 0x9e, 0xc2, 0x2e, 0x83, 0xa4, 0xdb, 0xaf, 0x4c,
	0xa3, 0x89, 0x3f, 0xf7, 0xdd, 0x4e, 0x80, 0x7d, 0x84, 0x59, 0x6e, 0xb4, 0xc6, 0x9c, 0xb0, 0xc8,
	0xd0, 0x1e, 0xd0, 0x3a, 0x1e, 0x79, 0xe9, 0x2f, 0xce, 0x3e, 0xc0, 0xcb, 0x07, 0xf6, 0xd3, 0x75,
	0x26, 0x78, 0xf3, 0x8c, 0x2e, 0x7e, 0x8d, 0x60, 0xbe, 0xee, 0x2f, 0x6a, 0xa3, 0xef, 0x8d, 0x55,
	0x92, 0x4a, 0xa3, 0x53, 0x74, 0xb5, 0xd1, 0x0e, 0xff, 0xe3, 0x25, 0x64, 0x30, 0xad, 0x65, 0x5b,
	0x19, 0x59, 0xf0, 0x30, 0x0e, 0x93, 0x8b, 0xeb, 0xcf, 0xe2, 0xfc, 0xcf, 0x11, 0xff, 0x2e, 0x26,
	0xbe, 0xf7, 0x67, 0xbf, 0x6a, 0xb2, 0x6d, 0x3a, 0x24, 0xcd, 0x6f, 0xe1, 0xf2, 0xf1, 0x82, 0xcd,
	0x20, 0xdc, 0x63, 0xeb, 0x8b, 0x47, 0x69, 0xf7, 0xca, 0x5e, 0xc3, 0xf8, 0x20, 0xab, 0x06, 0x8f,
	0x8d, 0xfa, 0xe1, 0x76, 0x74, 0x13, 0xdc, 0x4d, 0xfc, 0xf7, 0x7c, 0xfa, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0x04, 0x30, 0xea, 0xe2, 0x02, 0x00, 0x00,
}
