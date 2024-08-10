// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiler

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AuthSubject is an object representing the database table.
type AuthSubject struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Subject   string    `boil:"subject" json:"subject" toml:"subject" yaml:"subject"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *authSubjectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authSubjectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthSubjectColumns = struct {
	ID        string
	UserID    string
	Subject   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Subject:   "subject",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var AuthSubjectTableColumns = struct {
	ID        string
	UserID    string
	Subject   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "auth_subjects.id",
	UserID:    "auth_subjects.user_id",
	Subject:   "auth_subjects.subject",
	CreatedAt: "auth_subjects.created_at",
	UpdatedAt: "auth_subjects.updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AuthSubjectWhere = struct {
	ID        whereHelperint
	UserID    whereHelperint
	Subject   whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"auth_subjects\".\"id\""},
	UserID:    whereHelperint{field: "\"auth_subjects\".\"user_id\""},
	Subject:   whereHelperstring{field: "\"auth_subjects\".\"subject\""},
	CreatedAt: whereHelpertime_Time{field: "\"auth_subjects\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"auth_subjects\".\"updated_at\""},
}

// AuthSubjectRels is where relationship names are stored.
var AuthSubjectRels = struct {
	User string
}{
	User: "User",
}

// authSubjectR is where relationships are stored.
type authSubjectR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*authSubjectR) NewStruct() *authSubjectR {
	return &authSubjectR{}
}

func (r *authSubjectR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// authSubjectL is where Load methods for each relationship are stored.
type authSubjectL struct{}

var (
	authSubjectAllColumns            = []string{"id", "user_id", "subject", "created_at", "updated_at"}
	authSubjectColumnsWithoutDefault = []string{"user_id", "subject", "created_at", "updated_at"}
	authSubjectColumnsWithDefault    = []string{"id"}
	authSubjectPrimaryKeyColumns     = []string{"id"}
	authSubjectGeneratedColumns      = []string{"id"}
)

type (
	// AuthSubjectSlice is an alias for a slice of pointers to AuthSubject.
	// This should almost always be used instead of []AuthSubject.
	AuthSubjectSlice []*AuthSubject
	// AuthSubjectHook is the signature for custom AuthSubject hook methods
	AuthSubjectHook func(context.Context, boil.ContextExecutor, *AuthSubject) error

	authSubjectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authSubjectType                 = reflect.TypeOf(&AuthSubject{})
	authSubjectMapping              = queries.MakeStructMapping(authSubjectType)
	authSubjectPrimaryKeyMapping, _ = queries.BindMapping(authSubjectType, authSubjectMapping, authSubjectPrimaryKeyColumns)
	authSubjectInsertCacheMut       sync.RWMutex
	authSubjectInsertCache          = make(map[string]insertCache)
	authSubjectUpdateCacheMut       sync.RWMutex
	authSubjectUpdateCache          = make(map[string]updateCache)
	authSubjectUpsertCacheMut       sync.RWMutex
	authSubjectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authSubjectAfterSelectMu sync.Mutex
var authSubjectAfterSelectHooks []AuthSubjectHook

var authSubjectBeforeInsertMu sync.Mutex
var authSubjectBeforeInsertHooks []AuthSubjectHook
var authSubjectAfterInsertMu sync.Mutex
var authSubjectAfterInsertHooks []AuthSubjectHook

var authSubjectBeforeUpdateMu sync.Mutex
var authSubjectBeforeUpdateHooks []AuthSubjectHook
var authSubjectAfterUpdateMu sync.Mutex
var authSubjectAfterUpdateHooks []AuthSubjectHook

var authSubjectBeforeDeleteMu sync.Mutex
var authSubjectBeforeDeleteHooks []AuthSubjectHook
var authSubjectAfterDeleteMu sync.Mutex
var authSubjectAfterDeleteHooks []AuthSubjectHook

var authSubjectBeforeUpsertMu sync.Mutex
var authSubjectBeforeUpsertHooks []AuthSubjectHook
var authSubjectAfterUpsertMu sync.Mutex
var authSubjectAfterUpsertHooks []AuthSubjectHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthSubject) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthSubject) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthSubject) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthSubject) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthSubject) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthSubject) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthSubject) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthSubject) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthSubject) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authSubjectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthSubjectHook registers your hook function for all future operations.
func AddAuthSubjectHook(hookPoint boil.HookPoint, authSubjectHook AuthSubjectHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		authSubjectAfterSelectMu.Lock()
		authSubjectAfterSelectHooks = append(authSubjectAfterSelectHooks, authSubjectHook)
		authSubjectAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		authSubjectBeforeInsertMu.Lock()
		authSubjectBeforeInsertHooks = append(authSubjectBeforeInsertHooks, authSubjectHook)
		authSubjectBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		authSubjectAfterInsertMu.Lock()
		authSubjectAfterInsertHooks = append(authSubjectAfterInsertHooks, authSubjectHook)
		authSubjectAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		authSubjectBeforeUpdateMu.Lock()
		authSubjectBeforeUpdateHooks = append(authSubjectBeforeUpdateHooks, authSubjectHook)
		authSubjectBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		authSubjectAfterUpdateMu.Lock()
		authSubjectAfterUpdateHooks = append(authSubjectAfterUpdateHooks, authSubjectHook)
		authSubjectAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		authSubjectBeforeDeleteMu.Lock()
		authSubjectBeforeDeleteHooks = append(authSubjectBeforeDeleteHooks, authSubjectHook)
		authSubjectBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		authSubjectAfterDeleteMu.Lock()
		authSubjectAfterDeleteHooks = append(authSubjectAfterDeleteHooks, authSubjectHook)
		authSubjectAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		authSubjectBeforeUpsertMu.Lock()
		authSubjectBeforeUpsertHooks = append(authSubjectBeforeUpsertHooks, authSubjectHook)
		authSubjectBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		authSubjectAfterUpsertMu.Lock()
		authSubjectAfterUpsertHooks = append(authSubjectAfterUpsertHooks, authSubjectHook)
		authSubjectAfterUpsertMu.Unlock()
	}
}

// One returns a single authSubject record from the query.
func (q authSubjectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthSubject, error) {
	o := &AuthSubject{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: failed to execute a one query for auth_subjects")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthSubject records from the query.
func (q authSubjectQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthSubjectSlice, error) {
	var o []*AuthSubject

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "boiler: failed to assign all query results to AuthSubject slice")
	}

	if len(authSubjectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthSubject records in the query.
func (q authSubjectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to count auth_subjects rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authSubjectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "boiler: failed to check if auth_subjects exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AuthSubject) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authSubjectL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthSubject interface{}, mods queries.Applicator) error {
	var slice []*AuthSubject
	var object *AuthSubject

	if singular {
		var ok bool
		object, ok = maybeAuthSubject.(*AuthSubject)
		if !ok {
			object = new(AuthSubject)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthSubject)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthSubject))
			}
		}
	} else {
		s, ok := maybeAuthSubject.(*[]*AuthSubject)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthSubject)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthSubject))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &authSubjectR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authSubjectR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.AuthSubjects = append(foreign.R.AuthSubjects, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.AuthSubjects = append(foreign.R.AuthSubjects, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the authSubject to the related item.
// Sets o.R.User to related.
// Adds o to related.R.AuthSubjects.
func (o *AuthSubject) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_subjects\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, authSubjectPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &authSubjectR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			AuthSubjects: AuthSubjectSlice{o},
		}
	} else {
		related.R.AuthSubjects = append(related.R.AuthSubjects, o)
	}

	return nil
}

// AuthSubjects retrieves all the records using an executor.
func AuthSubjects(mods ...qm.QueryMod) authSubjectQuery {
	mods = append(mods, qm.From("\"auth_subjects\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"auth_subjects\".*"})
	}

	return authSubjectQuery{q}
}

// FindAuthSubject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthSubject(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AuthSubject, error) {
	authSubjectObj := &AuthSubject{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_subjects\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authSubjectObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "boiler: unable to select from auth_subjects")
	}

	if err = authSubjectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authSubjectObj, err
	}

	return authSubjectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthSubject) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("boiler: no auth_subjects provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authSubjectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authSubjectInsertCacheMut.RLock()
	cache, cached := authSubjectInsertCache[key]
	authSubjectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authSubjectAllColumns,
			authSubjectColumnsWithDefault,
			authSubjectColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, authSubjectGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(authSubjectType, authSubjectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authSubjectType, authSubjectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auth_subjects\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auth_subjects\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "boiler: unable to insert into auth_subjects")
	}

	if !cached {
		authSubjectInsertCacheMut.Lock()
		authSubjectInsertCache[key] = cache
		authSubjectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthSubject.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthSubject) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authSubjectUpdateCacheMut.RLock()
	cache, cached := authSubjectUpdateCache[key]
	authSubjectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authSubjectAllColumns,
			authSubjectPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, authSubjectGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("boiler: unable to update auth_subjects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_subjects\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authSubjectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authSubjectType, authSubjectMapping, append(wl, authSubjectPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update auth_subjects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by update for auth_subjects")
	}

	if !cached {
		authSubjectUpdateCacheMut.Lock()
		authSubjectUpdateCache[key] = cache
		authSubjectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authSubjectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all for auth_subjects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected for auth_subjects")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthSubjectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("boiler: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authSubjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auth_subjects\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authSubjectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to update all in authSubject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to retrieve rows affected all in update all authSubject")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthSubject) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("boiler: no auth_subjects provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authSubjectColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	authSubjectUpsertCacheMut.RLock()
	cache, cached := authSubjectUpsertCache[key]
	authSubjectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			authSubjectAllColumns,
			authSubjectColumnsWithDefault,
			authSubjectColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			authSubjectAllColumns,
			authSubjectPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, authSubjectGeneratedColumns)
		update = strmangle.SetComplement(update, authSubjectGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("boiler: unable to upsert auth_subjects, could not build update column list")
		}

		ret := strmangle.SetComplement(authSubjectAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(authSubjectPrimaryKeyColumns) == 0 {
				return errors.New("boiler: unable to upsert auth_subjects, could not build conflict column list")
			}

			conflict = make([]string, len(authSubjectPrimaryKeyColumns))
			copy(conflict, authSubjectPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"auth_subjects\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(authSubjectType, authSubjectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authSubjectType, authSubjectMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "boiler: unable to upsert auth_subjects")
	}

	if !cached {
		authSubjectUpsertCacheMut.Lock()
		authSubjectUpsertCache[key] = cache
		authSubjectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuthSubject record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthSubject) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("boiler: no AuthSubject provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authSubjectPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_subjects\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete from auth_subjects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by delete for auth_subjects")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authSubjectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("boiler: no authSubjectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from auth_subjects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for auth_subjects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthSubjectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authSubjectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authSubjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auth_subjects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authSubjectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "boiler: unable to delete all from authSubject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "boiler: failed to get rows affected by deleteall for auth_subjects")
	}

	if len(authSubjectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthSubject) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthSubject(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthSubjectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthSubjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authSubjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auth_subjects\".* FROM \"auth_subjects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authSubjectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "boiler: unable to reload all in AuthSubjectSlice")
	}

	*o = slice

	return nil
}

// AuthSubjectExists checks if the AuthSubject row exists.
func AuthSubjectExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auth_subjects\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "boiler: unable to check if auth_subjects exists")
	}

	return exists, nil
}

// Exists checks if the AuthSubject row exists.
func (o *AuthSubject) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AuthSubjectExists(ctx, exec, o.ID)
}
