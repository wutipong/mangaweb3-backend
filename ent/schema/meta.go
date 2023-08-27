package schema

import (
	"entgo.io/ent"
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
		field.Time("create_time"),
		field.Bool("favorite").Default(false),
		field.Ints("file_indices"),
		field.Bytes("thumbnail"),
		field.Bool("read"),
		field.Strings("tags"),
	}
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return nil
}
