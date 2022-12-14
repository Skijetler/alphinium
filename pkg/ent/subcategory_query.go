// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/category"
	"github.com/Skijetler/alphinium/pkg/ent/predicate"
	"github.com/Skijetler/alphinium/pkg/ent/subcategory"
	"github.com/Skijetler/alphinium/pkg/ent/thread"
)

// SubcategoryQuery is the builder for querying Subcategory entities.
type SubcategoryQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Subcategory
	withCategory *CategoryQuery
	withThreads  *ThreadQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubcategoryQuery builder.
func (sq *SubcategoryQuery) Where(ps ...predicate.Subcategory) *SubcategoryQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SubcategoryQuery) Limit(limit int) *SubcategoryQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SubcategoryQuery) Offset(offset int) *SubcategoryQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubcategoryQuery) Unique(unique bool) *SubcategoryQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SubcategoryQuery) Order(o ...OrderFunc) *SubcategoryQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCategory chains the current query on the "category" edge.
func (sq *SubcategoryQuery) QueryCategory() *CategoryQuery {
	query := &CategoryQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subcategory.CategoryTable, subcategory.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThreads chains the current query on the "threads" edge.
func (sq *SubcategoryQuery) QueryThreads() *ThreadQuery {
	query := &ThreadQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subcategory.ThreadsTable, subcategory.ThreadsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Subcategory entity from the query.
// Returns a *NotFoundError when no Subcategory was found.
func (sq *SubcategoryQuery) First(ctx context.Context) (*Subcategory, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubcategoryQuery) FirstX(ctx context.Context) *Subcategory {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Subcategory ID from the query.
// Returns a *NotFoundError when no Subcategory ID was found.
func (sq *SubcategoryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubcategoryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Subcategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Subcategory entity is found.
// Returns a *NotFoundError when no Subcategory entities are found.
func (sq *SubcategoryQuery) Only(ctx context.Context) (*Subcategory, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subcategory.Label}
	default:
		return nil, &NotSingularError{subcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubcategoryQuery) OnlyX(ctx context.Context) *Subcategory {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Subcategory ID in the query.
// Returns a *NotSingularError when more than one Subcategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubcategoryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subcategory.Label}
	default:
		err = &NotSingularError{subcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubcategoryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subcategories.
func (sq *SubcategoryQuery) All(ctx context.Context) ([]*Subcategory, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubcategoryQuery) AllX(ctx context.Context) []*Subcategory {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Subcategory IDs.
func (sq *SubcategoryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := sq.Select(subcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubcategoryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubcategoryQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubcategoryQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubcategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubcategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubcategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubcategoryQuery) Clone() *SubcategoryQuery {
	if sq == nil {
		return nil
	}
	return &SubcategoryQuery{
		config:       sq.config,
		limit:        sq.limit,
		offset:       sq.offset,
		order:        append([]OrderFunc{}, sq.order...),
		predicates:   append([]predicate.Subcategory{}, sq.predicates...),
		withCategory: sq.withCategory.Clone(),
		withThreads:  sq.withThreads.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubcategoryQuery) WithCategory(opts ...func(*CategoryQuery)) *SubcategoryQuery {
	query := &CategoryQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withCategory = query
	return sq
}

// WithThreads tells the query-builder to eager-load the nodes that are connected to
// the "threads" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubcategoryQuery) WithThreads(opts ...func(*ThreadQuery)) *SubcategoryQuery {
	query := &ThreadQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withThreads = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Subcategory.Query().
//		GroupBy(subcategory.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubcategoryQuery) GroupBy(field string, fields ...string) *SubcategoryGroupBy {
	grbuild := &SubcategoryGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = subcategory.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Subcategory.Query().
//		Select(subcategory.FieldName).
//		Scan(ctx, &v)
func (sq *SubcategoryQuery) Select(fields ...string) *SubcategorySelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SubcategorySelect{SubcategoryQuery: sq}
	selbuild.label = subcategory.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

func (sq *SubcategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !subcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubcategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Subcategory, error) {
	var (
		nodes       = []*Subcategory{}
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withCategory != nil,
			sq.withThreads != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Subcategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Subcategory{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCategory; query != nil {
		if err := sq.loadCategory(ctx, query, nodes, nil,
			func(n *Subcategory, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withThreads; query != nil {
		if err := sq.loadThreads(ctx, query, nodes,
			func(n *Subcategory) { n.Edges.Threads = []*Thread{} },
			func(n *Subcategory, e *Thread) { n.Edges.Threads = append(n.Edges.Threads, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubcategoryQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*Subcategory, init func(*Subcategory), assign func(*Subcategory, *Category)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Subcategory)
	for i := range nodes {
		fk := nodes[i].CategoryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SubcategoryQuery) loadThreads(ctx context.Context, query *ThreadQuery, nodes []*Subcategory, init func(*Subcategory), assign func(*Subcategory, *Thread)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Subcategory)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Thread(func(s *sql.Selector) {
		s.Where(sql.InValues(subcategory.ThreadsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SubcategoryID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "subcategory_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SubcategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubcategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SubcategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subcategory.Table,
			Columns: subcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: subcategory.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subcategory.FieldID)
		for i := range fields {
			if fields[i] != subcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubcategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subcategory.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = subcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubcategoryGroupBy is the group-by builder for Subcategory entities.
type SubcategoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubcategoryGroupBy) Aggregate(fns ...AggregateFunc) *SubcategoryGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SubcategoryGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SubcategoryGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !subcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SubcategoryGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SubcategorySelect is the builder for selecting fields of Subcategory entities.
type SubcategorySelect struct {
	*SubcategoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubcategorySelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SubcategoryQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SubcategorySelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
