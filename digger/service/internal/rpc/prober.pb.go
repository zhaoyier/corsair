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

type _RMTypeEnum struct {
	RmTypeUnknown RMType
	RmTypeShort   RMType
	RmTypeLong    RMType
	TotalSize     int
	List          []RMType
	ZeroList      []RMType
}

var RMTypeEnum = _RMTypeEnum{
	0,
	1,
	2,
	3,
	[]RMType{
		RMType_RmTypeShort,
		RMType_RmTypeLong,
	},
	[]RMType{
		RMType_RmTypeUnknown,
		RMType_RmTypeShort,
		RMType_RmTypeLong,
	},
}

func (x *_RMTypeEnum) Parse(short string) RMType {
	return RMType(RMType_value["RMType"+strings.TrimPrefix(short, "RMType")])
}

func (x RMType) IsRmTypeUnknown() bool {
	return x == RMTypeEnum.RmTypeUnknown
}

func (x RMType) BitHasRmTypeUnknown() bool {
	return x.BitHas(RMTypeEnum.RmTypeUnknown)
}

func (x RMType) BitOrRmTypeUnknown() RMType {
	x.BitOr(RMTypeEnum.RmTypeUnknown)
	return x
}

func (x RMType) GetRmTypeUnknown() RMType {
	return RMTypeEnum.RmTypeUnknown
}

func (x RMType) IsRmTypeShort() bool {
	return x == RMTypeEnum.RmTypeShort
}

func (x RMType) BitHasRmTypeShort() bool {
	return x.BitHas(RMTypeEnum.RmTypeShort)
}

func (x RMType) BitOrRmTypeShort() RMType {
	x.BitOr(RMTypeEnum.RmTypeShort)
	return x
}

func (x RMType) GetRmTypeShort() RMType {
	return RMTypeEnum.RmTypeShort
}

func (x RMType) IsRmTypeLong() bool {
	return x == RMTypeEnum.RmTypeLong
}

func (x RMType) BitHasRmTypeLong() bool {
	return x.BitHas(RMTypeEnum.RmTypeLong)
}

func (x RMType) BitOrRmTypeLong() RMType {
	x.BitOr(RMTypeEnum.RmTypeLong)
	return x
}

func (x RMType) GetRmTypeLong() RMType {
	return RMTypeEnum.RmTypeLong
}

func (x RMType) Valid() bool {
	if x == RMType_RmTypeUnknown {
		return false
	}
	return x.ZeroValid()
}

func (x RMType) ZeroValid() bool {
	_, ok := RMType_name[int32(x)]
	return ok
}
func (x RMType) Short() string {
	n := x.String()
	typ := "RMType"
	if len(n) > len(typ) {
		if n[:len(typ)] == typ {
			return n[len(typ):]
		}
	}
	return n

}
func (x RMType) BitString() []string {
	name := make([]string, 0, len(RMTypeEnum.List))
	for _, item := range RMTypeEnum.List {
		if x.BitHas(item) {
			name = append(name, item.Short())
		}
	}
	return name

}

func (x RMType) BitValid() bool {
	return x.BitHas(RMTypeEnum.List...)
}
func (x *RMType) BitOr(item RMType) {
	*x |= item
}
func (x RMType) BitHas(xs ...RMType) bool {
	for _, item := range xs {
		if x&item > 0 {
			return true
		}
	}
	return false
}

func (x *RMType) SetValue(v int32) {
	*x = RMType(v)
}

func (x RMType) GetValue() int32 {
	return int32(x)
}

func (x *RMType) Type() string {
	return "RMType"
}

type RMType int32

const (
	RMType_RmTypeUnknown RMType = 0
	RMType_RmTypeShort   RMType = 1
	RMType_RmTypeLong    RMType = 2
)

var RMType_name = map[int32]string{
	0: "RmTypeUnknown",
	1: "RmTypeShort",
	2: "RmTypeLong",
}

var RMType_value = map[string]int32{
	"RmTypeUnknown": 0,
	"RmTypeShort":   1,
	"RmTypeLong":    2,
}

