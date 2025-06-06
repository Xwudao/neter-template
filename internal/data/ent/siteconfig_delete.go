// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go-kitboxpro/internal/data/ent/predicate"
	"go-kitboxpro/internal/data/ent/siteconfig"
)

// SiteConfigDelete is the builder for deleting a SiteConfig entity.
type SiteConfigDelete struct {
	config
	hooks    []Hook
	mutation *SiteConfigMutation
}

// Where appends a list predicates to the SiteConfigDelete builder.
func (scd *SiteConfigDelete) Where(ps ...predicate.SiteConfig) *SiteConfigDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *SiteConfigDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *SiteConfigDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *SiteConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(siteconfig.Table, sqlgraph.NewFieldSpec(siteconfig.FieldID, field.TypeInt64))
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// SiteConfigDeleteOne is the builder for deleting a single SiteConfig entity.
type SiteConfigDeleteOne struct {
	scd *SiteConfigDelete
}

// Where appends a list predicates to the SiteConfigDelete builder.
func (scdo *SiteConfigDeleteOne) Where(ps ...predicate.SiteConfig) *SiteConfigDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *SiteConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{siteconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *SiteConfigDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
