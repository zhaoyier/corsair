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

	db.SetOnEnsureIndex(initGDNiusanIndex)

	RegisterEzOrmObjByID("digger", "GDNiusan", newGDNiusanFindByID)
	RegisterEzOrmObjRemove("digger", "GDNiusan", GDNiusanMgr.RemoveByID)

}

func initGDNiusanIndex() {
	session, collection := GDNiusanMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"SecurityCode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDNiusan SecurityCode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Niusan"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDNiusan Niusan error:" + err.Error())
	}

}

func newGDNiusanFindByID(id string) (result EzOrmObj, err error) {
	return GDNiusanMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDNiusan []func(obj EzOrmObj)
	updateCB_GDNiusan []func(obj EzOrmObj)
)

func GDNiusanAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDNiusan = append(insertCB_GDNiusan, cb)
}

func GDNiusanAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDNiusan = append(updateCB_GDNiusan, cb)
}

func (o *GDNiusan) Id() string {
	return o.ID.Hex()
}

func (o *GDNiusan) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDNiusanInsertCallback(o)
	} else {
		GDNiusanUpdateCallback(o)
	}

	return
}

func (o *GDNiusan) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDNiusanMgr.GetCol()
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
		GDNiusanInsertCallback(o)
	}
	return
}

func GDNiusanInsertCallback(o *GDNiusan) {
	for _, cb := range insertCB_GDNiusan {
		cb(o)
	}
}

func GDNiusanUpdateCallback(o *GDNiusan) {
	for _, cb := range updateCB_GDNiusan {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDNiusanMgr) FindOne(query interface{}, sortFields ...string) (result *GDNiusan, err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDNiusanSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GDNiusanSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GDNiusanMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDNiusanMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDNiusanSort(q, sortFields)
	return session, q
}

func (o *_GDNiusanMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDNiusanMgr.GetCol()
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
func (o *_GDNiusanMgr) FindBySecurityCode(SecurityCode int32, limit int, offset int, sortFields ...string) (result []*GDNiusan, err error) {
	query := db.M{
		"SecurityCode": SecurityCode,
	}
	session, q := GDNiusanMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDNiusanMgr) FindByNiusan(Niusan string, limit int, offset int, sortFields ...string) (result []*GDNiusan, err error) {
	query := db.M{
		"Niusan": Niusan,
	}
	session, q := GDNiusanMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDNiusanMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDNiusan, err error) {
	session, q := GDNiusanMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDNiusanMgr) FindAll(query interface{}, sortFields ...string) (result []*GDNiusan, err error) {
	session, q := GDNiusanMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDNiusanMgr) Has(query interface{}) bool {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDNiusanMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDNiusanMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDNiusanMgr) FindByIDs(id []string, sortFields ...string) (result []*GDNiusan, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDNiusanMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDNiusanMgr) FindByID(id string) (result *GDNiusan, err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDNiusanMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDNiusanMgr) RemoveByID(id string) (err error) {
	session, col := GDNiusanMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDNiusanMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDNiusan")
	}
	return getCol("digger", "digger.GDNiusan")
}

//Search

func (o *GDNiusan) IsSearchEnabled() bool {

	return false

}

//end search
