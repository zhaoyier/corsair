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

	db.SetOnEnsureIndex(initGPDailyIndex)

	RegisterEzOrmObjByID("digger", "GPDaily", newGPDailyFindByID)
	RegisterEzOrmObjRemove("digger", "GPDaily", GPDailyMgr.RemoveByID)

}

func initGPDailyIndex() {
	session, collection := GPDailyMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "CreateDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDaily SecucodeCreateDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDaily Name error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"UpdateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDaily UpdateDate error:" + err.Error())
	}

}

func newGPDailyFindByID(id string) (result EzOrmObj, err error) {
	return GPDailyMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPDaily []func(obj EzOrmObj)
	updateCB_GPDaily []func(obj EzOrmObj)
)

func GPDailyAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPDaily = append(insertCB_GPDaily, cb)
}

func GPDailyAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPDaily = append(updateCB_GPDaily, cb)
}

func (o *GPDaily) Id() string {
	return o.ID.Hex()
}

func (o *GPDaily) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPDailyInsertCallback(o)
	} else {
		GPDailyUpdateCallback(o)
	}

	return
}

func (o *GPDaily) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPDailyMgr.GetCol()
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
		GPDailyInsertCallback(o)
	}
	return
}

func GPDailyInsertCallback(o *GPDaily) {
	for _, cb := range insertCB_GPDaily {
		cb(o)
	}
}

func GPDailyUpdateCallback(o *GPDaily) {
	for _, cb := range updateCB_GPDaily {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPDailyMgr) FindOne(query interface{}, sortFields ...string) (result *GPDaily, err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPDailySort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPDailySort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPDailyMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPDailyMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPDailySort(q, sortFields)
	return session, q
}

func (o *_GPDailyMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPDailyMgr.GetCol()
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
func (o *_GPDailyMgr) FindOneBySecucodeCreateDate(Secucode string, CreateDate string) (result *GPDaily, err error) {
	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	session, q := GPDailyMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPDailyMgr) MustFindOneBySecucodeCreateDate(Secucode string, CreateDate string) (result *GPDaily) {
	result, _ = o.FindOneBySecucodeCreateDate(Secucode, CreateDate)
	if result == nil {
		result = GPDailyMgr.NewGPDaily()
		result.Secucode = Secucode
		result.CreateDate = CreateDate
		result.Save()
	}
	return
}

func (o *_GPDailyMgr) RemoveBySecucodeCreateDate(Secucode string, CreateDate string) (err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	return col.Remove(query)
}
func (o *_GPDailyMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPDaily, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPDailyMgr) FindByUpdateDate(UpdateDate int64, limit int, offset int, sortFields ...string) (result []*GPDaily, err error) {
	query := db.M{
		"UpdateDate": UpdateDate,
	}
	session, q := GPDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDailyMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPDaily, err error) {
	session, q := GPDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDailyMgr) FindAll(query interface{}, sortFields ...string) (result []*GPDaily, err error) {
	session, q := GPDailyMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDailyMgr) Has(query interface{}) bool {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPDailyMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPDailyMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPDailyMgr) FindByIDs(id []string, sortFields ...string) (result []*GPDaily, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPDailyMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPDailyMgr) FindByID(id string) (result *GPDaily, err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPDailyMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPDailyMgr) RemoveByID(id string) (err error) {
	session, col := GPDailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPDailyMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPDaily")
	}
	return getCol("digger", "digger.GPDaily")
}

//Search

func (o *GPDaily) IsSearchEnabled() bool {

	return false

}

//end search