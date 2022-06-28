// Code generated by entc, DO NOT EDIT.

package ent

import (
	"SafeSend/pkg/ent/accesstoken"
	"SafeSend/pkg/ent/predicate"
	"SafeSend/pkg/ent/user"
	"SafeSend/pkg/interfaces"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccessTokenUpdate is the builder for updating AccessToken entities.
type AccessTokenUpdate struct {
	config
	hooks    []Hook
	mutation *AccessTokenMutation
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atu *AccessTokenUpdate) Where(ps ...predicate.AccessToken) *AccessTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetTokenProvider sets the "token_provider" field.
func (atu *AccessTokenUpdate) SetTokenProvider(ip interfaces.TokenProvider) *AccessTokenUpdate {
	atu.mutation.SetTokenProvider(ip)
	return atu
}

// SetAccessToken sets the "access_token" field.
func (atu *AccessTokenUpdate) SetAccessToken(s string) *AccessTokenUpdate {
	atu.mutation.SetAccessToken(s)
	return atu
}

// SetRefreshToken sets the "refresh_token" field.
func (atu *AccessTokenUpdate) SetRefreshToken(s string) *AccessTokenUpdate {
	atu.mutation.SetRefreshToken(s)
	return atu
}

// SetExpiry sets the "expiry" field.
func (atu *AccessTokenUpdate) SetExpiry(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetExpiry(t)
	return atu
}

// SetDateCreated sets the "date_created" field.
func (atu *AccessTokenUpdate) SetDateCreated(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetDateCreated(t)
	return atu
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableDateCreated(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetDateCreated(*t)
	}
	return atu
}

// SetDateModified sets the "date_modified" field.
func (atu *AccessTokenUpdate) SetDateModified(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetDateModified(t)
	return atu
}

// SetNillableDateModified sets the "date_modified" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableDateModified(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetDateModified(*t)
	}
	return atu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (atu *AccessTokenUpdate) AddUserIDs(ids ...uuid.UUID) *AccessTokenUpdate {
	atu.mutation.AddUserIDs(ids...)
	return atu
}

