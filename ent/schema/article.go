package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Annotations of the Article.
func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Charset:   "utf8mb4",
			Collation: "utf8mb4_0900_ai_ci",
		}}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("author", KhumuUser.Type).Unique().Ref("articles"),
	}
}
