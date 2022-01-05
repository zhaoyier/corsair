package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPCycle struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `bson:"Name" json:"Name"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	ExpectMin    float64       `bson:"ExpectMin" json:"ExpectMin"`
	ExpectMax    float64       `bson:"ExpectMax" json:"ExpectMax"`
	ExpectStart  int64         `bson:"ExpectStart" json:"ExpectStart"`
	ExpectEnd    int64         `bson:"ExpectEnd" json:"ExpectEnd"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPCycleMgoFieldID           = "_id"
	GPCycleMgoFieldName         = "Name"
	GPCycleMgoFieldSecucode     = "Secucode"
	GPCycleMgoFieldDisabled     = "Disabled"
	GPCycleMgoFieldPresentPrice = "PresentPrice"
	GPCycleMgoFieldExpectMin    = "ExpectMin"
	GPCycleMgoFieldExpectMax    = "ExpectMax"
	GPCycleMgoFieldExpectStart  = "ExpectStart"
	GPCycleMgoFieldExpectEnd    = "ExpectEnd"
	GPCycleMgoFieldCreateDate   = "CreateDate"
	GPCycleMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GPCycleMgoSortFieldIDAsc  = "_id"
	GPCycleMgoSortFieldIDDesc = "-_id"
)

func (p *GPCycle) GetNameSpace() string {
	return "digger"
}

func (p *GPCycle) GetClassName() string {
	return "GPCycle"
}

type _GPCycleMgr struct {
}

var GPCycleMgr *_GPCycleMgr

// Get_GPCycleMgr returns the orm manager in case of its name starts with lower letter
func Get_GPCycleMgr() *_GPCycleMgr { return GPCycleMgr }

func (m *_GPCycleMgr) NewGPCycle() *GPCycle {
	rval := new(GPCycle)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
