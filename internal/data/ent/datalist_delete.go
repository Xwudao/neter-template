// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go-kitboxpro/internal/data/ent/datalist"
	"go-kitboxpro/internal/data/ent/predicate"
)

// DataListDelete is the builder for deleting a DataList entity.
type DataListDelete struct {
	config
	hooks    []Hook
	mutation *DataListMutation
}

// Where appends a list predicates to the DataListDelete builder.
func (dld *DataListDelete) Where(ps ...predicate.DataList) *DataListDelete {
	dld.mutation.Where(ps...)
	return dld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dld *DataListDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dld.sqlExec, dld.mutation, dld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dld *DataListDelete) ExecX(ctx context.Context) int {
	n, err := dld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dld *DataListDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(datalist.Table, sqlgraph.NewFieldSpec(datalist.FieldID, field.TypeInt64))
	if ps := dld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dld.mutation.done = true
	return affected, err
}

// DataListDeleteOne is the builder for deleting a single DataList entity.
type DataListDeleteOne struct {
	dld *DataListDelete
}

// Where appends a list predicates to the DataListDelete builder.
func (dldo *DataListDeleteOne) Where(ps ...predicate.DataList) *DataListDeleteOne {
	dldo.dld.mutation.Where(ps...)
	return dldo
}

// Exec executes the deletion query.
func (dldo *DataListDeleteOne) Exec(ctx context.Context) error {
	n, err := dldo.dld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{datalist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dldo *DataListDeleteOne) ExecX(ctx context.Context) {
	if err := dldo.Exec(ctx); err != nil {
		panic(err)
	}
}
