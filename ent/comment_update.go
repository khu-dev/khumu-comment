// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/predicate"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where adds a new predicate for the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetState sets the "state" field.
func (cu *CommentUpdate) SetState(s string) *CommentUpdate {
	cu.mutation.SetState(s)
	return cu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableState(s *string) *CommentUpdate {
	if s != nil {
		cu.SetState(*s)
	}
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (cu *CommentUpdate) SetAuthorID(id string) *CommentUpdate {
	cu.mutation.SetAuthorID(id)
	return cu
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableAuthorID(id *string) *CommentUpdate {
	if id != nil {
		cu = cu.SetAuthorID(*id)
	}
	return cu
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (cu *CommentUpdate) SetAuthor(k *KhumuUser) *CommentUpdate {
	return cu.SetAuthorID(k.ID)
}

// SetArticleID sets the "article" edge to the Article entity by ID.
func (cu *CommentUpdate) SetArticleID(id int) *CommentUpdate {
	cu.mutation.SetArticleID(id)
	return cu
}

// SetNillableArticleID sets the "article" edge to the Article entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableArticleID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetArticleID(*id)
	}
	return cu
}

// SetArticle sets the "article" edge to the Article entity.
func (cu *CommentUpdate) SetArticle(a *Article) *CommentUpdate {
	return cu.SetArticleID(a.ID)
}

// SetStudyArticleID sets the "studyArticle" edge to the StudyArticle entity by ID.
func (cu *CommentUpdate) SetStudyArticleID(id int) *CommentUpdate {
	cu.mutation.SetStudyArticleID(id)
	return cu
}

// SetNillableStudyArticleID sets the "studyArticle" edge to the StudyArticle entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableStudyArticleID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetStudyArticleID(*id)
	}
	return cu
}

// SetStudyArticle sets the "studyArticle" edge to the StudyArticle entity.
func (cu *CommentUpdate) SetStudyArticle(s *StudyArticle) *CommentUpdate {
	return cu.SetStudyArticleID(s.ID)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cu *CommentUpdate) SetParentID(id int) *CommentUpdate {
	cu.mutation.SetParentID(id)
	return cu
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableParentID(id *int) *CommentUpdate {
	if id != nil {
		cu = cu.SetParentID(*id)
	}
	return cu
}

// SetParent sets the "parent" edge to the Comment entity.
func (cu *CommentUpdate) SetParent(c *Comment) *CommentUpdate {
	return cu.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cu *CommentUpdate) AddChildIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddChildIDs(ids...)
	return cu
}

// AddChildren adds the "children" edges to the Comment entity.
func (cu *CommentUpdate) AddChildren(c ...*Comment) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChildIDs(ids...)
}

// AddLikeIDs adds the "like" edge to the LikeComment entity by IDs.
func (cu *CommentUpdate) AddLikeIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddLikeIDs(ids...)
	return cu
}

// AddLike adds the "like" edges to the LikeComment entity.
func (cu *CommentUpdate) AddLike(l ...*LikeComment) *CommentUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.AddLikeIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearAuthor clears the "author" edge to the KhumuUser entity.
func (cu *CommentUpdate) ClearAuthor() *CommentUpdate {
	cu.mutation.ClearAuthor()
	return cu
}

// ClearArticle clears the "article" edge to the Article entity.
func (cu *CommentUpdate) ClearArticle() *CommentUpdate {
	cu.mutation.ClearArticle()
	return cu
}

// ClearStudyArticle clears the "studyArticle" edge to the StudyArticle entity.
func (cu *CommentUpdate) ClearStudyArticle() *CommentUpdate {
	cu.mutation.ClearStudyArticle()
	return cu
}

// ClearParent clears the "parent" edge to the Comment entity.
func (cu *CommentUpdate) ClearParent() *CommentUpdate {
	cu.mutation.ClearParent()
	return cu
}

// ClearChildren clears all "children" edges to the Comment entity.
func (cu *CommentUpdate) ClearChildren() *CommentUpdate {
	cu.mutation.ClearChildren()
	return cu
}

// RemoveChildIDs removes the "children" edge to Comment entities by IDs.
func (cu *CommentUpdate) RemoveChildIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveChildIDs(ids...)
	return cu
}

// RemoveChildren removes "children" edges to Comment entities.
func (cu *CommentUpdate) RemoveChildren(c ...*Comment) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChildIDs(ids...)
}

// ClearLike clears all "like" edges to the LikeComment entity.
func (cu *CommentUpdate) ClearLike() *CommentUpdate {
	cu.mutation.ClearLike()
	return cu
}

// RemoveLikeIDs removes the "like" edge to LikeComment entities by IDs.
func (cu *CommentUpdate) RemoveLikeIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveLikeIDs(ids...)
	return cu
}

