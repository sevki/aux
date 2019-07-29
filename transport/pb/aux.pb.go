// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aux.proto

package pb

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

type Type int32

const (
	Type_A     Type = 0
	Type_AAAA  Type = 1
	Type_CNAME Type = 2
)

var Type_name = map[int32]string{
	0: "A",
	1: "AAAA",
	2: "CNAME",
}

var Type_value = map[string]int32{
	"A":     0,
	"AAAA":  1,
	"CNAME": 2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2316803b230a1705, []int{0}
}

type Query struct {
	Type                 Type     `protobuf:"varint,1,opt,name=type,proto3,enum=pb.Type" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_2316803b230a1705, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_A
}

func (m *Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Answer struct {
	Type                 Type     `protobuf:"varint,1,opt,name=type,proto3,enum=pb.Type" json:"type,omitempty"`
	Ttl                  int32    `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2316803b230a1705, []int{1}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_A
}

func (m *Answer) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Answer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Answer) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.Type", Type_name, Type_value)
	proto.RegisterType((*Query)(nil), "pb.Query")
	proto.RegisterType((*Answer)(nil), "pb.Answer")
}

func init() { proto.RegisterFile("aux.proto", fileDescriptor_2316803b230a1705) }

var fileDescriptor_2316803b230a1705 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xad, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xb2, 0xe4, 0x62, 0x0d, 0x2c, 0x4d,
	0x2d, 0xaa, 0x14, 0x92, 0xe1, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x33, 0xe2, 0xd0, 0x2b, 0x48, 0xd2, 0x0b, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0x8b, 0x0a, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x09,
	0x5c, 0x6c, 0x8e, 0x79, 0xc5, 0xe5, 0xa9, 0x45, 0x04, 0xf4, 0x0a, 0x70, 0x31, 0x97, 0x94, 0xe4,
	0x80, 0xb5, 0xb2, 0x06, 0x81, 0x98, 0x70, 0xd3, 0x98, 0x11, 0xa6, 0x81, 0xc4, 0x52, 0x12, 0x4b,
	0x12, 0x25, 0x58, 0x20, 0x62, 0x20, 0xb6, 0x96, 0x12, 0x17, 0x0b, 0xc8, 0x1c, 0x21, 0x56, 0x2e,
	0x46, 0x47, 0x01, 0x06, 0x21, 0x0e, 0x2e, 0x16, 0x47, 0x47, 0x47, 0x47, 0x01, 0x46, 0x21, 0x4e,
	0x2e, 0x56, 0x67, 0x3f, 0x47, 0x5f, 0x57, 0x01, 0x26, 0x23, 0x65, 0x2e, 0x66, 0xc7, 0xd2, 0x0a,
	0x21, 0x19, 0x2e, 0x66, 0x97, 0xcc, 0x74, 0x21, 0x4e, 0x90, 0xdd, 0x60, 0x0f, 0x49, 0x71, 0x81,
	0x98, 0x10, 0x07, 0x26, 0xb1, 0x81, 0x3d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x46, 0x67,
	0x0b, 0x59, 0xfd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuxClient is the client API for Aux service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuxClient interface {
	Dig(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error)
}

type auxClient struct {
	cc *grpc.ClientConn
}

func NewAuxClient(cc *grpc.ClientConn) AuxClient {
	return &auxClient{cc}
}

func (c *auxClient) Dig(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.Aux/Dig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuxServer is the server API for Aux service.
type AuxServer interface {
	Dig(context.Context, *Query) (*Answer, error)
}

// UnimplementedAuxServer can be embedded to have forward compatible implementations.
type UnimplementedAuxServer struct {
}

func (*UnimplementedAuxServer) Dig(ctx context.Context, req *Query) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dig not implemented")
}

func RegisterAuxServer(s *grpc.Server, srv AuxServer) {
	s.RegisterService(&_Aux_serviceDesc, srv)
}

func _Aux_Dig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuxServer).Dig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Aux/Dig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuxServer).Dig(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _Aux_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Aux",
	HandlerType: (*AuxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dig",
			Handler:    _Aux_Dig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aux.proto",
}
