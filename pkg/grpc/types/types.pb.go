// Code generated by protoc-gen-go-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	RpcFlowRequest
	RpcFlowResponse
*/
package types

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

type RpcFlowRequest struct {
	Request string `protobuf:"bytes,1,opt,name=Request" json:"Request,omitempty"`
}

func (m *RpcFlowRequest) Reset()                    { *m = RpcFlowRequest{} }
func (m *RpcFlowRequest) String() string            { return proto.CompactTextString(m) }
func (*RpcFlowRequest) ProtoMessage()               {}
func (*RpcFlowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpcFlowRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type RpcFlowResponse struct {
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *RpcFlowResponse) Reset()                    { *m = RpcFlowResponse{} }
func (m *RpcFlowResponse) String() string            { return proto.CompactTextString(m) }
func (*RpcFlowResponse) ProtoMessage()               {}
func (*RpcFlowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpcFlowResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*RpcFlowRequest)(nil), "types.RpcFlowRequest")
	proto.RegisterType((*RpcFlowResponse)(nil), "types.RpcFlowResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RpcFlow service

type RpcFlowClient interface {
	TestClinet(ctx context.Context, in *RpcFlowRequest, opts ...grpc.CallOption) (*RpcFlowResponse, error)
}

type rpcFlowClient struct {
	cc *grpc.ClientConn
}

func NewRpcFlowClient(cc *grpc.ClientConn) RpcFlowClient {
	return &rpcFlowClient{cc}
}

func (c *rpcFlowClient) TestClinet(ctx context.Context, in *RpcFlowRequest, opts ...grpc.CallOption) (*RpcFlowResponse, error) {
	out := new(RpcFlowResponse)
	err := grpc.Invoke(ctx, "/types.RpcFlow/TestClinet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcFlow service

type RpcFlowServer interface {
	TestClinet(context.Context, *RpcFlowRequest) (*RpcFlowResponse, error)
}

func RegisterRpcFlowServer(s *grpc.Server, srv RpcFlowServer) {
	s.RegisterService(&_RpcFlow_serviceDesc, srv)
}

func _RpcFlow_TestClinet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFlowServer).TestClinet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RpcFlow/TestClinet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFlowServer).TestClinet(ctx, req.(*RpcFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcFlow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.RpcFlow",
	HandlerType: (*RpcFlowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestClinet",
			Handler:    _RpcFlow_TestClinet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xb4, 0xb8, 0xf8, 0x82,
	0x0a, 0x92, 0xdd, 0x72, 0xf2, 0xcb, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8,
	0xd8, 0xa1, 0x4c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x49, 0x97, 0x8b, 0x1f,
	0xae, 0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x03, 0xc6, 0x86, 0xaa, 0x86,
	0xf3, 0x8d, 0x3c, 0xb8, 0xd8, 0xa1, 0xca, 0x85, 0x6c, 0xb9, 0xb8, 0x42, 0x52, 0x8b, 0x4b, 0x9c,
	0x73, 0x32, 0xf3, 0x52, 0x4b, 0x84, 0x44, 0xf5, 0x20, 0x0e, 0x41, 0xb5, 0x58, 0x4a, 0x0c, 0x5d,
	0x18, 0x62, 0x8e, 0x12, 0x83, 0x93, 0x1c, 0x97, 0x70, 0x56, 0x7e, 0x92, 0x6e, 0x72, 0x6a, 0x5e,
	0x49, 0x6a, 0x91, 0x5e, 0x41, 0x76, 0xba, 0x5e, 0x7a, 0x51, 0x41, 0xb2, 0x13, 0xc4, 0x0b, 0x01,
	0x8c, 0x49, 0x6c, 0x60, 0x2f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0x04, 0x5a, 0x2a,
	0xe1, 0x00, 0x00, 0x00,
}
