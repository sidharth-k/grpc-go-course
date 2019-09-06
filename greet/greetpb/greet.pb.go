// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LGRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LGRequest) Reset()         { *m = LGRequest{} }
func (m *LGRequest) String() string { return proto.CompactTextString(m) }
func (*LGRequest) ProtoMessage()    {}
func (*LGRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{3}
}

func (m *LGRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LGRequest.Unmarshal(m, b)
}
func (m *LGRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LGRequest.Marshal(b, m, deterministic)
}
func (m *LGRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LGRequest.Merge(m, src)
}
func (m *LGRequest) XXX_Size() int {
	return xxx_messageInfo_LGRequest.Size(m)
}
func (m *LGRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LGRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LGRequest proto.InternalMessageInfo

func (m *LGRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LGResposne struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LGResposne) Reset()         { *m = LGResposne{} }
func (m *LGResposne) String() string { return proto.CompactTextString(m) }
func (*LGResposne) ProtoMessage()    {}
func (*LGResposne) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{4}
}

func (m *LGResposne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LGResposne.Unmarshal(m, b)
}
func (m *LGResposne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LGResposne.Marshal(b, m, deterministic)
}
func (m *LGResposne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LGResposne.Merge(m, src)
}
func (m *LGResposne) XXX_Size() int {
	return xxx_messageInfo_LGResposne.Size(m)
}
func (m *LGResposne) XXX_DiscardUnknown() {
	xxx_messageInfo_LGResposne.DiscardUnknown(m)
}

var xxx_messageInfo_LGResposne proto.InternalMessageInfo

func (m *LGResposne) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*LGRequest)(nil), "greet.LGRequest")
	proto.RegisterType((*LGResposne)(nil), "greet.LGResposne")
}

func init() { proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871) }

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15,
	0xcc, 0x51, 0x72, 0xe3, 0xe2, 0x70, 0x07, 0x31, 0x32, 0xf3, 0xd2, 0x85, 0x64, 0xb9, 0xb8, 0xd2,
	0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc1, 0x22, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0xd2, 0x5c, 0x9c, 0x39, 0x89, 0x30, 0x59, 0x26,
	0xb0, 0x2c, 0x07, 0x48, 0x00, 0x24, 0xa9, 0x64, 0xcd, 0xc5, 0x03, 0x36, 0x27, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x9b, 0x8b, 0x23, 0x1d, 0x6a, 0x2e, 0xd8, 0x24, 0x6e, 0x23, 0x7e,
	0x3d, 0x88, 0xf5, 0x30, 0xeb, 0x82, 0xe0, 0x0a, 0x94, 0xd4, 0xb9, 0x78, 0xa1, 0x9a, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0xa0, 0xae,
	0x80, 0xf2, 0x94, 0x2c, 0xb8, 0x38, 0x7d, 0xdc, 0xc9, 0xb2, 0x42, 0x85, 0x8b, 0x0b, 0xa4, 0xb3,
	0xb8, 0x20, 0xbf, 0x38, 0x0f, 0xa7, 0xf9, 0x46, 0x55, 0x50, 0x5f, 0x04, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x0a, 0x99, 0x70, 0xb1, 0x82, 0xf9, 0x42, 0xc2, 0xc8, 0x26, 0x43, 0x1d, 0x20, 0x25,
	0x82, 0x2a, 0x08, 0x71, 0xbb, 0x12, 0x83, 0x90, 0x09, 0x17, 0xa7, 0x4f, 0x7e, 0x5e, 0x3a, 0x44,
	0xa7, 0x00, 0x54, 0x11, 0xdc, 0xdd, 0x52, 0x82, 0x48, 0x22, 0x10, 0xf7, 0x28, 0x31, 0x68, 0x30,
	0x3a, 0x71, 0x46, 0xb1, 0x43, 0xe3, 0x29, 0x89, 0x0d, 0x1c, 0x45, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x11, 0xe9, 0x5b, 0x35, 0xbf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//client streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LGRequest) error
	CloseAndRecv() (*LGResposne, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LGRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LGResposne, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LGResposne)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	//client streaming
	LongGreet(GreetService_LongGreetServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) LongGreet(srv GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LGResposne) error
	Recv() (*LGRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LGResposne) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LGRequest, error) {
	m := new(LGRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
