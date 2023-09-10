package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("name").NotEmpty().Unique().Annotations(
			entproto.Field(2),
		),
		field.Time("create_time").Annotations(
			entproto.Field(3),
		),
		field.Bool("favorite").Default(false).Annotations(
			entproto.Field(4),
		),
		field.Ints("file_indices").Annotations(
			entproto.Skip(),
		),
		field.Bytes("thumbnail").Optional().Sensitive().Annotations(
			entproto.Field(5),
		),
		field.Bool("read").Annotations(
			entproto.Field(6),
		),
		field.Bool("active").Default(true).Annotations(
			entproto.Field(7),
		),
	}
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type).Annotations(
			entproto.Skip(),
		),
	}
}

// Annotations of the Meta.
func (Meta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
