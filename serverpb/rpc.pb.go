// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverpb/rpc.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Req struct {
	Axis                 int32    `protobuf:"varint,1,opt,name=axis,proto3" json:"axis,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_781ae98e25678b32, []int{0}
}
func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (dst *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(dst, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

func (m *Req) GetAxis() int32 {
	if m != nil {
		return m.Axis
	}
	return 0
}

func (m *Req) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Reply struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_781ae98e25678b32, []int{1}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "rpc.Req")
	proto.RegisterType((*Reply)(nil), "rpc.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExtAxisClient is the client API for ExtAxis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExtAxisClient interface {
	// set ExtAxis pos to server
	GotoExtaxisPos(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error)
}

type extAxisClient struct {
	cc *grpc.ClientConn
}

func NewExtAxisClient(cc *grpc.ClientConn) ExtAxisClient {
	return &extAxisClient{cc}
}

func (c *extAxisClient) GotoExtaxisPos(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpc.ExtAxis/GotoExtaxisPos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtAxisServer is the server API for ExtAxis service.
type ExtAxisServer interface {
	// set ExtAxis pos to server
	GotoExtaxisPos(context.Context, *Req) (*Reply, error)
}

func RegisterExtAxisServer(s *grpc.Server, srv ExtAxisServer) {
	s.RegisterService(&_ExtAxis_serviceDesc, srv)
}

func _ExtAxis_GotoExtaxisPos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtAxisServer).GotoExtaxisPos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ExtAxis/GotoExtaxisPos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtAxisServer).GotoExtaxisPos(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExtAxis_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ExtAxis",
	HandlerType: (*ExtAxisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GotoExtaxisPos",
			Handler:    _ExtAxis_GotoExtaxisPos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb/rpc.proto",
}

func init() { proto.RegisterFile("serverpb/rpc.proto", fileDescriptor_rpc_781ae98e25678b32) }

var fileDescriptor_rpc_781ae98e25678b32 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x2a, 0x48, 0xd2, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2e, 0x2a, 0x48, 0x56, 0xd2, 0xe6, 0x62, 0x0e, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0x49, 0xac,
	0xc8, 0x2c, 0x96, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0xf3,
	0x4a, 0x73, 0x25, 0x98, 0xc0, 0x42, 0x20, 0xa6, 0x92, 0x24, 0x17, 0x6b, 0x50, 0x6a, 0x41, 0x4e,
	0x25, 0x48, 0xaa, 0x28, 0x15, 0xa2, 0x9a, 0x33, 0x08, 0xc4, 0x34, 0x32, 0xe6, 0x62, 0x77, 0xad,
	0x28, 0x71, 0x04, 0xe9, 0xd3, 0xe0, 0xe2, 0x73, 0xcf, 0x2f, 0xc9, 0x77, 0xad, 0x28, 0x01, 0x19,
	0x13, 0x90, 0x5f, 0x2c, 0xc4, 0xa1, 0x07, 0xb2, 0x35, 0x28, 0xb5, 0x50, 0x8a, 0x0b, 0xca, 0x2a,
	0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0xb3, 0x21, 0x2d, 0x9e, 0x00, 0x00, 0x00,
}
