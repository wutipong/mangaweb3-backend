// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
)

// HistoryUpdate is the builder for updating History entities.
type HistoryUpdate struct {
	config
	hooks    []Hook
	mutation *HistoryMutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (hu *HistoryUpdate) Where(ps ...predicate.History) *HistoryUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetCreateTime sets the "create_time" field.
func (hu *HistoryUpdate) SetCreateTime(t time.Time) *HistoryUpdate {
	hu.mutation.SetCreateTime(t)
	return hu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableCreateTime(t *time.Time) *HistoryUpdate {
	if t != nil {
		hu.SetCreateTime(*t)
	}
	return hu
}

// SetItemID sets the "item" edge to the Meta entity by ID.
func (hu *HistoryUpdate) SetItemID(id int) *HistoryUpdate {
	hu.mutation.SetItemID(id)
	return hu
}

// SetNillableItemID sets the "item" edge to the Meta entity by ID if the given value is not nil.
func (hu *HistoryUpdate) SetNillableItemID(id *int) *HistoryUpdate {
	if id != nil {
		hu = hu.SetItemID(*id)
	}
	return hu
}

// SetItem sets the "item" edge to the Meta entity.
func (hu *HistoryUpdate) SetItem(m *Meta) *HistoryUpdate {
	return hu.SetItemID(m.ID)
}

// Mutation returns the HistoryMutation object of the builder.
func (hu *HistoryUpdate) Mutation() *HistoryMutation {
	return hu.mutation
}

// ClearItem clears the "item" edge to the Meta entity.
func (hu *HistoryUpdate) ClearItem() *HistoryUpdate {
	hu.mutation.ClearItem()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistoryUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistoryUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hu *HistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.CreateTime(); ok {
		_spec.SetField(history.FieldCreateTime, field.TypeTime, value)
	}
	if hu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.ItemTable,
			Columns: []string{history.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.ItemTable,
			Columns: []string{history.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HistoryUpdateOne is the builder for updating a single History entity.
type HistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HistoryMutation
}

// SetCreateTime sets the "create_time" field.
func (huo *HistoryUpdateOne) SetCreateTime(t time.Time) *HistoryUpdateOne {
	huo.mutation.SetCreateTime(t)
	return huo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableCreateTime(t *time.Time) *HistoryUpdateOne {
	if t != nil {
		huo.SetCreateTime(*t)
	}
	return huo
}

// SetItemID sets the "item" edge to the Meta entity by ID.
func (huo *HistoryUpdateOne) SetItemID(id int) *HistoryUpdateOne {
	huo.mutation.SetItemID(id)
	return huo
}

// SetNillableItemID sets the "item" edge to the Meta entity by ID if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableItemID(id *int) *HistoryUpdateOne {
	if id != nil {
		huo = huo.SetItemID(*id)
	}
	return huo
}

// SetItem sets the "item" edge to the Meta entity.
func (huo *HistoryUpdateOne) SetItem(m *Meta) *HistoryUpdateOne {
	return huo.SetItemID(m.ID)
}

// Mutation returns the HistoryMutation object of the builder.
func (huo *HistoryUpdateOne) Mutation() *HistoryMutation {
	return huo.mutation
}

// ClearItem clears the "item" edge to the Meta entity.
func (huo *HistoryUpdateOne) ClearItem() *HistoryUpdateOne {
	huo.mutation.ClearItem()
	return huo
}

// Where appends a list predicates to the HistoryUpdate builder.
func (huo *HistoryUpdateOne) Where(ps ...predicate.History) *HistoryUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HistoryUpdateOne) Select(field string, fields ...string) *HistoryUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated History entity.
func (huo *HistoryUpdateOne) Save(ctx context.Context) (*History, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistoryUpdateOne) SaveX(ctx context.Context) *History {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistoryUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (huo *HistoryUpdateOne) sqlSave(ctx context.Context) (_node *History, err error) {
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "History.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for _, f := range fields {
			if !history.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.CreateTime(); ok {
		_spec.SetField(history.FieldCreateTime, field.TypeTime, value)
	}
	if huo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.ItemTable,
			Columns: []string{history.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   history.ItemTable,
			Columns: []string{history.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &History{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}