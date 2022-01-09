package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDAggregation struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name               string        `bson:"Name" json:"Name"`
	Secucode           string        `bson:"Secucode" json:"Secucode"`
	EndDate            int64         `bson:"EndDate" json:"EndDate"`
	TotalRatioAccum    int32         `bson:"TotalRatioAccum" json:"TotalRatioAccum"`
	PriceRatio         int32         `bson:"PriceRatio" json:"PriceRatio"`
	PriceMax           float64       `bson:"PriceMax" json:"PriceMax"`
	PriceMin           float64       `bson:"PriceMin" json:"PriceMin"`
	HolderTotalNum     int32         `bson:"HolderTotalNum" json:"HolderTotalNum"`
	HoldFocus          string        `bson:"HoldFocus" json:"HoldFocus"`
	HoldRatioTotal     int32         `bson:"HoldRatioTotal" json:"HoldRatioTotal"`
	FreeholdRatioTotal int32         `bson:"FreeholdRatioTotal" json:"FreeholdRatioTotal"`
	UpdateDate         int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew              bool
}

const (
	GDAggregationMgoFieldID                 = "_id"
	GDAggregationMgoFieldName               = "Name"
	GDAggregationMgoFieldSecucode           = "Secucode"
	GDAggregationMgoFieldEndDate            = "EndDate"
	GDAggregationMgoFieldTotalRatioAccum    = "TotalRatioAccum"
	GDAggregationMgoFieldPriceRatio         = "PriceRatio"
	GDAggregationMgoFieldPriceMax           = "PriceMax"
	GDAggregationMgoFieldPriceMin           = "PriceMin"
	GDAggregationMgoFieldHolderTotalNum     = "HolderTotalNum"
	GDAggregationMgoFieldHoldFocus          = "HoldFocus"
	GDAggregationMgoFieldHoldRatioTotal     = "HoldRatioTotal"
	GDAggregationMgoFieldFreeholdRatioTotal = "FreeholdRatioTotal"
	GDAggregationMgoFieldUpdateDate         = "UpdateDate"
)
const (
	GDAggregationMgoSortFieldIDAsc  = "_id"
	GDAggregationMgoSortFieldIDDesc = "-_id"
)

func (p *GDAggregation) GetNameSpace() string {
	return "digger"
}

func (p *GDAggregation) GetClassName() string {
	return "GDAggregation"
}

type _GDAggregationMgr struct {
}

var GDAggregationMgr *_GDAggregationMgr

// Get_GDAggregationMgr returns the orm manager in case of its name starts with lower letter
func Get_GDAggregationMgr() *_GDAggregationMgr { return GDAggregationMgr }

func (m *_GDAggregationMgr) NewGDAggregation() *GDAggregation {
	rval := new(GDAggregation)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
