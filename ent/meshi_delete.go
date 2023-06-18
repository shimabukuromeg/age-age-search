// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MeshiDelete is the builder for deleting a Meshi entity.
type MeshiDelete struct {
	config
	hooks    []Hook
	mutation *MeshiMutation
}

// Where appends a list predicates to the MeshiDelete builder.
func (md *MeshiDelete) Where(ps ...predicate.Meshi) *MeshiDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MeshiDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MeshiDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MeshiDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(meshi.Table, sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MeshiDeleteOne is the builder for deleting a single Meshi entity.
type MeshiDeleteOne struct {
	md *MeshiDelete
}

// Where appends a list predicates to the MeshiDelete builder.
func (mdo *MeshiDeleteOne) Where(ps ...predicate.Meshi) *MeshiDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MeshiDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{meshi.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MeshiDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}