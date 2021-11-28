package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDHoldValueIndex struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	ValueIndex    int32         `bson:"ValueIndex" json:"ValueIndex"`
	CumulantPrice string        `bson:"CumulantPrice" json:"CumulantPrice"`
	CumulantFocus string        `bson:"CumulantFocus" json:"CumulantFocus"`
	CumulantDate  string        `bson:"CumulantDate" json:"CumulantDate"`
	Disabled      bool          `bson:"Disabled" json:"Disabled"`
	Remark        string        `bson:"Remark" json:"Remark"`
	EndDate       string        `bson:"EndDate" json:"EndDate"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	isNew         bool
}

const (
	GDHoldValueIndexMgoFieldID            = "_id"
	GDHoldValueIndexMgoFieldSecucode      = "Secucode"
	GDHoldValueIndexMgoFieldValueIndex    = "ValueIndex"
	GDHoldValueIndexMgoFieldCumulantPrice = "CumulantPrice"
	GDHoldValueIndexMgoFieldCumulantFocus = "CumulantFocus"
	GDHoldValueIndexMgoFieldCumulantDate  = "CumulantDate"
	GDHoldValueIndexMgoFieldDisabled      = "Disabled"
	GDHoldValueIndexMgoFieldRemark        = "Remark"
	GDHoldValueIndexMgoFieldEndDate       = "EndDate"
	GDHoldValueIndexMgoFieldCreateDate    = "CreateDate"
)
const (
	GDHoldValueIndexMgoSortFieldIDAsc  = "_id"
	GDHoldValueIndexMgoSortFieldIDDesc = "-_id"
)

func (p *GDHoldValueIndex) GetNameSpace() string {
	return "digger"
}

func (p *GDHoldValueIndex) GetClassName() string {
	return "GDHoldValueIndex"
}

type _GDHoldValueIndexMgr struct {
}

var GDHoldValueIndexMgr *_GDHoldValueIndexMgr

// Get_GDHoldValueIndexMgr returns the orm manager in case of its name starts with lower letter
func Get_GDHoldValueIndexMgr() *_GDHoldValueIndexMgr { return GDHoldValueIndexMgr }

func (m *_GDHoldValueIndexMgr) NewGDHoldValueIndex() *GDHoldValueIndex {
	rval := new(GDHoldValueIndex)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
