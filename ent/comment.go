// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Kind holds the value of the "kind" field.
	Kind string `json:"kind,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges            CommentEdges `json:"edges"`
	article_id       *int
	parent_id        *int
	author_id        *string
	study_article_id *int
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Author holds the value of the author edge.
	Author *KhumuUser `json:"author,omitempty"`
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// StudyArticle holds the value of the studyArticle edge.
	StudyArticle *StudyArticle `json:"studyArticle,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Comment `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Comment `json:"children,omitempty"`
	// Like holds the value of the like edge.
	Like []*LikeComment `json:"like,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) AuthorOrErr() (*KhumuUser, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: khumuuser.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ArticleOrErr() (*Article, error) {
	if e.loadedTypes[1] {
		if e.Article == nil {
			// The edge article was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: article.Label}
		}
		return e.Article, nil
	}
	return nil, &NotLoadedError{edge: "article"}
}

// StudyArticleOrErr returns the StudyArticle value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) StudyArticleOrErr() (*StudyArticle, error) {
	if e.loadedTypes[2] {
		if e.StudyArticle == nil {
			// The edge studyArticle was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: studyarticle.Label}
		}
		return e.StudyArticle, nil
	}
	return nil, &NotLoadedError{edge: "studyArticle"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) ParentOrErr() (*Comment, error) {
	if e.loadedTypes[3] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: comment.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) ChildrenOrErr() ([]*Comment, error) {
	if e.loadedTypes[4] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// LikeOrErr returns the Like value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) LikeOrErr() ([]*LikeComment, error) {
	if e.loadedTypes[5] {
		return e.Like, nil
	}
	return nil, &NotLoadedError{edge: "like"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			values[i] = new(sql.NullInt64)
		case comment.FieldState, comment.FieldContent, comment.FieldKind:
			values[i] = new(sql.NullString)
		case comment.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case comment.ForeignKeys[0]: // article_id
			values[i] = new(sql.NullInt64)
		case comment.ForeignKeys[1]: // parent_id
			values[i] = new(sql.NullInt64)
		case comment.ForeignKeys[2]: // author_id
			values[i] = new(sql.NullString)
		case comment.ForeignKeys[3]: // study_article_id
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				c.State = value.String
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				c.Kind = value.String
			}
		case comment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field article_id", value)
			} else if value.Valid {
				c.article_id = new(int)
				*c.article_id = int(value.Int64)
			}
		case comment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field parent_id", value)
			} else if value.Valid {
				c.parent_id = new(int)
				*c.parent_id = int(value.Int64)
			}
		case comment.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value.Valid {
				c.author_id = new(string)
				*c.author_id = value.String
			}
		case comment.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field study_article_id", value)
			} else if value.Valid {
				c.study_article_id = new(int)
				*c.study_article_id = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the Comment entity.
func (c *Comment) QueryAuthor() *KhumuUserQuery {
	return (&CommentClient{config: c.config}).QueryAuthor(c)
}

// QueryArticle queries the "article" edge of the Comment entity.
func (c *Comment) QueryArticle() *ArticleQuery {
	return (&CommentClient{config: c.config}).QueryArticle(c)
}

// QueryStudyArticle queries the "studyArticle" edge of the Comment entity.
func (c *Comment) QueryStudyArticle() *StudyArticleQuery {
	return (&CommentClient{config: c.config}).QueryStudyArticle(c)
}

// QueryParent queries the "parent" edge of the Comment entity.
func (c *Comment) QueryParent() *CommentQuery {
	return (&CommentClient{config: c.config}).QueryParent(c)
}

// QueryChildren queries the "children" edge of the Comment entity.
func (c *Comment) QueryChildren() *CommentQuery {
	return (&CommentClient{config: c.config}).QueryChildren(c)
}

// QueryLike queries the "like" edge of the Comment entity.
func (c *Comment) QueryLike() *LikeCommentQuery {
	return (&CommentClient{config: c.config}).QueryLike(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", state=")
	builder.WriteString(c.State)
	builder.WriteString(", content=")
	builder.WriteString(c.Content)
	builder.WriteString(", kind=")
	builder.WriteString(c.Kind)
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
