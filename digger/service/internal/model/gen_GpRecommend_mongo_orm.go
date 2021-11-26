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

	db.SetOnEnsureIndex(initGpRecommendIndex)

	RegisterEzOrmObjByID("digger", "GpRecommend", newGpRecommendFindByID)
	RegisterEzOrmObjRemove("digger", "GpRecommend", GpRecommendMgr.RemoveByID)

}

func initGpRecommendIndex() {
	session, collection := GpRecommendMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GpRecommend SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Level"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GpRecommend Level error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GpRecommend EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GpRecommend Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GpRecommend CreateDate error:" + err.Error())
	}

}

func newGpRecommendFindByID(id string) (result EzOrmObj, err error) {
	return GpRecommendMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GpRecommend []func(obj EzOrmObj)
	updateCB_GpRecommend []func(obj EzOrmObj)
)

func GpRecommendAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GpRecommend = append(insertCB_GpRecommend, cb)
}

func GpRecommendAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GpRecommend = append(updateCB_GpRecommend, cb)
}

func (o *GpRecommend) Id() string {
	return o.ID.Hex()
}

func (o *GpRecommend) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GpRecommendInsertCallback(o)
	} else {
		GpRecommendUpdateCallback(o)
	}

	return
}

func (o *GpRecommend) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GpRecommendMgr.GetCol()
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
		GpRecommendInsertCallback(o)
	}
	return
}

func GpRecommendInsertCallback(o *GpRecommend) {
	for _, cb := range insertCB_GpRecommend {
		cb(o)
	}
}

func GpRecommendUpdateCallback(o *GpRecommend) {
	for _, cb := range updateCB_GpRecommend {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GpRecommendMgr) FindOne(query interface{}, sortFields ...string) (result *GpRecommend, err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GpRecommendSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GpRecommendSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GpRecommendMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GpRecommendMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GpRecommendSort(q, sortFields)
	return session, q
}

func (o *_GpRecommendMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GpRecommendMgr.GetCol()
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
func (o *_GpRecommendMgr) FindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GpRecommend, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GpRecommendMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GpRecommendMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate string) (result *GpRecommend) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = GpRecommendMgr.NewGpRecommend()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_GpRecommendMgr) RemoveBySecucodeEndDate(Secucode string, EndDate string) (err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
}
func (o *_GpRecommendMgr) FindByLevel(Level float64, limit int, offset int, sortFields ...string) (result []*GpRecommend, err error) {
	query := db.M{
		"Level": Level,
	}
	session, q := GpRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GpRecommendMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*GpRecommend, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := GpRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GpRecommendMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GpRecommend, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GpRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GpRecommendMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GpRecommend, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GpRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GpRecommendMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GpRecommend, err error) {
	session, q := GpRecommendMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GpRecommendMgr) FindAll(query interface{}, sortFields ...string) (result []*GpRecommend, err error) {
	session, q := GpRecommendMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GpRecommendMgr) Has(query interface{}) bool {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GpRecommendMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GpRecommendMgr) CountE(query interface{}) (result int, err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GpRecommendMgr) FindByIDs(id []string, sortFields ...string) (result []*GpRecommend, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GpRecommendMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GpRecommendMgr) FindByID(id string) (result *GpRecommend, err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GpRecommendMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GpRecommendMgr) RemoveByID(id string) (err error) {
	session, col := GpRecommendMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GpRecommendMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GpRecommend")
	}
	return getCol("digger", "digger.GpRecommend")
}

//Search

func (o *GpRecommend) IsSearchEnabled() bool {

	return false

}

//end search
