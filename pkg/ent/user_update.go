// Code generated by entc, DO NOT EDIT.

package ent

import (
	"SafeSend/pkg/ent/accesstoken"
	"SafeSend/pkg/ent/entity"
	"SafeSend/pkg/ent/group"
	"SafeSend/pkg/ent/predicate"
	"SafeSend/pkg/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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

// SetDisplayName sets the "displayName" field.
func (uu *UserUpdate) SetDisplayName(s string) *UserUpdate {
	uu.mutation.SetDisplayName(s)
	return uu
}

// SetDateAccessed sets the "dateAccessed" field.
func (uu *UserUpdate) SetDateAccessed(t time.Time) *UserUpdate {
	uu.mutation.SetDateAccessed(t)
	return uu
}

// SetNillableDateAccessed sets the "dateAccessed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDateAccessed(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDateAccessed(*t)
	}
	return uu
}

// ClearDateAccessed clears the value of the "dateAccessed" field.
func (uu *UserUpdate) ClearDateAccessed() *UserUpdate {
	uu.mutation.ClearDateAccessed()
	return uu
}

// SetDateCreated sets the "dateCreated" field.
func (uu *UserUpdate) SetDateCreated(t time.Time) *UserUpdate {
	uu.mutation.SetDateCreated(t)
	return uu
}

// SetNillableDateCreated sets the "dateCreated" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDateCreated(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDateCreated(*t)
	}
	return uu
}

// SetDateModified sets the "dateModified" field.
func (uu *UserUpdate) SetDateModified(t time.Time) *UserUpdate {
	uu.mutation.SetDateModified(t)
	return uu
}

// SetNillableDateModified sets the "dateModified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDateModified(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDateModified(*t)
	}
	return uu
}

// SetDeletedDate sets the "deletedDate" field.
func (uu *UserUpdate) SetDeletedDate(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedDate(t)
	return uu
}

// SetNillableDeletedDate sets the "deletedDate" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedDate(*t)
	}
	return uu
}

// ClearDeletedDate clears the value of the "deletedDate" field.
func (uu *UserUpdate) ClearDeletedDate() *UserUpdate {
	uu.mutation.ClearDeletedDate()
	return uu
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// SetEntitiesID sets the "entities" edge to the Entity entity by ID.
func (uu *UserUpdate) SetEntitiesID(id uuid.UUID) *UserUpdate {
	uu.mutation.SetEntitiesID(id)
	return uu
}

// SetNillableEntitiesID sets the "entities" edge to the Entity entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableEntitiesID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetEntitiesID(*id)
	}
	return uu
}

// SetEntities sets the "entities" edge to the Entity entity.
func (uu *UserUpdate) SetEntities(e *Entity) *UserUpdate {
	return uu.SetEntitiesID(e.ID)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the AccessToken entity by IDs.
func (uu *UserUpdate) AddAccessTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddAccessTokenIDs(ids...)
	return uu
}

// AddAccessTokens adds the "access_tokens" edges to the AccessToken entity.
func (uu *UserUpdate) AddAccessTokens(a ...*AccessToken) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAccessTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// ClearEntities clears the "entities" edge to the Entity entity.
func (uu *UserUpdate) ClearEntities() *UserUpdate {
	uu.mutation.ClearEntities()
	return uu
}

// ClearAccessTokens clears all "access_tokens" edges to the AccessToken entity.
func (uu *UserUpdate) ClearAccessTokens() *UserUpdate {
	uu.mutation.ClearAccessTokens()
	return uu
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to AccessToken entities by IDs.
func (uu *UserUpdate) RemoveAccessTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveAccessTokenIDs(ids...)
	return uu
}

// RemoveAccessTokens removes "access_tokens" edges to AccessToken entities.
func (uu *UserUpdate) RemoveAccessTokens(a ...*AccessToken) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAccessTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
	}
	if value, ok := uu.mutation.DateAccessed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateAccessed,
		})
	}
	if uu.mutation.DateAccessedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldDateAccessed,
		})
	}
	if value, ok := uu.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateCreated,
		})
	}
	if value, ok := uu.mutation.DateModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateModified,
		})
	}
	if value, ok := uu.mutation.DeletedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDeletedDate,
		})
	}
	if uu.mutation.DeletedDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldDeletedDate,
		})
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EntitiesTable,
			Columns: []string{user.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EntitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EntitiesTable,
			Columns: []string{user.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !uu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
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
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
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

// SetDisplayName sets the "displayName" field.
func (uuo *UserUpdateOne) SetDisplayName(s string) *UserUpdateOne {
	uuo.mutation.SetDisplayName(s)
	return uuo
}

// SetDateAccessed sets the "dateAccessed" field.
func (uuo *UserUpdateOne) SetDateAccessed(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDateAccessed(t)
	return uuo
}

// SetNillableDateAccessed sets the "dateAccessed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDateAccessed(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDateAccessed(*t)
	}
	return uuo
}

