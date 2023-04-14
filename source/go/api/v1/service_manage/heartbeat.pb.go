// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

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

type HeartbeatRecord struct {
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,json=instance_id,proto3" json:"instanceId,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Namespace            string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	LastHeartbeatSec     int64    `protobuf:"varint,6,opt,name=lastHeartbeatSec,json=last_heartbeat_sec,proto3" json:"lastHeartbeatSec,omitempty"`
	Exist                bool     `protobuf:"varint,7,opt,name=exist,proto3" json:"exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatRecord) Reset()         { *m = HeartbeatRecord{} }
func (m *HeartbeatRecord) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRecord) ProtoMessage()    {}
func (*HeartbeatRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{0}
}
func (m *HeartbeatRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRecord.Unmarshal(m, b)
}
func (m *HeartbeatRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRecord.Marshal(b, m, deterministic)
}
func (dst *HeartbeatRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRecord.Merge(dst, src)
}
func (m *HeartbeatRecord) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRecord.Size(m)
}
func (m *HeartbeatRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRecord.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRecord proto.InternalMessageInfo

func (m *HeartbeatRecord) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *HeartbeatRecord) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *HeartbeatRecord) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *HeartbeatRecord) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HeartbeatRecord) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HeartbeatRecord) GetLastHeartbeatSec() int64 {
	if m != nil {
		return m.LastHeartbeatSec
	}
	return 0
}

func (m *HeartbeatRecord) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

type Heartbeats struct {
	Heartbeats           []*Instance `protobuf:"bytes,1,rep,name=heartbeats,proto3" json:"heartbeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Heartbeats) Reset()         { *m = Heartbeats{} }
func (m *Heartbeats) String() string { return proto.CompactTextString(m) }
func (*Heartbeats) ProtoMessage()    {}
func (*Heartbeats) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{1}
}
func (m *Heartbeats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeats.Unmarshal(m, b)
}
func (m *Heartbeats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeats.Marshal(b, m, deterministic)
}
func (dst *Heartbeats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeats.Merge(dst, src)
}
func (m *Heartbeats) XXX_Size() int {
	return xxx_messageInfo_Heartbeats.Size(m)
}
func (m *Heartbeats) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeats.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeats proto.InternalMessageInfo

func (m *Heartbeats) GetHeartbeats() []*Instance {
	if m != nil {
		return m.Heartbeats
	}
	return nil
}

