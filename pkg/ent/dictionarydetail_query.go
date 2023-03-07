// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// DictionaryDetailQuery is the builder for querying DictionaryDetail entities.
type DictionaryDetailQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.DictionaryDetail
	withDictionaries *DictionaryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictionaryDetailQuery builder.
func (ddq *DictionaryDetailQuery) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailQuery {
	ddq.predicates = append(ddq.predicates, ps...)
	return ddq
}

// Limit the number of records to be returned by this query.
func (ddq *DictionaryDetailQuery) Limit(limit int) *DictionaryDetailQuery {
	ddq.ctx.Limit = &limit
	return ddq
}

// Offset to start from.
func (ddq *DictionaryDetailQuery) Offset(offset int) *DictionaryDetailQuery {
	ddq.ctx.Offset = &offset
	return ddq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ddq *DictionaryDetailQuery) Unique(unique bool) *DictionaryDetailQuery {
	ddq.ctx.Unique = &unique
	return ddq
}

// Order specifies how the records should be ordered.
func (ddq *DictionaryDetailQuery) Order(o ...OrderFunc) *DictionaryDetailQuery {
	ddq.order = append(ddq.order, o...)
	return ddq
}

// QueryDictionaries chains the current query on the "dictionaries" edge.
func (ddq *DictionaryDetailQuery) QueryDictionaries() *DictionaryQuery {
	query := (&DictionaryClient{config: ddq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ddq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ddq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dictionarydetail.Table, dictionarydetail.FieldID, selector),
			sqlgraph.To(dictionary.Table, dictionary.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dictionarydetail.DictionariesTable, dictionarydetail.DictionariesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ddq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DictionaryDetail entity from the query.
// Returns a *NotFoundError when no DictionaryDetail was found.
func (ddq *DictionaryDetailQuery) First(ctx context.Context) (*DictionaryDetail, error) {
	nodes, err := ddq.Limit(1).All(setContextOp(ctx, ddq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dictionarydetail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) FirstX(ctx context.Context) *DictionaryDetail {
	node, err := ddq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DictionaryDetail ID from the query.
// Returns a *NotFoundError when no DictionaryDetail ID was found.
func (ddq *DictionaryDetailQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ddq.Limit(1).IDs(setContextOp(ctx, ddq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dictionarydetail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ddq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DictionaryDetail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DictionaryDetail entity is found.
// Returns a *NotFoundError when no DictionaryDetail entities are found.
func (ddq *DictionaryDetailQuery) Only(ctx context.Context) (*DictionaryDetail, error) {
	nodes, err := ddq.Limit(2).All(setContextOp(ctx, ddq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dictionarydetail.Label}
	default:
		return nil, &NotSingularError{dictionarydetail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) OnlyX(ctx context.Context) *DictionaryDetail {
	node, err := ddq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DictionaryDetail ID in the query.
// Returns a *NotSingularError when more than one DictionaryDetail ID is found.
// Returns a *NotFoundError when no entities are found.
func (ddq *DictionaryDetailQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ddq.Limit(2).IDs(setContextOp(ctx, ddq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dictionarydetail.Label}
	default:
		err = &NotSingularError{dictionarydetail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ddq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DictionaryDetails.
func (ddq *DictionaryDetailQuery) All(ctx context.Context) ([]*DictionaryDetail, error) {
	ctx = setContextOp(ctx, ddq.ctx, "All")
	if err := ddq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DictionaryDetail, *DictionaryDetailQuery]()
	return withInterceptors[[]*DictionaryDetail](ctx, ddq, qr, ddq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) AllX(ctx context.Context) []*DictionaryDetail {
	nodes, err := ddq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DictionaryDetail IDs.
func (ddq *DictionaryDetailQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if ddq.ctx.Unique == nil && ddq.path != nil {
		ddq.Unique(true)
	}
	ctx = setContextOp(ctx, ddq.ctx, "IDs")
	if err = ddq.Select(dictionarydetail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ddq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ddq *DictionaryDetailQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ddq.ctx, "Count")
	if err := ddq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ddq, querierCount[*DictionaryDetailQuery](), ddq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) CountX(ctx context.Context) int {
	count, err := ddq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ddq *DictionaryDetailQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ddq.ctx, "Exist")
	switch _, err := ddq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ddq *DictionaryDetailQuery) ExistX(ctx context.Context) bool {
	exist, err := ddq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictionaryDetailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ddq *DictionaryDetailQuery) Clone() *DictionaryDetailQuery {
	if ddq == nil {
		return nil
	}
	return &DictionaryDetailQuery{
		config:           ddq.config,
		ctx:              ddq.ctx.Clone(),
		order:            append([]OrderFunc{}, ddq.order...),
		inters:           append([]Interceptor{}, ddq.inters...),
		predicates:       append([]predicate.DictionaryDetail{}, ddq.predicates...),
		withDictionaries: ddq.withDictionaries.Clone(),
		// clone intermediate query.
		sql:  ddq.sql.Clone(),
		path: ddq.path,
	}
}

// WithDictionaries tells the query-builder to eager-load the nodes that are connected to
// the "dictionaries" edge. The optional arguments are used to configure the query builder of the edge.
func (ddq *DictionaryDetailQuery) WithDictionaries(opts ...func(*DictionaryQuery)) *DictionaryDetailQuery {
	query := (&DictionaryClient{config: ddq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ddq.withDictionaries = query
	return ddq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DictionaryDetail.Query().
//		GroupBy(dictionarydetail.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ddq *DictionaryDetailQuery) GroupBy(field string, fields ...string) *DictionaryDetailGroupBy {
	ddq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DictionaryDetailGroupBy{build: ddq}
	grbuild.flds = &ddq.ctx.Fields
	grbuild.label = dictionarydetail.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DictionaryDetail.Query().
//		Select(dictionarydetail.FieldCreatedAt).
//		Scan(ctx, &v)
func (ddq *DictionaryDetailQuery) Select(fields ...string) *DictionaryDetailSelect {
	ddq.ctx.Fields = append(ddq.ctx.Fields, fields...)
	sbuild := &DictionaryDetailSelect{DictionaryDetailQuery: ddq}
	sbuild.label = dictionarydetail.Label
	sbuild.flds, sbuild.scan = &ddq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DictionaryDetailSelect configured with the given aggregations.
func (ddq *DictionaryDetailQuery) Aggregate(fns ...AggregateFunc) *DictionaryDetailSelect {
	return ddq.Select().Aggregate(fns...)
}

func (ddq *DictionaryDetailQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ddq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ddq); err != nil {
				return err
			}
		}
	}
	for _, f := range ddq.ctx.Fields {
		if !dictionarydetail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ddq.path != nil {
		prev, err := ddq.path(ctx)
		if err != nil {
			return err
		}
		ddq.sql = prev
	}
	return nil
}

func (ddq *DictionaryDetailQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DictionaryDetail, error) {
	var (
		nodes       = []*DictionaryDetail{}
		_spec       = ddq.querySpec()
		loadedTypes = [1]bool{
			ddq.withDictionaries != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DictionaryDetail).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DictionaryDetail{config: ddq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ddq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ddq.withDictionaries; query != nil {
		if err := ddq.loadDictionaries(ctx, query, nodes, nil,
			func(n *DictionaryDetail, e *Dictionary) { n.Edges.Dictionaries = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ddq *DictionaryDetailQuery) loadDictionaries(ctx context.Context, query *DictionaryQuery, nodes []*DictionaryDetail, init func(*DictionaryDetail), assign func(*DictionaryDetail, *Dictionary)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DictionaryDetail)
	for i := range nodes {
		fk := nodes[i].DictionaryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dictionary.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dictionary_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ddq *DictionaryDetailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ddq.querySpec()
	_spec.Node.Columns = ddq.ctx.Fields
	if len(ddq.ctx.Fields) > 0 {
		_spec.Unique = ddq.ctx.Unique != nil && *ddq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ddq.driver, _spec)
}

func (ddq *DictionaryDetailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dictionarydetail.Table, dictionarydetail.Columns, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64))
	_spec.From = ddq.sql
	if unique := ddq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ddq.path != nil {
		_spec.Unique = true
	}
	if fields := ddq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionarydetail.FieldID)
		for i := range fields {
			if fields[i] != dictionarydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ddq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ddq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ddq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ddq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ddq *DictionaryDetailQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ddq.driver.Dialect())
	t1 := builder.Table(dictionarydetail.Table)
	columns := ddq.ctx.Fields
	if len(columns) == 0 {
		columns = dictionarydetail.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ddq.sql != nil {
		selector = ddq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ddq.ctx.Unique != nil && *ddq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ddq.predicates {
		p(selector)
	}
	for _, p := range ddq.order {
		p(selector)
	}
	if offset := ddq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ddq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DictionaryDetailGroupBy is the group-by builder for DictionaryDetail entities.
type DictionaryDetailGroupBy struct {
	selector
	build *DictionaryDetailQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ddgb *DictionaryDetailGroupBy) Aggregate(fns ...AggregateFunc) *DictionaryDetailGroupBy {
	ddgb.fns = append(ddgb.fns, fns...)
	return ddgb
}

// Scan applies the selector query and scans the result into the given value.
func (ddgb *DictionaryDetailGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ddgb.build.ctx, "GroupBy")
	if err := ddgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictionaryDetailQuery, *DictionaryDetailGroupBy](ctx, ddgb.build, ddgb, ddgb.build.inters, v)
}

func (ddgb *DictionaryDetailGroupBy) sqlScan(ctx context.Context, root *DictionaryDetailQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ddgb.fns))
	for _, fn := range ddgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ddgb.flds)+len(ddgb.fns))
		for _, f := range *ddgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ddgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ddgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DictionaryDetailSelect is the builder for selecting fields of DictionaryDetail entities.
type DictionaryDetailSelect struct {
	*DictionaryDetailQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dds *DictionaryDetailSelect) Aggregate(fns ...AggregateFunc) *DictionaryDetailSelect {
	dds.fns = append(dds.fns, fns...)
	return dds
}

// Scan applies the selector query and scans the result into the given value.
func (dds *DictionaryDetailSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dds.ctx, "Select")
	if err := dds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictionaryDetailQuery, *DictionaryDetailSelect](ctx, dds.DictionaryDetailQuery, dds, dds.inters, v)
}

func (dds *DictionaryDetailSelect) sqlScan(ctx context.Context, root *DictionaryDetailQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dds.fns))
	for _, fn := range dds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
