// Code generated by protoc-gen-go.
// source: error.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*Error)(nil), "protocol.Error")
}

func init() { proto.RegisterFile("error.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x96,
	0x5c, 0xac, 0xae, 0x20, 0x09, 0x21, 0x09, 0x2e, 0xf6, 0xd4, 0xa2, 0x22, 0xe7, 0xfc, 0x94, 0x54,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x18, 0x57, 0x48, 0x8c, 0x8b, 0x2d, 0xb5, 0xa8, 0x28,
	0xb8, 0xa4, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x4b, 0x62, 0x03, 0x1b, 0x62,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x4f, 0x83, 0x83, 0x5a, 0x00, 0x00, 0x00,
}