func (x RMType) String() string {
	return proto.EnumName(RMType_name, int32(x))
}

func (RMType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{1}
}

//1准备,2开始, 3进行中,4结束,5放弃
type _RMStateEnum struct {
	Unknown   RMState
	Prepared  RMState
	Started   RMState
	First     RMState
	Second    RMState
	Third     RMState
	Over      RMState
	Abandoned RMState
	TotalSize int
	List      []RMState
	ZeroList  []RMState
}

var RMStateEnum = _RMStateEnum{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	[]RMState{
		RMState_RMStatePrepared,
		RMState_RMStateStarted,
		RMState_RMStateFirst,
		RMState_RMStateSecond,
		RMState_RMStateThird,
		RMState_RMStateOver,
		RMState_RMStateAbandoned,
	},
	[]RMState{
		RMState_RMStateUnknown,
		RMState_RMStatePrepared,
		RMState_RMStateStarted,
		RMState_RMStateFirst,
		RMState_RMStateSecond,
		RMState_RMStateThird,
		RMState_RMStateOver,
		RMState_RMStateAbandoned,
	},
}

func (x *_RMStateEnum) Parse(short string) RMState {
	return RMState(RMState_value["RMState"+strings.TrimPrefix(short, "RMState")])
}

func (x RMState) IsUnknown() bool {
	return x == RMStateEnum.Unknown
}

func (x RMState) GetUnknown() RMState {
	return RMStateEnum.Unknown
}

func (x RMState) IsPrepared() bool {
	return x == RMStateEnum.Prepared
}

func (x RMState) GetPrepared() RMState {
	return RMStateEnum.Prepared
}

func (x RMState) IsStarted() bool {
	return x == RMStateEnum.Started
}

func (x RMState) GetStarted() RMState {
	return RMStateEnum.Started
}

func (x RMState) IsFirst() bool {
	return x == RMStateEnum.First
}

func (x RMState) GetFirst() RMState {
	return RMStateEnum.First
}

func (x RMState) IsSecond() bool {
	return x == RMStateEnum.Second
}

func (x RMState) GetSecond() RMState {
	return RMStateEnum.Second
}

func (x RMState) IsThird() bool {
	return x == RMStateEnum.Third
}

func (x RMState) GetThird() RMState {
	return RMStateEnum.Third
}

func (x RMState) IsOver() bool {
	return x == RMStateEnum.Over
}

func (x RMState) GetOver() RMState {
	return RMStateEnum.Over
}

func (x RMState) IsAbandoned() bool {
	return x == RMStateEnum.Abandoned
}

func (x RMState) GetAbandoned() RMState {
	return RMStateEnum.Abandoned
}

func (x RMState) Valid() bool {
	if x == RMState_RMStateUnknown {
		return false
	}
	return x.ZeroValid()
}

func (x RMState) ZeroValid() bool {
	_, ok := RMState_name[int32(x)]
	return ok
}
func (x RMState) Short() string {
	n := x.String()
	typ := "RMState"
	if len(n) > len(typ) {
		if n[:len(typ)] == typ {
			return n[len(typ):]
		}
	}
	return n

}
func (x *RMState) SetValue(v int32) {
	*x = RMState(v)
}

func (x RMState) GetValue() int32 {
	return int32(x)
}

func (x *RMState) Type() string {
	return "RMState"
}

type RMState int32

const (
	RMState_RMStateUnknown   RMState = 0
	RMState_RMStatePrepared  RMState = 1
	RMState_RMStateStarted   RMState = 2
	RMState_RMStateFirst     RMState = 3
	RMState_RMStateSecond    RMState = 4
	RMState_RMStateThird     RMState = 5
	RMState_RMStateOver      RMState = 6
	RMState_RMStateAbandoned RMState = 7
)

var RMState_name = map[int32]string{
	0: "RMStateUnknown",
	1: "RMStatePrepared",
	2: "RMStateStarted",
	3: "RMStateFirst",
	4: "RMStateSecond",
	5: "RMStateThird",
	6: "RMStateOver",
	7: "RMStateAbandoned",
}

