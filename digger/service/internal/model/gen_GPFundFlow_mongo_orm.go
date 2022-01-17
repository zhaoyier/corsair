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

	db.SetOnEnsureIndex(initGPFundFlowIndex)

	RegisterEzOrmObjByID("digger", "GPFundFlow", newGPFundFlowFindByID)
	RegisterEzOrmObjRemove("digger", "GPFundFlow", GPFundFlowMgr.RemoveByID)

}

func initGPFundFlowIndex() {
	session, collection := GPFundFlowMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "FundDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFundFlow SecucodeFundDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFundFlow Name error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFundFlow Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"FundDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFundFlow FundDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"InflowRatio"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPFundFlow InflowRatio error:" + err.Error())
	}

}

func newGPFundFlowFindByID(id string) (result EzOrmObj, err error) {
	return GPFundFlowMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPFundFlow []func(obj EzOrmObj)
	updateCB_GPFundFlow []func(obj EzOrmObj)
)

func GPFundFlowAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPFundFlow = append(insertCB_GPFundFlow, cb)
}

func GPFundFlowAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPFundFlow = append(updateCB_GPFundFlow, cb)
}

func (o *GPFundFlow) Id() string {
	return o.ID.Hex()
}

func (o *GPFundFlow) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPFundFlowInsertCallback(o)
	} else {
		GPFundFlowUpdateCallback(o)
	}

	return
}

func (o *GPFundFlow) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPFundFlowMgr.GetCol()
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
		GPFundFlowInsertCallback(o)
	}
	return
}

func GPFundFlowInsertCallback(o *GPFundFlow) {
	for _, cb := range insertCB_GPFundFlow {
		cb(o)
	}
}

func GPFundFlowUpdateCallback(o *GPFundFlow) {
	for _, cb := range updateCB_GPFundFlow {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPFundFlowMgr) FindOne(query interface{}, sortFields ...string) (result *GPFundFlow, err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPFundFlowSort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPFundFlowSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPFundFlowMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPFundFlowMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPFundFlowSort(q, sortFields)
	return session, q
}

func (o *_GPFundFlowMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPFundFlowMgr.GetCol()
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
func (o *_GPFundFlowMgr) FindOneBySecucodeFundDate(Secucode string, FundDate int64) (result *GPFundFlow, err error) {
	query := db.M{
		"Secucode": Secucode,
		"FundDate": FundDate,
	}
	session, q := GPFundFlowMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPFundFlowMgr) MustFindOneBySecucodeFundDate(Secucode string, FundDate int64) (result *GPFundFlow) {
	result, _ = o.FindOneBySecucodeFundDate(Secucode, FundDate)
	if result == nil {
		result = GPFundFlowMgr.NewGPFundFlow()
		result.Secucode = Secucode
		result.FundDate = FundDate
		result.Save()
	}
	return
}

func (o *_GPFundFlowMgr) RemoveBySecucodeFundDate(Secucode string, FundDate int64) (err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"FundDate": FundDate,
	}
	return col.Remove(query)
}
func (o *_GPFundFlowMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPFundFlow, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPFundFlowMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPFundFlow, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPFundFlowMgr) FindByFundDate(FundDate int64, limit int, offset int, sortFields ...string) (result []*GPFundFlow, err error) {
	query := db.M{
		"FundDate": FundDate,
	}
	session, q := GPFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPFundFlowMgr) FindByInflowRatio(InflowRatio int32, limit int, offset int, sortFields ...string) (result []*GPFundFlow, err error) {
	query := db.M{
		"InflowRatio": InflowRatio,
	}
	session, q := GPFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFundFlowMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPFundFlow, err error) {
	session, q := GPFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFundFlowMgr) FindAll(query interface{}, sortFields ...string) (result []*GPFundFlow, err error) {
	session, q := GPFundFlowMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPFundFlowMgr) Has(query interface{}) bool {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPFundFlowMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPFundFlowMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPFundFlowMgr) FindByIDs(id []string, sortFields ...string) (result []*GPFundFlow, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPFundFlowMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPFundFlowMgr) FindByID(id string) (result *GPFundFlow, err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPFundFlowMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPFundFlowMgr) RemoveByID(id string) (err error) {
	session, col := GPFundFlowMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPFundFlowMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPFundFlow")
	}
	return getCol("digger", "digger.GPFundFlow")
}

//Search

func (o *GPFundFlow) IsSearchEnabled() bool {

	return false

}

//end search
