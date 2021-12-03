package schema

import (
	"encoding/json"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Annotations of the Comment.
func (Comment) Annotations() []schema.Annotation {
	return nil
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("state").Default("exists"),
		field.String("content"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", KhumuUser.Type).
			Ref("comments").
			Unique(),
		edge.From("article", Article.Type).
			Ref("comments").
			Unique(),
		edge.From("studyArticle", StudyArticle.Type).
			Ref("comments").
			Unique(),
		edge.From("parent", Comment.Type).
			Ref("children").
			Unique(),
		edge.To("children", Comment.Type).
			StorageKey(func(key *edge.StorageKey) {
				key.Columns = []string{"parent_id"}
			}),
		edge.To("like", LikeComment.Type).
			StorageKey(func(key *edge.StorageKey) {
				key.Columns = []string{"comment_id"}
			}),
	}
}

// MarshalBinary -
func (c *Comment) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary -
func (c *Comment) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}
