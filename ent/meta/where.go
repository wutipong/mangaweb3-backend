// Code generated by ent, DO NOT EDIT.

package meta

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Meta {
	return predicate.Meta(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Meta {
	return predicate.Meta(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Meta {
	return predicate.Meta(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Meta {
	return predicate.Meta(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Meta {
	return predicate.Meta(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Meta {
	return predicate.Meta(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldName, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldCreateTime, v))
}

// Favorite applies equality check predicate on the "favorite" field. It's identical to FavoriteEQ.
func Favorite(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldFavorite, v))
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldThumbnail, v))
}

// Read applies equality check predicate on the "read" field. It's identical to ReadEQ.
func Read(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldRead, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Meta {
	return predicate.Meta(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Meta {
	return predicate.Meta(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Meta {
	return predicate.Meta(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Meta {
	return predicate.Meta(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Meta {
	return predicate.Meta(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Meta {
	return predicate.Meta(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Meta {
	return predicate.Meta(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Meta {
	return predicate.Meta(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Meta {
	return predicate.Meta(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Meta {
	return predicate.Meta(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Meta {
	return predicate.Meta(sql.FieldContainsFold(FieldName, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Meta {
	return predicate.Meta(sql.FieldLTE(FieldCreateTime, v))
}

// FavoriteEQ applies the EQ predicate on the "favorite" field.
func FavoriteEQ(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldFavorite, v))
}

// FavoriteNEQ applies the NEQ predicate on the "favorite" field.
func FavoriteNEQ(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldFavorite, v))
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldThumbnail, v))
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldThumbnail, v))
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...[]byte) predicate.Meta {
	return predicate.Meta(sql.FieldIn(FieldThumbnail, vs...))
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...[]byte) predicate.Meta {
	return predicate.Meta(sql.FieldNotIn(FieldThumbnail, vs...))
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldGT(FieldThumbnail, v))
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldGTE(FieldThumbnail, v))
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldLT(FieldThumbnail, v))
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v []byte) predicate.Meta {
	return predicate.Meta(sql.FieldLTE(FieldThumbnail, v))
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.Meta {
	return predicate.Meta(sql.FieldIsNull(FieldThumbnail))
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.Meta {
	return predicate.Meta(sql.FieldNotNull(FieldThumbnail))
}

// ReadEQ applies the EQ predicate on the "read" field.
func ReadEQ(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldEQ(FieldRead, v))
}

// ReadNEQ applies the NEQ predicate on the "read" field.
func ReadNEQ(v bool) predicate.Meta {
	return predicate.Meta(sql.FieldNEQ(FieldRead, v))
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagsTable, TagsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		step := newTagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Meta) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Meta) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Meta) predicate.Meta {
	return predicate.Meta(func(s *sql.Selector) {
		p(s.Not())
	})
}
