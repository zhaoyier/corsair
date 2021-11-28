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

	db.SetOnEnsureIndex(initGPStubIncidentIndex)

	RegisterEzOrmObjByID("digger", "GPStubIncident", newGPStubIncidentFindByID)
	RegisterEzOrmObjRemove("digger", "GPStubIncident", GPStubIncidentMgr.RemoveByID)

}

func initGPStubIncidentIndex() {
	session, collection := GPStubIncidentMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "CreateDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStubIncident SecucodeCreateDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStubIncident Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"UpdateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStubIncident UpdateDate error:" + err.Error())
	}

}

func newGPStubIncidentFindByID(id string) (result EzOrmObj, err error) {
	return GPStubIncidentMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPStubIncident []func(obj EzOrmObj)
	updateCB_GPStubIncident []func(obj EzOrmObj)
)

func GPStubIncidentAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPStubIncident = append(insertCB_GPStubIncident, cb)
}

func GPStubIncidentAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPStubIncident = append(updateCB_GPStubIncident, cb)
}

func (o *GPStubIncident) Id() string {
	return o.ID.Hex()
}

func (o *GPStubIncident) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPStubIncidentInsertCallback(o)
	} else {
		GPStubIncidentUpdateCallback(o)
	}

	return
}

func (o *GPStubIncident) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPStubIncidentMgr.GetCol()
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
		GPStubIncidentInsertCallback(o)
	}
	return
}

func GPStubIncidentInsertCallback(o *GPStubIncident) {
	for _, cb := range insertCB_GPStubIncident {
		cb(o)
	}
}

func GPStubIncidentUpdateCallback(o *GPStubIncident) {
	for _, cb := range updateCB_GPStubIncident {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPStubIncidentMgr) FindOne(query interface{}, sortFields ...string) (result *GPStubIncident, err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPStubIncidentSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPStubIncidentSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPStubIncidentMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPStubIncidentMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPStubIncidentSort(q, sortFields)
	return session, q
}

func (o *_GPStubIncidentMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPStubIncidentMgr.GetCol()
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
func (o *_GPStubIncidentMgr) FindOneBySecucodeCreateDate(Secucode string, CreateDate int64) (result *GPStubIncident, err error) {
	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	session, q := GPStubIncidentMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPStubIncidentMgr) MustFindOneBySecucodeCreateDate(Secucode string, CreateDate int64) (result *GPStubIncident) {
	result, _ = o.FindOneBySecucodeCreateDate(Secucode, CreateDate)
	if result == nil {
		result = GPStubIncidentMgr.NewGPStubIncident()
		result.Secucode = Secucode
		result.CreateDate = CreateDate
		result.Save()
	}
	return
}

func (o *_GPStubIncidentMgr) RemoveBySecucodeCreateDate(Secucode string, CreateDate int64) (err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	return col.Remove(query)
}
func (o *_GPStubIncidentMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPStubIncident, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPStubIncidentMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPStubIncidentMgr) FindByUpdateDate(UpdateDate int64, limit int, offset int, sortFields ...string) (result []*GPStubIncident, err error) {
	query := db.M{
		"UpdateDate": UpdateDate,
	}
	session, q := GPStubIncidentMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStubIncidentMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPStubIncident, err error) {
	session, q := GPStubIncidentMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStubIncidentMgr) FindAll(query interface{}, sortFields ...string) (result []*GPStubIncident, err error) {
	session, q := GPStubIncidentMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStubIncidentMgr) Has(query interface{}) bool {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPStubIncidentMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPStubIncidentMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPStubIncidentMgr) FindByIDs(id []string, sortFields ...string) (result []*GPStubIncident, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPStubIncidentMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPStubIncidentMgr) FindByID(id string) (result *GPStubIncident, err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPStubIncidentMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPStubIncidentMgr) RemoveByID(id string) (err error) {
	session, col := GPStubIncidentMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPStubIncidentMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPStubIncident")
	}
	return getCol("digger", "digger.GPStubIncident")
}

//Search

func (o *GPStubIncident) IsSearchEnabled() bool {

	return false

}

//end search
