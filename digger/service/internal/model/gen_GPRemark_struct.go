package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPRemark struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Content    string        `bson:"Content" json:"Content"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPRemarkMgoFieldID         = "_id"
	GPRemarkMgoFieldContent    = "Content"
	GPRemarkMgoFieldUpdateDate = "UpdateDate"
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
