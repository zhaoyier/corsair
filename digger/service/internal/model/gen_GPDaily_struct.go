package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPDaily struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode   string        `bson:"Secucode" json:"Secucode"`
	Name       string        `bson:"Name" json:"Name"`
	Closing    float64       `bson:"Closing" json:"Closing"`
	Rise       float64       `bson:"Rise" json:"Rise"`
	PRise      float64       `bson:"PRise" json:"PRise"`
	Turnover   int64         `bson:"Turnover" json:"Turnover"`
	Business   float64       `bson:"Business" json:"Business"`
	Liangbi    float64       `bson:"Liangbi" json:"Liangbi"`
	MaxPrice   float64       `bson:"MaxPrice" json:"MaxPrice"`
	Opening    float64       `bson:"Opening" json:"Opening"`
	Market     int64         `bson:"Market" json:"Market"`
	Traded     int64         `bson:"Traded" json:"Traded"`
	BookRatio  float64       `bson:"BookRatio" json:"BookRatio"`
	CreateDate int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPDailyMgoFieldID         = "_id"
	GPDailyMgoFieldSecucode   = "Secucode"
	GPDailyMgoFieldName       = "Name"
	GPDailyMgoFieldClosing    = "Closing"
	GPDailyMgoFieldRise       = "Rise"
	GPDailyMgoFieldPRise      = "PRise"
	GPDailyMgoFieldTurnover   = "Turnover"
	GPDailyMgoFieldBusiness   = "Business"
	GPDailyMgoFieldLiangbi    = "Liangbi"
	GPDailyMgoFieldMaxPrice   = "MaxPrice"
	GPDailyMgoFieldOpening    = "Opening"
	GPDailyMgoFieldMarket     = "Market"
	GPDailyMgoFieldTraded     = "Traded"
	GPDailyMgoFieldBookRatio  = "BookRatio"
	GPDailyMgoFieldCreateDate = "CreateDate"
	GPDailyMgoFieldUpdateDate = "UpdateDate"
)
const (
	GPDailyMgoSortFieldIDAsc  = "_id"
	GPDailyMgoSortFieldIDDesc = "-_id"
)

func (p *GPDaily) GetNameSpace() string {
	return "digger"
}

func (p *GPDaily) GetClassName() string {
	return "GPDaily"
}

type _GPDailyMgr struct {
}

var GPDailyMgr *_GPDailyMgr

// Get_GPDailyMgr returns the orm manager in case of its name starts with lower letter
func Get_GPDailyMgr() *_GPDailyMgr { return GPDailyMgr }

func (m *_GPDailyMgr) NewGPDaily() *GPDaily {
	rval := new(GPDaily)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
