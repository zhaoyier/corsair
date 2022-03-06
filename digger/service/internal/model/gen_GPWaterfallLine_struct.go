package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPWaterfallLine struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `bson:"Name" json:"Name"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	MaxPDay      int64         `bson:"MaxPDay" json:"MaxPDay"`
	MaxPrice     float64       `bson:"MaxPrice" json:"MaxPrice"`
	MinPrice     float64       `bson:"MinPrice" json:"MinPrice"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	State        int32         `bson:"State" json:"State"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	Decrease     int32         `bson:"Decrease" json:"Decrease"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPWaterfallLineMgoFieldID           = "_id"
	GPWaterfallLineMgoFieldName         = "Name"
	GPWaterfallLineMgoFieldSecucode     = "Secucode"
	GPWaterfallLineMgoFieldMaxPDay      = "MaxPDay"
	GPWaterfallLineMgoFieldMaxPrice     = "MaxPrice"
	GPWaterfallLineMgoFieldMinPrice     = "MinPrice"
	GPWaterfallLineMgoFieldPresentPrice = "PresentPrice"
	GPWaterfallLineMgoFieldState        = "State"
	GPWaterfallLineMgoFieldDisabled     = "Disabled"
	GPWaterfallLineMgoFieldDecrease     = "Decrease"
	GPWaterfallLineMgoFieldCreateDate   = "CreateDate"
	GPWaterfallLineMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GPWaterfallLineMgoSortFieldIDAsc  = "_id"
	GPWaterfallLineMgoSortFieldIDDesc = "-_id"
)

func (p *GPWaterfallLine) GetNameSpace() string {
	return "digger"
}

func (p *GPWaterfallLine) GetClassName() string {
	return "GPWaterfallLine"
}

type _GPWaterfallLineMgr struct {
}

var GPWaterfallLineMgr *_GPWaterfallLineMgr

// Get_GPWaterfallLineMgr returns the orm manager in case of its name starts with lower letter
func Get_GPWaterfallLineMgr() *_GPWaterfallLineMgr { return GPWaterfallLineMgr }

func (m *_GPWaterfallLineMgr) NewGPWaterfallLine() *GPWaterfallLine {
	rval := new(GPWaterfallLine)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