// RemoveLike removes "like" edges to LikeComment entities.
func (cu *CommentUpdate) RemoveLike(l ...*LikeComment) *CommentUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.RemoveLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldState,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if cu.mutation.AuthorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AuthorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ArticleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ArticleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.StudyArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.StudyArticleTable,
			Columns: []string{comment.StudyArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyarticle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StudyArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.StudyArticleTable,
			Columns: []string{comment.StudyArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyarticle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedLikeIDs(); len(nodes) > 0 && !cu.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetState sets the "state" field.
func (cuo *CommentUpdateOne) SetState(s string) *CommentUpdateOne {
	cuo.mutation.SetState(s)
	return cuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableState(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetState(*s)
	}
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetAuthorID sets the "author" edge to the KhumuUser entity by ID.
func (cuo *CommentUpdateOne) SetAuthorID(id string) *CommentUpdateOne {
	cuo.mutation.SetAuthorID(id)
	return cuo
}

// SetNillableAuthorID sets the "author" edge to the KhumuUser entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAuthorID(id *string) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetAuthorID(*id)
	}
	return cuo
}

// SetAuthor sets the "author" edge to the KhumuUser entity.
func (cuo *CommentUpdateOne) SetAuthor(k *KhumuUser) *CommentUpdateOne {
	return cuo.SetAuthorID(k.ID)
}

// SetArticleID sets the "article" edge to the Article entity by ID.
func (cuo *CommentUpdateOne) SetArticleID(id int) *CommentUpdateOne {
	cuo.mutation.SetArticleID(id)
	return cuo
}

// SetNillableArticleID sets the "article" edge to the Article entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableArticleID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetArticleID(*id)
	}
	return cuo
}

// SetArticle sets the "article" edge to the Article entity.
func (cuo *CommentUpdateOne) SetArticle(a *Article) *CommentUpdateOne {
	return cuo.SetArticleID(a.ID)
}

// SetStudyArticleID sets the "studyArticle" edge to the StudyArticle entity by ID.
func (cuo *CommentUpdateOne) SetStudyArticleID(id int) *CommentUpdateOne {
	cuo.mutation.SetStudyArticleID(id)
	return cuo
}

// SetNillableStudyArticleID sets the "studyArticle" edge to the StudyArticle entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableStudyArticleID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetStudyArticleID(*id)
	}
	return cuo
}

// SetStudyArticle sets the "studyArticle" edge to the StudyArticle entity.
func (cuo *CommentUpdateOne) SetStudyArticle(s *StudyArticle) *CommentUpdateOne {
	return cuo.SetStudyArticleID(s.ID)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cuo *CommentUpdateOne) SetParentID(id int) *CommentUpdateOne {
	cuo.mutation.SetParentID(id)
	return cuo
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableParentID(id *int) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetParentID(*id)
	}
	return cuo
}

// SetParent sets the "parent" edge to the Comment entity.
func (cuo *CommentUpdateOne) SetParent(c *Comment) *CommentUpdateOne {
	return cuo.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cuo *CommentUpdateOne) AddChildIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddChildIDs(ids...)
	return cuo
}

// AddChildren adds the "children" edges to the Comment entity.
func (cuo *CommentUpdateOne) AddChildren(c ...*Comment) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChildIDs(ids...)
}

// AddLikeIDs adds the "like" edge to the LikeComment entity by IDs.
func (cuo *CommentUpdateOne) AddLikeIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddLikeIDs(ids...)
	return cuo
}

// AddLike adds the "like" edges to the LikeComment entity.
func (cuo *CommentUpdateOne) AddLike(l ...*LikeComment) *CommentUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.AddLikeIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearAuthor clears the "author" edge to the KhumuUser entity.
func (cuo *CommentUpdateOne) ClearAuthor() *CommentUpdateOne {
	cuo.mutation.ClearAuthor()
	return cuo
}

// ClearArticle clears the "article" edge to the Article entity.
func (cuo *CommentUpdateOne) ClearArticle() *CommentUpdateOne {
	cuo.mutation.ClearArticle()
	return cuo
}

// ClearStudyArticle clears the "studyArticle" edge to the StudyArticle entity.
func (cuo *CommentUpdateOne) ClearStudyArticle() *CommentUpdateOne {
	cuo.mutation.ClearStudyArticle()
	return cuo
}

// ClearParent clears the "parent" edge to the Comment entity.
func (cuo *CommentUpdateOne) ClearParent() *CommentUpdateOne {
	cuo.mutation.ClearParent()
	return cuo
}

// ClearChildren clears all "children" edges to the Comment entity.
func (cuo *CommentUpdateOne) ClearChildren() *CommentUpdateOne {
	cuo.mutation.ClearChildren()
	return cuo
}

// RemoveChildIDs removes the "children" edge to Comment entities by IDs.
func (cuo *CommentUpdateOne) RemoveChildIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveChildIDs(ids...)
	return cuo
}

// RemoveChildren removes "children" edges to Comment entities.
func (cuo *CommentUpdateOne) RemoveChildren(c ...*Comment) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChildIDs(ids...)
}

// ClearLike clears all "like" edges to the LikeComment entity.
func (cuo *CommentUpdateOne) ClearLike() *CommentUpdateOne {
	cuo.mutation.ClearLike()
	return cuo
}

// RemoveLikeIDs removes the "like" edge to LikeComment entities by IDs.
func (cuo *CommentUpdateOne) RemoveLikeIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveLikeIDs(ids...)
	return cuo
}

// RemoveLike removes "like" edges to LikeComment entities.
func (cuo *CommentUpdateOne) RemoveLike(l ...*LikeComment) *CommentUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.RemoveLikeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldState,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if cuo.mutation.AuthorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AuthorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ArticleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ArticleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.StudyArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.StudyArticleTable,
			Columns: []string{comment.StudyArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyarticle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StudyArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.StudyArticleTable,
			Columns: []string{comment.StudyArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyarticle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedLikeIDs(); len(nodes) > 0 && !cuo.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.LikeTable,
			Columns: []string{comment.LikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: likecomment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
