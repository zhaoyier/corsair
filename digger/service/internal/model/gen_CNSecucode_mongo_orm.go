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

	db.SetOnEnsureIndex(initCNSecucodeIndex)

	RegisterEzOrmObjByID("digger", "CNSecucode", newCNSecucodeFindByID)
	RegisterEzOrmObjRemove("digger", "CNSecucode", CNSecucodeMgr.RemoveByID)

}

func initCNSecucodeIndex() {
	session, collection := CNSecucodeMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.CNSecucode Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.CNSecucode Name error:" + err.Error())
	}

}

func newCNSecucodeFindByID(id string) (result EzOrmObj, err error) {
	return CNSecucodeMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_CNSecucode []func(obj EzOrmObj)
	updateCB_CNSecucode []func(obj EzOrmObj)
)

func CNSecucodeAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_CNSecucode = append(insertCB_CNSecucode, cb)
}

func CNSecucodeAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_CNSecucode = append(updateCB_CNSecucode, cb)
}

func (o *CNSecucode) Id() string {
	return o.ID.Hex()
}

func (o *CNSecucode) Save() (info *mgo.ChangeInfo, err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		CNSecucodeInsertCallback(o)
	} else {
		CNSecucodeUpdateCallback(o)
	}

	return
}

func (o *CNSecucode) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := CNSecucodeMgr.GetCol()
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
		CNSecucodeInsertCallback(o)
	}
	return
}

func CNSecucodeInsertCallback(o *CNSecucode) {
	for _, cb := range insertCB_CNSecucode {
		cb(o)
	}
}

func CNSecucodeUpdateCallback(o *CNSecucode) {
	for _, cb := range updateCB_CNSecucode {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_CNSecucodeMgr) FindOne(query interface{}, sortFields ...string) (result *CNSecucode, err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_CNSecucodeSort(q, sortFields)

	err = q.One(&result)
	return
}

// _CNSecucodeSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _CNSecucodeSort(q *mgo.Query, sortFields []string) {
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
func (o *_CNSecucodeMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := CNSecucodeMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_CNSecucodeSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_CNSecucodeMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := CNSecucodeMgr.GetCol()
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
func (o *_CNSecucodeMgr) FindOneBySecucode(Secucode string) (result *CNSecucode, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := CNSecucodeMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_CNSecucodeMgr) MustFindOneBySecucode(Secucode string) (result *CNSecucode) {
	result, _ = o.FindOneBySecucode(Secucode)
	if result == nil {
		result = CNSecucodeMgr.NewCNSecucode()
		result.Secucode = Secucode
		result.Save()
	}
	return
}

func (o *_CNSecucodeMgr) RemoveBySecucode(Secucode string) (err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
	}
	return col.Remove(query)
}
func (o *_CNSecucodeMgr) FindOneByName(Name string) (result *CNSecucode, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := CNSecucodeMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_CNSecucodeMgr) MustFindOneByName(Name string) (result *CNSecucode) {
	result, _ = o.FindOneByName(Name)
	if result == nil {
		result = CNSecucodeMgr.NewCNSecucode()
		result.Name = Name
		result.Save()
	}
	return
}

func (o *_CNSecucodeMgr) RemoveByName(Name string) (err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Name": Name,
	}
	return col.Remove(query)
}

func (o *_CNSecucodeMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*CNSecucode, err error) {
	session, q := CNSecucodeMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_CNSecucodeMgr) FindAll(query interface{}, sortFields ...string) (result []*CNSecucode, err error) {
	session, q := CNSecucodeMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_CNSecucodeMgr) Has(query interface{}) bool {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_CNSecucodeMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_CNSecucodeMgr) CountE(query interface{}) (result int, err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_CNSecucodeMgr) FindByIDs(id []string, sortFields ...string) (result []*CNSecucode, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return CNSecucodeMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_CNSecucodeMgr) FindByID(id string) (result *CNSecucode, err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_CNSecucodeMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_CNSecucodeMgr) RemoveByID(id string) (err error) {
	session, col := CNSecucodeMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_CNSecucodeMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.CNSecucode")
	}
	return getCol("digger", "digger.CNSecucode")
}

//Search

func (o *CNSecucode) IsSearchEnabled() bool {

	return false

}

//end search
