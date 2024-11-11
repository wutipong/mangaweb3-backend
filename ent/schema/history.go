package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// History holds the schema definition for the History entity.
type History struct {
	ent.Schema
}

// Fields of the History.
func (History) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Default(time.Now),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Meta.Type).Ref("histories").Unique(),
		edge.From("user", User.Type).Ref("histories").Unique(),
	}
}
