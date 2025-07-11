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
	"github.com/wutipong/mangaweb3-backend/ent/progress"
	"github.com/wutipong/mangaweb3-backend/ent/user"
)

// ProgressCreate is the builder for creating a Progress entity.
type ProgressCreate struct {
	config
	mutation *ProgressMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPage sets the "page" field.
func (pc *ProgressCreate) SetPage(i int) *ProgressCreate {
	pc.mutation.SetPage(i)
	return pc
}

// SetNillablePage sets the "page" field if the given value is not nil.
func (pc *ProgressCreate) SetNillablePage(i *int) *ProgressCreate {
	if i != nil {
		pc.SetPage(*i)
	}
	return pc
}

// SetItemID sets the "item_id" field.
func (pc *ProgressCreate) SetItemID(i int) *ProgressCreate {
	pc.mutation.SetItemID(i)
	return pc
}

// SetNillableItemID sets the "item_id" field if the given value is not nil.
func (pc *ProgressCreate) SetNillableItemID(i *int) *ProgressCreate {
	if i != nil {
		pc.SetItemID(*i)
	}
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *ProgressCreate) SetUserID(i int) *ProgressCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *ProgressCreate) SetNillableUserID(i *int) *ProgressCreate {
	if i != nil {
		pc.SetUserID(*i)
	}
	return pc
}

// SetItem sets the "item" edge to the Meta entity.
func (pc *ProgressCreate) SetItem(m *Meta) *ProgressCreate {
	return pc.SetItemID(m.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pc *ProgressCreate) SetUser(u *User) *ProgressCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the ProgressMutation object of the builder.
func (pc *ProgressCreate) Mutation() *ProgressMutation {
	return pc.mutation
}

// Save creates the Progress in the database.
func (pc *ProgressCreate) Save(ctx context.Context) (*Progress, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProgressCreate) SaveX(ctx context.Context) *Progress {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProgressCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProgressCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProgressCreate) defaults() {
	if _, ok := pc.mutation.Page(); !ok {
		v := progress.DefaultPage
		pc.mutation.SetPage(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProgressCreate) check() error {
	if _, ok := pc.mutation.Page(); !ok {
		return &ValidationError{Name: "page", err: errors.New(`ent: missing required field "Progress.page"`)}
	}
	return nil
}

func (pc *ProgressCreate) sqlSave(ctx context.Context) (*Progress, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProgressCreate) createSpec() (*Progress, *sqlgraph.CreateSpec) {
	var (
		_node = &Progress{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(progress.Table, sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Page(); ok {
		_spec.SetField(progress.FieldPage, field.TypeInt, value)
		_node.Page = value
	}
	if nodes := pc.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   progress.ItemTable,
			Columns: []string{progress.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ItemID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   progress.UserTable,
			Columns: []string{progress.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Progress.Create().
//		SetPage(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProgressUpsert) {
//			SetPage(v+v).
//		}).
//		Exec(ctx)
func (pc *ProgressCreate) OnConflict(opts ...sql.ConflictOption) *ProgressUpsertOne {
	pc.conflict = opts
	return &ProgressUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Progress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProgressCreate) OnConflictColumns(columns ...string) *ProgressUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProgressUpsertOne{
		create: pc,
	}
}

type (
	// ProgressUpsertOne is the builder for "upsert"-ing
	//  one Progress node.
	ProgressUpsertOne struct {
		create *ProgressCreate
	}

	// ProgressUpsert is the "OnConflict" setter.
	ProgressUpsert struct {
		*sql.UpdateSet
	}
)

// SetPage sets the "page" field.
func (u *ProgressUpsert) SetPage(v int) *ProgressUpsert {
	u.Set(progress.FieldPage, v)
	return u
}

// UpdatePage sets the "page" field to the value that was provided on create.
func (u *ProgressUpsert) UpdatePage() *ProgressUpsert {
	u.SetExcluded(progress.FieldPage)
	return u
}

// AddPage adds v to the "page" field.
func (u *ProgressUpsert) AddPage(v int) *ProgressUpsert {
	u.Add(progress.FieldPage, v)
	return u
}

// SetItemID sets the "item_id" field.
func (u *ProgressUpsert) SetItemID(v int) *ProgressUpsert {
	u.Set(progress.FieldItemID, v)
	return u
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ProgressUpsert) UpdateItemID() *ProgressUpsert {
	u.SetExcluded(progress.FieldItemID)
	return u
}

// ClearItemID clears the value of the "item_id" field.
func (u *ProgressUpsert) ClearItemID() *ProgressUpsert {
	u.SetNull(progress.FieldItemID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ProgressUpsert) SetUserID(v int) *ProgressUpsert {
	u.Set(progress.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProgressUpsert) UpdateUserID() *ProgressUpsert {
	u.SetExcluded(progress.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProgressUpsert) ClearUserID() *ProgressUpsert {
	u.SetNull(progress.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Progress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProgressUpsertOne) UpdateNewValues() *ProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Progress.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProgressUpsertOne) Ignore() *ProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProgressUpsertOne) DoNothing() *ProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProgressCreate.OnConflict
// documentation for more info.
func (u *ProgressUpsertOne) Update(set func(*ProgressUpsert)) *ProgressUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProgressUpsert{UpdateSet: update})
	}))
	return u
}

// SetPage sets the "page" field.
func (u *ProgressUpsertOne) SetPage(v int) *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.SetPage(v)
	})
}

// AddPage adds v to the "page" field.
func (u *ProgressUpsertOne) AddPage(v int) *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.AddPage(v)
	})
}

