// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// StudyArticleCreate is the builder for creating a StudyArticle entity.
type StudyArticleCreate struct {
	config
	mutation *StudyArticleMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (sac *StudyArticleCreate) SetID(i int) *StudyArticleCreate {
	sac.mutation.SetID(i)
	return sac
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (sac *StudyArticleCreate) AddCommentIDs(ids ...int) *StudyArticleCreate {
	sac.mutation.AddCommentIDs(ids...)
	return sac
}

// AddComments adds the "comments" edges to the Comment entity.
func (sac *StudyArticleCreate) AddComments(c ...*Comment) *StudyArticleCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sac.AddCommentIDs(ids...)
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (sac *StudyArticleCreate) SetAuthorID(id string) *StudyArticleCreate {
	sac.mutation.SetAuthorID(id)
	return sac
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (sac *StudyArticleCreate) SetNillableAuthorID(id *string) *StudyArticleCreate {
	if id != nil {
		sac = sac.SetAuthorID(*id)
	}
	return sac
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (sac *StudyArticleCreate) SetAuthor(k *KhumuUser) *StudyArticleCreate {
	return sac.SetAuthorID(k.ID)
}

// Mutation returns the StudyArticleMutation object of the builder.
func (sac *StudyArticleCreate) Mutation() *StudyArticleMutation {
	return sac.mutation
}

// Save creates the StudyArticle in the database.
func (sac *StudyArticleCreate) Save(ctx context.Context) (*StudyArticle, error) {
	var (
		err  error
		node *StudyArticle
	)
	if len(sac.hooks) == 0 {
		if err = sac.check(); err != nil {
			return nil, err
		}
		node, err = sac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sac.check(); err != nil {
				return nil, err
			}
			sac.mutation = mutation
			node, err = sac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sac.hooks) - 1; i >= 0; i-- {
			mut = sac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sac *StudyArticleCreate) SaveX(ctx context.Context) *StudyArticle {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sac *StudyArticleCreate) check() error {
	return nil
}

func (sac *StudyArticleCreate) sqlSave(ctx context.Context) (*StudyArticle, error) {
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (sac *StudyArticleCreate) createSpec() (*StudyArticle, *sqlgraph.CreateSpec) {
	var (
		_node = &StudyArticle{config: sac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: studyarticle.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyarticle.FieldID,
			},
		}
	)
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := sac.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   studyarticle.CommentsTable,
			Columns: []string{studyarticle.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   studyarticle.AuthorTable,
			Columns: []string{studyarticle.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: khumuuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.author_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StudyArticleCreateBulk is the builder for creating many StudyArticle entities in bulk.
type StudyArticleCreateBulk struct {
	config
	builders []*StudyArticleCreate
}

// Save creates the StudyArticle entities in the database.
func (sacb *StudyArticleCreateBulk) Save(ctx context.Context) ([]*StudyArticle, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*StudyArticle, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyArticleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *StudyArticleCreateBulk) SaveX(ctx context.Context) []*StudyArticle {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
