// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

// Favorite applies equality check predicate on the "favorite" field. It's identical to FavoriteEQ.
func Favorite(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldFavorite, v))
}

// Hidden applies equality check predicate on the "hidden" field. It's identical to HiddenEQ.
func Hidden(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldHidden, v))
}

// Thumbnail applies equality check predicate on the "thumbnail" field. It's identical to ThumbnailEQ.
func Thumbnail(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldThumbnail, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContainsFold(FieldName, v))
}

// FavoriteEQ applies the EQ predicate on the "favorite" field.
func FavoriteEQ(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldFavorite, v))
}

// FavoriteNEQ applies the NEQ predicate on the "favorite" field.
func FavoriteNEQ(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldFavorite, v))
}

// HiddenEQ applies the EQ predicate on the "hidden" field.
func HiddenEQ(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldHidden, v))
}

// HiddenNEQ applies the NEQ predicate on the "hidden" field.
func HiddenNEQ(v bool) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldHidden, v))
}

// ThumbnailEQ applies the EQ predicate on the "thumbnail" field.
func ThumbnailEQ(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldThumbnail, v))
}

// ThumbnailNEQ applies the NEQ predicate on the "thumbnail" field.
func ThumbnailNEQ(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldThumbnail, v))
}

// ThumbnailIn applies the In predicate on the "thumbnail" field.
func ThumbnailIn(vs ...[]byte) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldThumbnail, vs...))
}

// ThumbnailNotIn applies the NotIn predicate on the "thumbnail" field.
func ThumbnailNotIn(vs ...[]byte) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldThumbnail, vs...))
}

// ThumbnailGT applies the GT predicate on the "thumbnail" field.
func ThumbnailGT(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldThumbnail, v))
}

// ThumbnailGTE applies the GTE predicate on the "thumbnail" field.
func ThumbnailGTE(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldThumbnail, v))
}

// ThumbnailLT applies the LT predicate on the "thumbnail" field.
func ThumbnailLT(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldThumbnail, v))
}

// ThumbnailLTE applies the LTE predicate on the "thumbnail" field.
func ThumbnailLTE(v []byte) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldThumbnail, v))
}

// ThumbnailIsNil applies the IsNil predicate on the "thumbnail" field.
func ThumbnailIsNil() predicate.Tag {
	return predicate.Tag(sql.FieldIsNull(FieldThumbnail))
}

// ThumbnailNotNil applies the NotNil predicate on the "thumbnail" field.
func ThumbnailNotNil() predicate.Tag {
	return predicate.Tag(sql.FieldNotNull(FieldThumbnail))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
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
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		p(s.Not())
	})
}
