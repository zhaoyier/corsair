package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type Daily struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	Price      string        `bson:"Price" json:"Price"`
	Market     string        `bson:"Market" json:"Market"`
	Disabled   bool          `bson:"Disabled" json:"Disabled"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	isNew      bool
}

const (
	DailyMgoFieldID         = "_id"
	DailyMgoFieldSecucode   = "Secucode"
	DailyMgoFieldPrice      = "Price"
	DailyMgoFieldMarket     = "Market"
	DailyMgoFieldDisabled   = "Disabled"
	DailyMgoFieldCreateDate = "CreateDate"
)
const (
	DailyMgoSortFieldIDAsc  = "_id"
	DailyMgoSortFieldIDDesc = "-_id"
)

func (p *Daily) GetNameSpace() string {
	return "digger"
}

func (p *Daily) GetClassName() string {
	return "Daily"
}

type _DailyMgr struct {
}

var DailyMgr *_DailyMgr

// Get_DailyMgr returns the orm manager in case of its name starts with lower letter
func Get_DailyMgr() *_DailyMgr { return DailyMgr }

func (m *_DailyMgr) NewDaily() *Daily {
	rval := new(Daily)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
