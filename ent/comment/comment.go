// Code generated by entc, DO NOT EDIT.

package comment

import (
	"time"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldIsWrittenByArticleAuthor holds the string denoting the is_written_by_article_author field in the database.
	FieldIsWrittenByArticleAuthor = "is_written_by_article_author"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeArticle holds the string denoting the article edge name in mutations.
	EdgeArticle = "article"
	// EdgeStudyArticle holds the string denoting the studyarticle edge name in mutations.
	EdgeStudyArticle = "studyArticle"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeLike holds the string denoting the like edge name in mutations.
	EdgeLike = "like"
	// KhumuUserFieldID holds the string denoting the ID field of the KhumuUser.
	KhumuUserFieldID = "username"
	// Table holds the table name of the comment in the database.
	Table = "comment_comment"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "comment_comment"
	// AuthorInverseTable is the table name for the KhumuUser entity.
	// It exists in this package in order to avoid circular dependency with the "khumuuser" package.
	AuthorInverseTable = "user_khumuuser"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "author_id"
	// ArticleTable is the table the holds the article relation/edge.
	ArticleTable = "comment_comment"
	// ArticleInverseTable is the table name for the Article entity.
	// It exists in this package in order to avoid circular dependency with the "article" package.
	ArticleInverseTable = "article_article"
	// ArticleColumn is the table column denoting the article relation/edge.
	ArticleColumn = "article_id"
	// StudyArticleTable is the table the holds the studyArticle relation/edge.
	StudyArticleTable = "comment_comment"
	// StudyArticleInverseTable is the table name for the StudyArticle entity.
	// It exists in this package in order to avoid circular dependency with the "studyarticle" package.
	StudyArticleInverseTable = "article_studyarticle"
	// StudyArticleColumn is the table column denoting the studyArticle relation/edge.
	StudyArticleColumn = "study_article_id"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "comment_comment"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "comment_comment"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// LikeTable is the table the holds the like relation/edge.
	LikeTable = "comment_likecomment"
	// LikeInverseTable is the table name for the LikeComment entity.
	// It exists in this package in order to avoid circular dependency with the "likecomment" package.
	LikeInverseTable = "comment_likecomment"
	// LikeColumn is the table column denoting the like relation/edge.
	LikeColumn = "comment_id"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldContent,
	FieldKind,
	FieldIsWrittenByArticleAuthor,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comment_comment"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"article_id",
	"parent_id",
	"author_id",
	"study_article_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
	// DefaultKind holds the default value on creation for the "kind" field.
	DefaultKind string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
