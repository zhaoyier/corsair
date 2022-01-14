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

	db.SetOnEnsureIndex(initGPZhouQiIndex)

	RegisterEzOrmObjByID("digger", "GPZhouQi", newGPZhouQiFindByID)
	RegisterEzOrmObjRemove("digger", "GPZhouQi", GPZhouQiMgr.RemoveByID)

}

func initGPZhouQiIndex() {
	session, collection := GPZhouQiMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPZhouQi Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPZhouQi Name error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPZhouQi Secucode error:" + err.Error())
	}

}

func newGPZhouQiFindByID(id string) (result EzOrmObj, err error) {
	return GPZhouQiMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPZhouQi []func(obj EzOrmObj)
	updateCB_GPZhouQi []func(obj EzOrmObj)
)

func GPZhouQiAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPZhouQi = append(insertCB_GPZhouQi, cb)
}

func GPZhouQiAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPZhouQi = append(updateCB_GPZhouQi, cb)
}

func (o *GPZhouQi) Id() string {
	return o.ID.Hex()
}

func (o *GPZhouQi) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPZhouQiInsertCallback(o)
	} else {
		GPZhouQiUpdateCallback(o)
	}

	return
}

func (o *GPZhouQi) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPZhouQiMgr.GetCol()
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
		GPZhouQiInsertCallback(o)
	}
	return
}

func GPZhouQiInsertCallback(o *GPZhouQi) {
	for _, cb := range insertCB_GPZhouQi {
		cb(o)
	}
}

func GPZhouQiUpdateCallback(o *GPZhouQi) {
	for _, cb := range updateCB_GPZhouQi {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPZhouQiMgr) FindOne(query interface{}, sortFields ...string) (result *GPZhouQi, err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPZhouQiSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPZhouQiSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPZhouQiMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPZhouQiMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPZhouQiSort(q, sortFields)
	return session, q
}

func (o *_GPZhouQiMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPZhouQiMgr.GetCol()
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
func (o *_GPZhouQiMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPZhouQi, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPZhouQiMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPZhouQiMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPZhouQi, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPZhouQiMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPZhouQiMgr) FindOneBySecucode(Secucode string) (result *GPZhouQi, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPZhouQiMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPZhouQiMgr) MustFindOneBySecucode(Secucode string) (result *GPZhouQi) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = GPZhouQiMgr.NewGPZhouQi()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_GPZhouQiMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}

func (o *_GPZhouQiMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPZhouQi, err error) {
	session, q := GPZhouQiMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPZhouQiMgr) FindAll(query interface{}, sortFields ...string) (result []*GPZhouQi, err error) {
	session, q := GPZhouQiMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPZhouQiMgr) Has(query interface{}) bool {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPZhouQiMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPZhouQiMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPZhouQiMgr) FindByIDs(id []string, sortFields ...string) (result []*GPZhouQi, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPZhouQiMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPZhouQiMgr) FindByID(id string) (result *GPZhouQi, err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPZhouQiMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPZhouQiMgr) RemoveByID(id string) (err error) {
	session, col := GPZhouQiMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPZhouQiMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPZhouQi")
	}
	return getCol("digger", "digger.GPZhouQi")
}

//Search

func (o *GPZhouQi) IsSearchEnabled() bool {

	return false

}

//end search
