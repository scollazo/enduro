// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/storage/enums"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/aip"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/predicate"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/workflow"
	"github.com/google/uuid"
)

// AIPUpdate is the builder for updating AIP entities.
type AIPUpdate struct {
	config
	hooks    []Hook
	mutation *AIPMutation
}

// Where appends a list predicates to the AIPUpdate builder.
func (au *AIPUpdate) Where(ps ...predicate.AIP) *AIPUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AIPUpdate) SetName(s string) *AIPUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AIPUpdate) SetNillableName(s *string) *AIPUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetAipID sets the "aip_id" field.
func (au *AIPUpdate) SetAipID(u uuid.UUID) *AIPUpdate {
	au.mutation.SetAipID(u)
	return au
}

// SetNillableAipID sets the "aip_id" field if the given value is not nil.
func (au *AIPUpdate) SetNillableAipID(u *uuid.UUID) *AIPUpdate {
	if u != nil {
		au.SetAipID(*u)
	}
	return au
}

// SetLocationID sets the "location_id" field.
func (au *AIPUpdate) SetLocationID(i int) *AIPUpdate {
	au.mutation.SetLocationID(i)
	return au
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (au *AIPUpdate) SetNillableLocationID(i *int) *AIPUpdate {
	if i != nil {
		au.SetLocationID(*i)
	}
	return au
}

// ClearLocationID clears the value of the "location_id" field.
func (au *AIPUpdate) ClearLocationID() *AIPUpdate {
	au.mutation.ClearLocationID()
	return au
}

// SetStatus sets the "status" field.
func (au *AIPUpdate) SetStatus(es enums.AIPStatus) *AIPUpdate {
	au.mutation.SetStatus(es)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AIPUpdate) SetNillableStatus(es *enums.AIPStatus) *AIPUpdate {
	if es != nil {
		au.SetStatus(*es)
	}
	return au
}

// SetObjectKey sets the "object_key" field.
func (au *AIPUpdate) SetObjectKey(u uuid.UUID) *AIPUpdate {
	au.mutation.SetObjectKey(u)
	return au
}

// SetNillableObjectKey sets the "object_key" field if the given value is not nil.
func (au *AIPUpdate) SetNillableObjectKey(u *uuid.UUID) *AIPUpdate {
	if u != nil {
		au.SetObjectKey(*u)
	}
	return au
}

// SetLocation sets the "location" edge to the Location entity.
func (au *AIPUpdate) SetLocation(l *Location) *AIPUpdate {
	return au.SetLocationID(l.ID)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (au *AIPUpdate) AddWorkflowIDs(ids ...int) *AIPUpdate {
	au.mutation.AddWorkflowIDs(ids...)
	return au
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (au *AIPUpdate) AddWorkflows(w ...*Workflow) *AIPUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return au.AddWorkflowIDs(ids...)
}

// Mutation returns the AIPMutation object of the builder.
func (au *AIPUpdate) Mutation() *AIPMutation {
	return au.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (au *AIPUpdate) ClearLocation() *AIPUpdate {
	au.mutation.ClearLocation()
	return au
}

// ClearWorkflows clears all "workflows" edges to the Workflow entity.
func (au *AIPUpdate) ClearWorkflows() *AIPUpdate {
	au.mutation.ClearWorkflows()
	return au
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflow entities by IDs.
func (au *AIPUpdate) RemoveWorkflowIDs(ids ...int) *AIPUpdate {
	au.mutation.RemoveWorkflowIDs(ids...)
	return au
}

// RemoveWorkflows removes "workflows" edges to Workflow entities.
func (au *AIPUpdate) RemoveWorkflows(w ...*Workflow) *AIPUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return au.RemoveWorkflowIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AIPUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AIPUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AIPUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AIPUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AIPUpdate) check() error {
	if v, ok := au.mutation.Status(); ok {
		if err := aip.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "AIP.status": %w`, err)}
		}
	}
	return nil
}

func (au *AIPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(aip.Table, aip.Columns, sqlgraph.NewFieldSpec(aip.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(aip.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.AipID(); ok {
		_spec.SetField(aip.FieldAipID, field.TypeUUID, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(aip.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := au.mutation.ObjectKey(); ok {
		_spec.SetField(aip.FieldObjectKey, field.TypeUUID, value)
	}
	if au.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   aip.LocationTable,
			Columns: []string{aip.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   aip.LocationTable,
			Columns: []string{aip.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !au.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{aip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AIPUpdateOne is the builder for updating a single AIP entity.
type AIPUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AIPMutation
}

// SetName sets the "name" field.
func (auo *AIPUpdateOne) SetName(s string) *AIPUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AIPUpdateOne) SetNillableName(s *string) *AIPUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetAipID sets the "aip_id" field.
func (auo *AIPUpdateOne) SetAipID(u uuid.UUID) *AIPUpdateOne {
	auo.mutation.SetAipID(u)
	return auo
}

// SetNillableAipID sets the "aip_id" field if the given value is not nil.
func (auo *AIPUpdateOne) SetNillableAipID(u *uuid.UUID) *AIPUpdateOne {
	if u != nil {
		auo.SetAipID(*u)
	}
	return auo
}

// SetLocationID sets the "location_id" field.
func (auo *AIPUpdateOne) SetLocationID(i int) *AIPUpdateOne {
	auo.mutation.SetLocationID(i)
	return auo
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (auo *AIPUpdateOne) SetNillableLocationID(i *int) *AIPUpdateOne {
	if i != nil {
		auo.SetLocationID(*i)
	}
	return auo
}

// ClearLocationID clears the value of the "location_id" field.
func (auo *AIPUpdateOne) ClearLocationID() *AIPUpdateOne {
	auo.mutation.ClearLocationID()
	return auo
}

// SetStatus sets the "status" field.
func (auo *AIPUpdateOne) SetStatus(es enums.AIPStatus) *AIPUpdateOne {
	auo.mutation.SetStatus(es)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AIPUpdateOne) SetNillableStatus(es *enums.AIPStatus) *AIPUpdateOne {
	if es != nil {
		auo.SetStatus(*es)
	}
	return auo
}

// SetObjectKey sets the "object_key" field.
func (auo *AIPUpdateOne) SetObjectKey(u uuid.UUID) *AIPUpdateOne {
	auo.mutation.SetObjectKey(u)
	return auo
}

// SetNillableObjectKey sets the "object_key" field if the given value is not nil.
func (auo *AIPUpdateOne) SetNillableObjectKey(u *uuid.UUID) *AIPUpdateOne {
	if u != nil {
		auo.SetObjectKey(*u)
	}
	return auo
}

// SetLocation sets the "location" edge to the Location entity.
func (auo *AIPUpdateOne) SetLocation(l *Location) *AIPUpdateOne {
	return auo.SetLocationID(l.ID)
}

// AddWorkflowIDs adds the "workflows" edge to the Workflow entity by IDs.
func (auo *AIPUpdateOne) AddWorkflowIDs(ids ...int) *AIPUpdateOne {
	auo.mutation.AddWorkflowIDs(ids...)
	return auo
}

// AddWorkflows adds the "workflows" edges to the Workflow entity.
func (auo *AIPUpdateOne) AddWorkflows(w ...*Workflow) *AIPUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return auo.AddWorkflowIDs(ids...)
}

// Mutation returns the AIPMutation object of the builder.
func (auo *AIPUpdateOne) Mutation() *AIPMutation {
	return auo.mutation
}

// ClearLocation clears the "location" edge to the Location entity.
func (auo *AIPUpdateOne) ClearLocation() *AIPUpdateOne {
	auo.mutation.ClearLocation()
	return auo
}

// ClearWorkflows clears all "workflows" edges to the Workflow entity.
func (auo *AIPUpdateOne) ClearWorkflows() *AIPUpdateOne {
	auo.mutation.ClearWorkflows()
	return auo
}

// RemoveWorkflowIDs removes the "workflows" edge to Workflow entities by IDs.
func (auo *AIPUpdateOne) RemoveWorkflowIDs(ids ...int) *AIPUpdateOne {
	auo.mutation.RemoveWorkflowIDs(ids...)
	return auo
}

// RemoveWorkflows removes "workflows" edges to Workflow entities.
func (auo *AIPUpdateOne) RemoveWorkflows(w ...*Workflow) *AIPUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return auo.RemoveWorkflowIDs(ids...)
}

// Where appends a list predicates to the AIPUpdate builder.
func (auo *AIPUpdateOne) Where(ps ...predicate.AIP) *AIPUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AIPUpdateOne) Select(field string, fields ...string) *AIPUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated AIP entity.
func (auo *AIPUpdateOne) Save(ctx context.Context) (*AIP, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AIPUpdateOne) SaveX(ctx context.Context) *AIP {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AIPUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AIPUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AIPUpdateOne) check() error {
	if v, ok := auo.mutation.Status(); ok {
		if err := aip.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "AIP.status": %w`, err)}
		}
	}
	return nil
}

func (auo *AIPUpdateOne) sqlSave(ctx context.Context) (_node *AIP, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(aip.Table, aip.Columns, sqlgraph.NewFieldSpec(aip.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AIP.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, aip.FieldID)
		for _, f := range fields {
			if !aip.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != aip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(aip.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.AipID(); ok {
		_spec.SetField(aip.FieldAipID, field.TypeUUID, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(aip.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.ObjectKey(); ok {
		_spec.SetField(aip.FieldObjectKey, field.TypeUUID, value)
	}
	if auo.mutation.LocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   aip.LocationTable,
			Columns: []string{aip.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.LocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   aip.LocationTable,
			Columns: []string{aip.LocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedWorkflowsIDs(); len(nodes) > 0 && !auo.mutation.WorkflowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.WorkflowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   aip.WorkflowsTable,
			Columns: []string{aip.WorkflowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AIP{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{aip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
