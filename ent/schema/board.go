package schema

import "entgo.io/ent"

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return nil
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return nil
}
