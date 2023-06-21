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
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MunicipalityQuery is the builder for querying Municipality entities.
type MunicipalityQuery struct {
	config
	ctx             *QueryContext
	order           []municipality.OrderOption
	inters          []Interceptor
	predicates      []predicate.Municipality
	withMeshis      *MeshiQuery
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*Municipality) error
	withNamedMeshis map[string]*MeshiQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MunicipalityQuery builder.
func (mq *MunicipalityQuery) Where(ps ...predicate.Municipality) *MunicipalityQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MunicipalityQuery) Limit(limit int) *MunicipalityQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MunicipalityQuery) Offset(offset int) *MunicipalityQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MunicipalityQuery) Unique(unique bool) *MunicipalityQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MunicipalityQuery) Order(o ...municipality.OrderOption) *MunicipalityQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryMeshis chains the current query on the "meshis" edge.
func (mq *MunicipalityQuery) QueryMeshis() *MeshiQuery {
	query := (&MeshiClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(municipality.Table, municipality.FieldID, selector),
			sqlgraph.To(meshi.Table, meshi.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, municipality.MeshisTable, municipality.MeshisColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Municipality entity from the query.
// Returns a *NotFoundError when no Municipality was found.
func (mq *MunicipalityQuery) First(ctx context.Context) (*Municipality, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{municipality.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MunicipalityQuery) FirstX(ctx context.Context) *Municipality {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Municipality ID from the query.
// Returns a *NotFoundError when no Municipality ID was found.
func (mq *MunicipalityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{municipality.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MunicipalityQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Municipality entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Municipality entity is found.
// Returns a *NotFoundError when no Municipality entities are found.
func (mq *MunicipalityQuery) Only(ctx context.Context) (*Municipality, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{municipality.Label}
	default:
		return nil, &NotSingularError{municipality.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MunicipalityQuery) OnlyX(ctx context.Context) *Municipality {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Municipality ID in the query.
// Returns a *NotSingularError when more than one Municipality ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MunicipalityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{municipality.Label}
	default:
		err = &NotSingularError{municipality.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MunicipalityQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Municipalities.
func (mq *MunicipalityQuery) All(ctx context.Context) ([]*Municipality, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Municipality, *MunicipalityQuery]()
	return withInterceptors[[]*Municipality](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MunicipalityQuery) AllX(ctx context.Context) []*Municipality {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Municipality IDs.
func (mq *MunicipalityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(municipality.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MunicipalityQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MunicipalityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MunicipalityQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MunicipalityQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MunicipalityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MunicipalityQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MunicipalityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MunicipalityQuery) Clone() *MunicipalityQuery {
	if mq == nil {
		return nil
	}
	return &MunicipalityQuery{
		config:     mq.config,
		ctx:        mq.ctx.Clone(),
		order:      append([]municipality.OrderOption{}, mq.order...),
		inters:     append([]Interceptor{}, mq.inters...),
		predicates: append([]predicate.Municipality{}, mq.predicates...),
		withMeshis: mq.withMeshis.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithMeshis tells the query-builder to eager-load the nodes that are connected to
// the "meshis" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MunicipalityQuery) WithMeshis(opts ...func(*MeshiQuery)) *MunicipalityQuery {
	query := (&MeshiClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withMeshis = query
	return mq
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
//	client.Municipality.Query().
//		GroupBy(municipality.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MunicipalityQuery) GroupBy(field string, fields ...string) *MunicipalityGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MunicipalityGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = municipality.Label
	grbuild.scan = grbuild.Scan
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
//	client.Municipality.Query().
//		Select(municipality.FieldName).
//		Scan(ctx, &v)
func (mq *MunicipalityQuery) Select(fields ...string) *MunicipalitySelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MunicipalitySelect{MunicipalityQuery: mq}
	sbuild.label = municipality.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MunicipalitySelect configured with the given aggregations.
func (mq *MunicipalityQuery) Aggregate(fns ...AggregateFunc) *MunicipalitySelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MunicipalityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !municipality.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MunicipalityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Municipality, error) {
	var (
		nodes       = []*Municipality{}
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withMeshis != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Municipality).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Municipality{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withMeshis; query != nil {
		if err := mq.loadMeshis(ctx, query, nodes,
			func(n *Municipality) { n.Edges.Meshis = []*Meshi{} },
			func(n *Municipality, e *Meshi) { n.Edges.Meshis = append(n.Edges.Meshis, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mq.withNamedMeshis {
		if err := mq.loadMeshis(ctx, query, nodes,
			func(n *Municipality) { n.appendNamedMeshis(name) },
			func(n *Municipality, e *Meshi) { n.appendNamedMeshis(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range mq.loadTotal {
		if err := mq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MunicipalityQuery) loadMeshis(ctx context.Context, query *MeshiQuery, nodes []*Municipality, init func(*Municipality), assign func(*Municipality, *Meshi)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Municipality)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Meshi(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(municipality.MeshisColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.municipality_meshis
		if fk == nil {
			return fmt.Errorf(`foreign-key "municipality_meshis" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "municipality_meshis" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MunicipalityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MunicipalityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(municipality.Table, municipality.Columns, sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, municipality.FieldID)
		for i := range fields {
			if fields[i] != municipality.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MunicipalityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(municipality.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = municipality.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMeshis tells the query-builder to eager-load the nodes that are connected to the "meshis"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mq *MunicipalityQuery) WithNamedMeshis(name string, opts ...func(*MeshiQuery)) *MunicipalityQuery {
	query := (&MeshiClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mq.withNamedMeshis == nil {
		mq.withNamedMeshis = make(map[string]*MeshiQuery)
	}
	mq.withNamedMeshis[name] = query
	return mq
}

// MunicipalityGroupBy is the group-by builder for Municipality entities.
type MunicipalityGroupBy struct {
	selector
	build *MunicipalityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MunicipalityGroupBy) Aggregate(fns ...AggregateFunc) *MunicipalityGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MunicipalityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MunicipalityQuery, *MunicipalityGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MunicipalityGroupBy) sqlScan(ctx context.Context, root *MunicipalityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MunicipalitySelect is the builder for selecting fields of Municipality entities.
type MunicipalitySelect struct {
	*MunicipalityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MunicipalitySelect) Aggregate(fns ...AggregateFunc) *MunicipalitySelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MunicipalitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MunicipalityQuery, *MunicipalitySelect](ctx, ms.MunicipalityQuery, ms, ms.inters, v)
}

func (ms *MunicipalitySelect) sqlScan(ctx context.Context, root *MunicipalityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
