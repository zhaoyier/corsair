package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ezbuy/redis-orm/orm"
	"gopkg.in/go-playground/validator.v9"
)

var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
	_ validator.Validate
	_ context.Context
)

type OrderItemRemark struct {
	OrderItemId uint64    `db:"order_item_id"`
	Id          uint64    `db:"id"`
	Source      int32     `db:"source"`
	Remark      string    `db:"remark"`
	CreateBy    string    `db:"create_by"`
	CreateDate  time.Time `db:"create_date"`
	UpdateDate  time.Time `db:"update_date"`
}

var OrderItemRemarkColumns = struct {
	OrderItemId string
	Id          string
	Source      string
	Remark      string
	CreateBy    string
	CreateDate  string
	UpdateDate  string
}{
	"order_item_id",
	"id",
	"source",
	"remark",
	"create_by",
	"create_date",
	"update_date",
}

type _OrderItemRemarkMgr struct {
}

var OrderItemRemarkMgr *_OrderItemRemarkMgr

func (m *_OrderItemRemarkMgr) NewOrderItemRemark() *OrderItemRemark {
	return &OrderItemRemark{}
}

//! object function

func (obj *OrderItemRemark) GetNameSpace() string {
	return "model"
}

func (obj *OrderItemRemark) GetClassName() string {
	return "OrderItemRemark"
}

func (obj *OrderItemRemark) GetTableName() string {
	return "order_item_remark"
}

func (obj *OrderItemRemark) GetColumns() []string {
	columns := []string{
		"order_item_remark.`order_item_id`",
		"order_item_remark.`id`",
		"order_item_remark.`source`",
		"order_item_remark.`remark`",
		"order_item_remark.`create_by`",
		"order_item_remark.`create_date`",
		"order_item_remark.`update_date`",
	}
	return columns
}

func (obj *OrderItemRemark) GetNoneIncrementColumns() []string {
	columns := []string{
		"`order_item_id`",
		"`id`",
		"`source`",
		"`remark`",
		"`create_by`",
		"`create_date`",
		"`update_date`",
	}
	return columns
}

func (obj *OrderItemRemark) GetPrimaryKey() PrimaryKey {
	pk := OrderItemRemarkMgr.NewPrimaryKey()
	pk.OrderItemId = obj.OrderItemId
	pk.Id = obj.Id
	return pk
}

func (obj *OrderItemRemark) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type OrderItemIdIdOfOrderItemRemarkPK struct {
	OrderItemId uint64
	Id          uint64
}

func (m *_OrderItemRemarkMgr) NewPrimaryKey() *OrderItemIdIdOfOrderItemRemarkPK {
	return &OrderItemIdIdOfOrderItemRemarkPK{}
}

func (u *OrderItemIdIdOfOrderItemRemarkPK) Key() string {
	strs := []string{
		"OrderItemId",
		fmt.Sprint(u.OrderItemId),
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *OrderItemIdIdOfOrderItemRemarkPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
	}
	vOrderItemId, ok := kv["OrderItemId"]
	if !ok {
		return fmt.Errorf("key (%s) without (OrderItemId) field", key)
	}
	if err := orm.StringScan(vOrderItemId, &(u.OrderItemId)); err != nil {
		return err
	}
	vId, ok := kv["Id"]
	if !ok {
		return fmt.Errorf("key (%s) without (Id) field", key)
	}
	if err := orm.StringScan(vId, &(u.Id)); err != nil {
		return err
	}
	return nil
}

