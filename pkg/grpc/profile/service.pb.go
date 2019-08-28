// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package grpcProfile

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Profile struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	AccessLevel          int32    `protobuf:"varint,2,opt,name=AccessLevel,proto3" json:"AccessLevel,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ID                   string   `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Profile) GetAccessLevel() int32 {
	if m != nil {
		return m.AccessLevel
	}
	return 0
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ProfileRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequest.Unmarshal(m, b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ProfileRequest.Size(m)
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ProfileRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type SuccessResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessResponse) Reset()         { *m = SuccessResponse{} }
func (m *SuccessResponse) String() string { return proto.CompactTextString(m) }
func (*SuccessResponse) ProtoMessage()    {}
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *SuccessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessResponse.Unmarshal(m, b)
}
func (m *SuccessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessResponse.Marshal(b, m, deterministic)
}
func (m *SuccessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessResponse.Merge(m, src)
}
func (m *SuccessResponse) XXX_Size() int {
	return xxx_messageInfo_SuccessResponse.Size(m)
}
func (m *SuccessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessResponse proto.InternalMessageInfo

func (m *SuccessResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Profile)(nil), "grpcProfile.Profile")
	proto.RegisterType((*ProfileRequest)(nil), "grpcProfile.ProfileRequest")
	proto.RegisterType((*SuccessResponse)(nil), "grpcProfile.SuccessResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x2f, 0x2a, 0x48, 0x0e,
	0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x55, 0x4a, 0xe6, 0x62, 0x87, 0x32, 0x85, 0x84, 0xb8, 0x58,
	0x42, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x05, 0x2e,
	0x6e, 0xc7, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0x9f, 0xd4, 0xb2, 0xd4, 0x1c, 0x09, 0x26, 0xb0, 0x14,
	0xb2, 0x10, 0x48, 0x97, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98,
	0x2d, 0xc4, 0xc7, 0xc5, 0xe4, 0xe9, 0x22, 0xc1, 0x02, 0x16, 0x61, 0xf2, 0x74, 0x51, 0x32, 0xe1,
	0xe2, 0x83, 0x5a, 0x12, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x02, 0x55, 0xc1, 0x08, 0x53, 0x01,
	0xb7, 0x9b, 0x09, 0x61, 0xb7, 0x92, 0x36, 0x17, 0x7f, 0x70, 0x29, 0xd8, 0xa6, 0xa0, 0xd4, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0x76, 0xa8, 0x10, 0x58, 0x2f, 0x47, 0x10, 0x8c,
	0x6b, 0xf4, 0x88, 0x91, 0x0b, 0xd9, 0x5f, 0x42, 0xae, 0x5c, 0x7c, 0x41, 0xa9, 0x25, 0x45, 0x99,
	0xa9, 0x65, 0xa9, 0x6e, 0x45, 0xf9, 0xb9, 0x9e, 0x29, 0x42, 0xd2, 0x7a, 0x48, 0xf2, 0x7a, 0xa8,
	0xee, 0x91, 0x12, 0xc1, 0x26, 0xa9, 0xc4, 0x20, 0xe4, 0xca, 0xc5, 0x1b, 0x5a, 0x90, 0x92, 0x58,
	0x92, 0x0a, 0x33, 0x17, 0xab, 0x42, 0x29, 0x19, 0x14, 0x51, 0x34, 0x57, 0x43, 0x8c, 0x71, 0x2e,
	0x4a, 0xa5, 0xd4, 0x98, 0x24, 0x36, 0x70, 0x04, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x49,
	0x8a, 0xcd, 0x43, 0xd1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcProfileClient is the client API for GrpcProfile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcProfileClient interface {
	RetrieveFromId(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*SuccessResponse, error)
	CreateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type grpcProfileClient struct {
	cc *grpc.ClientConn
}

func NewGrpcProfileClient(cc *grpc.ClientConn) GrpcProfileClient {
	return &grpcProfileClient{cc}
}

func (c *grpcProfileClient) RetrieveFromId(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/grpcProfile.grpcProfile/RetrieveFromId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcProfileClient) UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/grpcProfile.grpcProfile/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcProfileClient) CreateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/grpcProfile.grpcProfile/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcProfileServer is the server API for GrpcProfile service.
type GrpcProfileServer interface {
	RetrieveFromId(context.Context, *ProfileRequest) (*Profile, error)
	UpdateProfile(context.Context, *Profile) (*SuccessResponse, error)
	CreateProfile(context.Context, *Profile) (*SuccessResponse, error)
}

// UnimplementedGrpcProfileServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcProfileServer struct {
}

func (*UnimplementedGrpcProfileServer) RetrieveFromId(ctx context.Context, req *ProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveFromId not implemented")
}
func (*UnimplementedGrpcProfileServer) UpdateProfile(ctx context.Context, req *Profile) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (*UnimplementedGrpcProfileServer) CreateProfile(ctx context.Context, req *Profile) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}

func RegisterGrpcProfileServer(s *grpc.Server, srv GrpcProfileServer) {
	s.RegisterService(&_GrpcProfile_serviceDesc, srv)
}

func _GrpcProfile_RetrieveFromId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcProfileServer).RetrieveFromId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcProfile.grpcProfile/RetrieveFromId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcProfileServer).RetrieveFromId(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcProfile_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcProfileServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcProfile.grpcProfile/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcProfileServer).UpdateProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcProfile_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcProfileServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcProfile.grpcProfile/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcProfileServer).CreateProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcProfile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcProfile.grpcProfile",
	HandlerType: (*GrpcProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveFromId",
			Handler:    _GrpcProfile_RetrieveFromId_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _GrpcProfile_UpdateProfile_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _GrpcProfile_CreateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
