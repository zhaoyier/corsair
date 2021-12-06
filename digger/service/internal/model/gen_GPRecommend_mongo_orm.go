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

	db.SetOnEnsureIndex(initGPRecommendIndex)

	RegisterEzOrmObjByID("digger", "GPRecommend", newGPRecommendFindByID)
	RegisterEzOrmObjRemove("digger", "GPRecommend", GPRecommendMgr.RemoveByID)

}

func initGPRecommendIndex() {
	session, collection := GPRecommendMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPRecommend SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "State"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPRecommend SecucodeState error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPRecommend CreateDate error:" + err.Error())
	}

}

func newGPRecommendFindByID(id string) (result EzOrmObj, err error) {
	return GPRecommendMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPRecommend []func(obj EzOrmObj)
	updateCB_GPRecommend []func(obj EzOrmObj)
)

func GPRecommendAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPRecommend = append(insertCB_GPRecommend, cb)
}

func GPRecommendAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPRecommend = append(updateCB_GPRecommend, cb)
}

func (o *GPRecommend) Id() string {
	return o.ID.Hex()
}

func (o *GPRecommend) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPRecommendInsertCallback(o)
	} else {
		GPRecommendUpdateCallback(o)
	}

	return
}

func (o *GPRecommend) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPRecommendMgr.GetCol()
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
		GPRecommendInsertCallback(o)
	}
	return
}

func GPRecommendInsertCallback(o *GPRecommend) {
	for _, cb := range insertCB_GPRecommend {
		cb(o)
	}
}

func GPRecommendUpdateCallback(o *GPRecommend) {
	for _, cb := range updateCB_GPRecommend {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPRecommendMgr) FindOne(query interface{}, sortFields ...string) (result *GPRecommend, err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPRecommendSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPRecommendSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPRecommendMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPRecommendMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPRecommendSort(q, sortFields)
	return session, q
}

func (o *_GPRecommendMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPRecommendMgr.GetCol()
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
func (o *_GPRecommendMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPRecommend, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPRecommendMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPRecommendMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPRecommend) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPRecommendMgr.NewGPRecommend()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPRecommendMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPRecommendMgr) FindBySecucodeState(Secucode string, State int32, limit int, offset int, sortFields ...string) (result []*GPRecommend, err error) {
	query := db.M{
		"Secucode": Secucode,
		"State":    State,
	}
	session, q := GPRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPRecommendMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GPRecommend, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GPRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPRecommendMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPRecommend, err error) {
	session, q := GPRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPRecommendMgr) FindAll(query interface{}, sortFields ...string) (result []*GPRecommend, err error) {
	session, q := GPRecommendMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPRecommendMgr) Has(query interface{}) bool {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPRecommendMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPRecommendMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPRecommendMgr) FindByIDs(id []string, sortFields ...string) (result []*GPRecommend, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPRecommendMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPRecommendMgr) FindByID(id string) (result *GPRecommend, err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPRecommendMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPRecommendMgr) RemoveByID(id string) (err error) {
	session, col := GPRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPRecommendMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPRecommend")
	}
	return getCol("digger", "digger.GPRecommend")
}

//Search

func (o *GPRecommend) IsSearchEnabled() bool {

	return false

}

//end search