type GetHeartbeatsRequest struct {
	InstanceIds          []string `protobuf:"bytes,1,rep,name=instanceIds,json=instance_ids,proto3" json:"instanceIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHeartbeatsRequest) Reset()         { *m = GetHeartbeatsRequest{} }
func (m *GetHeartbeatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetHeartbeatsRequest) ProtoMessage()    {}
func (*GetHeartbeatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{2}
}
func (m *GetHeartbeatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeartbeatsRequest.Unmarshal(m, b)
}
func (m *GetHeartbeatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeartbeatsRequest.Marshal(b, m, deterministic)
}
func (dst *GetHeartbeatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeartbeatsRequest.Merge(dst, src)
}
func (m *GetHeartbeatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetHeartbeatsRequest.Size(m)
}
func (m *GetHeartbeatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeartbeatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeartbeatsRequest proto.InternalMessageInfo

func (m *GetHeartbeatsRequest) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

type GetHeartbeatsResponse struct {
	Records              []*HeartbeatRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetHeartbeatsResponse) Reset()         { *m = GetHeartbeatsResponse{} }
func (m *GetHeartbeatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetHeartbeatsResponse) ProtoMessage()    {}
func (*GetHeartbeatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{3}
}
func (m *GetHeartbeatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeartbeatsResponse.Unmarshal(m, b)
}
func (m *GetHeartbeatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeartbeatsResponse.Marshal(b, m, deterministic)
}
func (dst *GetHeartbeatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeartbeatsResponse.Merge(dst, src)
}
func (m *GetHeartbeatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetHeartbeatsResponse.Size(m)
}
func (m *GetHeartbeatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeartbeatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeartbeatsResponse proto.InternalMessageInfo

func (m *GetHeartbeatsResponse) GetRecords() []*HeartbeatRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type DelHeartbeatsRequest struct {
	InstanceIds          []string `protobuf:"bytes,1,rep,name=instanceIds,json=instance_ids,proto3" json:"instanceIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelHeartbeatsRequest) Reset()         { *m = DelHeartbeatsRequest{} }
func (m *DelHeartbeatsRequest) String() string { return proto.CompactTextString(m) }
func (*DelHeartbeatsRequest) ProtoMessage()    {}
func (*DelHeartbeatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{4}
}
func (m *DelHeartbeatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelHeartbeatsRequest.Unmarshal(m, b)
}
func (m *DelHeartbeatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelHeartbeatsRequest.Marshal(b, m, deterministic)
}
func (dst *DelHeartbeatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelHeartbeatsRequest.Merge(dst, src)
}
func (m *DelHeartbeatsRequest) XXX_Size() int {
	return xxx_messageInfo_DelHeartbeatsRequest.Size(m)
}
func (m *DelHeartbeatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelHeartbeatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelHeartbeatsRequest proto.InternalMessageInfo

func (m *DelHeartbeatsRequest) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

type DelHeartbeatsResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelHeartbeatsResponse) Reset()         { *m = DelHeartbeatsResponse{} }
func (m *DelHeartbeatsResponse) String() string { return proto.CompactTextString(m) }
func (*DelHeartbeatsResponse) ProtoMessage()    {}
func (*DelHeartbeatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_9b3d15db5d8b5def, []int{5}
}
func (m *DelHeartbeatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelHeartbeatsResponse.Unmarshal(m, b)
}
func (m *DelHeartbeatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelHeartbeatsResponse.Marshal(b, m, deterministic)
}
func (dst *DelHeartbeatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelHeartbeatsResponse.Merge(dst, src)
}
func (m *DelHeartbeatsResponse) XXX_Size() int {
	return xxx_messageInfo_DelHeartbeatsResponse.Size(m)
}
func (m *DelHeartbeatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DelHeartbeatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DelHeartbeatsResponse proto.InternalMessageInfo

func (m *DelHeartbeatsResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DelHeartbeatsResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartbeatRecord)(nil), "v1.HeartbeatRecord")
	proto.RegisterType((*Heartbeats)(nil), "v1.Heartbeats")
	proto.RegisterType((*GetHeartbeatsRequest)(nil), "v1.GetHeartbeatsRequest")
	proto.RegisterType((*GetHeartbeatsResponse)(nil), "v1.GetHeartbeatsResponse")
	proto.RegisterType((*DelHeartbeatsRequest)(nil), "v1.DelHeartbeatsRequest")
	proto.RegisterType((*DelHeartbeatsResponse)(nil), "v1.DelHeartbeatsResponse")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_heartbeat_9b3d15db5d8b5def) }

