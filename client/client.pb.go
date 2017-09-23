// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	client/client.proto

It has these top-level messages:
	Client
*/
package client

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

type Client struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name" json:"last_name,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Uid       string `protobuf:"bytes,4,opt,name=uid,json=id" json:"uid,omitempty"`
	Sid       string `protobuf:"bytes,5,opt,name=sid" json:"sid,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Client) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Client) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Client) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func init() {
	proto.RegisterType((*Client)(nil), "client.Client")
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x52,
	0x2d, 0x17, 0x9b, 0x33, 0x98, 0x25, 0x24, 0xc7, 0xc5, 0x95, 0x96, 0x59, 0x54, 0x5c, 0x12, 0x9f,
	0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x24, 0xc3, 0xc5,
	0x99, 0x93, 0x08, 0x93, 0x66, 0x02, 0x4b, 0x23, 0x04, 0x84, 0x84, 0xb8, 0x58, 0xc0, 0x12, 0xcc,
	0x60, 0x09, 0x30, 0x5b, 0x88, 0x9f, 0x8b, 0xb9, 0x34, 0x33, 0x45, 0x82, 0x05, 0x2c, 0xc4, 0x94,
	0x99, 0x22, 0x24, 0xc0, 0xc5, 0x5c, 0x9c, 0x99, 0x22, 0xc1, 0x0a, 0x16, 0x00, 0x31, 0x93, 0xd8,
	0xc0, 0xae, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x91, 0xb6, 0x4e, 0xab, 0xa4, 0x00, 0x00,
	0x00,
}
