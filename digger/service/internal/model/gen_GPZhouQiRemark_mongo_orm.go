package model

import (
	"time"

	//3rd party libs
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//Own libs
	"github.com/ezbuy/ezorm/db"
	. "github.com/ezbuy/ezorm/orm"
)

var _ time.Time

func init() {

	RegisterEzOrmObjByID("digger", "GPZhouQiRemark", newGPZhouQiRemarkFindByID)
	RegisterEzOrmObjRemove("digger", "GPZhouQiRemark", GPZhouQiRemarkMgr.RemoveByID)

}

func newGPZhouQiRemarkFindByID(id string) (result EzOrmObj, err error) {
	return GPZhouQiRemarkMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPZhouQiRemark []func(obj EzOrmObj)
	updateCB_GPZhouQiRemark []func(obj EzOrmObj)
)

func GPZhouQiRemarkAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPZhouQiRemark = append(insertCB_GPZhouQiRemark, cb)
}

func GPZhouQiRemarkAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPZhouQiRemark = append(updateCB_GPZhouQiRemark, cb)
}

func (o *GPZhouQiRemark) Id() string {
	return o.ID.Hex()
}

func (o *GPZhouQiRemark) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPZhouQiRemarkInsertCallback(o)
	} else {
		GPZhouQiRemarkUpdateCallback(o)
	}

	return
}

func (o *GPZhouQiRemark) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	info, err := col.Upsert(query, db.M{"$setOnInsert": o})
	if err != nil {
		return
	}
	if info.Updated == 0 {
		saved = true
	}
	o.isNew = false
	if saved {
		GPZhouQiRemarkInsertCallback(o)
	}
	return
}

func GPZhouQiRemarkInsertCallback(o *GPZhouQiRemark) {
	for _, cb := range insertCB_GPZhouQiRemark {
		cb(o)
	}
}

func GPZhouQiRemarkUpdateCallback(o *GPZhouQiRemark) {
	for _, cb := range updateCB_GPZhouQiRemark {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPZhouQiRemarkMgr) FindOne(query interface{}, sortFields ...string) (result *GPZhouQiRemark, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPZhouQiRemarkSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPZhouQiRemarkSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPZhouQiRemarkMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPZhouQiRemarkSort(q, sortFields)
	return session, q
}

func (o *_GPZhouQiRemarkMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	if sortFields = XSortFieldsFilter(sortFields); len(sortFields) > 0 {
		q.Sort(sortFields...)
	}

	return session, q
}

func (o *_GPZhouQiRemarkMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPZhouQiRemark, err error) {
	session, q := GPZhouQiRemarkMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPZhouQiRemarkMgr) FindAll(query interface{}, sortFields ...string) (result []*GPZhouQiRemark, err error) {
	session, q := GPZhouQiRemarkMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPZhouQiRemarkMgr) Has(query interface{}) bool {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPZhouQiRemarkMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPZhouQiRemarkMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPZhouQiRemarkMgr) FindByIDs(id []string, sortFields ...string) (result []*GPZhouQiRemark, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPZhouQiRemarkMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPZhouQiRemarkMgr) FindByID(id string) (result *GPZhouQiRemark, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPZhouQiRemarkMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPZhouQiRemarkMgr) RemoveByID(id string) (err error) {
	session, col := GPZhouQiRemarkMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPZhouQiRemarkMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("ezseller", "digger.GPZhouQiRemark")
	}
	return getCol("ezseller", "digger.GPZhouQiRemark")
}

//Search

func (o *GPZhouQiRemark) IsSearchEnabled() bool {

	return false

}

//end search
