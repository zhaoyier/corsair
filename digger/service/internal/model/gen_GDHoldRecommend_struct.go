package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GDHoldRecommend struct {
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
	GDHoldRecommendMgoFieldID            = "_id"
	GDHoldRecommendMgoFieldSecucode      = "Secucode"
	GDHoldRecommendMgoFieldLevel         = "Level"
	GDHoldRecommendMgoFieldCumulantPrice = "CumulantPrice"
	GDHoldRecommendMgoFieldCumulantFocus = "CumulantFocus"
	GDHoldRecommendMgoFieldCumulantDate  = "CumulantDate"
	GDHoldRecommendMgoFieldDisabled      = "Disabled"
	GDHoldRecommendMgoFieldRemark        = "Remark"
	GDHoldRecommendMgoFieldEndDate       = "EndDate"
	GDHoldRecommendMgoFieldCreateDate    = "CreateDate"
)
const (
	GDHoldRecommendMgoSortFieldIDAsc  = "_id"
	GDHoldRecommendMgoSortFieldIDDesc = "-_id"
)

func (p *GDHoldRecommend) GetNameSpace() string {
	return "digger"
}

func (p *GDHoldRecommend) GetClassName() string {
	return "GDHoldRecommend"
}

type _GDHoldRecommendMgr struct {
}

var GDHoldRecommendMgr *_GDHoldRecommendMgr

// Get_GDHoldRecommendMgr returns the orm manager in case of its name starts with lower letter
func Get_GDHoldRecommendMgr() *_GDHoldRecommendMgr { return GDHoldRecommendMgr }

func (m *_GDHoldRecommendMgr) NewGDHoldRecommend() *GDHoldRecommend {
	rval := new(GDHoldRecommend)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
