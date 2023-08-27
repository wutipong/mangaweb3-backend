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
	"github.com/wutipong/mangaweb3-backend/ent/meta"
)

// MetaCreate is the builder for creating a Meta entity.
type MetaCreate struct {
	config
	mutation *MetaMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (mc *MetaCreate) SetName(s string) *MetaCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetCreateTime sets the "create_time" field.
func (mc *MetaCreate) SetCreateTime(t time.Time) *MetaCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetFavorite sets the "favorite" field.
func (mc *MetaCreate) SetFavorite(b bool) *MetaCreate {
	mc.mutation.SetFavorite(b)
	return mc
}

// SetNillableFavorite sets the "favorite" field if the given value is not nil.
func (mc *MetaCreate) SetNillableFavorite(b *bool) *MetaCreate {
	if b != nil {
		mc.SetFavorite(*b)
	}
	return mc
}

// SetFileIndices sets the "file_indices" field.
func (mc *MetaCreate) SetFileIndices(i []int) *MetaCreate {
	mc.mutation.SetFileIndices(i)
	return mc
}

// SetThumbnail sets the "thumbnail" field.
func (mc *MetaCreate) SetThumbnail(b []byte) *MetaCreate {
	mc.mutation.SetThumbnail(b)
	return mc
}

// SetRead sets the "read" field.
func (mc *MetaCreate) SetRead(b bool) *MetaCreate {
	mc.mutation.SetRead(b)
	return mc
}

// SetTags sets the "tags" field.
func (mc *MetaCreate) SetTags(s []string) *MetaCreate {
	mc.mutation.SetTags(s)
	return mc
}

// Mutation returns the MetaMutation object of the builder.
func (mc *MetaCreate) Mutation() *MetaMutation {
	return mc.mutation
}

// Save creates the Meta in the database.
func (mc *MetaCreate) Save(ctx context.Context) (*Meta, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetaCreate) SaveX(ctx context.Context) *Meta {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetaCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetaCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MetaCreate) defaults() {
	if _, ok := mc.mutation.Favorite(); !ok {
		v := meta.DefaultFavorite
		mc.mutation.SetFavorite(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetaCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Meta.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := meta.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Meta.name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Meta.create_time"`)}
	}
	if _, ok := mc.mutation.Favorite(); !ok {
		return &ValidationError{Name: "favorite", err: errors.New(`ent: missing required field "Meta.favorite"`)}
	}
	if _, ok := mc.mutation.FileIndices(); !ok {
		return &ValidationError{Name: "file_indices", err: errors.New(`ent: missing required field "Meta.file_indices"`)}
	}
	if _, ok := mc.mutation.Thumbnail(); !ok {
		return &ValidationError{Name: "thumbnail", err: errors.New(`ent: missing required field "Meta.thumbnail"`)}
	}
	if _, ok := mc.mutation.Read(); !ok {
		return &ValidationError{Name: "read", err: errors.New(`ent: missing required field "Meta.read"`)}
	}
	if _, ok := mc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "Meta.tags"`)}
	}
	return nil
}

func (mc *MetaCreate) sqlSave(ctx context.Context) (*Meta, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MetaCreate) createSpec() (*Meta, *sqlgraph.CreateSpec) {
	var (
		_node = &Meta{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(meta.Table, sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt))
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(meta.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(meta.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.Favorite(); ok {
		_spec.SetField(meta.FieldFavorite, field.TypeBool, value)
		_node.Favorite = value
	}
	if value, ok := mc.mutation.FileIndices(); ok {
		_spec.SetField(meta.FieldFileIndices, field.TypeJSON, value)
		_node.FileIndices = value
	}
	if value, ok := mc.mutation.Thumbnail(); ok {
		_spec.SetField(meta.FieldThumbnail, field.TypeBytes, value)
		_node.Thumbnail = value
	}
	if value, ok := mc.mutation.Read(); ok {
		_spec.SetField(meta.FieldRead, field.TypeBool, value)
		_node.Read = value
	}
	if value, ok := mc.mutation.Tags(); ok {
		_spec.SetField(meta.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Meta.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetaUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mc *MetaCreate) OnConflict(opts ...sql.ConflictOption) *MetaUpsertOne {
	mc.conflict = opts
	return &MetaUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Meta.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MetaCreate) OnConflictColumns(columns ...string) *MetaUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MetaUpsertOne{
		create: mc,
	}
}

type (
	// MetaUpsertOne is the builder for "upsert"-ing
	//  one Meta node.
	MetaUpsertOne struct {
		create *MetaCreate
	}

	// MetaUpsert is the "OnConflict" setter.
	MetaUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *MetaUpsert) SetName(v string) *MetaUpsert {
	u.Set(meta.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetaUpsert) UpdateName() *MetaUpsert {
	u.SetExcluded(meta.FieldName)
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *MetaUpsert) SetCreateTime(v time.Time) *MetaUpsert {
	u.Set(meta.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *MetaUpsert) UpdateCreateTime() *MetaUpsert {
	u.SetExcluded(meta.FieldCreateTime)
	return u
}

// SetFavorite sets the "favorite" field.
func (u *MetaUpsert) SetFavorite(v bool) *MetaUpsert {
	u.Set(meta.FieldFavorite, v)
	return u
}

// UpdateFavorite sets the "favorite" field to the value that was provided on create.
func (u *MetaUpsert) UpdateFavorite() *MetaUpsert {
	u.SetExcluded(meta.FieldFavorite)
	return u
}

// SetFileIndices sets the "file_indices" field.
func (u *MetaUpsert) SetFileIndices(v []int) *MetaUpsert {
	u.Set(meta.FieldFileIndices, v)
	return u
}

// UpdateFileIndices sets the "file_indices" field to the value that was provided on create.
func (u *MetaUpsert) UpdateFileIndices() *MetaUpsert {
	u.SetExcluded(meta.FieldFileIndices)
	return u
}

// SetThumbnail sets the "thumbnail" field.
func (u *MetaUpsert) SetThumbnail(v []byte) *MetaUpsert {
	u.Set(meta.FieldThumbnail, v)
	return u
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *MetaUpsert) UpdateThumbnail() *MetaUpsert {
	u.SetExcluded(meta.FieldThumbnail)
	return u
}

// SetRead sets the "read" field.
func (u *MetaUpsert) SetRead(v bool) *MetaUpsert {
	u.Set(meta.FieldRead, v)
	return u
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *MetaUpsert) UpdateRead() *MetaUpsert {
	u.SetExcluded(meta.FieldRead)
	return u
}

// SetTags sets the "tags" field.
func (u *MetaUpsert) SetTags(v []string) *MetaUpsert {
	u.Set(meta.FieldTags, v)
	return u
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *MetaUpsert) UpdateTags() *MetaUpsert {
	u.SetExcluded(meta.FieldTags)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Meta.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MetaUpsertOne) UpdateNewValues() *MetaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Meta.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MetaUpsertOne) Ignore() *MetaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetaUpsertOne) DoNothing() *MetaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetaCreate.OnConflict
// documentation for more info.
func (u *MetaUpsertOne) Update(set func(*MetaUpsert)) *MetaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetaUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MetaUpsertOne) SetName(v string) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateName() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateName()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *MetaUpsertOne) SetCreateTime(v time.Time) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateCreateTime() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateCreateTime()
	})
}

// SetFavorite sets the "favorite" field.
func (u *MetaUpsertOne) SetFavorite(v bool) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetFavorite(v)
	})
}

// UpdateFavorite sets the "favorite" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateFavorite() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateFavorite()
	})
}

// SetFileIndices sets the "file_indices" field.
func (u *MetaUpsertOne) SetFileIndices(v []int) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetFileIndices(v)
	})
}

// UpdateFileIndices sets the "file_indices" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateFileIndices() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateFileIndices()
	})
}

// SetThumbnail sets the "thumbnail" field.
func (u *MetaUpsertOne) SetThumbnail(v []byte) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetThumbnail(v)
	})
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateThumbnail() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateThumbnail()
	})
}

// SetRead sets the "read" field.
func (u *MetaUpsertOne) SetRead(v bool) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateRead() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateRead()
	})
}

// SetTags sets the "tags" field.
func (u *MetaUpsertOne) SetTags(v []string) *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *MetaUpsertOne) UpdateTags() *MetaUpsertOne {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateTags()
	})
}

// Exec executes the query.
func (u *MetaUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetaCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetaUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MetaUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MetaUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MetaCreateBulk is the builder for creating many Meta entities in bulk.
type MetaCreateBulk struct {
	config
	builders []*MetaCreate
	conflict []sql.ConflictOption
}

// Save creates the Meta entities in the database.
func (mcb *MetaCreateBulk) Save(ctx context.Context) ([]*Meta, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Meta, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetaCreateBulk) SaveX(ctx context.Context) []*Meta {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetaCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetaCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Meta.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetaUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (mcb *MetaCreateBulk) OnConflict(opts ...sql.ConflictOption) *MetaUpsertBulk {
	mcb.conflict = opts
	return &MetaUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Meta.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MetaCreateBulk) OnConflictColumns(columns ...string) *MetaUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MetaUpsertBulk{
		create: mcb,
	}
}

// MetaUpsertBulk is the builder for "upsert"-ing
// a bulk of Meta nodes.
type MetaUpsertBulk struct {
	create *MetaCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Meta.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MetaUpsertBulk) UpdateNewValues() *MetaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Meta.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MetaUpsertBulk) Ignore() *MetaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetaUpsertBulk) DoNothing() *MetaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetaCreateBulk.OnConflict
// documentation for more info.
func (u *MetaUpsertBulk) Update(set func(*MetaUpsert)) *MetaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetaUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *MetaUpsertBulk) SetName(v string) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateName() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateName()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *MetaUpsertBulk) SetCreateTime(v time.Time) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateCreateTime() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateCreateTime()
	})
}

// SetFavorite sets the "favorite" field.
func (u *MetaUpsertBulk) SetFavorite(v bool) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetFavorite(v)
	})
}

// UpdateFavorite sets the "favorite" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateFavorite() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateFavorite()
	})
}

// SetFileIndices sets the "file_indices" field.
func (u *MetaUpsertBulk) SetFileIndices(v []int) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetFileIndices(v)
	})
}

// UpdateFileIndices sets the "file_indices" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateFileIndices() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateFileIndices()
	})
}

// SetThumbnail sets the "thumbnail" field.
func (u *MetaUpsertBulk) SetThumbnail(v []byte) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetThumbnail(v)
	})
}

// UpdateThumbnail sets the "thumbnail" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateThumbnail() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateThumbnail()
	})
}

// SetRead sets the "read" field.
func (u *MetaUpsertBulk) SetRead(v bool) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetRead(v)
	})
}

// UpdateRead sets the "read" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateRead() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateRead()
	})
}

// SetTags sets the "tags" field.
func (u *MetaUpsertBulk) SetTags(v []string) *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.SetTags(v)
	})
}

// UpdateTags sets the "tags" field to the value that was provided on create.
func (u *MetaUpsertBulk) UpdateTags() *MetaUpsertBulk {
	return u.Update(func(s *MetaUpsert) {
		s.UpdateTags()
	})
}

// Exec executes the query.
func (u *MetaUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MetaCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetaCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetaUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}