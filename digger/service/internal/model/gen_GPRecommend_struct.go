package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPRecommend struct {
	ID         bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Secucode   string            `bson:"Secucode" json:"Secucode"`
	RmIndex    int32             `bson:"RmIndex" json:"RmIndex"`
	RmType     int32             `bson:"RmType" json:"RmType"`
	MonthDrop  int32             `bson:"MonthDrop" json:"MonthDrop"`
	RmDate     int64             `bson:"RmDate" json:"RmDate"`
	RmPhase    int32             `bson:"RmPhase" json:"RmPhase"`
	State      int32             `bson:"State" json:"State"`
	Remark     map[string]string `bson:"Remark" json:"Remark"`
	Disabled   bool              `bson:"Disabled" json:"Disabled"`
	CreateDate int64             `bson:"CreateDate" json:"CreateDate"`
	UpdateBy   string            `bson:"UpdateBy" json:"UpdateBy"`
	UpdateDate int64             `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	GPRecommendMgoFieldID         = "_id"
	GPRecommendMgoFieldSecucode   = "Secucode"
	GPRecommendMgoFieldRmIndex    = "RmIndex"
	GPRecommendMgoFieldRmType     = "RmType"
	GPRecommendMgoFieldMonthDrop  = "MonthDrop"
	GPRecommendMgoFieldRmDate     = "RmDate"
	GPRecommendMgoFieldRmPhase    = "RmPhase"
	GPRecommendMgoFieldState      = "State"
	GPRecommendMgoFieldRemark     = "Remark"
	GPRecommendMgoFieldDisabled   = "Disabled"
	GPRecommendMgoFieldCreateDate = "CreateDate"
	GPRecommendMgoFieldUpdateBy   = "UpdateBy"
	GPRecommendMgoFieldUpdateDate = "UpdateDate"
)
const (
	GPRecommendMgoSortFieldIDAsc  = "_id"
	GPRecommendMgoSortFieldIDDesc = "-_id"
)

func (p *GPRecommend) GetNameSpace() string {
	return "digger"
}

func (p *GPRecommend) GetClassName() string {
	return "GPRecommend"
}

type _GPRecommendMgr struct {
}

var GPRecommendMgr *_GPRecommendMgr

// Get_GPRecommendMgr returns the orm manager in case of its name starts with lower letter
func Get_GPRecommendMgr() *_GPRecommendMgr { return GPRecommendMgr }

func (m *_GPRecommendMgr) NewGPRecommend() *GPRecommend {
	rval := new(GPRecommend)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
