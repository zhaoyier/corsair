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

	db.SetOnEnsureIndex(initGPWaterfallLineIndex)

	RegisterEzOrmObjByID("digger", "GPWaterfallLine", newGPWaterfallLineFindByID)
	RegisterEzOrmObjRemove("digger", "GPWaterfallLine", GPWaterfallLineMgr.RemoveByID)

}

func initGPWaterfallLineIndex() {
	session, collection := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPWaterfallLine Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPWaterfallLine Name error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPWaterfallLine Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPWaterfallLine Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPWaterfallLine CreateDate error:" + err.Error())
	}

}

func newGPWaterfallLineFindByID(id string) (result EzOrmObj, err error) {
	return GPWaterfallLineMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPWaterfallLine []func(obj EzOrmObj)
	updateCB_GPWaterfallLine []func(obj EzOrmObj)
)

func GPWaterfallLineAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPWaterfallLine = append(insertCB_GPWaterfallLine, cb)
}

func GPWaterfallLineAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPWaterfallLine = append(updateCB_GPWaterfallLine, cb)
}

func (o *GPWaterfallLine) Id() string {
	return o.ID.Hex()
}

func (o *GPWaterfallLine) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPWaterfallLineInsertCallback(o)
	} else {
		GPWaterfallLineUpdateCallback(o)
	}

	return
}

func (o *GPWaterfallLine) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
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
		GPWaterfallLineInsertCallback(o)
	}
	return
}

func GPWaterfallLineInsertCallback(o *GPWaterfallLine) {
	for _, cb := range insertCB_GPWaterfallLine {
		cb(o)
	}
}

func GPWaterfallLineUpdateCallback(o *GPWaterfallLine) {
	for _, cb := range updateCB_GPWaterfallLine {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPWaterfallLineMgr) FindOne(query interface{}, sortFields ...string) (result *GPWaterfallLine, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPWaterfallLineSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPWaterfallLineSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPWaterfallLineMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPWaterfallLineMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPWaterfallLineSort(q, sortFields)
	return session, q
}

func (o *_GPWaterfallLineMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPWaterfallLineMgr.GetCol()
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
func (o *_GPWaterfallLineMgr) FindOneBySecucode(Secucode string) (result *GPWaterfallLine, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPWaterfallLineMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPWaterfallLineMgr) MustFindOneBySecucode(Secucode string) (result *GPWaterfallLine) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = GPWaterfallLineMgr.NewGPWaterfallLine()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_GPWaterfallLineMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}
func (o *_GPWaterfallLineMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPWaterfallLine, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPWaterfallLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPWaterfallLineMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPWaterfallLine, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPWaterfallLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPWaterfallLineMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GPWaterfallLine, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GPWaterfallLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPWaterfallLineMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GPWaterfallLine, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GPWaterfallLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPWaterfallLineMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPWaterfallLine, err error) {
	session, q := GPWaterfallLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPWaterfallLineMgr) FindAll(query interface{}, sortFields ...string) (result []*GPWaterfallLine, err error) {
	session, q := GPWaterfallLineMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPWaterfallLineMgr) Has(query interface{}) bool {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPWaterfallLineMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPWaterfallLineMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPWaterfallLineMgr) FindByIDs(id []string, sortFields ...string) (result []*GPWaterfallLine, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPWaterfallLineMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPWaterfallLineMgr) FindByID(id string) (result *GPWaterfallLine, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPWaterfallLineMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPWaterfallLineMgr) RemoveByID(id string) (err error) {
	session, col := GPWaterfallLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPWaterfallLineMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPWaterfallLine")
	}
	return getCol("digger", "digger.GPWaterfallLine")
}

//Search

func (o *GPWaterfallLine) IsSearchEnabled() bool {

	return false

}

//end search
