package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDLongLine struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	Name          string        `bson:"Name" json:"Name"`
	ValueIndex    int32         `bson:"ValueIndex" json:"ValueIndex"`
	CumulantPrice string        `bson:"CumulantPrice" json:"CumulantPrice"`
	CumulantFocus string        `bson:"CumulantFocus" json:"CumulantFocus"`
	CumulantDate  string        `bson:"CumulantDate" json:"CumulantDate"`
	GDReduceRatio string        `bson:"GDReduceRatio" json:"GDReduceRatio"`
	Disabled      bool          `bson:"Disabled" json:"Disabled"`
	Remark        string        `bson:"Remark" json:"Remark"`
	EndDate       string        `bson:"EndDate" json:"EndDate"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	isNew         bool
}

const (
	GDLongLineMgoFieldID            = "_id"
	GDLongLineMgoFieldSecucode      = "Secucode"
	GDLongLineMgoFieldName          = "Name"
	GDLongLineMgoFieldValueIndex    = "ValueIndex"
	GDLongLineMgoFieldCumulantPrice = "CumulantPrice"
	GDLongLineMgoFieldCumulantFocus = "CumulantFocus"
	GDLongLineMgoFieldCumulantDate  = "CumulantDate"
	GDLongLineMgoFieldGDReduceRatio = "GDReduceRatio"
	GDLongLineMgoFieldDisabled      = "Disabled"
	GDLongLineMgoFieldRemark        = "Remark"
	GDLongLineMgoFieldEndDate       = "EndDate"
	GDLongLineMgoFieldCreateDate    = "CreateDate"
)
const (
	GDLongLineMgoSortFieldIDAsc  = "_id"
	GDLongLineMgoSortFieldIDDesc = "-_id"
)

func (p *GDLongLine) GetNameSpace() string {
	return "digger"
}

func (p *GDLongLine) GetClassName() string {
	return "GDLongLine"
}

type _GDLongLineMgr struct {
}

var GDLongLineMgr *_GDLongLineMgr

// Get_GDLongLineMgr returns the orm manager in case of its name starts with lower letter
func Get_GDLongLineMgr() *_GDLongLineMgr { return GDLongLineMgr }

func (m *_GDLongLineMgr) NewGDLongLine() *GDLongLine {
	rval := new(GDLongLine)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
