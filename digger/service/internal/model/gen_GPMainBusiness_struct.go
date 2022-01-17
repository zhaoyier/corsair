package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPMainBusiness struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `bson:"Name" json:"Name"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	Content    string        `bson:"Content" json:"Content"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	isNew      bool
}

const (
	GPMainBusinessMgoFieldID         = "_id"
	GPMainBusinessMgoFieldName       = "Name"
	GPMainBusinessMgoFieldSecucode   = "Secucode"
	GPMainBusinessMgoFieldContent    = "Content"
	GPMainBusinessMgoFieldCreateDate = "CreateDate"
)
const (
	GPMainBusinessMgoSortFieldIDAsc  = "_id"
	GPMainBusinessMgoSortFieldIDDesc = "-_id"
)

func (p *GPMainBusiness) GetNameSpace() string {
	return "digger"
}

func (p *GPMainBusiness) GetClassName() string {
	return "GPMainBusiness"
}

type _GPMainBusinessMgr struct {
}

var GPMainBusinessMgr *_GPMainBusinessMgr

// Get_GPMainBusinessMgr returns the orm manager in case of its name starts with lower letter
func Get_GPMainBusinessMgr() *_GPMainBusinessMgr { return GPMainBusinessMgr }

func (m *_GPMainBusinessMgr) NewGPMainBusiness() *GPMainBusiness {
	rval := new(GPMainBusiness)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
