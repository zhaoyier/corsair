package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDRenshu struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode           string        `bson:"Secucode" json:"Secucode"`
	SecurityCode       string        `bson:"SecurityCode" json:"SecurityCode"`
	EndDate            int64         `bson:"EndDate" json:"EndDate"`
	HolderTotalNum     int32         `bson:"HolderTotalNum" json:"HolderTotalNum"`
	TotalNumRatio      int32         `bson:"TotalNumRatio" json:"TotalNumRatio"`
	AvgFreeShares      int32         `bson:"AvgFreeShares" json:"AvgFreeShares"`
	AvgFreesharesRatio int32         `bson:"AvgFreesharesRatio" json:"AvgFreesharesRatio"`
	HoldFocus          string        `bson:"HoldFocus" json:"HoldFocus"`
	Price              int32         `bson:"Price" json:"Price"`
	AvgHoldAmt         int32         `bson:"AvgHoldAmt" json:"AvgHoldAmt"`
	HoldRatioTotal     int32         `bson:"HoldRatioTotal" json:"HoldRatioTotal"`
	FreeholdRatioTotal int32         `bson:"FreeholdRatioTotal" json:"FreeholdRatioTotal"`
	CreateDate         int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate         int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew              bool
}

const (
	GDRenshuMgoFieldID                 = "_id"
	GDRenshuMgoFieldSecucode           = "Secucode"
	GDRenshuMgoFieldSecurityCode       = "SecurityCode"
	GDRenshuMgoFieldEndDate            = "EndDate"
	GDRenshuMgoFieldHolderTotalNum     = "HolderTotalNum"
	GDRenshuMgoFieldTotalNumRatio      = "TotalNumRatio"
	GDRenshuMgoFieldAvgFreeShares      = "AvgFreeShares"
	GDRenshuMgoFieldAvgFreesharesRatio = "AvgFreesharesRatio"
	GDRenshuMgoFieldHoldFocus          = "HoldFocus"
	GDRenshuMgoFieldPrice              = "Price"
	GDRenshuMgoFieldAvgHoldAmt         = "AvgHoldAmt"
	GDRenshuMgoFieldHoldRatioTotal     = "HoldRatioTotal"
	GDRenshuMgoFieldFreeholdRatioTotal = "FreeholdRatioTotal"
	GDRenshuMgoFieldCreateDate         = "CreateDate"
	GDRenshuMgoFieldUpdateDate         = "UpdateDate"
)
const (
	GDRenshuMgoSortFieldIDAsc  = "_id"
	GDRenshuMgoSortFieldIDDesc = "-_id"
)

func (p *GDRenshu) GetNameSpace() string {
	return "digger"
}

func (p *GDRenshu) GetClassName() string {
	return "GDRenshu"
}

type _GDRenshuMgr struct {
}

var GDRenshuMgr *_GDRenshuMgr

// Get_GDRenshuMgr returns the orm manager in case of its name starts with lower letter
func Get_GDRenshuMgr() *_GDRenshuMgr { return GDRenshuMgr }

func (m *_GDRenshuMgr) NewGDRenshu() *GDRenshu {
	rval := new(GDRenshu)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
