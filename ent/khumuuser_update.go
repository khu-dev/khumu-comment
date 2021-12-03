// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

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

// KhumuUserUpdate is the builder for updating KhumuUser entities.
type KhumuUserUpdate struct {
	config
	hooks    []Hook
	mutation *KhumuUserMutation
}

// Where adds a new predicate for the KhumuUserUpdate builder.
func (kuu *KhumuUserUpdate) Where(ps ...predicate.KhumuUser) *KhumuUserUpdate {
	kuu.mutation.predicates = append(kuu.mutation.predicates, ps...)
	return kuu
}

// SetNickname sets the "nickname" field.
func (kuu *KhumuUserUpdate) SetNickname(s string) *KhumuUserUpdate {
	kuu.mutation.SetNickname(s)
	return kuu
}

// SetStatus sets the "status" field.
func (kuu *KhumuUserUpdate) SetStatus(s string) *KhumuUserUpdate {
	kuu.mutation.SetStatus(s)
	return kuu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (kuu *KhumuUserUpdate) SetNillableStatus(s *string) *KhumuUserUpdate {
	if s != nil {
		kuu.SetStatus(*s)
	}
	return kuu
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (kuu *KhumuUserUpdate) AddCommentIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.AddCommentIDs(ids...)
	return kuu
}

// AddComments adds the "comments" edges to the Comment entity.
func (kuu *KhumuUserUpdate) AddComments(c ...*Comment) *KhumuUserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return kuu.AddCommentIDs(ids...)
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (kuu *KhumuUserUpdate) AddArticleIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.AddArticleIDs(ids...)
	return kuu
}

// AddArticles adds the "articles" edges to the Article entity.
func (kuu *KhumuUserUpdate) AddArticles(a ...*Article) *KhumuUserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return kuu.AddArticleIDs(ids...)
}

// AddStudyArticleIDs adds the "studyArticles" edge to the StudyArticle entity by IDs.
func (kuu *KhumuUserUpdate) AddStudyArticleIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.AddStudyArticleIDs(ids...)
	return kuu
}

// AddStudyArticles adds the "studyArticles" edges to the StudyArticle entity.
func (kuu *KhumuUserUpdate) AddStudyArticles(s ...*StudyArticle) *KhumuUserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return kuu.AddStudyArticleIDs(ids...)
}

// AddLikeIDs adds the "like" edge to the LikeComment entity by IDs.
func (kuu *KhumuUserUpdate) AddLikeIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.AddLikeIDs(ids...)
	return kuu
}

// AddLike adds the "like" edges to the LikeComment entity.
func (kuu *KhumuUserUpdate) AddLike(l ...*LikeComment) *KhumuUserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return kuu.AddLikeIDs(ids...)
}

// Mutation returns the KhumuUserMutation object of the builder.
func (kuu *KhumuUserUpdate) Mutation() *KhumuUserMutation {
	return kuu.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (kuu *KhumuUserUpdate) ClearComments() *KhumuUserUpdate {
	kuu.mutation.ClearComments()
	return kuu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (kuu *KhumuUserUpdate) RemoveCommentIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.RemoveCommentIDs(ids...)
	return kuu
}

// RemoveComments removes "comments" edges to Comment entities.
func (kuu *KhumuUserUpdate) RemoveComments(c ...*Comment) *KhumuUserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return kuu.RemoveCommentIDs(ids...)
}

// ClearArticles clears all "articles" edges to the Article entity.
func (kuu *KhumuUserUpdate) ClearArticles() *KhumuUserUpdate {
	kuu.mutation.ClearArticles()
	return kuu
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (kuu *KhumuUserUpdate) RemoveArticleIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.RemoveArticleIDs(ids...)
	return kuu
}

// RemoveArticles removes "articles" edges to Article entities.
func (kuu *KhumuUserUpdate) RemoveArticles(a ...*Article) *KhumuUserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return kuu.RemoveArticleIDs(ids...)
}

// ClearStudyArticles clears all "studyArticles" edges to the StudyArticle entity.
func (kuu *KhumuUserUpdate) ClearStudyArticles() *KhumuUserUpdate {
	kuu.mutation.ClearStudyArticles()
	return kuu
}

// RemoveStudyArticleIDs removes the "studyArticles" edge to StudyArticle entities by IDs.
func (kuu *KhumuUserUpdate) RemoveStudyArticleIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.RemoveStudyArticleIDs(ids...)
	return kuu
}

// RemoveStudyArticles removes "studyArticles" edges to StudyArticle entities.
func (kuu *KhumuUserUpdate) RemoveStudyArticles(s ...*StudyArticle) *KhumuUserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return kuu.RemoveStudyArticleIDs(ids...)
}

// ClearLike clears all "like" edges to the LikeComment entity.
func (kuu *KhumuUserUpdate) ClearLike() *KhumuUserUpdate {
	kuu.mutation.ClearLike()
	return kuu
}

