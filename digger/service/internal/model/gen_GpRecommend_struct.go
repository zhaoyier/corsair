package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GpRecommend struct {
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
	GpRecommendMgoFieldID            = "_id"
	GpRecommendMgoFieldSecucode      = "Secucode"
	GpRecommendMgoFieldLevel         = "Level"
	GpRecommendMgoFieldCumulantPrice = "CumulantPrice"
	GpRecommendMgoFieldCumulantFocus = "CumulantFocus"
	GpRecommendMgoFieldCumulantDate  = "CumulantDate"
	GpRecommendMgoFieldDisabled      = "Disabled"
	GpRecommendMgoFieldRemark        = "Remark"
	GpRecommendMgoFieldEndDate       = "EndDate"
	GpRecommendMgoFieldCreateDate    = "CreateDate"
)
const (
	GpRecommendMgoSortFieldIDAsc  = "_id"
	GpRecommendMgoSortFieldIDDesc = "-_id"
)

func (p *GpRecommend) GetNameSpace() string {
	return "digger"
}

func (p *GpRecommend) GetClassName() string {
	return "GpRecommend"
}

type _GpRecommendMgr struct {
}

var GpRecommendMgr *_GpRecommendMgr

// Get_GpRecommendMgr returns the orm manager in case of its name starts with lower letter
func Get_GpRecommendMgr() *_GpRecommendMgr { return GpRecommendMgr }

func (m *_GpRecommendMgr) NewGpRecommend() *GpRecommend {
	rval := new(GpRecommend)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
