package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPRecommend struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Secucode      string        `bson:"Secucode" json:"Secucode"`
	Level         float64       `bson:"Level" json:"Level"`
	CumulantPrice string        `bson:"CumulantPrice" json:"CumulantPrice"`
	CumulantFocus string        `bson:"CumulantFocus" json:"CumulantFocus"`
	CumulantDate  string        `bson:"CumulantDate" json:"CumulantDate"`
	Disabled      bool          `bson:"Disabled" json:"Disabled"`
	Remark        string        `bson:"Remark" json:"Remark"`
	EndDate       string        `bson:"EndDate" json:"EndDate"`
	CreateDate    int64         `bson:"CreateDate" json:"CreateDate"`
	isNew         bool
}

const (
	GPRecommendMgoFieldID            = "_id"
	GPRecommendMgoFieldSecucode      = "Secucode"
	GPRecommendMgoFieldLevel         = "Level"
	GPRecommendMgoFieldCumulantPrice = "CumulantPrice"
	GPRecommendMgoFieldCumulantFocus = "CumulantFocus"
	GPRecommendMgoFieldCumulantDate  = "CumulantDate"
	GPRecommendMgoFieldDisabled      = "Disabled"
	GPRecommendMgoFieldRemark        = "Remark"
	GPRecommendMgoFieldEndDate       = "EndDate"
	GPRecommendMgoFieldCreateDate    = "CreateDate"
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
