// Code generated by protoc-gen-go.
// source: qtypes.proto
// DO NOT EDIT!

/*
Package qtypes is a generated protocol buffer package.

It is generated from these files:
	qtypes.proto

It has these top-level messages:
	String
	Int64
	Uint64
	Float64
	Timestamp
*/
package qtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryType int32

const (
	QueryType_NULL             QueryType = 0
	QueryType_EQUAL            QueryType = 1
	QueryType_GREATER          QueryType = 2
	QueryType_GREATER_EQUAL    QueryType = 3
	QueryType_LESS             QueryType = 4
	QueryType_LESS_EQUAL       QueryType = 5
	QueryType_IN               QueryType = 6
	QueryType_BETWEEN          QueryType = 7
	QueryType_HAS_PREFIX       QueryType = 8
	QueryType_HAS_SUFFIX       QueryType = 9
	QueryType_SUBSTRING        QueryType = 10
	QueryType_PATTERN          QueryType = 11
	QueryType_MIN_LENGTH       QueryType = 12
	QueryType_MAX_LENGTH       QueryType = 13
	QueryType_OVERLAP          QueryType = 14
	QueryType_CONTAINS         QueryType = 15
	QueryType_IS_CONTAINED_BY  QueryType = 16
	QueryType_HAS_ELEMENT      QueryType = 17
	QueryType_HAS_ANY_ELEMENT  QueryType = 18
	QueryType_HAS_ALL_ELEMENTS QueryType = 19
)

var QueryType_name = map[int32]string{
	0:  "NULL",
	1:  "EQUAL",
	2:  "GREATER",
	3:  "GREATER_EQUAL",
	4:  "LESS",
	5:  "LESS_EQUAL",
	6:  "IN",
	7:  "BETWEEN",
	8:  "HAS_PREFIX",
	9:  "HAS_SUFFIX",
	10: "SUBSTRING",
	11: "PATTERN",
	12: "MIN_LENGTH",
	13: "MAX_LENGTH",
	14: "OVERLAP",
	15: "CONTAINS",
	16: "IS_CONTAINED_BY",
	17: "HAS_ELEMENT",
	18: "HAS_ANY_ELEMENT",
	19: "HAS_ALL_ELEMENTS",
}
var QueryType_value = map[string]int32{
	"NULL":             0,
	"EQUAL":            1,
	"GREATER":          2,
	"GREATER_EQUAL":    3,
	"LESS":             4,
	"LESS_EQUAL":       5,
	"IN":               6,
	"BETWEEN":          7,
	"HAS_PREFIX":       8,
	"HAS_SUFFIX":       9,
	"SUBSTRING":        10,
	"PATTERN":          11,
	"MIN_LENGTH":       12,
	"MAX_LENGTH":       13,
	"OVERLAP":          14,
	"CONTAINS":         15,
	"IS_CONTAINED_BY":  16,
	"HAS_ELEMENT":      17,
	"HAS_ANY_ELEMENT":  18,
	"HAS_ALL_ELEMENTS": 19,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// String ...
type String struct {
	Values      []string  `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Valid       bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation    bool      `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type        QueryType `protobuf:"varint,4,opt,name=type,enum=qtypes.QueryType" json:"type,omitempty"`
	Insensitive bool      `protobuf:"varint,5,opt,name=insensitive" json:"insensitive,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *String) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *String) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *String) GetNegation() bool {
	if m != nil {
		return m.Negation
	}
	return false
}

func (m *String) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_NULL
}

func (m *String) GetInsensitive() bool {
	if m != nil {
		return m.Insensitive
	}
	return false
}

// Int64 ...
type Int64 struct {
	Values   []int64   `protobuf:"varint,1,rep,packed,name=values" json:"values,omitempty"`
	Valid    bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool      `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     QueryType `protobuf:"varint,4,opt,name=type,enum=qtypes.QueryType" json:"type,omitempty"`
}

func (m *Int64) Reset()                    { *m = Int64{} }
func (m *Int64) String() string            { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()               {}
func (*Int64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Int64) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Int64) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Int64) GetNegation() bool {
	if m != nil {
		return m.Negation
	}
	return false
}

func (m *Int64) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_NULL
}

