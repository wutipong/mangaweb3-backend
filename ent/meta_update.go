// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
)

// MetaUpdate is the builder for updating Meta entities.
type MetaUpdate struct {
	config
	hooks    []Hook
	mutation *MetaMutation
}

// Where appends a list predicates to the MetaUpdate builder.
func (mu *MetaUpdate) Where(ps ...predicate.Meta) *MetaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MetaUpdate) SetName(s string) *MetaUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableName(s *string) *MetaUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetCreateTime sets the "create_time" field.
func (mu *MetaUpdate) SetCreateTime(t time.Time) *MetaUpdate {
	mu.mutation.SetCreateTime(t)
	return mu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableCreateTime(t *time.Time) *MetaUpdate {
	if t != nil {
		mu.SetCreateTime(*t)
	}
	return mu
}

// SetFavorite sets the "favorite" field.
func (mu *MetaUpdate) SetFavorite(b bool) *MetaUpdate {
	mu.mutation.SetFavorite(b)
	return mu
}

// SetNillableFavorite sets the "favorite" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableFavorite(b *bool) *MetaUpdate {
	if b != nil {
		mu.SetFavorite(*b)
	}
	return mu
}

// SetFileIndices sets the "file_indices" field.
func (mu *MetaUpdate) SetFileIndices(i []int) *MetaUpdate {
	mu.mutation.SetFileIndices(i)
	return mu
}

// AppendFileIndices appends i to the "file_indices" field.
func (mu *MetaUpdate) AppendFileIndices(i []int) *MetaUpdate {
	mu.mutation.AppendFileIndices(i)
	return mu
}

// SetThumbnail sets the "thumbnail" field.
func (mu *MetaUpdate) SetThumbnail(b []byte) *MetaUpdate {
	mu.mutation.SetThumbnail(b)
	return mu
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (mu *MetaUpdate) ClearThumbnail() *MetaUpdate {
	mu.mutation.ClearThumbnail()
	return mu
}

// SetRead sets the "read" field.
func (mu *MetaUpdate) SetRead(b bool) *MetaUpdate {
	mu.mutation.SetRead(b)
	return mu
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableRead(b *bool) *MetaUpdate {
	if b != nil {
		mu.SetRead(*b)
	}
	return mu
}

// SetActive sets the "active" field.
func (mu *MetaUpdate) SetActive(b bool) *MetaUpdate {
	mu.mutation.SetActive(b)
	return mu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableActive(b *bool) *MetaUpdate {
	if b != nil {
		mu.SetActive(*b)
	}
	return mu
}

// SetObjectType sets the "object_type" field.
func (mu *MetaUpdate) SetObjectType(mt meta.ObjectType) *MetaUpdate {
	mu.mutation.SetObjectType(mt)
	return mu
}

// SetNillableObjectType sets the "object_type" field if the given value is not nil.
func (mu *MetaUpdate) SetNillableObjectType(mt *meta.ObjectType) *MetaUpdate {
	if mt != nil {
		mu.SetObjectType(*mt)
	}
	return mu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (mu *MetaUpdate) AddTagIDs(ids ...int) *MetaUpdate {
	mu.mutation.AddTagIDs(ids...)
	return mu
}

// AddTags adds the "tags" edges to the Tag entity.
func (mu *MetaUpdate) AddTags(t ...*Tag) *MetaUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTagIDs(ids...)
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (mu *MetaUpdate) AddHistoryIDs(ids ...int) *MetaUpdate {
	mu.mutation.AddHistoryIDs(ids...)
	return mu
}

// AddHistories adds the "histories" edges to the History entity.
func (mu *MetaUpdate) AddHistories(h ...*History) *MetaUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return mu.AddHistoryIDs(ids...)
}

// Mutation returns the MetaMutation object of the builder.
func (mu *MetaUpdate) Mutation() *MetaMutation {
	return mu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (mu *MetaUpdate) ClearTags() *MetaUpdate {
	mu.mutation.ClearTags()
	return mu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (mu *MetaUpdate) RemoveTagIDs(ids ...int) *MetaUpdate {
	mu.mutation.RemoveTagIDs(ids...)
	return mu
}

// RemoveTags removes "tags" edges to Tag entities.
func (mu *MetaUpdate) RemoveTags(t ...*Tag) *MetaUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTagIDs(ids...)
}

// ClearHistories clears all "histories" edges to the History entity.
func (mu *MetaUpdate) ClearHistories() *MetaUpdate {
	mu.mutation.ClearHistories()
	return mu
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (mu *MetaUpdate) RemoveHistoryIDs(ids ...int) *MetaUpdate {
	mu.mutation.RemoveHistoryIDs(ids...)
	return mu
}

// RemoveHistories removes "histories" edges to History entities.
func (mu *MetaUpdate) RemoveHistories(h ...*History) *MetaUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return mu.RemoveHistoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MetaUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := meta.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Meta.name": %w`, err)}
		}
	}
	if v, ok := mu.mutation.ObjectType(); ok {
		if err := meta.ObjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "object_type", err: fmt.Errorf(`ent: validator failed for field "Meta.object_type": %w`, err)}
		}
	}
	return nil
}

