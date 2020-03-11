// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GiftsRequest struct {
	Quantity             int64    `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Gifts                string   `protobuf:"bytes,2,opt,name=gifts,proto3" json:"gifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GiftsRequest) Reset()         { *m = GiftsRequest{} }
func (m *GiftsRequest) String() string { return proto.CompactTextString(m) }
func (*GiftsRequest) ProtoMessage()    {}
func (*GiftsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *GiftsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GiftsRequest.Unmarshal(m, b)
}
func (m *GiftsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GiftsRequest.Marshal(b, m, deterministic)
}
func (m *GiftsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GiftsRequest.Merge(m, src)
}
func (m *GiftsRequest) XXX_Size() int {
	return xxx_messageInfo_GiftsRequest.Size(m)
}
func (m *GiftsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GiftsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GiftsRequest proto.InternalMessageInfo

func (m *GiftsRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *GiftsRequest) GetGifts() string {
	if m != nil {
		return m.Gifts
	}
	return ""
}

type GiftsReply struct {
	Quantity             int64    `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Gifts                string   `protobuf:"bytes,2,opt,name=gifts,proto3" json:"gifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GiftsReply) Reset()         { *m = GiftsReply{} }
func (m *GiftsReply) String() string { return proto.CompactTextString(m) }
func (*GiftsReply) ProtoMessage()    {}
func (*GiftsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *GiftsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GiftsReply.Unmarshal(m, b)
}
func (m *GiftsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GiftsReply.Marshal(b, m, deterministic)
}
func (m *GiftsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GiftsReply.Merge(m, src)
}
func (m *GiftsReply) XXX_Size() int {
	return xxx_messageInfo_GiftsReply.Size(m)
}
func (m *GiftsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GiftsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GiftsReply proto.InternalMessageInfo

func (m *GiftsReply) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *GiftsReply) GetGifts() string {
	if m != nil {
		return m.Gifts
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*GiftsRequest)(nil), "helloworld.GiftsRequest")
	proto.RegisterType((*GiftsReply)(nil), "helloworld.GiftsReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0xab, 0xe4, 0xc0, 0xc5, 0xe3, 0x9e, 0x99, 0x56, 0x52, 0x0c, 0x33,
	0x4b, 0x8a, 0x8b, 0xa3, 0xb0, 0x34, 0x31, 0xaf, 0x24, 0xb3, 0xa4, 0x12, 0xac, 0x94, 0x39, 0x08,
	0xce, 0x17, 0x12, 0xe1, 0x62, 0x4d, 0x07, 0xa9, 0x95, 0x60, 0x02, 0x9b, 0x01, 0xe1, 0x28, 0xd9,
	0x71, 0x71, 0x41, 0x4d, 0x00, 0xd9, 0x44, 0xb2, 0x7e, 0xa3, 0x3e, 0x46, 0x2e, 0x76, 0xf7, 0xa2,
	0xd4, 0xd4, 0x92, 0xd4, 0x22, 0x21, 0x3b, 0x2e, 0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0xc3, 0x85, 0x24,
	0xf4, 0x90, 0x02, 0x01, 0xd9, 0xbf, 0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x18, 0x84,
	0x9c, 0xb9, 0x38, 0x83, 0x53, 0xf3, 0x52, 0xc0, 0xee, 0x41, 0x35, 0x00, 0xd9, 0x93, 0xa8, 0x06,
	0x20, 0x1c, 0xaf, 0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x71, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x2e, 0x07, 0x2d, 0x85, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SendGifts(ctx context.Context, opts ...grpc.CallOption) (Greeter_SendGiftsClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendGifts(ctx context.Context, opts ...grpc.CallOption) (Greeter_SendGiftsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/helloworld.Greeter/SendGifts", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSendGiftsClient{stream}
	return x, nil
}

type Greeter_SendGiftsClient interface {
	Send(*GiftsRequest) error
	Recv() (*GiftsReply, error)
	grpc.ClientStream
}

type greeterSendGiftsClient struct {
	grpc.ClientStream
}

func (x *greeterSendGiftsClient) Send(m *GiftsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSendGiftsClient) Recv() (*GiftsReply, error) {
	m := new(GiftsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SendGifts(Greeter_SendGiftsServer) error
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendGifts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SendGifts(&greeterSendGiftsServer{stream})
}

type Greeter_SendGiftsServer interface {
	Send(*GiftsReply) error
	Recv() (*GiftsRequest, error)
	grpc.ServerStream
}

type greeterSendGiftsServer struct {
	grpc.ServerStream
}

func (x *greeterSendGiftsServer) Send(m *GiftsReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSendGiftsServer) Recv() (*GiftsRequest, error) {
	m := new(GiftsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendGifts",
			Handler:       _Greeter_SendGifts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
