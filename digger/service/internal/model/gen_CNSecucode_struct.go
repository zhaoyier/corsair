package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type CNSecucode struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	SecurityCode int32         `bson:"SecurityCode" json:"SecurityCode"`
	Focus        bool          `bson:"Focus" json:"Focus"`
	Remark       string        `bson:"Remark" json:"Remark"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	CNSecucodeMgoFieldID           = "_id"
	CNSecucodeMgoFieldSecucode     = "Secucode"
	CNSecucodeMgoFieldSecurityCode = "SecurityCode"
	CNSecucodeMgoFieldFocus        = "Focus"
	CNSecucodeMgoFieldRemark       = "Remark"
	CNSecucodeMgoFieldCreateDate   = "CreateDate"
	CNSecucodeMgoFieldUpdateDate   = "UpdateDate"
)
const (
	CNSecucodeMgoSortFieldIDAsc  = "_id"
	CNSecucodeMgoSortFieldIDDesc = "-_id"
)

func (p *CNSecucode) GetNameSpace() string {
	return "digger"
}

func (p *CNSecucode) GetClassName() string {
	return "CNSecucode"
}

type _CNSecucodeMgr struct {
}

var CNSecucodeMgr *_CNSecucodeMgr

// Get_CNSecucodeMgr returns the orm manager in case of its name starts with lower letter
func Get_CNSecucodeMgr() *_CNSecucodeMgr { return CNSecucodeMgr }

func (m *_CNSecucodeMgr) NewCNSecucode() *CNSecucode {
	rval := new(CNSecucode)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
