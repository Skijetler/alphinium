// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Skijetler/alphinium/pkg/ent/post"
	"github.com/Skijetler/alphinium/pkg/ent/predicate"
	"github.com/Skijetler/alphinium/pkg/ent/subcategory"
	"github.com/Skijetler/alphinium/pkg/ent/thread"
)

// ThreadUpdate is the builder for updating Thread entities.
type ThreadUpdate struct {
	config
	hooks    []Hook
	mutation *ThreadMutation
}

// Where appends a list predicates to the ThreadUpdate builder.
func (tu *ThreadUpdate) Where(ps ...predicate.Thread) *ThreadUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *ThreadUpdate) SetName(s string) *ThreadUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescriptionID sets the "description_id" field.
func (tu *ThreadUpdate) SetDescriptionID(u uint64) *ThreadUpdate {
	tu.mutation.SetDescriptionID(u)
	return tu
}

// SetSubcategoryID sets the "subcategory_id" field.
func (tu *ThreadUpdate) SetSubcategoryID(u uint64) *ThreadUpdate {
	tu.mutation.SetSubcategoryID(u)
	return tu
}

// SetSubcategory sets the "subcategory" edge to the Subcategory entity.
func (tu *ThreadUpdate) SetSubcategory(s *Subcategory) *ThreadUpdate {
	return tu.SetSubcategoryID(s.ID)
}

// SetDescription sets the "description" edge to the Post entity.
func (tu *ThreadUpdate) SetDescription(p *Post) *ThreadUpdate {
	return tu.SetDescriptionID(p.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (tu *ThreadUpdate) AddPostIDs(ids ...uint64) *ThreadUpdate {
	tu.mutation.AddPostIDs(ids...)
	return tu
}

// AddPosts adds the "posts" edges to the Post entity.
func (tu *ThreadUpdate) AddPosts(p ...*Post) *ThreadUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddPostIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tu *ThreadUpdate) Mutation() *ThreadMutation {
	return tu.mutation
}

// ClearSubcategory clears the "subcategory" edge to the Subcategory entity.
func (tu *ThreadUpdate) ClearSubcategory() *ThreadUpdate {
	tu.mutation.ClearSubcategory()
	return tu
}

// ClearDescription clears the "description" edge to the Post entity.
func (tu *ThreadUpdate) ClearDescription() *ThreadUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// ClearPosts clears all "posts" edges to the Post entity.
func (tu *ThreadUpdate) ClearPosts() *ThreadUpdate {
	tu.mutation.ClearPosts()
	return tu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (tu *ThreadUpdate) RemovePostIDs(ids ...uint64) *ThreadUpdate {
	tu.mutation.RemovePostIDs(ids...)
	return tu
}

// RemovePosts removes "posts" edges to Post entities.
func (tu *ThreadUpdate) RemovePosts(p ...*Post) *ThreadUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemovePostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThreadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThreadUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThreadUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThreadUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *ThreadUpdate) check() error {
	if _, ok := tu.mutation.SubcategoryID(); tu.mutation.SubcategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.subcategory"`)
	}
	if _, ok := tu.mutation.DescriptionID(); tu.mutation.DescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.description"`)
	}
	return nil
}

func (tu *ThreadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thread.Table,
			Columns: thread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: thread.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thread.FieldName,
		})
	}
	if tu.mutation.SubcategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.SubcategoryTable,
			Columns: []string{thread.SubcategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: subcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SubcategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.SubcategoryTable,
			Columns: []string{thread.SubcategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: subcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.DescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.DescriptionTable,
			Columns: []string{thread.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.DescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.DescriptionTable,
			Columns: []string{thread.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !tu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ThreadUpdateOne is the builder for updating a single Thread entity.
type ThreadUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThreadMutation
}

// SetName sets the "name" field.
func (tuo *ThreadUpdateOne) SetName(s string) *ThreadUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescriptionID sets the "description_id" field.
func (tuo *ThreadUpdateOne) SetDescriptionID(u uint64) *ThreadUpdateOne {
	tuo.mutation.SetDescriptionID(u)
	return tuo
}

// SetSubcategoryID sets the "subcategory_id" field.
func (tuo *ThreadUpdateOne) SetSubcategoryID(u uint64) *ThreadUpdateOne {
	tuo.mutation.SetSubcategoryID(u)
	return tuo
}

// SetSubcategory sets the "subcategory" edge to the Subcategory entity.
func (tuo *ThreadUpdateOne) SetSubcategory(s *Subcategory) *ThreadUpdateOne {
	return tuo.SetSubcategoryID(s.ID)
}

// SetDescription sets the "description" edge to the Post entity.
func (tuo *ThreadUpdateOne) SetDescription(p *Post) *ThreadUpdateOne {
	return tuo.SetDescriptionID(p.ID)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (tuo *ThreadUpdateOne) AddPostIDs(ids ...uint64) *ThreadUpdateOne {
	tuo.mutation.AddPostIDs(ids...)
	return tuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (tuo *ThreadUpdateOne) AddPosts(p ...*Post) *ThreadUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddPostIDs(ids...)
}

// Mutation returns the ThreadMutation object of the builder.
func (tuo *ThreadUpdateOne) Mutation() *ThreadMutation {
	return tuo.mutation
}

// ClearSubcategory clears the "subcategory" edge to the Subcategory entity.
func (tuo *ThreadUpdateOne) ClearSubcategory() *ThreadUpdateOne {
	tuo.mutation.ClearSubcategory()
	return tuo
}

// ClearDescription clears the "description" edge to the Post entity.
func (tuo *ThreadUpdateOne) ClearDescription() *ThreadUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// ClearPosts clears all "posts" edges to the Post entity.
func (tuo *ThreadUpdateOne) ClearPosts() *ThreadUpdateOne {
	tuo.mutation.ClearPosts()
	return tuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (tuo *ThreadUpdateOne) RemovePostIDs(ids ...uint64) *ThreadUpdateOne {
	tuo.mutation.RemovePostIDs(ids...)
	return tuo
}

// RemovePosts removes "posts" edges to Post entities.
func (tuo *ThreadUpdateOne) RemovePosts(p ...*Post) *ThreadUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemovePostIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ThreadUpdateOne) Select(field string, fields ...string) *ThreadUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Thread entity.
func (tuo *ThreadUpdateOne) Save(ctx context.Context) (*Thread, error) {
	var (
		err  error
		node *Thread
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThreadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Thread)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ThreadMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThreadUpdateOne) SaveX(ctx context.Context) *Thread {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThreadUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThreadUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *ThreadUpdateOne) check() error {
	if _, ok := tuo.mutation.SubcategoryID(); tuo.mutation.SubcategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.subcategory"`)
	}
	if _, ok := tuo.mutation.DescriptionID(); tuo.mutation.DescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Thread.description"`)
	}
	return nil
}

func (tuo *ThreadUpdateOne) sqlSave(ctx context.Context) (_node *Thread, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   thread.Table,
			Columns: thread.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: thread.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Thread.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thread.FieldID)
		for _, f := range fields {
			if !thread.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != thread.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: thread.FieldName,
		})
	}
	if tuo.mutation.SubcategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.SubcategoryTable,
			Columns: []string{thread.SubcategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: subcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SubcategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thread.SubcategoryTable,
			Columns: []string{thread.SubcategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: subcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.DescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.DescriptionTable,
			Columns: []string{thread.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.DescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   thread.DescriptionTable,
			Columns: []string{thread.DescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !tuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   thread.PostsTable,
			Columns: []string{thread.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Thread{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{thread.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
