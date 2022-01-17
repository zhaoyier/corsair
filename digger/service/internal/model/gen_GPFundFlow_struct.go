package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPFundFlow struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name          string        `bson:"Name" json:"Name"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	FundDate      int64         `bson:"FundDate" json:"FundDate"`
	Inflow        int32         `bson:"Inflow" json:"Inflow"`
	InflowRatio   int32         `bson:"InflowRatio" json:"InflowRatio"`
	PresentPrice  int32         `bson:"PresentPrice" json:"PresentPrice"`
	IncreaseRatio int32         `bson:"IncreaseRatio" json:"IncreaseRatio"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	isNew         bool
}

const (
	GPFundFlowMgoFieldID            = "_id"
	GPFundFlowMgoFieldName          = "Name"
	GPFundFlowMgoFieldSecucode      = "Secucode"
	GPFundFlowMgoFieldFundDate      = "FundDate"
	GPFundFlowMgoFieldInflow        = "Inflow"
	GPFundFlowMgoFieldInflowRatio   = "InflowRatio"
	GPFundFlowMgoFieldPresentPrice  = "PresentPrice"
	GPFundFlowMgoFieldIncreaseRatio = "IncreaseRatio"
	GPFundFlowMgoFieldCreateDate    = "CreateDate"
)
const (
	GPFundFlowMgoSortFieldIDAsc  = "_id"
	GPFundFlowMgoSortFieldIDDesc = "-_id"
)

func (p *GPFundFlow) GetNameSpace() string {
	return "digger"
}

func (p *GPFundFlow) GetClassName() string {
	return "GPFundFlow"
}

type _GPFundFlowMgr struct {
}

var GPFundFlowMgr *_GPFundFlowMgr

// Get_GPFundFlowMgr returns the orm manager in case of its name starts with lower letter
func Get_GPFundFlowMgr() *_GPFundFlowMgr { return GPFundFlowMgr }

func (m *_GPFundFlowMgr) NewGPFundFlow() *GPFundFlow {
	rval := new(GPFundFlow)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
