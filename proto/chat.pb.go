// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type Message struct {
	// Types that are valid to be assigned to Msg:
	//	*Message_Text
	//	*Message_Join
	Msg                  isMessage_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_7c8c3af5b2228ffc, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_Text struct {
	Text *Text `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type Message_Join struct {
	Join *Join `protobuf:"bytes,3,opt,name=join,proto3,oneof"`
}

func (*Message_Text) isMessage_Msg() {}

func (*Message_Join) isMessage_Msg() {}

func (m *Message) GetMsg() isMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Message) GetText() *Text {
	if x, ok := m.GetMsg().(*Message_Text); ok {
		return x.Text
	}
	return nil
}

func (m *Message) GetJoin() *Join {
	if x, ok := m.GetMsg().(*Message_Join); ok {
		return x.Join
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Text)(nil),
		(*Message_Join)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// msg
	switch x := m.Msg.(type) {
	case *Message_Text:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *Message_Join:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Join); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Msg has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // msg.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Text)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Text{msg}
		return true, err
	case 3: // msg.join
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Join)
		err := b.DecodeMessage(msg)
		m.Msg = &Message_Join{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// msg
	switch x := m.Msg.(type) {
	case *Message_Text:
		s := proto.Size(x.Text)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Join:
		s := proto.Size(x.Join)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Event struct {
	Time *timestamp.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Who  string               `protobuf:"bytes,2,opt,name=who,proto3" json:"who,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*Event_Text
	//	*Event_Join
	//	*Event_Leave
	Msg                  isEvent_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_7c8c3af5b2228ffc, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Event) GetWho() string {
	if m != nil {
		return m.Who
	}
	return ""
}

type isEvent_Msg interface {
	isEvent_Msg()
}

type Event_Text struct {
	Text *Text `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type Event_Join struct {
	Join *Join `protobuf:"bytes,4,opt,name=join,proto3,oneof"`
}

type Event_Leave struct {
	Leave *Leave `protobuf:"bytes,5,opt,name=leave,proto3,oneof"`
}

func (*Event_Text) isEvent_Msg() {}

func (*Event_Join) isEvent_Msg() {}

func (*Event_Leave) isEvent_Msg() {}

func (m *Event) GetMsg() isEvent_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Event) GetText() *Text {
	if x, ok := m.GetMsg().(*Event_Text); ok {
		return x.Text
	}
	return nil
}

func (m *Event) GetJoin() *Join {
	if x, ok := m.GetMsg().(*Event_Join); ok {
		return x.Join
	}
	return nil
}

func (m *Event) GetLeave() *Leave {
	if x, ok := m.GetMsg().(*Event_Leave); ok {
		return x.Leave
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Text)(nil),
		(*Event_Join)(nil),
		(*Event_Leave)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// msg
	switch x := m.Msg.(type) {
	case *Event_Text:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *Event_Join:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Join); err != nil {
			return err
		}
	case *Event_Leave:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leave); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Msg has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 3: // msg.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Text)
		err := b.DecodeMessage(msg)
		m.Msg = &Event_Text{msg}
		return true, err
	case 4: // msg.join
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Join)
		err := b.DecodeMessage(msg)
		m.Msg = &Event_Join{msg}
		return true, err
	case 5: // msg.leave
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Leave)
		err := b.DecodeMessage(msg)
		m.Msg = &Event_Leave{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// msg
	switch x := m.Msg.(type) {
	case *Event_Text:
		s := proto.Size(x.Text)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Join:
		s := proto.Size(x.Join)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Leave:
		s := proto.Size(x.Leave)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_7c8c3af5b2228ffc, []int{2}
}
func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (dst *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(dst, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Join struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pass                 string   `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Join) Reset()         { *m = Join{} }
func (m *Join) String() string { return proto.CompactTextString(m) }
func (*Join) ProtoMessage()    {}
func (*Join) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_7c8c3af5b2228ffc, []int{3}
}
func (m *Join) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Join.Unmarshal(m, b)
}
func (m *Join) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Join.Marshal(b, m, deterministic)
}
func (dst *Join) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Join.Merge(dst, src)
}
func (m *Join) XXX_Size() int {
	return xxx_messageInfo_Join.Size(m)
}
func (m *Join) XXX_DiscardUnknown() {
	xxx_messageInfo_Join.DiscardUnknown(m)
}

var xxx_messageInfo_Join proto.InternalMessageInfo

func (m *Join) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Join) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type Leave struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Leave) Reset()         { *m = Leave{} }
func (m *Leave) String() string { return proto.CompactTextString(m) }
func (*Leave) ProtoMessage()    {}
func (*Leave) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_7c8c3af5b2228ffc, []int{4}
}
func (m *Leave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Leave.Unmarshal(m, b)
}
func (m *Leave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Leave.Marshal(b, m, deterministic)
}
func (dst *Leave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leave.Merge(dst, src)
}
func (m *Leave) XXX_Size() int {
	return xxx_messageInfo_Leave.Size(m)
}
func (m *Leave) XXX_DiscardUnknown() {
	xxx_messageInfo_Leave.DiscardUnknown(m)
}

