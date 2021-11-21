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

type GDNiusan struct {
	Id           int64     `db:"id"`
	SecurityCode int32     `db:"security_code"`
	Niusan       string    `db:"niusan"`
	Disabled     bool      `db:"disabled"`
	CreateDate   time.Time `db:"create_date"`
	UpdateDate   time.Time `db:"update_date"`
}

var GDNiusanColumns = struct {
	Id           string
	SecurityCode string
	Niusan       string
	Disabled     string
	CreateDate   string
	UpdateDate   string
}{
	"id",
	"security_code",
	"niusan",
	"disabled",
	"create_date",
	"update_date",
}

type _GDNiusanMgr struct {
}

var GDNiusanMgr *_GDNiusanMgr

func (m *_GDNiusanMgr) NewGDNiusan() *GDNiusan {
	return &GDNiusan{}
}

//! object function

func (obj *GDNiusan) GetNameSpace() string {
	return "model"
}

func (obj *GDNiusan) GetClassName() string {
	return "GDNiusan"
}

func (obj *GDNiusan) GetTableName() string {
	return "gd_niusan"
}

func (obj *GDNiusan) GetColumns() []string {
	columns := []string{
		"gd_niusan.`id`",
		"gd_niusan.`security_code`",
		"gd_niusan.`niusan`",
		"gd_niusan.`disabled`",
		"gd_niusan.`create_date`",
		"gd_niusan.`update_date`",
	}
	return columns
}

func (obj *GDNiusan) GetNoneIncrementColumns() []string {
	columns := []string{
		"`security_code`",
		"`niusan`",
		"`disabled`",
		"`create_date`",
		"`update_date`",
	}
	return columns
}

func (obj *GDNiusan) GetPrimaryKey() PrimaryKey {
	pk := GDNiusanMgr.NewPrimaryKey()
	pk.Id = obj.Id
	pk.Id = obj.Id
	return pk
}

func (obj *GDNiusan) Validate() error {
	validate := validator.New()
	return validate.Struct(obj)
}

//! primary key

type IdIdOfGDNiusanPK struct {
	Id int64
	Id int64
}

func (m *_GDNiusanMgr) NewPrimaryKey() *IdIdOfGDNiusanPK {
	return &IdIdOfGDNiusanPK{}
}

func (u *IdIdOfGDNiusanPK) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
		"Id",
		fmt.Sprint(u.Id),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdIdOfGDNiusanPK) Parse(key string) error {
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
	vId, ok := kv["Id"]
	if !ok {
		return fmt.Errorf("key (%s) without (Id) field", key)
	}
	if err := orm.StringScan(vId, &(u.Id)); err != nil {
		return err
	}
	return nil
}

func (u *IdIdOfGDNiusanPK) SQLFormat() string {
	conditions := []string{
		"`id` = ?",
		"`id` = ?",
	}
	return orm.SQLWhere(conditions)
}

func (u *IdIdOfGDNiusanPK) SQLParams() []interface{} {
	return []interface{}{
		u.Id,
		u.Id,
	}
}

func (u *IdIdOfGDNiusanPK) Columns() []string {
	return []string{
		"`id`",
		"`id`",
	}
}

//! uniques

//! indexes

type SecurityCodeOfGDNiusanIDX struct {
	SecurityCode int32
	offset       int
	limit        int
}

