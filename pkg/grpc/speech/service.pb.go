// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto
package grpcSpeech

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SpeechRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpeechRequest) Reset()         { *m = SpeechRequest{} }
func (m *SpeechRequest) String() string { return proto.CompactTextString(m) }
func (*SpeechRequest) ProtoMessage()    {}
func (*SpeechRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *SpeechRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeechRequest.Unmarshal(m, b)
}
func (m *SpeechRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeechRequest.Marshal(b, m, deterministic)
}
func (m *SpeechRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeechRequest.Merge(m, src)
}
func (m *SpeechRequest) XXX_Size() int {
	return xxx_messageInfo_SpeechRequest.Size(m)
}
func (m *SpeechRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeechRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpeechRequest proto.InternalMessageInfo

func (m *SpeechRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SpeechResponse struct {
	Duration             int32    `protobuf:"varint,1,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpeechResponse) Reset()         { *m = SpeechResponse{} }
func (m *SpeechResponse) String() string { return proto.CompactTextString(m) }
func (*SpeechResponse) ProtoMessage()    {}
func (*SpeechResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *SpeechResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeechResponse.Unmarshal(m, b)
}
func (m *SpeechResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeechResponse.Marshal(b, m, deterministic)
}
func (m *SpeechResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeechResponse.Merge(m, src)
}
func (m *SpeechResponse) XXX_Size() int {
	return xxx_messageInfo_SpeechResponse.Size(m)
}
func (m *SpeechResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeechResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpeechResponse proto.InternalMessageInfo

func (m *SpeechResponse) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *SpeechResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SpeechRequest)(nil), "SpeechRequest")
	proto.RegisterType((*SpeechResponse)(nil), "SpeechResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe4, 0xe2, 0x0d, 0x2e, 0x48,
	0x4d, 0x4d, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xdc,
	0xb8, 0xf8, 0x60, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x5c, 0x4a,
	0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xc0, 0x8a, 0x59, 0x83, 0xe0, 0x7c, 0x64, 0x73, 0x98, 0x50,
	0xcc, 0x31, 0x2a, 0xe3, 0xe2, 0x4a, 0x2f, 0x2a, 0x48, 0x86, 0x98, 0x25, 0x64, 0xc2, 0xc5, 0xef,
	0x92, 0x5a, 0x92, 0x5a, 0x94, 0x9b, 0x99, 0x97, 0x0a, 0x15, 0xe2, 0xd3, 0x43, 0x71, 0x92, 0x14,
	0xbf, 0x1e, 0xaa, 0xbd, 0x4a, 0x0c, 0x42, 0x46, 0x5c, 0xbc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0x61, 0x89, 0x45, 0x99, 0x89, 0x79, 0x25, 0x44, 0xe8, 0x49, 0x62, 0x03, 0xfb, 0xd8, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x09, 0x04, 0xb1, 0x02, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcSpeechClient is the client API for GrpcSpeech service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcSpeechClient interface {
	DetermineSpeech(ctx context.Context, in *SpeechRequest, opts ...grpc.CallOption) (*SpeechResponse, error)
	CreateVariant(ctx context.Context, in *SpeechRequest, opts ...grpc.CallOption) (*SpeechResponse, error)
}

type grpcSpeechClient struct {
	cc *grpc.ClientConn
}

func NewGrpcSpeechClient(cc *grpc.ClientConn) GrpcSpeechClient {
	return &grpcSpeechClient{cc}
}

func (c *grpcSpeechClient) DetermineSpeech(ctx context.Context, in *SpeechRequest, opts ...grpc.CallOption) (*SpeechResponse, error) {
	out := new(SpeechResponse)
	err := c.cc.Invoke(ctx, "/grpcSpeech/DetermineSpeech", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcSpeechClient) CreateVariant(ctx context.Context, in *SpeechRequest, opts ...grpc.CallOption) (*SpeechResponse, error) {
	out := new(SpeechResponse)
	err := c.cc.Invoke(ctx, "/grpcSpeech/CreateVariant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcSpeechServer is the server API for GrpcSpeech service.
type GrpcSpeechServer interface {
	DetermineSpeech(context.Context, *SpeechRequest) (*SpeechResponse, error)
	CreateVariant(context.Context, *SpeechRequest) (*SpeechResponse, error)
}

// UnimplementedGrpcSpeechServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcSpeechServer struct {
}

func (*UnimplementedGrpcSpeechServer) DetermineSpeech(ctx context.Context, req *SpeechRequest) (*SpeechResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetermineSpeech not implemented")
}
func (*UnimplementedGrpcSpeechServer) CreateVariant(ctx context.Context, req *SpeechRequest) (*SpeechResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVariant not implemented")
}

func RegisterGrpcSpeechServer(s *grpc.Server, srv GrpcSpeechServer) {
	s.RegisterService(&_GrpcSpeech_serviceDesc, srv)
}

func _GrpcSpeech_DetermineSpeech_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeechRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcSpeechServer).DetermineSpeech(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcSpeech/DetermineSpeech",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcSpeechServer).DetermineSpeech(ctx, req.(*SpeechRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcSpeech_CreateVariant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeechRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcSpeechServer).CreateVariant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcSpeech/CreateVariant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcSpeechServer).CreateVariant(ctx, req.(*SpeechRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcSpeech_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcSpeech",
	HandlerType: (*GrpcSpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetermineSpeech",
			Handler:    _GrpcSpeech_DetermineSpeech_Handler,
		},
		{
			MethodName: "CreateVariant",
			Handler:    _GrpcSpeech_CreateVariant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
