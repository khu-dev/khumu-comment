package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LikeComment holds the schema definition for the LikeComment entity.
type LikeComment struct {
	ent.Schema
}

// Annotations of the LikeComment.
func (LikeComment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "comment_likecomment"},
	}
}

// Fields of the LikeComment.
func (LikeComment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
	}
}

// Edges of the LikeComment.
func (LikeComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("likedBy", KhumuUser.Type).Ref("like").Unique(),
		edge.From("about", Comment.Type).Ref("like").Unique(),
	}
}
