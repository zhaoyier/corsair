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

	db.SetOnEnsureIndex(initGPPromptBuyIndex)

	RegisterEzOrmObjByID("digger", "GPPromptBuy", newGPPromptBuyFindByID)
	RegisterEzOrmObjRemove("digger", "GPPromptBuy", GPPromptBuyMgr.RemoveByID)

}

func initGPPromptBuyIndex() {
	session, collection := GPPromptBuyMgr.GetCol()
	defer session.Close()

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode", "Disabled"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPPromptBuy SecucodeDisabled error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Secucode"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPPromptBuy Secucode error:" + err.Error())
	}

	if err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"Name"},
		Background: true,
		Sparse:     true,
	}); err != nil {
		panic("ensureIndex digger.GPPromptBuy Name error:" + err.Error())
	}

}

func newGPPromptBuyFindByID(id string) (result EzOrmObj, err error) {
	return GPPromptBuyMgr.FindByID(id)
}

//mongo methods
var (
	insertCB_GPPromptBuy []func(obj EzOrmObj)
	updateCB_GPPromptBuy []func(obj EzOrmObj)
)

func GPPromptBuyAddInsertCallback(cb func(obj EzOrmObj)) {
	insertCB_GPPromptBuy = append(insertCB_GPPromptBuy, cb)
}

func GPPromptBuyAddUpdateCallback(cb func(obj EzOrmObj)) {
	updateCB_GPPromptBuy = append(updateCB_GPPromptBuy, cb)
}

func (o *GPPromptBuy) Id() string {
	return o.ID.Hex()
}

func (o *GPPromptBuy) Save() (info *mgo.ChangeInfo, err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	isNew := o.isNew

	info, err = col.UpsertId(o.ID, o)
	o.isNew = false

	if isNew {
		GPPromptBuyInsertCallback(o)
	} else {
		GPPromptBuyUpdateCallback(o)
	}

	return
}

func (o *GPPromptBuy) InsertUnique(query interface{}) (saved bool, err error) {
	session, col := GPPromptBuyMgr.GetCol()
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
		GPPromptBuyInsertCallback(o)
	}
	return
}

func GPPromptBuyInsertCallback(o *GPPromptBuy) {
	for _, cb := range insertCB_GPPromptBuy {
		cb(o)
	}
}

func GPPromptBuyUpdateCallback(o *GPPromptBuy) {
	for _, cb := range updateCB_GPPromptBuy {
		cb(o)
	}
}

//foreigh keys

//Collection Manage methods

func (o *_GPPromptBuyMgr) FindOne(query interface{}, sortFields ...string) (result *GPPromptBuy, err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	q := col.Find(query)

	_GPPromptBuySort(q, sortFields)

	err = q.One(&result)
	return
}

func _GPPromptBuySort(q *mgo.Query, sortFields []string) {
	sortFields = XSortFieldsFilter(sortFields)
	if len(sortFields) > 0 {
		q.Sort(sortFields...)
		return
	}

	q.Sort("-_id")
}

func (o *_GPPromptBuyMgr) Query(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPPromptBuyMgr.GetCol()
	q := col.Find(query)
	if limit > 0 {
		q.Limit(limit)
	}
	if offset > 0 {
		q.Skip(offset)
	}

	_GPPromptBuySort(q, sortFields)
	return session, q
}

func (o *_GPPromptBuyMgr) NQuery(query interface{}, limit, offset int, sortFields []string) (*mgo.Session, *mgo.Query) {
	session, col := GPPromptBuyMgr.GetCol()
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
func (o *_GPPromptBuyMgr) FindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPPromptBuy, err error) {
	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	session, q := GPPromptBuyMgr.NQuery(query, 1, 0, nil)
	defer session.Close()
	err = q.One(&result)
	return
}

func (o *_GPPromptBuyMgr) MustFindOneBySecucodeDisabled(Secucode string, Disabled bool) (result *GPPromptBuy) {
	result, _ = o.FindOneBySecucodeDisabled(Secucode, Disabled)
	if result == nil {
		result = GPPromptBuyMgr.NewGPPromptBuy()
		result.Secucode = Secucode
		result.Disabled = Disabled
		result.Save()
	}
	return
}

func (o *_GPPromptBuyMgr) RemoveBySecucodeDisabled(Secucode string, Disabled bool) (err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	query := db.M{
		"Secucode": Secucode,
		"Disabled": Disabled,
	}
	return col.Remove(query)
}
func (o *_GPPromptBuyMgr) FindBySecucode(Secucode string, limit int, offset int, sortFields ...string) (result []*GPPromptBuy, err error) {
	query := db.M{
		"Secucode": Secucode,
	}
	session, q := GPPromptBuyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}
func (o *_GPPromptBuyMgr) FindByName(Name string, limit int, offset int, sortFields ...string) (result []*GPPromptBuy, err error) {
	query := db.M{
		"Name": Name,
	}
	session, q := GPPromptBuyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPPromptBuyMgr) Find(query interface{}, limit int, offset int, sortFields ...string) (result []*GPPromptBuy, err error) {
	session, q := GPPromptBuyMgr.Query(query, limit, offset, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPPromptBuyMgr) FindAll(query interface{}, sortFields ...string) (result []*GPPromptBuy, err error) {
	session, q := GPPromptBuyMgr.Query(query, -1, -1, sortFields)
	defer session.Close()
	err = q.All(&result)
	return
}

func (o *_GPPromptBuyMgr) Has(query interface{}) bool {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	var ret interface{}
	err := col.Find(query).One(&ret)
	if err != nil || ret == nil {
		return false
	}
	return true
}

func (o *_GPPromptBuyMgr) Count(query interface{}) (result int) {
	result, _ = o.CountE(query)
	return
}

func (o *_GPPromptBuyMgr) CountE(query interface{}) (result int, err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	result, err = col.Find(query).Count()
	return
}

func (o *_GPPromptBuyMgr) FindByIDs(id []string, sortFields ...string) (result []*GPPromptBuy, err error) {
	ids := make([]bson.ObjectId, 0, len(id))
	for _, i := range id {
		if bson.IsObjectIdHex(i) {
			ids = append(ids, bson.ObjectIdHex(i))
		}
	}
	return GPPromptBuyMgr.FindAll(db.M{"_id": db.M{"$in": ids}}, sortFields...)
}

func (m *_GPPromptBuyMgr) FindByID(id string) (result *GPPromptBuy, err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

func (m *_GPPromptBuyMgr) RemoveAll(query interface{}) (info *mgo.ChangeInfo, err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	return col.RemoveAll(query)
}

func (m *_GPPromptBuyMgr) RemoveByID(id string) (err error) {
	session, col := GPPromptBuyMgr.GetCol()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		err = mgo.ErrNotFound
		return
	}
	err = col.RemoveId(bson.ObjectIdHex(id))

	return
}

func (m *_GPPromptBuyMgr) GetCol() (session *mgo.Session, col *mgo.Collection) {
	if mgoInstances == nil {
		return db.GetCol("digger", "digger.GPPromptBuy")
	}
	return getCol("digger", "digger.GPPromptBuy")
}

//Search

func (o *GPPromptBuy) IsSearchEnabled() bool {

	return false

}

//end search