func (u *SecurityCodeOfGDNiusanIDX) Key() string {
	strs := []string{
		"SecurityCode",
		fmt.Sprint(u.SecurityCode),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *SecurityCodeOfGDNiusanIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`security_code` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *SecurityCodeOfGDNiusanIDX) SQLParams() []interface{} {
	return []interface{}{
		u.SecurityCode,
	}
}

func (u *SecurityCodeOfGDNiusanIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *SecurityCodeOfGDNiusanIDX) Limit(n int) {
	u.limit = n
}

func (u *SecurityCodeOfGDNiusanIDX) Offset(n int) {
	u.offset = n
}

func (u *SecurityCodeOfGDNiusanIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *SecurityCodeOfGDNiusanIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

type NiusanOfGDNiusanIDX struct {
	Niusan string
	offset int
	limit  int
}

func (u *NiusanOfGDNiusanIDX) Key() string {
	strs := []string{
		"Niusan",
		fmt.Sprint(u.Niusan),
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *NiusanOfGDNiusanIDX) SQLFormat(limit bool) string {
	conditions := []string{
		"`niusan` = ?",
	}
	if limit {
		return fmt.Sprintf("%s %s", orm.SQLWhere(conditions), orm.SQLOffsetLimit(u.offset, u.limit))
	}
	return orm.SQLWhere(conditions)
}

func (u *NiusanOfGDNiusanIDX) SQLParams() []interface{} {
	return []interface{}{
		u.Niusan,
	}
}

func (u *NiusanOfGDNiusanIDX) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *NiusanOfGDNiusanIDX) Limit(n int) {
	u.limit = n
}

func (u *NiusanOfGDNiusanIDX) Offset(n int) {
	u.offset = n
}

func (u *NiusanOfGDNiusanIDX) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *NiusanOfGDNiusanIDX) IDXRelation(store *orm.RedisStore) IndexRelation {
	return nil
}

//! ranges

type IdIdOfGDNiusanRNG struct {
	Id           int64
	IdBegin      int64
	IdEnd        int64
	offset       int
	limit        int
	includeBegin bool
	includeEnd   bool
	revert       bool
}

func (u *IdIdOfGDNiusanRNG) Key() string {
	strs := []string{
		"Id",
		fmt.Sprint(u.Id),
		"Id",
	}
	return fmt.Sprintf("%s", strings.Join(strs, ":"))
}

func (u *IdIdOfGDNiusanRNG) beginOp() string {
	if u.includeBegin {
		return ">="
	}
	return ">"
}
func (u *IdIdOfGDNiusanRNG) endOp() string {
	if u.includeBegin {
		return "<="
	}
	return "<"
}

func (u *IdIdOfGDNiusanRNG) SQLFormat(limit bool) string {
	conditions := []string{}
	conditions = append(conditions, "`id` = ?")
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

func (u *IdIdOfGDNiusanRNG) SQLParams() []interface{} {
	params := []interface{}{
		u.Id,
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

func (u *IdIdOfGDNiusanRNG) SQLLimit() int {
	if u.limit > 0 {
		return u.limit
	}
	return -1
}

func (u *IdIdOfGDNiusanRNG) Limit(n int) {
	u.limit = n
}

func (u *IdIdOfGDNiusanRNG) Offset(n int) {
	u.offset = n
}

func (u *IdIdOfGDNiusanRNG) PositionOffsetLimit(len int) (int, int) {
	if u.limit <= 0 {
		return 0, len
	}
	if u.offset+u.limit > len {
		return u.offset, len
	}
	return u.offset, u.limit
}

func (u *IdIdOfGDNiusanRNG) Begin() int64 {
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

func (u *IdIdOfGDNiusanRNG) End() int64 {
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

func (u *IdIdOfGDNiusanRNG) Revert(b bool) {
	u.revert = b
}

func (u *IdIdOfGDNiusanRNG) IncludeBegin(f bool) {
	u.includeBegin = f
}

func (u *IdIdOfGDNiusanRNG) IncludeEnd(f bool) {
	u.includeEnd = f
}

func (u *IdIdOfGDNiusanRNG) RNGRelation(store *orm.RedisStore) RangeRelation {
	return nil
}

type _GDNiusanDBMgr struct {
	db orm.DB
}

func (m *_GDNiusanMgr) DB(db orm.DB) *_GDNiusanDBMgr {
	return GDNiusanDBMgr(db)
}

func GDNiusanDBMgr(db orm.DB) *_GDNiusanDBMgr {
	if db == nil {
		panic(fmt.Errorf("GDNiusanDBMgr init need db"))
	}
	return &_GDNiusanDBMgr{db: db}
}

func (m *_GDNiusanDBMgr) Search(where string, orderby string, limit string, args ...interface{}) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQL(query, args...)
}

func (m *_GDNiusanDBMgr) SearchContext(ctx context.Context, where string, orderby string, limit string, args ...interface{}) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()

	if limit = strings.ToUpper(strings.TrimSpace(limit)); limit != "" && !strings.HasPrefix(limit, "LIMIT") {
		limit = "LIMIT " + limit
	}

	conditions := []string{where, orderby, limit}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), strings.Join(conditions, " "))
	return m.FetchBySQLContext(ctx, query, args...)
}

func (m *_GDNiusanDBMgr) SearchConditions(conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	q := fmt.Sprintf("SELECT %s FROM gd_niusan %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQL(q, args...)
}

func (m *_GDNiusanDBMgr) SearchConditionsContext(ctx context.Context, conditions []string, orderby string, offset int, limit int, args ...interface{}) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	q := fmt.Sprintf("SELECT %s FROM gd_niusan %s %s %s",
		strings.Join(obj.GetColumns(), ","),
		orm.SQLWhere(conditions),
		orderby,
		orm.SQLOffsetLimit(offset, limit))

	return m.FetchBySQLContext(ctx, q, args...)
}

func (m *_GDNiusanDBMgr) SearchCount(where string, args ...interface{}) (int64, error) {
	return m.queryCount(where, args...)
}

func (m *_GDNiusanDBMgr) SearchCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, where, args...)
}