var xxx_messageInfo_Leave proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*Event)(nil), "proto.Event")
	proto.RegisterType((*Text)(nil), "proto.Text")
	proto.RegisterType((*Join)(nil), "proto.Join")
	proto.RegisterType((*Leave)(nil), "proto.Leave")
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
	Messages(ctx context.Context, opts ...grpc.CallOption) (Chat_MessagesClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Messages(ctx context.Context, opts ...grpc.CallOption) (Chat_MessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/proto.Chat/Messages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatMessagesClient{stream}
	return x, nil
}

type Chat_MessagesClient interface {
	Send(*Message) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type chatMessagesClient struct {
	grpc.ClientStream
}

func (x *chatMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatMessagesClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Messages(Chat_MessagesServer) error
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Messages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Messages(&chatMessagesServer{stream})
}

type Chat_MessagesServer interface {
	Send(*Event) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatMessagesServer struct {
	grpc.ServerStream
}

func (x *chatMessagesServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Messages",
			Handler:       _Chat_Messages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_7c8c3af5b2228ffc) }

var fileDescriptor_chat_7c8c3af5b2228ffc = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x1b, 0x9b, 0xb8, 0xee, 0xac, 0x88, 0xe4, 0x54, 0x7a, 0x51, 0x83, 0x87, 0x3d, 0x65,
	0x65, 0x05, 0x1f, 0x40, 0x11, 0x16, 0xd1, 0x4b, 0x58, 0xbc, 0x67, 0x25, 0xb6, 0x95, 0x6d, 0x53,
	0x4c, 0x5c, 0xf7, 0xe1, 0x7c, 0x38, 0x99, 0x69, 0x2a, 0x78, 0x10, 0x4f, 0x9d, 0xce, 0xff, 0x75,
	0xe6, 0xeb, 0x00, 0xbc, 0xd4, 0x36, 0xea, 0xfe, 0xdd, 0x47, 0x2f, 0x05, 0x3d, 0xca, 0xb3, 0xca,
	0xfb, 0x6a, 0xeb, 0x16, 0xf4, 0xb6, 0xf9, 0x78, 0x5d, 0xc4, 0xa6, 0x75, 0x21, 0xda, 0xb6, 0x1f,
	0x38, 0xf5, 0x0c, 0x93, 0x27, 0x17, 0x82, 0xad, 0x9c, 0xbc, 0x00, 0x1e, 0xdd, 0x3e, 0x16, 0x07,
	0xe7, 0x6c, 0x3e, 0x5b, 0xce, 0x06, 0x40, 0xaf, 0xdd, 0x3e, 0xae, 0x32, 0x43, 0x11, 0x22, 0x6f,
	0xbe, 0xe9, 0x8a, 0xfc, 0x17, 0xf2, 0xe0, 0x9b, 0x0e, 0x11, 0x8c, 0x6e, 0x05, 0xe4, 0x6d, 0xa8,
	0xd4, 0x17, 0x03, 0x71, 0xbf, 0x73, 0x5d, 0x94, 0x1a, 0x38, 0x2e, 0x2d, 0x18, 0x7d, 0x53, 0xea,
	0xc1, 0x48, 0x8f, 0x46, 0x7a, 0x3d, 0x1a, 0x19, 0xe2, 0xe4, 0x29, 0xe4, 0x9f, 0xb5, 0x27, 0x8b,
	0xa9, 0xc1, 0xf2, 0x47, 0x2c, 0xff, 0x5f, 0x8c, 0xff, 0x29, 0x26, 0x2f, 0x41, 0x6c, 0x9d, 0xdd,
	0xb9, 0x42, 0x10, 0x73, 0x9c, 0x98, 0x47, 0xec, 0xad, 0x32, 0x33, 0x84, 0xa3, 0x7e, 0x09, 0x1c,
	0xe7, 0x4b, 0x99, 0x56, 0x33, 0xb2, 0xa1, 0x5a, 0x69, 0xe0, 0x38, 0x18, 0xb3, 0xce, 0xa6, 0x1f,
	0x9b, 0x1a, 0xaa, 0xb1, 0xd7, 0xdb, 0x10, 0x92, 0x3d, 0xd5, 0x6a, 0x02, 0x82, 0x96, 0x2c, 0x6f,
	0x80, 0xdf, 0xd5, 0x16, 0x2f, 0x72, 0x94, 0x6e, 0x1e, 0xe4, 0x49, 0xd2, 0x48, 0x8d, 0x72, 0xd4,
	0xa2, 0xdb, 0xa9, 0x6c, 0xce, 0xae, 0xd8, 0xe6, 0x90, 0x5a, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xe5, 0xc4, 0x0a, 0xe0, 0x01, 0x00, 0x00,
}
