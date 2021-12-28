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

	db.SetOnEnsureIndex(initGDsdltIndex)

	RegisterEzOrmObjByID("digger", "GDsdlt", newGDsdltFindByID)
	RegisterEzOrmObjRemove("digger", "GDsdlt", GDsdltMgr.RemoveByID)

}

func initGDsdltIndex() {
	session, collection := GDsdltMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDsdlt SecucodeEndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDsdlt CreateDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "EndDate", "HolderName"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDsdlt SecucodeEndDateHolderName error:" + err.Error())
	}

}

func newGDsdltFindByID(id string) (result EzOrmObj, err error) {
	return GDsdltMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDsdlt []func(obj EzOrmObj)
	updateCB_GDsdlt []func(obj EzOrmObj)
)

func GDsdltAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDsdlt = append(insertCB_GDsdlt, cb)
}

func GDsdltAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDsdlt = append(updateCB_GDsdlt, cb)
}

func (o *GDsdlt) Id() string {
	return o.ID.Hex()
}

func (o *GDsdlt) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDsdltInsertCallback(o)
	} else {
		GDsdltUpdateCallback(o)
	}

	return
}

func (o *GDsdlt) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDsdltMgr.GetCol()
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
		GDsdltInsertCallback(o)
	}
	return
}

func GDsdltInsertCallback(o *GDsdlt) {
	for _, cb := range insertCB_GDsdlt {
		cb(o)
	}
}

func GDsdltUpdateCallback(o *GDsdlt) {
	for _, cb := range updateCB_GDsdlt {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDsdltMgr) FindOne(query interface{}, sortFields ...string) (result *GDsdlt, err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDsdltSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDsdltSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDsdltMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDsdltMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDsdltSort(q, sortFields)
	return session, q
}

func (o *_GDsdltMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDsdltMgr.GetCol()
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
func (o *_GDsdltMgr) FindBySecucodeEndDate(Secucode string, EndDate int64, limit int, offset int, sortFields ...string) (result []*GDsdlt, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDsdltMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDsdltMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GDsdlt, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GDsdltMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDsdltMgr) FindOneBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (result *GDsdlt, err error) {
	query := db.M{
		"Secucode":   Secucode,
		"EndDate":    EndDate,
		"HolderName": HolderName,
	}
	session, q := GDsdltMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDsdltMgr) MustFindOneBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (result *GDsdlt) {
	result, _ = o.FindOneBySecucodeEndDateHolderName(Secucode, EndDate, HolderName)
	if result == nil {
		result = GDsdltMgr.NewGDsdlt()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.HolderName = HolderName
		result.Save()
	}
	return
}

func (o *_GDsdltMgr) RemoveBySecucodeEndDateHolderName(Secucode string, EndDate int64, HolderName string) (err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode":   Secucode,
		"EndDate":    EndDate,
		"HolderName": HolderName,
	}
	return col.Remove(query)
}

func (o *_GDsdltMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDsdlt, err error) {
	session, q := GDsdltMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDsdltMgr) FindAll(query interface{}, sortFields ...string) (result []*GDsdlt, err error) {
	session, q := GDsdltMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDsdltMgr) Has(query interface{}) bool {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDsdltMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDsdltMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDsdltMgr) FindByIDs(id []string, sortFields ...string) (result []*GDsdlt, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDsdltMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDsdltMgr) FindByID(id string) (result *GDsdlt, err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDsdltMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDsdltMgr) RemoveByID(id string) (err error) {
	session, col := GDsdltMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDsdltMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDsdlt")
	}
	return getCol("digger", "digger.GDsdlt")
}

//Search

func (o *GDsdlt) IsSearchEnabled() bool {

	return false

}

//end search
