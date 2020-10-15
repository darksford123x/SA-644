// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/darksford123x/repairinvoice-record/ent/device"
	"github.com/darksford123x/repairinvoice-record/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUname sets the Uname field.
func (uc *UserCreate) SetUname(s string) *UserCreate {
	uc.mutation.SetUname(s)
	return uc
}

// AddDeviceIDs adds the device edge to Device by ids.
func (uc *UserCreate) AddDeviceIDs(ids ...int) *UserCreate {
	uc.mutation.AddDeviceIDs(ids...)
	return uc
}

// AddDevice adds the device edges to Device.
func (uc *UserCreate) AddDevice(d ...*Device) *UserCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDeviceIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.Uname(); !ok {
		return nil, &ValidationError{Name: "Uname", err: errors.New("ent: missing required field \"Uname\"")}
	}
	if v, ok := uc.mutation.Uname(); ok {
		if err := user.UnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Uname", err: fmt.Errorf("ent: validator failed for field \"Uname\": %w", err)}
		}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Uname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUname,
		})
		u.Uname = value
	}
	if nodes := uc.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DeviceTable,
			Columns: []string{user.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: device.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}
