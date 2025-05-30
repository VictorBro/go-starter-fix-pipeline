// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// AccessToken is an object representing the database table.
type AccessToken struct {
	Token      string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	ValidUntil time.Time `boil:"valid_until" json:"valid_until" toml:"valid_until" yaml:"valid_until"`
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *accessTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accessTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccessTokenColumns = struct {
	Token      string
	ValidUntil string
	UserID     string
	CreatedAt  string
	UpdatedAt  string
}{
	Token:      "token",
	ValidUntil: "valid_until",
	UserID:     "user_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var AccessTokenTableColumns = struct {
	Token      string
	ValidUntil string
	UserID     string
	CreatedAt  string
	UpdatedAt  string
}{
	Token:      "access_tokens.token",
	ValidUntil: "access_tokens.valid_until",
	UserID:     "access_tokens.user_id",
	CreatedAt:  "access_tokens.created_at",
	UpdatedAt:  "access_tokens.updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod      { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod    { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod   { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod   { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) SIMILAR(x string) qm.QueryMod { return qm.Where(w.field+" SIMILAR TO ?", x) }
func (w whereHelperstring) NSIMILAR(x string) qm.QueryMod {
	return qm.Where(w.field+" NOT SIMILAR TO ?", x)
}
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

var AccessTokenWhere = struct {
	Token      whereHelperstring
	ValidUntil whereHelpertime_Time
	UserID     whereHelperstring
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	Token:      whereHelperstring{field: "\"access_tokens\".\"token\""},
	ValidUntil: whereHelpertime_Time{field: "\"access_tokens\".\"valid_until\""},
	UserID:     whereHelperstring{field: "\"access_tokens\".\"user_id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"access_tokens\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"access_tokens\".\"updated_at\""},
}

// AccessTokenRels is where relationship names are stored.
var AccessTokenRels = struct {
	User string
}{
	User: "User",
}

// accessTokenR is where relationships are stored.
type accessTokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*accessTokenR) NewStruct() *accessTokenR {
	return &accessTokenR{}
}

func (r *accessTokenR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// accessTokenL is where Load methods for each relationship are stored.
type accessTokenL struct{}

var (
	accessTokenAllColumns            = []string{"token", "valid_until", "user_id", "created_at", "updated_at"}
	accessTokenColumnsWithoutDefault = []string{"valid_until", "user_id", "created_at", "updated_at"}
	accessTokenColumnsWithDefault    = []string{"token"}
	accessTokenPrimaryKeyColumns     = []string{"token"}
	accessTokenGeneratedColumns      = []string{}
)

type (
	// AccessTokenSlice is an alias for a slice of pointers to AccessToken.
	// This should almost always be used instead of []AccessToken.
	AccessTokenSlice []*AccessToken

	accessTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accessTokenType                 = reflect.TypeOf(&AccessToken{})
	accessTokenMapping              = queries.MakeStructMapping(accessTokenType)
	accessTokenPrimaryKeyMapping, _ = queries.BindMapping(accessTokenType, accessTokenMapping, accessTokenPrimaryKeyColumns)
	accessTokenInsertCacheMut       sync.RWMutex
	accessTokenInsertCache          = make(map[string]insertCache)
	accessTokenUpdateCacheMut       sync.RWMutex
	accessTokenUpdateCache          = make(map[string]updateCache)
	accessTokenUpsertCacheMut       sync.RWMutex
	accessTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single accessToken record from the query.
func (q accessTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccessToken, error) {
	o := &AccessToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for access_tokens")
	}

	return o, nil
}

// All returns all AccessToken records from the query.
func (q accessTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccessTokenSlice, error) {
	var o []*AccessToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AccessToken slice")
	}

	return o, nil
}

// Count returns the count of all AccessToken records in the query.
func (q accessTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count access_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accessTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if access_tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AccessToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (accessTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccessToken interface{}, mods queries.Applicator) error {
	var slice []*AccessToken
	var object *AccessToken

	if singular {
		var ok bool
		object, ok = maybeAccessToken.(*AccessToken)
		if !ok {
			object = new(AccessToken)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAccessToken)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAccessToken))
			}
		}
	} else {
		s, ok := maybeAccessToken.(*[]*AccessToken)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAccessToken)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAccessToken))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &accessTokenR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accessTokenR{}
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.AccessTokens = append(foreign.R.AccessTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.AccessTokens = append(foreign.R.AccessTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the accessToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.AccessTokens.
func (o *AccessToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"access_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, accessTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Token}

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
		o.R = &accessTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			AccessTokens: AccessTokenSlice{o},
		}
	} else {
		related.R.AccessTokens = append(related.R.AccessTokens, o)
	}

	return nil
}

