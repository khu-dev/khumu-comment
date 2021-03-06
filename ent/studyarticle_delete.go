// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/predicate"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// StudyArticleDelete is the builder for deleting a StudyArticle entity.
type StudyArticleDelete struct {
	config
	hooks    []Hook
	mutation *StudyArticleMutation
}

// Where adds a new predicate to the StudyArticleDelete builder.
func (sad *StudyArticleDelete) Where(ps ...predicate.StudyArticle) *StudyArticleDelete {
	sad.mutation.predicates = append(sad.mutation.predicates, ps...)
	return sad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sad *StudyArticleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sad.hooks) == 0 {
		affected, err = sad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sad.mutation = mutation
			affected, err = sad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sad.hooks) - 1; i >= 0; i-- {
			mut = sad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sad *StudyArticleDelete) ExecX(ctx context.Context) int {
	n, err := sad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sad *StudyArticleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: studyarticle.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyarticle.FieldID,
			},
		},
	}
	if ps := sad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sad.driver, _spec)
}

// StudyArticleDeleteOne is the builder for deleting a single StudyArticle entity.
type StudyArticleDeleteOne struct {
	sad *StudyArticleDelete
}

// Exec executes the deletion query.
func (sado *StudyArticleDeleteOne) Exec(ctx context.Context) error {
	n, err := sado.sad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{studyarticle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sado *StudyArticleDeleteOne) ExecX(ctx context.Context) {
	sado.sad.ExecX(ctx)
}
