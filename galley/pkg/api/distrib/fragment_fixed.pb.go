// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: distrib/fragment.proto

/*
Package distrib is a generated protocol buffer package.

It is generated from these files:
	distrib/fragment.proto

It has these top-level messages:
	Fragment
*/
package distrib

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Fragment struct {
	Id      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content []*google_protobuf.Any `protobuf:"bytes,2,rep,name=content" json:"content,omitempty"`
}

func (m *Fragment) Reset()                    { *m = Fragment{} }
func (m *Fragment) String() string            { return proto.CompactTextString(m) }
func (*Fragment) ProtoMessage()               {}
func (*Fragment) Descriptor() ([]byte, []int) { return fileDescriptorFragment, []int{0} }

func (m *Fragment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Fragment) GetContent() []*google_protobuf.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Fragment)(nil), "istio.component.config.Fragment")
}

func init() { proto.RegisterFile("distrib/fragment.proto", fileDescriptorFragment) }

var fileDescriptorFragment = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x4f, 0x0b, 0x82, 0x40,
	0x10, 0x47, 0xd1, 0xa0, 0x3f, 0x1b, 0x74, 0x90, 0x10, 0xeb, 0x24, 0x41, 0xe4, 0x69, 0x07, 0xea,
	0x13, 0xd4, 0xa1, 0x43, 0x47, 0x8f, 0xdd, 0x56, 0x5d, 0x97, 0x21, 0x9d, 0x59, 0x74, 0x3b, 0xf8,
	0xed, 0x23, 0x57, 0x6f, 0xc3, 0xf0, 0x7b, 0x8f, 0x27, 0xe2, 0x0a, 0x7b, 0xd7, 0x61, 0x01, 0x75,
	0xa7, 0x4c, 0xab, 0xc9, 0x49, 0xdb, 0xb1, 0xe3, 0x28, 0xc6, 0xde, 0x21, 0xcb, 0x92, 0x5b, 0xcb,
	0xf4, 0x7f, 0x97, 0x4c, 0x35, 0x9a, 0xe3, 0xc1, 0x30, 0x9b, 0x46, 0xc3, 0xb8, 0x2a, 0xbe, 0x35,
	0x28, 0x1a, 0x3c, 0x72, 0x7a, 0x89, 0xf5, 0x73, 0x92, 0x44, 0x3b, 0x11, 0x62, 0x95, 0x04, 0x69,
	0x90, 0x6d, 0xf2, 0x10, 0xab, 0x48, 0x8a, 0x55, 0xc9, 0xe4, 0x34, 0xb9, 0x24, 0x4c, 0x17, 0xd9,
	0xf6, 0xba, 0x97, 0x5e, 0x24, 0x67, 0x91, 0xbc, 0xd3, 0x90, 0xcf, 0xa3, 0xc7, 0xe5, 0x7d, 0xf6,
	0x01, 0xc8, 0x30, 0x1e, 0x60, 0x54, 0xd3, 0xe8, 0x01, 0xec, 0xc7, 0x80, 0xb2, 0x08, 0x53, 0x76,
	0xb1, 0x1c, 0xf9, 0xdb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xdd, 0x8d, 0xd9, 0xc8, 0x00, 0x00,
	0x00,
}