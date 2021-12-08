package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPShortLine struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	Name         string        `bson:"Name" json:"Name"`
	MDecrease    int32         `bson:"MDecrease" json:"MDecrease"`
	TDecrease    int32         `bson:"TDecrease" json:"TDecrease"`
	DecreaseTag  int32         `bson:"DecreaseTag" json:"DecreaseTag"`
	MaxPrice     float64       `bson:"MaxPrice" json:"MaxPrice"`
	MinPrice     float64       `bson:"MinPrice" json:"MinPrice"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	MaxPDay      int64         `bson:"MaxPDay" json:"MaxPDay"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	Remark       string        `bson:"Remark" json:"Remark"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPShortLineMgoFieldID           = "_id"
	GPShortLineMgoFieldSecucode     = "Secucode"
	GPShortLineMgoFieldName         = "Name"
	GPShortLineMgoFieldMDecrease    = "MDecrease"
	GPShortLineMgoFieldTDecrease    = "TDecrease"
	GPShortLineMgoFieldDecreaseTag  = "DecreaseTag"
	GPShortLineMgoFieldMaxPrice     = "MaxPrice"
	GPShortLineMgoFieldMinPrice     = "MinPrice"
	GPShortLineMgoFieldPresentPrice = "PresentPrice"
	GPShortLineMgoFieldMaxPDay      = "MaxPDay"
	GPShortLineMgoFieldDisabled     = "Disabled"
	GPShortLineMgoFieldRemark       = "Remark"
	GPShortLineMgoFieldCreateDate   = "CreateDate"
	GPShortLineMgoFieldUpdateDate   = "UpdateDate"
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
