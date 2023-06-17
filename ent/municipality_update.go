// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MunicipalityUpdate is the builder for updating Municipality entities.
type MunicipalityUpdate struct {
	config
	hooks    []Hook
	mutation *MunicipalityMutation
}

// Where appends a list predicates to the MunicipalityUpdate builder.
func (mu *MunicipalityUpdate) Where(ps ...predicate.Municipality) *MunicipalityUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MunicipalityUpdate) SetName(s string) *MunicipalityUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MunicipalityUpdate) SetCreatedAt(t time.Time) *MunicipalityUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MunicipalityUpdate) SetNillableCreatedAt(t *time.Time) *MunicipalityUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// AddMeshiIDs adds the "meshis" edge to the Meshi entity by IDs.
func (mu *MunicipalityUpdate) AddMeshiIDs(ids ...int) *MunicipalityUpdate {
	mu.mutation.AddMeshiIDs(ids...)
	return mu
}

// AddMeshis adds the "meshis" edges to the Meshi entity.
func (mu *MunicipalityUpdate) AddMeshis(m ...*Meshi) *MunicipalityUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMeshiIDs(ids...)
}

// Mutation returns the MunicipalityMutation object of the builder.
func (mu *MunicipalityUpdate) Mutation() *MunicipalityMutation {
	return mu.mutation
}

// ClearMeshis clears all "meshis" edges to the Meshi entity.
func (mu *MunicipalityUpdate) ClearMeshis() *MunicipalityUpdate {
	mu.mutation.ClearMeshis()
	return mu
}

// RemoveMeshiIDs removes the "meshis" edge to Meshi entities by IDs.
func (mu *MunicipalityUpdate) RemoveMeshiIDs(ids ...int) *MunicipalityUpdate {
	mu.mutation.RemoveMeshiIDs(ids...)
	return mu
}

// RemoveMeshis removes "meshis" edges to Meshi entities.
func (mu *MunicipalityUpdate) RemoveMeshis(m ...*Meshi) *MunicipalityUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMeshiIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MunicipalityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MunicipalityUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MunicipalityUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MunicipalityUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MunicipalityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(municipality.Table, municipality.Columns, sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(municipality.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(municipality.FieldCreatedAt, field.TypeTime, value)
	}
	if mu.mutation.MeshisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMeshisIDs(); len(nodes) > 0 && !mu.mutation.MeshisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MeshisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{municipality.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MunicipalityUpdateOne is the builder for updating a single Municipality entity.
type MunicipalityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MunicipalityMutation
}

// SetName sets the "name" field.
func (muo *MunicipalityUpdateOne) SetName(s string) *MunicipalityUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MunicipalityUpdateOne) SetCreatedAt(t time.Time) *MunicipalityUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MunicipalityUpdateOne) SetNillableCreatedAt(t *time.Time) *MunicipalityUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// AddMeshiIDs adds the "meshis" edge to the Meshi entity by IDs.
func (muo *MunicipalityUpdateOne) AddMeshiIDs(ids ...int) *MunicipalityUpdateOne {
	muo.mutation.AddMeshiIDs(ids...)
	return muo
}

// AddMeshis adds the "meshis" edges to the Meshi entity.
func (muo *MunicipalityUpdateOne) AddMeshis(m ...*Meshi) *MunicipalityUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMeshiIDs(ids...)
}

// Mutation returns the MunicipalityMutation object of the builder.
func (muo *MunicipalityUpdateOne) Mutation() *MunicipalityMutation {
	return muo.mutation
}

// ClearMeshis clears all "meshis" edges to the Meshi entity.
func (muo *MunicipalityUpdateOne) ClearMeshis() *MunicipalityUpdateOne {
	muo.mutation.ClearMeshis()
	return muo
}

// RemoveMeshiIDs removes the "meshis" edge to Meshi entities by IDs.
func (muo *MunicipalityUpdateOne) RemoveMeshiIDs(ids ...int) *MunicipalityUpdateOne {
	muo.mutation.RemoveMeshiIDs(ids...)
	return muo
}

// RemoveMeshis removes "meshis" edges to Meshi entities.
func (muo *MunicipalityUpdateOne) RemoveMeshis(m ...*Meshi) *MunicipalityUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMeshiIDs(ids...)
}

// Where appends a list predicates to the MunicipalityUpdate builder.
func (muo *MunicipalityUpdateOne) Where(ps ...predicate.Municipality) *MunicipalityUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MunicipalityUpdateOne) Select(field string, fields ...string) *MunicipalityUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Municipality entity.
func (muo *MunicipalityUpdateOne) Save(ctx context.Context) (*Municipality, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MunicipalityUpdateOne) SaveX(ctx context.Context) *Municipality {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MunicipalityUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MunicipalityUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MunicipalityUpdateOne) sqlSave(ctx context.Context) (_node *Municipality, err error) {
	_spec := sqlgraph.NewUpdateSpec(municipality.Table, municipality.Columns, sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Municipality.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, municipality.FieldID)
		for _, f := range fields {
			if !municipality.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != municipality.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(municipality.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(municipality.FieldCreatedAt, field.TypeTime, value)
	}
	if muo.mutation.MeshisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMeshisIDs(); len(nodes) > 0 && !muo.mutation.MeshisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MeshisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   municipality.MeshisTable,
			Columns: []string{municipality.MeshisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Municipality{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{municipality.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
