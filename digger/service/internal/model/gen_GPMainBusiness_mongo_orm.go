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

	db.SetOnEnsureIndex(initGPMainBusinessIndex)

	RegisterEzOrmObjByID("digger", "GPMainBusiness", newGPMainBusinessFindByID)
	RegisterEzOrmObjRemove("digger", "GPMainBusiness", GPMainBusinessMgr.RemoveByID)

}

func initGPMainBusinessIndex() {
	session, collection := GPMainBusinessMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPMainBusiness Secucode error:" + err.Error())
	}

}

func newGPMainBusinessFindByID(id string) (result EzOrmObj, err error) {
	return GPMainBusinessMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPMainBusiness []func(obj EzOrmObj)
	updateCB_GPMainBusiness []func(obj EzOrmObj)
)

func GPMainBusinessAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPMainBusiness = append(insertCB_GPMainBusiness, cb)
}

func GPMainBusinessAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPMainBusiness = append(updateCB_GPMainBusiness, cb)
}

func (o *GPMainBusiness) Id() string {
	return o.ID.Hex()
}

func (o *GPMainBusiness) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPMainBusinessInsertCallback(o)
	} else {
		GPMainBusinessUpdateCallback(o)
	}

	return
}

func (o *GPMainBusiness) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPMainBusinessMgr.GetCol()
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
		GPMainBusinessInsertCallback(o)
	}
	return
}

func GPMainBusinessInsertCallback(o *GPMainBusiness) {
	for _, cb := range insertCB_GPMainBusiness {
		cb(o)
	}
}

func GPMainBusinessUpdateCallback(o *GPMainBusiness) {
	for _, cb := range updateCB_GPMainBusiness {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPMainBusinessMgr) FindOne(query interface{}, sortFields ...string) (result *GPMainBusiness, err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPMainBusinessSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPMainBusinessSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPMainBusinessMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPMainBusinessMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPMainBusinessSort(q, sortFields)
	return session, q
}

func (o *_GPMainBusinessMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPMainBusinessMgr.GetCol()
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
func (o *_GPMainBusinessMgr) FindOneBySecucode(Secucode string) (result *GPMainBusiness, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPMainBusinessMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPMainBusinessMgr) MustFindOneBySecucode(Secucode string) (result *GPMainBusiness) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = GPMainBusinessMgr.NewGPMainBusiness()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_GPMainBusinessMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}

func (o *_GPMainBusinessMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPMainBusiness, err error) {
	session, q := GPMainBusinessMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPMainBusinessMgr) FindAll(query interface{}, sortFields ...string) (result []*GPMainBusiness, err error) {
	session, q := GPMainBusinessMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPMainBusinessMgr) Has(query interface{}) bool {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPMainBusinessMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPMainBusinessMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPMainBusinessMgr) FindByIDs(id []string, sortFields ...string) (result []*GPMainBusiness, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPMainBusinessMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPMainBusinessMgr) FindByID(id string) (result *GPMainBusiness, err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPMainBusinessMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPMainBusinessMgr) RemoveByID(id string) (err error) {
	session, col := GPMainBusinessMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPMainBusinessMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPMainBusiness")
	}
	return getCol("digger", "digger.GPMainBusiness")
}

//Search

func (o *GPMainBusiness) IsSearchEnabled() bool {

	return false

}

//end search
