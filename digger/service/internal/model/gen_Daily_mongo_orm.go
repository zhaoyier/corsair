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

	db.SetOnEnsureIndex(initDailyIndex)

	RegisterEzOrmObjByID("digger", "Daily", newDailyFindByID)
	RegisterEzOrmObjRemove("digger", "Daily", DailyMgr.RemoveByID)

}

func initDailyIndex() {
	session, collection := DailyMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.Daily SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.Daily Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.Daily EndDate error:" + err.Error())
	}

}

func newDailyFindByID(id string) (result EzOrmObj, err error) {
	return DailyMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_Daily []func(obj EzOrmObj)
	updateCB_Daily []func(obj EzOrmObj)
)

func DailyAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_Daily = append(insertCB_Daily, cb)
}

func DailyAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_Daily = append(updateCB_Daily, cb)
}

func (o *Daily) Id() string {
	return o.ID.Hex()
}

func (o *Daily) Save() (info *mgo.ChangeInfo, err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		DailyInsertCallback(o)
	} else {
		DailyUpdateCallback(o)
	}

	return
}

func (o *Daily) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := DailyMgr.GetCol()
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
		DailyInsertCallback(o)
	}
	return
}

func DailyInsertCallback(o *Daily) {
	for _, cb := range insertCB_Daily {
		cb(o)
	}
}

func DailyUpdateCallback(o *Daily) {
	for _, cb := range updateCB_Daily {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_DailyMgr) FindOne(query interface{}, sortFields ...string) (result *Daily, err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_DailySort(q, sortFields)

	err = q.One(&result)
	return
}

func _DailySort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_DailyMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := DailyMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_DailySort(q, sortFields)
	return session, q
}

func (o *_DailyMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := DailyMgr.GetCol()
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
func (o *_DailyMgr) FindOneBySecucodeEndDate(Secucode string, EndDate string) (result *Daily, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := DailyMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_DailyMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate string) (result *Daily) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = DailyMgr.NewDaily()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_DailyMgr) RemoveBySecucodeEndDate(Secucode string, EndDate string) (err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}
func (o *_DailyMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*Daily, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := DailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_DailyMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*Daily, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := DailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_DailyMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*Daily, err error) {
	session, q := DailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_DailyMgr) FindAll(query interface{}, sortFields ...string) (result []*Daily, err error) {
	session, q := DailyMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_DailyMgr) Has(query interface{}) bool {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_DailyMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_DailyMgr) CountE(query interface{}) (result int, err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_DailyMgr) FindByIDs(id []string, sortFields ...string) (result []*Daily, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return DailyMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_DailyMgr) FindByID(id string) (result *Daily, err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_DailyMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_DailyMgr) RemoveByID(id string) (err error) {
	session, col := DailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_DailyMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.Daily")
	}
	return getCol("digger", "digger.Daily")
}

//Search

func (o *Daily) IsSearchEnabled() bool {

	return false

}

//end search
