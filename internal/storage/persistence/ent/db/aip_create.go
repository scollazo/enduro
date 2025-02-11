// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/aip"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/types"
	"github.com/google/uuid"
)

// AIPCreate is the builder for creating a AIP entity.
type AIPCreate struct {
	config
	mutation *AIPMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AIPCreate) SetName(s string) *AIPCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetAipID sets the "aip_id" field.
func (ac *AIPCreate) SetAipID(u uuid.UUID) *AIPCreate {
	ac.mutation.SetAipID(u)
	return ac
}

// SetLocationID sets the "location_id" field.
func (ac *AIPCreate) SetLocationID(i int) *AIPCreate {
	ac.mutation.SetLocationID(i)
	return ac
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (ac *AIPCreate) SetNillableLocationID(i *int) *AIPCreate {
	if i != nil {
		ac.SetLocationID(*i)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AIPCreate) SetStatus(ts types.AIPStatus) *AIPCreate {
	ac.mutation.SetStatus(ts)
	return ac
}

// SetObjectKey sets the "object_key" field.
func (ac *AIPCreate) SetObjectKey(u uuid.UUID) *AIPCreate {
	ac.mutation.SetObjectKey(u)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AIPCreate) SetCreatedAt(t time.Time) *AIPCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AIPCreate) SetNillableCreatedAt(t *time.Time) *AIPCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetLocation sets the "location" edge to the Location entity.
func (ac *AIPCreate) SetLocation(l *Location) *AIPCreate {
	return ac.SetLocationID(l.ID)
}

// Mutation returns the AIPMutation object of the builder.
func (ac *AIPCreate) Mutation() *AIPMutation {
	return ac.mutation
}

// Save creates the AIP in the database.
func (ac *AIPCreate) Save(ctx context.Context) (*AIP, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AIPCreate) SaveX(ctx context.Context) *AIP {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AIPCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AIPCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AIPCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := aip.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AIPCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "AIP.name"`)}
	}
	if _, ok := ac.mutation.AipID(); !ok {
		return &ValidationError{Name: "aip_id", err: errors.New(`db: missing required field "AIP.aip_id"`)}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "AIP.status"`)}
	}
	if v, ok := ac.mutation.Status(); ok {
		if err := aip.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "AIP.status": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ObjectKey(); !ok {
		return &ValidationError{Name: "object_key", err: errors.New(`db: missing required field "AIP.object_key"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "AIP.created_at"`)}
	}
	return nil
}

func (ac *AIPCreate) sqlSave(ctx context.Context) (*AIP, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AIPCreate) createSpec() (*AIP, *sqlgraph.CreateSpec) {
	var (
		_node = &AIP{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(aip.Table, sqlgraph.NewFieldSpec(aip.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(aip.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.AipID(); ok {
		_spec.SetField(aip.FieldAipID, field.TypeUUID, value)
		_node.AipID = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(aip.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.ObjectKey(); ok {
		_spec.SetField(aip.FieldObjectKey, field.TypeUUID, value)
		_node.ObjectKey = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(aip.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.LocationIDs(); len(nodes) > 0 {
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
		_node.LocationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AIPCreateBulk is the builder for creating many AIP entities in bulk.
type AIPCreateBulk struct {
	config
	err      error
	builders []*AIPCreate
}

// Save creates the AIP entities in the database.
func (acb *AIPCreateBulk) Save(ctx context.Context) ([]*AIP, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*AIP, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AIPMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AIPCreateBulk) SaveX(ctx context.Context) []*AIP {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AIPCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AIPCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