var RMState_value = map[string]int32{
	"RMStateUnknown":   0,
	"RMStatePrepared":  1,
	"RMStateStarted":   2,
	"RMStateFirst":     3,
	"RMStateSecond":    4,
	"RMStateThird":     5,
	"RMStateOver":      6,
	"RMStateAbandoned": 7,
}

func (x RMState) String() string {
	return proto.EnumName(RMState_name, int32(x))
}

func (RMState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{2}
}

type _FunctionTypeEnum struct {
	Unknown     FunctionType
	CodeList    FunctionType
	Shareholder FunctionType
	LongLine    FunctionType
	ShortLine   FunctionType
	Recommend   FunctionType
	TotalSize   int
	List        []FunctionType
	ZeroList    []FunctionType
}

var FunctionTypeEnum = _FunctionTypeEnum{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	[]FunctionType{
		FunctionType_FunctionTypeCodeList,
		FunctionType_FunctionTypeShareholder,
		FunctionType_FunctionTypeLongLine,
		FunctionType_FunctionTypeShortLine,
		FunctionType_FunctionTypeRecommend,
	},
	[]FunctionType{
		FunctionType_FunctionTypeUnknown,
		FunctionType_FunctionTypeCodeList,
		FunctionType_FunctionTypeShareholder,
		FunctionType_FunctionTypeLongLine,
		FunctionType_FunctionTypeShortLine,
		FunctionType_FunctionTypeRecommend,
	},
}

func (x *_FunctionTypeEnum) Parse(short string) FunctionType {
	return FunctionType(FunctionType_value["FunctionType"+strings.TrimPrefix(short, "FunctionType")])
}

func (x FunctionType) IsUnknown() bool {
	return x == FunctionTypeEnum.Unknown
}

func (x FunctionType) GetUnknown() FunctionType {
	return FunctionTypeEnum.Unknown
}

func (x FunctionType) IsCodeList() bool {
	return x == FunctionTypeEnum.CodeList
}

func (x FunctionType) GetCodeList() FunctionType {
	return FunctionTypeEnum.CodeList
}

func (x FunctionType) IsShareholder() bool {
	return x == FunctionTypeEnum.Shareholder
}

func (x FunctionType) GetShareholder() FunctionType {
	return FunctionTypeEnum.Shareholder
}

func (x FunctionType) IsLongLine() bool {
	return x == FunctionTypeEnum.LongLine
}

func (x FunctionType) GetLongLine() FunctionType {
	return FunctionTypeEnum.LongLine
}

func (x FunctionType) IsShortLine() bool {
	return x == FunctionTypeEnum.ShortLine
}

func (x FunctionType) GetShortLine() FunctionType {
	return FunctionTypeEnum.ShortLine
}

func (x FunctionType) IsRecommend() bool {
	return x == FunctionTypeEnum.Recommend
}

func (x FunctionType) GetRecommend() FunctionType {
	return FunctionTypeEnum.Recommend
}

func (x FunctionType) Valid() bool {
	if x == FunctionType_FunctionTypeUnknown {
		return false
	}
	return x.ZeroValid()
}

func (x FunctionType) ZeroValid() bool {
	_, ok := FunctionType_name[int32(x)]
	return ok
}
func (x FunctionType) Short() string {
	n := x.String()
	typ := "FunctionType"
	if len(n) > len(typ) {
		if n[:len(typ)] == typ {
			return n[len(typ):]
		}
	}
	return n

}
func (x *FunctionType) SetValue(v int32) {
	*x = FunctionType(v)
}

func (x FunctionType) GetValue() int32 {
	return int32(x)
}

func (x *FunctionType) Type() string {
	return "FunctionType"
}

type FunctionType int32

