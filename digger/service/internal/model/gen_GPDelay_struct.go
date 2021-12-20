package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPDelay struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	Name          string        `bson:"Name" json:"Name"`
	Disabled      bool          `bson:"Disabled" json:"Disabled"`
	DecreaseTag   int32         `bson:"DecreaseTag" json:"DecreaseTag"`
	EffectiveDate string        `bson:"EffectiveDate" json:"EffectiveDate"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate    int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew         bool
}

const (
	GPDelayMgoFieldID            = "_id"
	GPDelayMgoFieldSecucode      = "Secucode"
	GPDelayMgoFieldName          = "Name"
	GPDelayMgoFieldDisabled      = "Disabled"
	GPDelayMgoFieldDecreaseTag   = "DecreaseTag"
	GPDelayMgoFieldEffectiveDate = "EffectiveDate"
	GPDelayMgoFieldCreateDate    = "CreateDate"
	GPDelayMgoFieldUpdateDate    = "UpdateDate"
)
const (
	GPDelayMgoSortFieldIDAsc  = "_id"
	GPDelayMgoSortFieldIDDesc = "-_id"
)

func (p *GPDelay) GetNameSpace() string {
	return "digger"
}

func (p *GPDelay) GetClassName() string {
	return "GPDelay"
}

type _GPDelayMgr struct {
}

var GPDelayMgr *_GPDelayMgr

// Get_GPDelayMgr returns the orm manager in case of its name starts with lower letter
func Get_GPDelayMgr() *_GPDelayMgr { return GPDelayMgr }

func (m *_GPDelayMgr) NewGPDelay() *GPDelay {
	rval := new(GPDelay)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
