package schema

import (
	"entgo.io/ent"
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
	return nil
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.From("author", KhumuUser.Type).Unique().Ref("articles"),
	}
}