func (mu *MetaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(meta.Table, meta.Columns, sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(meta.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreateTime(); ok {
		_spec.SetField(meta.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Favorite(); ok {
		_spec.SetField(meta.FieldFavorite, field.TypeBool, value)
	}
	if value, ok := mu.mutation.FileIndices(); ok {
		_spec.SetField(meta.FieldFileIndices, field.TypeJSON, value)
	}
	if value, ok := mu.mutation.AppendedFileIndices(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, meta.FieldFileIndices, value)
		})
	}
	if value, ok := mu.mutation.Thumbnail(); ok {
		_spec.SetField(meta.FieldThumbnail, field.TypeBytes, value)
	}
	if mu.mutation.ThumbnailCleared() {
		_spec.ClearField(meta.FieldThumbnail, field.TypeBytes)
	}
	if value, ok := mu.mutation.Read(); ok {
		_spec.SetField(meta.FieldRead, field.TypeBool, value)
	}
	if value, ok := mu.mutation.Active(); ok {
		_spec.SetField(meta.FieldActive, field.TypeBool, value)
	}
	if value, ok := mu.mutation.ObjectType(); ok {
		_spec.SetField(meta.FieldObjectType, field.TypeEnum, value)
	}
	if mu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !mu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !mu.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meta.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetaUpdateOne is the builder for updating a single Meta entity.
type MetaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetaMutation
}

// SetName sets the "name" field.
func (muo *MetaUpdateOne) SetName(s string) *MetaUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableName(s *string) *MetaUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetCreateTime sets the "create_time" field.
func (muo *MetaUpdateOne) SetCreateTime(t time.Time) *MetaUpdateOne {
	muo.mutation.SetCreateTime(t)
	return muo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableCreateTime(t *time.Time) *MetaUpdateOne {
	if t != nil {
		muo.SetCreateTime(*t)
	}
	return muo
}

// SetFavorite sets the "favorite" field.
func (muo *MetaUpdateOne) SetFavorite(b bool) *MetaUpdateOne {
	muo.mutation.SetFavorite(b)
	return muo
}

// SetNillableFavorite sets the "favorite" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableFavorite(b *bool) *MetaUpdateOne {
	if b != nil {
		muo.SetFavorite(*b)
	}
	return muo
}

// SetFileIndices sets the "file_indices" field.
func (muo *MetaUpdateOne) SetFileIndices(i []int) *MetaUpdateOne {
	muo.mutation.SetFileIndices(i)
	return muo
}

// AppendFileIndices appends i to the "file_indices" field.
func (muo *MetaUpdateOne) AppendFileIndices(i []int) *MetaUpdateOne {
	muo.mutation.AppendFileIndices(i)
	return muo
}

// SetThumbnail sets the "thumbnail" field.
func (muo *MetaUpdateOne) SetThumbnail(b []byte) *MetaUpdateOne {
	muo.mutation.SetThumbnail(b)
	return muo
}

// ClearThumbnail clears the value of the "thumbnail" field.
func (muo *MetaUpdateOne) ClearThumbnail() *MetaUpdateOne {
	muo.mutation.ClearThumbnail()
	return muo
}

// SetRead sets the "read" field.
func (muo *MetaUpdateOne) SetRead(b bool) *MetaUpdateOne {
	muo.mutation.SetRead(b)
	return muo
}

// SetNillableRead sets the "read" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableRead(b *bool) *MetaUpdateOne {
	if b != nil {
		muo.SetRead(*b)
	}
	return muo
}

