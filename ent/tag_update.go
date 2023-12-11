// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TagUpdate) SetNillableName(s *string) *TagUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetFavorite sets the "favorite" field.
func (tu *TagUpdate) SetFavorite(b bool) *TagUpdate {
	tu.mutation.SetFavorite(b)
	return tu
}

// SetNillableFavorite sets the "favorite" field if the given value is not nil.
func (tu *TagUpdate) SetNillableFavorite(b *bool) *TagUpdate {
	if b != nil {
		tu.SetFavorite(*b)
	}
	return tu
}

// SetHidden sets the "hidden" field.
func (tu *TagUpdate) SetHidden(b bool) *TagUpdate {
	tu.mutation.SetHidden(b)
	return tu
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (tu *TagUpdate) SetNillableHidden(b *bool) *TagUpdate {
	if b != nil {
		tu.SetHidden(*b)
	}
	return tu
}

// SetThumbnail sets the "thumbnail" field.
func (tu *TagUpdate) SetThumbnail(b []byte) *TagUpdate {
	tu.mutation.SetThumbnail(b)
	return tu
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (tu *TagUpdate) ClearThumbnail() *TagUpdate {
	tu.mutation.ClearThumbnail()
	return tu
}

// AddMetumIDs adds the "meta" edge to the Meta entity by IDs.
func (tu *TagUpdate) AddMetumIDs(ids ...int) *TagUpdate {
	tu.mutation.AddMetumIDs(ids...)
	return tu
}

// AddMeta adds the "meta" edges to the Meta entity.
func (tu *TagUpdate) AddMeta(m ...*Meta) *TagUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.AddMetumIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearMeta clears all "meta" edges to the Meta entity.
func (tu *TagUpdate) ClearMeta() *TagUpdate {
	tu.mutation.ClearMeta()
	return tu
}

// RemoveMetumIDs removes the "meta" edge to Meta entities by IDs.
func (tu *TagUpdate) RemoveMetumIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveMetumIDs(ids...)
	return tu
}

// RemoveMeta removes "meta" edges to Meta entities.
func (tu *TagUpdate) RemoveMeta(m ...*Meta) *TagUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.RemoveMetumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Favorite(); ok {
		_spec.SetField(tag.FieldFavorite, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Hidden(); ok {
		_spec.SetField(tag.FieldHidden, field.TypeBool, value)
	}
	if value, ok := tu.mutation.Thumbnail(); ok {
		_spec.SetField(tag.FieldThumbnail, field.TypeBytes, value)
	}
	if tu.mutation.ThumbnailCleared() {
		_spec.ClearField(tag.FieldThumbnail, field.TypeBytes)
	}
	if tu.mutation.MetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedMetaIDs(); len(nodes) > 0 && !tu.mutation.MetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.MetaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableName(s *string) *TagUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetFavorite sets the "favorite" field.
func (tuo *TagUpdateOne) SetFavorite(b bool) *TagUpdateOne {
	tuo.mutation.SetFavorite(b)
	return tuo
}

// SetNillableFavorite sets the "favorite" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableFavorite(b *bool) *TagUpdateOne {
	if b != nil {
		tuo.SetFavorite(*b)
	}
	return tuo
}

// SetHidden sets the "hidden" field.
func (tuo *TagUpdateOne) SetHidden(b bool) *TagUpdateOne {
	tuo.mutation.SetHidden(b)
	return tuo
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableHidden(b *bool) *TagUpdateOne {
	if b != nil {
		tuo.SetHidden(*b)
	}
	return tuo
}

// SetThumbnail sets the "thumbnail" field.
func (tuo *TagUpdateOne) SetThumbnail(b []byte) *TagUpdateOne {
	tuo.mutation.SetThumbnail(b)
	return tuo
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (tuo *TagUpdateOne) ClearThumbnail() *TagUpdateOne {
	tuo.mutation.ClearThumbnail()
	return tuo
}

// AddMetumIDs adds the "meta" edge to the Meta entity by IDs.
func (tuo *TagUpdateOne) AddMetumIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddMetumIDs(ids...)
	return tuo
}

// AddMeta adds the "meta" edges to the Meta entity.
func (tuo *TagUpdateOne) AddMeta(m ...*Meta) *TagUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.AddMetumIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearMeta clears all "meta" edges to the Meta entity.
func (tuo *TagUpdateOne) ClearMeta() *TagUpdateOne {
	tuo.mutation.ClearMeta()
	return tuo
}

// RemoveMetumIDs removes the "meta" edge to Meta entities by IDs.
func (tuo *TagUpdateOne) RemoveMetumIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveMetumIDs(ids...)
	return tuo
}

// RemoveMeta removes "meta" edges to Meta entities.
func (tuo *TagUpdateOne) RemoveMeta(m ...*Meta) *TagUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.RemoveMetumIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (tuo *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Favorite(); ok {
		_spec.SetField(tag.FieldFavorite, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Hidden(); ok {
		_spec.SetField(tag.FieldHidden, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.Thumbnail(); ok {
		_spec.SetField(tag.FieldThumbnail, field.TypeBytes, value)
	}
	if tuo.mutation.ThumbnailCleared() {
		_spec.ClearField(tag.FieldThumbnail, field.TypeBytes)
	}
	if tuo.mutation.MetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedMetaIDs(); len(nodes) > 0 && !tuo.mutation.MetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.MetaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.MetaTable,
			Columns: tag.MetaPrimaryKey,
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
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
