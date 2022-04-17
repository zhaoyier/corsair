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

	db.SetOnEnsureIndex(initGDAggregationIndex)

	RegisterEzOrmObjByID("digger", "GDAggregation", newGDAggregationFindByID)
	RegisterEzOrmObjRemove("digger", "GDAggregation", GDAggregationMgr.RemoveByID)

}

func initGDAggregationIndex() {
	session, collection := GDAggregationMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"TotalRatioAccum"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation TotalRatioAccum error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"PriceRatio"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation PriceRatio error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"HoldRatioTotal"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation HoldRatioTotal error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"FreeholdRatioTotal"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDAggregation FreeholdRatioTotal error:" + err.Error())
	}

}

func newGDAggregationFindByID(id string) (result EzOrmObj, err error) {
	return GDAggregationMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GDAggregation []func(obj EzOrmObj)
	updateCB_GDAggregation []func(obj EzOrmObj)
)

func GDAggregationAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GDAggregation = append(insertCB_GDAggregation, cb)
}

func GDAggregationAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GDAggregation = append(updateCB_GDAggregation, cb)
}

func (o *GDAggregation) Id() string {
	return o.ID.Hex()
}

func (o *GDAggregation) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GDAggregationInsertCallback(o)
	} else {
		GDAggregationUpdateCallback(o)
	}

	return
}

func (o *GDAggregation) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GDAggregationMgr.GetCol()
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
		GDAggregationInsertCallback(o)
	}
	return
}

func GDAggregationInsertCallback(o *GDAggregation) {
	for _, cb := range insertCB_GDAggregation {
		cb(o)
	}
}

func GDAggregationUpdateCallback(o *GDAggregation) {
	for _, cb := range updateCB_GDAggregation {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GDAggregationMgr) FindOne(query interface{}, sortFields ...string) (result *GDAggregation, err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GDAggregationSort(q, sortFields)

	err = q.One(&result)
	return
}

// _GDAggregationSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _GDAggregationSort(q *mgo.Query, sortFields []string) {
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
func (o *_GDAggregationMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDAggregationMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GDAggregationSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_GDAggregationMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GDAggregationMgr.GetCol()
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
func (o *_GDAggregationMgr) FindOneBySecucode(Secucode string) (result *GDAggregation, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GDAggregationMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDAggregationMgr) MustFindOneBySecucode(Secucode string) (result *GDAggregation) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = GDAggregationMgr.NewGDAggregation()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_GDAggregationMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}
func (o *_GDAggregationMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDAggregationMgr) FindByTotalRatioAccum(TotalRatioAccum int32, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	query := db.M{
		"TotalRatioAccum": TotalRatioAccum,
	}
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDAggregationMgr) FindByPriceRatio(PriceRatio int32, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	query := db.M{
		"PriceRatio": PriceRatio,
	}
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDAggregationMgr) FindByHoldRatioTotal(HoldRatioTotal int32, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	query := db.M{
		"HoldRatioTotal": HoldRatioTotal,
	}
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDAggregationMgr) FindByFreeholdRatioTotal(FreeholdRatioTotal int32, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	query := db.M{
		"FreeholdRatioTotal": FreeholdRatioTotal,
	}
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDAggregationMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GDAggregation, err error) {
	session, q := GDAggregationMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDAggregationMgr) FindAll(query interface{}, sortFields ...string) (result []*GDAggregation, err error) {
	session, q := GDAggregationMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GDAggregationMgr) Has(query interface{}) bool {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GDAggregationMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GDAggregationMgr) CountE(query interface{}) (result int, err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GDAggregationMgr) FindByIDs(id []string, sortFields ...string) (result []*GDAggregation, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GDAggregationMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GDAggregationMgr) FindByID(id string) (result *GDAggregation, err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GDAggregationMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GDAggregationMgr) RemoveByID(id string) (err error) {
	session, col := GDAggregationMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GDAggregationMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GDAggregation")
	}
	return getCol("digger", "digger.GDAggregation")
}

//Search

func (o *GDAggregation) IsSearchEnabled() bool {

	return false

}

//end search
