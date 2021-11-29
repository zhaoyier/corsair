package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type SinaDaily struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	EndDate    string        `bson:"EndDate" json:"EndDate"`
	Price      float64       `bson:"Price" json:"Price"`
	Highest    float64       `bson:"Highest" json:"Highest"`
	Minimum    float64       `bson:"Minimum" json:"Minimum"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	isNew      bool
}

const (
	SinaDailyMgoFieldID         = "_id"
	SinaDailyMgoFieldSecucode   = "Secucode"
	SinaDailyMgoFieldEndDate    = "EndDate"
	SinaDailyMgoFieldPrice      = "Price"
	SinaDailyMgoFieldHighest    = "Highest"
	SinaDailyMgoFieldMinimum    = "Minimum"
	SinaDailyMgoFieldCreateDate = "CreateDate"
)
const (
	SinaDailyMgoSortFieldIDAsc  = "_id"
	SinaDailyMgoSortFieldIDDesc = "-_id"
)

func (p *SinaDaily) GetNameSpace() string {
	return "digger"
}

func (p *SinaDaily) GetClassName() string {
	return "SinaDaily"
}

type _SinaDailyMgr struct {
}

var SinaDailyMgr *_SinaDailyMgr

// Get_SinaDailyMgr returns the orm manager in case of its name starts with lower letter
func Get_SinaDailyMgr() *_SinaDailyMgr { return SinaDailyMgr }

func (m *_SinaDailyMgr) NewSinaDaily() *SinaDaily {
	rval := new(SinaDaily)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
