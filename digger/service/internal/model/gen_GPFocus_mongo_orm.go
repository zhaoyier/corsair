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

	db.SetOnEnsureIndex(initGPFocusIndex)

	RegisterEzOrmObjByID("digger", "GPFocus", newGPFocusFindByID)
	RegisterEzOrmObjRemove("digger", "GPFocus", GPFocusMgr.RemoveByID)

}

func initGPFocusIndex() {
	session, collection := GPFocusMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFocus SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFocus Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFocus Name error:" + err.Error())
	}

}

func newGPFocusFindByID(id string) (result EzOrmObj, err error) {
	return GPFocusMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPFocus []func(obj EzOrmObj)
	updateCB_GPFocus []func(obj EzOrmObj)
)

func GPFocusAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPFocus = append(insertCB_GPFocus, cb)
}

func GPFocusAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPFocus = append(updateCB_GPFocus, cb)
}

func (o *GPFocus) Id() string {
	return o.ID.Hex()
}

func (o *GPFocus) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPFocusInsertCallback(o)
	} else {
		GPFocusUpdateCallback(o)
	}

	return
}

func (o *GPFocus) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPFocusMgr.GetCol()
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
		GPFocusInsertCallback(o)
	}
	return
}

func GPFocusInsertCallback(o *GPFocus) {
	for _, cb := range insertCB_GPFocus {
		cb(o)
	}
}

func GPFocusUpdateCallback(o *GPFocus) {
	for _, cb := range updateCB_GPFocus {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPFocusMgr) FindOne(query interface{}, sortFields ...string) (result *GPFocus, err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPFocusSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPFocusSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPFocusMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPFocusMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPFocusSort(q, sortFields)
	return session, q
}

func (o *_GPFocusMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPFocusMgr.GetCol()
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
func (o *_GPFocusMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPFocus, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPFocusMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPFocusMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPFocus) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPFocusMgr.NewGPFocus()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPFocusMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPFocusMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPFocus, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPFocusMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPFocusMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPFocus, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPFocusMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFocusMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPFocus, err error) {
	session, q := GPFocusMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFocusMgr) FindAll(query interface{}, sortFields ...string) (result []*GPFocus, err error) {
	session, q := GPFocusMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFocusMgr) Has(query interface{}) bool {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPFocusMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPFocusMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPFocusMgr) FindByIDs(id []string, sortFields ...string) (result []*GPFocus, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPFocusMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPFocusMgr) FindByID(id string) (result *GPFocus, err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPFocusMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPFocusMgr) RemoveByID(id string) (err error) {
	session, col := GPFocusMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPFocusMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPFocus")
	}
	return getCol("digger", "digger.GPFocus")
}

//Search

func (o *GPFocus) IsSearchEnabled() bool {

	return false

}

//end search