func (m *_GDNiusanDBMgr) SearchConditionsCount(conditions []string, args ...interface{}) (int64, error) {
	return m.queryCount(orm.SQLWhere(conditions), args...)
}

func (m *_GDNiusanDBMgr) SearchConditionsCountContext(ctx context.Context, conditions []string, args ...interface{}) (int64, error) {
	return m.queryCountContext(ctx, orm.SQLWhere(conditions), args...)
}

func (m *_GDNiusanDBMgr) FetchBySQL(q string, args ...interface{}) (results []*GDNiusan, err error) {
	rows, err := m.db.Query(q, args...)
	if err != nil {
		return nil, fmt.Errorf("GDNiusan fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result GDNiusan
		err = rows.Scan(&(result.Id), &(result.SecurityCode), &(result.Niusan), &(result.Disabled), &CreateDate, &UpdateDate)
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
		return nil, fmt.Errorf("GDNiusan fetch result error: %v", err)
	}
	return
}

func (m *_GDNiusanDBMgr) FetchBySQLContext(ctx context.Context, q string, args ...interface{}) (results []*GDNiusan, err error) {
	rows, err := m.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("GDNiusan fetch error: %v", err)
	}
	defer rows.Close()

	var CreateDate int64
	var UpdateDate int64

	for rows.Next() {
		var result GDNiusan
		err = rows.Scan(&(result.Id), &(result.SecurityCode), &(result.Niusan), &(result.Disabled), &CreateDate, &UpdateDate)
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
		return nil, fmt.Errorf("GDNiusan fetch result error: %v", err)
	}
	return
}
func (m *_GDNiusanDBMgr) Exist(pk PrimaryKey) (bool, error) {
	c, err := m.queryCount(pk.SQLFormat(), pk.SQLParams()...)
	if err != nil {
		return false, err
	}
	return (c != 0), nil
}

// Deprecated: Use FetchByPrimaryKey instead.
func (m *_GDNiusanDBMgr) Fetch(pk PrimaryKey) (*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDNiusan fetch record not found")
}

// err not found check
func (m *_GDNiusanDBMgr) IsErrNotFound(err error) bool {
	return strings.Contains(err.Error(), "not found") || err == sql.ErrNoRows
}

