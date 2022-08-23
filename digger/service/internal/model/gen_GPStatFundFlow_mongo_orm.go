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

	db.SetOnEnsureIndex(initGPStatFundFlowIndex)

	RegisterEzOrmObjByID("digger", "GPStatFundFlow", newGPStatFundFlowFindByID)
	RegisterEzOrmObjRemove("digger", "GPStatFundFlow", GPStatFundFlowMgr.RemoveByID)

}

func initGPStatFundFlowIndex() {
	session, collection := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStatFundFlow Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Rising"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStatFundFlow Rising error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"UpdateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPStatFundFlow UpdateDate error:" + err.Error())
	}

}

func newGPStatFundFlowFindByID(id string) (result EzOrmObj, err error) {
	return GPStatFundFlowMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPStatFundFlow []func(obj EzOrmObj)
	updateCB_GPStatFundFlow []func(obj EzOrmObj)
)

func GPStatFundFlowAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPStatFundFlow = append(insertCB_GPStatFundFlow, cb)
}

func GPStatFundFlowAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPStatFundFlow = append(updateCB_GPStatFundFlow, cb)
}

func (o *GPStatFundFlow) Id() string {
	return o.ID.Hex()
}

func (o *GPStatFundFlow) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPStatFundFlowInsertCallback(o)
	} else {
		GPStatFundFlowUpdateCallback(o)
	}

	return
}

func (o *GPStatFundFlow) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
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
		GPStatFundFlowInsertCallback(o)
	}
	return
}

func GPStatFundFlowInsertCallback(o *GPStatFundFlow) {
	for _, cb := range insertCB_GPStatFundFlow {
		cb(o)
	}
}

func GPStatFundFlowUpdateCallback(o *GPStatFundFlow) {
	for _, cb := range updateCB_GPStatFundFlow {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPStatFundFlowMgr) FindOne(query interface{}, sortFields ...string) (result *GPStatFundFlow, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPStatFundFlowSort(q, sortFields)

	err = q.One(&result)
	return
}

// _GPStatFundFlowSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _GPStatFundFlowSort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

// Query 按照查询条件、分页、排序等构建 MongoDB 查询对象，默认情况按照插入倒序返回全量数据
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则使用 `-_id` 作为排序条件（注意：如果表数据量很大，请显式传递该字段，否则会发生慢查询）
func (o *_GPStatFundFlowMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPStatFundFlowMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPStatFundFlowSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_GPStatFundFlowMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPStatFundFlowMgr.GetCol()
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
func (o *_GPStatFundFlowMgr) FindOneBySecucode(Secucode string) (result *GPStatFundFlow, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPStatFundFlowMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPStatFundFlowMgr) MustFindOneBySecucode(Secucode string) (result *GPStatFundFlow) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = GPStatFundFlowMgr.NewGPStatFundFlow()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_GPStatFundFlowMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}
func (o *_GPStatFundFlowMgr) FindByRising(Rising int32, limit int, offset int, sortFields ...string) (result []*GPStatFundFlow, err error) {
	query := db.M{
		"Rising": Rising,
	}
	session, q := GPStatFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPStatFundFlowMgr) FindByUpdateDate(UpdateDate int64, limit int, offset int, sortFields ...string) (result []*GPStatFundFlow, err error) {
	query := db.M{
		"UpdateDate": UpdateDate,
	}
	session, q := GPStatFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStatFundFlowMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPStatFundFlow, err error) {
	session, q := GPStatFundFlowMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStatFundFlowMgr) FindAll(query interface{}, sortFields ...string) (result []*GPStatFundFlow, err error) {
	session, q := GPStatFundFlowMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPStatFundFlowMgr) Has(query interface{}) bool {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPStatFundFlowMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPStatFundFlowMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPStatFundFlowMgr) FindByIDs(id []string, sortFields ...string) (result []*GPStatFundFlow, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPStatFundFlowMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPStatFundFlowMgr) FindByID(id string) (result *GPStatFundFlow, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPStatFundFlowMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPStatFundFlowMgr) RemoveByID(id string) (err error) {
	session, col := GPStatFundFlowMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPStatFundFlowMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPStatFundFlow")
	}
	return getCol("digger", "digger.GPStatFundFlow")
}

//Search

func (o *GPStatFundFlow) IsSearchEnabled() bool {

	return false

}

//end search
