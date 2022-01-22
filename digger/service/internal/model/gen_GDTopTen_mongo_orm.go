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

	db.SetOnEnsureIndex(initGDTopTenIndex)

	RegisterEzOrmObjByID("digger", "GDTopTen", newGDTopTenFindByID)
	RegisterEzOrmObjRemove("digger", "GDTopTen", GDTopTenMgr.RemoveByID)

}

func initGDTopTenIndex() {
	session, collection := GDTopTenMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate", "HolderName"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDTopTen SecucodeEndDateHolderName error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDTopTen SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDTopTen CreateDate error:" + err.Error())
	}

}

func newGDTopTenFindByID(id string) (result EzOrmObj, err error) {
	return GDTopTenMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDTopTen []func(obj EzOrmObj)
	updateCB_GDTopTen []func(obj EzOrmObj)
)

func GDTopTenAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDTopTen = append(insertCB_GDTopTen, cb)
}

func GDTopTenAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDTopTen = append(updateCB_GDTopTen, cb)
}

func (o *GDTopTen) Id() string {
	return o.ID.Hex()
}

func (o *GDTopTen) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDTopTenInsertCallback(o)
	} else {
		GDTopTenUpdateCallback(o)
	}

	return
}

func (o *GDTopTen) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDTopTenMgr.GetCol()
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
		GDTopTenInsertCallback(o)
	}
	return
}

func GDTopTenInsertCallback(o *GDTopTen) {
	for _, cb := range insertCB_GDTopTen {
		cb(o)
	}
}

func GDTopTenUpdateCallback(o *GDTopTen) {
	for _, cb := range updateCB_GDTopTen {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDTopTenMgr) FindOne(query interface{}, sortFields ...string) (result *GDTopTen, err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDTopTenSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDTopTenSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDTopTenMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDTopTenMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDTopTenSort(q, sortFields)
	return session, q
}

func (o *_GDTopTenMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDTopTenMgr.GetCol()
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
func (o *_GDTopTenMgr) FindOneBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (result *GDTopTen, err error) {
	query := db.M{
		"Secucode":   Secucode,
		"EndDate":    EndDate,
		"HolderName": HolderName,
	}
	session, q := GDTopTenMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDTopTenMgr) MustFindOneBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (result *GDTopTen) {
	result, _ = o.FindOneBySecucodeEndDateHolderName(Secucode, EndDate, HolderName)
	if result == nil {
		result = GDTopTenMgr.NewGDTopTen()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.HolderName = HolderName
		result.Save()
	}
	return
}

func (o *_GDTopTenMgr) RemoveBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode":   Secucode,
		"EndDate":    EndDate,
		"HolderName": HolderName,
	}
	return col.Remove(query)
}
func (o *_GDTopTenMgr) FindBySecucodeEndDate(Secucode string, EndDate int64, limit int, offset int, sortFields ...string) (result []*GDTopTen, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDTopTenMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDTopTenMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GDTopTen, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GDTopTenMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDTopTenMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDTopTen, err error) {
	session, q := GDTopTenMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDTopTenMgr) FindAll(query interface{}, sortFields ...string) (result []*GDTopTen, err error) {
	session, q := GDTopTenMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDTopTenMgr) Has(query interface{}) bool {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDTopTenMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDTopTenMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDTopTenMgr) FindByIDs(id []string, sortFields ...string) (result []*GDTopTen, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDTopTenMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDTopTenMgr) FindByID(id string) (result *GDTopTen, err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDTopTenMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDTopTenMgr) RemoveByID(id string) (err error) {
	session, col := GDTopTenMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDTopTenMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDTopTen")
	}
	return getCol("digger", "digger.GDTopTen")
}

//Search

func (o *GDTopTen) IsSearchEnabled() bool {

	return false

}

//end search
