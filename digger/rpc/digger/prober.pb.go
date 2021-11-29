// Code generated by protoc-gen-go. DO NOT EDIT.
// source: digger/prober.proto

package digger

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	strings "strings"
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
	GPList     EastMoneyType
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
	5,
	[]EastMoneyType{
		EastMoneyType_EastMoneyTypeHolder,
		EastMoneyType_EastMoneyTypeNews,
		EastMoneyType_EastMoneyTypeOperations,
		EastMoneyType_EastMoneyTypeGPList,
	},
	[]EastMoneyType{
		EastMoneyType_EastMoneyTypeUnknown,
		EastMoneyType_EastMoneyTypeHolder,
		EastMoneyType_EastMoneyTypeNews,
		EastMoneyType_EastMoneyTypeOperations,
		EastMoneyType_EastMoneyTypeGPList,
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

func (x EastMoneyType) IsGPList() bool {
	return x == EastMoneyTypeEnum.GPList
}

func (x EastMoneyType) GetGPList() EastMoneyType {
	return EastMoneyTypeEnum.GPList
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
	EastMoneyType_EastMoneyTypeGPList     EastMoneyType = 4
)

var EastMoneyType_name = map[int32]string{
	0: "EastMoneyTypeUnknown",
	1: "EastMoneyTypeHolder",
	2: "EastMoneyTypeNews",
	3: "EastMoneyTypeOperations",
	4: "EastMoneyTypeGPList",
}

var EastMoneyType_value = map[string]int32{
	"EastMoneyTypeUnknown":    0,
	"EastMoneyTypeHolder":     1,
	"EastMoneyTypeNews":       2,
	"EastMoneyTypeOperations": 3,
	"EastMoneyTypeGPList":     4,
}

func (x EastMoneyType) String() string {
	return proto.EnumName(EastMoneyType_name, int32(x))
}

func (EastMoneyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{0}
}

//权重单元
type WeightUnit struct {
	//
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value"`
	//
	Accum float64 `protobuf:"fixed64,2,opt,name=accum,proto3" json:"accum"`
	//
	Consecutive int32 `protobuf:"varint,3,opt,name=consecutive,proto3" json:"consecutive"`
	//
	Counter int32 `protobuf:"varint,4,opt,name=counter,proto3" json:"counter"`
	//
	Present float64 `protobuf:"fixed64,5,opt,name=present,proto3" json:"present"`
	//
	SubNew bool `protobuf:"varint,6,opt,name=subNew,proto3" json:"subNew"`
}

func (m *WeightUnit) Validate() error {
	return nil
}

func (m *WeightUnit) Reset()         { *m = WeightUnit{} }
func (m *WeightUnit) String() string { return proto.CompactTextString(m) }
func (*WeightUnit) ProtoMessage()    {}
func (*WeightUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{0}
}

func (m *WeightUnit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightUnit.Unmarshal(m, b)
}
func (m *WeightUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightUnit.Marshal(b, m, deterministic)
}
func (m *WeightUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightUnit.Merge(m, src)
}
func (m *WeightUnit) XXX_Size() int {
	return xxx_messageInfo_WeightUnit.Size(m)
}
func (m *WeightUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightUnit.DiscardUnknown(m)
}

var xxx_messageInfo_WeightUnit proto.InternalMessageInfo

func (m *WeightUnit) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *WeightUnit) GetAccum() float64 {
	if m != nil {
		return m.Accum
	}
	return 0
}

func (m *WeightUnit) GetConsecutive() int32 {
	if m != nil {
		return m.Consecutive
	}
	return 0
}

func (m *WeightUnit) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *WeightUnit) GetPresent() float64 {
	if m != nil {
		return m.Present
	}
	return 0
}

func (m *WeightUnit) GetSubNew() bool {
	if m != nil {
		return m.SubNew
	}
	return false
}

func init() {
	proto.RegisterEnum("digger.EastMoneyType", EastMoneyType_name, EastMoneyType_value)
	proto.RegisterType((*WeightUnit)(nil), "digger.WeightUnit")
}

func init() { proto.RegisterFile("digger/prober.proto", fileDescriptor_f34ec7a831e4a2b5) }

var fileDescriptor_f34ec7a831e4a2b5 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x46, 0x9d, 0xb6, 0x89, 0x7a, 0x8b, 0x30, 0x4e, 0xab, 0x1d, 0x70, 0x13, 0x5c, 0x05, 0x17,
	0xed, 0xc2, 0xa5, 0x3b, 0x41, 0x74, 0xa1, 0x55, 0x82, 0x45, 0x70, 0x97, 0x4c, 0x2f, 0x71, 0xb0,
	0xce, 0x0c, 0xf3, 0xd3, 0xd0, 0xa7, 0xf0, 0x2d, 0x7c, 0x4e, 0x49, 0x62, 0xc0, 0xd0, 0xe5, 0x39,
	0x07, 0xbe, 0xc5, 0x07, 0x93, 0xb5, 0x2c, 0x4b, 0xb4, 0x0b, 0x63, 0x75, 0x81, 0x76, 0x6e, 0xac,
	0xf6, 0x9a, 0xc5, 0xad, 0xbc, 0xfc, 0x21, 0x00, 0x6f, 0x28, 0xcb, 0x0f, 0xbf, 0x52, 0xd2, 0xb3,
	0x29, 0x44, 0xdb, 0x7c, 0x13, 0x90, 0x93, 0x84, 0xa4, 0x24, 0x6b, 0xa1, 0xb6, 0xb9, 0x10, 0xe1,
	0x8b, 0x0f, 0x5a, 0xdb, 0x00, 0x4b, 0x60, 0x2c, 0xb4, 0x72, 0x28, 0x82, 0x97, 0x5b, 0xe4, 0xc3,
	0x84, 0xa4, 0x51, 0xf6, 0x5f, 0x31, 0x0e, 0x87, 0x42, 0x07, 0xe5, 0xd1, 0xf2, 0x51, 0x53, 0x3b,
	0xac, 0x8b, 0xb1, 0xe8, 0x50, 0x79, 0x1e, 0x35, 0x9b, 0x1d, 0xb2, 0x73, 0x88, 0x5d, 0x28, 0x96,
	0x58, 0xf1, 0x38, 0x21, 0xe9, 0x51, 0xf6, 0x47, 0x57, 0xdf, 0x04, 0x4e, 0xee, 0x72, 0xe7, 0x9f,
	0xb4, 0xc2, 0xdd, 0xeb, 0xce, 0xd4, 0xeb, 0xd3, 0x9e, 0x58, 0xa9, 0x4f, 0xa5, 0x2b, 0x45, 0x0f,
	0xd8, 0x0c, 0x26, 0xbd, 0xf2, 0xa0, 0x37, 0x6b, 0xb4, 0x94, 0xb0, 0x33, 0x38, 0xed, 0x85, 0x25,
	0x56, 0x8e, 0x0e, 0xd8, 0x05, 0xcc, 0x7a, 0xfa, 0xd9, 0xa0, 0xcd, 0xbd, 0xd4, 0xca, 0xd1, 0xe1,
	0xde, 0xd8, 0xfd, 0xcb, 0xa3, 0x74, 0x9e, 0x8e, 0x6e, 0xc7, 0xef, 0xc7, 0xf3, 0xc5, 0x4d, 0xfb,
	0x63, 0x11, 0x37, 0xb7, 0x5e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x48, 0x4d, 0x96, 0xf9, 0x6d,
	0x01, 0x00, 0x00,
}
