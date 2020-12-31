// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatty.proto

package chattyproto

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

type Empty struct {
	Empty                string   `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1db09a1aceff296f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetEmpty() string {
	if m != nil {
		return m.Empty
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "chattyproto.empty")
}

func init() {
	proto.RegisterFile("chatty.proto", fileDescriptor_1db09a1aceff296f)
}

var fileDescriptor_1db09a1aceff296f = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xc1, 0x4a, 0xc3, 0x50,
	0x10, 0x45, 0x29, 0xa2, 0xe0, 0x44, 0xa1, 0x3c, 0x8a, 0x4a, 0x55, 0x10, 0x57, 0xae, 0xb2, 0xd0,
	0x2f, 0x90, 0x80, 0x48, 0x21, 0x9b, 0xa4, 0xae, 0xe5, 0x35, 0xb9, 0x98, 0x80, 0x99, 0x89, 0x6f,
	0xc6, 0x42, 0xff, 0xd9, 0x8f, 0x10, 0xfa, 0x22, 0x48, 0xb7, 0xc9, 0x6e, 0xce, 0xdc, 0xe1, 0x5c,
	0x18, 0x3a, 0xab, 0x1a, 0x6f, 0xb6, 0x4b, 0xfb, 0x20, 0x26, 0x2e, 0x89, 0xb4, 0x87, 0xe5, 0x22,
	0xc2, 0x7b, 0x07, 0x55, 0xff, 0x81, 0x78, 0x72, 0x7f, 0x4b, 0xc7, 0xe8, 0x7a, 0xdb, 0xb9, 0xc5,
	0x30, 0x5c, 0xcd, 0xee, 0x66, 0x0f, 0xa7, 0x45, 0x84, 0xc7, 0x9f, 0x23, 0x4a, 0xde, 0x14, 0xa1,
	0x44, 0xd8, 0xb6, 0x15, 0x5c, 0x4e, 0xf3, 0x2c, 0xc0, 0x1b, 0x32, 0x61, 0x46, 0x65, 0xad, 0xb0,
	0xbb, 0x4e, 0xff, 0xd5, 0xa4, 0x05, 0xbe, 0xbe, 0xa1, 0x96, 0xc7, 0x96, 0xe5, 0xcd, 0x41, 0xa8,
	0xbd, 0xb0, 0x62, 0x48, 0xdd, 0x2b, 0x25, 0x25, 0xb8, 0xfe, 0xc3, 0x11, 0xa6, 0x92, 0x2e, 0x0a,
	0xf4, 0x12, 0x2c, 0x6b, 0x3c, 0x33, 0x3e, 0x75, 0x08, 0x74, 0x8c, 0x74, 0x45, 0xe7, 0x2f, 0xb0,
	0xaa, 0x99, 0xc2, 0x95, 0xd3, 0x7c, 0x25, 0x2d, 0x67, 0xc2, 0x5b, 0x04, 0xf5, 0x63, 0x3f, 0xb7,
	0xa6, 0xcb, 0xe7, 0xba, 0xce, 0xd1, 0x6d, 0x10, 0xd6, 0x32, 0x91, 0x75, 0x73, 0xb2, 0x5f, 0x3f,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x54, 0x30, 0x5f, 0x47, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateConnection(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	SendMessage(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	ReportChannelsMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	FetchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	JoinConversation(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
	AddMemberToConversation(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateConnection(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendMessage(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReportChannelsMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/ReportChannelsMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FetchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/FetchMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) JoinConversation(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/JoinConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddMemberToConversation(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/chattyproto.UserService/AddMemberToConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateConnection(context.Context, *RequestMessage) (*ResponseMessage, error)
	SendMessage(context.Context, *RequestMessage) (*ResponseMessage, error)
	ReportChannelsMessages(context.Context, *RequestMessage) (*ResponseMessage, error)
	FetchMessages(context.Context, *RequestMessage) (*ResponseMessage, error)
	JoinConversation(context.Context, *RequestMessage) (*ResponseMessage, error)
	AddMemberToConversation(context.Context, *RequestMessage) (*ResponseMessage, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateConnection(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (*UnimplementedUserServiceServer) SendMessage(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedUserServiceServer) ReportChannelsMessages(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportChannelsMessages not implemented")
}
func (*UnimplementedUserServiceServer) FetchMessages(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMessages not implemented")
}
func (*UnimplementedUserServiceServer) JoinConversation(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConversation not implemented")
}
func (*UnimplementedUserServiceServer) AddMemberToConversation(ctx context.Context, req *RequestMessage) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberToConversation not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateConnection(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendMessage(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReportChannelsMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReportChannelsMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/ReportChannelsMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReportChannelsMessages(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FetchMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FetchMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/FetchMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FetchMessages(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_JoinConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).JoinConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/JoinConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).JoinConversation(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddMemberToConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddMemberToConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chattyproto.UserService/AddMemberToConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddMemberToConversation(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chattyproto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConnection",
			Handler:    _UserService_CreateConnection_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _UserService_SendMessage_Handler,
		},
		{
			MethodName: "ReportChannelsMessages",
			Handler:    _UserService_ReportChannelsMessages_Handler,
		},
		{
			MethodName: "FetchMessages",
			Handler:    _UserService_FetchMessages_Handler,
		},
		{
			MethodName: "JoinConversation",
			Handler:    _UserService_JoinConversation_Handler,
		},
		{
			MethodName: "AddMemberToConversation",
			Handler:    _UserService_AddMemberToConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatty.proto",
}
