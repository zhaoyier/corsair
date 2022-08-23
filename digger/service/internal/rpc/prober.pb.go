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
	Unknown      FunctionType
	CodeList     FunctionType
	Shareholder  FunctionType
	LongLine     FunctionType
	ShortLine    FunctionType
	Recommend    FunctionType
	FundFlow     FunctionType
	Focus        FunctionType
	ZhouQi       FunctionType
	Waterfall    FunctionType
	StatFundFlow FunctionType
	TotalSize    int
	List         []FunctionType
	ZeroList     []FunctionType
}

var FunctionTypeEnum = _FunctionTypeEnum{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	10,
	11,
	12,
	13,
	11,
	[]FunctionType{
		FunctionType_FunctionTypeCodeList,
		FunctionType_FunctionTypeShareholder,
		FunctionType_FunctionTypeLongLine,
		FunctionType_FunctionTypeShortLine,
		FunctionType_FunctionTypeRecommend,
		FunctionType_FunctionTypeFundFlow,
		FunctionType_FunctionTypeFocus,
		FunctionType_FunctionTypeZhouQi,
		FunctionType_FunctionTypeWaterfall,
		FunctionType_FunctionTypeStatFundFlow,
	},
	[]FunctionType{
		FunctionType_FunctionTypeUnknown,
		FunctionType_FunctionTypeCodeList,
		FunctionType_FunctionTypeShareholder,
		FunctionType_FunctionTypeLongLine,
		FunctionType_FunctionTypeShortLine,
		FunctionType_FunctionTypeRecommend,
		FunctionType_FunctionTypeFundFlow,
		FunctionType_FunctionTypeFocus,
		FunctionType_FunctionTypeZhouQi,
		FunctionType_FunctionTypeWaterfall,
		FunctionType_FunctionTypeStatFundFlow,
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

func (x FunctionType) IsFundFlow() bool {
	return x == FunctionTypeEnum.FundFlow
}

func (x FunctionType) GetFundFlow() FunctionType {
	return FunctionTypeEnum.FundFlow
}

func (x FunctionType) IsFocus() bool {
	return x == FunctionTypeEnum.Focus
}

func (x FunctionType) GetFocus() FunctionType {
	return FunctionTypeEnum.Focus
}

func (x FunctionType) IsZhouQi() bool {
	return x == FunctionTypeEnum.ZhouQi
}

func (x FunctionType) GetZhouQi() FunctionType {
	return FunctionTypeEnum.ZhouQi
}

func (x FunctionType) IsWaterfall() bool {
	return x == FunctionTypeEnum.Waterfall
}

func (x FunctionType) GetWaterfall() FunctionType {
	return FunctionTypeEnum.Waterfall
}

func (x FunctionType) IsStatFundFlow() bool {
	return x == FunctionTypeEnum.StatFundFlow
}

func (x FunctionType) GetStatFundFlow() FunctionType {
	return FunctionTypeEnum.StatFundFlow
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
	FunctionType_FunctionTypeUnknown      FunctionType = 0
	FunctionType_FunctionTypeCodeList     FunctionType = 1
	FunctionType_FunctionTypeShareholder  FunctionType = 2
	FunctionType_FunctionTypeLongLine     FunctionType = 3
	FunctionType_FunctionTypeShortLine    FunctionType = 4
	FunctionType_FunctionTypeRecommend    FunctionType = 5
	FunctionType_FunctionTypeFundFlow     FunctionType = 6
	FunctionType_FunctionTypeFocus        FunctionType = 10
	FunctionType_FunctionTypeZhouQi       FunctionType = 11
	FunctionType_FunctionTypeWaterfall    FunctionType = 12
	FunctionType_FunctionTypeStatFundFlow FunctionType = 13
)

var FunctionType_name = map[int32]string{
	0:  "FunctionTypeUnknown",
	1:  "FunctionTypeCodeList",
	2:  "FunctionTypeShareholder",
	3:  "FunctionTypeLongLine",
	4:  "FunctionTypeShortLine",
	5:  "FunctionTypeRecommend",
	6:  "FunctionTypeFundFlow",
	10: "FunctionTypeFocus",
	11: "FunctionTypeZhouQi",
	12: "FunctionTypeWaterfall",
	13: "FunctionTypeStatFundFlow",
}

var FunctionType_value = map[string]int32{
	"FunctionTypeUnknown":      0,
	"FunctionTypeCodeList":     1,
	"FunctionTypeShareholder":  2,
	"FunctionTypeLongLine":     3,
	"FunctionTypeShortLine":    4,
	"FunctionTypeRecommend":    5,
	"FunctionTypeFundFlow":     6,
	"FunctionTypeFocus":        10,
	"FunctionTypeZhouQi":       11,
	"FunctionTypeWaterfall":    12,
	"FunctionTypeStatFundFlow": 13,
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
	//
	Indexes []int32 `protobuf:"varint,7,rep,packed,name=indexes,proto3" json:"indexes"`
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

func (m *WeightUnit) GetIndexes() []int32 {
	if m != nil {
		return m.Indexes
	}
	return nil
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

type PresentPrice struct {
	//
	Price float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price"`
	//
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp"`
}

func (m *PresentPrice) Validate() error {
	return nil
}

func (m *PresentPrice) Reset()         { *m = PresentPrice{} }
func (m *PresentPrice) String() string { return proto.CompactTextString(m) }
func (*PresentPrice) ProtoMessage()    {}
func (*PresentPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ec7a831e4a2b5, []int{2}
}

func (m *PresentPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresentPrice.Unmarshal(m, b)
}
func (m *PresentPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresentPrice.Marshal(b, m, deterministic)
}
func (m *PresentPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresentPrice.Merge(m, src)
}
func (m *PresentPrice) XXX_Size() int {
	return xxx_messageInfo_PresentPrice.Size(m)
}
func (m *PresentPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_PresentPrice.DiscardUnknown(m)
}

var xxx_messageInfo_PresentPrice proto.InternalMessageInfo

func (m *PresentPrice) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PresentPrice) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
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
	proto.RegisterType((*PresentPrice)(nil), "digger.PresentPrice")
}

func init() { proto.RegisterFile("digger/prober.proto", fileDescriptor_f34ec7a831e4a2b5) }

var fileDescriptor_f34ec7a831e4a2b5 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x04, 0x03, 0x02, 0x01, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x6e, 0x13, 0x31,
	0x14, 0x66, 0xf2, 0x4b, 0x5f, 0x92, 0xd6, 0x75, 0xd2, 0x76, 0x10, 0x5d, 0x44, 0x5d, 0xa0, 0x28,
	0x8b, 0x54, 0x82, 0x0d, 0x12, 0x6c, 0x28, 0x22, 0xb0, 0xe8, 0x4f, 0x98, 0xb4, 0xaa, 0xd4, 0x9d,
	0x33, 0xf3, 0x9a, 0x8c, 0x98, 0xd8, 0x23, 0xdb, 0x93, 0xd0, 0x53, 0x70, 0x09, 0x0e, 0xc0, 0x11,
	0x38, 0x00, 0x87, 0x42, 0xb6, 0x27, 0x24, 0x13, 0xda, 0x5d, 0xbe, 0x9f, 0xf7, 0xe6, 0xbd, 0xcf,
	0x76, 0xa0, 0x1d, 0xc5, 0xd3, 0x29, 0xca, 0xd3, 0x54, 0x8a, 0x09, 0xca, 0x41, 0x2a, 0x85, 0x16,
	0xb4, 0xe6, 0xc8, 0x93, 0xdf, 0x1e, 0xc0, 0x2d, 0xc6, 0xd3, 0x99, 0xbe, 0xe1, 0xb1, 0xa6, 0x1d,
	0xa8, 0x2e, 0x58, 0x92, 0xa1, 0xef, 0x75, 0xbd, 0x9e, 0x17, 0x38, 0x60, 0x58, 0x16, 0x86, 0xd9,
	0xdc, 0x2f, 0x39, 0xd6, 0x02, 0xda, 0x85, 0x46, 0x28, 0xb8, 0xc2, 0x30, 0xd3, 0xf1, 0x02, 0xfd,
	0x72, 0xd7, 0xeb, 0x55, 0x83, 0x4d, 0x8a, 0xfa, 0x50, 0x0f, 0x45, 0xc6, 0x35, 0x4a, 0xbf, 0x62,
	0xd5, 0x15, 0x34, 0x4a, 0x2a, 0x51, 0x21, 0xd7, 0x7e, 0xd5, 0xf6, 0x5c, 0x41, 0x7a, 0x08, 0x35,
	0x95, 0x4d, 0x2e, 0x71, 0xe9, 0xd7, 0xba, 0x5e, 0xef, 0x79, 0x90, 0x23, 0x53, 0x11, 0xf3, 0x08,
	0xbf, 0xa3, 0xf2, 0xeb, 0xdd, 0xb2, 0xe9, 0x95, 0xc3, 0x93, 0x3f, 0xa5, 0xd5, 0x0a, 0x41, 0x96,
	0xd8, 0x61, 0x63, 0x65, 0xea, 0x3d, 0x5b, 0xef, 0x00, 0xed, 0x41, 0x35, 0x95, 0x71, 0x88, 0x76,
	0x85, 0xc6, 0x6b, 0x3a, 0x70, 0xfb, 0x0f, 0xd6, 0xbb, 0x07, 0xce, 0x60, 0x9c, 0xf7, 0x22, 0xcc,
	0x94, 0x5d, 0xe8, 0x09, 0xa7, 0x35, 0xd0, 0xb7, 0xd0, 0xd2, 0x42, 0xb3, 0xe4, 0x32, 0x9b, 0x07,
	0x4c, 0xc7, 0xc2, 0x2e, 0xf9, 0x78, 0x45, 0xd1, 0x48, 0xcf, 0x80, 0xb2, 0xc5, 0x74, 0x28, 0x11,
	0xd5, 0x8c, 0x49, 0x54, 0xae, 0xbc, 0xfa, 0x64, 0xf9, 0x23, 0x6e, 0xfa, 0x0a, 0x76, 0x67, 0x22,
	0x89, 0x2c, 0xb8, 0x36, 0xdd, 0x6d, 0x60, 0x5e, 0xb0, 0xc5, 0xd2, 0x01, 0xd0, 0x7b, 0x89, 0xb8,
	0xe5, 0xad, 0x5b, 0xef, 0x23, 0xca, 0xc9, 0x19, 0x34, 0x47, 0xee, 0x2c, 0x46, 0x36, 0x8f, 0xce,
	0x2a, 0xb9, 0xfc, 0x4a, 0xb8, 0x94, 0x8e, 0x61, 0x47, 0xc7, 0x73, 0x54, 0x9a, 0xcd, 0x53, 0x9b,
	0x69, 0x39, 0x58, 0x13, 0xfd, 0x1f, 0x1e, 0xb4, 0x3e, 0x31, 0xa5, 0x2f, 0x04, 0xc7, 0x87, 0xeb,
	0x87, 0xd4, 0x5c, 0x85, 0x4e, 0x81, 0xb8, 0xe1, 0xdf, 0xb8, 0x58, 0x72, 0xf2, 0x8c, 0x1e, 0x41,
	0xbb, 0xa0, 0x7c, 0x11, 0x49, 0x84, 0x92, 0x78, 0xf4, 0x00, 0xf6, 0x0b, 0xc2, 0x25, 0x2e, 0x15,
	0x29, 0xd1, 0x97, 0x70, 0x54, 0xa0, 0xaf, 0x52, 0x94, 0x66, 0x7a, 0xae, 0x48, 0xf9, 0xbf, 0x66,
	0x9f, 0x47, 0xe7, 0xb1, 0xd2, 0xa4, 0xd2, 0x7f, 0x0f, 0xb5, 0xe0, 0xc2, 0x4e, 0xb2, 0x0f, 0xad,
	0x60, 0x5e, 0x1c, 0x61, 0x0f, 0x1a, 0x8e, 0x1a, 0xcf, 0x84, 0xd4, 0xc4, 0xa3, 0xbb, 0x00, 0x8e,
	0x38, 0x17, 0x7c, 0x4a, 0x4a, 0xfd, 0x9f, 0x1e, 0xd4, 0x83, 0x8b, 0xb1, 0x66, 0x1a, 0x29, 0x85,
	0xdd, 0xfc, 0xe7, 0xba, 0x41, 0x1b, 0xf6, 0x72, 0x6e, 0x24, 0x31, 0x65, 0x12, 0x23, 0xe2, 0x6d,
	0x18, 0xc7, 0x9a, 0x49, 0x8d, 0x11, 0x29, 0x51, 0x02, 0xcd, 0x9c, 0x1b, 0xc6, 0x52, 0x69, 0x52,
	0xb6, 0xe3, 0xe4, 0x2e, 0x0c, 0x05, 0x8f, 0x48, 0x65, 0xc3, 0x74, 0x3d, 0x8b, 0x65, 0x44, 0xaa,
	0x76, 0x40, 0xc7, 0x5c, 0x2d, 0x50, 0x92, 0x1a, 0xed, 0x00, 0xc9, 0x89, 0x0f, 0x13, 0xc6, 0x23,
	0xc1, 0x31, 0x22, 0xf5, 0xfe, 0xaf, 0x12, 0x34, 0x87, 0x19, 0x0f, 0x4d, 0x1a, 0x76, 0xd7, 0x23,
	0x68, 0x6f, 0xe2, 0xf5, 0xc0, 0x3e, 0x74, 0x36, 0x85, 0x8f, 0x22, 0x42, 0x1b, 0x94, 0x67, 0xe2,
	0xdd, 0x54, 0xc6, 0xe6, 0xc6, 0xcd, 0xdc, 0x91, 0x94, 0xb6, 0xcb, 0x4c, 0x3a, 0xe7, 0x31, 0x47,
	0x52, 0xa6, 0x2f, 0xe0, 0xa0, 0x58, 0x26, 0xa4, 0xb6, 0x52, 0x65, 0x5b, 0x0a, 0x30, 0x14, 0xf3,
	0x39, 0x72, 0xb3, 0xd7, 0x56, 0xbf, 0x61, 0xc6, 0xa3, 0x61, 0x22, 0x96, 0xa4, 0x66, 0x0e, 0xbf,
	0xa0, 0x98, 0x07, 0x47, 0x80, 0x1e, 0x02, 0xdd, 0xa4, 0xef, 0x66, 0x22, 0xfb, 0x1a, 0x93, 0xc6,
	0xf6, 0x37, 0x6e, 0x99, 0x46, 0x79, 0xcf, 0x92, 0x84, 0x34, 0xe9, 0x31, 0xf8, 0x85, 0xc9, 0x34,
	0xd3, 0xff, 0xbe, 0xd3, 0x3a, 0x6b, 0xdc, 0xed, 0x0c, 0x4e, 0xdf, 0xb9, 0x17, 0x37, 0xa9, 0xd9,
	0xff, 0xc6, 0x37, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x42, 0x87, 0x46, 0x32, 0x05, 0x00,
	0x00,
}
