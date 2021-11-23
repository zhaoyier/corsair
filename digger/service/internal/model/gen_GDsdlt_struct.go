package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDsdlt struct {
	ID                bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode          string        `bson:"Secucode" json:"Secucode"`
	EndDate           int64         `bson:"EndDate" json:"EndDate"`
	HolderRank        int32         `bson:"HolderRank" json:"HolderRank"`
	HolderName        string        `bson:"HolderName" json:"HolderName"`
	HolderType        string        `bson:"HolderType" json:"HolderType"`
	HoldNum           int32         `bson:"HoldNum" json:"HoldNum"`
	FreeHoldnumRation float64       `bson:"FreeHoldnumRation" json:"FreeHoldnumRation"`
	HoldNumChange     string        `bson:"HoldNumChange" json:"HoldNumChange"`
	CreateDate        int64         `bson:"CreateDate" json:"CreateDate"`
	isNew             bool
}

const (
	GDsdltMgoFieldID                = "_id"
	GDsdltMgoFieldSecucode          = "Secucode"
	GDsdltMgoFieldEndDate           = "EndDate"
	GDsdltMgoFieldHolderRank        = "HolderRank"
	GDsdltMgoFieldHolderName        = "HolderName"
	GDsdltMgoFieldHolderType        = "HolderType"
	GDsdltMgoFieldHoldNum           = "HoldNum"
	GDsdltMgoFieldFreeHoldnumRation = "FreeHoldnumRation"
	GDsdltMgoFieldHoldNumChange     = "HoldNumChange"
	GDsdltMgoFieldCreateDate        = "CreateDate"
)
const (
	GDsdltMgoSortFieldIDAsc  = "_id"
	GDsdltMgoSortFieldIDDesc = "-_id"
)

func (p *GDsdlt) GetNameSpace() string {
	return "digger"
}

func (p *GDsdlt) GetClassName() string {
	return "GDsdlt"
}

type _GDsdltMgr struct {
}

var GDsdltMgr *_GDsdltMgr

// Get_GDsdltMgr returns the orm manager in case of its name starts with lower letter
func Get_GDsdltMgr() *_GDsdltMgr { return GDsdltMgr }

func (m *_GDsdltMgr) NewGDsdlt() *GDsdlt {
	rval := new(GDsdlt)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
