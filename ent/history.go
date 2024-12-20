// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/user"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges          HistoryEdges `json:"edges"`
	meta_histories *int
	user_histories *int
	selectValues   sql.SelectValues
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Item holds the value of the item edge.
	Item *Meta `json:"item,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) ItemOrErr() (*Meta, error) {
	if e.Item != nil {
		return e.Item, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: meta.Label}
	}
	return nil, &NotLoadedError{edge: "item"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistoryEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			values[i] = new(sql.NullInt64)
		case history.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case history.ForeignKeys[0]: // meta_histories
			values[i] = new(sql.NullInt64)
		case history.ForeignKeys[1]: // user_histories
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case history.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				h.CreateTime = value.Time
			}
		case history.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field meta_histories", value)
			} else if value.Valid {
				h.meta_histories = new(int)
				*h.meta_histories = int(value.Int64)
			}
		case history.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_histories", value)
			} else if value.Valid {
				h.user_histories = new(int)
				*h.user_histories = int(value.Int64)
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the History.
// This includes values selected through modifiers, order, etc.
func (h *History) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryItem queries the "item" edge of the History entity.
func (h *History) QueryItem() *MetaQuery {
	return NewHistoryClient(h.config).QueryItem(h)
}

// QueryUser queries the "user" edge of the History entity.
func (h *History) QueryUser() *UserQuery {
	return NewHistoryClient(h.config).QueryUser(h)
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return NewHistoryClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("create_time=")
	builder.WriteString(h.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History
