package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type KhumuUser struct {
	ent.Schema
}

// Annotations of the User.
func (KhumuUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_0900_ai_ci",
		}}
}

// Fields of the KhumuUser.
func (KhumuUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("username"),
		field.String("nickname"),
		field.String("status").Default("exists"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the KhumuUser.
func (KhumuUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type).StorageKey(func(key *edge.StorageKey) {
			key.Columns = []string{"author_id"}
		}),
		edge.To("articles", Article.Type).StorageKey(func(key *edge.StorageKey) {
			key.Columns = []string{"author_id"}
		}),
		edge.To("studyArticles", StudyArticle.Type).StorageKey(func(key *edge.StorageKey) {
			key.Columns = []string{"author_id"}
		}),
		edge.To("like", LikeComment.Type).
			StorageKey(func(key *edge.StorageKey) {
				key.Columns = []string{"user_id"}
			}),
	}
}
