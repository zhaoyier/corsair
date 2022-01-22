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

	db.SetOnEnsureIndex(initSinaDailyIndex)

	RegisterEzOrmObjByID("digger", "SinaDaily", newSinaDailyFindByID)
	RegisterEzOrmObjRemove("digger", "SinaDaily", SinaDailyMgr.RemoveByID)

}

func initSinaDailyIndex() {
	session, collection := SinaDailyMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.SinaDaily Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.SinaDaily EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.SinaDaily SecucodeEndDate error:" + err.Error())
	}

}

func newSinaDailyFindByID(id string) (result EzOrmObj, err error) {
	return SinaDailyMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_SinaDaily []func(obj EzOrmObj)
	updateCB_SinaDaily []func(obj EzOrmObj)
)

func SinaDailyAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_SinaDaily = append(insertCB_SinaDaily, cb)
}

func SinaDailyAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_SinaDaily = append(updateCB_SinaDaily, cb)
}

func (o *SinaDaily) Id() string {
	return o.ID.Hex()
}

func (o *SinaDaily) Save() (info *mgo.ChangeInfo, err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		SinaDailyInsertCallback(o)
	} else {
		SinaDailyUpdateCallback(o)
	}

	return
}

func (o *SinaDaily) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := SinaDailyMgr.GetCol()
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
		SinaDailyInsertCallback(o)
	}
	return
}

func SinaDailyInsertCallback(o *SinaDaily) {
	for _, cb := range insertCB_SinaDaily {
		cb(o)
	}
}

func SinaDailyUpdateCallback(o *SinaDaily) {
	for _, cb := range updateCB_SinaDaily {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_SinaDailyMgr) FindOne(query interface{}, sortFields ...string) (result *SinaDaily, err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_SinaDailySort(q, sortFields)

	err = q.One(&result)
	return
}

func _SinaDailySort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_SinaDailyMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := SinaDailyMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_SinaDailySort(q, sortFields)
	return session, q
}

func (o *_SinaDailyMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := SinaDailyMgr.GetCol()
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
func (o *_SinaDailyMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*SinaDaily, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := SinaDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_SinaDailyMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*SinaDaily, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := SinaDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_SinaDailyMgr) FindOneBySecucodeEndDate(Secucode string, EndDate string) (result *SinaDaily, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := SinaDailyMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_SinaDailyMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate string) (result *SinaDaily) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = SinaDailyMgr.NewSinaDaily()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_SinaDailyMgr) RemoveBySecucodeEndDate(Secucode string, EndDate string) (err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}

func (o *_SinaDailyMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*SinaDaily, err error) {
	session, q := SinaDailyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_SinaDailyMgr) FindAll(query interface{}, sortFields ...string) (result []*SinaDaily, err error) {
	session, q := SinaDailyMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_SinaDailyMgr) Has(query interface{}) bool {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_SinaDailyMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_SinaDailyMgr) CountE(query interface{}) (result int, err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_SinaDailyMgr) FindByIDs(id []string, sortFields ...string) (result []*SinaDaily, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return SinaDailyMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_SinaDailyMgr) FindByID(id string) (result *SinaDaily, err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_SinaDailyMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_SinaDailyMgr) RemoveByID(id string) (err error) {
	session, col := SinaDailyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_SinaDailyMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.SinaDaily")
	}
	return getCol("digger", "digger.SinaDaily")
}

//Search

func (o *SinaDaily) IsSearchEnabled() bool {

	return false

}

//end search