// ClearDateAccessed clears the value of the "dateAccessed" field.
func (uuo *UserUpdateOne) ClearDateAccessed() *UserUpdateOne {
	uuo.mutation.ClearDateAccessed()
	return uuo
}

// SetDateCreated sets the "dateCreated" field.
func (uuo *UserUpdateOne) SetDateCreated(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDateCreated(t)
	return uuo
}

// SetNillableDateCreated sets the "dateCreated" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDateCreated(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDateCreated(*t)
	}
	return uuo
}

// SetDateModified sets the "dateModified" field.
func (uuo *UserUpdateOne) SetDateModified(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDateModified(t)
	return uuo
}

// SetNillableDateModified sets the "dateModified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDateModified(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDateModified(*t)
	}
	return uuo
}

// SetDeletedDate sets the "deletedDate" field.
func (uuo *UserUpdateOne) SetDeletedDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedDate(t)
	return uuo
}

// SetNillableDeletedDate sets the "deletedDate" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedDate(*t)
	}
	return uuo
}

// ClearDeletedDate clears the value of the "deletedDate" field.
func (uuo *UserUpdateOne) ClearDeletedDate() *UserUpdateOne {
	uuo.mutation.ClearDeletedDate()
	return uuo
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// SetEntitiesID sets the "entities" edge to the Entity entity by ID.
func (uuo *UserUpdateOne) SetEntitiesID(id uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetEntitiesID(id)
	return uuo
}

// SetNillableEntitiesID sets the "entities" edge to the Entity entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEntitiesID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetEntitiesID(*id)
	}
	return uuo
}

// SetEntities sets the "entities" edge to the Entity entity.
func (uuo *UserUpdateOne) SetEntities(e *Entity) *UserUpdateOne {
	return uuo.SetEntitiesID(e.ID)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the AccessToken entity by IDs.
func (uuo *UserUpdateOne) AddAccessTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddAccessTokenIDs(ids...)
	return uuo
}

// AddAccessTokens adds the "access_tokens" edges to the AccessToken entity.
func (uuo *UserUpdateOne) AddAccessTokens(a ...*AccessToken) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAccessTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// ClearEntities clears the "entities" edge to the Entity entity.
func (uuo *UserUpdateOne) ClearEntities() *UserUpdateOne {
	uuo.mutation.ClearEntities()
	return uuo
}

// ClearAccessTokens clears all "access_tokens" edges to the AccessToken entity.
func (uuo *UserUpdateOne) ClearAccessTokens() *UserUpdateOne {
	uuo.mutation.ClearAccessTokens()
	return uuo
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to AccessToken entities by IDs.
func (uuo *UserUpdateOne) RemoveAccessTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveAccessTokenIDs(ids...)
	return uuo
}

// RemoveAccessTokens removes "access_tokens" edges to AccessToken entities.
func (uuo *UserUpdateOne) RemoveAccessTokens(a ...*AccessToken) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAccessTokenIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
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

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
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
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.DisplayName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldDisplayName,
		})
	}
	if value, ok := uuo.mutation.DateAccessed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateAccessed,
		})
	}
	if uuo.mutation.DateAccessedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldDateAccessed,
		})
	}
	if value, ok := uuo.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateCreated,
		})
	}
	if value, ok := uuo.mutation.DateModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDateModified,
		})
	}
	if value, ok := uuo.mutation.DeletedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldDeletedDate,
		})
	}
	if uuo.mutation.DeletedDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldDeletedDate,
		})
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: []string{user.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.EntitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EntitiesTable,
			Columns: []string{user.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entity.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EntitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.EntitiesTable,
			Columns: []string{user.EntitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: entity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !uuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.AccessTokensTable,
			Columns: user.AccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: accesstoken.FieldID,
				},
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
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
