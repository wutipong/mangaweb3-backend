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
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
	"github.com/wutipong/mangaweb3-backend/ent/user"
)

// HistoryQuery is the builder for querying History entities.
type HistoryQuery struct {
	config
	ctx        *QueryContext
	order      []history.OrderOption
	inters     []Interceptor
	predicates []predicate.History
	withItem   *MetaQuery
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HistoryQuery builder.
func (hq *HistoryQuery) Where(ps ...predicate.History) *HistoryQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit the number of records to be returned by this query.
func (hq *HistoryQuery) Limit(limit int) *HistoryQuery {
	hq.ctx.Limit = &limit
	return hq
}

// Offset to start from.
func (hq *HistoryQuery) Offset(offset int) *HistoryQuery {
	hq.ctx.Offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HistoryQuery) Unique(unique bool) *HistoryQuery {
	hq.ctx.Unique = &unique
	return hq
}

// Order specifies how the records should be ordered.
func (hq *HistoryQuery) Order(o ...history.OrderOption) *HistoryQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryItem chains the current query on the "item" edge.
func (hq *HistoryQuery) QueryItem() *MetaQuery {
	query := (&MetaClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(history.Table, history.FieldID, selector),
			sqlgraph.To(meta.Table, meta.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, history.ItemTable, history.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (hq *HistoryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(history.Table, history.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, history.UserTable, history.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first History entity from the query.
// Returns a *NotFoundError when no History was found.
func (hq *HistoryQuery) First(ctx context.Context) (*History, error) {
	nodes, err := hq.Limit(1).All(setContextOp(ctx, hq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{history.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HistoryQuery) FirstX(ctx context.Context) *History {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first History ID from the query.
// Returns a *NotFoundError when no History ID was found.
func (hq *HistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(setContextOp(ctx, hq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{history.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single History entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one History entity is found.
// Returns a *NotFoundError when no History entities are found.
func (hq *HistoryQuery) Only(ctx context.Context) (*History, error) {
	nodes, err := hq.Limit(2).All(setContextOp(ctx, hq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{history.Label}
	default:
		return nil, &NotSingularError{history.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HistoryQuery) OnlyX(ctx context.Context) *History {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only History ID in the query.
// Returns a *NotSingularError when more than one History ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(setContextOp(ctx, hq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{history.Label}
	default:
		err = &NotSingularError{history.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Histories.
func (hq *HistoryQuery) All(ctx context.Context) ([]*History, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryAll)
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*History, *HistoryQuery]()
	return withInterceptors[[]*History](ctx, hq, qr, hq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hq *HistoryQuery) AllX(ctx context.Context) []*History {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of History IDs.
func (hq *HistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hq.ctx.Unique == nil && hq.path != nil {
		hq.Unique(true)
	}
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryIDs)
	if err = hq.Select(history.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryCount)
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hq, querierCount[*HistoryQuery](), hq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HistoryQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hq.ctx, ent.OpQueryExist)
	switch _, err := hq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HistoryQuery) Clone() *HistoryQuery {
	if hq == nil {
		return nil
	}
	return &HistoryQuery{
		config:     hq.config,
		ctx:        hq.ctx.Clone(),
		order:      append([]history.OrderOption{}, hq.order...),
		inters:     append([]Interceptor{}, hq.inters...),
		predicates: append([]predicate.History{}, hq.predicates...),
		withItem:   hq.withItem.Clone(),
		withUser:   hq.withUser.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HistoryQuery) WithItem(opts ...func(*MetaQuery)) *HistoryQuery {
	query := (&MetaClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withItem = query
	return hq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HistoryQuery) WithUser(opts ...func(*UserQuery)) *HistoryQuery {
	query := (&UserClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withUser = query
	return hq
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
//	client.History.Query().
//		GroupBy(history.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hq *HistoryQuery) GroupBy(field string, fields ...string) *HistoryGroupBy {
	hq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HistoryGroupBy{build: hq}
	grbuild.flds = &hq.ctx.Fields
	grbuild.label = history.Label
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
//	client.History.Query().
//		Select(history.FieldCreateTime).
//		Scan(ctx, &v)
func (hq *HistoryQuery) Select(fields ...string) *HistorySelect {
	hq.ctx.Fields = append(hq.ctx.Fields, fields...)
	sbuild := &HistorySelect{HistoryQuery: hq}
	sbuild.label = history.Label
	sbuild.flds, sbuild.scan = &hq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HistorySelect configured with the given aggregations.
func (hq *HistoryQuery) Aggregate(fns ...AggregateFunc) *HistorySelect {
	return hq.Select().Aggregate(fns...)
}

func (hq *HistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hq); err != nil {
				return err
			}
		}
	}
	for _, f := range hq.ctx.Fields {
		if !history.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*History, error) {
	var (
		nodes       = []*History{}
		withFKs     = hq.withFKs
		_spec       = hq.querySpec()
		loadedTypes = [2]bool{
			hq.withItem != nil,
			hq.withUser != nil,
		}
	)
	if hq.withItem != nil || hq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, history.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*History).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &History{config: hq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hq.withItem; query != nil {
		if err := hq.loadItem(ctx, query, nodes, nil,
			func(n *History, e *Meta) { n.Edges.Item = e }); err != nil {
			return nil, err
		}
	}
	if query := hq.withUser; query != nil {
		if err := hq.loadUser(ctx, query, nodes, nil,
			func(n *History, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hq *HistoryQuery) loadItem(ctx context.Context, query *MetaQuery, nodes []*History, init func(*History), assign func(*History, *Meta)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*History)
	for i := range nodes {
		if nodes[i].meta_histories == nil {
			continue
		}
		fk := *nodes[i].meta_histories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(meta.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "meta_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hq *HistoryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*History, init func(*History), assign func(*History, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*History)
	for i := range nodes {
		if nodes[i].user_histories == nil {
			continue
		}
		fk := *nodes[i].user_histories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hq *HistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	_spec.Node.Columns = hq.ctx.Fields
	if len(hq.ctx.Fields) > 0 {
		_spec.Unique = hq.ctx.Unique != nil && *hq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	_spec.From = hq.sql
	if unique := hq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hq.path != nil {
		_spec.Unique = true
	}
	if fields := hq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for i := range fields {
			if fields[i] != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(history.Table)
	columns := hq.ctx.Fields
	if len(columns) == 0 {
		columns = history.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.ctx.Unique != nil && *hq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HistoryGroupBy is the group-by builder for History entities.
type HistoryGroupBy struct {
	selector
	build *HistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HistoryGroupBy) Aggregate(fns ...AggregateFunc) *HistoryGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the selector query and scans the result into the given value.
func (hgb *HistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hgb.build.ctx, ent.OpQueryGroupBy)
	if err := hgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HistoryQuery, *HistoryGroupBy](ctx, hgb.build, hgb, hgb.build.inters, v)
}

func (hgb *HistoryGroupBy) sqlScan(ctx context.Context, root *HistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hgb.flds)+len(hgb.fns))
		for _, f := range *hgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HistorySelect is the builder for selecting fields of History entities.
type HistorySelect struct {
	*HistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hs *HistorySelect) Aggregate(fns ...AggregateFunc) *HistorySelect {
	hs.fns = append(hs.fns, fns...)
	return hs
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hs.ctx, ent.OpQuerySelect)
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HistoryQuery, *HistorySelect](ctx, hs.HistoryQuery, hs, hs.inters, v)
}

func (hs *HistorySelect) sqlScan(ctx context.Context, root *HistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hs.fns))
	for _, fn := range hs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
