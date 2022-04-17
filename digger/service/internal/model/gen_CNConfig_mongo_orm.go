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

	RegisterEzOrmObjByID("digger", "CNConfig", newCNConfigFindByID)
	RegisterEzOrmObjRemove("digger", "CNConfig", CNConfigMgr.RemoveByID)

}

func newCNConfigFindByID(id string) (result EzOrmObj, err error) {
	return CNConfigMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_CNConfig []func(obj EzOrmObj)
	updateCB_CNConfig []func(obj EzOrmObj)
)

func CNConfigAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_CNConfig = append(insertCB_CNConfig, cb)
}

func CNConfigAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_CNConfig = append(updateCB_CNConfig, cb)
}

func (o *CNConfig) Id() string {
	return o.ID.Hex()
}

func (o *CNConfig) Save() (info *mgo.ChangeInfo, err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		CNConfigInsertCallback(o)
	} else {
		CNConfigUpdateCallback(o)
	}

	return
}

func (o *CNConfig) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := CNConfigMgr.GetCol()
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
		CNConfigInsertCallback(o)
	}
	return
}

func CNConfigInsertCallback(o *CNConfig) {
	for _, cb := range insertCB_CNConfig {
		cb(o)
	}
}

func CNConfigUpdateCallback(o *CNConfig) {
	for _, cb := range updateCB_CNConfig {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_CNConfigMgr) FindOne(query interface{}, sortFields ...string) (result *CNConfig, err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_CNConfigSort(q, sortFields)

	err = q.One(&result)
	return
}

// _CNConfigSort 将排序字段应用到查询对象中，如果找不到有效的排序字段，则默认使用 `-_id` 作为排序字段
func _CNConfigSort(q *mgo.Query, sortFields []string) {
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
func (o *_CNConfigMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := CNConfigMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_CNConfigSort(q, sortFields)
	return session, q
}

// NQuery 按照查询条件、分页、排序等构建 MongoDB 查询对象，如果不指定排序字段，则 MongoDB
// 会按照引擎中的存储顺序返回（Natural-Order）， 不保证返回数据保持插入顺序或插入倒序。
// 建议仅在保证返回数据唯一的情况下使用
// Ref: https://docs.mongodb.com/manual/reference/method/cursor.sort/#return-in-natural-order
//   - 如果 limit 小于等于 0，则忽略该参数
//   - 如果 offset 小于等于 0，则忽略该参数
//   - 如果 sortFields 为空或全为非法值，则忽略该参数
func (o *_CNConfigMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := CNConfigMgr.GetCol()
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

func (o *_CNConfigMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*CNConfig, err error) {
	session, q := CNConfigMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_CNConfigMgr) FindAll(query interface{}, sortFields ...string) (result []*CNConfig, err error) {
	session, q := CNConfigMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_CNConfigMgr) Has(query interface{}) bool {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_CNConfigMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_CNConfigMgr) CountE(query interface{}) (result int, err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_CNConfigMgr) FindByIDs(id []string, sortFields ...string) (result []*CNConfig, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return CNConfigMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_CNConfigMgr) FindByID(id string) (result *CNConfig, err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_CNConfigMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_CNConfigMgr) RemoveByID(id string) (err error) {
	session, col := CNConfigMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_CNConfigMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.CNConfig")
	}
	return getCol("digger", "digger.CNConfig")
}

//Search

func (o *CNConfig) IsSearchEnabled() bool {

	return false

}

//end search