const (
	FunctionType_FunctionTypeUnknown     FunctionType = 0
	FunctionType_FunctionTypeCodeList    FunctionType = 1
	FunctionType_FunctionTypeShareholder FunctionType = 2
	FunctionType_FunctionTypeLongLine    FunctionType = 3
	FunctionType_FunctionTypeShortLine   FunctionType = 4
	FunctionType_FunctionTypeRecommend   FunctionType = 5
)

var FunctionType_name = map[int32]string{
	0: "FunctionTypeUnknown",
	1: "FunctionTypeCodeList",
	2: "FunctionTypeShareholder",
	3: "FunctionTypeLongLine",
	4: "FunctionTypeShortLine",
	5: "FunctionTypeRecommend",
}

var FunctionType_value = map[string]int32{
	"FunctionTypeUnknown":     0,
	"FunctionTypeCodeList":    1,
	"FunctionTypeShareholder": 2,
	"FunctionTypeLongLine":    3,
	"FunctionTypeShortLine":   4,
	"FunctionTypeRecommend":   5,
}

func (x FunctionType) String() string {
	return proto.EnumName(FunctionType_name, int32(x))
}

func (FunctionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{3}
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

// 权重规则
type WeightRule struct {
	//
	IsNew bool `protobuf:"varint,1,opt,name=isNew,proto3" json:"isNew"`
	//
	Price *WeightUnit `protobuf:"bytes,2,opt,name=price,proto3" json:"price"`
	//
	Focus *WeightUnit `protobuf:"bytes,3,opt,name=focus,proto3" json:"focus"`
	//
	TotalNumRatio *WeightUnit `protobuf:"bytes,4,opt,name=totalNumRatio,proto3" json:"totalNumRatio"`
	//
	AvgFreesharesRatio *WeightUnit `protobuf:"bytes,5,opt,name=avgFreesharesRatio,proto3" json:"avgFreesharesRatio"`
	//
	HoldRatioTotal float64 `protobuf:"fixed64,6,opt,name=holdRatioTotal,proto3" json:"holdRatioTotal"`
	//
	FreeholdRatioTotal float64 `protobuf:"fixed64,7,opt,name=freeholdRatioTotal,proto3" json:"freeholdRatioTotal"`
}

func (m *WeightRule) Validate() error {
	return nil
}

func (m *WeightRule) Reset()         { *m = WeightRule{} }
func (m *WeightRule) String() string { return proto.CompactTextString(m) }
func (*WeightRule) ProtoMessage()    {}
func (*WeightRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{1}
}

func (m *WeightRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightRule.Unmarshal(m, b)
}
func (m *WeightRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightRule.Marshal(b, m, deterministic)
}
func (m *WeightRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightRule.Merge(m, src)
}
func (m *WeightRule) XXX_Size() int {
	return xxx_messageInfo_WeightRule.Size(m)
}
func (m *WeightRule) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightRule.DiscardUnknown(m)
}

var xxx_messageInfo_WeightRule proto.InternalMessageInfo

func (m *WeightRule) GetIsNew() bool {
	if m != nil {
		return m.IsNew
	}
	return false
}

func (m *WeightRule) GetPrice() *WeightUnit {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *WeightRule) GetFocus() *WeightUnit {
	if m != nil {
		return m.Focus
	}
	return nil
}

func (m *WeightRule) GetTotalNumRatio() *WeightUnit {
	if m != nil {
		return m.TotalNumRatio
	}
	return nil
}

func (m *WeightRule) GetAvgFreesharesRatio() *WeightUnit {
	if m != nil {
		return m.AvgFreesharesRatio
	}
	return nil
}

func (m *WeightRule) GetHoldRatioTotal() float64 {
	if m != nil {
		return m.HoldRatioTotal
	}
	return 0
}

func (m *WeightRule) GetFreeholdRatioTotal() float64 {
	if m != nil {
		return m.FreeholdRatioTotal
	}
	return 0
}

func init() {
	proto.RegisterEnum("digger.EastMoneyType", EastMoneyType_name, EastMoneyType_value)
	proto.RegisterEnum("digger.RMType", RMType_name, RMType_value)
	proto.RegisterEnum("digger.RMState", RMState_name, RMState_value)
	proto.RegisterEnum("digger.FunctionType", FunctionType_name, FunctionType_value)
	proto.RegisterType((*WeightUnit)(nil), "digger.WeightUnit")
	proto.RegisterType((*WeightRule)(nil), "digger.WeightRule")
}

