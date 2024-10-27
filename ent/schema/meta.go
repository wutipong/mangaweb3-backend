package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Meta holds the schema definition for the Meta entity.
type Meta struct {
	ent.Schema
}

// Fields of the Meta.
func (Meta) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.Time("create_time").Default(time.Now),
		field.Bool("favorite").Default(false),
		field.Ints("file_indices").Default([]int{}),
		field.Bytes("thumbnail").Optional().Sensitive(),
		field.Bool("read").Default(false),
		field.Bool("active").Default(true),
		field.Enum("container_type").Values("zip", "directory").Default("zip"),
	}
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("histories", History.Type),
	}
}
