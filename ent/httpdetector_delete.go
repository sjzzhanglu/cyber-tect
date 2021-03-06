// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetector"
	"github.com/vicanso/cybertect/ent/predicate"
)

// HTTPDetectorDelete is the builder for deleting a HTTPDetector entity.
type HTTPDetectorDelete struct {
	config
	hooks    []Hook
	mutation *HTTPDetectorMutation
}

// Where adds a new predicate to the HTTPDetectorDelete builder.
func (hdd *HTTPDetectorDelete) Where(ps ...predicate.HTTPDetector) *HTTPDetectorDelete {
	hdd.mutation.predicates = append(hdd.mutation.predicates, ps...)
	return hdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hdd *HTTPDetectorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hdd.hooks) == 0 {
		affected, err = hdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HTTPDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hdd.mutation = mutation
			affected, err = hdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hdd.hooks) - 1; i >= 0; i-- {
			mut = hdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdd *HTTPDetectorDelete) ExecX(ctx context.Context) int {
	n, err := hdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hdd *HTTPDetectorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: httpdetector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httpdetector.FieldID,
			},
		},
	}
	if ps := hdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hdd.driver, _spec)
}

// HTTPDetectorDeleteOne is the builder for deleting a single HTTPDetector entity.
type HTTPDetectorDeleteOne struct {
	hdd *HTTPDetectorDelete
}

// Exec executes the deletion query.
func (hddo *HTTPDetectorDeleteOne) Exec(ctx context.Context) error {
	n, err := hddo.hdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{httpdetector.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hddo *HTTPDetectorDeleteOne) ExecX(ctx context.Context) {
	hddo.hdd.ExecX(ctx)
}
