// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetState sets the "state" field.
func (cc *CommentCreate) SetState(s string) *CommentCreate {
	cc.mutation.SetState(s)
	return cc
}

// SetContent sets the "content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetKind sets the "kind" field.
func (cc *CommentCreate) SetKind(s string) *CommentCreate {
	cc.mutation.SetKind(s)
	return cc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (cc *CommentCreate) SetNillableKind(s *string) *CommentCreate {
	if s != nil {
		cc.SetKind(*s)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(i int) *CommentCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (cc *CommentCreate) SetAuthorID(id string) *CommentCreate {
	cc.mutation.SetAuthorID(id)
	return cc
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableAuthorID(id *string) *CommentCreate {
	if id != nil {
		cc = cc.SetAuthorID(*id)
	}
	return cc
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (cc *CommentCreate) SetAuthor(k *KhumuUser) *CommentCreate {
	return cc.SetAuthorID(k.ID)
}

// SetArticleID sets the "article" edge to the Article entity by ID.
func (cc *CommentCreate) SetArticleID(id int) *CommentCreate {
	cc.mutation.SetArticleID(id)
	return cc
}

// SetNillableArticleID sets the "article" edge to the Article entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableArticleID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetArticleID(*id)
	}
	return cc
}

// SetArticle sets the "article" edge to the Article entity.
func (cc *CommentCreate) SetArticle(a *Article) *CommentCreate {
	return cc.SetArticleID(a.ID)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cc *CommentCreate) SetParentID(id int) *CommentCreate {
	cc.mutation.SetParentID(id)
	return cc
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableParentID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetParentID(*id)
	}
	return cc
}

// SetParent sets the "parent" edge to the Comment entity.
func (cc *CommentCreate) SetParent(c *Comment) *CommentCreate {
	return cc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cc *CommentCreate) AddChildIDs(ids ...int) *CommentCreate {
	cc.mutation.AddChildIDs(ids...)
	return cc
}

// AddChildren adds the "children" edges to the Comment entity.
func (cc *CommentCreate) AddChildren(c ...*Comment) *CommentCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChildIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.Kind(); !ok {
		v := comment.DefaultKind
		cc.mutation.SetKind(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required field \"state\"")}
	}
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New("ent: missing required field \"content\"")}
	}
	if _, ok := cc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New("ent: missing required field \"kind\"")}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldState,
		})
		_node.State = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := cc.mutation.Kind(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldKind,
		})
		_node.Kind = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := cc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.AuthorTable,
			Columns: []string{comment.AuthorColumn},
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
	if nodes := cc.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ArticleTable,
			Columns: []string{comment.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.article_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
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
		_node.parent_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
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
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