// RemoveLikeIDs removes the "like" edge to LikeComment entities by IDs.
func (kuu *KhumuUserUpdate) RemoveLikeIDs(ids ...int) *KhumuUserUpdate {
	kuu.mutation.RemoveLikeIDs(ids...)
	return kuu
}

// RemoveLike removes "like" edges to LikeComment entities.
func (kuu *KhumuUserUpdate) RemoveLike(l ...*LikeComment) *KhumuUserUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return kuu.RemoveLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kuu *KhumuUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(kuu.hooks) == 0 {
		affected, err = kuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KhumuUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kuu.mutation = mutation
			affected, err = kuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(kuu.hooks) - 1; i >= 0; i-- {
			mut = kuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuu *KhumuUserUpdate) SaveX(ctx context.Context) int {
	affected, err := kuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kuu *KhumuUserUpdate) Exec(ctx context.Context) error {
	_, err := kuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuu *KhumuUserUpdate) ExecX(ctx context.Context) {
	if err := kuu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kuu *KhumuUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   khumuuser.Table,
			Columns: khumuuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: khumuuser.FieldID,
			},
		},
	}
	if ps := kuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuu.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: khumuuser.FieldNickname,
		})
	}
	if value, ok := kuu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: khumuuser.FieldStatus,
		})
	}
	if kuu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if nodes := kuu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !kuu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if nodes := kuu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if kuu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
	if nodes := kuu.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !kuu.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuu.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
	if kuu.mutation.StudyArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
	if nodes := kuu.mutation.RemovedStudyArticlesIDs(); len(nodes) > 0 && !kuu.mutation.StudyArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuu.mutation.StudyArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
	if kuu.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	if nodes := kuu.mutation.RemovedLikeIDs(); len(nodes) > 0 && !kuu.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	if nodes := kuu.mutation.LikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, kuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{khumuuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// KhumuUserUpdateOne is the builder for updating a single KhumuUser entity.
type KhumuUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KhumuUserMutation
}

// SetNickname sets the "nickname" field.
func (kuuo *KhumuUserUpdateOne) SetNickname(s string) *KhumuUserUpdateOne {
	kuuo.mutation.SetNickname(s)
	return kuuo
}

// SetStatus sets the "status" field.
func (kuuo *KhumuUserUpdateOne) SetStatus(s string) *KhumuUserUpdateOne {
	kuuo.mutation.SetStatus(s)
	return kuuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (kuuo *KhumuUserUpdateOne) SetNillableStatus(s *string) *KhumuUserUpdateOne {
	if s != nil {
		kuuo.SetStatus(*s)
	}
	return kuuo
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (kuuo *KhumuUserUpdateOne) AddCommentIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.AddCommentIDs(ids...)
	return kuuo
}

// AddComments adds the "comments" edges to the Comment entity.
func (kuuo *KhumuUserUpdateOne) AddComments(c ...*Comment) *KhumuUserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return kuuo.AddCommentIDs(ids...)
}

// AddArticleIDs adds the "articles" edge to the Article entity by IDs.
func (kuuo *KhumuUserUpdateOne) AddArticleIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.AddArticleIDs(ids...)
	return kuuo
}

// AddArticles adds the "articles" edges to the Article entity.
func (kuuo *KhumuUserUpdateOne) AddArticles(a ...*Article) *KhumuUserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return kuuo.AddArticleIDs(ids...)
}

// AddStudyArticleIDs adds the "studyArticles" edge to the StudyArticle entity by IDs.
func (kuuo *KhumuUserUpdateOne) AddStudyArticleIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.AddStudyArticleIDs(ids...)
	return kuuo
}

// AddStudyArticles adds the "studyArticles" edges to the StudyArticle entity.
func (kuuo *KhumuUserUpdateOne) AddStudyArticles(s ...*StudyArticle) *KhumuUserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return kuuo.AddStudyArticleIDs(ids...)
}

// AddLikeIDs adds the "like" edge to the LikeComment entity by IDs.
func (kuuo *KhumuUserUpdateOne) AddLikeIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.AddLikeIDs(ids...)
	return kuuo
}

// AddLike adds the "like" edges to the LikeComment entity.
func (kuuo *KhumuUserUpdateOne) AddLike(l ...*LikeComment) *KhumuUserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return kuuo.AddLikeIDs(ids...)
}

// Mutation returns the KhumuUserMutation object of the builder.
func (kuuo *KhumuUserUpdateOne) Mutation() *KhumuUserMutation {
	return kuuo.mutation
}

// ClearComments clears all "comments" edges to the Comment entity.
func (kuuo *KhumuUserUpdateOne) ClearComments() *KhumuUserUpdateOne {
	kuuo.mutation.ClearComments()
	return kuuo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (kuuo *KhumuUserUpdateOne) RemoveCommentIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.RemoveCommentIDs(ids...)
	return kuuo
}

