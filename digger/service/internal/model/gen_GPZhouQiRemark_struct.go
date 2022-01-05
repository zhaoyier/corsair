package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPZhouQiRemark struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Remark     string        `bson:"Remark" json:"Remark"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPZhouQiRemarkMgoFieldID         = "_id"
	GPZhouQiRemarkMgoFieldRemark     = "Remark"
	GPZhouQiRemarkMgoFieldUpdateDate = "UpdateDate"
)
const (
	GPZhouQiRemarkMgoSortFieldIDAsc  = "_id"
	GPZhouQiRemarkMgoSortFieldIDDesc = "-_id"
)

func (p *GPZhouQiRemark) GetNameSpace() string {
	return "digger"
}

func (p *GPZhouQiRemark) GetClassName() string {
	return "GPZhouQiRemark"
}

type _GPZhouQiRemarkMgr struct {
}

var GPZhouQiRemarkMgr *_GPZhouQiRemarkMgr

// Get_GPZhouQiRemarkMgr returns the orm manager in case of its name starts with lower letter
func Get_GPZhouQiRemarkMgr() *_GPZhouQiRemarkMgr { return GPZhouQiRemarkMgr }

func (m *_GPZhouQiRemarkMgr) NewGPZhouQiRemark() *GPZhouQiRemark {
	rval := new(GPZhouQiRemark)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
