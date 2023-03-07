// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/ent/role"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(t time.Time) *RoleCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoleCreate) SetStatus(u uint8) *RoleCreate {
	rc.mutation.SetStatus(u)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoleCreate) SetNillableStatus(u *uint8) *RoleCreate {
	if u != nil {
		rc.SetStatus(*u)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetCode sets the "code" field.
func (rc *RoleCreate) SetCode(s string) *RoleCreate {
	rc.mutation.SetCode(s)
	return rc
}

// SetDefaultRouter sets the "default_router" field.
func (rc *RoleCreate) SetDefaultRouter(s string) *RoleCreate {
	rc.mutation.SetDefaultRouter(s)
	return rc
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDefaultRouter(s *string) *RoleCreate {
	if s != nil {
		rc.SetDefaultRouter(*s)
	}
	return rc
}

// SetRemark sets the "remark" field.
func (rc *RoleCreate) SetRemark(s string) *RoleCreate {
	rc.mutation.SetRemark(s)
	return rc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (rc *RoleCreate) SetNillableRemark(s *string) *RoleCreate {
	if s != nil {
		rc.SetRemark(*s)
	}
	return rc
}

// SetSort sets the "sort" field.
func (rc *RoleCreate) SetSort(u uint32) *RoleCreate {
	rc.mutation.SetSort(u)
	return rc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (rc *RoleCreate) SetNillableSort(u *uint32) *RoleCreate {
	if u != nil {
		rc.SetSort(*u)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(u uint64) *RoleCreate {
	rc.mutation.SetID(u)
	return rc
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (rc *RoleCreate) AddMenuIDs(ids ...uint64) *RoleCreate {
	rc.mutation.AddMenuIDs(ids...)
	return rc
}

// AddMenus adds the "menus" edges to the Menu entity.
func (rc *RoleCreate) AddMenus(m ...*Menu) *RoleCreate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rc *RoleCreate) AddUserIDs(ids ...uuid.UUID) *RoleCreate {
	rc.mutation.AddUserIDs(ids...)
	return rc
}

// AddUsers adds the "users" edges to the User entity.
func (rc *RoleCreate) AddUsers(u ...*User) *RoleCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	rc.defaults()
	return withHooks[*Role, RoleMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := role.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := role.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := role.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.DefaultRouter(); !ok {
		v := role.DefaultDefaultRouter
		rc.mutation.SetDefaultRouter(v)
	}
	if _, ok := rc.mutation.Remark(); !ok {
		v := role.DefaultRemark
		rc.mutation.SetRemark(v)
	}
	if _, ok := rc.mutation.Sort(); !ok {
		v := role.DefaultSort
		rc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Role.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Role.updated_at"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Role.name"`)}
	}
	if _, ok := rc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Role.code"`)}
	}
	if _, ok := rc.mutation.DefaultRouter(); !ok {
		return &ValidationError{Name: "default_router", err: errors.New(`ent: missing required field "Role.default_router"`)}
	}
	if _, ok := rc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "Role.remark"`)}
	}
	if _, ok := rc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Role.sort"`)}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rc.mutation.DefaultRouter(); ok {
		_spec.SetField(role.FieldDefaultRouter, field.TypeString, value)
		_node.DefaultRouter = value
	}
	if value, ok := rc.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := rc.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if nodes := rc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	builders []*RoleCreate
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
