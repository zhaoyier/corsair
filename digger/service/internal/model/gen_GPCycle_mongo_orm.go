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

	db.SetOnEnsureIndex(initGPCycleIndex)

	RegisterEzOrmObjByID("digger", "GPCycle", newGPCycleFindByID)
	RegisterEzOrmObjRemove("digger", "GPCycle", GPCycleMgr.RemoveByID)

}

func initGPCycleIndex() {
	session, collection := GPCycleMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPCycle SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPCycle Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPCycle Name error:" + err.Error())
	}

}

func newGPCycleFindByID(id string) (result EzOrmObj, err error) {
	return GPCycleMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPCycle []func(obj EzOrmObj)
	updateCB_GPCycle []func(obj EzOrmObj)
)

func GPCycleAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPCycle = append(insertCB_GPCycle, cb)
}

func GPCycleAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPCycle = append(updateCB_GPCycle, cb)
}

func (o *GPCycle) Id() string {
	return o.ID.Hex()
}

func (o *GPCycle) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPCycleInsertCallback(o)
	} else {
		GPCycleUpdateCallback(o)
	}

	return
}

func (o *GPCycle) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPCycleMgr.GetCol()
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
		GPCycleInsertCallback(o)
	}
	return
}

func GPCycleInsertCallback(o *GPCycle) {
	for _, cb := range insertCB_GPCycle {
		cb(o)
	}
}

func GPCycleUpdateCallback(o *GPCycle) {
	for _, cb := range updateCB_GPCycle {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPCycleMgr) FindOne(query interface{}, sortFields ...string) (result *GPCycle, err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPCycleSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPCycleSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPCycleMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPCycleMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPCycleSort(q, sortFields)
	return session, q
}

func (o *_GPCycleMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPCycleMgr.GetCol()
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
func (o *_GPCycleMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPCycle, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPCycleMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPCycleMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPCycle) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPCycleMgr.NewGPCycle()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPCycleMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPCycleMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPCycle, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPCycleMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPCycleMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPCycle, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPCycleMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPCycleMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPCycle, err error) {
	session, q := GPCycleMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPCycleMgr) FindAll(query interface{}, sortFields ...string) (result []*GPCycle, err error) {
	session, q := GPCycleMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPCycleMgr) Has(query interface{}) bool {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPCycleMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPCycleMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPCycleMgr) FindByIDs(id []string, sortFields ...string) (result []*GPCycle, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPCycleMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPCycleMgr) FindByID(id string) (result *GPCycle, err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPCycleMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPCycleMgr) RemoveByID(id string) (err error) {
	session, col := GPCycleMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPCycleMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPCycle")
	}
	return getCol("digger", "digger.GPCycle")
}

//Search

func (o *GPCycle) IsSearchEnabled() bool {

	return false

}

//end search
