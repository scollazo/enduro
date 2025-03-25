// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/aip"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/predicate"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/workflow"
)

// AIPQuery is the builder for querying AIP entities.
type AIPQuery struct {
	config
	ctx           *QueryContext
	order         []aip.OrderOption
	inters        []Interceptor
	predicates    []predicate.AIP
	withLocation  *LocationQuery
	withWorkflows *WorkflowQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AIPQuery builder.
func (aq *AIPQuery) Where(ps ...predicate.AIP) *AIPQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AIPQuery) Limit(limit int) *AIPQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AIPQuery) Offset(offset int) *AIPQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AIPQuery) Unique(unique bool) *AIPQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AIPQuery) Order(o ...aip.OrderOption) *AIPQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryLocation chains the current query on the "location" edge.
func (aq *AIPQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(aip.Table, aip.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, aip.LocationTable, aip.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkflows chains the current query on the "workflows" edge.
func (aq *AIPQuery) QueryWorkflows() *WorkflowQuery {
	query := (&WorkflowClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(aip.Table, aip.FieldID, selector),
			sqlgraph.To(workflow.Table, workflow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, aip.WorkflowsTable, aip.WorkflowsColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AIP entity from the query.
// Returns a *NotFoundError when no AIP was found.
func (aq *AIPQuery) First(ctx context.Context) (*AIP, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{aip.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AIPQuery) FirstX(ctx context.Context) *AIP {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AIP ID from the query.
// Returns a *NotFoundError when no AIP ID was found.
func (aq *AIPQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{aip.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AIPQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AIP entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AIP entity is found.
// Returns a *NotFoundError when no AIP entities are found.
func (aq *AIPQuery) Only(ctx context.Context) (*AIP, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{aip.Label}
	default:
		return nil, &NotSingularError{aip.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AIPQuery) OnlyX(ctx context.Context) *AIP {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AIP ID in the query.
// Returns a *NotSingularError when more than one AIP ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AIPQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{aip.Label}
	default:
		err = &NotSingularError{aip.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AIPQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AIPs.
func (aq *AIPQuery) All(ctx context.Context) ([]*AIP, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AIP, *AIPQuery]()
	return withInterceptors[[]*AIP](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AIPQuery) AllX(ctx context.Context) []*AIP {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AIP IDs.
func (aq *AIPQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(aip.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AIPQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AIPQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AIPQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AIPQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AIPQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AIPQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AIPQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AIPQuery) Clone() *AIPQuery {
	if aq == nil {
		return nil
	}
	return &AIPQuery{
		config:        aq.config,
		ctx:           aq.ctx.Clone(),
		order:         append([]aip.OrderOption{}, aq.order...),
		inters:        append([]Interceptor{}, aq.inters...),
		predicates:    append([]predicate.AIP{}, aq.predicates...),
		withLocation:  aq.withLocation.Clone(),
		withWorkflows: aq.withWorkflows.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AIPQuery) WithLocation(opts ...func(*LocationQuery)) *AIPQuery {
	query := (&LocationClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withLocation = query
	return aq
}

// WithWorkflows tells the query-builder to eager-load the nodes that are connected to
// the "workflows" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AIPQuery) WithWorkflows(opts ...func(*WorkflowQuery)) *AIPQuery {
	query := (&WorkflowClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withWorkflows = query
	return aq
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
//	client.AIP.Query().
//		GroupBy(aip.FieldName).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (aq *AIPQuery) GroupBy(field string, fields ...string) *AIPGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AIPGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = aip.Label
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
//	client.AIP.Query().
//		Select(aip.FieldName).
//		Scan(ctx, &v)
func (aq *AIPQuery) Select(fields ...string) *AIPSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AIPSelect{AIPQuery: aq}
	sbuild.label = aip.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AIPSelect configured with the given aggregations.
func (aq *AIPQuery) Aggregate(fns ...AggregateFunc) *AIPSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AIPQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !aip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AIPQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AIP, error) {
	var (
		nodes       = []*AIP{}
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withLocation != nil,
			aq.withWorkflows != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AIP).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AIP{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withLocation; query != nil {
		if err := aq.loadLocation(ctx, query, nodes, nil,
			func(n *AIP, e *Location) { n.Edges.Location = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withWorkflows; query != nil {
		if err := aq.loadWorkflows(ctx, query, nodes,
			func(n *AIP) { n.Edges.Workflows = []*Workflow{} },
			func(n *AIP, e *Workflow) { n.Edges.Workflows = append(n.Edges.Workflows, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AIPQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*AIP, init func(*AIP), assign func(*AIP, *Location)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AIP)
	for i := range nodes {
		fk := nodes[i].LocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AIPQuery) loadWorkflows(ctx context.Context, query *WorkflowQuery, nodes []*AIP, init func(*AIP), assign func(*AIP, *Workflow)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*AIP)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workflow.FieldAipID)
	}
	query.Where(predicate.Workflow(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(aip.WorkflowsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AipID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "aip_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AIPQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AIPQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(aip.Table, aip.Columns, sqlgraph.NewFieldSpec(aip.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, aip.FieldID)
		for i := range fields {
			if fields[i] != aip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withLocation != nil {
			_spec.Node.AddColumnOnce(aip.FieldLocationID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AIPQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(aip.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = aip.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AIPGroupBy is the group-by builder for AIP entities.
type AIPGroupBy struct {
	selector
	build *AIPQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AIPGroupBy) Aggregate(fns ...AggregateFunc) *AIPGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AIPGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AIPQuery, *AIPGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AIPGroupBy) sqlScan(ctx context.Context, root *AIPQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AIPSelect is the builder for selecting fields of AIP entities.
type AIPSelect struct {
	*AIPQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AIPSelect) Aggregate(fns ...AggregateFunc) *AIPSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AIPSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AIPQuery, *AIPSelect](ctx, as.AIPQuery, as, as.inters, v)
}

func (as *AIPSelect) sqlScan(ctx context.Context, root *AIPQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
