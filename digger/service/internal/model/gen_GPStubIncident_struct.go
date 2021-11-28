package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPStubIncident struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	PriceDiff  int32         `bson:"PriceDiff" json:"PriceDiff"`
	Decrease   int32         `bson:"Decrease" json:"Decrease"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPStubIncidentMgoFieldID         = "_id"
	GPStubIncidentMgoFieldSecucode   = "Secucode"
	GPStubIncidentMgoFieldPriceDiff  = "PriceDiff"
	GPStubIncidentMgoFieldDecrease   = "Decrease"
	GPStubIncidentMgoFieldCreateDate = "CreateDate"
	GPStubIncidentMgoFieldUpdateDate = "UpdateDate"
)
const (
	GPStubIncidentMgoSortFieldIDAsc  = "_id"
	GPStubIncidentMgoSortFieldIDDesc = "-_id"
)

func (p *GPStubIncident) GetNameSpace() string {
	return "digger"
}

func (p *GPStubIncident) GetClassName() string {
	return "GPStubIncident"
}

type _GPStubIncidentMgr struct {
}

var GPStubIncidentMgr *_GPStubIncidentMgr

// Get_GPStubIncidentMgr returns the orm manager in case of its name starts with lower letter
func Get_GPStubIncidentMgr() *_GPStubIncidentMgr { return GPStubIncidentMgr }

func (m *_GPStubIncidentMgr) NewGPStubIncident() *GPStubIncident {
	rval := new(GPStubIncident)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
