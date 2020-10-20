// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/darksford123x/repairinvoice-record/ent/predicate"
	"github.com/darksford123x/repairinvoice-record/ent/repairinvoice"
	"github.com/darksford123x/repairinvoice-record/ent/statusr"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// StatusRUpdate is the builder for updating StatusR entities.
type StatusRUpdate struct {
	config
	hooks      []Hook
	mutation   *StatusRMutation
	predicates []predicate.StatusR
}

// Where adds a new predicate for the builder.
func (sr *StatusRUpdate) Where(ps ...predicate.StatusR) *StatusRUpdate {
	sr.predicates = append(sr.predicates, ps...)
	return sr
}

// SetSname sets the Sname field.
func (sr *StatusRUpdate) SetSname(s string) *StatusRUpdate {
	sr.mutation.SetSname(s)
	return sr
}

// AddRepairInformationIDs adds the repair_information edge to RepairInvoice by ids.
func (sr *StatusRUpdate) AddRepairInformationIDs(ids ...int) *StatusRUpdate {
	sr.mutation.AddRepairInformationIDs(ids...)
	return sr
}

// AddRepairInformation adds the repair_information edges to RepairInvoice.
func (sr *StatusRUpdate) AddRepairInformation(r ...*RepairInvoice) *StatusRUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sr.AddRepairInformationIDs(ids...)
}

// Mutation returns the StatusRMutation object of the builder.
func (sr *StatusRUpdate) Mutation() *StatusRMutation {
	return sr.mutation
}

// RemoveRepairInformationIDs removes the repair_information edge to RepairInvoice by ids.
func (sr *StatusRUpdate) RemoveRepairInformationIDs(ids ...int) *StatusRUpdate {
	sr.mutation.RemoveRepairInformationIDs(ids...)
	return sr
}

// RemoveRepairInformation removes repair_information edges to RepairInvoice.
func (sr *StatusRUpdate) RemoveRepairInformation(r ...*RepairInvoice) *StatusRUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sr.RemoveRepairInformationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (sr *StatusRUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(sr.hooks) == 0 {
		affected, err = sr.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusRMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sr.mutation = mutation
			affected, err = sr.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sr.hooks) - 1; i >= 0; i-- {
			mut = sr.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sr.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sr *StatusRUpdate) SaveX(ctx context.Context) int {
	affected, err := sr.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sr *StatusRUpdate) Exec(ctx context.Context) error {
	_, err := sr.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sr *StatusRUpdate) ExecX(ctx context.Context) {
	if err := sr.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sr *StatusRUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusr.Table,
			Columns: statusr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusr.FieldID,
			},
		},
	}
	if ps := sr.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sr.mutation.Sname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statusr.FieldSname,
		})
	}
	if nodes := sr.mutation.RemovedRepairInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusr.RepairInformationTable,
			Columns: []string{statusr.RepairInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sr.mutation.RepairInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusr.RepairInformationTable,
			Columns: []string{statusr.RepairInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sr.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusr.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StatusRUpdateOne is the builder for updating a single StatusR entity.
type StatusRUpdateOne struct {
	config
	hooks    []Hook
	mutation *StatusRMutation
}

// SetSname sets the Sname field.
func (sro *StatusRUpdateOne) SetSname(s string) *StatusRUpdateOne {
	sro.mutation.SetSname(s)
	return sro
}

// AddRepairInformationIDs adds the repair_information edge to RepairInvoice by ids.
func (sro *StatusRUpdateOne) AddRepairInformationIDs(ids ...int) *StatusRUpdateOne {
	sro.mutation.AddRepairInformationIDs(ids...)
	return sro
}

// AddRepairInformation adds the repair_information edges to RepairInvoice.
func (sro *StatusRUpdateOne) AddRepairInformation(r ...*RepairInvoice) *StatusRUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sro.AddRepairInformationIDs(ids...)
}

// Mutation returns the StatusRMutation object of the builder.
func (sro *StatusRUpdateOne) Mutation() *StatusRMutation {
	return sro.mutation
}

// RemoveRepairInformationIDs removes the repair_information edge to RepairInvoice by ids.
func (sro *StatusRUpdateOne) RemoveRepairInformationIDs(ids ...int) *StatusRUpdateOne {
	sro.mutation.RemoveRepairInformationIDs(ids...)
	return sro
}

// RemoveRepairInformation removes repair_information edges to RepairInvoice.
func (sro *StatusRUpdateOne) RemoveRepairInformation(r ...*RepairInvoice) *StatusRUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return sro.RemoveRepairInformationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (sro *StatusRUpdateOne) Save(ctx context.Context) (*StatusR, error) {

	var (
		err  error
		node *StatusR
	)
	if len(sro.hooks) == 0 {
		node, err = sro.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusRMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sro.mutation = mutation
			node, err = sro.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sro.hooks) - 1; i >= 0; i-- {
			mut = sro.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sro.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sro *StatusRUpdateOne) SaveX(ctx context.Context) *StatusR {
	s, err := sro.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (sro *StatusRUpdateOne) Exec(ctx context.Context) error {
	_, err := sro.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sro *StatusRUpdateOne) ExecX(ctx context.Context) {
	if err := sro.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sro *StatusRUpdateOne) sqlSave(ctx context.Context) (s *StatusR, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusr.Table,
			Columns: statusr.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusr.FieldID,
			},
		},
	}
	id, ok := sro.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing StatusR.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := sro.mutation.Sname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statusr.FieldSname,
		})
	}
	if nodes := sro.mutation.RemovedRepairInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusr.RepairInformationTable,
			Columns: []string{statusr.RepairInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sro.mutation.RepairInformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statusr.RepairInformationTable,
			Columns: []string{statusr.RepairInformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: repairinvoice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &StatusR{config: sro.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, sro.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statusr.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
