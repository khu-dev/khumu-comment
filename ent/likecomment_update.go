// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/predicate"
)

// LikeCommentUpdate is the builder for updating LikeComment entities.
type LikeCommentUpdate struct {
	config
	hooks    []Hook
	mutation *LikeCommentMutation
}

// Where adds a new predicate for the LikeCommentUpdate builder.
func (lcu *LikeCommentUpdate) Where(ps ...predicate.LikeComment) *LikeCommentUpdate {
	lcu.mutation.predicates = append(lcu.mutation.predicates, ps...)
	return lcu
}

// SetLikedByID sets the "likedBy" edge to the KhumuUser entity by ID.
func (lcu *LikeCommentUpdate) SetLikedByID(id string) *LikeCommentUpdate {
	lcu.mutation.SetLikedByID(id)
	return lcu
}

// SetNillableLikedByID sets the "likedBy" edge to the KhumuUser entity by ID if the given value is not nil.
func (lcu *LikeCommentUpdate) SetNillableLikedByID(id *string) *LikeCommentUpdate {
	if id != nil {
		lcu = lcu.SetLikedByID(*id)
	}
	return lcu
}

// SetLikedBy sets the "likedBy" edge to the KhumuUser entity.
func (lcu *LikeCommentUpdate) SetLikedBy(k *KhumuUser) *LikeCommentUpdate {
	return lcu.SetLikedByID(k.ID)
}

// SetAboutID sets the "about" edge to the Comment entity by ID.
func (lcu *LikeCommentUpdate) SetAboutID(id int) *LikeCommentUpdate {
	lcu.mutation.SetAboutID(id)
	return lcu
}

// SetNillableAboutID sets the "about" edge to the Comment entity by ID if the given value is not nil.
func (lcu *LikeCommentUpdate) SetNillableAboutID(id *int) *LikeCommentUpdate {
	if id != nil {
		lcu = lcu.SetAboutID(*id)
	}
	return lcu
}

// SetAbout sets the "about" edge to the Comment entity.
func (lcu *LikeCommentUpdate) SetAbout(c *Comment) *LikeCommentUpdate {
	return lcu.SetAboutID(c.ID)
}

// Mutation returns the LikeCommentMutation object of the builder.
func (lcu *LikeCommentUpdate) Mutation() *LikeCommentMutation {
	return lcu.mutation
}

// ClearLikedBy clears the "likedBy" edge to the KhumuUser entity.
func (lcu *LikeCommentUpdate) ClearLikedBy() *LikeCommentUpdate {
	lcu.mutation.ClearLikedBy()
	return lcu
}