// primary key
func (m *_GDNiusanDBMgr) FetchByPrimaryKey(id int64, id int64) (*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	pk := &IdIdOfGDNiusanPK{
		Id: id,
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQL(query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDNiusan fetch record not found")
}

func (m *_GDNiusanDBMgr) FetchByPrimaryKeyContext(ctx context.Context, id int64, id int64) (*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	pk := &IdIdOfGDNiusanPK{
		Id: id,
		Id: id,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), pk.SQLFormat())
	objs, err := m.FetchBySQLContext(ctx, query, pk.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDNiusan fetch record not found")
}

// indexes

func (m *_GDNiusanDBMgr) FindBySecurityCode(securityCode int32, limit int, offset int) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &SecurityCodeOfGDNiusanIDX{
		SecurityCode: securityCode,
		limit:        limit,
		offset:       offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindBySecurityCodeContext(ctx context.Context, securityCode int32, limit int, offset int) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &SecurityCodeOfGDNiusanIDX{
		SecurityCode: securityCode,
		limit:        limit,
		offset:       offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindAllBySecurityCode(securityCode int32) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &SecurityCodeOfGDNiusanIDX{
		SecurityCode: securityCode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindAllBySecurityCodeContext(ctx context.Context, securityCode int32) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &SecurityCodeOfGDNiusanIDX{
		SecurityCode: securityCode,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindBySecurityCodeGroup(items []int32) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan where `security_code` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDNiusanDBMgr) FindBySecurityCodeGroupContext(ctx context.Context, items []int32) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan where `security_code` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

func (m *_GDNiusanDBMgr) FindByNiusan(niusan string, limit int, offset int) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &NiusanOfGDNiusanIDX{
		Niusan: niusan,
		limit:  limit,
		offset: offset,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindByNiusanContext(ctx context.Context, niusan string, limit int, offset int) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &NiusanOfGDNiusanIDX{
		Niusan: niusan,
		limit:  limit,
		offset: offset,
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindAllByNiusan(niusan string) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &NiusanOfGDNiusanIDX{
		Niusan: niusan,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQL(query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindAllByNiusanContext(ctx context.Context, niusan string) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	idx_ := &NiusanOfGDNiusanIDX{
		Niusan: niusan,
	}

	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), idx_.SQLFormat(true))
	return m.FetchBySQLContext(ctx, query, idx_.SQLParams()...)
}

func (m *_GDNiusanDBMgr) FindByNiusanGroup(items []string) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan where `niusan` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQL(query, params...)
}

func (m *_GDNiusanDBMgr) FindByNiusanGroupContext(ctx context.Context, items []string) ([]*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	if len(items) == 0 {
		return nil, nil
	}
	params := make([]interface{}, 0, len(items))
	for _, item := range items {
		params = append(params, item)
	}
	query := fmt.Sprintf("SELECT %s FROM gd_niusan where `niusan` in (?", strings.Join(obj.GetColumns(), ",")) +
		strings.Repeat(",?", len(items)-1) + ")"
	return m.FetchBySQLContext(ctx, query, params...)
}

// uniques

func (m *_GDNiusanDBMgr) FindOne(unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimit(unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDNiusan find record not found")
}

func (m *_GDNiusanDBMgr) FindOneContext(ctx context.Context, unique Unique) (PrimaryKey, error) {
	objs, err := m.queryLimitContext(ctx, unique.SQLFormat(true), unique.SQLLimit(), unique.SQLParams()...)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return objs[0], nil
	}
	return nil, fmt.Errorf("GDNiusan find record not found")
}

// Deprecated: Use FetchByXXXUnique instead.
func (m *_GDNiusanDBMgr) FindOneFetch(unique Unique) (*GDNiusan, error) {
	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), unique.SQLFormat(true))
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
func (m *_GDNiusanDBMgr) Find(index Index) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(index.SQLFormat(true), index.SQLLimit(), index.SQLParams()...)
	return total, pks, err
}

func (m *_GDNiusanDBMgr) FindFetch(index Index) (int64, []*GDNiusan, error) {
	total, err := m.queryCount(index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDNiusanDBMgr) FindFetchContext(ctx context.Context, index Index) (int64, []*GDNiusan, error) {
	total, err := m.queryCountContext(ctx, index.SQLFormat(false), index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}

	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), index.SQLFormat(true))
	results, err := m.FetchBySQL(query, index.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDNiusanDBMgr) Range(scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimit(scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_GDNiusanDBMgr) RangeContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	pks, err := m.queryLimitContext(ctx, scope.SQLFormat(true), scope.SQLLimit(), scope.SQLParams()...)
	return total, pks, err
}

func (m *_GDNiusanDBMgr) RangeFetch(scope Range) (int64, []*GDNiusan, error) {
	total, err := m.queryCount(scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQL(query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDNiusanDBMgr) RangeFetchContext(ctx context.Context, scope Range) (int64, []*GDNiusan, error) {
	total, err := m.queryCountContext(ctx, scope.SQLFormat(false), scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	obj := GDNiusanMgr.NewGDNiusan()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(obj.GetColumns(), ","), scope.SQLFormat(true))
	results, err := m.FetchBySQLContext(ctx, query, scope.SQLParams()...)
	if err != nil {
		return total, nil, err
	}
	return total, results, nil
}

func (m *_GDNiusanDBMgr) RangeRevert(scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.Range(scope)
}

func (m *_GDNiusanDBMgr) RangeRevertContext(ctx context.Context, scope Range) (int64, []PrimaryKey, error) {
	scope.Revert(true)
	return m.RangeContext(ctx, scope)
}

func (m *_GDNiusanDBMgr) RangeRevertFetch(scope Range) (int64, []*GDNiusan, error) {
	scope.Revert(true)
	return m.RangeFetch(scope)
}

func (m *_GDNiusanDBMgr) RangeRevertFetchContext(ctx context.Context, scope Range) (int64, []*GDNiusan, error) {
	scope.Revert(true)
	return m.RangeFetchContext(ctx, scope)
}

func (m *_GDNiusanDBMgr) queryLimit(where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := GDNiusanMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("GDNiusan query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := GDNiusanMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id), &(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("GDNiusan query limit result error: %v", err)
	}
	return
}

func (m *_GDNiusanDBMgr) queryLimitContext(ctx context.Context, where string, limit int, args ...interface{}) (results []PrimaryKey, err error) {
	pk := GDNiusanMgr.NewPrimaryKey()
	query := fmt.Sprintf("SELECT %s FROM gd_niusan %s", strings.Join(pk.Columns(), ","), where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("GDNiusan query limit error: %v", err)
	}
	defer rows.Close()

	offset := 0

	for rows.Next() {
		if limit >= 0 && offset >= limit {
			break
		}
		offset++

		result := GDNiusanMgr.NewPrimaryKey()
		err = rows.Scan(&(result.Id), &(result.Id))
		if err != nil {
			m.db.SetError(err)
			return nil, err
		}

		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		m.db.SetError(err)
		return nil, fmt.Errorf("GDNiusan query limit result error: %v", err)
	}
	return
}

func (m *_GDNiusanDBMgr) queryCount(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM gd_niusan %s", where)
	rows, err := m.db.Query(query, args...)
	if err != nil {
		return 0, fmt.Errorf("GDNiusan query count error: %v", err)
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

func (m *_GDNiusanDBMgr) queryCountContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("SELECT count(`id`) FROM gd_niusan %s", where)
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("GDNiusan query count error: %v", err)
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

func (m *_GDNiusanDBMgr) BatchCreate(objs []*GDNiusan) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*5)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(5, "?"), ",")))
		values = append(values, obj.SecurityCode)
		values = append(values, obj.Niusan)
		values = append(values, obj.Disabled)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO gd_niusan(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
	result, err := m.db.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) BatchCreateContext(ctx context.Context, objs []*GDNiusan) (int64, error) {
	if len(objs) == 0 {
		return 0, nil
	}

	params := make([]string, 0, len(objs))
	values := make([]interface{}, 0, len(objs)*5)
	for _, obj := range objs {
		params = append(params, fmt.Sprintf("(%s)", strings.Join(orm.NewStringSlice(5, "?"), ",")))
		values = append(values, obj.SecurityCode)
		values = append(values, obj.Niusan)
		values = append(values, obj.Disabled)
		values = append(values, obj.CreateDate.Unix())
		values = append(values, obj.UpdateDate.Unix())
	}
	query := fmt.Sprintf("INSERT INTO gd_niusan(%s) VALUES %s", strings.Join(objs[0].GetNoneIncrementColumns(), ","), strings.Join(params, ","))
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
func (m *_GDNiusanDBMgr) UpdateBySQL(set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE gd_niusan SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE gd_niusan SET %s WHERE %s", set, where)
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
func (m *_GDNiusanDBMgr) UpdateBySQLContext(ctx context.Context, set, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE gd_niusan SET %s", set)
	if where != "" {
		query = fmt.Sprintf("UPDATE gd_niusan SET %s WHERE %s", set, where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) Create(obj *GDNiusan) (int64, error) {
	params := orm.NewStringSlice(5, "?")
	q := fmt.Sprintf("INSERT INTO gd_niusan(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 6)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.Niusan)
	values = append(values, obj.Disabled)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) CreateContext(ctx context.Context, obj *GDNiusan) (int64, error) {
	params := orm.NewStringSlice(5, "?")
	q := fmt.Sprintf("INSERT INTO gd_niusan(%s) VALUES(%s)",
		strings.Join(obj.GetNoneIncrementColumns(), ","),
		strings.Join(params, ","))

	values := make([]interface{}, 0, 6)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.Niusan)
	values = append(values, obj.Disabled)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) Update(obj *GDNiusan) (int64, error) {
	columns := []string{
		"`security_code` = ?",
		"`niusan` = ?",
		"`disabled` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE gd_niusan SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 6-2)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.Niusan)
	values = append(values, obj.Disabled)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.Exec(q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) UpdateContext(ctx context.Context, obj *GDNiusan) (int64, error) {
	columns := []string{
		"`security_code` = ?",
		"`niusan` = ?",
		"`disabled` = ?",
		"`create_date` = ?",
		"`update_date` = ?",
	}

	pk := obj.GetPrimaryKey()
	q := fmt.Sprintf("UPDATE gd_niusan SET %s %s", strings.Join(columns, ","), pk.SQLFormat())
	values := make([]interface{}, 0, 6-2)
	values = append(values, obj.SecurityCode)
	values = append(values, obj.Niusan)
	values = append(values, obj.Disabled)
	values = append(values, obj.CreateDate.Unix())
	values = append(values, obj.UpdateDate.Unix())
	values = append(values, pk.SQLParams()...)

	result, err := m.db.ExecContext(ctx, q, values...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) Save(obj *GDNiusan) (int64, error) {
	affected, err := m.Update(obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.Create(obj)
	}
	return affected, err
}

func (m *_GDNiusanDBMgr) SaveContext(ctx context.Context, obj *GDNiusan) (int64, error) {
	affected, err := m.UpdateContext(ctx, obj)
	if err != nil {
		return affected, err
	}
	if affected == 0 {
		return m.CreateContext(ctx, obj)
	}
	return affected, err
}

func (m *_GDNiusanDBMgr) Delete(obj *GDNiusan) (int64, error) {
	return m.DeleteByPrimaryKey(obj.Id, obj.Id)
}

func (m *_GDNiusanDBMgr) DeleteContext(ctx context.Context, obj *GDNiusan) (int64, error) {
	return m.DeleteByPrimaryKeyContext(ctx, obj.Id, obj.Id)
}

func (m *_GDNiusanDBMgr) DeleteByPrimaryKey(id int64, id int64) (int64, error) {
	pk := &IdIdOfGDNiusanPK{
		Id: id,
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM gd_niusan %s", pk.SQLFormat())
	result, err := m.db.Exec(q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) DeleteByPrimaryKeyContext(ctx context.Context, id int64, id int64) (int64, error) {
	pk := &IdIdOfGDNiusanPK{
		Id: id,
		Id: id,
	}
	q := fmt.Sprintf("DELETE FROM gd_niusan %s", pk.SQLFormat())
	result, err := m.db.ExecContext(ctx, q, pk.SQLParams()...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) DeleteBySQL(where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM gd_niusan")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM gd_niusan WHERE %s", where)
	}
	result, err := m.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *_GDNiusanDBMgr) DeleteBySQLContext(ctx context.Context, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("DELETE FROM gd_niusan")
	if where != "" {
		query = fmt.Sprintf("DELETE FROM gd_niusan WHERE %s", where)
	}
	result, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