// AccessTokens retrieves all the records using an executor.
func AccessTokens(mods ...qm.QueryMod) accessTokenQuery {
	mods = append(mods, qm.From("\"access_tokens\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"access_tokens\".*"})
	}

	return accessTokenQuery{q}
}

// FindAccessToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccessToken(ctx context.Context, exec boil.ContextExecutor, token string, selectCols ...string) (*AccessToken, error) {
	accessTokenObj := &AccessToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"access_tokens\" where \"token\"=$1", sel,
	)

	q := queries.Raw(query, token)

	err := q.Bind(ctx, exec, accessTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from access_tokens")
	}

	return accessTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccessToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no access_tokens provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(accessTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accessTokenInsertCacheMut.RLock()
	cache, cached := accessTokenInsertCache[key]
	accessTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accessTokenAllColumns,
			accessTokenColumnsWithDefault,
			accessTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accessTokenType, accessTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accessTokenType, accessTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"access_tokens\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"access_tokens\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into access_tokens")
	}

	if !cached {
		accessTokenInsertCacheMut.Lock()
		accessTokenInsertCache[key] = cache
		accessTokenInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AccessToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccessToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	accessTokenUpdateCacheMut.RLock()
	cache, cached := accessTokenUpdateCache[key]
	accessTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accessTokenAllColumns,
			accessTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update access_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"access_tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, accessTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accessTokenType, accessTokenMapping, append(wl, accessTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update access_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for access_tokens")
	}

	if !cached {
		accessTokenUpdateCacheMut.Lock()
		accessTokenUpdateCache[key] = cache
		accessTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q accessTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for access_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for access_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccessTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"access_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, accessTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in accessToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all accessToken")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccessToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no access_tokens provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(accessTokenColumnsWithDefault, o)

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

	accessTokenUpsertCacheMut.RLock()
	cache, cached := accessTokenUpsertCache[key]
	accessTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			accessTokenAllColumns,
			accessTokenColumnsWithDefault,
			accessTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			accessTokenAllColumns,
			accessTokenPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert access_tokens, could not build update column list")
		}

		ret := strmangle.SetComplement(accessTokenAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(accessTokenPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert access_tokens, could not build conflict column list")
			}

			conflict = make([]string, len(accessTokenPrimaryKeyColumns))
			copy(conflict, accessTokenPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"access_tokens\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(accessTokenType, accessTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accessTokenType, accessTokenMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert access_tokens")
	}

	if !cached {
		accessTokenUpsertCacheMut.Lock()
		accessTokenUpsertCache[key] = cache
		accessTokenUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AccessToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccessToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AccessToken provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accessTokenPrimaryKeyMapping)
	sql := "DELETE FROM \"access_tokens\" WHERE \"token\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from access_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for access_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accessTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accessTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from access_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for access_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccessTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"access_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accessTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from accessToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for access_tokens")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AccessToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccessToken(ctx, exec, o.Token)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccessTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccessTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accessTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"access_tokens\".* FROM \"access_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, accessTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccessTokenSlice")
	}

	*o = slice

	return nil
}

// AccessTokenExists checks if the AccessToken row exists.
func AccessTokenExists(ctx context.Context, exec boil.ContextExecutor, token string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"access_tokens\" where \"token\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, token)
	}
	row := exec.QueryRowContext(ctx, sql, token)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if access_tokens exists")
	}

	return exists, nil
}

// Exists checks if the AccessToken row exists.
func (o *AccessToken) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AccessTokenExists(ctx, exec, o.Token)
}
