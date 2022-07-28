// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"example/golang-rest-api/ent/note"
	"example/golang-rest-api/ent/predicate"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NoteDelete is the builder for deleting a Note entity.
type NoteDelete struct {
	config
	hooks    []Hook
	mutation *NoteMutation
}

// Where appends a list predicates to the NoteDelete builder.
func (nd *NoteDelete) Where(ps ...predicate.Note) *NoteDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NoteDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nd.hooks) == 0 {
		affected, err = nd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nd.mutation = mutation
			affected, err = nd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nd.hooks) - 1; i >= 0; i-- {
			if nd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NoteDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: note.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: note.FieldID,
			},
		},
	}
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// NoteDeleteOne is the builder for deleting a single Note entity.
type NoteDeleteOne struct {
	nd *NoteDelete
}

// Exec executes the deletion query.
func (ndo *NoteDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{note.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NoteDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