var fileDescriptor_heartbeat_9b3d15db5d8b5def = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0x69, 0x43, 0x26, 0x8d, 0x8a, 0x4c, 0x2b, 0x59, 0x08, 0x89, 0x25, 0xa7, 0x3d,
	0x14, 0xaf, 0x02, 0x07, 0x04, 0x17, 0x24, 0x54, 0x01, 0xbd, 0x9a, 0x1b, 0x97, 0xc8, 0x71, 0xa6,
	0x89, 0xa5, 0xac, 0x6d, 0x3c, 0xce, 0x8a, 0x1b, 0xff, 0x92, 0xdf, 0x83, 0xd6, 0xfb, 0x91, 0x92,
	0x63, 0x6f, 0xe3, 0x37, 0x7e, 0x6f, 0xde, 0x3c, 0x1b, 0xae, 0x76, 0xa8, 0x42, 0x5c, 0xa3, 0x8a,
	0xc2, 0x07, 0x17, 0x1d, 0x3b, 0xab, 0x97, 0x2f, 0xe7, 0x84, 0xa1, 0x36, 0x1a, 0x5b, 0x68, 0xf1,
	0x37, 0x83, 0xab, 0xef, 0xfd, 0x35, 0x89, 0xda, 0x85, 0x0d, 0x7b, 0x0d, 0x60, 0x2c, 0x45, 0x65,
	0x35, 0xde, 0x6f, 0x78, 0x96, 0x67, 0xc5, 0x54, 0xce, 0x7a, 0x64, 0x65, 0x36, 0x8c, 0xc3, 0xa4,
	0x53, 0xe1, 0x67, 0xa9, 0xdb, 0x1f, 0xd9, 0x2b, 0x98, 0x5a, 0x55, 0x21, 0x79, 0xa5, 0x91, 0x8f,
	0x52, 0xef, 0x08, 0x30, 0x06, 0xe3, 0x9d, 0xa3, 0xc8, 0xc7, 0xa9, 0x91, 0xea, 0x06, 0xf3, 0x2e,
	0x44, 0x7e, 0x9e, 0x67, 0xc5, 0x5c, 0xa6, 0x9a, 0xdd, 0xc2, 0xf3, 0xbd, 0xa2, 0x38, 0xf8, 0xfa,
	0x81, 0x9a, 0x5f, 0xe4, 0x59, 0x31, 0x92, 0xac, 0xc1, 0x57, 0xc3, 0x5e, 0x2b, 0x42, 0xcd, 0xae,
	0xe1, 0x1c, 0x7f, 0x1b, 0x8a, 0x7c, 0x92, 0x67, 0xc5, 0x33, 0xd9, 0x1e, 0x16, 0x9f, 0x00, 0x06,
	0x3e, 0xb1, 0x5b, 0x80, 0x81, 0x44, 0x3c, 0xcb, 0x47, 0xc5, 0xec, 0xdd, 0xa5, 0xa8, 0x97, 0xe2,
	0xbe, 0x5b, 0x4b, 0x3e, 0xea, 0x2f, 0x3e, 0xc2, 0xf5, 0x37, 0x3c, 0x8e, 0x27, 0x89, 0xbf, 0x0e,
	0x48, 0x91, 0xbd, 0x81, 0xd9, 0x31, 0x98, 0x56, 0x66, 0x2a, 0x2f, 0x1f, 0x25, 0x43, 0x8b, 0xaf,
	0x70, 0x73, 0x42, 0x25, 0xef, 0x2c, 0x21, 0x7b, 0x0b, 0x93, 0x90, 0xe2, 0xed, 0xc7, 0xbf, 0x68,
	0xc6, 0x9f, 0x44, 0x2f, 0xfb, 0x3b, 0x8d, 0x85, 0x3b, 0xdc, 0x3f, 0xc9, 0xc2, 0x67, 0xb8, 0x39,
	0xa1, 0x76, 0x16, 0x18, 0x8c, 0xb5, 0xdb, 0x60, 0x7a, 0xd1, 0xb9, 0x4c, 0x75, 0x83, 0x19, 0xfb,
	0xe0, 0xba, 0x77, 0x4c, 0xf5, 0x97, 0x3f, 0xf0, 0x41, 0xbb, 0x4a, 0x44, 0xb4, 0x1a, 0x6d, 0x14,
	0xde, 0xed, 0x55, 0x30, 0x24, 0xc8, 0xa3, 0x36, 0x0f, 0x46, 0xab, 0x68, 0x9c, 0x15, 0xca, 0x9b,
	0x66, 0x81, 0xfe, 0x37, 0x55, 0xca, 0xaa, 0x2d, 0xfe, 0xbc, 0xdb, 0x9a, 0xb8, 0x3b, 0xac, 0x85,
	0x76, 0x55, 0xd9, 0xf1, 0x2a, 0xa4, 0x5d, 0xf9, 0x1f, 0xb7, 0x24, 0x77, 0x08, 0x1a, 0xcb, 0xad,
	0x2b, 0x95, 0x37, 0x65, 0xbd, 0x2c, 0x3b, 0x95, 0x55, 0xab, 0xb2, 0xbe, 0x48, 0x7f, 0xf3, 0xfd,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x03, 0xeb, 0x70, 0xc1, 0x02, 0x00, 0x00,
}