// SetActive sets the "active" field.
func (muo *MetaUpdateOne) SetActive(b bool) *MetaUpdateOne {
	muo.mutation.SetActive(b)
	return muo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableActive(b *bool) *MetaUpdateOne {
	if b != nil {
		muo.SetActive(*b)
	}
	return muo
}

// SetObjectType sets the "object_type" field.
func (muo *MetaUpdateOne) SetObjectType(mt meta.ObjectType) *MetaUpdateOne {
	muo.mutation.SetObjectType(mt)
	return muo
}

// SetNillableObjectType sets the "object_type" field if the given value is not nil.
func (muo *MetaUpdateOne) SetNillableObjectType(mt *meta.ObjectType) *MetaUpdateOne {
	if mt != nil {
		muo.SetObjectType(*mt)
	}
	return muo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (muo *MetaUpdateOne) AddTagIDs(ids ...int) *MetaUpdateOne {
	muo.mutation.AddTagIDs(ids...)
	return muo
}

// AddTags adds the "tags" edges to the Tag entity.
func (muo *MetaUpdateOne) AddTags(t ...*Tag) *MetaUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTagIDs(ids...)
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (muo *MetaUpdateOne) AddHistoryIDs(ids ...int) *MetaUpdateOne {
	muo.mutation.AddHistoryIDs(ids...)
	return muo
}

// AddHistories adds the "histories" edges to the History entity.
func (muo *MetaUpdateOne) AddHistories(h ...*History) *MetaUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return muo.AddHistoryIDs(ids...)
}

// Mutation returns the MetaMutation object of the builder.
func (muo *MetaUpdateOne) Mutation() *MetaMutation {
	return muo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (muo *MetaUpdateOne) ClearTags() *MetaUpdateOne {
	muo.mutation.ClearTags()
	return muo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (muo *MetaUpdateOne) RemoveTagIDs(ids ...int) *MetaUpdateOne {
	muo.mutation.RemoveTagIDs(ids...)
	return muo
}

// RemoveTags removes "tags" edges to Tag entities.
func (muo *MetaUpdateOne) RemoveTags(t ...*Tag) *MetaUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTagIDs(ids...)
}

// ClearHistories clears all "histories" edges to the History entity.
func (muo *MetaUpdateOne) ClearHistories() *MetaUpdateOne {
	muo.mutation.ClearHistories()
	return muo
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (muo *MetaUpdateOne) RemoveHistoryIDs(ids ...int) *MetaUpdateOne {
	muo.mutation.RemoveHistoryIDs(ids...)
	return muo
}

// RemoveHistories removes "histories" edges to History entities.
func (muo *MetaUpdateOne) RemoveHistories(h ...*History) *MetaUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return muo.RemoveHistoryIDs(ids...)
}

// Where appends a list predicates to the MetaUpdate builder.
func (muo *MetaUpdateOne) Where(ps ...predicate.Meta) *MetaUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetaUpdateOne) Select(field string, fields ...string) *MetaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Meta entity.
func (muo *MetaUpdateOne) Save(ctx context.Context) (*Meta, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetaUpdateOne) SaveX(ctx context.Context) *Meta {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MetaUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := meta.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Meta.name": %w`, err)}
		}
	}
	if v, ok := muo.mutation.ObjectType(); ok {
		if err := meta.ObjectTypeValidator(v); err != nil {
			return &ValidationError{Name: "object_type", err: fmt.Errorf(`ent: validator failed for field "Meta.object_type": %w`, err)}
		}
	}
	return nil
}

func (muo *MetaUpdateOne) sqlSave(ctx context.Context) (_node *Meta, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(meta.Table, meta.Columns, sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Meta.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meta.FieldID)
		for _, f := range fields {
			if !meta.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(meta.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreateTime(); ok {
		_spec.SetField(meta.FieldCreateTime, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Favorite(); ok {
		_spec.SetField(meta.FieldFavorite, field.TypeBool, value)
	}
	if value, ok := muo.mutation.FileIndices(); ok {
		_spec.SetField(meta.FieldFileIndices, field.TypeJSON, value)
	}
	if value, ok := muo.mutation.AppendedFileIndices(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, meta.FieldFileIndices, value)
		})
	}
	if value, ok := muo.mutation.Thumbnail(); ok {
		_spec.SetField(meta.FieldThumbnail, field.TypeBytes, value)
	}
	if muo.mutation.ThumbnailCleared() {
		_spec.ClearField(meta.FieldThumbnail, field.TypeBytes)
	}
	if value, ok := muo.mutation.Read(); ok {
		_spec.SetField(meta.FieldRead, field.TypeBool, value)
	}
	if value, ok := muo.mutation.Active(); ok {
		_spec.SetField(meta.FieldActive, field.TypeBool, value)
	}
	if value, ok := muo.mutation.ObjectType(); ok {
		_spec.SetField(meta.FieldObjectType, field.TypeEnum, value)
	}
	if muo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !muo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   meta.TagsTable,
			Columns: meta.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !muo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meta.HistoriesTable,
			Columns: []string{meta.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Meta{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meta.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
