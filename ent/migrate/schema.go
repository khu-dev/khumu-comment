// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticleArticleColumns holds the columns for the "article_article" table.
	ArticleArticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "images", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
	}
	// ArticleArticleTable holds the schema information for the "article_article" table.
	ArticleArticleTable = &schema.Table{
		Name:       "article_article",
		Columns:    ArticleArticleColumns,
		PrimaryKey: []*schema.Column{ArticleArticleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_article_user_khumuuser_articles",
				Columns:    []*schema.Column{ArticleArticleColumns[4]},
				RefColumns: []*schema.Column{UserKhumuuserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BoardsColumns holds the columns for the "boards" table.
	BoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// BoardsTable holds the schema information for the "boards" table.
	BoardsTable = &schema.Table{
		Name:        "boards",
		Columns:     BoardsColumns,
		PrimaryKey:  []*schema.Column{BoardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CommentCommentColumns holds the columns for the "comment_comment" table.
	CommentCommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "state", Type: field.TypeString, Default: "exists"},
		{Name: "content", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString, Default: "anonymous"},
		{Name: "is_written_by_article_author", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "article_id", Type: field.TypeInt, Nullable: true},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
		{Name: "study_article_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentCommentTable holds the schema information for the "comment_comment" table.
	CommentCommentTable = &schema.Table{
		Name:       "comment_comment",
		Columns:    CommentCommentColumns,
		PrimaryKey: []*schema.Column{CommentCommentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_comment_article_article_comments",
				Columns:    []*schema.Column{CommentCommentColumns[6]},
				RefColumns: []*schema.Column{ArticleArticleColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_comment_comment_comment_children",
				Columns:    []*schema.Column{CommentCommentColumns[7]},
				RefColumns: []*schema.Column{CommentCommentColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_comment_user_khumuuser_comments",
				Columns:    []*schema.Column{CommentCommentColumns[8]},
				RefColumns: []*schema.Column{UserKhumuuserColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_comment_article_studyarticle_comments",
				Columns:    []*schema.Column{CommentCommentColumns[9]},
				RefColumns: []*schema.Column{ArticleStudyarticleColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserKhumuuserColumns holds the columns for the "user_khumuuser" table.
	UserKhumuuserColumns = []*schema.Column{
		{Name: "username", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "student_number", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeString, Default: "exists"},
	}
	// UserKhumuuserTable holds the schema information for the "user_khumuuser" table.
	UserKhumuuserTable = &schema.Table{
		Name:        "user_khumuuser",
		Columns:     UserKhumuuserColumns,
		PrimaryKey:  []*schema.Column{UserKhumuuserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CommentLikecommentColumns holds the columns for the "comment_likecomment" table.
	CommentLikecommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// CommentLikecommentTable holds the schema information for the "comment_likecomment" table.
	CommentLikecommentTable = &schema.Table{
		Name:       "comment_likecomment",
		Columns:    CommentLikecommentColumns,
		PrimaryKey: []*schema.Column{CommentLikecommentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_likecomment_comment_comment_like",
				Columns:    []*schema.Column{CommentLikecommentColumns[1]},
				RefColumns: []*schema.Column{CommentCommentColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_likecomment_user_khumuuser_like",
				Columns:    []*schema.Column{CommentLikecommentColumns[2]},
				RefColumns: []*schema.Column{UserKhumuuserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ArticleStudyarticleColumns holds the columns for the "article_studyarticle" table.
	ArticleStudyarticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
	}
	// ArticleStudyarticleTable holds the schema information for the "article_studyarticle" table.
	ArticleStudyarticleTable = &schema.Table{
		Name:       "article_studyarticle",
		Columns:    ArticleStudyarticleColumns,
		PrimaryKey: []*schema.Column{ArticleStudyarticleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_studyarticle_user_khumuuser_studyArticles",
				Columns:    []*schema.Column{ArticleStudyarticleColumns[1]},
				RefColumns: []*schema.Column{UserKhumuuserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticleArticleTable,
		BoardsTable,
		CommentCommentTable,
		UserKhumuuserTable,
		CommentLikecommentTable,
		ArticleStudyarticleTable,
	}
)

func init() {
	ArticleArticleTable.ForeignKeys[0].RefTable = UserKhumuuserTable
	ArticleArticleTable.Annotation = &entsql.Annotation{
		Table: "article_article",
	}
	CommentCommentTable.ForeignKeys[0].RefTable = ArticleArticleTable
	CommentCommentTable.ForeignKeys[1].RefTable = CommentCommentTable
	CommentCommentTable.ForeignKeys[2].RefTable = UserKhumuuserTable
	CommentCommentTable.ForeignKeys[3].RefTable = ArticleStudyarticleTable
	CommentCommentTable.Annotation = &entsql.Annotation{
		Table: "comment_comment",
	}
	UserKhumuuserTable.Annotation = &entsql.Annotation{
		Table: "user_khumuuser",
	}
	CommentLikecommentTable.ForeignKeys[0].RefTable = CommentCommentTable
	CommentLikecommentTable.ForeignKeys[1].RefTable = UserKhumuuserTable
	CommentLikecommentTable.Annotation = &entsql.Annotation{
		Table: "comment_likecomment",
	}
	ArticleStudyarticleTable.ForeignKeys[0].RefTable = UserKhumuuserTable
	ArticleStudyarticleTable.Annotation = &entsql.Annotation{
		Table: "article_studyarticle",
	}
}
