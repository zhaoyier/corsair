package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPShortLine struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	PriceDiff  int32         `bson:"PriceDiff" json:"PriceDiff"`
	Decrease   int32         `bson:"Decrease" json:"Decrease"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPShortLineMgoFieldID         = "_id"
	GPShortLineMgoFieldSecucode   = "Secucode"
	GPShortLineMgoFieldPriceDiff  = "PriceDiff"
	GPShortLineMgoFieldDecrease   = "Decrease"
	GPShortLineMgoFieldCreateDate = "CreateDate"
	GPShortLineMgoFieldUpdateDate = "UpdateDate"
)
const (
	GPShortLineMgoSortFieldIDAsc  = "_id"
	GPShortLineMgoSortFieldIDDesc = "-_id"
)

func (p *GPShortLine) GetNameSpace() string {
	return "digger"
}

func (p *GPShortLine) GetClassName() string {
	return "GPShortLine"
}

type _GPShortLineMgr struct {
}

var GPShortLineMgr *_GPShortLineMgr

// Get_GPShortLineMgr returns the orm manager in case of its name starts with lower letter
func Get_GPShortLineMgr() *_GPShortLineMgr { return GPShortLineMgr }

func (m *_GPShortLineMgr) NewGPShortLine() *GPShortLine {
	rval := new(GPShortLine)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
