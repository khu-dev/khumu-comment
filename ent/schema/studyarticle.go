package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StudyArticle holds the schema definition for the StudyArticle entity.
type StudyArticle struct {
	ent.Schema
}

// Annotations of the StudyArticle.
func (StudyArticle) Annotations() []schema.Annotation {
	return nil
}

// Fields of the StudyArticle.
func (StudyArticle) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
	}
}

// Edges of the StudyArticle.
func (StudyArticle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type).
			StorageKey(func(key *edge.StorageKey) {
				key.Columns = []string{"study_article_id"}
			}),
		// 참조하는 엔티티가 From
		// 자기가 한 개만 참조하면 Unique
		edge.From("author", KhumuUser.Type).
			Ref("studyArticles").
			Unique(),
	}
}
