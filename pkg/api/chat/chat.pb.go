// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/chat/chat.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import client "github.com/minimalchat/daemon/pkg/api/client"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Chat struct {
	CreationTime         *timestamp.Timestamp `protobuf:"bytes,1,opt,name=creation_time,proto3" json:"creation_time,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_time,json=update_time,proto3" json:"updated_time,omitempty"`
	Open                 bool                 `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
	Uid                  string               `protobuf:"bytes,4,opt,name=uid,json=id,proto3" json:"uid,omitempty"`
	Client               *client.Client       `protobuf:"bytes,5,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Chat) Reset()         { *m = Chat{} }
func (m *Chat) String() string { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()    {}
func (*Chat) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_c4d626344bc83633, []int{0}
}
func (m *Chat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chat.Unmarshal(m, b)
}
func (m *Chat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chat.Marshal(b, m, deterministic)
}
func (dst *Chat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chat.Merge(dst, src)
}
func (m *Chat) XXX_Size() int {
	return xxx_messageInfo_Chat.Size(m)
}
func (m *Chat) XXX_DiscardUnknown() {
	xxx_messageInfo_Chat.DiscardUnknown(m)
}

var xxx_messageInfo_Chat proto.InternalMessageInfo

func (m *Chat) GetCreationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Chat) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

func (m *Chat) GetOpen() bool {
	if m != nil {
		return m.Open
	}
	return false
}

func (m *Chat) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Chat) GetClient() *client.Client {
	if m != nil {
		return m.Client
	}
	return nil
}

type Message struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author               string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Chat                 string               `protobuf:"bytes,4,opt,name=chat,proto3" json:"chat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_c4d626344bc83633, []int{1}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Message) GetChat() string {
	if m != nil {
		return m.Chat
	}
	return ""
}

func init() {
	proto.RegisterType((*Chat)(nil), "chat.Chat")
	proto.RegisterType((*Message)(nil), "chat.Message")
}

func init() { proto.RegisterFile("pkg/api/chat/chat.proto", fileDescriptor_chat_c4d626344bc83633) }

var fileDescriptor_chat_c4d626344bc83633 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0xe5, 0x36, 0xa4, 0xe4, 0xca, 0x87, 0xe4, 0x01, 0xac, 0x30, 0x10, 0x75, 0x40, 0x99,
	0x1c, 0x09, 0x16, 0x16, 0x24, 0xa4, 0xce, 0x2c, 0x16, 0x3b, 0x72, 0x13, 0x93, 0x58, 0xb4, 0xb1,
	0x95, 0x5c, 0x1e, 0x82, 0xf7, 0xe3, 0x81, 0x90, 0x3f, 0x52, 0xc4, 0x04, 0x8b, 0xef, 0xce, 0x77,
	0xf7, 0xbb, 0xff, 0xd9, 0x70, 0x6d, 0x3f, 0xda, 0x4a, 0x5a, 0x5d, 0xd5, 0x9d, 0x44, 0x7f, 0x70,
	0x3b, 0x18, 0x34, 0x34, 0x71, 0x7e, 0x7e, 0xdb, 0x1a, 0xd3, 0xee, 0x55, 0xe5, 0xef, 0x76, 0xd3,
	0x7b, 0x85, 0xfa, 0xa0, 0x46, 0x94, 0x07, 0x1b, 0xca, 0xf2, 0x9b, 0x63, 0xff, 0x5e, 0xab, 0x1e,
	0xa3, 0x09, 0xc9, 0xcd, 0x17, 0x81, 0x64, 0xdb, 0x49, 0xa4, 0xcf, 0x70, 0x5e, 0x0f, 0x4a, 0xa2,
	0x36, 0xfd, 0x9b, 0x23, 0x30, 0x52, 0x90, 0x72, 0x7d, 0x9f, 0xf3, 0x80, 0xe7, 0x33, 0x9e, 0xbf,
	0xce, 0x78, 0xf1, 0xbb, 0x81, 0x3e, 0xc1, 0xd9, 0x64, 0x1b, 0x89, 0xaa, 0x09, 0x80, 0xc5, 0x9f,
	0x80, 0x75, 0xa8, 0x0f, 0xed, 0x14, 0x12, 0x63, 0x55, 0xcf, 0x96, 0x05, 0x29, 0x4f, 0x85, 0xf7,
	0xe9, 0x25, 0x2c, 0x27, 0xdd, 0xb0, 0xa4, 0x20, 0x65, 0x26, 0x16, 0xba, 0xa1, 0x77, 0x90, 0x06,
	0xf9, 0xec, 0xc4, 0xd3, 0x2f, 0x78, 0xdc, 0x66, 0xeb, 0x8d, 0x88, 0xd9, 0xcd, 0x27, 0x81, 0xd5,
	0x8b, 0x1a, 0x47, 0xd9, 0x2a, 0xfa, 0x08, 0xd9, 0xf1, 0x49, 0xfe, 0xb1, 0xd5, 0x4f, 0x31, 0x65,
	0xb0, 0xaa, 0x4d, 0x8f, 0x6e, 0xdc, 0xc2, 0x4b, 0x98, 0x43, 0x7a, 0x05, 0xa9, 0x9c, 0xb0, 0x33,
	0x83, 0x97, 0x9b, 0x89, 0x18, 0xb9, 0x25, 0xdc, 0xa7, 0x44, 0xc5, 0xde, 0xdf, 0xa5, 0x7e, 0xc8,
	0xc3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x8f, 0xac, 0x05, 0xc8, 0x01, 0x00, 0x00,
}
