// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/category"
	"github.com/Skijetler/alphinium/pkg/ent/subcategory"
	"github.com/Skijetler/alphinium/pkg/ent/thread"
)

// SubcategoryCreate is the builder for creating a Subcategory entity.
type SubcategoryCreate struct {
	config
	mutation *SubcategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SubcategoryCreate) SetName(s string) *SubcategoryCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SubcategoryCreate) SetDescription(s string) *SubcategoryCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetCategoryID sets the "category_id" field.
func (sc *SubcategoryCreate) SetCategoryID(u uint64) *SubcategoryCreate {
	sc.mutation.SetCategoryID(u)
	return sc
}

// SetID sets the "id" field.
func (sc *SubcategoryCreate) SetID(u uint64) *SubcategoryCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetCategory sets the "category" edge to the Category entity.
func (sc *SubcategoryCreate) SetCategory(c *Category) *SubcategoryCreate {
	return sc.SetCategoryID(c.ID)
}

// AddThreadIDs adds the "threads" edge to the Thread entity by IDs.
func (sc *SubcategoryCreate) AddThreadIDs(ids ...uint64) *SubcategoryCreate {
	sc.mutation.AddThreadIDs(ids...)
	return sc
}

// AddThreads adds the "threads" edges to the Thread entity.
func (sc *SubcategoryCreate) AddThreads(t ...*Thread) *SubcategoryCreate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddThreadIDs(ids...)
}

// Mutation returns the SubcategoryMutation object of the builder.
func (sc *SubcategoryCreate) Mutation() *SubcategoryMutation {
	return sc.mutation
}

// Save creates the Subcategory in the database.
func (sc *SubcategoryCreate) Save(ctx context.Context) (*Subcategory, error) {
	var (
		err  error
		node *Subcategory
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubcategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Subcategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubcategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubcategoryCreate) SaveX(ctx context.Context) *Subcategory {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubcategoryCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubcategoryCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubcategoryCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subcategory.name"`)}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Subcategory.description"`)}
	}
	if _, ok := sc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "Subcategory.category_id"`)}
	}
	if _, ok := sc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "Subcategory.category"`)}
	}
	return nil
}

func (sc *SubcategoryCreate) sqlSave(ctx context.Context) (*Subcategory, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (sc *SubcategoryCreate) createSpec() (*Subcategory, *sqlgraph.CreateSpec) {
	var (
		_node = &Subcategory{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subcategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: subcategory.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subcategory.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subcategory.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := sc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ThreadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ThreadsTable,
			Columns: []string{subcategory.ThreadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: thread.FieldID,
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

// SubcategoryCreateBulk is the builder for creating many Subcategory entities in bulk.
type SubcategoryCreateBulk struct {
	config
	builders []*SubcategoryCreate
}

// Save creates the Subcategory entities in the database.
func (scb *SubcategoryCreateBulk) Save(ctx context.Context) ([]*Subcategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subcategory, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubcategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubcategoryCreateBulk) SaveX(ctx context.Context) []*Subcategory {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubcategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubcategoryCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
