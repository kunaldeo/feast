// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/FeatureRow.proto

package types // import "github.com/gojektech/feast/go-feast-proto/feast/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureRow struct {
	EntityKey            string               `protobuf:"bytes,1,opt,name=entityKey,proto3" json:"entityKey,omitempty"`
	Features             []*Feature           `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	EventTimestamp       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	EntityName           string               `protobuf:"bytes,4,opt,name=entityName,proto3" json:"entityName,omitempty"`
	Granularity          Granularity_Enum     `protobuf:"varint,5,opt,name=granularity,proto3,enum=feast.types.Granularity_Enum" json:"granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureRow) Reset()         { *m = FeatureRow{} }
func (m *FeatureRow) String() string { return proto.CompactTextString(m) }
func (*FeatureRow) ProtoMessage()    {}
func (*FeatureRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureRow_6b06ea938f8ece4f, []int{0}
}
func (m *FeatureRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureRow.Unmarshal(m, b)
}
func (m *FeatureRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureRow.Marshal(b, m, deterministic)
}
func (dst *FeatureRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureRow.Merge(dst, src)
}
func (m *FeatureRow) XXX_Size() int {
	return xxx_messageInfo_FeatureRow.Size(m)
}
func (m *FeatureRow) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureRow.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureRow proto.InternalMessageInfo

func (m *FeatureRow) GetEntityKey() string {
	if m != nil {
		return m.EntityKey
	}
	return ""
}

func (m *FeatureRow) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureRow) GetEventTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.EventTimestamp
	}
	return nil
}

func (m *FeatureRow) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *FeatureRow) GetGranularity() Granularity_Enum {
	if m != nil {
		return m.Granularity
	}
	return Granularity_NONE
}

func init() {
	proto.RegisterType((*FeatureRow)(nil), "feast.types.FeatureRow")
}

func init() {
	proto.RegisterFile("feast/types/FeatureRow.proto", fileDescriptor_FeatureRow_6b06ea938f8ece4f)
}

var fileDescriptor_FeatureRow_6b06ea938f8ece4f = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xf4, 0x30,
	0x10, 0xc5, 0xe9, 0xee, 0xf7, 0x89, 0x9b, 0xc2, 0x0a, 0x41, 0xb0, 0x2e, 0xbb, 0x5a, 0x3c, 0xf5,
	0xb2, 0x89, 0x54, 0x3c, 0x0b, 0x05, 0xf5, 0x20, 0x88, 0x14, 0xf5, 0xe0, 0x2d, 0x5d, 0xa6, 0xd9,
	0x6a, 0xdb, 0x94, 0x76, 0xa2, 0xf4, 0xe8, 0x7f, 0x2e, 0xa6, 0xed, 0x36, 0x8a, 0xb7, 0x30, 0xbf,
	0x37, 0xc3, 0x7b, 0x2f, 0x64, 0x99, 0x82, 0x68, 0x90, 0x63, 0x5b, 0x41, 0xc3, 0x6f, 0x40, 0xa0,
	0xae, 0x21, 0x56, 0x1f, 0xac, 0xaa, 0x15, 0x2a, 0xea, 0x1a, 0xca, 0x0c, 0x5d, 0x9c, 0x4a, 0xa5,
	0x64, 0x0e, 0xdc, 0xa0, 0x44, 0xa7, 0x1c, 0xb3, 0x02, 0x1a, 0x14, 0x45, 0xd5, 0xa9, 0x17, 0xc7,
	0x7f, 0xdc, 0xea, 0xd1, 0xca, 0x46, 0xb7, 0xb5, 0x28, 0x75, 0x2e, 0xea, 0x0c, 0xdb, 0x1e, 0x1f,
	0xd9, 0xf8, 0x59, 0xe4, 0xba, 0xdf, 0x3b, 0xfb, 0x9c, 0x10, 0x32, 0xba, 0xa2, 0x4b, 0x32, 0x83,
	0x12, 0x33, 0x6c, 0xef, 0xa0, 0xf5, 0x1c, 0xdf, 0x09, 0x66, 0xf1, 0x38, 0xa0, 0xe7, 0x64, 0x3f,
	0xed, 0xb4, 0x8d, 0x37, 0xf1, 0xa7, 0x81, 0x1b, 0x1e, 0x32, 0x2b, 0x00, 0x1b, 0x0e, 0xed, 0x54,
	0x34, 0x22, 0x73, 0x78, 0x87, 0x12, 0x1f, 0x87, 0x24, 0xde, 0xd4, 0x77, 0x02, 0x37, 0x5c, 0xb0,
	0x2e, 0x2b, 0x1b, 0xb2, 0xb2, 0x9d, 0x22, 0xfe, 0xb5, 0x41, 0x4f, 0x08, 0xe9, 0x2c, 0xdc, 0x8b,
	0x02, 0xbc, 0x7f, 0xc6, 0x94, 0x35, 0xa1, 0x57, 0xc4, 0x95, 0x63, 0x60, 0xef, 0xbf, 0xef, 0x04,
	0xf3, 0x70, 0xf5, 0xc3, 0x98, 0x5d, 0xc8, 0x75, 0xa9, 0x8b, 0xd8, 0xde, 0x88, 0x9e, 0x88, 0xfd,
	0x0d, 0xd1, 0xc1, 0xd8, 0xc7, 0xc3, 0xb7, 0xbb, 0x97, 0x4b, 0x99, 0xe1, 0x56, 0x27, 0x6c, 0xa3,
	0x0a, 0x2e, 0xd5, 0x2b, 0xbc, 0x21, 0x6c, 0xb6, 0xbc, 0x6b, 0x54, 0xaa, 0xb5, 0x79, 0xac, 0x4d,
	0x10, 0x6e, 0xd5, 0x9c, 0xec, 0x99, 0xd1, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x18,
	0xb6, 0xac, 0x02, 0x02, 0x00, 0x00,
}