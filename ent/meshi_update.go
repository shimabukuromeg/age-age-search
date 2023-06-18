// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// MeshiUpdate is the builder for updating Meshi entities.
type MeshiUpdate struct {
	config
	hooks    []Hook
	mutation *MeshiMutation
}

// Where appends a list predicates to the MeshiUpdate builder.
func (mu *MeshiUpdate) Where(ps ...predicate.Meshi) *MeshiUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetArticleID sets the "article_id" field.
func (mu *MeshiUpdate) SetArticleID(s string) *MeshiUpdate {
	mu.mutation.SetArticleID(s)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MeshiUpdate) SetTitle(s string) *MeshiUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableTitle(s *string) *MeshiUpdate {
	if s != nil {
		mu.SetTitle(*s)
	}
	return mu
}

// SetImageURL sets the "image_url" field.
func (mu *MeshiUpdate) SetImageURL(s string) *MeshiUpdate {
	mu.mutation.SetImageURL(s)
	return mu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableImageURL(s *string) *MeshiUpdate {
	if s != nil {
		mu.SetImageURL(*s)
	}
	return mu
}

// SetStoreName sets the "store_name" field.
func (mu *MeshiUpdate) SetStoreName(s string) *MeshiUpdate {
	mu.mutation.SetStoreName(s)
	return mu
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableStoreName(s *string) *MeshiUpdate {
	if s != nil {
		mu.SetStoreName(*s)
	}
	return mu
}

// SetAddress sets the "address" field.
func (mu *MeshiUpdate) SetAddress(s string) *MeshiUpdate {
	mu.mutation.SetAddress(s)
	return mu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableAddress(s *string) *MeshiUpdate {
	if s != nil {
		mu.SetAddress(*s)
	}
	return mu
}

// SetSiteURL sets the "site_url" field.
func (mu *MeshiUpdate) SetSiteURL(s string) *MeshiUpdate {
	mu.mutation.SetSiteURL(s)
	return mu
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableSiteURL(s *string) *MeshiUpdate {
	if s != nil {
		mu.SetSiteURL(*s)
	}
	return mu
}

// SetPublishedDate sets the "published_date" field.
func (mu *MeshiUpdate) SetPublishedDate(t time.Time) *MeshiUpdate {
	mu.mutation.SetPublishedDate(t)
	return mu
}

// SetLatitude sets the "latitude" field.
func (mu *MeshiUpdate) SetLatitude(f float64) *MeshiUpdate {
	mu.mutation.ResetLatitude()
	mu.mutation.SetLatitude(f)
	return mu
}

// AddLatitude adds f to the "latitude" field.
func (mu *MeshiUpdate) AddLatitude(f float64) *MeshiUpdate {
	mu.mutation.AddLatitude(f)
	return mu
}

// SetLongitude sets the "longitude" field.
func (mu *MeshiUpdate) SetLongitude(f float64) *MeshiUpdate {
	mu.mutation.ResetLongitude()
	mu.mutation.SetLongitude(f)
	return mu
}

// AddLongitude adds f to the "longitude" field.
func (mu *MeshiUpdate) AddLongitude(f float64) *MeshiUpdate {
	mu.mutation.AddLongitude(f)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MeshiUpdate) SetCreatedAt(t time.Time) *MeshiUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MeshiUpdate) SetNillableCreatedAt(t *time.Time) *MeshiUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetMunicipalityID sets the "municipality" edge to the Municipality entity by ID.
func (mu *MeshiUpdate) SetMunicipalityID(id int) *MeshiUpdate {
	mu.mutation.SetMunicipalityID(id)
	return mu
}

// SetNillableMunicipalityID sets the "municipality" edge to the Municipality entity by ID if the given value is not nil.
func (mu *MeshiUpdate) SetNillableMunicipalityID(id *int) *MeshiUpdate {
	if id != nil {
		mu = mu.SetMunicipalityID(*id)
	}
	return mu
}

// SetMunicipality sets the "municipality" edge to the Municipality entity.
func (mu *MeshiUpdate) SetMunicipality(m *Municipality) *MeshiUpdate {
	return mu.SetMunicipalityID(m.ID)
}

// Mutation returns the MeshiMutation object of the builder.
func (mu *MeshiUpdate) Mutation() *MeshiMutation {
	return mu.mutation
}

// ClearMunicipality clears the "municipality" edge to the Municipality entity.
func (mu *MeshiUpdate) ClearMunicipality() *MeshiUpdate {
	mu.mutation.ClearMunicipality()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MeshiUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MeshiUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MeshiUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MeshiUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MeshiUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(meshi.Table, meshi.Columns, sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.ArticleID(); ok {
		_spec.SetField(meshi.FieldArticleID, field.TypeString, value)
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.SetField(meshi.FieldTitle, field.TypeString, value)
	}
	if value, ok := mu.mutation.ImageURL(); ok {
		_spec.SetField(meshi.FieldImageURL, field.TypeString, value)
	}
	if value, ok := mu.mutation.StoreName(); ok {
		_spec.SetField(meshi.FieldStoreName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Address(); ok {
		_spec.SetField(meshi.FieldAddress, field.TypeString, value)
	}
	if value, ok := mu.mutation.SiteURL(); ok {
		_spec.SetField(meshi.FieldSiteURL, field.TypeString, value)
	}
	if value, ok := mu.mutation.PublishedDate(); ok {
		_spec.SetField(meshi.FieldPublishedDate, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Latitude(); ok {
		_spec.SetField(meshi.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedLatitude(); ok {
		_spec.AddField(meshi.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.Longitude(); ok {
		_spec.SetField(meshi.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.AddedLongitude(); ok {
		_spec.AddField(meshi.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(meshi.FieldCreatedAt, field.TypeTime, value)
	}
	if mu.mutation.MunicipalityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meshi.MunicipalityTable,
			Columns: []string{meshi.MunicipalityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MunicipalityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meshi.MunicipalityTable,
			Columns: []string{meshi.MunicipalityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meshi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MeshiUpdateOne is the builder for updating a single Meshi entity.
type MeshiUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MeshiMutation
}

// SetArticleID sets the "article_id" field.
func (muo *MeshiUpdateOne) SetArticleID(s string) *MeshiUpdateOne {
	muo.mutation.SetArticleID(s)
	return muo
}

// SetTitle sets the "title" field.
func (muo *MeshiUpdateOne) SetTitle(s string) *MeshiUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableTitle(s *string) *MeshiUpdateOne {
	if s != nil {
		muo.SetTitle(*s)
	}
	return muo
}

// SetImageURL sets the "image_url" field.
func (muo *MeshiUpdateOne) SetImageURL(s string) *MeshiUpdateOne {
	muo.mutation.SetImageURL(s)
	return muo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableImageURL(s *string) *MeshiUpdateOne {
	if s != nil {
		muo.SetImageURL(*s)
	}
	return muo
}

// SetStoreName sets the "store_name" field.
func (muo *MeshiUpdateOne) SetStoreName(s string) *MeshiUpdateOne {
	muo.mutation.SetStoreName(s)
	return muo
}

// SetNillableStoreName sets the "store_name" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableStoreName(s *string) *MeshiUpdateOne {
	if s != nil {
		muo.SetStoreName(*s)
	}
	return muo
}

// SetAddress sets the "address" field.
func (muo *MeshiUpdateOne) SetAddress(s string) *MeshiUpdateOne {
	muo.mutation.SetAddress(s)
	return muo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableAddress(s *string) *MeshiUpdateOne {
	if s != nil {
		muo.SetAddress(*s)
	}
	return muo
}

// SetSiteURL sets the "site_url" field.
func (muo *MeshiUpdateOne) SetSiteURL(s string) *MeshiUpdateOne {
	muo.mutation.SetSiteURL(s)
	return muo
}

// SetNillableSiteURL sets the "site_url" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableSiteURL(s *string) *MeshiUpdateOne {
	if s != nil {
		muo.SetSiteURL(*s)
	}
	return muo
}

// SetPublishedDate sets the "published_date" field.
func (muo *MeshiUpdateOne) SetPublishedDate(t time.Time) *MeshiUpdateOne {
	muo.mutation.SetPublishedDate(t)
	return muo
}

// SetLatitude sets the "latitude" field.
func (muo *MeshiUpdateOne) SetLatitude(f float64) *MeshiUpdateOne {
	muo.mutation.ResetLatitude()
	muo.mutation.SetLatitude(f)
	return muo
}

// AddLatitude adds f to the "latitude" field.
func (muo *MeshiUpdateOne) AddLatitude(f float64) *MeshiUpdateOne {
	muo.mutation.AddLatitude(f)
	return muo
}

// SetLongitude sets the "longitude" field.
func (muo *MeshiUpdateOne) SetLongitude(f float64) *MeshiUpdateOne {
	muo.mutation.ResetLongitude()
	muo.mutation.SetLongitude(f)
	return muo
}

// AddLongitude adds f to the "longitude" field.
func (muo *MeshiUpdateOne) AddLongitude(f float64) *MeshiUpdateOne {
	muo.mutation.AddLongitude(f)
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MeshiUpdateOne) SetCreatedAt(t time.Time) *MeshiUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableCreatedAt(t *time.Time) *MeshiUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetMunicipalityID sets the "municipality" edge to the Municipality entity by ID.
func (muo *MeshiUpdateOne) SetMunicipalityID(id int) *MeshiUpdateOne {
	muo.mutation.SetMunicipalityID(id)
	return muo
}

// SetNillableMunicipalityID sets the "municipality" edge to the Municipality entity by ID if the given value is not nil.
func (muo *MeshiUpdateOne) SetNillableMunicipalityID(id *int) *MeshiUpdateOne {
	if id != nil {
		muo = muo.SetMunicipalityID(*id)
	}
	return muo
}

// SetMunicipality sets the "municipality" edge to the Municipality entity.
func (muo *MeshiUpdateOne) SetMunicipality(m *Municipality) *MeshiUpdateOne {
	return muo.SetMunicipalityID(m.ID)
}

// Mutation returns the MeshiMutation object of the builder.
func (muo *MeshiUpdateOne) Mutation() *MeshiMutation {
	return muo.mutation
}

// ClearMunicipality clears the "municipality" edge to the Municipality entity.
func (muo *MeshiUpdateOne) ClearMunicipality() *MeshiUpdateOne {
	muo.mutation.ClearMunicipality()
	return muo
}

// Where appends a list predicates to the MeshiUpdate builder.
func (muo *MeshiUpdateOne) Where(ps ...predicate.Meshi) *MeshiUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MeshiUpdateOne) Select(field string, fields ...string) *MeshiUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Meshi entity.
func (muo *MeshiUpdateOne) Save(ctx context.Context) (*Meshi, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MeshiUpdateOne) SaveX(ctx context.Context) *Meshi {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MeshiUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MeshiUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MeshiUpdateOne) sqlSave(ctx context.Context) (_node *Meshi, err error) {
	_spec := sqlgraph.NewUpdateSpec(meshi.Table, meshi.Columns, sqlgraph.NewFieldSpec(meshi.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Meshi.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, meshi.FieldID)
		for _, f := range fields {
			if !meshi.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != meshi.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.ArticleID(); ok {
		_spec.SetField(meshi.FieldArticleID, field.TypeString, value)
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.SetField(meshi.FieldTitle, field.TypeString, value)
	}
	if value, ok := muo.mutation.ImageURL(); ok {
		_spec.SetField(meshi.FieldImageURL, field.TypeString, value)
	}
	if value, ok := muo.mutation.StoreName(); ok {
		_spec.SetField(meshi.FieldStoreName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Address(); ok {
		_spec.SetField(meshi.FieldAddress, field.TypeString, value)
	}
	if value, ok := muo.mutation.SiteURL(); ok {
		_spec.SetField(meshi.FieldSiteURL, field.TypeString, value)
	}
	if value, ok := muo.mutation.PublishedDate(); ok {
		_spec.SetField(meshi.FieldPublishedDate, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Latitude(); ok {
		_spec.SetField(meshi.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedLatitude(); ok {
		_spec.AddField(meshi.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.Longitude(); ok {
		_spec.SetField(meshi.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.AddedLongitude(); ok {
		_spec.AddField(meshi.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(meshi.FieldCreatedAt, field.TypeTime, value)
	}
	if muo.mutation.MunicipalityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meshi.MunicipalityTable,
			Columns: []string{meshi.MunicipalityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MunicipalityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   meshi.MunicipalityTable,
			Columns: []string{meshi.MunicipalityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipality.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Meshi{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{meshi.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}