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

	db.SetOnEnsureIndex(initGDLongLineIndex)

	RegisterEzOrmObjByID("digger", "GDLongLine", newGDLongLineFindByID)
	RegisterEzOrmObjRemove("digger", "GDLongLine", GDLongLineMgr.RemoveByID)

}

func initGDLongLineIndex() {
	session, collection := GDLongLineMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"ValueIndex"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDLongLine ValueIndex error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDLongLine EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDLongLine Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDLongLine CreateDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDLongLine SecucodeEndDate error:" + err.Error())
	}

}

func newGDLongLineFindByID(id string) (result EzOrmObj, err error) {
	return GDLongLineMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDLongLine []func(obj EzOrmObj)
	updateCB_GDLongLine []func(obj EzOrmObj)
)

func GDLongLineAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDLongLine = append(insertCB_GDLongLine, cb)
}

func GDLongLineAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDLongLine = append(updateCB_GDLongLine, cb)
}

func (o *GDLongLine) Id() string {
	return o.ID.Hex()
}

func (o *GDLongLine) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDLongLineInsertCallback(o)
	} else {
		GDLongLineUpdateCallback(o)
	}

	return
}

func (o *GDLongLine) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDLongLineMgr.GetCol()
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
		GDLongLineInsertCallback(o)
	}
	return
}

func GDLongLineInsertCallback(o *GDLongLine) {
	for _, cb := range insertCB_GDLongLine {
		cb(o)
	}
}

func GDLongLineUpdateCallback(o *GDLongLine) {
	for _, cb := range updateCB_GDLongLine {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDLongLineMgr) FindOne(query interface{}, sortFields ...string) (result *GDLongLine, err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDLongLineSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDLongLineSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDLongLineMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDLongLineMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDLongLineSort(q, sortFields)
	return session, q
}

func (o *_GDLongLineMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDLongLineMgr.GetCol()
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
func (o *_GDLongLineMgr) FindByValueIndex(ValueIndex int32, limit int, offset int, sortFields ...string) (result []*GDLongLine, err error) {
	query := db.M{
		"ValueIndex": ValueIndex,
	}
	session, q := GDLongLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDLongLineMgr) FindByEndDate(EndDate int64, limit int, offset int, sortFields ...string) (result []*GDLongLine, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := GDLongLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDLongLineMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GDLongLine, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GDLongLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDLongLineMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GDLongLine, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GDLongLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDLongLineMgr) FindOneBySecucodeEndDate(Secucode string, EndDate int64) (result *GDLongLine, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDLongLineMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDLongLineMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate int64) (result *GDLongLine) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = GDLongLineMgr.NewGDLongLine()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_GDLongLineMgr) RemoveBySecucodeEndDate(Secucode string, EndDate int64) (err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}

func (o *_GDLongLineMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDLongLine, err error) {
	session, q := GDLongLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDLongLineMgr) FindAll(query interface{}, sortFields ...string) (result []*GDLongLine, err error) {
	session, q := GDLongLineMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDLongLineMgr) Has(query interface{}) bool {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDLongLineMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDLongLineMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDLongLineMgr) FindByIDs(id []string, sortFields ...string) (result []*GDLongLine, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDLongLineMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDLongLineMgr) FindByID(id string) (result *GDLongLine, err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDLongLineMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDLongLineMgr) RemoveByID(id string) (err error) {
	session, col := GDLongLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDLongLineMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDLongLine")
	}
	return getCol("digger", "digger.GDLongLine")
}

//Search

func (o *GDLongLine) IsSearchEnabled() bool {

	return false

}

//end search
