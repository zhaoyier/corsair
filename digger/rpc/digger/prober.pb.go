// Code generated by protoc-gen-go. DO NOT EDIT.
// source: digger/prober.proto

package digger

import (
	context1 "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc1 "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	strings "strings"
)

import (
	mservice "git.ezbuy.me/ezbuy/base/dist/mservice"
	ezcommon "git.ezbuy.me/ezbuy/base/dist/proto/common"
	context "git.ezbuy.me/ezbuy/base/misc/context"
	"net/http"
	"sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = strings.TrimPrefix

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type _EastMoneyTypeEnum struct {
	Unknown    EastMoneyType
	Holder     EastMoneyType
	News       EastMoneyType
	Operations EastMoneyType
	TotalSize  int
	List       []EastMoneyType
	ZeroList   []EastMoneyType
}

var EastMoneyTypeEnum = _EastMoneyTypeEnum{
	0,
	1,
	2,
	3,
	4,
	[]EastMoneyType{
		EastMoneyType_EastMoneyTypeHolder,
		EastMoneyType_EastMoneyTypeNews,
		EastMoneyType_EastMoneyTypeOperations,
	},
	[]EastMoneyType{
		EastMoneyType_EastMoneyTypeUnknown,
		EastMoneyType_EastMoneyTypeHolder,
		EastMoneyType_EastMoneyTypeNews,
		EastMoneyType_EastMoneyTypeOperations,
	},
}

func (x *_EastMoneyTypeEnum) Parse(short string) EastMoneyType {
	return EastMoneyType(EastMoneyType_value["EastMoneyType"+strings.TrimPrefix(short, "EastMoneyType")])
}

func (x EastMoneyType) IsUnknown() bool {
	return x == EastMoneyTypeEnum.Unknown
}

func (x EastMoneyType) GetUnknown() EastMoneyType {
	return EastMoneyTypeEnum.Unknown
}

func (x EastMoneyType) IsHolder() bool {
	return x == EastMoneyTypeEnum.Holder
}

func (x EastMoneyType) GetHolder() EastMoneyType {
	return EastMoneyTypeEnum.Holder
}

func (x EastMoneyType) IsNews() bool {
	return x == EastMoneyTypeEnum.News
}

func (x EastMoneyType) GetNews() EastMoneyType {
	return EastMoneyTypeEnum.News
}

func (x EastMoneyType) IsOperations() bool {
	return x == EastMoneyTypeEnum.Operations
}

func (x EastMoneyType) GetOperations() EastMoneyType {
	return EastMoneyTypeEnum.Operations
}

func (x EastMoneyType) Valid() bool {
	if x == EastMoneyType_EastMoneyTypeUnknown {
		return false
	}
	return x.ZeroValid()
}

func (x EastMoneyType) ZeroValid() bool {
	_, ok := EastMoneyType_name[int32(x)]
	return ok
}
func (x EastMoneyType) Short() string {
	n := x.String()
	typ := "EastMoneyType"
	if len(n) > len(typ) {
		if n[:len(typ)] == typ {
			return n[len(typ):]
		}
	}
	return n

}
func (x *EastMoneyType) SetValue(v int32) {
	*x = EastMoneyType(v)
}

func (x EastMoneyType) GetValue() int32 {
	return int32(x)
}

func (x *EastMoneyType) Type() string {
	return "EastMoneyType"
}

type EastMoneyType int32

const (
	EastMoneyType_EastMoneyTypeUnknown    EastMoneyType = 0
	EastMoneyType_EastMoneyTypeHolder     EastMoneyType = 1
	EastMoneyType_EastMoneyTypeNews       EastMoneyType = 2
	EastMoneyType_EastMoneyTypeOperations EastMoneyType = 3
)

var EastMoneyType_name = map[int32]string{
	0: "EastMoneyTypeUnknown",
	1: "EastMoneyTypeHolder",
	2: "EastMoneyTypeNews",
	3: "EastMoneyTypeOperations",
}

var EastMoneyType_value = map[string]int32{
	"EastMoneyTypeUnknown":    0,
	"EastMoneyTypeHolder":     1,
	"EastMoneyTypeNews":       2,
	"EastMoneyTypeOperations": 3,
}

func (x EastMoneyType) String() string {
	return proto.EnumName(EastMoneyType_name, int32(x))
}

func (EastMoneyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{0}
}

type ShareholderResp struct {
	//
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (m *ShareholderResp) Validate() error {
	return nil
}

func (m *ShareholderResp) Reset()         { *m = ShareholderResp{} }
func (m *ShareholderResp) String() string { return proto.CompactTextString(m) }
func (*ShareholderResp) ProtoMessage()    {}
func (*ShareholderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{0}
}

func (m *ShareholderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareholderResp.Unmarshal(m, b)
}
func (m *ShareholderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareholderResp.Marshal(b, m, deterministic)
}
func (m *ShareholderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareholderResp.Merge(m, src)
}
func (m *ShareholderResp) XXX_Size() int {
	return xxx_messageInfo_ShareholderResp.Size(m)
}
func (m *ShareholderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareholderResp.DiscardUnknown(m)
}

var xxx_messageInfo_ShareholderResp proto.InternalMessageInfo

func (m *ShareholderResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNameReq struct {
}

func (m *GetNameReq) Validate() error {
	return nil
}

func (m *GetNameReq) Reset()         { *m = GetNameReq{} }
func (m *GetNameReq) String() string { return proto.CompactTextString(m) }
func (*GetNameReq) ProtoMessage()    {}
func (*GetNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{1}
}

func (m *GetNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNameReq.Unmarshal(m, b)
}
func (m *GetNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNameReq.Marshal(b, m, deterministic)
}
func (m *GetNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNameReq.Merge(m, src)
}
func (m *GetNameReq) XXX_Size() int {
	return xxx_messageInfo_GetNameReq.Size(m)
}
func (m *GetNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetNameReq proto.InternalMessageInfo

type GetNameResp struct {
}

func (m *GetNameResp) Validate() error {
	return nil
}

func (m *GetNameResp) Reset()         { *m = GetNameResp{} }
func (m *GetNameResp) String() string { return proto.CompactTextString(m) }
func (*GetNameResp) ProtoMessage()    {}
func (*GetNameResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{2}
}

func (m *GetNameResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNameResp.Unmarshal(m, b)
}
func (m *GetNameResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNameResp.Marshal(b, m, deterministic)
}
func (m *GetNameResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNameResp.Merge(m, src)
}
func (m *GetNameResp) XXX_Size() int {
	return xxx_messageInfo_GetNameResp.Size(m)
}
func (m *GetNameResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNameResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetNameResp proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("digger.EastMoneyType", EastMoneyType_name, EastMoneyType_value)
	proto.RegisterType((*ShareholderResp)(nil), "digger.ShareholderResp")
	proto.RegisterType((*GetNameReq)(nil), "digger.GetNameReq")
	proto.RegisterType((*GetNameResp)(nil), "digger.GetNameResp")
}

func init() { proto.RegisterFile("digger/prober.proto", fileDescriptor_f34ec7a831e4a2b5) }

var fileDescriptor_f34ec7a831e4a2b5 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xc9, 0x4c, 0x4f,
	0x4f, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0xa9, 0x72, 0xf1, 0x07, 0x67, 0x24, 0x16, 0xa5, 0x66, 0xe4, 0xe7, 0xa4,
	0xa4, 0x16, 0x05, 0xa5, 0x16, 0x17, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x3c, 0x5c, 0x5c, 0xee, 0xa9, 0x25, 0x7e, 0x89,
	0xb9, 0xa9, 0x41, 0xa9, 0x85, 0x4a, 0xbc, 0x5c, 0xdc, 0x70, 0x5e, 0x71, 0x81, 0x56, 0x19, 0x17,
	0xaf, 0x6b, 0x62, 0x71, 0x89, 0x6f, 0x7e, 0x5e, 0x6a, 0x65, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x04,
	0x97, 0x08, 0x8a, 0x40, 0x68, 0x5e, 0x76, 0x5e, 0x7e, 0x79, 0x9e, 0x00, 0x83, 0x90, 0x38, 0x97,
	0x30, 0x8a, 0x8c, 0x07, 0xd8, 0x5a, 0x01, 0x46, 0x21, 0x51, 0x2e, 0x41, 0x14, 0x09, 0xbf, 0xd4,
	0xf2, 0x62, 0x01, 0x26, 0x21, 0x69, 0x2e, 0x71, 0x14, 0x61, 0xff, 0x82, 0xd4, 0xa2, 0xc4, 0x92,
	0xcc, 0xfc, 0xbc, 0x62, 0x01, 0x66, 0x23, 0x1b, 0x2e, 0xb6, 0x00, 0xb0, 0x9f, 0x84, 0x8c, 0xb8,
	0xd8, 0xa1, 0x0e, 0x12, 0x12, 0xd2, 0x83, 0xf8, 0x4c, 0x0f, 0xe1, 0x5e, 0x29, 0x61, 0x0c, 0xb1,
	0xe2, 0x82, 0x24, 0x36, 0x70, 0x40, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x74, 0xf3, 0x00,
	0x44, 0x1f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context1.Context
var _ grpc1.ClientConn
var _ http.Client
var _ ezcommon.Empty

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion6

const ServiceProber = "digger.Prober"

var ServiceMethodProber = []string{
	"GetName",
}

// Client API for Prober service

// ProberClient is the client API for Prober service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProberClient interface {
	GetName(ctx context.T, in *GetNameReq, opts ...grpc1.CallOption) (*GetNameResp, error)
}

type proberClient struct {
	cc *grpc1.ClientConn
}

var proberClientOnce sync.Once
var proberClientInstance ProberClient

func GetProber() ProberClient {
	proberClientOnce.Do(func() {
		var err error
		proberClientInstance, err = NewProber()
		if err != nil {
			panic(err)
		}
	})
	return proberClientInstance
}

func NewProber() (ProberClient, error) {
	cfg := mservice.DefaultClientConfig()
	cfg.Desc = OptionProber
	cli, err := mservice.NewClientEx(ServiceProber, cfg)
	if err != nil {
		return nil, err
	}
	return NewProberClient(cli.GRPC), nil
}

func NewProberClient(cc *grpc1.ClientConn) ProberClient {
	return &proberClient{cc}
}

func (c *proberClient) GetName(ctx context.T, in *GetNameReq, opts ...grpc1.CallOption) (*GetNameResp, error) {
	out := new(GetNameResp)
	err := c.cc.Invoke(ctx, "/digger.Prober/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProberServer is the server API for Prober service.
type ProberServer interface {
	GetName(context.T, *GetNameReq) (*GetNameResp, error)
}

func RegisterProberGrpc(s ProberGrpcRegister) {
	s.RegisterService(ServiceProber, ServiceMethodProber)
	RegisterProberServer(s.GetServer(), s)
}

func RegisterProberGrpcImpl(s mservice.GrpcRegister, impl ProberServer) {
	s.RegisterService(ServiceProber, ServiceMethodProber)
	RegisterProberServer(s.GetServer(), impl)
}

// UnimplementedProberServer can be embedded to have forward compatible implementations.
type UnimplementedProberServer struct {
}

func (*UnimplementedProberServer) GetName(ctx context1.Context, req *GetNameReq) (*GetNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}

func RegisterProberServer(s *grpc1.Server, srv ProberServer) {
	s.RegisterService(&_Prober_serviceDesc, srv)
}

type ProberMServicer interface {
	RegisterServiceDesc(*ezcommon.ServiceOpDesc)
	ProberServer
	mservice.WebApiRegister
	mservice.GrpcRegister
	mservice.WebSocketRegister
}

type ProberGrpcRegister interface {
	ProberServer
	mservice.GrpcRegister
}

func RegisterProber(s ProberMServicer) {
	s.RegisterServiceDesc(OptionProber)
	RegisterProberGrpc(s)
	RegisterProberWebapiEx(s)
	RegisterProberWebSocket(s)
}

type ProberWebApiRegister interface {
	ProberServer
	mservice.WebApiRegister
}

func RegisterProberWebApi(s ProberWebApiRegister) {
	RegisterProberWebApiImpl(s, s)
}

func RegisterProberWebApiImpl(s mservice.WebApiRegister, impl ProberServer) {
	wrap := &ProberWebapi{server: impl, register: s}
	_ = wrap
	s.WebApiRegister("/api/digger.Prober/GetName", wrap.GetName)
}

func RegisterProberWebapiEx(s ProberWebApiRegister) {
	wrap := &ProberWebapi{server: s, register: s}
	_ = wrap
	s.WebApiRegisterMethod("digger.Prober", "GetName", wrap.GetName)
}

type ProberWebapi struct {
	server   ProberServer
	register mservice.WebApiRegister
}

func (s *ProberWebapi) GetName(ctx *context.T, w http.ResponseWriter, req *http.Request) {
	params := new(GetNameReq)
	if err := s.register.WebApiDecode(ctx, req, params); err != nil {
		s.register.WebApiHandleResp(ctx, w, nil, err)
		return
	}
	resp, err := s.server.GetName(*ctx, params)
	s.register.WebApiHandleResp(ctx, w, resp, err)
}

type ProberWebSocketApi interface {
}

type ProberWebSocketRegister interface {
	ProberWebSocketApi
	mservice.WebSocketRegister
}

func RegisterProberWebSocket(s ProberWebSocketRegister) {
}

type ProberWebSocket struct {
	api      ProberWebSocketApi
	register mservice.WebSocketRegister
}

func _Prober_GetName_Handler(srv interface{}, ctx context1.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if err := in.Validate(); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProberServer).GetName(context.From(ctx), in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/digger.Prober/GetName",
	}
	handler := func(ctx context1.Context, req interface{}) (interface{}, error) {
		return srv.(ProberServer).GetName(context.From(ctx), req.(*GetNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Prober_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "digger.Prober",
	HandlerType: (*ProberServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _Prober_GetName_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "digger/prober.proto",
}

var OptionProber = ezcommon.GenOption([]byte{
	// 26 bytes of Option
	0x0a, 0x0d, 0x64, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x1a,
	0x09, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
})
