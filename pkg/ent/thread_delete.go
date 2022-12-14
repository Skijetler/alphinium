// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/predicate"
	"github.com/Skijetler/alphinium/pkg/ent/thread"
)

// ThreadDelete is the builder for deleting a Thread entity.
type ThreadDelete struct {
	config
	hooks    []Hook
	mutation *ThreadMutation
}

// Where appends a list predicates to the ThreadDelete builder.
func (td *ThreadDelete) Where(ps ...predicate.Thread) *ThreadDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *ThreadDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			if td.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *ThreadDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *ThreadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: thread.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: thread.FieldID,
			},
		},
	}
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ThreadDeleteOne is the builder for deleting a single Thread entity.
type ThreadDeleteOne struct {
	td *ThreadDelete
}

// Exec executes the deletion query.
func (tdo *ThreadDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{thread.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *ThreadDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
