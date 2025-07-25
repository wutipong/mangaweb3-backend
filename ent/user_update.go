// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wutipong/mangaweb3-backend/ent/history"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
	"github.com/wutipong/mangaweb3-backend/ent/predicate"
	"github.com/wutipong/mangaweb3-backend/ent/progress"
	"github.com/wutipong/mangaweb3-backend/ent/tag"
	"github.com/wutipong/mangaweb3-backend/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetActive sets the "active" field.
func (uu *UserUpdate) SetActive(b bool) *UserUpdate {
	uu.mutation.SetActive(b)
	return uu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetActive(*b)
	}
	return uu
}

// AddFavoriteItemIDs adds the "favorite_items" edge to the Meta entity by IDs.
func (uu *UserUpdate) AddFavoriteItemIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoriteItemIDs(ids...)
	return uu
}

// AddFavoriteItems adds the "favorite_items" edges to the Meta entity.
func (uu *UserUpdate) AddFavoriteItems(m ...*Meta) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddFavoriteItemIDs(ids...)
}

// AddFavoriteTagIDs adds the "favorite_tags" edge to the Tag entity by IDs.
func (uu *UserUpdate) AddFavoriteTagIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoriteTagIDs(ids...)
	return uu
}

// AddFavoriteTags adds the "favorite_tags" edges to the Tag entity.
func (uu *UserUpdate) AddFavoriteTags(t ...*Tag) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddFavoriteTagIDs(ids...)
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (uu *UserUpdate) AddHistoryIDs(ids ...int) *UserUpdate {
	uu.mutation.AddHistoryIDs(ids...)
	return uu
}

// AddHistories adds the "histories" edges to the History entity.
func (uu *UserUpdate) AddHistories(h ...*History) *UserUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uu.AddHistoryIDs(ids...)
}

// AddProgresIDs adds the "progress" edge to the Progress entity by IDs.
func (uu *UserUpdate) AddProgresIDs(ids ...int) *UserUpdate {
	uu.mutation.AddProgresIDs(ids...)
	return uu
}

// AddProgress adds the "progress" edges to the Progress entity.
func (uu *UserUpdate) AddProgress(p ...*Progress) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddProgresIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearFavoriteItems clears all "favorite_items" edges to the Meta entity.
func (uu *UserUpdate) ClearFavoriteItems() *UserUpdate {
	uu.mutation.ClearFavoriteItems()
	return uu
}

// RemoveFavoriteItemIDs removes the "favorite_items" edge to Meta entities by IDs.
func (uu *UserUpdate) RemoveFavoriteItemIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoriteItemIDs(ids...)
	return uu
}

// RemoveFavoriteItems removes "favorite_items" edges to Meta entities.
func (uu *UserUpdate) RemoveFavoriteItems(m ...*Meta) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveFavoriteItemIDs(ids...)
}

// ClearFavoriteTags clears all "favorite_tags" edges to the Tag entity.
func (uu *UserUpdate) ClearFavoriteTags() *UserUpdate {
	uu.mutation.ClearFavoriteTags()
	return uu
}

// RemoveFavoriteTagIDs removes the "favorite_tags" edge to Tag entities by IDs.
func (uu *UserUpdate) RemoveFavoriteTagIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoriteTagIDs(ids...)
	return uu
}

// RemoveFavoriteTags removes "favorite_tags" edges to Tag entities.
func (uu *UserUpdate) RemoveFavoriteTags(t ...*Tag) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveFavoriteTagIDs(ids...)
}

// ClearHistories clears all "histories" edges to the History entity.
func (uu *UserUpdate) ClearHistories() *UserUpdate {
	uu.mutation.ClearHistories()
	return uu
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (uu *UserUpdate) RemoveHistoryIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveHistoryIDs(ids...)
	return uu
}

// RemoveHistories removes "histories" edges to History entities.
func (uu *UserUpdate) RemoveHistories(h ...*History) *UserUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uu.RemoveHistoryIDs(ids...)
}

// ClearProgress clears all "progress" edges to the Progress entity.
func (uu *UserUpdate) ClearProgress() *UserUpdate {
	uu.mutation.ClearProgress()
	return uu
}

// RemoveProgresIDs removes the "progress" edge to Progress entities by IDs.
func (uu *UserUpdate) RemoveProgresIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveProgresIDs(ids...)
	return uu
}

// RemoveProgress removes "progress" edges to Progress entities.
func (uu *UserUpdate) RemoveProgress(p ...*Progress) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveProgresIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if uu.mutation.FavoriteItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFavoriteItemsIDs(); len(nodes) > 0 && !uu.mutation.FavoriteItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
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
	if nodes := uu.mutation.FavoriteItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
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
	if uu.mutation.FavoriteTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFavoriteTagsIDs(); len(nodes) > 0 && !uu.mutation.FavoriteTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
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
	if nodes := uu.mutation.FavoriteTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
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
	if uu.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !uu.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
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
	if nodes := uu.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
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
	if uu.mutation.ProgressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedProgressIDs(); len(nodes) > 0 && !uu.mutation.ProgressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProgressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetActive sets the "active" field.
