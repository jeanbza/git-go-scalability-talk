// Code generated by protoc-gen-go.
// source: application/model/types.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Some generic request with a message
type Request struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "model.Request")
	proto.RegisterType((*Empty)(nil), "model.Empty")
}

func init() { proto.RegisterFile("application/model/types.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x28, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x2f, 0xa9,
	0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x0b, 0x29, 0x29, 0x73,
	0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0xec, 0x5c, 0xac,
	0xae, 0xb9, 0x05, 0x25, 0x95, 0x49, 0x6c, 0x60, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0x16, 0xd5, 0x5f, 0x5c, 0x00, 0x00, 0x00,
}