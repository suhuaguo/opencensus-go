// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package testdata is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	FooRequest
	FooResponse
*/
package testdata

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

type FooRequest struct {
	Fail bool `protobuf:"varint,1,opt,name=fail" json:"fail,omitempty"`
}

func (m *FooRequest) Reset()                    { *m = FooRequest{} }
func (m *FooRequest) String() string            { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()               {}
func (*FooRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FooRequest) GetFail() bool {
	if m != nil {
		return m.Fail
	}
	return false
}

type FooResponse struct {
}

func (m *FooResponse) Reset()                    { *m = FooResponse{} }
func (m *FooResponse) String() string            { return proto.CompactTextString(m) }
func (*FooResponse) ProtoMessage()               {}
func (*FooResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*FooRequest)(nil), "testdata.FooRequest")
	proto.RegisterType((*FooResponse)(nil), "testdata.FooResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Foo service

type FooClient interface {
	Single(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
	Multiple(ctx context.Context, opts ...grpc.CallOption) (Foo_MultipleClient, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Single(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := grpc.Invoke(ctx, "/testdata.Foo/Single", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fooClient) Multiple(ctx context.Context, opts ...grpc.CallOption) (Foo_MultipleClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Foo_serviceDesc.Streams[0], c.cc, "/testdata.Foo/Multiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &fooMultipleClient{stream}
	return x, nil
}

type Foo_MultipleClient interface {
	Send(*FooRequest) error
	Recv() (*FooResponse, error)
	grpc.ClientStream
}

type fooMultipleClient struct {
	grpc.ClientStream
}

func (x *fooMultipleClient) Send(m *FooRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fooMultipleClient) Recv() (*FooResponse, error) {
	m := new(FooResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Foo service

type FooServer interface {
	Single(context.Context, *FooRequest) (*FooResponse, error)
	Multiple(Foo_MultipleServer) error
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Single_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Single(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Foo/Single",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Single(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Foo_Multiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FooServer).Multiple(&fooMultipleServer{stream})
}

type Foo_MultipleServer interface {
	Send(*FooResponse) error
	Recv() (*FooRequest, error)
	grpc.ServerStream
}

type fooMultipleServer struct {
	grpc.ServerStream
}

func (x *fooMultipleServer) Send(m *FooResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fooMultipleServer) Recv() (*FooRequest, error) {
	m := new(FooRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Single",
			Handler:    _Foo_Single_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Multiple",
			Handler:       _Foo_Multiple_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xb1, 0x53, 0x12, 0x4b, 0x12, 0x95, 0x14,
	0xb8, 0xb8, 0xdc, 0xf2, 0xf3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58,
	0xd2, 0x12, 0x33, 0x73, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0xc0, 0x6c, 0x25, 0x5e, 0x2e,
	0x6e, 0xb0, 0x8a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3, 0x4a, 0x2e, 0x66, 0xb7, 0xfc, 0x7c,
	0x21, 0x53, 0x2e, 0xb6, 0xe0, 0xcc, 0xbc, 0xf4, 0x9c, 0x54, 0x21, 0x11, 0x3d, 0x98, 0x61, 0x7a,
	0x08, 0x93, 0xa4, 0x44, 0xd1, 0x44, 0x21, 0xba, 0x85, 0xac, 0xb9, 0x38, 0x7c, 0x4b, 0x73, 0x4a,
	0x32, 0x0b, 0x48, 0xd4, 0xa8, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0xbc, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x60, 0x13, 0xc6, 0xca, 0x00, 0x00, 0x00,
}
//go:generate ./generate.sh