func (u *OrderItemIdIdOfOrderItemRemarkPK) SQLFormat() string {
	conditions := []string{
		"`order_item_id` = ?",
		"`id` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *OrderItemIdIdOfOrderItemRemarkPK) SQLParams() []interface{} {
	return []interface{}{
		u.OrderItemId,
		u.Id,
	}
}

func (u *OrderItemIdIdOfOrderItemRemarkPK) Columns() []string {
	return []string{
		"`order_item_id`",
		"`id`",
	}
}

//! uniques

//! indexes

type OrderItemIdOfOrderItemRemarkIDX struct {
	OrderItemId uint64
	offset      int
	limit       int
}

func (u *OrderItemIdOfOrderItemRemarkIDX) Key() string {
	strs := []string{
		"OrderItemId",
		fmt.Sprint(u.OrderItemId),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *OrderItemIdOfOrderItemRemarkIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`order_item_id` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *OrderItemIdOfOrderItemRemarkIDX) SQLParams() []interface{} {
	return []interface{}{
		u.OrderItemId,
	}
}

func (u *OrderItemIdOfOrderItemRemarkIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *OrderItemIdOfOrderItemRemarkIDX) Limit(n int) {
	u.limit = n
}

func (u *OrderItemIdOfOrderItemRemarkIDX) Offset(n int) {
	u.offset = n
}

func (u *OrderItemIdOfOrderItemRemarkIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *OrderItemIdOfOrderItemRemarkIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type OrderItemIdIdOfOrderItemRemarkRNG struct {
	OrderItemId  uint64
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) Key() string {
	strs := []string{
		"OrderItemId",
		fmt.Sprint(u.OrderItemId),
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *OrderItemIdIdOfOrderItemRemarkRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	conditions = append(conditions, "`order_item_id` = ?")
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			conditions = append(conditions, fmt.Sprintf("`id` %s ?", u.beginOp()))
		}
		if u.IdEnd != -1 {
			conditions = append(conditions, fmt.Sprintf("`id` %s ?", u.endOp()))
		}
	}
	if limit {
		return fmt.Sprintf("%s %s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`id`", u.revert), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOrderBy("`id`", u.revert))
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) SQLParams() []interface{} {
	params := []interface{}{
		u.OrderItemId,
	}
	if u.IdBegin != u.IdEnd {
		if u.IdBegin != -1 {
			params = append(params, u.IdBegin)
		}
		if u.IdEnd != -1 {
			params = append(params, u.IdEnd)
		}
	}
	return params
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) Limit(n int) {
	u.limit = n
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) Offset(n int) {
	u.offset = n
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) Begin() int64 {
	start := u.IdBegin
	if start == -1 || start == 0 {
		start = 0
	}
	if start > 0 {
		if !u.includeBegin {
			start = start + 1
		}
	}
	return start
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) End() int64 {
	stop := u.IdEnd
	if stop == 0 || stop == -1 {
		stop = -1
	}
	if stop > 0 {
		if !u.includeBegin {
			stop = stop - 1
		}
	}
	return stop
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) Revert(b bool) {
	u.revert = b
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *OrderItemIdIdOfOrderItemRemarkRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

type _OrderItemRemarkDBMgr struct {
	db orm.DB
}

func (m *_OrderItemRemarkMgr) DB(db orm.DB) *_OrderItemRemarkDBMgr {
	return OrderItemRemarkDBMgr(db)
}

func OrderItemRemarkDBMgr(db orm.DB) *_OrderItemRemarkDBMgr {
	if db == nil {
		panic(fmt.Errorf("OrderItemRemarkDBMgr init need db"))
	}
	return &_OrderItemRemarkDBMgr{db: db}
}

func (m *_OrderItemRemarkDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	q := fmt.Sprintf("SELECT %s FROM order_item_remark %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	q := fmt.Sprintf("SELECT %s FROM order_item_remark %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_OrderItemRemarkDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_OrderItemRemarkDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_OrderItemRemarkDBMgr) FetchBySQL(q string, args ...interface{}) (results []*OrderItemRemark, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("OrderItemRemark fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result OrderItemRemark
		err = rows.Scan(&(result.OrderItemId), &(result.Id), &(result.Source), &(result.Remark), &(result.CreateBy), &CreateDate, &UpdateDate)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.CreateDate = time.Unix(CreateDate, 0)
		result.UpdateDate = time.Unix(UpdateDate, 0)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("OrderItemRemark fetch result error: %v", err)
	}
	return
}

func (m *_OrderItemRemarkDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*OrderItemRemark, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("OrderItemRemark fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result OrderItemRemark
		err = rows.Scan(&(result.OrderItemId), &(result.Id), &(result.Source), &(result.Remark), &(result.CreateBy), &CreateDate, &UpdateDate)
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		result.CreateDate = time.Unix(CreateDate, 0)
		result.UpdateDate = time.Unix(UpdateDate, 0)

		results = append(results, &result)
	}
	if err = rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("OrderItemRemark fetch result error: %v", err)
	}
	return
}
func (m *_OrderItemRemarkDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_OrderItemRemarkDBMgr) Fetch(pk PrimaryKey) (*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("OrderItemRemark fetch record not found")
}

// err not found check
func (m *_OrderItemRemarkDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_OrderItemRemarkDBMgr) FetchByPrimaryKey(orderItemId uint64, id uint64) (*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	pk := &OrderItemIdIdOfOrderItemRemarkPK{
		OrderItemId: orderItemId,
		Id:          id,
	}

	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("OrderItemRemark fetch record not found")
}

func (m *_OrderItemRemarkDBMgr) FetchByPrimaryKeyContext(ctx context.Context, orderItemId uint64, id uint64) (*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	pk := &OrderItemIdIdOfOrderItemRemarkPK{
		OrderItemId: orderItemId,
		Id:          id,
	}

	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("OrderItemRemark fetch record not found")
}

// indexes

func (m *_OrderItemRemarkDBMgr) FindByOrderItemId(orderItemId uint64, limit int, offset int) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	idx_ := &OrderItemIdOfOrderItemRemarkIDX{
		OrderItemId: orderItemId,
		limit:       limit,
		offset:      offset,
	}

	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_OrderItemRemarkDBMgr) FindByOrderItemIdContext(ctx context.Context, orderItemId uint64, limit int, offset int) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	idx_ := &OrderItemIdOfOrderItemRemarkIDX{
		OrderItemId: orderItemId,
		limit:       limit,
		offset:      offset,
	}
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_OrderItemRemarkDBMgr) FindAllByOrderItemId(orderItemId uint64) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	idx_ := &OrderItemIdOfOrderItemRemarkIDX{
		OrderItemId: orderItemId,
	}

	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_OrderItemRemarkDBMgr) FindAllByOrderItemIdContext(ctx context.Context, orderItemId uint64) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	idx_ := &OrderItemIdOfOrderItemRemarkIDX{
		OrderItemId: orderItemId,
	}

	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_OrderItemRemarkDBMgr) FindByOrderItemIdGroup(items []uint64) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM order_item_remark where `order_item_id` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_OrderItemRemarkDBMgr) FindByOrderItemIdGroupContext(ctx context.Context, items []uint64) ([]*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM order_item_remark where `order_item_id` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

func (m *_OrderItemRemarkDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("OrderItemRemark find record not found")
}

func (m *_OrderItemRemarkDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("OrderItemRemark find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_OrderItemRemarkDBMgr) FindOneFetch(unique Unique) (*OrderItemRemark, error) {
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
	objs, err := m.FetchBySQL(query, unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("none record")
}

// Deprecated: Use FindByXXXUnique instead.
func (m *_OrderItemRemarkDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_OrderItemRemarkDBMgr) FindFetch(index Index) (int64, []*OrderItemRemark, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OrderItemRemarkDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*OrderItemRemark, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OrderItemRemarkDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_OrderItemRemarkDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_OrderItemRemarkDBMgr) RangeFetch(scope Range) (int64, []*OrderItemRemark, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OrderItemRemarkDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*OrderItemRemark, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := OrderItemRemarkMgr.NewOrderItemRemark()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_OrderItemRemarkDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_OrderItemRemarkDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_OrderItemRemarkDBMgr) RangeRevertFetch(scope Range) (int64, []*OrderItemRemark, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_OrderItemRemarkDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*OrderItemRemark, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_OrderItemRemarkDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := OrderItemRemarkMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("OrderItemRemark query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := OrderItemRemarkMgr.NewPrimaryKey()
		err = rows.Scan(&(result.OrderItemId), &(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("OrderItemRemark query limit result error: %v", err)
	}
	return
}

func (m *_OrderItemRemarkDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := OrderItemRemarkMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM order_item_remark %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("OrderItemRemark query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := OrderItemRemarkMgr.NewPrimaryKey()
		err = rows.Scan(&(result.OrderItemId), &(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("OrderItemRemark query limit result error: %v", err)
	}
	return
}

func (m *_OrderItemRemarkDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`order_item_id`) FROM order_item_remark %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("OrderItemRemark query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			m.db.SetError(err)
			return 0, err
		}
		break
	}
	return count, nil
}

func (m *_OrderItemRemarkDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`order_item_id`) FROM order_item_remark %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("OrderItemRemark query count error: %v", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			m.db.SetError(err)
			return 0, err
		}
		break
	}
	return count, nil
}

func (m *_OrderItemRemarkDBMgr) BatchCreate(objs []*OrderItemRemark) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*7)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(7, "?"), ",")))
		values = append(values, obj.OrderItemId)
		values = append(values, obj.Id)
		values = append(values, obj.Source)
		values = append(values, obj.Remark)
		values = append(values, obj.CreateBy)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO order_item_remark(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) BatchCreateContext(ctx context.Context, objs []*OrderItemRemark) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*7)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(7, "?"), ",")))
		values = append(values, obj.OrderItemId)
		values = append(values, obj.Id)
		values = append(values, obj.Source)
		values = append(values, obj.Remark)
		values = append(values, obj.CreateBy)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO order_item_remark(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.ExecContext(ctx, query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (m *_OrderItemRemarkDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE order_item_remark SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE order_item_remark SET %s WHERE %s", set, where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// argument example:
// set:"a=?, b=?"
// where:"c=? and d=?"
// params:[]interface{}{"a", "b", "c", "d"}...
func (m *_OrderItemRemarkDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE order_item_remark SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE order_item_remark SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) Create(obj *OrderItemRemark) (int64, error) {
	params := orm.NewStringSlice(7, "?")
	q := fmt.Sprintf("INSERT INTO order_item_remark(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 7)
	values = append(values, obj.OrderItemId)
	values = append(values, obj.Id)
	values = append(values, obj.Source)
	values = append(values, obj.Remark)
	values = append(values, obj.CreateBy)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) CreateContext(ctx context.Context, obj *OrderItemRemark) (int64, error) {
	params := orm.NewStringSlice(7, "?")
	q := fmt.Sprintf("INSERT INTO order_item_remark(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 7)
	values = append(values, obj.OrderItemId)
	values = append(values, obj.Id)
	values = append(values, obj.Source)
	values = append(values, obj.Remark)
	values = append(values, obj.CreateBy)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) Update(obj *OrderItemRemark) (int64, error) {
	columns := []string{
		"`source` = ?",
		"`remark` = ?",
		"`create_by` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE order_item_remark SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 7-2)
	values = append(values, obj.Source)
	values = append(values, obj.Remark)
	values = append(values, obj.CreateBy)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) UpdateContext(ctx context.Context, obj *OrderItemRemark) (int64, error) {
	columns := []string{
		"`source` = ?",
		"`remark` = ?",
		"`create_by` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE order_item_remark SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 7-2)
	values = append(values, obj.Source)
	values = append(values, obj.Remark)
	values = append(values, obj.CreateBy)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) Save(obj *OrderItemRemark) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_OrderItemRemarkDBMgr) SaveContext(ctx context.Context, obj *OrderItemRemark) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_OrderItemRemarkDBMgr) Delete(obj *OrderItemRemark) (int64, error) {
	return m.DeleteByPrimaryKey(obj.OrderItemId, obj.Id)
}

func (m *_OrderItemRemarkDBMgr) DeleteContext(ctx context.Context, obj *OrderItemRemark) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.OrderItemId, obj.Id)
}

func (m *_OrderItemRemarkDBMgr) DeleteByPrimaryKey(orderItemId uint64, id uint64) (int64, error) {
	pk := &OrderItemIdIdOfOrderItemRemarkPK{
		OrderItemId: orderItemId,
		Id:          id,
	}
	q := fmt.Sprintf("DELETE FROM order_item_remark %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, orderItemId uint64, id uint64) (int64, error) {
	pk := &OrderItemIdIdOfOrderItemRemarkPK{
		OrderItemId: orderItemId,
		Id:          id,
	}
	q := fmt.Sprintf("DELETE FROM order_item_remark %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM order_item_remark")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM order_item_remark WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_OrderItemRemarkDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM order_item_remark")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM order_item_remark WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
