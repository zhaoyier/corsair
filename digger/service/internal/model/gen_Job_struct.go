package model

import "gopkg.in/mgo.v2/bson"

import "time"

var _ time.Time

type Job struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Msg        []string      `bson:"Msg" json:"Msg"`
	CreateDate string        `bson:"CreateDate" json:"CreateDate"`
	UpdateDate int64         `bson:"UpdateDate" json:"UpdateDate"`
	isNew      bool
}

const (
	JobMgoFieldID         = "_id"
	JobMgoFieldMsg        = "Msg"
	JobMgoFieldCreateDate = "CreateDate"
	JobMgoFieldUpdateDate = "UpdateDate"
)
const (
	JobMgoSortFieldIDAsc  = "_id"
	JobMgoSortFieldIDDesc = "-_id"
)

func (p *Job) GetNameSpace() string {
	return "digger"
}

func (p *Job) GetClassName() string {
	return "Job"
}

type _JobMgr struct {
}

var JobMgr *_JobMgr

// Get_JobMgr returns the orm manager in case of its name starts with lower letter
func Get_JobMgr() *_JobMgr { return JobMgr }

func (m *_JobMgr) NewJob() *Job {
	rval := new(Job)
	rval.isNew = true
	rval.ID = bson.NewObjectId()

	return rval
}
