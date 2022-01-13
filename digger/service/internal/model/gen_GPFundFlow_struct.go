package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPFundFlow struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `bson:"Name" json:"Name"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	Five       int32         `bson:"Five" json:"Five"`
	Ten        int32         `bson:"Ten" json:"Ten"`
	Twenty     int32         `bson:"Twenty" json:"Twenty"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPFundFlowMgoFieldID         = "_id"
	GPFundFlowMgoFieldName       = "Name"
	GPFundFlowMgoFieldSecucode   = "Secucode"
	GPFundFlowMgoFieldFive       = "Five"
	GPFundFlowMgoFieldTen        = "Ten"
	GPFundFlowMgoFieldTwenty     = "Twenty"
	GPFundFlowMgoFieldUpdateDate = "UpdateDate"
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
