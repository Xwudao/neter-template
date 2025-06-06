// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go-kitboxpro/internal/data/ent/datalist"
	"go-kitboxpro/internal/data/ent/predicate"
)

// DataListQuery is the builder for querying DataList entities.
type DataListQuery struct {
	config
	ctx        *QueryContext
	order      []datalist.OrderOption
	inters     []Interceptor
	predicates []predicate.DataList
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DataListQuery builder.
func (dlq *DataListQuery) Where(ps ...predicate.DataList) *DataListQuery {
	dlq.predicates = append(dlq.predicates, ps...)
	return dlq
}

// Limit the number of records to be returned by this query.
func (dlq *DataListQuery) Limit(limit int) *DataListQuery {
	dlq.ctx.Limit = &limit
	return dlq
}

// Offset to start from.
func (dlq *DataListQuery) Offset(offset int) *DataListQuery {
	dlq.ctx.Offset = &offset
	return dlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dlq *DataListQuery) Unique(unique bool) *DataListQuery {
	dlq.ctx.Unique = &unique
	return dlq
}

// Order specifies how the records should be ordered.
func (dlq *DataListQuery) Order(o ...datalist.OrderOption) *DataListQuery {
	dlq.order = append(dlq.order, o...)
	return dlq
}

// First returns the first DataList entity from the query.
// Returns a *NotFoundError when no DataList was found.
func (dlq *DataListQuery) First(ctx context.Context) (*DataList, error) {
	nodes, err := dlq.Limit(1).All(setContextOp(ctx, dlq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{datalist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dlq *DataListQuery) FirstX(ctx context.Context) *DataList {
	node, err := dlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DataList ID from the query.
// Returns a *NotFoundError when no DataList ID was found.
func (dlq *DataListQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dlq.Limit(1).IDs(setContextOp(ctx, dlq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{datalist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dlq *DataListQuery) FirstIDX(ctx context.Context) int64 {
	id, err := dlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DataList entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DataList entity is found.
// Returns a *NotFoundError when no DataList entities are found.
func (dlq *DataListQuery) Only(ctx context.Context) (*DataList, error) {
	nodes, err := dlq.Limit(2).All(setContextOp(ctx, dlq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{datalist.Label}
	default:
		return nil, &NotSingularError{datalist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dlq *DataListQuery) OnlyX(ctx context.Context) *DataList {
	node, err := dlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DataList ID in the query.
// Returns a *NotSingularError when more than one DataList ID is found.
// Returns a *NotFoundError when no entities are found.
func (dlq *DataListQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dlq.Limit(2).IDs(setContextOp(ctx, dlq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{datalist.Label}
	default:
		err = &NotSingularError{datalist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dlq *DataListQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := dlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DataLists.
func (dlq *DataListQuery) All(ctx context.Context) ([]*DataList, error) {
	ctx = setContextOp(ctx, dlq.ctx, ent.OpQueryAll)
	if err := dlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DataList, *DataListQuery]()
	return withInterceptors[[]*DataList](ctx, dlq, qr, dlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dlq *DataListQuery) AllX(ctx context.Context) []*DataList {
	nodes, err := dlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DataList IDs.
func (dlq *DataListQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if dlq.ctx.Unique == nil && dlq.path != nil {
		dlq.Unique(true)
	}
	ctx = setContextOp(ctx, dlq.ctx, ent.OpQueryIDs)
	if err = dlq.Select(datalist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dlq *DataListQuery) IDsX(ctx context.Context) []int64 {
	ids, err := dlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dlq *DataListQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dlq.ctx, ent.OpQueryCount)
	if err := dlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dlq, querierCount[*DataListQuery](), dlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dlq *DataListQuery) CountX(ctx context.Context) int {
	count, err := dlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dlq *DataListQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dlq.ctx, ent.OpQueryExist)
	switch _, err := dlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dlq *DataListQuery) ExistX(ctx context.Context) bool {
	exist, err := dlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DataListQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dlq *DataListQuery) Clone() *DataListQuery {
	if dlq == nil {
		return nil
	}
	return &DataListQuery{
		config:     dlq.config,
		ctx:        dlq.ctx.Clone(),
		order:      append([]datalist.OrderOption{}, dlq.order...),
		inters:     append([]Interceptor{}, dlq.inters...),
		predicates: append([]predicate.DataList{}, dlq.predicates...),
		// clone intermediate query.
		sql:  dlq.sql.Clone(),
		path: dlq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DataList.Query().
//		GroupBy(datalist.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dlq *DataListQuery) GroupBy(field string, fields ...string) *DataListGroupBy {
	dlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DataListGroupBy{build: dlq}
	grbuild.flds = &dlq.ctx.Fields
	grbuild.label = datalist.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.DataList.Query().
//		Select(datalist.FieldCreateTime).
//		Scan(ctx, &v)
func (dlq *DataListQuery) Select(fields ...string) *DataListSelect {
	dlq.ctx.Fields = append(dlq.ctx.Fields, fields...)
	sbuild := &DataListSelect{DataListQuery: dlq}
	sbuild.label = datalist.Label
	sbuild.flds, sbuild.scan = &dlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DataListSelect configured with the given aggregations.
func (dlq *DataListQuery) Aggregate(fns ...AggregateFunc) *DataListSelect {
	return dlq.Select().Aggregate(fns...)
}

func (dlq *DataListQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dlq); err != nil {
				return err
			}
		}
	}
	for _, f := range dlq.ctx.Fields {
		if !datalist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dlq.path != nil {
		prev, err := dlq.path(ctx)
		if err != nil {
			return err
		}
		dlq.sql = prev
	}
	return nil
}

func (dlq *DataListQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DataList, error) {
	var (
		nodes = []*DataList{}
		_spec = dlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DataList).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DataList{config: dlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dlq *DataListQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dlq.querySpec()
	_spec.Node.Columns = dlq.ctx.Fields
	if len(dlq.ctx.Fields) > 0 {
		_spec.Unique = dlq.ctx.Unique != nil && *dlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dlq.driver, _spec)
}

func (dlq *DataListQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(datalist.Table, datalist.Columns, sqlgraph.NewFieldSpec(datalist.FieldID, field.TypeInt64))
	_spec.From = dlq.sql
	if unique := dlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dlq.path != nil {
		_spec.Unique = true
	}
	if fields := dlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, datalist.FieldID)
		for i := range fields {
			if fields[i] != datalist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dlq *DataListQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dlq.driver.Dialect())
	t1 := builder.Table(datalist.Table)
	columns := dlq.ctx.Fields
	if len(columns) == 0 {
		columns = datalist.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dlq.sql != nil {
		selector = dlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dlq.ctx.Unique != nil && *dlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dlq.predicates {
		p(selector)
	}
	for _, p := range dlq.order {
		p(selector)
	}
	if offset := dlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DataListGroupBy is the group-by builder for DataList entities.
type DataListGroupBy struct {
	selector
	build *DataListQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dlgb *DataListGroupBy) Aggregate(fns ...AggregateFunc) *DataListGroupBy {
	dlgb.fns = append(dlgb.fns, fns...)
	return dlgb
}

// Scan applies the selector query and scans the result into the given value.
func (dlgb *DataListGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dlgb.build.ctx, ent.OpQueryGroupBy)
	if err := dlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DataListQuery, *DataListGroupBy](ctx, dlgb.build, dlgb, dlgb.build.inters, v)
}

func (dlgb *DataListGroupBy) sqlScan(ctx context.Context, root *DataListQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dlgb.fns))
	for _, fn := range dlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dlgb.flds)+len(dlgb.fns))
		for _, f := range *dlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DataListSelect is the builder for selecting fields of DataList entities.
type DataListSelect struct {
	*DataListQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dls *DataListSelect) Aggregate(fns ...AggregateFunc) *DataListSelect {
	dls.fns = append(dls.fns, fns...)
	return dls
}

// Scan applies the selector query and scans the result into the given value.
func (dls *DataListSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dls.ctx, ent.OpQuerySelect)
	if err := dls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DataListQuery, *DataListSelect](ctx, dls.DataListQuery, dls, dls.inters, v)
}

func (dls *DataListSelect) sqlScan(ctx context.Context, root *DataListQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dls.fns))
	for _, fn := range dls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
