package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPFocus struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `bson:"Name" json:"Name"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	FocusPrice   float64       `bson:"FocusPrice" json:"FocusPrice"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	ExpectPrice  float64       `bson:"ExpectPrice" json:"ExpectPrice"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPFocusMgoFieldID           = "_id"
	GPFocusMgoFieldName         = "Name"
	GPFocusMgoFieldSecucode     = "Secucode"
	GPFocusMgoFieldDisabled     = "Disabled"
	GPFocusMgoFieldFocusPrice   = "FocusPrice"
	GPFocusMgoFieldPresentPrice = "PresentPrice"
	GPFocusMgoFieldExpectPrice  = "ExpectPrice"
	GPFocusMgoFieldCreateDate   = "CreateDate"
	GPFocusMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GPFocusMgoSortFieldIDAsc  = "_id"
	GPFocusMgoSortFieldIDDesc = "-_id"
)

func (p *GPFocus) GetNameSpace() string {
	return "digger"
}

func (p *GPFocus) GetClassName() string {
	return "GPFocus"
}

type _GPFocusMgr struct {
}

var GPFocusMgr *_GPFocusMgr

// Get_GPFocusMgr returns the orm manager in case of its name starts with lower letter
func Get_GPFocusMgr() *_GPFocusMgr { return GPFocusMgr }

func (m *_GPFocusMgr) NewGPFocus() *GPFocus {
	rval := new(GPFocus)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
