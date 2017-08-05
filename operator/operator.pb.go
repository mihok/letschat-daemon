// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operator/operator.proto

/*
Package operator is a generated protocol buffer package.

It is generated from these files:
	operator/operator.proto

It has these top-level messages:
	Operator
*/
package operator

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

type Operator struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name" json:"last_name,omitempty"`
	UserName  string `protobuf:"bytes,3,opt,name=user_name,json=username" json:"user_name,omitempty"`
	Uid       string `protobuf:"bytes,4,opt,name=uid,json=id" json:"uid,omitempty"`
}

func (m *Operator) Reset()                    { *m = Operator{} }
func (m *Operator) String() string            { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()               {}
func (*Operator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Operator) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Operator) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Operator) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Operator) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*Operator)(nil), "operator.Operator")
}

func init() { proto.RegisterFile("operator/operator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x2f, 0x48, 0x2d,
	0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0xa5, 0x32, 0x2e, 0x0e, 0x7f, 0x28, 0x5b, 0x48, 0x8e, 0x8b, 0x2b, 0x2d, 0xb3, 0xa8, 0xb8,
	0x24, 0x3e, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0x44, 0x48,
	0x86, 0x8b, 0x33, 0x27, 0x11, 0x26, 0xcd, 0x04, 0x96, 0x46, 0x08, 0x08, 0x49, 0x73, 0x71, 0x96,
	0x16, 0xa7, 0x16, 0x41, 0x64, 0x99, 0xc1, 0xb2, 0x1c, 0x20, 0x01, 0xb0, 0x24, 0x3f, 0x17, 0x73,
	0x69, 0x66, 0x8a, 0x04, 0x0b, 0x58, 0x98, 0x29, 0x33, 0x25, 0x89, 0x0d, 0xec, 0x10, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x2c, 0xa3, 0x1c, 0xa3, 0x00, 0x00, 0x00,
}