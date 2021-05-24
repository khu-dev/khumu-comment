// Code generated by entc, DO NOT EDIT.

package khumuuser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/khu-dev/khumu-comment/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
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
func IDGT(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// StudentNumber applies equality check predicate on the "student_number" field. It's identical to StudentNumberEQ.
func StudentNumber(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentNumber), v))
	})
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// StudentNumberEQ applies the EQ predicate on the "student_number" field.
func StudentNumberEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberNEQ applies the NEQ predicate on the "student_number" field.
func StudentNumberNEQ(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberIn applies the In predicate on the "student_number" field.
func StudentNumberIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStudentNumber), v...))
	})
}

// StudentNumberNotIn applies the NotIn predicate on the "student_number" field.
func StudentNumberNotIn(vs ...string) predicate.KhumuUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KhumuUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStudentNumber), v...))
	})
}

// StudentNumberGT applies the GT predicate on the "student_number" field.
func StudentNumberGT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberGTE applies the GTE predicate on the "student_number" field.
func StudentNumberGTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberLT applies the LT predicate on the "student_number" field.
func StudentNumberLT(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberLTE applies the LTE predicate on the "student_number" field.
func StudentNumberLTE(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberContains applies the Contains predicate on the "student_number" field.
func StudentNumberContains(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberHasPrefix applies the HasPrefix predicate on the "student_number" field.
func StudentNumberHasPrefix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberHasSuffix applies the HasSuffix predicate on the "student_number" field.
func StudentNumberHasSuffix(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberIsNil applies the IsNil predicate on the "student_number" field.
func StudentNumberIsNil() predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStudentNumber)))
	})
}

// StudentNumberNotNil applies the NotNil predicate on the "student_number" field.
func StudentNumberNotNil() predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStudentNumber)))
	})
}

// StudentNumberEqualFold applies the EqualFold predicate on the "student_number" field.
func StudentNumberEqualFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStudentNumber), v))
	})
}

// StudentNumberContainsFold applies the ContainsFold predicate on the "student_number" field.
func StudentNumberContainsFold(v string) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStudentNumber), v))
	})
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActive), v))
	})
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsActive), v))
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsTable, CommentFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CommentsInverseTable, CommentFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArticles applies the HasEdge predicate on the "articles" edge.
func HasArticles() predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticlesTable, ArticleFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticlesTable, ArticlesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticlesWith applies the HasEdge predicate on the "articles" edge with a given conditions (other predicates).
func HasArticlesWith(preds ...predicate.Article) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticlesInverseTable, ArticleFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ArticlesTable, ArticlesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLike applies the HasEdge predicate on the "like" edge.
func HasLike() predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LikeTable, LikeCommentFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LikeTable, LikeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikeWith applies the HasEdge predicate on the "like" edge with a given conditions (other predicates).
func HasLikeWith(preds ...predicate.LikeComment) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LikeInverseTable, LikeCommentFieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LikeTable, LikeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KhumuUser) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KhumuUser) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
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
func Not(p predicate.KhumuUser) predicate.KhumuUser {
	return predicate.KhumuUser(func(s *sql.Selector) {
		p(s.Not())
	})
}