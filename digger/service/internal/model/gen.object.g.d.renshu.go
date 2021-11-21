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

type GDRenshu struct {
	Id                 int64     `db:"id"`
	Secucode           string    `db:"secucode"`
	SecurityCode       int32     `db:"security_code"`
	EndDate            string    `db:"end_date"`
	HolderTotalNum     float64   `db:"holder_total_num"`
	TotalNumRatio      float64   `db:"total_num_ratio"`
	AvgFreeShares      float64   `db:"avg_free_shares"`
	AvgFreesharesRatio float64   `db:"avg_freeshares_ratio"`
	HoldFocus          string    `db:"hold_focus"`
	Price              float64   `db:"price"`
	AvgHoldAmt         float64   `db:"avg_hold_amt"`
	HoldRatioTotal     float64   `db:"hold_ratio_total"`
	FreeholdRatioTotal float64   `db:"freehold_ratio_total"`
	CreateDate         time.Time `db:"create_date"`
	UpdateDate         time.Time `db:"update_date"`
}

var GDRenshuColumns = struct {
	Id                 string
	Secucode           string
	SecurityCode       string
	EndDate            string
	HolderTotalNum     string
	TotalNumRatio      string
	AvgFreeShares      string
	AvgFreesharesRatio string
	HoldFocus          string
	Price              string
	AvgHoldAmt         string
	HoldRatioTotal     string
	FreeholdRatioTotal string
	CreateDate         string
	UpdateDate         string
}{
	"id",
	"secucode",
	"security_code",
	"end_date",
	"holder_total_num",
	"total_num_ratio",
	"avg_free_shares",
	"avg_freeshares_ratio",
	"hold_focus",
	"price",
	"avg_hold_amt",
	"hold_ratio_total",
	"freehold_ratio_total",
	"create_date",
	"update_date",
}

type _GDRenshuMgr struct {
}

var GDRenshuMgr *_GDRenshuMgr

func (m *_GDRenshuMgr) NewGDRenshu() *GDRenshu {
	return &GDRenshu{}
}

//! object function

func (obj *GDRenshu) GetNameSpace() string {
	return "model"
}

func (obj *GDRenshu) GetClassName() string {
	return "GDRenshu"
}

func (obj *GDRenshu) GetTableName() string {
	return "gd_renshu"
}

func (obj *GDRenshu) GetColumns() []string {
	columns := []string{
		"gd_renshu.`id`",
		"gd_renshu.`secucode`",
		"gd_renshu.`security_code`",
		"gd_renshu.`end_date`",
		"gd_renshu.`holder_total_num`",
		"gd_renshu.`total_num_ratio`",
		"gd_renshu.`avg_free_shares`",
		"gd_renshu.`avg_freeshares_ratio`",
		"gd_renshu.`hold_focus`",
		"gd_renshu.`price`",
		"gd_renshu.`avg_hold_amt`",
		"gd_renshu.`hold_ratio_total`",
		"gd_renshu.`freehold_ratio_total`",
		"gd_renshu.`create_date`",
		"gd_renshu.`update_date`",
	}
	return columns
}

func (obj *GDRenshu) GetNoneIncrementColumns() []string {
	columns := []string{
		"`id`",
		"`secucode`",
		"`security_code`",
		"`end_date`",
		"`holder_total_num`",
		"`total_num_ratio`",
		"`avg_free_shares`",
		"`avg_freeshares_ratio`",
		"`hold_focus`",
		"`price`",
		"`avg_hold_amt`",
		"`hold_ratio_total`",
		"`freehold_ratio_total`",
		"`create_date`",
		"`update_date`",
	}
	return columns
}

func (obj *GDRenshu) GetPrimaryKey() PrimaryKey {
	pk := GDRenshuMgr.NewPrimaryKey()
	pk.Id = obj.Id
	return pk
}

func (obj *GDRenshu) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type IdOfGDRenshuPK struct {
	Id int64
}

