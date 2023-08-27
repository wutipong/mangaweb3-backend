// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/schema"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	metaFields := schema.Meta{}.Fields()
	_ = metaFields
	// metaDescName is the schema descriptor for name field.
	metaDescName := metaFields[0].Descriptor()
	// meta.NameValidator is a validator for the "name" field. It is called by the builders before save.
	meta.NameValidator = metaDescName.Validators[0].(func(string) error)
	// metaDescFavorite is the schema descriptor for favorite field.
	metaDescFavorite := metaFields[2].Descriptor()
	// meta.DefaultFavorite holds the default value on creation for the favorite field.
	meta.DefaultFavorite = metaDescFavorite.Default.(bool)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[0].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescFavorite is the schema descriptor for favorite field.
	tagDescFavorite := tagFields[1].Descriptor()
	// tag.DefaultFavorite holds the default value on creation for the favorite field.
	tag.DefaultFavorite = tagDescFavorite.Default.(bool)
	// tagDescHidden is the schema descriptor for hidden field.
	tagDescHidden := tagFields[2].Descriptor()
	// tag.DefaultHidden holds the default value on creation for the hidden field.
	tag.DefaultHidden = tagDescHidden.Default.(bool)
}