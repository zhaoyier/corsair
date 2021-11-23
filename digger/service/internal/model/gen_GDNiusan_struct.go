package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDNiusan struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	SecurityCode string        `bson:"SecurityCode" json:"SecurityCode"`
	Niusan       string        `bson:"Niusan" json:"Niusan"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GDNiusanMgoFieldID           = "_id"
	GDNiusanMgoFieldSecurityCode = "SecurityCode"
	GDNiusanMgoFieldNiusan       = "Niusan"
	GDNiusanMgoFieldDisabled     = "Disabled"
	GDNiusanMgoFieldCreateDate   = "CreateDate"
	GDNiusanMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GDNiusanMgoSortFieldIDAsc  = "_id"
	GDNiusanMgoSortFieldIDDesc = "-_id"
)

func (p *GDNiusan) GetNameSpace() string {
	return "digger"
}

func (p *GDNiusan) GetClassName() string {
	return "GDNiusan"
}

type _GDNiusanMgr struct {
}

var GDNiusanMgr *_GDNiusanMgr

// Get_GDNiusanMgr returns the orm manager in case of its name starts with lower letter
func Get_GDNiusanMgr() *_GDNiusanMgr { return GDNiusanMgr }

func (m *_GDNiusanMgr) NewGDNiusan() *GDNiusan {
	rval := new(GDNiusan)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
