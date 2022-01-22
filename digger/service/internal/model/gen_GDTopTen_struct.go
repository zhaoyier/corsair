package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDTopTen struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	EndDate       int64         `bson:"EndDate" json:"EndDate"`
	HolderRank    int32         `bson:"HolderRank" json:"HolderRank"`
	HolderName    string        `bson:"HolderName" json:"HolderName"`
	HolderType    string        `bson:"HolderType" json:"HolderType"`
	HoldNum       int32         `bson:"HoldNum" json:"HoldNum"`
	TopType       int32         `bson:"TopType" json:"TopType"`
	HoldnumRation float64       `bson:"HoldnumRation" json:"HoldnumRation"`
	HoldNumChange string        `bson:"HoldNumChange" json:"HoldNumChange"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	isNew         bool
}

const (
	GDTopTenMgoFieldID            = "_id"
	GDTopTenMgoFieldSecucode      = "Secucode"
	GDTopTenMgoFieldEndDate       = "EndDate"
	GDTopTenMgoFieldHolderRank    = "HolderRank"
	GDTopTenMgoFieldHolderName    = "HolderName"
	GDTopTenMgoFieldHolderType    = "HolderType"
	GDTopTenMgoFieldHoldNum       = "HoldNum"
	GDTopTenMgoFieldTopType       = "TopType"
	GDTopTenMgoFieldHoldnumRation = "HoldnumRation"
	GDTopTenMgoFieldHoldNumChange = "HoldNumChange"
	GDTopTenMgoFieldCreateDate    = "CreateDate"
)
const (
	GDTopTenMgoSortFieldIDAsc  = "_id"
	GDTopTenMgoSortFieldIDDesc = "-_id"
)

func (p *GDTopTen) GetNameSpace() string {
	return "digger"
}

func (p *GDTopTen) GetClassName() string {
	return "GDTopTen"
}

type _GDTopTenMgr struct {
}

var GDTopTenMgr *_GDTopTenMgr

// Get_GDTopTenMgr returns the orm manager in case of its name starts with lower letter
func Get_GDTopTenMgr() *_GDTopTenMgr { return GDTopTenMgr }

func (m *_GDTopTenMgr) NewGDTopTen() *GDTopTen {
	rval := new(GDTopTen)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
