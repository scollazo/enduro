// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/predicate"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/preservationaction"
)

// PkgQuery is the builder for querying Pkg entities.
type PkgQuery struct {
	config
	ctx                     *QueryContext
	order                   []pkg.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Pkg
	withPreservationActions *PreservationActionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PkgQuery builder.
func (pq *PkgQuery) Where(ps ...predicate.Pkg) *PkgQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PkgQuery) Limit(limit int) *PkgQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PkgQuery) Offset(offset int) *PkgQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PkgQuery) Unique(unique bool) *PkgQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PkgQuery) Order(o ...pkg.OrderOption) *PkgQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPreservationActions chains the current query on the "preservation_actions" edge.
func (pq *PkgQuery) QueryPreservationActions() *PreservationActionQuery {
	query := (&PreservationActionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pkg.Table, pkg.FieldID, selector),
			sqlgraph.To(preservationaction.Table, preservationaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pkg.PreservationActionsTable, pkg.PreservationActionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pkg entity from the query.
// Returns a *NotFoundError when no Pkg was found.
func (pq *PkgQuery) First(ctx context.Context) (*Pkg, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pkg.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PkgQuery) FirstX(ctx context.Context) *Pkg {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pkg ID from the query.
// Returns a *NotFoundError when no Pkg ID was found.
func (pq *PkgQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pkg.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PkgQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pkg entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pkg entity is found.
// Returns a *NotFoundError when no Pkg entities are found.
func (pq *PkgQuery) Only(ctx context.Context) (*Pkg, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pkg.Label}
	default:
		return nil, &NotSingularError{pkg.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PkgQuery) OnlyX(ctx context.Context) *Pkg {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pkg ID in the query.
// Returns a *NotSingularError when more than one Pkg ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PkgQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pkg.Label}
	default:
		err = &NotSingularError{pkg.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PkgQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pkgs.
func (pq *PkgQuery) All(ctx context.Context) ([]*Pkg, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pkg, *PkgQuery]()
	return withInterceptors[[]*Pkg](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PkgQuery) AllX(ctx context.Context) []*Pkg {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pkg IDs.
func (pq *PkgQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(pkg.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PkgQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PkgQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PkgQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PkgQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PkgQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PkgQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PkgQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PkgQuery) Clone() *PkgQuery {
	if pq == nil {
		return nil
	}
	return &PkgQuery{
		config:                  pq.config,
		ctx:                     pq.ctx.Clone(),
		order:                   append([]pkg.OrderOption{}, pq.order...),
		inters:                  append([]Interceptor{}, pq.inters...),
		predicates:              append([]predicate.Pkg{}, pq.predicates...),
		withPreservationActions: pq.withPreservationActions.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPreservationActions tells the query-builder to eager-load the nodes that are connected to
// the "preservation_actions" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PkgQuery) WithPreservationActions(opts ...func(*PreservationActionQuery)) *PkgQuery {
	query := (&PreservationActionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPreservationActions = query
	return pq
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
//	client.Pkg.Query().
//		GroupBy(pkg.FieldName).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (pq *PkgQuery) GroupBy(field string, fields ...string) *PkgGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PkgGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pkg.Label
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
//	client.Pkg.Query().
//		Select(pkg.FieldName).
//		Scan(ctx, &v)
func (pq *PkgQuery) Select(fields ...string) *PkgSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PkgSelect{PkgQuery: pq}
	sbuild.label = pkg.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PkgSelect configured with the given aggregations.
func (pq *PkgQuery) Aggregate(fns ...AggregateFunc) *PkgSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PkgQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pkg.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PkgQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pkg, error) {
	var (
		nodes       = []*Pkg{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withPreservationActions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pkg).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pkg{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPreservationActions; query != nil {
		if err := pq.loadPreservationActions(ctx, query, nodes,
			func(n *Pkg) { n.Edges.PreservationActions = []*PreservationAction{} },
			func(n *Pkg, e *PreservationAction) {
				n.Edges.PreservationActions = append(n.Edges.PreservationActions, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PkgQuery) loadPreservationActions(ctx context.Context, query *PreservationActionQuery, nodes []*Pkg, init func(*Pkg), assign func(*Pkg, *PreservationAction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Pkg)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(preservationaction.FieldPackageID)
	}
	query.Where(predicate.PreservationAction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(pkg.PreservationActionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PackageID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "package_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PkgQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PkgQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pkg.Table, pkg.Columns, sqlgraph.NewFieldSpec(pkg.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pkg.FieldID)
		for i := range fields {
			if fields[i] != pkg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PkgQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pkg.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pkg.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PkgGroupBy is the group-by builder for Pkg entities.
type PkgGroupBy struct {
	selector
	build *PkgQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PkgGroupBy) Aggregate(fns ...AggregateFunc) *PkgGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PkgGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PkgQuery, *PkgGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PkgGroupBy) sqlScan(ctx context.Context, root *PkgQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PkgSelect is the builder for selecting fields of Pkg entities.
type PkgSelect struct {
	*PkgQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PkgSelect) Aggregate(fns ...AggregateFunc) *PkgSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PkgSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PkgQuery, *PkgSelect](ctx, ps.PkgQuery, ps, ps.inters, v)
}

func (ps *PkgSelect) sqlScan(ctx context.Context, root *PkgQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
