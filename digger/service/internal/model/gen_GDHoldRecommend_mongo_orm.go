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

	db.SetOnEnsureIndex(initGDHoldRecommendIndex)

	RegisterEzOrmObjByID("digger", "GDHoldRecommend", newGDHoldRecommendFindByID)
	RegisterEzOrmObjRemove("digger", "GDHoldRecommend", GDHoldRecommendMgr.RemoveByID)

}

func initGDHoldRecommendIndex() {
	session, collection := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldRecommend SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Level"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldRecommend Level error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldRecommend EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldRecommend Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDHoldRecommend CreateDate error:" + err.Error())
	}

}

func newGDHoldRecommendFindByID(id string) (result EzOrmObj, err error) {
	return GDHoldRecommendMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDHoldRecommend []func(obj EzOrmObj)
	updateCB_GDHoldRecommend []func(obj EzOrmObj)
)

func GDHoldRecommendAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDHoldRecommend = append(insertCB_GDHoldRecommend, cb)
}

func GDHoldRecommendAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDHoldRecommend = append(updateCB_GDHoldRecommend, cb)
}

func (o *GDHoldRecommend) Id() string {
	return o.ID.Hex()
}

func (o *GDHoldRecommend) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDHoldRecommendInsertCallback(o)
	} else {
		GDHoldRecommendUpdateCallback(o)
	}

	return
}

func (o *GDHoldRecommend) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
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
		GDHoldRecommendInsertCallback(o)
	}
	return
}

func GDHoldRecommendInsertCallback(o *GDHoldRecommend) {
	for _, cb := range insertCB_GDHoldRecommend {
		cb(o)
	}
}

func GDHoldRecommendUpdateCallback(o *GDHoldRecommend) {
	for _, cb := range updateCB_GDHoldRecommend {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDHoldRecommendMgr) FindOne(query interface{}, sortFields ...string) (result *GDHoldRecommend, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDHoldRecommendSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDHoldRecommendSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDHoldRecommendMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDHoldRecommendMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDHoldRecommendSort(q, sortFields)
	return session, q
}

func (o *_GDHoldRecommendMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDHoldRecommendMgr.GetCol()
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
func (o *_GDHoldRecommendMgr) FindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GDHoldRecommend, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDHoldRecommendMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDHoldRecommendMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GDHoldRecommend) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = GDHoldRecommendMgr.NewGDHoldRecommend()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_GDHoldRecommendMgr) RemoveBySecucodeEndDate(Secucode string, EndDate string) (err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}
func (o *_GDHoldRecommendMgr) FindByLevel(Level float64, limit int, offset int, sortFields ...string) (result []*GDHoldRecommend, err error) {
	query := db.M{
		"Level": Level,
	}
	session, q := GDHoldRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldRecommendMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*GDHoldRecommend, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := GDHoldRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldRecommendMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GDHoldRecommend, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GDHoldRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDHoldRecommendMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GDHoldRecommend, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GDHoldRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldRecommendMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDHoldRecommend, err error) {
	session, q := GDHoldRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldRecommendMgr) FindAll(query interface{}, sortFields ...string) (result []*GDHoldRecommend, err error) {
	session, q := GDHoldRecommendMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDHoldRecommendMgr) Has(query interface{}) bool {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDHoldRecommendMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDHoldRecommendMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDHoldRecommendMgr) FindByIDs(id []string, sortFields ...string) (result []*GDHoldRecommend, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDHoldRecommendMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDHoldRecommendMgr) FindByID(id string) (result *GDHoldRecommend, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDHoldRecommendMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDHoldRecommendMgr) RemoveByID(id string) (err error) {
	session, col := GDHoldRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDHoldRecommendMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDHoldRecommend")
	}
	return getCol("digger", "digger.GDHoldRecommend")
}

//Search

func (o *GDHoldRecommend) IsSearchEnabled() bool {

	return false

}

//end search
