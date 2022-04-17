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

	db.SetOnEnsureIndex(initGPManualDecreaseIndex)

	RegisterEzOrmObjByID("digger", "GPManualDecrease", newGPManualDecreaseFindByID)
	RegisterEzOrmObjRemove("digger", "GPManualDecrease", GPManualDecreaseMgr.RemoveByID)

}

func initGPManualDecreaseIndex() {
	session, collection := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPManualDecrease SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPManualDecrease Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPManualDecrease Name error:" + err.Error())
	}

}

func newGPManualDecreaseFindByID(id string) (result EzOrmObj, err error) {
	return GPManualDecreaseMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPManualDecrease []func(obj EzOrmObj)
	updateCB_GPManualDecrease []func(obj EzOrmObj)
)

func GPManualDecreaseAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPManualDecrease = append(insertCB_GPManualDecrease, cb)
}

func GPManualDecreaseAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPManualDecrease = append(updateCB_GPManualDecrease, cb)
}

func (o *GPManualDecrease) Id() string {
	return o.ID.Hex()
}

func (o *GPManualDecrease) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPManualDecreaseInsertCallback(o)
	} else {
		GPManualDecreaseUpdateCallback(o)
	}

	return
}

func (o *GPManualDecrease) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
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
		GPManualDecreaseInsertCallback(o)
	}
	return
}

func GPManualDecreaseInsertCallback(o *GPManualDecrease) {
	for _, cb := range insertCB_GPManualDecrease {
		cb(o)
	}
}

func GPManualDecreaseUpdateCallback(o *GPManualDecrease) {
	for _, cb := range updateCB_GPManualDecrease {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPManualDecreaseMgr) FindOne(query interface{}, sortFields ...string) (result *GPManualDecrease, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPManualDecreaseSort(q, sortFields)

	err = q.One(&result)
	return
}

// _GPManualDecreaseSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _GPManualDecreaseSort(q *mgo.Query, sortFields []string) {
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
func (o *_GPManualDecreaseMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPManualDecreaseMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPManualDecreaseSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_GPManualDecreaseMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPManualDecreaseMgr.GetCol()
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
func (o *_GPManualDecreaseMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPManualDecrease, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPManualDecreaseMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPManualDecreaseMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPManualDecrease) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPManualDecreaseMgr.NewGPManualDecrease()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPManualDecreaseMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPManualDecreaseMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPManualDecrease, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPManualDecreaseMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPManualDecreaseMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPManualDecrease, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPManualDecreaseMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPManualDecreaseMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPManualDecrease, err error) {
	session, q := GPManualDecreaseMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPManualDecreaseMgr) FindAll(query interface{}, sortFields ...string) (result []*GPManualDecrease, err error) {
	session, q := GPManualDecreaseMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPManualDecreaseMgr) Has(query interface{}) bool {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPManualDecreaseMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPManualDecreaseMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPManualDecreaseMgr) FindByIDs(id []string, sortFields ...string) (result []*GPManualDecrease, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPManualDecreaseMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPManualDecreaseMgr) FindByID(id string) (result *GPManualDecrease, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPManualDecreaseMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPManualDecreaseMgr) RemoveByID(id string) (err error) {
	session, col := GPManualDecreaseMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPManualDecreaseMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPManualDecrease")
	}
	return getCol("digger", "digger.GPManualDecrease")
}

//Search

func (o *GPManualDecrease) IsSearchEnabled() bool {

	return false

}

//end search
