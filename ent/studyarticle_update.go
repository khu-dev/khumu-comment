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
	"github.com/khu-dev/khumu-comment/ent/predicate"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// StudyArticleUpdate is the builder for updating StudyArticle entities.
type StudyArticleUpdate struct {
	config
	hooks    []Hook
	mutation *StudyArticleMutation
}

// Where adds a new predicate for the StudyArticleUpdate builder.
func (sau *StudyArticleUpdate) Where(ps ...predicate.StudyArticle) *StudyArticleUpdate {
	sau.mutation.predicates = append(sau.mutation.predicates, ps...)
	return sau
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (sau *StudyArticleUpdate) AddCommentIDs(ids ...int) *StudyArticleUpdate {
	sau.mutation.AddCommentIDs(ids...)
	return sau
}

// AddComments adds the "comments" edges to the Comment entity.
func (sau *StudyArticleUpdate) AddComments(c ...*Comment) *StudyArticleUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sau.AddCommentIDs(ids...)
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (sau *StudyArticleUpdate) SetAuthorID(id string) *StudyArticleUpdate {
	sau.mutation.SetAuthorID(id)
	return sau
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (sau *StudyArticleUpdate) SetNillableAuthorID(id *string) *StudyArticleUpdate {
	if id != nil {
		sau = sau.SetAuthorID(*id)
	}
	return sau
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (sau *StudyArticleUpdate) SetAuthor(k *KhumuUser) *StudyArticleUpdate {
	return sau.SetAuthorID(k.ID)
}

// Mutation returns the StudyArticleMutation object of the builder.
func (sau *StudyArticleUpdate) Mutation() *StudyArticleMutation {
	return sau.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (sau *StudyArticleUpdate) ClearComments() *StudyArticleUpdate {
	sau.mutation.ClearComments()
	return sau
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (sau *StudyArticleUpdate) RemoveCommentIDs(ids ...int) *StudyArticleUpdate {
	sau.mutation.RemoveCommentIDs(ids...)
	return sau
}

// RemoveComments removes "comments" edges to Comment entities.
func (sau *StudyArticleUpdate) RemoveComments(c ...*Comment) *StudyArticleUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sau.RemoveCommentIDs(ids...)
}

// ClearAuthor clears the "author" edge to the KhumuUser entity.
func (sau *StudyArticleUpdate) ClearAuthor() *StudyArticleUpdate {
	sau.mutation.ClearAuthor()
	return sau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sau *StudyArticleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sau.hooks) == 0 {
		affected, err = sau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sau.mutation = mutation
			affected, err = sau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sau.hooks) - 1; i >= 0; i-- {
			mut = sau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sau *StudyArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := sau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sau *StudyArticleUpdate) Exec(ctx context.Context) error {
	_, err := sau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sau *StudyArticleUpdate) ExecX(ctx context.Context) {
	if err := sau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sau *StudyArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyarticle.Table,
			Columns: studyarticle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyarticle.FieldID,
			},
		},
	}
	if ps := sau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if sau.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !sau.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sau.mutation.AuthorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sau.mutation.AuthorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studyarticle.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StudyArticleUpdateOne is the builder for updating a single StudyArticle entity.
type StudyArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudyArticleMutation
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (sauo *StudyArticleUpdateOne) AddCommentIDs(ids ...int) *StudyArticleUpdateOne {
	sauo.mutation.AddCommentIDs(ids...)
	return sauo
}

// AddComments adds the "comments" edges to the Comment entity.
func (sauo *StudyArticleUpdateOne) AddComments(c ...*Comment) *StudyArticleUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sauo.AddCommentIDs(ids...)
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (sauo *StudyArticleUpdateOne) SetAuthorID(id string) *StudyArticleUpdateOne {
	sauo.mutation.SetAuthorID(id)
	return sauo
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (sauo *StudyArticleUpdateOne) SetNillableAuthorID(id *string) *StudyArticleUpdateOne {
	if id != nil {
		sauo = sauo.SetAuthorID(*id)
	}
	return sauo
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (sauo *StudyArticleUpdateOne) SetAuthor(k *KhumuUser) *StudyArticleUpdateOne {
	return sauo.SetAuthorID(k.ID)
}

// Mutation returns the StudyArticleMutation object of the builder.
func (sauo *StudyArticleUpdateOne) Mutation() *StudyArticleMutation {
	return sauo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (sauo *StudyArticleUpdateOne) ClearComments() *StudyArticleUpdateOne {
	sauo.mutation.ClearComments()
	return sauo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (sauo *StudyArticleUpdateOne) RemoveCommentIDs(ids ...int) *StudyArticleUpdateOne {
	sauo.mutation.RemoveCommentIDs(ids...)
	return sauo
}

// RemoveComments removes "comments" edges to Comment entities.
func (sauo *StudyArticleUpdateOne) RemoveComments(c ...*Comment) *StudyArticleUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sauo.RemoveCommentIDs(ids...)
}

// ClearAuthor clears the "author" edge to the KhumuUser entity.
func (sauo *StudyArticleUpdateOne) ClearAuthor() *StudyArticleUpdateOne {
	sauo.mutation.ClearAuthor()
	return sauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sauo *StudyArticleUpdateOne) Select(field string, fields ...string) *StudyArticleUpdateOne {
	sauo.fields = append([]string{field}, fields...)
	return sauo
}

// Save executes the query and returns the updated StudyArticle entity.
func (sauo *StudyArticleUpdateOne) Save(ctx context.Context) (*StudyArticle, error) {
	var (
		err  error
		node *StudyArticle
	)
	if len(sauo.hooks) == 0 {
		node, err = sauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudyArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sauo.mutation = mutation
			node, err = sauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sauo.hooks) - 1; i >= 0; i-- {
			mut = sauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sauo *StudyArticleUpdateOne) SaveX(ctx context.Context) *StudyArticle {
	node, err := sauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sauo *StudyArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := sauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sauo *StudyArticleUpdateOne) ExecX(ctx context.Context) {
	if err := sauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sauo *StudyArticleUpdateOne) sqlSave(ctx context.Context) (_node *StudyArticle, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyarticle.Table,
			Columns: studyarticle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyarticle.FieldID,
			},
		},
	}
	id, ok := sauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing StudyArticle.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studyarticle.FieldID)
		for _, f := range fields {
			if !studyarticle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != studyarticle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if sauo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !sauo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sauo.mutation.AuthorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sauo.mutation.AuthorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &StudyArticle{config: sauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{studyarticle.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}