func init() { proto.RegisterFile("digger/prober.proto", fileDescriptor_f34ec7a831e4a2b5) }

var fileDescriptor_f34ec7a831e4a2b5 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xff, 0xaf, 0x13, 0x3b, 0x7f, 0x26, 0xfd, 0xd8, 0x6e, 0x52, 0x6a, 0xc4, 0x25, 0xea, 0x01,
	0x45, 0x39, 0xa4, 0x12, 0x5c, 0x90, 0xe0, 0x42, 0x11, 0x85, 0x43, 0x92, 0x56, 0x4e, 0x2a, 0x24,
	0x6e, 0x8e, 0x3d, 0x8d, 0x2d, 0x92, 0x5d, 0x6b, 0x77, 0x9d, 0xa8, 0x4f, 0xc1, 0x4b, 0x20, 0x8e,
	0x3c, 0x09, 0x0f, 0x85, 0xbc, 0xeb, 0x28, 0x76, 0x48, 0x6f, 0xfe, 0x7d, 0xcc, 0xf8, 0x37, 0xb3,
	0xab, 0x85, 0x4e, 0x9c, 0x2e, 0x16, 0x28, 0xaf, 0x32, 0x29, 0xe6, 0x28, 0x87, 0x99, 0x14, 0x5a,
	0x30, 0xcf, 0x92, 0x97, 0xbf, 0x08, 0xc0, 0x57, 0x4c, 0x17, 0x89, 0xbe, 0xe7, 0xa9, 0x66, 0x5d,
	0x70, 0xd7, 0xe1, 0x32, 0x47, 0x9f, 0xf4, 0x48, 0x9f, 0x04, 0x16, 0x14, 0x6c, 0x18, 0x45, 0xf9,
	0xca, 0x77, 0x2c, 0x6b, 0x00, 0xeb, 0x41, 0x3b, 0x12, 0x5c, 0x61, 0x94, 0xeb, 0x74, 0x8d, 0x7e,
	0xa3, 0x47, 0xfa, 0x6e, 0x50, 0xa5, 0x98, 0x0f, 0xad, 0x48, 0xe4, 0x5c, 0xa3, 0xf4, 0x9b, 0x46,
	0xdd, 0xc2, 0x42, 0xc9, 0x24, 0x2a, 0xe4, 0xda, 0x77, 0x4d, 0xcf, 0x2d, 0x64, 0xcf, 0xc1, 0x53,
	0xf9, 0x7c, 0x82, 0x1b, 0xdf, 0xeb, 0x91, 0xfe, 0xff, 0x41, 0x89, 0x2e, 0xff, 0x38, 0xdb, 0xa0,
	0x41, 0xbe, 0x34, 0x91, 0x52, 0x55, 0xb8, 0x88, 0x71, 0x59, 0xc0, 0xfa, 0xe0, 0x66, 0x32, 0x8d,
	0xd0, 0x04, 0x6d, 0xbf, 0x66, 0x43, 0x3b, 0xe5, 0x70, 0x37, 0x61, 0x60, 0x0d, 0x85, 0xf3, 0x41,
	0x44, 0xb9, 0x32, 0xb1, 0x9f, 0x70, 0x1a, 0x03, 0x7b, 0x0b, 0xc7, 0x5a, 0xe8, 0x70, 0x39, 0xc9,
	0x57, 0x41, 0xa8, 0x53, 0x61, 0x46, 0x39, 0x5c, 0x51, 0x37, 0xb2, 0x6b, 0x60, 0xe1, 0x7a, 0x71,
	0x23, 0x11, 0x55, 0x12, 0x4a, 0x54, 0xb6, 0xdc, 0x7d, 0xb2, 0xfc, 0x80, 0x9b, 0xbd, 0x82, 0x93,
	0x44, 0x2c, 0x63, 0x03, 0x66, 0x45, 0x77, 0xb3, 0x16, 0x12, 0xec, 0xb1, 0x6c, 0x08, 0xec, 0x41,
	0x22, 0xee, 0x79, 0x5b, 0xc6, 0x7b, 0x40, 0x19, 0xfc, 0x20, 0x70, 0xfc, 0x29, 0x54, 0x7a, 0x2c,
	0x38, 0x3e, 0xce, 0x1e, 0xb3, 0xe2, 0xb0, 0xba, 0x35, 0xe2, 0x9e, 0x7f, 0xe7, 0x62, 0xc3, 0xe9,
	0x7f, 0xec, 0x02, 0x3a, 0x35, 0xe5, 0x8b, 0x58, 0xc6, 0x28, 0x29, 0x61, 0xe7, 0x70, 0x56, 0x13,
	0x26, 0xb8, 0x51, 0xd4, 0x61, 0x2f, 0xe1, 0xa2, 0x46, 0xdf, 0x66, 0x28, 0x8b, 0x3f, 0x73, 0x45,
	0x1b, 0xff, 0x34, 0xfb, 0x7c, 0x37, 0x4a, 0x95, 0xa6, 0xcd, 0xc1, 0x7b, 0xf0, 0x82, 0xb1, 0x49,
	0x72, 0x06, 0xc7, 0xc1, 0xaa, 0x1e, 0xe1, 0x14, 0xda, 0x96, 0x9a, 0x26, 0x42, 0x6a, 0x4a, 0xd8,
	0x09, 0x80, 0x25, 0x46, 0x82, 0x2f, 0xa8, 0x33, 0xf8, 0x49, 0xa0, 0x15, 0x8c, 0xa7, 0x3a, 0xd4,
	0xc8, 0x18, 0x9c, 0x94, 0x9f, 0xbb, 0x06, 0x1d, 0x38, 0x2d, 0xb9, 0x3b, 0x89, 0x59, 0x28, 0x31,
	0xa6, 0xa4, 0x62, 0x9c, 0xea, 0x50, 0x6a, 0x8c, 0xa9, 0xc3, 0x28, 0x1c, 0x95, 0xdc, 0x4d, 0x2a,
	0x95, 0xa6, 0x0d, 0x13, 0xa7, 0x74, 0x61, 0x24, 0x78, 0x4c, 0x9b, 0x15, 0xd3, 0x2c, 0x49, 0x65,
	0x4c, 0x5d, 0x13, 0xd0, 0x32, 0xb7, 0x6b, 0x94, 0xd4, 0x63, 0x5d, 0xa0, 0x25, 0xf1, 0x61, 0x1e,
	0xf2, 0x58, 0x70, 0x8c, 0x69, 0x6b, 0xf0, 0x9b, 0xc0, 0xd1, 0x4d, 0xce, 0xa3, 0x62, 0x1b, 0x66,
	0xd6, 0x0b, 0xe8, 0x54, 0xf1, 0x2e, 0xb0, 0x0f, 0xdd, 0xaa, 0xf0, 0x51, 0xc4, 0x68, 0x16, 0x45,
	0x8a, 0xf5, 0x56, 0x95, 0x69, 0x71, 0x5b, 0x12, 0x7b, 0x24, 0xce, 0x7e, 0x59, 0xb1, 0x9d, 0x51,
	0xca, 0x91, 0x36, 0xd8, 0x0b, 0x38, 0xaf, 0x97, 0x09, 0xa9, 0x8d, 0xd4, 0xdc, 0x97, 0x02, 0x8c,
	0xc4, 0x6a, 0x85, 0x3c, 0xa6, 0xee, 0x75, 0xfb, 0xdb, 0xb3, 0xe1, 0xd5, 0x3b, 0x7b, 0x57, 0xe7,
	0x9e, 0x79, 0x3b, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x79, 0x59, 0xd9, 0xa4, 0x52, 0x04,
	0x00, 0x00,
}
