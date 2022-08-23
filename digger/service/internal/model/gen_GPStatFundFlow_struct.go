package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPStatFundFlow struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `bson:"Name" json:"Name"`
	Secucode    string        `bson:"Secucode" json:"Secucode"`
	Twenty      int64         `bson:"Twenty" json:"Twenty"`
	TwentyRatio int64         `bson:"TwentyRatio" json:"TwentyRatio"`
	Ten         int64         `bson:"Ten" json:"Ten"`
	TenRatio    int64         `bson:"TenRatio" json:"TenRatio"`
	Five        int64         `bson:"Five" json:"Five"`
	FiveRatio   int64         `bson:"FiveRatio" json:"FiveRatio"`
	Three       int64         `bson:"Three" json:"Three"`
	ThreeRatio  int64         `bson:"ThreeRatio" json:"ThreeRatio"`
	Traded      int64         `bson:"Traded" json:"Traded"`
	Inflow      string        `bson:"Inflow" json:"Inflow"`
	Rising      int32         `bson:"Rising" json:"Rising"`
	MonthRise   int32         `bson:"MonthRise" json:"MonthRise"`
	TermRise    int32         `bson:"TermRise" json:"TermRise"`
	UpdateDate  int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew       bool
}

const (
	GPStatFundFlowMgoFieldID          = "_id"
	GPStatFundFlowMgoFieldName        = "Name"
	GPStatFundFlowMgoFieldSecucode    = "Secucode"
	GPStatFundFlowMgoFieldTwenty      = "Twenty"
	GPStatFundFlowMgoFieldTwentyRatio = "TwentyRatio"
	GPStatFundFlowMgoFieldTen         = "Ten"
	GPStatFundFlowMgoFieldTenRatio    = "TenRatio"
	GPStatFundFlowMgoFieldFive        = "Five"
	GPStatFundFlowMgoFieldFiveRatio   = "FiveRatio"
	GPStatFundFlowMgoFieldThree       = "Three"
	GPStatFundFlowMgoFieldThreeRatio  = "ThreeRatio"
	GPStatFundFlowMgoFieldTraded      = "Traded"
	GPStatFundFlowMgoFieldInflow      = "Inflow"
	GPStatFundFlowMgoFieldRising      = "Rising"
	GPStatFundFlowMgoFieldMonthRise   = "MonthRise"
	GPStatFundFlowMgoFieldTermRise    = "TermRise"
	GPStatFundFlowMgoFieldUpdateDate  = "UpdateDate"
)
const (
	GPStatFundFlowMgoSortFieldIDAsc  = "_id"
	GPStatFundFlowMgoSortFieldIDDesc = "-_id"
)

func (p *GPStatFundFlow) GetNameSpace() string {
	return "digger"
}

func (p *GPStatFundFlow) GetClassName() string {
	return "GPStatFundFlow"
}

type _GPStatFundFlowMgr struct {
}

var GPStatFundFlowMgr *_GPStatFundFlowMgr

// Get_GPStatFundFlowMgr returns the orm manager in case of its name starts with lower letter
func Get_GPStatFundFlowMgr() *_GPStatFundFlowMgr { return GPStatFundFlowMgr }

func (m *_GPStatFundFlowMgr) NewGPStatFundFlow() *GPStatFundFlow {
	rval := new(GPStatFundFlow)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
