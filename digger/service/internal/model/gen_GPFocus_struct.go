package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPFocus struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `bson:"Name" json:"Name"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	State        int32         `bson:"State" json:"State"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	ExpectPrice  float64       `bson:"ExpectPrice" json:"ExpectPrice"`
	ExpectDate   int64         `bson:"ExpectDate" json:"ExpectDate"`
	Remarks      []GPRemark    `bson:"Remarks" json:"Remarks"`
	MainBusiness string        `bson:"MainBusiness" json:"MainBusiness"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPFocusMgoFieldID           = "_id"
	GPFocusMgoFieldName         = "Name"
	GPFocusMgoFieldSecucode     = "Secucode"
	GPFocusMgoFieldState        = "State"
	GPFocusMgoFieldDisabled     = "Disabled"
	GPFocusMgoFieldPresentPrice = "PresentPrice"
	GPFocusMgoFieldExpectPrice  = "ExpectPrice"
	GPFocusMgoFieldExpectDate   = "ExpectDate"
	GPFocusMgoFieldRemarks      = "Remarks"
	GPFocusMgoFieldMainBusiness = "MainBusiness"
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