// ClearAbout clears the "about" edge to the Comment entity.
func (lcu *LikeCommentUpdate) ClearAbout() *LikeCommentUpdate {
	lcu.mutation.ClearAbout()
	return lcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lcu *LikeCommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lcu.hooks) == 0 {
		affected, err = lcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lcu.mutation = mutation
			affected, err = lcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lcu.hooks) - 1; i >= 0; i-- {
			mut = lcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lcu *LikeCommentUpdate) SaveX(ctx context.Context) int {
	affected, err := lcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lcu *LikeCommentUpdate) Exec(ctx context.Context) error {
	_, err := lcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcu *LikeCommentUpdate) ExecX(ctx context.Context) {
	if err := lcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lcu *LikeCommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   likecomment.Table,
			Columns: likecomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likecomment.FieldID,
			},
		},
	}
	if ps := lcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lcu.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.LikedByTable,
			Columns: []string{likecomment.LikedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: khumuuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.LikedByTable,
			Columns: []string{likecomment.LikedByColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lcu.mutation.AboutCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.AboutTable,
			Columns: []string{likecomment.AboutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.AboutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.AboutTable,
			Columns: []string{likecomment.AboutColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{likecomment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LikeCommentUpdateOne is the builder for updating a single LikeComment entity.
type LikeCommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeCommentMutation
}

// SetLikedByID sets the "likedBy" edge to the KhumuUser entity by ID.
func (lcuo *LikeCommentUpdateOne) SetLikedByID(id string) *LikeCommentUpdateOne {
	lcuo.mutation.SetLikedByID(id)
	return lcuo
}

// SetNillableLikedByID sets the "likedBy" edge to the KhumuUser entity by ID if the given value is not nil.
func (lcuo *LikeCommentUpdateOne) SetNillableLikedByID(id *string) *LikeCommentUpdateOne {
	if id != nil {
		lcuo = lcuo.SetLikedByID(*id)
	}
	return lcuo
}

// SetLikedBy sets the "likedBy" edge to the KhumuUser entity.
func (lcuo *LikeCommentUpdateOne) SetLikedBy(k *KhumuUser) *LikeCommentUpdateOne {
	return lcuo.SetLikedByID(k.ID)
}

// SetAboutID sets the "about" edge to the Comment entity by ID.
func (lcuo *LikeCommentUpdateOne) SetAboutID(id int) *LikeCommentUpdateOne {
	lcuo.mutation.SetAboutID(id)
	return lcuo
}

// SetNillableAboutID sets the "about" edge to the Comment entity by ID if the given value is not nil.
func (lcuo *LikeCommentUpdateOne) SetNillableAboutID(id *int) *LikeCommentUpdateOne {
	if id != nil {
		lcuo = lcuo.SetAboutID(*id)
	}
	return lcuo
}

// SetAbout sets the "about" edge to the Comment entity.
func (lcuo *LikeCommentUpdateOne) SetAbout(c *Comment) *LikeCommentUpdateOne {
	return lcuo.SetAboutID(c.ID)
}

// Mutation returns the LikeCommentMutation object of the builder.
func (lcuo *LikeCommentUpdateOne) Mutation() *LikeCommentMutation {
	return lcuo.mutation
}

// ClearLikedBy clears the "likedBy" edge to the KhumuUser entity.
func (lcuo *LikeCommentUpdateOne) ClearLikedBy() *LikeCommentUpdateOne {
	lcuo.mutation.ClearLikedBy()
	return lcuo
}

// ClearAbout clears the "about" edge to the Comment entity.
func (lcuo *LikeCommentUpdateOne) ClearAbout() *LikeCommentUpdateOne {
	lcuo.mutation.ClearAbout()
	return lcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lcuo *LikeCommentUpdateOne) Select(field string, fields ...string) *LikeCommentUpdateOne {
	lcuo.fields = append([]string{field}, fields...)
	return lcuo
}

// Save executes the query and returns the updated LikeComment entity.
func (lcuo *LikeCommentUpdateOne) Save(ctx context.Context) (*LikeComment, error) {
	var (
		err  error
		node *LikeComment
	)
	if len(lcuo.hooks) == 0 {
		node, err = lcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lcuo.mutation = mutation
			node, err = lcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lcuo.hooks) - 1; i >= 0; i-- {
			mut = lcuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lcuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (lcuo *LikeCommentUpdateOne) SaveX(ctx context.Context) *LikeComment {
	node, err := lcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lcuo *LikeCommentUpdateOne) Exec(ctx context.Context) error {
	_, err := lcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcuo *LikeCommentUpdateOne) ExecX(ctx context.Context) {
	if err := lcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lcuo *LikeCommentUpdateOne) sqlSave(ctx context.Context) (_node *LikeComment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   likecomment.Table,
			Columns: likecomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: likecomment.FieldID,
			},
		},
	}
	id, ok := lcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LikeComment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := lcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, likecomment.FieldID)
		for _, f := range fields {
			if !likecomment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != likecomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lcuo.mutation.LikedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.LikedByTable,
			Columns: []string{likecomment.LikedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: khumuuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.LikedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.LikedByTable,
			Columns: []string{likecomment.LikedByColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lcuo.mutation.AboutCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.AboutTable,
			Columns: []string{likecomment.AboutColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.AboutIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   likecomment.AboutTable,
			Columns: []string{likecomment.AboutColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &LikeComment{config: lcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{likecomment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