func (uuo *UserUpdateOne) SetActive(b bool) *UserUpdateOne {
	uuo.mutation.SetActive(b)
	return uuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetActive(*b)
	}
	return uuo
}

// AddFavoriteItemIDs adds the "favorite_items" edge to the Meta entity by IDs.
func (uuo *UserUpdateOne) AddFavoriteItemIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoriteItemIDs(ids...)
	return uuo
}

// AddFavoriteItems adds the "favorite_items" edges to the Meta entity.
func (uuo *UserUpdateOne) AddFavoriteItems(m ...*Meta) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddFavoriteItemIDs(ids...)
}

// AddFavoriteTagIDs adds the "favorite_tags" edge to the Tag entity by IDs.
func (uuo *UserUpdateOne) AddFavoriteTagIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoriteTagIDs(ids...)
	return uuo
}

// AddFavoriteTags adds the "favorite_tags" edges to the Tag entity.
func (uuo *UserUpdateOne) AddFavoriteTags(t ...*Tag) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddFavoriteTagIDs(ids...)
}

// AddHistoryIDs adds the "histories" edge to the History entity by IDs.
func (uuo *UserUpdateOne) AddHistoryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddHistoryIDs(ids...)
	return uuo
}

// AddHistories adds the "histories" edges to the History entity.
func (uuo *UserUpdateOne) AddHistories(h ...*History) *UserUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uuo.AddHistoryIDs(ids...)
}

// AddProgresIDs adds the "progress" edge to the Progress entity by IDs.
func (uuo *UserUpdateOne) AddProgresIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddProgresIDs(ids...)
	return uuo
}

// AddProgress adds the "progress" edges to the Progress entity.
func (uuo *UserUpdateOne) AddProgress(p ...*Progress) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddProgresIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearFavoriteItems clears all "favorite_items" edges to the Meta entity.
func (uuo *UserUpdateOne) ClearFavoriteItems() *UserUpdateOne {
	uuo.mutation.ClearFavoriteItems()
	return uuo
}

// RemoveFavoriteItemIDs removes the "favorite_items" edge to Meta entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoriteItemIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoriteItemIDs(ids...)
	return uuo
}

// RemoveFavoriteItems removes "favorite_items" edges to Meta entities.
func (uuo *UserUpdateOne) RemoveFavoriteItems(m ...*Meta) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveFavoriteItemIDs(ids...)
}

// ClearFavoriteTags clears all "favorite_tags" edges to the Tag entity.
func (uuo *UserUpdateOne) ClearFavoriteTags() *UserUpdateOne {
	uuo.mutation.ClearFavoriteTags()
	return uuo
}

// RemoveFavoriteTagIDs removes the "favorite_tags" edge to Tag entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoriteTagIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoriteTagIDs(ids...)
	return uuo
}

// RemoveFavoriteTags removes "favorite_tags" edges to Tag entities.
func (uuo *UserUpdateOne) RemoveFavoriteTags(t ...*Tag) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveFavoriteTagIDs(ids...)
}

// ClearHistories clears all "histories" edges to the History entity.
func (uuo *UserUpdateOne) ClearHistories() *UserUpdateOne {
	uuo.mutation.ClearHistories()
	return uuo
}

// RemoveHistoryIDs removes the "histories" edge to History entities by IDs.
func (uuo *UserUpdateOne) RemoveHistoryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveHistoryIDs(ids...)
	return uuo
}

// RemoveHistories removes "histories" edges to History entities.
func (uuo *UserUpdateOne) RemoveHistories(h ...*History) *UserUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return uuo.RemoveHistoryIDs(ids...)
}

// ClearProgress clears all "progress" edges to the Progress entity.
func (uuo *UserUpdateOne) ClearProgress() *UserUpdateOne {
	uuo.mutation.ClearProgress()
	return uuo
}

// RemoveProgresIDs removes the "progress" edge to Progress entities by IDs.
func (uuo *UserUpdateOne) RemoveProgresIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveProgresIDs(ids...)
	return uuo
}

// RemoveProgress removes "progress" edges to Progress entities.
func (uuo *UserUpdateOne) RemoveProgress(p ...*Progress) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveProgresIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if uuo.mutation.FavoriteItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(meta.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFavoriteItemsIDs(); len(nodes) > 0 && !uuo.mutation.FavoriteItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
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
	if nodes := uuo.mutation.FavoriteItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteItemsTable,
			Columns: user.FavoriteItemsPrimaryKey,
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
	if uuo.mutation.FavoriteTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFavoriteTagsIDs(); len(nodes) > 0 && !uuo.mutation.FavoriteTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
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
	if nodes := uuo.mutation.FavoriteTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteTagsTable,
			Columns: user.FavoriteTagsPrimaryKey,
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
	if uuo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedHistoriesIDs(); len(nodes) > 0 && !uuo.mutation.HistoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
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
	if nodes := uuo.mutation.HistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.HistoriesTable,
			Columns: []string{user.HistoriesColumn},
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
	if uuo.mutation.ProgressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedProgressIDs(); len(nodes) > 0 && !uuo.mutation.ProgressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProgressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProgressTable,
			Columns: []string{user.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
