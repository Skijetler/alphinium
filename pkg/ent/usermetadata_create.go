// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/user"
	"github.com/Skijetler/alphinium/pkg/ent/usermetadata"
)

// UserMetadataCreate is the builder for creating a UserMetadata entity.
type UserMetadataCreate struct {
	config
	mutation *UserMetadataMutation
	hooks    []Hook
}

// SetNameColor sets the "name_color" field.
func (umc *UserMetadataCreate) SetNameColor(s string) *UserMetadataCreate {
	umc.mutation.SetNameColor(s)
	return umc
}

// SetTitle sets the "title" field.
func (umc *UserMetadataCreate) SetTitle(s string) *UserMetadataCreate {
	umc.mutation.SetTitle(s)
	return umc
}

// SetGender sets the "gender" field.
func (umc *UserMetadataCreate) SetGender(s string) *UserMetadataCreate {
	umc.mutation.SetGender(s)
	return umc
}

// SetLastOnline sets the "last_online" field.
func (umc *UserMetadataCreate) SetLastOnline(t time.Time) *UserMetadataCreate {
	umc.mutation.SetLastOnline(t)
	return umc
}

// SetUserID sets the "user_id" field.
func (umc *UserMetadataCreate) SetUserID(u uint64) *UserMetadataCreate {
	umc.mutation.SetUserID(u)
	return umc
}

// SetID sets the "id" field.
func (umc *UserMetadataCreate) SetID(u uint64) *UserMetadataCreate {
	umc.mutation.SetID(u)
	return umc
}

// SetUser sets the "user" edge to the User entity.
func (umc *UserMetadataCreate) SetUser(u *User) *UserMetadataCreate {
	return umc.SetUserID(u.ID)
}

// Mutation returns the UserMetadataMutation object of the builder.
func (umc *UserMetadataCreate) Mutation() *UserMetadataMutation {
	return umc.mutation
}

// Save creates the UserMetadata in the database.
func (umc *UserMetadataCreate) Save(ctx context.Context) (*UserMetadata, error) {
	var (
		err  error
		node *UserMetadata
	)
	if len(umc.hooks) == 0 {
		if err = umc.check(); err != nil {
			return nil, err
		}
		node, err = umc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMetadataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = umc.check(); err != nil {
				return nil, err
			}
			umc.mutation = mutation
			if node, err = umc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(umc.hooks) - 1; i >= 0; i-- {
			if umc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = umc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, umc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserMetadata)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMetadataMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (umc *UserMetadataCreate) SaveX(ctx context.Context) *UserMetadata {
	v, err := umc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umc *UserMetadataCreate) Exec(ctx context.Context) error {
	_, err := umc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umc *UserMetadataCreate) ExecX(ctx context.Context) {
	if err := umc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (umc *UserMetadataCreate) check() error {
	if _, ok := umc.mutation.NameColor(); !ok {
		return &ValidationError{Name: "name_color", err: errors.New(`ent: missing required field "UserMetadata.name_color"`)}
	}
	if _, ok := umc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "UserMetadata.title"`)}
	}
	if _, ok := umc.mutation.Gender(); !ok {
		return &ValidationError{Name: "gender", err: errors.New(`ent: missing required field "UserMetadata.gender"`)}
	}
	if _, ok := umc.mutation.LastOnline(); !ok {
		return &ValidationError{Name: "last_online", err: errors.New(`ent: missing required field "UserMetadata.last_online"`)}
	}
	if _, ok := umc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserMetadata.user_id"`)}
	}
	if _, ok := umc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserMetadata.user"`)}
	}
	return nil
}

func (umc *UserMetadataCreate) sqlSave(ctx context.Context) (*UserMetadata, error) {
	_node, _spec := umc.createSpec()
	if err := sqlgraph.CreateNode(ctx, umc.driver, _spec); err != nil {
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

func (umc *UserMetadataCreate) createSpec() (*UserMetadata, *sqlgraph.CreateSpec) {
	var (
		_node = &UserMetadata{config: umc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usermetadata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: usermetadata.FieldID,
			},
		}
	)
	if id, ok := umc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := umc.mutation.NameColor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usermetadata.FieldNameColor,
		})
		_node.NameColor = value
	}
	if value, ok := umc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usermetadata.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := umc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usermetadata.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := umc.mutation.LastOnline(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usermetadata.FieldLastOnline,
		})
		_node.LastOnline = value
	}
	if nodes := umc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usermetadata.UserTable,
			Columns: []string{usermetadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
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

// UserMetadataCreateBulk is the builder for creating many UserMetadata entities in bulk.
type UserMetadataCreateBulk struct {
	config
	builders []*UserMetadataCreate
}

// Save creates the UserMetadata entities in the database.
func (umcb *UserMetadataCreateBulk) Save(ctx context.Context) ([]*UserMetadata, error) {
	specs := make([]*sqlgraph.CreateSpec, len(umcb.builders))
	nodes := make([]*UserMetadata, len(umcb.builders))
	mutators := make([]Mutator, len(umcb.builders))
	for i := range umcb.builders {
		func(i int, root context.Context) {
			builder := umcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, umcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, umcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, umcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (umcb *UserMetadataCreateBulk) SaveX(ctx context.Context) []*UserMetadata {
	v, err := umcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (umcb *UserMetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := umcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (umcb *UserMetadataCreateBulk) ExecX(ctx context.Context) {
	if err := umcb.Exec(ctx); err != nil {
		panic(err)
	}
}