// UpdatePage sets the "page" field to the value that was provided on create.
func (u *ProgressUpsertOne) UpdatePage() *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdatePage()
	})
}

// SetItemID sets the "item_id" field.
func (u *ProgressUpsertOne) SetItemID(v int) *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.SetItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ProgressUpsertOne) UpdateItemID() *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdateItemID()
	})
}

// ClearItemID clears the value of the "item_id" field.
func (u *ProgressUpsertOne) ClearItemID() *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.ClearItemID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProgressUpsertOne) SetUserID(v int) *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProgressUpsertOne) UpdateUserID() *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProgressUpsertOne) ClearUserID() *ProgressUpsertOne {
	return u.Update(func(s *ProgressUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *ProgressUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProgressCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProgressUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProgressUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProgressUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProgressCreateBulk is the builder for creating many Progress entities in bulk.
type ProgressCreateBulk struct {
	config
	err      error
	builders []*ProgressCreate
	conflict []sql.ConflictOption
}

// Save creates the Progress entities in the database.
func (pcb *ProgressCreateBulk) Save(ctx context.Context) ([]*Progress, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Progress, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProgressMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProgressCreateBulk) SaveX(ctx context.Context) []*Progress {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProgressCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProgressCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Progress.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProgressUpsert) {
//			SetPage(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProgressCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProgressUpsertBulk {
	pcb.conflict = opts
	return &ProgressUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Progress.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProgressCreateBulk) OnConflictColumns(columns ...string) *ProgressUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProgressUpsertBulk{
		create: pcb,
	}
}

// ProgressUpsertBulk is the builder for "upsert"-ing
// a bulk of Progress nodes.
type ProgressUpsertBulk struct {
	create *ProgressCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Progress.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProgressUpsertBulk) UpdateNewValues() *ProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Progress.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProgressUpsertBulk) Ignore() *ProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProgressUpsertBulk) DoNothing() *ProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProgressCreateBulk.OnConflict
// documentation for more info.
func (u *ProgressUpsertBulk) Update(set func(*ProgressUpsert)) *ProgressUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProgressUpsert{UpdateSet: update})
	}))
	return u
}

// SetPage sets the "page" field.
func (u *ProgressUpsertBulk) SetPage(v int) *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.SetPage(v)
	})
}

// AddPage adds v to the "page" field.
func (u *ProgressUpsertBulk) AddPage(v int) *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.AddPage(v)
	})
}

// UpdatePage sets the "page" field to the value that was provided on create.
func (u *ProgressUpsertBulk) UpdatePage() *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdatePage()
	})
}

// SetItemID sets the "item_id" field.
func (u *ProgressUpsertBulk) SetItemID(v int) *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.SetItemID(v)
	})
}

// UpdateItemID sets the "item_id" field to the value that was provided on create.
func (u *ProgressUpsertBulk) UpdateItemID() *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdateItemID()
	})
}

// ClearItemID clears the value of the "item_id" field.
func (u *ProgressUpsertBulk) ClearItemID() *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.ClearItemID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ProgressUpsertBulk) SetUserID(v int) *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ProgressUpsertBulk) UpdateUserID() *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ProgressUpsertBulk) ClearUserID() *ProgressUpsertBulk {
	return u.Update(func(s *ProgressUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *ProgressUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProgressCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProgressCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProgressUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
