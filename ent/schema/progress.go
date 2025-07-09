package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Progress holds the schema definition for the user progress entity.
type Progress struct {
	ent.Schema
}

// Fields of the Tag.
func (Progress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("page").Default(-1),
	}
}

// Edges of the Progress.
func (Progress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Meta.Type).Ref("progress").Unique(),
		edge.From("user", User.Type).Ref("progress").Unique(),
	}
}
