// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dictdetail"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictDetailDelete is the builder for deleting a DictDetail entity.
type DictDetailDelete struct {
	config
	hooks    []Hook
	mutation *DictDetailMutation
}

// Where appends a list predicates to the DictDetailDelete builder.
func (ddd *DictDetailDelete) Where(ps ...predicate.DictDetail) *DictDetailDelete {
	ddd.mutation.Where(ps...)
	return ddd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ddd *DictDetailDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ddd.sqlExec, ddd.mutation, ddd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ddd *DictDetailDelete) ExecX(ctx context.Context) int {
	n, err := ddd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ddd *DictDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dictdetail.Table, sqlgraph.NewFieldSpec(dictdetail.FieldID, field.TypeUint32))
	if ps := ddd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ddd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ddd.mutation.done = true
	return affected, err
}

// DictDetailDeleteOne is the builder for deleting a single DictDetail entity.
type DictDetailDeleteOne struct {
	ddd *DictDetailDelete
}

// Where appends a list predicates to the DictDetailDelete builder.
func (dddo *DictDetailDeleteOne) Where(ps ...predicate.DictDetail) *DictDetailDeleteOne {
	dddo.ddd.mutation.Where(ps...)
	return dddo
}

// Exec executes the deletion query.
func (dddo *DictDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := dddo.ddd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dictdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dddo *DictDetailDeleteOne) ExecX(ctx context.Context) {
	if err := dddo.Exec(ctx); err != nil {
		panic(err)
	}
}
