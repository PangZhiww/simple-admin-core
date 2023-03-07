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
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// DictionaryQuery is the builder for querying Dictionary entities.
type DictionaryQuery struct {
	config
	ctx                   *QueryContext
	order                 []OrderFunc
	inters                []Interceptor
	predicates            []predicate.Dictionary
	withDictionaryDetails *DictionaryDetailQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictionaryQuery builder.
func (dq *DictionaryQuery) Where(ps ...predicate.Dictionary) *DictionaryQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DictionaryQuery) Limit(limit int) *DictionaryQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DictionaryQuery) Offset(offset int) *DictionaryQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DictionaryQuery) Unique(unique bool) *DictionaryQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DictionaryQuery) Order(o ...OrderFunc) *DictionaryQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDictionaryDetails chains the current query on the "dictionary_details" edge.
func (dq *DictionaryQuery) QueryDictionaryDetails() *DictionaryDetailQuery {
	query := (&DictionaryDetailClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dictionary.Table, dictionary.FieldID, selector),
			sqlgraph.To(dictionarydetail.Table, dictionarydetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dictionary.DictionaryDetailsTable, dictionary.DictionaryDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dictionary entity from the query.
// Returns a *NotFoundError when no Dictionary was found.
func (dq *DictionaryQuery) First(ctx context.Context) (*Dictionary, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dictionary.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DictionaryQuery) FirstX(ctx context.Context) *Dictionary {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dictionary ID from the query.
// Returns a *NotFoundError when no Dictionary ID was found.
func (dq *DictionaryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dictionary.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DictionaryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dictionary entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dictionary entity is found.
// Returns a *NotFoundError when no Dictionary entities are found.
func (dq *DictionaryQuery) Only(ctx context.Context) (*Dictionary, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dictionary.Label}
	default:
		return nil, &NotSingularError{dictionary.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DictionaryQuery) OnlyX(ctx context.Context) *Dictionary {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dictionary ID in the query.
// Returns a *NotSingularError when more than one Dictionary ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DictionaryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dictionary.Label}
	default:
		err = &NotSingularError{dictionary.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DictionaryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dictionaries.
func (dq *DictionaryQuery) All(ctx context.Context) ([]*Dictionary, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Dictionary, *DictionaryQuery]()
	return withInterceptors[[]*Dictionary](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DictionaryQuery) AllX(ctx context.Context) []*Dictionary {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dictionary IDs.
func (dq *DictionaryQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(dictionary.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DictionaryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DictionaryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DictionaryQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DictionaryQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DictionaryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DictionaryQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictionaryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DictionaryQuery) Clone() *DictionaryQuery {
	if dq == nil {
		return nil
	}
	return &DictionaryQuery{
		config:                dq.config,
		ctx:                   dq.ctx.Clone(),
		order:                 append([]OrderFunc{}, dq.order...),
		inters:                append([]Interceptor{}, dq.inters...),
		predicates:            append([]predicate.Dictionary{}, dq.predicates...),
		withDictionaryDetails: dq.withDictionaryDetails.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDictionaryDetails tells the query-builder to eager-load the nodes that are connected to
// the "dictionary_details" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DictionaryQuery) WithDictionaryDetails(opts ...func(*DictionaryDetailQuery)) *DictionaryQuery {
	query := (&DictionaryDetailClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDictionaryDetails = query
	return dq
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
//	client.Dictionary.Query().
//		GroupBy(dictionary.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DictionaryQuery) GroupBy(field string, fields ...string) *DictionaryGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DictionaryGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = dictionary.Label
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
//	client.Dictionary.Query().
//		Select(dictionary.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DictionaryQuery) Select(fields ...string) *DictionarySelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DictionarySelect{DictionaryQuery: dq}
	sbuild.label = dictionary.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DictionarySelect configured with the given aggregations.
func (dq *DictionaryQuery) Aggregate(fns ...AggregateFunc) *DictionarySelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DictionaryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !dictionary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DictionaryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dictionary, error) {
	var (
		nodes       = []*Dictionary{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDictionaryDetails != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Dictionary).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Dictionary{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDictionaryDetails; query != nil {
		if err := dq.loadDictionaryDetails(ctx, query, nodes,
			func(n *Dictionary) { n.Edges.DictionaryDetails = []*DictionaryDetail{} },
			func(n *Dictionary, e *DictionaryDetail) {
				n.Edges.DictionaryDetails = append(n.Edges.DictionaryDetails, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DictionaryQuery) loadDictionaryDetails(ctx context.Context, query *DictionaryDetailQuery, nodes []*Dictionary, init func(*Dictionary), assign func(*Dictionary, *DictionaryDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Dictionary)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.DictionaryDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(dictionary.DictionaryDetailsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DictionaryID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dictionary_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DictionaryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DictionaryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dictionary.Table, dictionary.Columns, sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionary.FieldID)
		for i := range fields {
			if fields[i] != dictionary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DictionaryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dictionary.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = dictionary.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DictionaryGroupBy is the group-by builder for Dictionary entities.
type DictionaryGroupBy struct {
	selector
	build *DictionaryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DictionaryGroupBy) Aggregate(fns ...AggregateFunc) *DictionaryGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DictionaryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictionaryQuery, *DictionaryGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DictionaryGroupBy) sqlScan(ctx context.Context, root *DictionaryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DictionarySelect is the builder for selecting fields of Dictionary entities.
type DictionarySelect struct {
	*DictionaryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DictionarySelect) Aggregate(fns ...AggregateFunc) *DictionarySelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DictionarySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictionaryQuery, *DictionarySelect](ctx, ds.DictionaryQuery, ds, ds.inters, v)
}

func (ds *DictionarySelect) sqlScan(ctx context.Context, root *DictionaryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
