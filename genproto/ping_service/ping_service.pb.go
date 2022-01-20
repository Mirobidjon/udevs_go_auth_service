// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping_service.proto

package ping_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PongResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69f5a10807d3eaeb, []int{0}
}

func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (m *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(m, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PongResponse)(nil), "ping_service.PongResponse")
}

func init() { proto.RegisterFile("ping_service.proto", fileDescriptor_69f5a10807d3eaeb) }

var fileDescriptor_69f5a10807d3eaeb = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xc8, 0xcc, 0x4b,
	0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x93, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xcb, 0x25, 0x95, 0xa6,
	0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x94, 0x2a, 0x69, 0x70, 0xf1, 0x04, 0xe4, 0xe7, 0xa5,
	0x07, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xde, 0x5c, 0xdc,
	0x01, 0x99, 0x79, 0xe9, 0xc1, 0x10, 0x53, 0x85, 0x6c, 0xb8, 0x58, 0x40, 0x5c, 0x21, 0x31, 0x3d,
	0x88, 0xf1, 0x7a, 0x30, 0xe3, 0xf5, 0x5c, 0x41, 0xc6, 0x4b, 0x49, 0xe9, 0xa1, 0x38, 0x0c, 0xd9,
	0x12, 0x25, 0x06, 0x27, 0xf1, 0x28, 0xd1, 0xf4, 0xd4, 0x3c, 0xb0, 0x16, 0x7d, 0x64, 0x75, 0x49,
	0x6c, 0x60, 0x31, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x82, 0x05, 0x2f, 0xd7, 0x00,
	0x00, 0x00,
}
