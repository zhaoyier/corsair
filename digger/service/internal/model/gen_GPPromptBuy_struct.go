package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPPromptBuy struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	Name         string        `bson:"Name" json:"Name"`
	MinPrice     float64       `bson:"MinPrice" json:"MinPrice"`
	Start        int64         `bson:"Start" json:"Start"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPPromptBuyMgoFieldID           = "_id"
	GPPromptBuyMgoFieldSecucode     = "Secucode"
	GPPromptBuyMgoFieldName         = "Name"
	GPPromptBuyMgoFieldMinPrice     = "MinPrice"
	GPPromptBuyMgoFieldStart        = "Start"
	GPPromptBuyMgoFieldPresentPrice = "PresentPrice"
	GPPromptBuyMgoFieldDisabled     = "Disabled"
	GPPromptBuyMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GPPromptBuyMgoSortFieldIDAsc  = "_id"
	GPPromptBuyMgoSortFieldIDDesc = "-_id"
)

func (p *GPPromptBuy) GetNameSpace() string {
	return "digger"
}

func (p *GPPromptBuy) GetClassName() string {
	return "GPPromptBuy"
}

type _GPPromptBuyMgr struct {
}

var GPPromptBuyMgr *_GPPromptBuyMgr

// Get_GPPromptBuyMgr returns the orm manager in case of its name starts with lower letter
func Get_GPPromptBuyMgr() *_GPPromptBuyMgr { return GPPromptBuyMgr }

func (m *_GPPromptBuyMgr) NewGPPromptBuy() *GPPromptBuy {
	rval := new(GPPromptBuy)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