// Uint64 ...
type Uint64 struct {
	Values   []uint64  `protobuf:"varint,1,rep,packed,name=values" json:"values,omitempty"`
	Valid    bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool      `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     QueryType `protobuf:"varint,4,opt,name=type,enum=qtypes.QueryType" json:"type,omitempty"`
}

func (m *Uint64) Reset()                    { *m = Uint64{} }
func (m *Uint64) String() string            { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()               {}
func (*Uint64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Uint64) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Uint64) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Uint64) GetNegation() bool {
	if m != nil {
		return m.Negation
	}
	return false
}

func (m *Uint64) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_NULL
}

// Float64 ...
type Float64 struct {
	Values   []float64 `protobuf:"fixed64,1,rep,packed,name=values" json:"values,omitempty"`
	Valid    bool      `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool      `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     QueryType `protobuf:"varint,4,opt,name=type,enum=qtypes.QueryType" json:"type,omitempty"`
}

func (m *Float64) Reset()                    { *m = Float64{} }
func (m *Float64) String() string            { return proto.CompactTextString(m) }
func (*Float64) ProtoMessage()               {}
func (*Float64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Float64) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Float64) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Float64) GetNegation() bool {
	if m != nil {
		return m.Negation
	}
	return false
}

func (m *Float64) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_NULL
}

// Timestamp ...
type Timestamp struct {
	Values   []*google_protobuf.Timestamp `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Valid    bool                         `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool                         `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     QueryType                    `protobuf:"varint,4,opt,name=type,enum=qtypes.QueryType" json:"type,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Timestamp) GetValues() []*google_protobuf.Timestamp {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Timestamp) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Timestamp) GetNegation() bool {
	if m != nil {
		return m.Negation
	}
	return false
}

func (m *Timestamp) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_NULL
}

func init() {
	proto.RegisterType((*String)(nil), "qtypes.String")
	proto.RegisterType((*Int64)(nil), "qtypes.Int64")
	proto.RegisterType((*Uint64)(nil), "qtypes.Uint64")
	proto.RegisterType((*Float64)(nil), "qtypes.Float64")
	proto.RegisterType((*Timestamp)(nil), "qtypes.Timestamp")
	proto.RegisterEnum("qtypes.QueryType", QueryType_name, QueryType_value)
}

func init() { proto.RegisterFile("qtypes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcb, 0x6e, 0xda, 0x40,
	0x18, 0x85, 0x6b, 0x2e, 0x0e, 0xfe, 0xb9, 0x0d, 0x93, 0xa8, 0x42, 0x6c, 0x8a, 0x90, 0x2a, 0xa1,
	0x2e, 0x8c, 0x44, 0xab, 0xee, 0x4d, 0x3b, 0x10, 0x4b, 0xce, 0x84, 0x78, 0x86, 0x36, 0xe9, 0x06,
	0x99, 0xd4, 0xa5, 0xa3, 0x80, 0xed, 0xe2, 0x31, 0x2d, 0x79, 0x92, 0x4a, 0x7d, 0x90, 0xbe, 0x5e,
	0x35, 0xbe, 0xa0, 0xde, 0xb6, 0xec, 0xfc, 0x9d, 0xff, 0xcc, 0x9c, 0x23, 0xcf, 0x0f, 0x8d, 0x2f,
	0xf2, 0x10, 0xf9, 0xb1, 0x19, 0xed, 0x42, 0x19, 0x62, 0x3d, 0xa3, 0xde, 0xb3, 0x75, 0x18, 0xae,
	0x37, 0xfe, 0x28, 0x55, 0x57, 0xc9, 0xa7, 0x91, 0x14, 0x5b, 0x3f, 0x96, 0xde, 0x36, 0xca, 0x8c,
	0x83, 0x1f, 0x1a, 0xe8, 0x4c, 0xee, 0x44, 0xb0, 0xc6, 0x4f, 0x41, 0xdf, 0x7b, 0x9b, 0xc4, 0x8f,
	0xbb, 0x5a, 0xbf, 0x3c, 0x34, 0xdc, 0x9c, 0xf0, 0x05, 0x54, 0xf7, 0xde, 0x46, 0x7c, 0xec, 0x96,
	0xfa, 0xda, 0xb0, 0xe6, 0x66, 0x80, 0x7b, 0x50, 0x0b, 0xfc, 0xb5, 0x27, 0x45, 0x18, 0x74, 0xcb,
	0xe9, 0xe0, 0xc8, 0xf8, 0x39, 0x54, 0x54, 0x7c, 0xb7, 0xd2, 0xd7, 0x86, 0xad, 0x71, 0xc7, 0xcc,
	0xab, 0xdd, 0x24, 0xfe, 0xee, 0xc0, 0x0f, 0x91, 0xef, 0xa6, 0x63, 0xdc, 0x87, 0xba, 0x08, 0x62,
	0x3f, 0x88, 0x85, 0x14, 0x7b, 0xbf, 0x5b, 0x4d, 0x6f, 0xf9, 0x5d, 0x1a, 0x7c, 0x83, 0xaa, 0x1d,
	0xc8, 0xd7, 0xaf, 0xfe, 0xea, 0x56, 0x3e, 0x79, 0xb7, 0xc1, 0x01, 0xf4, 0x85, 0xf8, 0x4f, 0x74,
	0xe5, 0xf4, 0xd1, 0x8f, 0x70, 0x36, 0xdd, 0x84, 0xde, 0xbf, 0xd9, 0xda, 0xe9, 0xb3, 0xbf, 0x6b,
	0x60, 0xf0, 0x62, 0x45, 0xf0, 0xf8, 0x8f, 0xf8, 0xfa, 0xb8, 0x67, 0x66, 0xeb, 0x64, 0x16, 0xeb,
	0x64, 0x1e, 0xbd, 0x27, 0xaf, 0xf6, 0xe2, 0x67, 0x09, 0x8c, 0xa3, 0x86, 0x6b, 0x50, 0xa1, 0x0b,
	0xc7, 0x41, 0x4f, 0xb0, 0x01, 0x55, 0x72, 0xb3, 0xb0, 0x1c, 0xa4, 0xe1, 0x3a, 0x9c, 0xcd, 0x5c,
	0x62, 0x71, 0xe2, 0xa2, 0x12, 0xee, 0x40, 0x33, 0x87, 0x65, 0x36, 0x2f, 0xab, 0x43, 0x0e, 0x61,
	0x0c, 0x55, 0x70, 0x0b, 0x40, 0x7d, 0xe5, 0x93, 0x2a, 0xd6, 0xa1, 0x64, 0x53, 0xa4, 0xab, 0x1b,
	0x26, 0x84, 0xbf, 0x27, 0x84, 0xa2, 0x33, 0x65, 0xba, 0xb4, 0xd8, 0x72, 0xee, 0x92, 0xa9, 0x7d,
	0x8b, 0x6a, 0x05, 0xb3, 0xc5, 0x54, 0xb1, 0x81, 0x9b, 0x60, 0xb0, 0xc5, 0x84, 0x71, 0xd7, 0xa6,
	0x33, 0x04, 0xea, 0xec, 0xdc, 0xe2, 0x9c, 0xb8, 0x14, 0xd5, 0x95, 0xf7, 0xca, 0xa6, 0x4b, 0x87,
	0xd0, 0x19, 0xbf, 0x44, 0x8d, 0x94, 0xad, 0xdb, 0x82, 0x9b, 0xca, 0x7c, 0xfd, 0x8e, 0xb8, 0x8e,
	0x35, 0x47, 0x2d, 0xdc, 0x80, 0xda, 0x9b, 0x6b, 0xca, 0x2d, 0x9b, 0x32, 0xd4, 0xc6, 0xe7, 0xd0,
	0xb6, 0xd9, 0x32, 0x17, 0xc8, 0xdb, 0xe5, 0xe4, 0x0e, 0x21, 0xdc, 0x86, 0xba, 0xca, 0x26, 0x0e,
	0xb9, 0x22, 0x94, 0xa3, 0x8e, 0x72, 0x29, 0xc1, 0xa2, 0x77, 0x47, 0x11, 0xe3, 0x0b, 0x40, 0xa9,
	0xe8, 0x38, 0x85, 0xc8, 0xd0, 0xf9, 0x64, 0xf0, 0xa1, 0xbf, 0x16, 0xf2, 0x73, 0xb2, 0x32, 0xef,
	0xc3, 0xed, 0x28, 0x12, 0xa1, 0xdc, 0x3d, 0x84, 0x5f, 0xbd, 0xcd, 0xfd, 0x63, 0xf2, 0x30, 0xca,
	0xfe, 0xf6, 0x4a, 0x4f, 0x9f, 0xf4, 0xe5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x02, 0x9c,
	0x84, 0x47, 0x04, 0x00, 0x00,
}