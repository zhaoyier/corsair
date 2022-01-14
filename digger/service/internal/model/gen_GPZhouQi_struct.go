package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type GPZhouQi struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `bson:"Name" json:"Name"`
	Secucode     string        `bson:"Secucode" json:"Secucode"`
	State        int32         `bson:"State" json:"State"`
	Disabled     bool          `bson:"Disabled" json:"Disabled"`
	PresentPrice float64       `bson:"PresentPrice" json:"PresentPrice"`
	ExpectMin    float64       `bson:"ExpectMin" json:"ExpectMin"`
	ExpectMax    float64       `bson:"ExpectMax" json:"ExpectMax"`
	ExpectStart  int64         `bson:"ExpectStart" json:"ExpectStart"`
	ExpectEnd    int64         `bson:"ExpectEnd" json:"ExpectEnd"`
	BuyingCount  int32         `bson:"BuyingCount" json:"BuyingCount"`
	MainBusiness string        `bson:"MainBusiness" json:"MainBusiness"`
	CreateDate   int64         `bson:"CreateDate" json:"CreateDate"`
	UpdateDate   int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew        bool
}

const (
	GPZhouQiMgoFieldID           = "_id"
	GPZhouQiMgoFieldName         = "Name"
	GPZhouQiMgoFieldSecucode     = "Secucode"
	GPZhouQiMgoFieldState        = "State"
	GPZhouQiMgoFieldDisabled     = "Disabled"
	GPZhouQiMgoFieldPresentPrice = "PresentPrice"
	GPZhouQiMgoFieldExpectMin    = "ExpectMin"
	GPZhouQiMgoFieldExpectMax    = "ExpectMax"
	GPZhouQiMgoFieldExpectStart  = "ExpectStart"
	GPZhouQiMgoFieldExpectEnd    = "ExpectEnd"
	GPZhouQiMgoFieldBuyingCount  = "BuyingCount"
	GPZhouQiMgoFieldMainBusiness = "MainBusiness"
	GPZhouQiMgoFieldCreateDate   = "CreateDate"
	GPZhouQiMgoFieldUpdateDate   = "UpdateDate"
)
const (
	GPZhouQiMgoSortFieldIDAsc  = "_id"
	GPZhouQiMgoSortFieldIDDesc = "-_id"
)

func (p *GPZhouQi) GetNameSpace() string {
	return "digger"
}

func (p *GPZhouQi) GetClassName() string {
	return "GPZhouQi"
}

type _GPZhouQiMgr struct {
}

var GPZhouQiMgr *_GPZhouQiMgr

// Get_GPZhouQiMgr returns the orm manager in case of its name starts with lower letter
func Get_GPZhouQiMgr() *_GPZhouQiMgr { return GPZhouQiMgr }

func (m *_GPZhouQiMgr) NewGPZhouQi() *GPZhouQi {
	rval := new(GPZhouQi)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
