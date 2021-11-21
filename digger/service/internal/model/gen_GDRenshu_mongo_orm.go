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

	db.SetOnEnsureIndex(initGDRenshuIndex)

	RegisterEzOrmObjByID("digger", "GDRenshu", newGDRenshuFindByID)
	RegisterEzOrmObjRemove("digger", "GDRenshu", GDRenshuMgr.RemoveByID)

}

func initGDRenshuIndex() {
	session, collection := GDRenshuMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDRenshu Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"SecurityCode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDRenshu SecurityCode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"EndDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDRenshu EndDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"HolderTotalNum"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDRenshu HolderTotalNum error:" + err.Error())
	}

}

func newGDRenshuFindByID(id string) (result EzOrmObj, err error) {
	return GDRenshuMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDRenshu []func(obj EzOrmObj)
	updateCB_GDRenshu []func(obj EzOrmObj)
)

func GDRenshuAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDRenshu = append(insertCB_GDRenshu, cb)
}

func GDRenshuAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDRenshu = append(updateCB_GDRenshu, cb)
}

func (o *GDRenshu) Id() string {
	return o.ID.Hex()
}

func (o *GDRenshu) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDRenshuInsertCallback(o)
	} else {
		GDRenshuUpdateCallback(o)
	}

	return
}

func (o *GDRenshu) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDRenshuMgr.GetCol()
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
		GDRenshuInsertCallback(o)
	}
	return
}

func GDRenshuInsertCallback(o *GDRenshu) {
	for _, cb := range insertCB_GDRenshu {
		cb(o)
	}
}

func GDRenshuUpdateCallback(o *GDRenshu) {
	for _, cb := range updateCB_GDRenshu {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDRenshuMgr) FindOne(query interface{}, sortFields ...string) (result *GDRenshu, err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDRenshuSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDRenshuSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDRenshuMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDRenshuMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDRenshuSort(q, sortFields)
	return session, q
}

func (o *_GDRenshuMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDRenshuMgr.GetCol()
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
func (o *_GDRenshuMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDRenshuMgr) FindBySecurityCode(SecurityCode int32, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	query := db.M{
		"SecurityCode": SecurityCode,
	}
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDRenshuMgr) FindByEndDate(EndDate string, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	query := db.M{
		"EndDate": EndDate,
	}
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDRenshuMgr) FindByHolderTotalNum(HolderTotalNum float64, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	query := db.M{
		"HolderTotalNum": HolderTotalNum,
	}
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDRenshuMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDRenshuMgr) FindAll(query interface{}, sortFields ...string) (result []*GDRenshu, err error) {
	session, q := GDRenshuMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDRenshuMgr) Has(query interface{}) bool {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDRenshuMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDRenshuMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDRenshuMgr) FindByIDs(id []string, sortFields ...string) (result []*GDRenshu, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDRenshuMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDRenshuMgr) FindByID(id string) (result *GDRenshu, err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDRenshuMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDRenshuMgr) RemoveByID(id string) (err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDRenshuMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDRenshu")
	}
	return getCol("digger", "digger.GDRenshu")
}

//Search

func (o *GDRenshu) IsSearchEnabled() bool {

	return false

}

//end search
