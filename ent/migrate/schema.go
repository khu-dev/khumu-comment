// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "articles_khumu_users_articles",
				Columns:    []*schema.Column{ArticlesColumns[2]},
				RefColumns: []*schema.Column{KhumuUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "state", Type: field.TypeString, Default: "exists"},
		{Name: "content", Type: field.TypeString},
		{Name: "kind", Type: field.TypeString, Default: "anonymous"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "article_comments", Type: field.TypeInt, Nullable: true},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
		{Name: "study_article_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_articles_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_comments_children",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_khumu_users_comments",
				Columns:    []*schema.Column{CommentsColumns[7]},
				RefColumns: []*schema.Column{KhumuUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_study_articles_comments",
				Columns:    []*schema.Column{CommentsColumns[8]},
				RefColumns: []*schema.Column{StudyArticlesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// KhumuUsersColumns holds the columns for the "khumu_users" table.
	KhumuUsersColumns = []*schema.Column{
		{Name: "username", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "status", Type: field.TypeString, Default: "exists"},
		{Name: "created_at", Type: field.TypeTime},
	}
	// KhumuUsersTable holds the schema information for the "khumu_users" table.
	KhumuUsersTable = &schema.Table{
		Name:        "khumu_users",
		Columns:     KhumuUsersColumns,
		PrimaryKey:  []*schema.Column{KhumuUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// LikeCommentsColumns holds the columns for the "like_comments" table.
	LikeCommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeString, Nullable: true},
	}
	// LikeCommentsTable holds the schema information for the "like_comments" table.
	LikeCommentsTable = &schema.Table{
		Name:       "like_comments",
		Columns:    LikeCommentsColumns,
		PrimaryKey: []*schema.Column{LikeCommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "like_comments_comments_like",
				Columns:    []*schema.Column{LikeCommentsColumns[1]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "like_comments_khumu_users_like",
				Columns:    []*schema.Column{LikeCommentsColumns[2]},
				RefColumns: []*schema.Column{KhumuUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudyArticlesColumns holds the columns for the "study_articles" table.
	StudyArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author_id", Type: field.TypeString, Nullable: true},
	}
	// StudyArticlesTable holds the schema information for the "study_articles" table.
	StudyArticlesTable = &schema.Table{
		Name:       "study_articles",
		Columns:    StudyArticlesColumns,
		PrimaryKey: []*schema.Column{StudyArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "study_articles_khumu_users_studyArticles",
				Columns:    []*schema.Column{StudyArticlesColumns[1]},
				RefColumns: []*schema.Column{KhumuUsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		CommentsTable,
		KhumuUsersTable,
		LikeCommentsTable,
		StudyArticlesTable,
	}
)

func init() {
	ArticlesTable.ForeignKeys[0].RefTable = KhumuUsersTable
	ArticlesTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_0900_ai_ci",
	}
	CommentsTable.ForeignKeys[0].RefTable = ArticlesTable
	CommentsTable.ForeignKeys[1].RefTable = CommentsTable
	CommentsTable.ForeignKeys[2].RefTable = KhumuUsersTable
	CommentsTable.ForeignKeys[3].RefTable = StudyArticlesTable
	CommentsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_0900_ai_ci",
	}
	KhumuUsersTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_0900_ai_ci",
	}
	LikeCommentsTable.ForeignKeys[0].RefTable = CommentsTable
	LikeCommentsTable.ForeignKeys[1].RefTable = KhumuUsersTable
	LikeCommentsTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_0900_ai_ci",
	}
	StudyArticlesTable.ForeignKeys[0].RefTable = KhumuUsersTable
	StudyArticlesTable.Annotation = &entsql.Annotation{
		Charset:   "utf8mb4",
		Collation: "utf8mb4_0900_ai_ci",
	}
}
