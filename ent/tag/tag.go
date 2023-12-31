// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFavorite holds the string denoting the favorite field in the database.
	FieldFavorite = "favorite"
	// FieldHidden holds the string denoting the hidden field in the database.
	FieldHidden = "hidden"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// EdgeMeta holds the string denoting the meta edge name in mutations.
	EdgeMeta = "meta"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// MetaTable is the table that holds the meta relation/edge. The primary key declared below.
	MetaTable = "meta_tags"
	// MetaInverseTable is the table name for the Meta entity.
	// It exists in this package in order to avoid circular dependency with the "meta" package.
	MetaInverseTable = "meta"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFavorite,
	FieldHidden,
	FieldThumbnail,
}

var (
	// MetaPrimaryKey and MetaColumn2 are the table columns denoting the
	// primary key for the meta relation (M2M).
	MetaPrimaryKey = []string{"meta_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultFavorite holds the default value on creation for the "favorite" field.
	DefaultFavorite bool
	// DefaultHidden holds the default value on creation for the "hidden" field.
	DefaultHidden bool
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFavorite orders the results by the favorite field.
func ByFavorite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavorite, opts...).ToFunc()
}

// ByHidden orders the results by the hidden field.
func ByHidden(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHidden, opts...).ToFunc()
}

// ByMetaCount orders the results by meta count.
func ByMetaCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMetaStep(), opts...)
	}
}

// ByMeta orders the results by meta terms.
func ByMeta(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetaStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMetaStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetaInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, MetaTable, MetaPrimaryKey...),
	)
}
