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

	db.SetOnEnsureIndex(initGPDelayIndex)

	RegisterEzOrmObjByID("digger", "GPDelay", newGPDelayFindByID)
	RegisterEzOrmObjRemove("digger", "GPDelay", GPDelayMgr.RemoveByID)

}

func initGPDelayIndex() {
	session, collection := GPDelayMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDelay SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDelay Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPDelay Name error:" + err.Error())
	}

}

func newGPDelayFindByID(id string) (result EzOrmObj, err error) {
	return GPDelayMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPDelay []func(obj EzOrmObj)
	updateCB_GPDelay []func(obj EzOrmObj)
)

func GPDelayAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPDelay = append(insertCB_GPDelay, cb)
}

func GPDelayAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPDelay = append(updateCB_GPDelay, cb)
}

func (o *GPDelay) Id() string {
	return o.ID.Hex()
}

func (o *GPDelay) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPDelayInsertCallback(o)
	} else {
		GPDelayUpdateCallback(o)
	}

	return
}

func (o *GPDelay) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPDelayMgr.GetCol()
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
		GPDelayInsertCallback(o)
	}
	return
}

func GPDelayInsertCallback(o *GPDelay) {
	for _, cb := range insertCB_GPDelay {
		cb(o)
	}
}

func GPDelayUpdateCallback(o *GPDelay) {
	for _, cb := range updateCB_GPDelay {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPDelayMgr) FindOne(query interface{}, sortFields ...string) (result *GPDelay, err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPDelaySort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPDelaySort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPDelayMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPDelayMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPDelaySort(q, sortFields)
	return session, q
}

func (o *_GPDelayMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPDelayMgr.GetCol()
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
func (o *_GPDelayMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPDelay, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPDelayMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPDelayMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPDelay) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPDelayMgr.NewGPDelay()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPDelayMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPDelayMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPDelay, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPDelayMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPDelayMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPDelay, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPDelayMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDelayMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPDelay, err error) {
	session, q := GPDelayMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDelayMgr) FindAll(query interface{}, sortFields ...string) (result []*GPDelay, err error) {
	session, q := GPDelayMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPDelayMgr) Has(query interface{}) bool {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPDelayMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPDelayMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPDelayMgr) FindByIDs(id []string, sortFields ...string) (result []*GPDelay, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPDelayMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPDelayMgr) FindByID(id string) (result *GPDelay, err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPDelayMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPDelayMgr) RemoveByID(id string) (err error) {
	session, col := GPDelayMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPDelayMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPDelay")
	}
	return getCol("digger", "digger.GPDelay")
}

//Search

func (o *GPDelay) IsSearchEnabled() bool {

	return false

}

//end search
