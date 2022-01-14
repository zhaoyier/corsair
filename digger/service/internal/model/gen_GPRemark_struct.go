package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPRemark struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `bson:"Name" json:"Name"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	Remark     string        `bson:"Remark" json:"Remark"`
	Disabled   bool          `bson:"Disabled" json:"Disabled"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	isNew      bool
}

const (
	GPRemarkMgoFieldID         = "_id"
	GPRemarkMgoFieldName       = "Name"
	GPRemarkMgoFieldSecucode   = "Secucode"
	GPRemarkMgoFieldRemark     = "Remark"
	GPRemarkMgoFieldDisabled   = "Disabled"
	GPRemarkMgoFieldCreateDate = "CreateDate"
)
const (
	GPRemarkMgoSortFieldIDAsc  = "_id"
	GPRemarkMgoSortFieldIDDesc = "-_id"
)

func (p *GPRemark) GetNameSpace() string {
	return "digger"
}

func (p *GPRemark) GetClassName() string {
	return "GPRemark"
}

type _GPRemarkMgr struct {
}

var GPRemarkMgr *_GPRemarkMgr

// Get_GPRemarkMgr returns the orm manager in case of its name starts with lower letter
func Get_GPRemarkMgr() *_GPRemarkMgr { return GPRemarkMgr }

func (m *_GPRemarkMgr) NewGPRemark() *GPRemark {
	rval := new(GPRemark)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
