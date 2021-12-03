package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPRecommend struct {
	ID           bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Secucode     string            `bson:"Secucode" json:"Secucode"`
	RMIndex      int32             `bson:"RMIndex" json:"RMIndex"`
	RMType       int32             `bson:"RMType" json:"RMType"`
	Decrease     int32             `bson:"Decrease" json:"Decrease"`
	DecreaseTag  int32             `bson:"DecreaseTag" json:"DecreaseTag"`
	DecreaseDay  string            `bson:"DecreaseDay" json:"DecreaseDay"`
	GDDecrease   int32             `bson:"GDDecrease" json:"GDDecrease"`
	MaxPrice     float64           `bson:"MaxPrice" json:"MaxPrice"`
	MaxDay       int32             `bson:"MaxDay" json:"MaxDay"`
	PresentPrice float64           `bson:"PresentPrice" json:"PresentPrice"`
	RMPrice      string            `bson:"RMPrice" json:"RMPrice"`
	State        int32             `bson:"State" json:"State"`
	Remark       map[string]string `bson:"Remark" json:"Remark"`
	Disabled     bool              `bson:"Disabled" json:"Disabled"`
	CreateDate   int64             `bson:"CreateDate" json:"CreateDate"`
	UpdateBy     string            `bson:"UpdateBy" json:"UpdateBy"`
	UpdateDate   int64             `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPRecommendMgoFieldID           = "_id"
	GPRecommendMgoFieldSecucode     = "Secucode"
	GPRecommendMgoFieldRMIndex      = "RMIndex"
	GPRecommendMgoFieldRMType       = "RMType"
	GPRecommendMgoFieldDecrease     = "Decrease"
	GPRecommendMgoFieldDecreaseTag  = "DecreaseTag"
	GPRecommendMgoFieldDecreaseDay  = "DecreaseDay"
	GPRecommendMgoFieldGDDecrease   = "GDDecrease"
	GPRecommendMgoFieldMaxPrice     = "MaxPrice"
	GPRecommendMgoFieldMaxDay       = "MaxDay"
	GPRecommendMgoFieldPresentPrice = "PresentPrice"
	GPRecommendMgoFieldRMPrice      = "RMPrice"
	GPRecommendMgoFieldState        = "State"
	GPRecommendMgoFieldRemark       = "Remark"
	GPRecommendMgoFieldDisabled     = "Disabled"
	GPRecommendMgoFieldCreateDate   = "CreateDate"
	GPRecommendMgoFieldUpdateBy     = "UpdateBy"
	GPRecommendMgoFieldUpdateDate   = "UpdateDate"
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
