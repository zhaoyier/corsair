package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type CNConfig struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DecreaseTag    int32         `bson:"DecreaseTag" json:"DecreaseTag"`
	DecreasePeriod int32         `bson:"DecreasePeriod" json:"DecreasePeriod"`
	PriceInterval  int32         `bson:"PriceInterval" json:"PriceInterval"`
	UpdateBy       string        `bson:"UpdateBy" json:"UpdateBy"`
	UpdateDate     int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew          bool
}

const (
	CNConfigMgoFieldID             = "_id"
	CNConfigMgoFieldDecreaseTag    = "DecreaseTag"
	CNConfigMgoFieldDecreasePeriod = "DecreasePeriod"
	CNConfigMgoFieldPriceInterval  = "PriceInterval"
	CNConfigMgoFieldUpdateBy       = "UpdateBy"
	CNConfigMgoFieldUpdateDate     = "UpdateDate"
)
const (
	CNConfigMgoSortFieldIDAsc  = "_id"
	CNConfigMgoSortFieldIDDesc = "-_id"
)

func (p *CNConfig) GetNameSpace() string {
	return "digger"
}

func (p *CNConfig) GetClassName() string {
	return "CNConfig"
}

type _CNConfigMgr struct {
}

var CNConfigMgr *_CNConfigMgr

// Get_CNConfigMgr returns the orm manager in case of its name starts with lower letter
func Get_CNConfigMgr() *_CNConfigMgr { return CNConfigMgr }

func (m *_CNConfigMgr) NewCNConfig() *CNConfig {
	rval := new(CNConfig)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
