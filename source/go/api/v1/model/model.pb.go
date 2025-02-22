// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model // import "github.com/polarismesh/specification/source/go/api/v1/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchString_MatchStringType int32

const (
	// Equivalent match
	MatchString_EXACT MatchString_MatchStringType = 0
	// Regular match
	MatchString_REGEX MatchString_MatchStringType = 1
	// Not equals match
	MatchString_NOT_EQUALS MatchString_MatchStringType = 2
	// Include match
	MatchString_IN MatchString_MatchStringType = 3
	// Not include match
	MatchString_NOT_IN MatchString_MatchStringType = 4
	// Range match
	MatchString_RANGE MatchString_MatchStringType = 5
)

var MatchString_MatchStringType_name = map[int32]string{
	0: "EXACT",
	1: "REGEX",
	2: "NOT_EQUALS",
	3: "IN",
	4: "NOT_IN",
	5: "RANGE",
}
var MatchString_MatchStringType_value = map[string]int32{
	"EXACT":      0,
	"REGEX":      1,
	"NOT_EQUALS": 2,
	"IN":         3,
	"NOT_IN":     4,
	"RANGE":      5,
}

func (x MatchString_MatchStringType) String() string {
	return proto.EnumName(MatchString_MatchStringType_name, int32(x))
}
func (MatchString_MatchStringType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{1, 0}
}

type MatchString_ValueType int32

const (
	MatchString_TEXT      MatchString_ValueType = 0
	MatchString_PARAMETER MatchString_ValueType = 1
	MatchString_VARIABLE  MatchString_ValueType = 2
)

var MatchString_ValueType_name = map[int32]string{
	0: "TEXT",
	1: "PARAMETER",
	2: "VARIABLE",
}
var MatchString_ValueType_value = map[string]int32{
	"TEXT":      0,
	"PARAMETER": 1,
	"VARIABLE":  2,
}

func (x MatchString_ValueType) String() string {
	return proto.EnumName(MatchString_ValueType_name, int32(x))
}
func (MatchString_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{1, 1}
}

type Location struct {
	Region               *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Campus               *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetRegion() *wrapperspb.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Location) GetZone() *wrapperspb.StringValue {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Location) GetCampus() *wrapperspb.StringValue {
	if m != nil {
		return m.Campus
	}
	return nil
}