func (m *_GDRenshuMgr) NewPrimaryKey() *IdOfGDRenshuPK {
	return &IdOfGDRenshuPK{}
}

func (u *IdOfGDRenshuPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfGDRenshuPK) Parse(key string) error {
	arr := strings.Split(key, ":")
	if len(arr)%2 != 0 {
		return fmt.Errorf("key (%s) format error", key)
	}
	kv := map[string]string{}
	for i := 0; i < len(arr)/2; i++ {
		kv[arr[2*i]] = arr[2*i+1]
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

func (u *IdOfGDRenshuPK) SQLFormat() string {
	conditions := []string{
		"`id` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdOfGDRenshuPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
	}
}

func (u *IdOfGDRenshuPK) Columns() []string {
	return []string{
		"`id`",
	}
}

//! uniques

//! indexes

type SecurityCodeOfGDRenshuIDX struct {
	SecurityCode int32
	offset       int
	limit        int
}

func (u *SecurityCodeOfGDRenshuIDX) Key() string {
	strs := []string{
		"SecurityCode",
		fmt.Sprint(u.SecurityCode),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *SecurityCodeOfGDRenshuIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`security_code` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *SecurityCodeOfGDRenshuIDX) SQLParams() []interface{} {
	return []interface{}{
		u.SecurityCode,
	}
}

func (u *SecurityCodeOfGDRenshuIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SecurityCodeOfGDRenshuIDX) Limit(n int) {
	u.limit = n
}

func (u *SecurityCodeOfGDRenshuIDX) Offset(n int) {
	u.offset = n
}

func (u *SecurityCodeOfGDRenshuIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *SecurityCodeOfGDRenshuIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

type SecucodeOfGDRenshuIDX struct {
	Secucode string
	offset   int
	limit    int
}

func (u *SecucodeOfGDRenshuIDX) Key() string {
	strs := []string{
		"Secucode",
		fmt.Sprint(u.Secucode),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *SecucodeOfGDRenshuIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`secucode` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *SecucodeOfGDRenshuIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Secucode,
	}
}

func (u *SecucodeOfGDRenshuIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SecucodeOfGDRenshuIDX) Limit(n int) {
	u.limit = n
}

func (u *SecucodeOfGDRenshuIDX) Offset(n int) {
	u.offset = n
}

func (u *SecucodeOfGDRenshuIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *SecucodeOfGDRenshuIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

type HolderTotalNumOfGDRenshuIDX struct {
	HolderTotalNum float64
	offset         int
	limit          int
}

func (u *HolderTotalNumOfGDRenshuIDX) Key() string {
	strs := []string{
		"HolderTotalNum",
		fmt.Sprint(u.HolderTotalNum),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *HolderTotalNumOfGDRenshuIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`holder_total_num` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *HolderTotalNumOfGDRenshuIDX) SQLParams() []interface{} {
	return []interface{}{
		u.HolderTotalNum,
	}
}

func (u *HolderTotalNumOfGDRenshuIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *HolderTotalNumOfGDRenshuIDX) Limit(n int) {
	u.limit = n
}

func (u *HolderTotalNumOfGDRenshuIDX) Offset(n int) {
	u.offset = n
}

func (u *HolderTotalNumOfGDRenshuIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *HolderTotalNumOfGDRenshuIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

type EndDateOfGDRenshuIDX struct {
	EndDate string
	offset  int
	limit   int
}

func (u *EndDateOfGDRenshuIDX) Key() string {
	strs := []string{
		"EndDate",
		fmt.Sprint(u.EndDate),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *EndDateOfGDRenshuIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`end_date` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *EndDateOfGDRenshuIDX) SQLParams() []interface{} {
	return []interface{}{
		u.EndDate,
	}
}

func (u *EndDateOfGDRenshuIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *EndDateOfGDRenshuIDX) Limit(n int) {
	u.limit = n
}

func (u *EndDateOfGDRenshuIDX) Offset(n int) {
	u.offset = n
}

func (u *EndDateOfGDRenshuIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *EndDateOfGDRenshuIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type IdOfGDRenshuRNG struct {
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdOfGDRenshuRNG) Key() string {
	strs := []string{
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdOfGDRenshuRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdOfGDRenshuRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdOfGDRenshuRNG) SQLFormat(limit bool) string {
	conditions := []string{}
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

func (u *IdOfGDRenshuRNG) SQLParams() []interface{} {
	params := []interface{}{}
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

func (u *IdOfGDRenshuRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdOfGDRenshuRNG) Limit(n int) {
	u.limit = n
}

func (u *IdOfGDRenshuRNG) Offset(n int) {
	u.offset = n
}

func (u *IdOfGDRenshuRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdOfGDRenshuRNG) Begin() int64 {
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

func (u *IdOfGDRenshuRNG) End() int64 {
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

func (u *IdOfGDRenshuRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdOfGDRenshuRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdOfGDRenshuRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdOfGDRenshuRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

type _GDRenshuDBMgr struct {
	db orm.DB
}

func (m *_GDRenshuMgr) DB(db orm.DB) *_GDRenshuDBMgr {
	return GDRenshuDBMgr(db)
}

func GDRenshuDBMgr(db orm.DB) *_GDRenshuDBMgr {
	if db == nil {
		panic(fmt.Errorf("GDRenshuDBMgr init need db"))
	}
	return &_GDRenshuDBMgr{db: db}
}

func (m *_GDRenshuDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_GDRenshuDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_GDRenshuDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	q := fmt.Sprintf("SELECT %s FROM gd_renshu %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_GDRenshuDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	q := fmt.Sprintf("SELECT %s FROM gd_renshu %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_GDRenshuDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_GDRenshuDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_GDRenshuDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_GDRenshuDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_GDRenshuDBMgr) FetchBySQL(q string, args ...interface{}) (results []*GDRenshu, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("GDRenshu fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result GDRenshu
		err = rows.Scan(&(result.Id), &(result.Secucode), &(result.SecurityCode), &(result.EndDate), &(result.HolderTotalNum), &(result.TotalNumRatio), &(result.AvgFreeShares), &(result.AvgFreesharesRatio), &(result.HoldFocus), &(result.Price), &(result.AvgHoldAmt), &(result.HoldRatioTotal), &(result.FreeholdRatioTotal), &CreateDate, &UpdateDate)
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
		return nil, fmt.Errorf("GDRenshu fetch result error: %v", err)
	}
	return
}

func (m *_GDRenshuDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*GDRenshu, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("GDRenshu fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result GDRenshu
		err = rows.Scan(&(result.Id), &(result.Secucode), &(result.SecurityCode), &(result.EndDate), &(result.HolderTotalNum), &(result.TotalNumRatio), &(result.AvgFreeShares), &(result.AvgFreesharesRatio), &(result.HoldFocus), &(result.Price), &(result.AvgHoldAmt), &(result.HoldRatioTotal), &(result.FreeholdRatioTotal), &CreateDate, &UpdateDate)
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
		return nil, fmt.Errorf("GDRenshu fetch result error: %v", err)
	}
	return
}
func (m *_GDRenshuDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_GDRenshuDBMgr) Fetch(pk PrimaryKey) (*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDRenshu fetch record not found")
}

// err not found check
func (m *_GDRenshuDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_GDRenshuDBMgr) FetchByPrimaryKey(id int64) (*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	pk := &IdOfGDRenshuPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDRenshu fetch record not found")
}

func (m *_GDRenshuDBMgr) FetchByPrimaryKeyContext(ctx context.Context, id int64) (*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	pk := &IdOfGDRenshuPK{
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDRenshu fetch record not found")
}

func (m *_GDRenshuDBMgr) FetchByPrimaryKeys(ids []int64) ([]*GDRenshu, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQL(query, params...)
}

func (m *_GDRenshuDBMgr) FetchByPrimaryKeysContext(ctx context.Context, ids []int64) ([]*GDRenshu, error) {
	size := len(ids)
	if size == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, size)
	for _, pk := range ids {
		params = append(params, pk)
	}
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu WHERE `id` IN (?%s)", strings.Join(obj.GetColumns(), ","),
		strings.Repeat(",?", size-1))
	return m.FetchBySQLContext(ctx, query, params...)
}

// indexes

func (m *_GDRenshuDBMgr) FindBySecurityCode(securityCode int32, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecurityCodeOfGDRenshuIDX{
		SecurityCode: securityCode,
		limit:        limit,
		offset:       offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindBySecurityCodeContext(ctx context.Context, securityCode int32, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecurityCodeOfGDRenshuIDX{
		SecurityCode: securityCode,
		limit:        limit,
		offset:       offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllBySecurityCode(securityCode int32) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecurityCodeOfGDRenshuIDX{
		SecurityCode: securityCode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllBySecurityCodeContext(ctx context.Context, securityCode int32) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecurityCodeOfGDRenshuIDX{
		SecurityCode: securityCode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindBySecurityCodeGroup(items []int32) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `security_code` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDRenshuDBMgr) FindBySecurityCodeGroupContext(ctx context.Context, items []int32) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `security_code` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

func (m *_GDRenshuDBMgr) FindBySecucode(secucode string, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecucodeOfGDRenshuIDX{
		Secucode: secucode,
		limit:    limit,
		offset:   offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindBySecucodeContext(ctx context.Context, secucode string, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecucodeOfGDRenshuIDX{
		Secucode: secucode,
		limit:    limit,
		offset:   offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllBySecucode(secucode string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecucodeOfGDRenshuIDX{
		Secucode: secucode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllBySecucodeContext(ctx context.Context, secucode string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &SecucodeOfGDRenshuIDX{
		Secucode: secucode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindBySecucodeGroup(items []string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `secucode` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDRenshuDBMgr) FindBySecucodeGroupContext(ctx context.Context, items []string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `secucode` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

func (m *_GDRenshuDBMgr) FindByHolderTotalNum(holderTotalNum float64, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &HolderTotalNumOfGDRenshuIDX{
		HolderTotalNum: holderTotalNum,
		limit:          limit,
		offset:         offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindByHolderTotalNumContext(ctx context.Context, holderTotalNum float64, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &HolderTotalNumOfGDRenshuIDX{
		HolderTotalNum: holderTotalNum,
		limit:          limit,
		offset:         offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllByHolderTotalNum(holderTotalNum float64) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &HolderTotalNumOfGDRenshuIDX{
		HolderTotalNum: holderTotalNum,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllByHolderTotalNumContext(ctx context.Context, holderTotalNum float64) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &HolderTotalNumOfGDRenshuIDX{
		HolderTotalNum: holderTotalNum,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindByHolderTotalNumGroup(items []float64) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `holder_total_num` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDRenshuDBMgr) FindByHolderTotalNumGroupContext(ctx context.Context, items []float64) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `holder_total_num` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

func (m *_GDRenshuDBMgr) FindByEndDate(endDate string, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &EndDateOfGDRenshuIDX{
		EndDate: endDate,
		limit:   limit,
		offset:  offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindByEndDateContext(ctx context.Context, endDate string, limit int, offset int) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &EndDateOfGDRenshuIDX{
		EndDate: endDate,
		limit:   limit,
		offset:  offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllByEndDate(endDate string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &EndDateOfGDRenshuIDX{
		EndDate: endDate,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindAllByEndDateContext(ctx context.Context, endDate string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	idx_ := &EndDateOfGDRenshuIDX{
		EndDate: endDate,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDRenshuDBMgr) FindByEndDateGroup(items []string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `end_date` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDRenshuDBMgr) FindByEndDateGroupContext(ctx context.Context, items []string) ([]*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_renshu where `end_date` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

func (m *_GDRenshuDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDRenshu find record not found")
}

func (m *_GDRenshuDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDRenshu find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_GDRenshuDBMgr) FindOneFetch(unique Unique) (*GDRenshu, error) {
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_GDRenshuDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_GDRenshuDBMgr) FindFetch(index Index) (int64, []*GDRenshu, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDRenshuDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*GDRenshu, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDRenshuDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_GDRenshuDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_GDRenshuDBMgr) RangeFetch(scope Range) (int64, []*GDRenshu, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDRenshuDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*GDRenshu, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := GDRenshuMgr.NewGDRenshu()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDRenshuDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_GDRenshuDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_GDRenshuDBMgr) RangeRevertFetch(scope Range) (int64, []*GDRenshu, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_GDRenshuDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*GDRenshu, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_GDRenshuDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := GDRenshuMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("GDRenshu query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := GDRenshuMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("GDRenshu query limit result error: %v", err)
	}
	return
}

func (m *_GDRenshuDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := GDRenshuMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM gd_renshu %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("GDRenshu query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := GDRenshuMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("GDRenshu query limit result error: %v", err)
	}
	return
}

func (m *_GDRenshuDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM gd_renshu %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("GDRenshu query count error: %v", err)
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

func (m *_GDRenshuDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM gd_renshu %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("GDRenshu query count error: %v", err)
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

func (m *_GDRenshuDBMgr) BatchCreate(objs []*GDRenshu) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*15)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(15, "?"), ",")))
		values = append(values, obj.Id)
		values = append(values, obj.Secucode)
		values = append(values, obj.SecurityCode)
		values = append(values, obj.EndDate)
		values = append(values, obj.HolderTotalNum)
		values = append(values, obj.TotalNumRatio)
		values = append(values, obj.AvgFreeShares)
		values = append(values, obj.AvgFreesharesRatio)
		values = append(values, obj.HoldFocus)
		values = append(values, obj.Price)
		values = append(values, obj.AvgHoldAmt)
		values = append(values, obj.HoldRatioTotal)
		values = append(values, obj.FreeholdRatioTotal)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO gd_renshu(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) BatchCreateContext(ctx context.Context, objs []*GDRenshu) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*15)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(15, "?"), ",")))
		values = append(values, obj.Id)
		values = append(values, obj.Secucode)
		values = append(values, obj.SecurityCode)
		values = append(values, obj.EndDate)
		values = append(values, obj.HolderTotalNum)
		values = append(values, obj.TotalNumRatio)
		values = append(values, obj.AvgFreeShares)
		values = append(values, obj.AvgFreesharesRatio)
		values = append(values, obj.HoldFocus)
		values = append(values, obj.Price)
		values = append(values, obj.AvgHoldAmt)
		values = append(values, obj.HoldRatioTotal)
		values = append(values, obj.FreeholdRatioTotal)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO gd_renshu(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_GDRenshuDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE gd_renshu SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE gd_renshu SET %s WHERE %s", set, where)
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
func (m *_GDRenshuDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE gd_renshu SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE gd_renshu SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) Create(obj *GDRenshu) (int64, error) {
	params := orm.NewStringSlice(15, "?")
	q := fmt.Sprintf("INSERT INTO gd_renshu(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 15)
	values = append(values, obj.Id)
	values = append(values, obj.Secucode)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.EndDate)
	values = append(values, obj.HolderTotalNum)
	values = append(values, obj.TotalNumRatio)
	values = append(values, obj.AvgFreeShares)
	values = append(values, obj.AvgFreesharesRatio)
	values = append(values, obj.HoldFocus)
	values = append(values, obj.Price)
	values = append(values, obj.AvgHoldAmt)
	values = append(values, obj.HoldRatioTotal)
	values = append(values, obj.FreeholdRatioTotal)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) CreateContext(ctx context.Context, obj *GDRenshu) (int64, error) {
	params := orm.NewStringSlice(15, "?")
	q := fmt.Sprintf("INSERT INTO gd_renshu(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 15)
	values = append(values, obj.Id)
	values = append(values, obj.Secucode)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.EndDate)
	values = append(values, obj.HolderTotalNum)
	values = append(values, obj.TotalNumRatio)
	values = append(values, obj.AvgFreeShares)
	values = append(values, obj.AvgFreesharesRatio)
	values = append(values, obj.HoldFocus)
	values = append(values, obj.Price)
	values = append(values, obj.AvgHoldAmt)
	values = append(values, obj.HoldRatioTotal)
	values = append(values, obj.FreeholdRatioTotal)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) Update(obj *GDRenshu) (int64, error) {
	columns := []string{
		"`secucode` = ?",
		"`security_code` = ?",
		"`end_date` = ?",
		"`holder_total_num` = ?",
		"`total_num_ratio` = ?",
		"`avg_free_shares` = ?",
		"`avg_freeshares_ratio` = ?",
		"`hold_focus` = ?",
		"`price` = ?",
		"`avg_hold_amt` = ?",
		"`hold_ratio_total` = ?",
		"`freehold_ratio_total` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE gd_renshu SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 15-1)
	values = append(values, obj.Secucode)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.EndDate)
	values = append(values, obj.HolderTotalNum)
	values = append(values, obj.TotalNumRatio)
	values = append(values, obj.AvgFreeShares)
	values = append(values, obj.AvgFreesharesRatio)
	values = append(values, obj.HoldFocus)
	values = append(values, obj.Price)
	values = append(values, obj.AvgHoldAmt)
	values = append(values, obj.HoldRatioTotal)
	values = append(values, obj.FreeholdRatioTotal)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) UpdateContext(ctx context.Context, obj *GDRenshu) (int64, error) {
	columns := []string{
		"`secucode` = ?",
		"`security_code` = ?",
		"`end_date` = ?",
		"`holder_total_num` = ?",
		"`total_num_ratio` = ?",
		"`avg_free_shares` = ?",
		"`avg_freeshares_ratio` = ?",
		"`hold_focus` = ?",
		"`price` = ?",
		"`avg_hold_amt` = ?",
		"`hold_ratio_total` = ?",
		"`freehold_ratio_total` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE gd_renshu SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 15-1)
	values = append(values, obj.Secucode)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.EndDate)
	values = append(values, obj.HolderTotalNum)
	values = append(values, obj.TotalNumRatio)
	values = append(values, obj.AvgFreeShares)
	values = append(values, obj.AvgFreesharesRatio)
	values = append(values, obj.HoldFocus)
	values = append(values, obj.Price)
	values = append(values, obj.AvgHoldAmt)
	values = append(values, obj.HoldRatioTotal)
	values = append(values, obj.FreeholdRatioTotal)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) Save(obj *GDRenshu) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_GDRenshuDBMgr) SaveContext(ctx context.Context, obj *GDRenshu) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_GDRenshuDBMgr) Delete(obj *GDRenshu) (int64, error) {
	return m.DeleteByPrimaryKey(obj.Id)
}

func (m *_GDRenshuDBMgr) DeleteContext(ctx context.Context, obj *GDRenshu) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.Id)
}

func (m *_GDRenshuDBMgr) DeleteByPrimaryKey(id int64) (int64, error) {
	pk := &IdOfGDRenshuPK{
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM gd_renshu %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, id int64) (int64, error) {
	pk := &IdOfGDRenshuPK{
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM gd_renshu %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM gd_renshu")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM gd_renshu WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDRenshuDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM gd_renshu")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM gd_renshu WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
