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

	db.SetOnEnsureIndex(initGPShortLineIndex)

	RegisterEzOrmObjByID("digger", "GPShortLine", newGPShortLineFindByID)
	RegisterEzOrmObjRemove("digger", "GPShortLine", GPShortLineMgr.RemoveByID)

}

func initGPShortLineIndex() {
	session, collection := GPShortLineMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "CreateDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPShortLine SecucodeCreateDate error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPShortLine Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Disabled"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPShortLine Disabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"CreateDate"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPShortLine CreateDate error:" + err.Error())
	}

}

func newGPShortLineFindByID(id string) (result EzOrmObj, err error) {
	return GPShortLineMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPShortLine []func(obj EzOrmObj)
	updateCB_GPShortLine []func(obj EzOrmObj)
)

func GPShortLineAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPShortLine = append(insertCB_GPShortLine, cb)
}

func GPShortLineAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPShortLine = append(updateCB_GPShortLine, cb)
}

func (o *GPShortLine) Id() string {
	return o.ID.Hex()
}

func (o *GPShortLine) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPShortLineInsertCallback(o)
	} else {
		GPShortLineUpdateCallback(o)
	}

	return
}

func (o *GPShortLine) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPShortLineMgr.GetCol()
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
		GPShortLineInsertCallback(o)
	}
	return
}

func GPShortLineInsertCallback(o *GPShortLine) {
	for _, cb := range insertCB_GPShortLine {
		cb(o)
	}
}

func GPShortLineUpdateCallback(o *GPShortLine) {
	for _, cb := range updateCB_GPShortLine {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPShortLineMgr) FindOne(query interface{}, sortFields ...string) (result *GPShortLine, err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPShortLineSort(q, sortFields)

	err = q.One(&result)
	return
}

// _GPShortLineSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _GPShortLineSort(q *mgo.Query, sortFields []string) {
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
func (o *_GPShortLineMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPShortLineMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPShortLineSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_GPShortLineMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPShortLineMgr.GetCol()
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
func (o *_GPShortLineMgr) FindOneBySecucodeCreateDate(Secucode string, CreateDate int64) (result *GPShortLine, err error) {
	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	session, q := GPShortLineMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPShortLineMgr) MustFindOneBySecucodeCreateDate(Secucode string, CreateDate int64) (result *GPShortLine) {
	result, _ = o.FindOneBySecucodeCreateDate(Secucode, CreateDate)
	if result == nil {
		result = GPShortLineMgr.NewGPShortLine()
		result.Secucode = Secucode
		result.CreateDate = CreateDate
		result.Save()
	}
	return
}

func (o *_GPShortLineMgr) RemoveBySecucodeCreateDate(Secucode string, CreateDate int64) (err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode":   Secucode,
		"CreateDate": CreateDate,
	}
	return col.Remove(query)
}
func (o *_GPShortLineMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPShortLine, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPShortLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPShortLineMgr) FindByDisabled(Disabled bool, limit int, offset int, sortFields ...string) (result []*GPShortLine, err error) {
	query := db.M{
		"Disabled": Disabled,
	}
	session, q := GPShortLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPShortLineMgr) FindByCreateDate(CreateDate int64, limit int, offset int, sortFields ...string) (result []*GPShortLine, err error) {
	query := db.M{
		"CreateDate": CreateDate,
	}
	session, q := GPShortLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPShortLineMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPShortLine, err error) {
	session, q := GPShortLineMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPShortLineMgr) FindAll(query interface{}, sortFields ...string) (result []*GPShortLine, err error) {
	session, q := GPShortLineMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPShortLineMgr) Has(query interface{}) bool {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPShortLineMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPShortLineMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPShortLineMgr) FindByIDs(id []string, sortFields ...string) (result []*GPShortLine, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPShortLineMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPShortLineMgr) FindByID(id string) (result *GPShortLine, err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPShortLineMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPShortLineMgr) RemoveByID(id string) (err error) {
	session, col := GPShortLineMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPShortLineMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPShortLine")
	}
	return getCol("digger", "digger.GPShortLine")
}

//Search

func (o *GPShortLine) IsSearchEnabled() bool {

	return false

}

//end search