type MatchString struct {
	Type                 MatchString_MatchStringType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.MatchString_MatchStringType" json:"type,omitempty"`
	Value                *wrapperspb.StringValue     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValueType            MatchString_ValueType       `protobuf:"varint,3,opt,name=value_type,proto3,enum=v1.MatchString_ValueType" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchString) Reset()         { *m = MatchString{} }
func (m *MatchString) String() string { return proto.CompactTextString(m) }
func (*MatchString) ProtoMessage()    {}
func (*MatchString) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{1}
}
func (m *MatchString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchString.Unmarshal(m, b)
}
func (m *MatchString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchString.Marshal(b, m, deterministic)
}
func (dst *MatchString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchString.Merge(dst, src)
}
func (m *MatchString) XXX_Size() int {
	return xxx_messageInfo_MatchString.Size(m)
}
func (m *MatchString) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchString.DiscardUnknown(m)
}

var xxx_messageInfo_MatchString proto.InternalMessageInfo

func (m *MatchString) GetType() MatchString_MatchStringType {
	if m != nil {
		return m.Type
	}
	return MatchString_EXACT
}

func (m *MatchString) GetValue() *wrapperspb.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MatchString) GetValueType() MatchString_ValueType {
	if m != nil {
		return m.ValueType
	}
	return MatchString_TEXT
}

type StringList struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{2}
}
func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (dst *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(dst, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// 汇总查询数据
type Summary struct {
	// 服务总数
	TotalServiceCount uint32 `protobuf:"varint,1,opt,name=total_service_count,proto3" json:"total_service_count,omitempty"`
	// 健康实例总数
	TotalHealthInstanceCount uint32 `protobuf:"varint,2,opt,name=total_health_instance_count,proto3" json:"total_health_instance_count,omitempty"`
	// 实例总数
	TotalInstanceCount   uint32   `protobuf:"varint,3,opt,name=total_instance_count,proto3" json:"total_instance_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{3}
}
func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (dst *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(dst, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetTotalServiceCount() uint32 {
	if m != nil {
		return m.TotalServiceCount
	}
	return 0
}

func (m *Summary) GetTotalHealthInstanceCount() uint32 {
	if m != nil {
		return m.TotalHealthInstanceCount
	}
	return 0
}

func (m *Summary) GetTotalInstanceCount() uint32 {
	if m != nil {
		return m.TotalInstanceCount
	}
	return 0
}

type ClientLabel struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *MatchString `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClientLabel) Reset()         { *m = ClientLabel{} }
func (m *ClientLabel) String() string { return proto.CompactTextString(m) }
func (*ClientLabel) ProtoMessage()    {}
func (*ClientLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_89bde6c9fe109cf9, []int{4}
}
func (m *ClientLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientLabel.Unmarshal(m, b)
}
func (m *ClientLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientLabel.Marshal(b, m, deterministic)
}
func (dst *ClientLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientLabel.Merge(dst, src)
}
func (m *ClientLabel) XXX_Size() int {
	return xxx_messageInfo_ClientLabel.Size(m)
}
func (m *ClientLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ClientLabel proto.InternalMessageInfo

func (m *ClientLabel) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ClientLabel) GetValue() *MatchString {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "v1.Location")
	proto.RegisterType((*MatchString)(nil), "v1.MatchString")
	proto.RegisterType((*StringList)(nil), "v1.StringList")
	proto.RegisterType((*Summary)(nil), "v1.Summary")
	proto.RegisterType((*ClientLabel)(nil), "v1.ClientLabel")
	proto.RegisterEnum("v1.MatchString_MatchStringType", MatchString_MatchStringType_name, MatchString_MatchStringType_value)
	proto.RegisterEnum("v1.MatchString_ValueType", MatchString_ValueType_name, MatchString_ValueType_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_89bde6c9fe109cf9) }

var fileDescriptor_model_89bde6c9fe109cf9 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x8f, 0xd2, 0x4c,
	0x14, 0xc6, 0xdf, 0xb6, 0xc0, 0x0b, 0x07, 0x77, 0xb7, 0x19, 0x8d, 0xc1, 0x3f, 0xd1, 0x4d, 0xa3,
	0xc9, 0x5e, 0x4d, 0x17, 0xd6, 0x1b, 0xe3, 0x8d, 0x65, 0x53, 0x37, 0x24, 0x80, 0x6b, 0x41, 0x42,
	0xbc, 0x21, 0x43, 0x9d, 0x2d, 0x13, 0xdb, 0xce, 0xa4, 0x33, 0xc5, 0xe0, 0x27, 0xf2, 0xd6, 0x0f,
	0xe0, 0x77, 0x33, 0x33, 0xc5, 0x0d, 0x4b, 0x36, 0x86, 0x2b, 0x0e, 0x73, 0x9e, 0xdf, 0x79, 0x9e,
	0xf4, 0xcc, 0x40, 0x3b, 0xe3, 0x5f, 0x69, 0x8a, 0x45, 0xc1, 0x15, 0x47, 0xf6, 0xba, 0xfb, 0xf4,
	0x45, 0xc2, 0x79, 0x92, 0x52, 0xdf, 0x9c, 0x2c, 0xcb, 0x1b, 0xff, 0x7b, 0x41, 0x84, 0xa0, 0x85,
	0xac, 0x34, 0xde, 0x4f, 0x0b, 0x9a, 0x43, 0x1e, 0x13, 0xc5, 0x78, 0x8e, 0xde, 0x40, 0xa3, 0xa0,
	0x09, 0xe3, 0x79, 0xc7, 0x3a, 0xb5, 0xce, 0xda, 0xbd, 0xe7, 0xb8, 0xa2, 0xf1, 0x5f, 0x1a, 0x4f,
	0x54, 0xc1, 0xf2, 0x64, 0x46, 0xd2, 0x92, 0x46, 0x5b, 0x2d, 0x3a, 0x87, 0xda, 0x0f, 0x9e, 0xd3,
	0x8e, 0x7d, 0x00, 0x63, 0x94, 0xda, 0x27, 0x26, 0x99, 0x28, 0x65, 0xc7, 0x39, 0xc4, 0xa7, 0xd2,
	0x7a, 0xbf, 0x6d, 0x68, 0x8f, 0x88, 0x8a, 0x57, 0x55, 0x13, 0x5d, 0x40, 0x4d, 0x6d, 0x04, 0x35,
	0x59, 0x8f, 0x7b, 0x2f, 0xf1, 0xba, 0x8b, 0x77, 0xda, 0xbb, 0xf5, 0x74, 0x23, 0x68, 0x64, 0xc4,
	0xa8, 0x07, 0xf5, 0xb5, 0x9e, 0x7a, 0x50, 0xda, 0x4a, 0x8a, 0xde, 0x02, 0x98, 0x62, 0x61, 0xec,
	0x1c, 0x63, 0xf7, 0x64, 0xdf, 0xce, 0x10, 0xc6, 0x68, 0x47, 0xec, 0xcd, 0xe0, 0x64, 0x2f, 0x07,
	0x6a, 0x41, 0x3d, 0x9c, 0x07, 0x97, 0x53, 0xf7, 0x3f, 0x5d, 0x46, 0xe1, 0x55, 0x38, 0x77, 0x2d,
	0x74, 0x0c, 0x30, 0xfe, 0x38, 0x5d, 0x84, 0x9f, 0x3e, 0x07, 0xc3, 0x89, 0x6b, 0xa3, 0x06, 0xd8,
	0x83, 0xb1, 0xeb, 0x20, 0x80, 0x86, 0x3e, 0x1f, 0x8c, 0xdd, 0x9a, 0x91, 0x07, 0xe3, 0xab, 0xd0,
	0xad, 0x7b, 0x3d, 0x68, 0xdd, 0x1a, 0xa2, 0x26, 0xd4, 0xa6, 0xe1, 0x5c, 0x0f, 0x3c, 0x82, 0xd6,
	0x75, 0x10, 0x05, 0xa3, 0x70, 0x1a, 0x46, 0xae, 0x85, 0x1e, 0x40, 0x73, 0x16, 0x44, 0x83, 0xa0,
	0x3f, 0x0c, 0x5d, 0xdb, 0x7b, 0x05, 0x50, 0xc5, 0x18, 0x32, 0xa9, 0xd0, 0x63, 0x68, 0x98, 0x9c,
	0xb2, 0x63, 0x9d, 0x3a, 0x67, 0xad, 0x68, 0xfb, 0xcf, 0xfb, 0x65, 0xc1, 0xff, 0x93, 0x32, 0xcb,
	0x48, 0xb1, 0x41, 0xe7, 0xf0, 0x50, 0x71, 0x45, 0xd2, 0x85, 0xa4, 0xc5, 0x9a, 0xc5, 0x74, 0x11,
	0xf3, 0x32, 0x57, 0xe6, 0x83, 0x1f, 0x45, 0xf7, 0xb5, 0xd0, 0x7b, 0x78, 0x56, 0x1d, 0xaf, 0x28,
	0x49, 0xd5, 0x6a, 0xc1, 0x72, 0xa9, 0x48, 0x7e, 0x4b, 0xda, 0x86, 0xfc, 0x97, 0x04, 0xf5, 0xe0,
	0x51, 0xd5, 0xde, 0x43, 0x1d, 0x83, 0xde, 0xdb, 0xf3, 0x3e, 0x40, 0xfb, 0x32, 0x65, 0x34, 0x57,
	0x43, 0xb2, 0xa4, 0x29, 0x72, 0xc1, 0xf9, 0x46, 0x37, 0x26, 0x66, 0x2b, 0xd2, 0x25, 0x7a, 0x7d,
	0x77, 0xeb, 0x27, 0x7b, 0xcb, 0xdb, 0x2e, 0xba, 0xbf, 0x01, 0x1c, 0xf3, 0x0c, 0x2b, 0x9a, 0xc7,
	0x34, 0x57, 0x58, 0xf0, 0x94, 0x14, 0x4c, 0x62, 0x29, 0x68, 0xcc, 0x6e, 0x58, 0xf5, 0x48, 0x30,
	0x11, 0x4c, 0xe3, 0xe6, 0xa1, 0xf5, 0x61, 0xa4, 0x7f, 0xae, 0xf5, 0xe5, 0xf9, 0xf2, 0x2e, 0x61,
	0x6a, 0x55, 0x2e, 0xf5, 0x08, 0x7f, 0x8b, 0x66, 0x54, 0xae, 0xfc, 0x3b, 0xb8, 0x2f, 0x79, 0x59,
	0xc4, 0xd4, 0x4f, 0xb8, 0x4f, 0x04, 0xf3, 0xd7, 0x5d, 0xdf, 0x0c, 0x5a, 0x36, 0xcc, 0x05, 0xbc,
	0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x33, 0xbc, 0xe4, 0x5b, 0xc1, 0x03, 0x00, 0x00,
}
