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
		Key:        []string{"Secucode", "EndDate"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GDRenshu SecucodeEndDate error:" + err.Error())
	}

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

// _GDRenshuSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _GDRenshuSort(q *mgo.Query, sortFields []string) {
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

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
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
func (o *_GDRenshuMgr) FindOneBySecucodeEndDate(Secucode string, EndDate int64) (result *GDRenshu, err error) {
	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	session, q := GDRenshuMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GDRenshuMgr) MustFindOneBySecucodeEndDate(Secucode string, EndDate int64) (result *GDRenshu) {
	result, _ = o.FindOneBySecucodeEndDate(Secucode, EndDate)
	if result == nil {
		result = GDRenshuMgr.NewGDRenshu()
		result.Secucode = Secucode
		result.EndDate = EndDate
		result.Save()
	}
	return
}

func (o *_GDRenshuMgr) RemoveBySecucodeEndDate(Secucode string, EndDate int64) (err error) {
	session, col := GDRenshuMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"EndDate":  EndDate,
	}
	return col.Remove(query)
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
func (o *_GDRenshuMgr) FindBySecurityCode(SecurityCode string, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
	query := db.M{
		"SecurityCode": SecurityCode,
	}
	session, q := GDRenshuMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GDRenshuMgr) FindByEndDate(EndDate int64, limit int, offset int, sortFields ...string) (result []*GDRenshu, err error) {
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
