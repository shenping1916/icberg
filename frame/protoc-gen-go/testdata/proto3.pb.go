// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3.proto

/*
Package proto3 is a generated protocol buffer package.

It is generated from these files:
	proto3.proto

It has these top-level messages:
	Request
	Book
*/
package proto3

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

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "SWEET",
	1: "SOUR",
	2: "UMAMI",
	3: "GOPHERLICIOUS",
}
var Request_Flavour_value = map[string]int32{
	"SWEET":         0,
	"SOUR":          1,
	"UMAMI":         2,
	"GOPHERLICIOUS": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}
func (Request_Flavour) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Request struct {
	Name     string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty" xml:"name,omitempty"`
	Key      []int64         `protobuf:"varint,2,rep,packed,name=key" json:"key,omitempty" xml:"key,omitempty"`
	Taste    Request_Flavour `protobuf:"varint,3,opt,name=taste,enum=proto3.Request_Flavour" json:"taste,omitempty" xml:"taste,omitempty"`
	Book     *Book           `protobuf:"bytes,4,opt,name=book" json:"book,omitempty" xml:"book,omitempty"`
	Unpacked []int64         `protobuf:"varint,5,rep,packed,name=unpacked" json:"unpacked,omitempty" xml:"unpacked,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty" xml:"title,omitempty"`
	RawData []byte `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty" xml:"raw_data,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x50, 0xbb, 0x4e, 0xc3, 0x40,
	0x10, 0xe4, 0xfc, 0x20, 0xce, 0x62, 0xd0, 0xb1, 0x42, 0xe2, 0x68, 0xd0, 0xc9, 0x95, 0x1b, 0x5c,
	0x84, 0x82, 0x86, 0x86, 0x80, 0x01, 0x4b, 0x44, 0x46, 0x67, 0x2c, 0x4a, 0xb4, 0x21, 0x57, 0x20,
	0x87, 0x5c, 0x70, 0xce, 0x44, 0xfc, 0x2c, 0xdf, 0x82, 0xfc, 0x08, 0xd5, 0xce, 0xcc, 0x8e, 0x66,
	0xb4, 0x0b, 0xe1, 0xba, 0x36, 0xd6, 0x5c, 0x26, 0xdd, 0xc0, 0xfd, 0x9e, 0x45, 0xbf, 0x0c, 0x46,
	0x4a, 0x7f, 0x35, 0x7a, 0x63, 0x11, 0xc1, 0x5b, 0xd1, 0xa7, 0x16, 0x4c, 0xb2, 0x78, 0xac, 0x3a,
	0x8c, 0x1c, 0xdc, 0x4a, 0xff, 0x08, 0x47, 0xba, 0xb1, 0xab, 0x5a, 0x88, 0x17, 0xe0, 0x5b, 0xda,
	0x58, 0x2d, 0x5c, 0xc9, 0xe2, 0xa3, 0xc9, 0x69, 0x32, 0xe4, 0x0e, 0x29, 0xc9, 0xfd, 0x92, 0xbe,
	0x4d, 0x53, 0xab, 0xde, 0x85, 0x12, 0xbc, 0xb9, 0x31, 0x95, 0xf0, 0x24, 0x8b, 0x0f, 0x26, 0xe1,
	0xce, 0x3d, 0x35, 0xa6, 0x52, 0xdd, 0x06, 0xcf, 0x21, 0x68, 0x56, 0x6b, 0x7a, 0xaf, 0xf4, 0x42,
	0xf8, 0x6d, 0xcf, 0xd4, 0xe1, 0x4c, 0xfd, 0x6b, 0xd1, 0x35, 0x8c, 0x86, 0x4c, 0x1c, 0x83, 0x5f,
	0xbc, 0xa6, 0xe9, 0x0b, 0xdf, 0xc3, 0x00, 0xbc, 0x22, 0x2f, 0x15, 0x67, 0xad, 0x58, 0xce, 0x6e,
	0x66, 0x19, 0x77, 0xf0, 0x18, 0x0e, 0x1f, 0xf2, 0xe7, 0xc7, 0x54, 0x3d, 0x65, 0xb7, 0x59, 0x5e,
	0x16, 0xdc, 0x8d, 0xae, 0xc0, 0x6b, 0xbb, 0xf0, 0x04, 0x7c, 0xfb, 0x61, 0x97, 0xbb, 0xeb, 0x7a,
	0x82, 0x67, 0x10, 0xd4, 0xb4, 0x7d, 0x5b, 0x90, 0x25, 0xe1, 0x48, 0x16, 0x87, 0x6a, 0x54, 0xd3,
	0xf6, 0x8e, 0x2c, 0xcd, 0x87, 0x0f, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x94, 0xa9, 0x6c, 0x37,
	0x38, 0x01, 0x00, 0x00,
}
