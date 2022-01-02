package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPManualDecrease struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode    string        `bson:"Secucode" json:"Secucode"`
	Name        string        `bson:"Name" json:"Name"`
	Disabled    bool          `bson:"Disabled" json:"Disabled"`
	DecreaseTag int32         `bson:"DecreaseTag" json:"DecreaseTag"`
	EffectStart int64         `bson:"EffectStart" json:"EffectStart"`
	EffectEnd   int64         `bson:"EffectEnd" json:"EffectEnd"`
	CreateDate  int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate  int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew       bool
}

const (
	GPManualDecreaseMgoFieldID          = "_id"
	GPManualDecreaseMgoFieldSecucode    = "Secucode"
	GPManualDecreaseMgoFieldName        = "Name"
	GPManualDecreaseMgoFieldDisabled    = "Disabled"
	GPManualDecreaseMgoFieldDecreaseTag = "DecreaseTag"
	GPManualDecreaseMgoFieldEffectStart = "EffectStart"
	GPManualDecreaseMgoFieldEffectEnd   = "EffectEnd"
	GPManualDecreaseMgoFieldCreateDate  = "CreateDate"
	GPManualDecreaseMgoFieldUpdateDate  = "UpdateDate"
)
const (
	GPManualDecreaseMgoSortFieldIDAsc  = "_id"
	GPManualDecreaseMgoSortFieldIDDesc = "-_id"
)

func (p *GPManualDecrease) GetNameSpace() string {
	return "digger"
}

func (p *GPManualDecrease) GetClassName() string {
	return "GPManualDecrease"
}

type _GPManualDecreaseMgr struct {
}

var GPManualDecreaseMgr *_GPManualDecreaseMgr

// Get_GPManualDecreaseMgr returns the orm manager in case of its name starts with lower letter
func Get_GPManualDecreaseMgr() *_GPManualDecreaseMgr { return GPManualDecreaseMgr }

func (m *_GPManualDecreaseMgr) NewGPManualDecrease() *GPManualDecrease {
	rval := new(GPManualDecrease)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
