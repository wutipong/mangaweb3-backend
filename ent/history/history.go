// Code generated by ent, DO NOT EDIT.

package history

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// Table holds the table name of the history in the database.
	Table = "histories"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "histories"
	// ItemInverseTable is the table name for the Meta entity.
	// It exists in this package in order to avoid circular dependency with the "meta" package.
	ItemInverseTable = "meta"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "meta_histories"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"meta_histories",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
)

// OrderOption defines the ordering options for the History queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByItemField orders the results by item field.
func ByItemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemStep(), sql.OrderByField(field, opts...))
	}
}
func newItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ItemTable, ItemColumn),
	)
}
