// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	messages "github.com/CastyLab/grpc.proto/messages"
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

type GetMessagesRequest struct {
	ReceiverId           string               `protobuf:"bytes,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	AuthRequest          *AuthenticateRequest `protobuf:"bytes,2,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetMessagesRequest) Reset()         { *m = GetMessagesRequest{} }
func (m *GetMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessagesRequest) ProtoMessage()    {}
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9d17ea2f814373, []int{0}
}

func (m *GetMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessagesRequest.Unmarshal(m, b)
}
func (m *GetMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessagesRequest.Marshal(b, m, deterministic)
}
func (m *GetMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessagesRequest.Merge(m, src)
}
func (m *GetMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessagesRequest.Size(m)
}
func (m *GetMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessagesRequest proto.InternalMessageInfo

func (m *GetMessagesRequest) GetReceiverId() string {
	if m != nil {
		return m.ReceiverId
	}
	return ""
}

func (m *GetMessagesRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type GetMessagesResponse struct {
	Code                 int64               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string              `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string              `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*messages.Message `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetMessagesResponse) Reset()         { *m = GetMessagesResponse{} }
func (m *GetMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*GetMessagesResponse) ProtoMessage()    {}
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9d17ea2f814373, []int{1}
}

func (m *GetMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessagesResponse.Unmarshal(m, b)
}
func (m *GetMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessagesResponse.Marshal(b, m, deterministic)
}
func (m *GetMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessagesResponse.Merge(m, src)
}
func (m *GetMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_GetMessagesResponse.Size(m)
}
func (m *GetMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessagesResponse proto.InternalMessageInfo

func (m *GetMessagesResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetMessagesResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetMessagesResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetMessagesResponse) GetResult() []*messages.Message {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateMessageRequest struct {
	RecieverId           string               `protobuf:"bytes,1,opt,name=reciever_id,json=recieverId,proto3" json:"reciever_id,omitempty"`
	Content              string               `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthRequest          *AuthenticateRequest `protobuf:"bytes,3,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateMessageRequest) Reset()         { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()    {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9d17ea2f814373, []int{2}
}

func (m *CreateMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageRequest.Unmarshal(m, b)
}
func (m *CreateMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageRequest.Marshal(b, m, deterministic)
}
func (m *CreateMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageRequest.Merge(m, src)
}
func (m *CreateMessageRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMessageRequest.Size(m)
}
func (m *CreateMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageRequest proto.InternalMessageInfo

func (m *CreateMessageRequest) GetRecieverId() string {
	if m != nil {
		return m.RecieverId
	}
	return ""
}

func (m *CreateMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateMessageRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type CreateMessageResponse struct {
	Code                 int64             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string            `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string            `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Result               *messages.Message `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateMessageResponse) Reset()         { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()    {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c9d17ea2f814373, []int{3}
}

func (m *CreateMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessageResponse.Unmarshal(m, b)
}
func (m *CreateMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessageResponse.Marshal(b, m, deterministic)
}
func (m *CreateMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessageResponse.Merge(m, src)
}
func (m *CreateMessageResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMessageResponse.Size(m)
}
func (m *CreateMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessageResponse proto.InternalMessageInfo

func (m *CreateMessageResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateMessageResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateMessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateMessageResponse) GetResult() *messages.Message {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMessagesRequest)(nil), "proto.GetMessagesRequest")
	proto.RegisterType((*GetMessagesResponse)(nil), "proto.GetMessagesResponse")
	proto.RegisterType((*CreateMessageRequest)(nil), "proto.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "proto.CreateMessageResponse")
}

func init() { proto.RegisterFile("message.service.proto", fileDescriptor_6c9d17ea2f814373) }

var fileDescriptor_6c9d17ea2f814373 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0x4d, 0xad, 0x74, 0xa2, 0x14, 0x57, 0x5b, 0x62, 0x14, 0x2c, 0x3d, 0xd5, 0x4b, 0x84,
	0x7a, 0xf6, 0x20, 0x1e, 0x6a, 0x05, 0x2f, 0x11, 0xcf, 0x25, 0x4d, 0x06, 0x1b, 0xd0, 0xa4, 0xee,
	0x4c, 0xfa, 0x09, 0xe2, 0xd1, 0xbf, 0xf0, 0x37, 0xa5, 0x9b, 0xd9, 0x43, 0x6d, 0x10, 0x3c, 0x78,
	0x4a, 0x66, 0xde, 0xdb, 0x79, 0xfb, 0xde, 0x0e, 0xf4, 0x5e, 0x91, 0x28, 0x79, 0xc6, 0x88, 0xd0,
	0xac, 0xf2, 0x14, 0xa3, 0xa5, 0x29, 0xb9, 0xd4, 0xbb, 0xf6, 0x13, 0xc2, 0x3c, 0x21, 0x69, 0x85,
	0x7d, 0x61, 0xd2, 0xa5, 0x3b, 0x62, 0xfb, 0x43, 0x06, 0x3d, 0x41, 0x7e, 0x10, 0x30, 0xc6, 0xb7,
	0x0a, 0x89, 0xf5, 0x39, 0xf8, 0x06, 0x53, 0xcc, 0x57, 0x68, 0x66, 0x79, 0x16, 0xa8, 0x81, 0x1a,
	0x75, 0x62, 0x70, 0xad, 0x69, 0xa6, 0xaf, 0x61, 0x3f, 0xa9, 0x78, 0x31, 0x33, 0xf5, 0x81, 0x60,
	0x67, 0xa0, 0x46, 0xfe, 0x38, 0xac, 0x87, 0x46, 0x37, 0x15, 0x2f, 0xb0, 0xe0, 0x3c, 0x4d, 0x18,
	0x65, 0x64, 0xec, 0xaf, 0xf9, 0x52, 0x0c, 0xdf, 0x15, 0x1c, 0x6d, 0xc8, 0xd2, 0xb2, 0x2c, 0x08,
	0xb5, 0x86, 0x56, 0x5a, 0x66, 0x68, 0x05, 0xbd, 0xd8, 0xfe, 0xeb, 0x3e, 0xb4, 0x89, 0x13, 0xae,
	0xc8, 0x8a, 0x74, 0x62, 0xa9, 0x74, 0x00, 0x7b, 0x62, 0x25, 0x68, 0x59, 0xc0, 0x95, 0xfa, 0x02,
	0xda, 0x06, 0xa9, 0x7a, 0xe1, 0xc0, 0x1b, 0x78, 0x23, 0x7f, 0x7c, 0x18, 0x39, 0xf3, 0x91, 0x28,
	0xc6, 0x42, 0x18, 0x7e, 0x2a, 0x38, 0xbe, 0x35, 0x98, 0x30, 0x3a, 0x64, 0x23, 0x81, 0x1c, 0xb7,
	0x12, 0xb0, 0xad, 0x69, 0xb6, 0x96, 0x4f, 0xcb, 0x82, 0xb1, 0x60, 0xb9, 0x97, 0x2b, 0xb7, 0xb2,
	0xf1, 0xfe, 0x96, 0xcd, 0x87, 0x82, 0xde, 0x8f, 0x2b, 0xfd, 0x5b, 0x3a, 0xea, 0xd7, 0x74, 0xc6,
	0x5f, 0x0a, 0xba, 0xee, 0x8d, 0x1e, 0xeb, 0x0d, 0xd3, 0x77, 0xd0, 0x9d, 0x20, 0x3f, 0x11, 0x1a,
	0x87, 0xe8, 0x13, 0xb1, 0xb6, 0xbd, 0x48, 0x61, 0xd8, 0x04, 0x89, 0x9d, 0x7b, 0x38, 0xd8, 0xf0,
	0xa9, 0x4f, 0x85, 0xdc, 0xf4, 0x20, 0xe1, 0x59, 0x33, 0x58, 0xcf, 0x9a, 0xb7, 0x2d, 0x78, 0xf5,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x40, 0x3c, 0x74, 0xf1, 0x11, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagesServiceClient is the client API for MessagesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagesServiceClient interface {
	GetUserMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
}

type messagesServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessagesServiceClient(cc *grpc.ClientConn) MessagesServiceClient {
	return &messagesServiceClient{cc}
}

func (c *messagesServiceClient) GetUserMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/proto.MessagesService/GetUserMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.MessagesService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServiceServer is the server API for MessagesService service.
type MessagesServiceServer interface {
	GetUserMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
}

// UnimplementedMessagesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessagesServiceServer struct {
}

func (*UnimplementedMessagesServiceServer) GetUserMessages(ctx context.Context, req *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMessages not implemented")
}
func (*UnimplementedMessagesServiceServer) CreateMessage(ctx context.Context, req *CreateMessageRequest) (*CreateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}

func RegisterMessagesServiceServer(s *grpc.Server, srv MessagesServiceServer) {
	s.RegisterService(&_MessagesService_serviceDesc, srv)
}

func _MessagesService_GetUserMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).GetUserMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessagesService/GetUserMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).GetUserMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagesService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MessagesService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessagesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MessagesService",
	HandlerType: (*MessagesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserMessages",
			Handler:    _MessagesService_GetUserMessages_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _MessagesService_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.service.proto",
}