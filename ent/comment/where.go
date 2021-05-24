// Code generated by entc, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/khu-dev/khumu-comment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Kind applies equality check predicate on the "kind" field. It's identical to KindEQ.
func Kind(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKind), v))
	})
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKind), v...))
	})
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...string) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKind), v...))
	})
}

// KindGT applies the GT predicate on the "kind" field.
func KindGT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKind), v))
	})
}

// KindGTE applies the GTE predicate on the "kind" field.
func KindGTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKind), v))
	})
}

// KindLT applies the LT predicate on the "kind" field.
func KindLT(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKind), v))
	})
}

// KindLTE applies the LTE predicate on the "kind" field.
func KindLTE(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKind), v))
	})
}

// KindContains applies the Contains predicate on the "kind" field.
func KindContains(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKind), v))
	})
}

// KindHasPrefix applies the HasPrefix predicate on the "kind" field.
func KindHasPrefix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKind), v))
	})
}

// KindHasSuffix applies the HasSuffix predicate on the "kind" field.
func KindHasSuffix(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKind), v))
	})
}

// KindEqualFold applies the EqualFold predicate on the "kind" field.
func KindEqualFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKind), v))
	})
}

// KindContainsFold applies the ContainsFold predicate on the "kind" field.
func KindContainsFold(v string) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKind), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Comment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, KhumuUserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.KhumuUser) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, KhumuUserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArticle applies the HasEdge predicate on the "article" edge.
func HasArticle() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleWith applies the HasEdge predicate on the "article" edge with a given conditions (other predicates).
func HasArticleWith(preds ...predicate.Article) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChildrenTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		p(s.Not())
	})
}
