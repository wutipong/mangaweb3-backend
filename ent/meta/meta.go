// Code generated by ent, DO NOT EDIT.

package meta

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the meta type in the database.
	Label = "meta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldFavorite holds the string denoting the favorite field in the database.
	FieldFavorite = "favorite"
	// FieldFileIndices holds the string denoting the file_indices field in the database.
	FieldFileIndices = "file_indices"
	// FieldThumbnail holds the string denoting the thumbnail field in the database.
	FieldThumbnail = "thumbnail"
	// FieldRead holds the string denoting the read field in the database.
	FieldRead = "read"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldObjectType holds the string denoting the object_type field in the database.
	FieldObjectType = "object_type"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeHistories holds the string denoting the histories edge name in mutations.
	EdgeHistories = "histories"
	// Table holds the table name of the meta in the database.
	Table = "meta"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "meta_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// HistoriesTable is the table that holds the histories relation/edge.
	HistoriesTable = "histories"
	// HistoriesInverseTable is the table name for the History entity.
	// It exists in this package in order to avoid circular dependency with the "history" package.
	HistoriesInverseTable = "histories"
	// HistoriesColumn is the table column denoting the histories relation/edge.
	HistoriesColumn = "meta_histories"
)

// Columns holds all SQL columns for meta fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreateTime,
	FieldFavorite,
	FieldFileIndices,
	FieldThumbnail,
	FieldRead,
	FieldActive,
	FieldObjectType,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"meta_id", "tag_id"}
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultFavorite holds the default value on creation for the "favorite" field.
	DefaultFavorite bool
	// DefaultFileIndices holds the default value on creation for the "file_indices" field.
	DefaultFileIndices []int
	// DefaultRead holds the default value on creation for the "read" field.
	DefaultRead bool
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
)

// ObjectType defines the type for the "object_type" enum field.
type ObjectType string

// ObjectTypeZip is the default value of the ObjectType enum.
const DefaultObjectType = ObjectTypeZip

// ObjectType values.
const (
	ObjectTypeZip       ObjectType = "zip"
	ObjectTypeDirectory ObjectType = "directory"
)

func (ot ObjectType) String() string {
	return string(ot)
}

// ObjectTypeValidator is a validator for the "object_type" field enum values. It is called by the builders before save.
func ObjectTypeValidator(ot ObjectType) error {
	switch ot {
	case ObjectTypeZip, ObjectTypeDirectory:
		return nil
	default:
		return fmt.Errorf("meta: invalid enum value for object_type field: %q", ot)
	}
}

// OrderOption defines the ordering options for the Meta queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByFavorite orders the results by the favorite field.
func ByFavorite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFavorite, opts...).ToFunc()
}

// ByRead orders the results by the read field.
func ByRead(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRead, opts...).ToFunc()
}

// ByActive orders the results by the active field.
func ByActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActive, opts...).ToFunc()
}

// ByObjectType orders the results by the object_type field.
func ByObjectType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldObjectType, opts...).ToFunc()
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHistoriesCount orders the results by histories count.
func ByHistoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHistoriesStep(), opts...)
	}
}

// ByHistories orders the results by histories terms.
func ByHistories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHistoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newHistoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HistoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HistoriesTable, HistoriesColumn),
	)
}
