package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty().Annotations(
			entproto.Field(2),
		),
		field.Bool("favorite").Default(false).Annotations(
			entproto.Field(3),
		),
		field.Bool("hidden").Default(false).Annotations(
			entproto.Field(4),
		),
		field.Bytes("thumbnail").Optional().Sensitive().Annotations(
			entproto.Field(5),
		),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("meta", Meta.Type).Ref("tags").Annotations(
			entproto.Skip(),
		),
	}
}

// Annotations of the Tag.
func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
