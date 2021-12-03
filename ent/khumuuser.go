// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
)

// KhumuUser is the model entity for the KhumuUser schema.
type KhumuUser struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KhumuUserQuery when eager-loading is set.
	Edges KhumuUserEdges `json:"edges"`
}

// KhumuUserEdges holds the relations/edges for other nodes in the graph.
type KhumuUserEdges struct {
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Articles holds the value of the articles edge.
	Articles []*Article `json:"articles,omitempty"`
	// StudyArticles holds the value of the studyArticles edge.
	StudyArticles []*StudyArticle `json:"studyArticles,omitempty"`
	// Like holds the value of the like edge.
	Like []*LikeComment `json:"like,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e KhumuUserEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[0] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// ArticlesOrErr returns the Articles value or an error if the edge
// was not loaded in eager-loading.
func (e KhumuUserEdges) ArticlesOrErr() ([]*Article, error) {
	if e.loadedTypes[1] {
		return e.Articles, nil
	}
	return nil, &NotLoadedError{edge: "articles"}
}

// StudyArticlesOrErr returns the StudyArticles value or an error if the edge
// was not loaded in eager-loading.
func (e KhumuUserEdges) StudyArticlesOrErr() ([]*StudyArticle, error) {
	if e.loadedTypes[2] {
		return e.StudyArticles, nil
	}
	return nil, &NotLoadedError{edge: "studyArticles"}
}

// LikeOrErr returns the Like value or an error if the edge
// was not loaded in eager-loading.
func (e KhumuUserEdges) LikeOrErr() ([]*LikeComment, error) {
	if e.loadedTypes[3] {
		return e.Like, nil
	}
	return nil, &NotLoadedError{edge: "like"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KhumuUser) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case khumuuser.FieldID, khumuuser.FieldNickname, khumuuser.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KhumuUser", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KhumuUser fields.
func (ku *KhumuUser) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case khumuuser.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ku.ID = value.String
			}
		case khumuuser.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				ku.Nickname = value.String
			}
		case khumuuser.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ku.Status = value.String
			}
		}
	}
	return nil
}

// QueryComments queries the "comments" edge of the KhumuUser entity.
func (ku *KhumuUser) QueryComments() *CommentQuery {
	return (&KhumuUserClient{config: ku.config}).QueryComments(ku)
}

// QueryArticles queries the "articles" edge of the KhumuUser entity.
func (ku *KhumuUser) QueryArticles() *ArticleQuery {
	return (&KhumuUserClient{config: ku.config}).QueryArticles(ku)
}

// QueryStudyArticles queries the "studyArticles" edge of the KhumuUser entity.
func (ku *KhumuUser) QueryStudyArticles() *StudyArticleQuery {
	return (&KhumuUserClient{config: ku.config}).QueryStudyArticles(ku)
}

// QueryLike queries the "like" edge of the KhumuUser entity.
func (ku *KhumuUser) QueryLike() *LikeCommentQuery {
	return (&KhumuUserClient{config: ku.config}).QueryLike(ku)
}

// Update returns a builder for updating this KhumuUser.
// Note that you need to call KhumuUser.Unwrap() before calling this method if this KhumuUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ku *KhumuUser) Update() *KhumuUserUpdateOne {
	return (&KhumuUserClient{config: ku.config}).UpdateOne(ku)
}

// Unwrap unwraps the KhumuUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ku *KhumuUser) Unwrap() *KhumuUser {
	tx, ok := ku.config.driver.(*txDriver)
	if !ok {
		panic("ent: KhumuUser is not a transactional entity")
	}
	ku.config.driver = tx.drv
	return ku
}

// String implements the fmt.Stringer.
func (ku *KhumuUser) String() string {
	var builder strings.Builder
	builder.WriteString("KhumuUser(")
	builder.WriteString(fmt.Sprintf("id=%v", ku.ID))
	builder.WriteString(", nickname=")
	builder.WriteString(ku.Nickname)
	builder.WriteString(", status=")
	builder.WriteString(ku.Status)
	builder.WriteByte(')')
	return builder.String()
}

// KhumuUsers is a parsable slice of KhumuUser.
type KhumuUsers []*KhumuUser

func (ku KhumuUsers) config(cfg config) {
	for _i := range ku {
		ku[_i].config = cfg
	}
}
