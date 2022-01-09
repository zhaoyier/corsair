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

	RegisterEzOrmObjByID("digger", "GPRemark", newGPRemarkFindByID)
	RegisterEzOrmObjRemove("digger", "GPRemark", GPRemarkMgr.RemoveByID)

}

func newGPRemarkFindByID(id string) (result EzOrmObj, err error) {
	return GPRemarkMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPRemark []func(obj EzOrmObj)
	updateCB_GPRemark []func(obj EzOrmObj)
)

func GPRemarkAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPRemark = append(insertCB_GPRemark, cb)
}

func GPRemarkAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPRemark = append(updateCB_GPRemark, cb)
}

func (o *GPRemark) Id() string {
	return o.ID.Hex()
}

func (o *GPRemark) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPRemarkInsertCallback(o)
	} else {
		GPRemarkUpdateCallback(o)
	}

	return
}

func (o *GPRemark) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPRemarkMgr.GetCol()
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
		GPRemarkInsertCallback(o)
	}
	return
}

func GPRemarkInsertCallback(o *GPRemark) {
	for _, cb := range insertCB_GPRemark {
		cb(o)
	}
}

func GPRemarkUpdateCallback(o *GPRemark) {
	for _, cb := range updateCB_GPRemark {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPRemarkMgr) FindOne(query interface{}, sortFields ...string) (result *GPRemark, err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPRemarkSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPRemarkSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPRemarkMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPRemarkMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPRemarkSort(q, sortFields)
	return session, q
}

func (o *_GPRemarkMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPRemarkMgr.GetCol()
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

func (o *_GPRemarkMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPRemark, err error) {
	session, q := GPRemarkMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPRemarkMgr) FindAll(query interface{}, sortFields ...string) (result []*GPRemark, err error) {
	session, q := GPRemarkMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPRemarkMgr) Has(query interface{}) bool {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPRemarkMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPRemarkMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPRemarkMgr) FindByIDs(id []string, sortFields ...string) (result []*GPRemark, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPRemarkMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPRemarkMgr) FindByID(id string) (result *GPRemark, err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPRemarkMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPRemarkMgr) RemoveByID(id string) (err error) {
	session, col := GPRemarkMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPRemarkMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("", "digger.GPRemark")
	}
	return getCol("", "digger.GPRemark")
}

//Search

func (o *GPRemark) IsSearchEnabled() bool {

	return false

}

//end search