// RemoveComments removes "comments" edges to Comment entities.
func (kuuo *KhumuUserUpdateOne) RemoveComments(c ...*Comment) *KhumuUserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return kuuo.RemoveCommentIDs(ids...)
}

// ClearArticles clears all "articles" edges to the Article entity.
func (kuuo *KhumuUserUpdateOne) ClearArticles() *KhumuUserUpdateOne {
	kuuo.mutation.ClearArticles()
	return kuuo
}

// RemoveArticleIDs removes the "articles" edge to Article entities by IDs.
func (kuuo *KhumuUserUpdateOne) RemoveArticleIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.RemoveArticleIDs(ids...)
	return kuuo
}

// RemoveArticles removes "articles" edges to Article entities.
func (kuuo *KhumuUserUpdateOne) RemoveArticles(a ...*Article) *KhumuUserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return kuuo.RemoveArticleIDs(ids...)
}

// ClearStudyArticles clears all "studyArticles" edges to the StudyArticle entity.
func (kuuo *KhumuUserUpdateOne) ClearStudyArticles() *KhumuUserUpdateOne {
	kuuo.mutation.ClearStudyArticles()
	return kuuo
}

// RemoveStudyArticleIDs removes the "studyArticles" edge to StudyArticle entities by IDs.
func (kuuo *KhumuUserUpdateOne) RemoveStudyArticleIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.RemoveStudyArticleIDs(ids...)
	return kuuo
}

// RemoveStudyArticles removes "studyArticles" edges to StudyArticle entities.
func (kuuo *KhumuUserUpdateOne) RemoveStudyArticles(s ...*StudyArticle) *KhumuUserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return kuuo.RemoveStudyArticleIDs(ids...)
}

// ClearLike clears all "like" edges to the LikeComment entity.
func (kuuo *KhumuUserUpdateOne) ClearLike() *KhumuUserUpdateOne {
	kuuo.mutation.ClearLike()
	return kuuo
}

// RemoveLikeIDs removes the "like" edge to LikeComment entities by IDs.
func (kuuo *KhumuUserUpdateOne) RemoveLikeIDs(ids ...int) *KhumuUserUpdateOne {
	kuuo.mutation.RemoveLikeIDs(ids...)
	return kuuo
}

// RemoveLike removes "like" edges to LikeComment entities.
func (kuuo *KhumuUserUpdateOne) RemoveLike(l ...*LikeComment) *KhumuUserUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return kuuo.RemoveLikeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuuo *KhumuUserUpdateOne) Select(field string, fields ...string) *KhumuUserUpdateOne {
	kuuo.fields = append([]string{field}, fields...)
	return kuuo
}

// Save executes the query and returns the updated KhumuUser entity.
func (kuuo *KhumuUserUpdateOne) Save(ctx context.Context) (*KhumuUser, error) {
	var (
		err  error
		node *KhumuUser
	)
	if len(kuuo.hooks) == 0 {
		node, err = kuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KhumuUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kuuo.mutation = mutation
			node, err = kuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuuo.hooks) - 1; i >= 0; i-- {
			mut = kuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuuo *KhumuUserUpdateOne) SaveX(ctx context.Context) *KhumuUser {
	node, err := kuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuuo *KhumuUserUpdateOne) Exec(ctx context.Context) error {
	_, err := kuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuuo *KhumuUserUpdateOne) ExecX(ctx context.Context) {
	if err := kuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (kuuo *KhumuUserUpdateOne) sqlSave(ctx context.Context) (_node *KhumuUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   khumuuser.Table,
			Columns: khumuuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: khumuuser.FieldID,
			},
		},
	}
	id, ok := kuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing KhumuUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := kuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, khumuuser.FieldID)
		for _, f := range fields {
			if !khumuuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != khumuuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuuo.mutation.Nickname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: khumuuser.FieldNickname,
		})
	}
	if value, ok := kuuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: khumuuser.FieldStatus,
		})
	}
	if kuuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if nodes := kuuo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !kuuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if nodes := kuuo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.CommentsTable,
			Columns: []string{khumuuser.CommentsColumn},
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
	if kuuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
	if nodes := kuuo.mutation.RemovedArticlesIDs(); len(nodes) > 0 && !kuuo.mutation.ArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuuo.mutation.ArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.ArticlesTable,
			Columns: []string{khumuuser.ArticlesColumn},
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
	if kuuo.mutation.StudyArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
	if nodes := kuuo.mutation.RemovedStudyArticlesIDs(); len(nodes) > 0 && !kuuo.mutation.StudyArticlesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kuuo.mutation.StudyArticlesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.StudyArticlesTable,
			Columns: []string{khumuuser.StudyArticlesColumn},
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
	if kuuo.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	if nodes := kuuo.mutation.RemovedLikeIDs(); len(nodes) > 0 && !kuuo.mutation.LikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	if nodes := kuuo.mutation.LikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   khumuuser.LikeTable,
			Columns: []string{khumuuser.LikeColumn},
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
	_node = &KhumuUser{config: kuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{khumuuser.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
