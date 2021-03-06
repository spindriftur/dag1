// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer.proto

/*
Package peers is a generated protocol buffer package.

It is generated from these files:
	peer.proto

It has these top-level messages:
	PeerMessage
*/
package peers

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

type PeerMessage struct {
	NetAddr   string `protobuf:"bytes,1,opt,name=NetAddr,json=netAddr" json:"NetAddr,omitempty"`
	PubKeyHex string `protobuf:"bytes,2,opt,name=PubKeyHex,json=pubKeyHex" json:"PubKeyHex,omitempty"`
}

func (m *PeerMessage) Reset()                    { *m = PeerMessage{} }
func (m *PeerMessage) String() string            { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()               {}
func (*PeerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PeerMessage) GetNetAddr() string {
	if m != nil {
		return m.NetAddr
	}
	return ""
}

func (m *PeerMessage) GetPubKeyHex() string {
	if m != nil {
		return m.PubKeyHex
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerMessage)(nil), "peers.PeerMessage")
}

func init() { proto.RegisterFile("peer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x4d, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0xb1, 0x8b, 0x95, 0x5c, 0xb9, 0xb8, 0x03,
	0x52, 0x53, 0x8b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0xfd, 0x52,
	0x4b, 0x1c, 0x53, 0x52, 0x8a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xd8, 0xf3, 0x20, 0x5c,
	0x21, 0x19, 0x2e, 0xce, 0x80, 0xd2, 0x24, 0xef, 0xd4, 0x4a, 0x8f, 0xd4, 0x0a, 0x09, 0x26, 0xb0,
	0x1c, 0x67, 0x01, 0x4c, 0x20, 0x89, 0x0d, 0x6c, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef,
	0x73, 0xb3, 0xed, 0x62, 0x00, 0x00, 0x00,
}
