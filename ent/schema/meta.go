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
		field.Bool("read").Default(false),
		field.Bool("active").Default(true),
		field.Bool("hidden").Default(false),
		field.Enum("container_type").Values("zip", "directory").Default("zip"),
		field.Int("thumbnail_index").Default(0).Optional(),
		field.Int("thumbnail_x").Default(0).Optional(),
		field.Int("thumbnail_y").Default(0).Optional(),
		field.Int("thumbnail_width").Default(0).Optional(),
		field.Int("thumbnail_height").Default(0).Optional(),
	}
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("histories", History.Type),
		edge.From("user", User.Type).Ref("favorite_items"),
	}
}
