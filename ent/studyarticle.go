// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// StudyArticle is the model entity for the StudyArticle schema.
type StudyArticle struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudyArticleQuery when eager-loading is set.
	Edges     StudyArticleEdges `json:"edges"`
	author_id *string
}

// StudyArticleEdges holds the relations/edges for other nodes in the graph.
type StudyArticleEdges struct {
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Author holds the value of the author edge.
	Author *KhumuUser `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e StudyArticleEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[0] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudyArticleEdges) AuthorOrErr() (*KhumuUser, error) {
	if e.loadedTypes[1] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: khumuuser.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StudyArticle) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case studyarticle.FieldID:
			values[i] = new(sql.NullInt64)
		case studyarticle.ForeignKeys[0]: // author_id
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type StudyArticle", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StudyArticle fields.
func (sa *StudyArticle) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case studyarticle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sa.ID = int(value.Int64)
		case studyarticle.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value.Valid {
				sa.author_id = new(string)
				*sa.author_id = value.String
			}
		}
	}
	return nil
}

// QueryComments queries the "comments" edge of the StudyArticle entity.
func (sa *StudyArticle) QueryComments() *CommentQuery {
	return (&StudyArticleClient{config: sa.config}).QueryComments(sa)
}

// QueryAuthor queries the "author" edge of the StudyArticle entity.
func (sa *StudyArticle) QueryAuthor() *KhumuUserQuery {
	return (&StudyArticleClient{config: sa.config}).QueryAuthor(sa)
}

// Update returns a builder for updating this StudyArticle.
// Note that you need to call StudyArticle.Unwrap() before calling this method if this StudyArticle
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *StudyArticle) Update() *StudyArticleUpdateOne {
	return (&StudyArticleClient{config: sa.config}).UpdateOne(sa)
}

// Unwrap unwraps the StudyArticle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *StudyArticle) Unwrap() *StudyArticle {
	tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: StudyArticle is not a transactional entity")
	}
	sa.config.driver = tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *StudyArticle) String() string {
	var builder strings.Builder
	builder.WriteString("StudyArticle(")
	builder.WriteString(fmt.Sprintf("id=%v", sa.ID))
	builder.WriteByte(')')
	return builder.String()
}

// StudyArticles is a parsable slice of StudyArticle.
type StudyArticles []*StudyArticle

func (sa StudyArticles) config(cfg config) {
	for _i := range sa {
		sa[_i].config = cfg
	}
}