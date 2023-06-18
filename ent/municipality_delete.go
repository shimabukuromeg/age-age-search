// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MunicipalityDelete is the builder for deleting a Municipality entity.
type MunicipalityDelete struct {
	config
	hooks    []Hook
	mutation *MunicipalityMutation
}

// Where appends a list predicates to the MunicipalityDelete builder.
func (md *MunicipalityDelete) Where(ps ...predicate.Municipality) *MunicipalityDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MunicipalityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MunicipalityDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MunicipalityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(municipality.Table, sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt))
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

// MunicipalityDeleteOne is the builder for deleting a single Municipality entity.
type MunicipalityDeleteOne struct {
	md *MunicipalityDelete
}

// Where appends a list predicates to the MunicipalityDelete builder.
func (mdo *MunicipalityDeleteOne) Where(ps ...predicate.Municipality) *MunicipalityDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MunicipalityDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{municipality.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MunicipalityDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}