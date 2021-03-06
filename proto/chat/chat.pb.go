// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

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

// Request 请求数据格式
type Request struct {
	MainId               int32    `protobuf:"varint,1,opt,name=mainId,proto3" json:"mainId,omitempty"`
	SubId                int32    `protobuf:"varint,2,opt,name=subId,proto3" json:"subId,omitempty"`
	RequestId            int32    `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Input                string   `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMainId() int32 {
	if m != nil {
		return m.MainId
	}
	return 0
}

func (m *Request) GetSubId() int32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

func (m *Request) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Request) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

// Response 响应数据格式
type Response struct {
	MainId               int32    `protobuf:"varint,1,opt,name=mainId,proto3" json:"mainId,omitempty"`
	SubId                int32    `protobuf:"varint,2,opt,name=subId,proto3" json:"subId,omitempty"`
	RequestId            int32    `protobuf:"varint,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Output               string   `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMainId() int32 {
	if m != nil {
		return m.MainId
	}
	return 0
}

func (m *Response) GetSubId() int32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

func (m *Response) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "chat.Request")
	proto.RegisterType((*Response)(nil), "chat.Response")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xb2, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x72, 0x13, 0x33, 0xf3, 0x3c, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x21, 0x11, 0x2e, 0xd6, 0xe2, 0xd2, 0x24, 0xcf,
	0x14, 0x09, 0x26, 0xb0, 0x30, 0x84, 0x23, 0x24, 0xc3, 0xc5, 0x59, 0x04, 0xd1, 0xe8, 0x99, 0x22,
	0xc1, 0x0c, 0x96, 0x41, 0x08, 0x80, 0xf4, 0x64, 0xe6, 0x15, 0x94, 0x96, 0x48, 0xb0, 0x28, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x79, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x54, 0xb5, 0x4d, 0x8c, 0x8b, 0x2d, 0xbf, 0xb4, 0x04, 0x61, 0x1d, 0x94, 0x67, 0x64, 0xc1,
	0xc5, 0xe2, 0x9c, 0x91, 0x58, 0x22, 0x64, 0xc0, 0xc5, 0xe9, 0x94, 0x99, 0x12, 0x5c, 0x52, 0x94,
	0x9a, 0x98, 0x2b, 0xc4, 0xab, 0x07, 0x0e, 0x04, 0xa8, 0xaf, 0xa5, 0xf8, 0x60, 0x5c, 0x88, 0xbb,
	0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x61, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x13, 0xb7, 0x0b, 0x1d, 0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	BidStream(ctx context.Context, opts ...grpc.CallOption) (Chat_BidStreamClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) BidStream(ctx context.Context, opts ...grpc.CallOption) (Chat_BidStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/BidStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBidStreamClient{stream}
	return x, nil
}

type Chat_BidStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type chatBidStreamClient struct {
	grpc.ClientStream
}

func (x *chatBidStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatBidStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	BidStream(Chat_BidStreamServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_BidStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).BidStream(&chatBidStreamServer{stream})
}

type Chat_BidStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type chatBidStreamServer struct {
	grpc.ServerStream
}

func (x *chatBidStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatBidStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidStream",
			Handler:       _Chat_BidStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
