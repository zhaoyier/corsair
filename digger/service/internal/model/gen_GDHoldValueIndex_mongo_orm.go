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

	db.SetOnEnsureIndex(initGDHoldValueIndexIndex)

	RegisterEzOrmObjByID("digger", "GDHoldValueIndex", newGDHoldValueIndexFindByID)
	RegisterEzOrmObjRemove("digger", "GDHoldValueIndex", GDHoldValueIndexMgr.RemoveByID)

}

func initGDHoldValueIndexIndex() {
	session, collection := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldValueIndex SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"ValueIndex"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldValueIndex ValueIndex error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldValueIndex EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldValueIndex Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldValueIndex CreateDate error:" + err.Error())
	}

}

func newGDHoldValueIndexFindByID(id string) (result EzOrmObj, err error) {
	return GDHoldValueIndexMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDHoldValueIndex []func(obj EzOrmObj)
	updateCB_GDHoldValueIndex []func(obj EzOrmObj)
)

func GDHoldValueIndexAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDHoldValueIndex = append(insertCB_GDHoldValueIndex, cb)
}

func GDHoldValueIndexAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDHoldValueIndex = append(updateCB_GDHoldValueIndex, cb)
}

func (o *GDHoldValueIndex) Id() string {
	return o.ID.Hex()
}

func (o *GDHoldValueIndex) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDHoldValueIndexInsertCallback(o)
	} else {
		GDHoldValueIndexUpdateCallback(o)
	}

	return
}

func (o *GDHoldValueIndex) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
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
		GDHoldValueIndexInsertCallback(o)
	}
	return
}

func GDHoldValueIndexInsertCallback(o *GDHoldValueIndex) {
	for _, cb := range insertCB_GDHoldValueIndex {
		cb(o)
	}
}

func GDHoldValueIndexUpdateCallback(o *GDHoldValueIndex) {
	for _, cb := range updateCB_GDHoldValueIndex {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDHoldValueIndexMgr) FindOne(query interface{}, sortFields ...string) (result *GDHoldValueIndex, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDHoldValueIndexSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDHoldValueIndexSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDHoldValueIndexMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDHoldValueIndexMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDHoldValueIndexSort(q, sortFields)
	return session, q
}

func (o *_GDHoldValueIndexMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDHoldValueIndexMgr.GetCol()
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
func (o *_GDHoldValueIndexMgr) FindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GDHoldValueIndex, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDHoldValueIndexMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDHoldValueIndexMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GDHoldValueIndex) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = GDHoldValueIndexMgr.NewGDHoldValueIndex()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_GDHoldValueIndexMgr) RemoveBySecucodeEndDate(Secucode string, EndDate string) (err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}
func (o *_GDHoldValueIndexMgr) FindByValueIndex(ValueIndex int32, limit int, offset int, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	query := db.M{
		"ValueIndex": ValueIndex,
	}
	session, q := GDHoldValueIndexMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldValueIndexMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := GDHoldValueIndexMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldValueIndexMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GDHoldValueIndexMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldValueIndexMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GDHoldValueIndexMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldValueIndexMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	session, q := GDHoldValueIndexMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldValueIndexMgr) FindAll(query interface{}, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	session, q := GDHoldValueIndexMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldValueIndexMgr) Has(query interface{}) bool {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDHoldValueIndexMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDHoldValueIndexMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDHoldValueIndexMgr) FindByIDs(id []string, sortFields ...string) (result []*GDHoldValueIndex, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDHoldValueIndexMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDHoldValueIndexMgr) FindByID(id string) (result *GDHoldValueIndex, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDHoldValueIndexMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDHoldValueIndexMgr) RemoveByID(id string) (err error) {
	session, col := GDHoldValueIndexMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDHoldValueIndexMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDHoldValueIndex")
	}
	return getCol("digger", "digger.GDHoldValueIndex")
}

//Search

func (o *GDHoldValueIndex) IsSearchEnabled() bool {

	return false

}

//end search