// AddUsers adds the "users" edges to the User entity.
func (atu *AccessTokenUpdate) AddUsers(u ...*User) *AccessTokenUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atu.AddUserIDs(ids...)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atu *AccessTokenUpdate) Mutation() *AccessTokenMutation {
	return atu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (atu *AccessTokenUpdate) ClearUsers() *AccessTokenUpdate {
	atu.mutation.ClearUsers()
	return atu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (atu *AccessTokenUpdate) RemoveUserIDs(ids ...uuid.UUID) *AccessTokenUpdate {
	atu.mutation.RemoveUserIDs(ids...)
	return atu
}

// RemoveUsers removes "users" edges to User entities.
func (atu *AccessTokenUpdate) RemoveUsers(u ...*User) *AccessTokenUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AccessTokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccessTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AccessTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AccessTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AccessTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AccessTokenUpdate) check() error {
	if v, ok := atu.mutation.TokenProvider(); ok {
		if err := accesstoken.TokenProviderValidator(v); err != nil {
			return &ValidationError{Name: "token_provider", err: fmt.Errorf(`ent: validator failed for field "AccessToken.token_provider": %w`, err)}
		}
	}
	if v, ok := atu.mutation.AccessToken(); ok {
		if err := accesstoken.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "AccessToken.access_token": %w`, err)}
		}
	}
	if v, ok := atu.mutation.RefreshToken(); ok {
		if err := accesstoken.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "AccessToken.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (atu *AccessTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accesstoken.Table,
			Columns: accesstoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accesstoken.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.TokenProvider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: accesstoken.FieldTokenProvider,
		})
	}
	if value, ok := atu.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesstoken.FieldAccessToken,
		})
	}
	if value, ok := atu.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesstoken.FieldRefreshToken,
		})
	}
	if value, ok := atu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldExpiry,
		})
	}
	if value, ok := atu.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldDateCreated,
		})
	}
	if value, ok := atu.mutation.DateModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldDateModified,
		})
	}
	if atu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !atu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AccessTokenUpdateOne is the builder for updating a single AccessToken entity.
type AccessTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessTokenMutation
}

// SetTokenProvider sets the "token_provider" field.
func (atuo *AccessTokenUpdateOne) SetTokenProvider(ip interfaces.TokenProvider) *AccessTokenUpdateOne {
	atuo.mutation.SetTokenProvider(ip)
	return atuo
}

// SetAccessToken sets the "access_token" field.
func (atuo *AccessTokenUpdateOne) SetAccessToken(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetAccessToken(s)
	return atuo
}

// SetRefreshToken sets the "refresh_token" field.
func (atuo *AccessTokenUpdateOne) SetRefreshToken(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetRefreshToken(s)
	return atuo
}

// SetExpiry sets the "expiry" field.
func (atuo *AccessTokenUpdateOne) SetExpiry(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetExpiry(t)
	return atuo
}

// SetDateCreated sets the "date_created" field.
func (atuo *AccessTokenUpdateOne) SetDateCreated(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetDateCreated(t)
	return atuo
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableDateCreated(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetDateCreated(*t)
	}
	return atuo
}

// SetDateModified sets the "date_modified" field.
func (atuo *AccessTokenUpdateOne) SetDateModified(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetDateModified(t)
	return atuo
}

// SetNillableDateModified sets the "date_modified" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableDateModified(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetDateModified(*t)
	}
	return atuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (atuo *AccessTokenUpdateOne) AddUserIDs(ids ...uuid.UUID) *AccessTokenUpdateOne {
	atuo.mutation.AddUserIDs(ids...)
	return atuo
}

// AddUsers adds the "users" edges to the User entity.
func (atuo *AccessTokenUpdateOne) AddUsers(u ...*User) *AccessTokenUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atuo.AddUserIDs(ids...)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atuo *AccessTokenUpdateOne) Mutation() *AccessTokenMutation {
	return atuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (atuo *AccessTokenUpdateOne) ClearUsers() *AccessTokenUpdateOne {
	atuo.mutation.ClearUsers()
	return atuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (atuo *AccessTokenUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *AccessTokenUpdateOne {
	atuo.mutation.RemoveUserIDs(ids...)
	return atuo
}

// RemoveUsers removes "users" edges to User entities.
func (atuo *AccessTokenUpdateOne) RemoveUsers(u ...*User) *AccessTokenUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atuo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AccessTokenUpdateOne) Select(field string, fields ...string) *AccessTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AccessToken entity.
func (atuo *AccessTokenUpdateOne) Save(ctx context.Context) (*AccessToken, error) {
	var (
		err  error
		node *AccessToken
	)
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccessTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) SaveX(ctx context.Context) *AccessToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AccessTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AccessTokenUpdateOne) check() error {
	if v, ok := atuo.mutation.TokenProvider(); ok {
		if err := accesstoken.TokenProviderValidator(v); err != nil {
			return &ValidationError{Name: "token_provider", err: fmt.Errorf(`ent: validator failed for field "AccessToken.token_provider": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.AccessToken(); ok {
		if err := accesstoken.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "AccessToken.access_token": %w`, err)}
		}
	}
	if v, ok := atuo.mutation.RefreshToken(); ok {
		if err := accesstoken.RefreshTokenValidator(v); err != nil {
			return &ValidationError{Name: "refresh_token", err: fmt.Errorf(`ent: validator failed for field "AccessToken.refresh_token": %w`, err)}
		}
	}
	return nil
}

func (atuo *AccessTokenUpdateOne) sqlSave(ctx context.Context) (_node *AccessToken, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accesstoken.Table,
			Columns: accesstoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: accesstoken.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesstoken.FieldID)
		for _, f := range fields {
			if !accesstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.TokenProvider(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: accesstoken.FieldTokenProvider,
		})
	}
	if value, ok := atuo.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesstoken.FieldAccessToken,
		})
	}
	if value, ok := atuo.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesstoken.FieldRefreshToken,
		})
	}
	if value, ok := atuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldExpiry,
		})
	}
	if value, ok := atuo.mutation.DateCreated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldDateCreated,
		})
	}
	if value, ok := atuo.mutation.DateModified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: accesstoken.FieldDateModified,
		})
	}
	if atuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !atuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   accesstoken.UsersTable,
			Columns: accesstoken.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccessToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
