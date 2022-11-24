// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/predicate"
	"github.com/Skijetler/alphinium/pkg/ent/user"
	"github.com/Skijetler/alphinium/pkg/ent/usermetadata"
)

// UserMetadataQuery is the builder for querying UserMetadata entities.
type UserMetadataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserMetadata
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserMetadataQuery builder.
func (umq *UserMetadataQuery) Where(ps ...predicate.UserMetadata) *UserMetadataQuery {
	umq.predicates = append(umq.predicates, ps...)
	return umq
}

// Limit adds a limit step to the query.
func (umq *UserMetadataQuery) Limit(limit int) *UserMetadataQuery {
	umq.limit = &limit
	return umq
}

// Offset adds an offset step to the query.
func (umq *UserMetadataQuery) Offset(offset int) *UserMetadataQuery {
	umq.offset = &offset
	return umq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (umq *UserMetadataQuery) Unique(unique bool) *UserMetadataQuery {
	umq.unique = &unique
	return umq
}

// Order adds an order step to the query.
func (umq *UserMetadataQuery) Order(o ...OrderFunc) *UserMetadataQuery {
	umq.order = append(umq.order, o...)
	return umq
}

// QueryUser chains the current query on the "user" edge.
func (umq *UserMetadataQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: umq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := umq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := umq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usermetadata.Table, usermetadata.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, usermetadata.UserTable, usermetadata.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(umq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserMetadata entity from the query.
// Returns a *NotFoundError when no UserMetadata was found.
func (umq *UserMetadataQuery) First(ctx context.Context) (*UserMetadata, error) {
	nodes, err := umq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usermetadata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (umq *UserMetadataQuery) FirstX(ctx context.Context) *UserMetadata {
	node, err := umq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserMetadata ID from the query.
// Returns a *NotFoundError when no UserMetadata ID was found.
func (umq *UserMetadataQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = umq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usermetadata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (umq *UserMetadataQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := umq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserMetadata entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserMetadata entity is found.
// Returns a *NotFoundError when no UserMetadata entities are found.
func (umq *UserMetadataQuery) Only(ctx context.Context) (*UserMetadata, error) {
	nodes, err := umq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usermetadata.Label}
	default:
		return nil, &NotSingularError{usermetadata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (umq *UserMetadataQuery) OnlyX(ctx context.Context) *UserMetadata {
	node, err := umq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserMetadata ID in the query.
// Returns a *NotSingularError when more than one UserMetadata ID is found.
// Returns a *NotFoundError when no entities are found.
func (umq *UserMetadataQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = umq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usermetadata.Label}
	default:
		err = &NotSingularError{usermetadata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (umq *UserMetadataQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := umq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserMetadataSlice.
func (umq *UserMetadataQuery) All(ctx context.Context) ([]*UserMetadata, error) {
	if err := umq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return umq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (umq *UserMetadataQuery) AllX(ctx context.Context) []*UserMetadata {
	nodes, err := umq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserMetadata IDs.
func (umq *UserMetadataQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := umq.Select(usermetadata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (umq *UserMetadataQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := umq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (umq *UserMetadataQuery) Count(ctx context.Context) (int, error) {
	if err := umq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return umq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (umq *UserMetadataQuery) CountX(ctx context.Context) int {
	count, err := umq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (umq *UserMetadataQuery) Exist(ctx context.Context) (bool, error) {
	if err := umq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return umq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (umq *UserMetadataQuery) ExistX(ctx context.Context) bool {
	exist, err := umq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserMetadataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (umq *UserMetadataQuery) Clone() *UserMetadataQuery {
	if umq == nil {
		return nil
	}
	return &UserMetadataQuery{
		config:     umq.config,
		limit:      umq.limit,
		offset:     umq.offset,
		order:      append([]OrderFunc{}, umq.order...),
		predicates: append([]predicate.UserMetadata{}, umq.predicates...),
		withUser:   umq.withUser.Clone(),
		// clone intermediate query.
		sql:    umq.sql.Clone(),
		path:   umq.path,
		unique: umq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (umq *UserMetadataQuery) WithUser(opts ...func(*UserQuery)) *UserMetadataQuery {
	query := &UserQuery{config: umq.config}
	for _, opt := range opts {
		opt(query)
	}
	umq.withUser = query
	return umq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NameColor string `json:"name_color,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserMetadata.Query().
//		GroupBy(usermetadata.FieldNameColor).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (umq *UserMetadataQuery) GroupBy(field string, fields ...string) *UserMetadataGroupBy {
	grbuild := &UserMetadataGroupBy{config: umq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := umq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return umq.sqlQuery(ctx), nil
	}
	grbuild.label = usermetadata.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NameColor string `json:"name_color,omitempty"`
//	}
//
//	client.UserMetadata.Query().
//		Select(usermetadata.FieldNameColor).
//		Scan(ctx, &v)
func (umq *UserMetadataQuery) Select(fields ...string) *UserMetadataSelect {
	umq.fields = append(umq.fields, fields...)
	selbuild := &UserMetadataSelect{UserMetadataQuery: umq}
	selbuild.label = usermetadata.Label
	selbuild.flds, selbuild.scan = &umq.fields, selbuild.Scan
	return selbuild
}

func (umq *UserMetadataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range umq.fields {
		if !usermetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if umq.path != nil {
		prev, err := umq.path(ctx)
		if err != nil {
			return err
		}
		umq.sql = prev
	}
	return nil
}

func (umq *UserMetadataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserMetadata, error) {
	var (
		nodes       = []*UserMetadata{}
		_spec       = umq.querySpec()
		loadedTypes = [1]bool{
			umq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserMetadata).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserMetadata{config: umq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, umq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := umq.withUser; query != nil {
		if err := umq.loadUser(ctx, query, nodes, nil,
			func(n *UserMetadata, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (umq *UserMetadataQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserMetadata, init func(*UserMetadata), assign func(*UserMetadata, *User)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*UserMetadata)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (umq *UserMetadataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := umq.querySpec()
	_spec.Node.Columns = umq.fields
	if len(umq.fields) > 0 {
		_spec.Unique = umq.unique != nil && *umq.unique
	}
	return sqlgraph.CountNodes(ctx, umq.driver, _spec)
}

func (umq *UserMetadataQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := umq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (umq *UserMetadataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usermetadata.Table,
			Columns: usermetadata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: usermetadata.FieldID,
			},
		},
		From:   umq.sql,
		Unique: true,
	}
	if unique := umq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := umq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usermetadata.FieldID)
		for i := range fields {
			if fields[i] != usermetadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := umq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := umq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := umq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := umq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (umq *UserMetadataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(umq.driver.Dialect())
	t1 := builder.Table(usermetadata.Table)
	columns := umq.fields
	if len(columns) == 0 {
		columns = usermetadata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if umq.sql != nil {
		selector = umq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if umq.unique != nil && *umq.unique {
		selector.Distinct()
	}
	for _, p := range umq.predicates {
		p(selector)
	}
	for _, p := range umq.order {
		p(selector)
	}
	if offset := umq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := umq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserMetadataGroupBy is the group-by builder for UserMetadata entities.
type UserMetadataGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (umgb *UserMetadataGroupBy) Aggregate(fns ...AggregateFunc) *UserMetadataGroupBy {
	umgb.fns = append(umgb.fns, fns...)
	return umgb
}

// Scan applies the group-by query and scans the result into the given value.
func (umgb *UserMetadataGroupBy) Scan(ctx context.Context, v any) error {
	query, err := umgb.path(ctx)
	if err != nil {
		return err
	}
	umgb.sql = query
	return umgb.sqlScan(ctx, v)
}

func (umgb *UserMetadataGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range umgb.fields {
		if !usermetadata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := umgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := umgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (umgb *UserMetadataGroupBy) sqlQuery() *sql.Selector {
	selector := umgb.sql.Select()
	aggregation := make([]string, 0, len(umgb.fns))
	for _, fn := range umgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(umgb.fields)+len(umgb.fns))
		for _, f := range umgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(umgb.fields...)...)
}

// UserMetadataSelect is the builder for selecting fields of UserMetadata entities.
type UserMetadataSelect struct {
	*UserMetadataQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ums *UserMetadataSelect) Scan(ctx context.Context, v any) error {
	if err := ums.prepareQuery(ctx); err != nil {
		return err
	}
	ums.sql = ums.UserMetadataQuery.sqlQuery(ctx)
	return ums.sqlScan(ctx, v)
}

func (ums *UserMetadataSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ums.sql.Query()
	if err := ums.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
