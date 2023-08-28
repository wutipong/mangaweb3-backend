// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MetaColumns holds the columns for the "meta" table.
	MetaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "favorite", Type: field.TypeBool, Default: false},
		{Name: "file_indices", Type: field.TypeJSON},
		{Name: "thumbnail", Type: field.TypeBytes, Nullable: true},
		{Name: "read", Type: field.TypeBool},
		{Name: "tags", Type: field.TypeJSON},
	}
	// MetaTable holds the schema information for the "meta" table.
	MetaTable = &schema.Table{
		Name:       "meta",
		Columns:    MetaColumns,
		PrimaryKey: []*schema.Column{MetaColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "favorite", Type: field.TypeBool, Default: false},
		{Name: "hidden", Type: field.TypeBool, Default: false},
		{Name: "thumbnail", Type: field.TypeBytes, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MetaTable,
		TagsTable,
	}
)

func init() {
}
