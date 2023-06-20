// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
)

// Municipality is the model entity for the Municipality schema.
type Municipality struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MunicipalityQuery when eager-loading is set.
	Edges        MunicipalityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MunicipalityEdges holds the relations/edges for other nodes in the graph.
type MunicipalityEdges struct {
	// Meshis holds the value of the meshis edge.
	Meshis []*Meshi `json:"meshis,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedMeshis map[string][]*Meshi
}

// MeshisOrErr returns the Meshis value or an error if the edge
// was not loaded in eager-loading.
func (e MunicipalityEdges) MeshisOrErr() ([]*Meshi, error) {
	if e.loadedTypes[0] {
		return e.Meshis, nil
	}
	return nil, &NotLoadedError{edge: "meshis"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Municipality) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case municipality.FieldID:
			values[i] = new(sql.NullInt64)
		case municipality.FieldName:
			values[i] = new(sql.NullString)
		case municipality.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Municipality fields.
func (m *Municipality) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case municipality.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case municipality.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case municipality.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Municipality.
// This includes values selected through modifiers, order, etc.
func (m *Municipality) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMeshis queries the "meshis" edge of the Municipality entity.
func (m *Municipality) QueryMeshis() *MeshiQuery {
	return NewMunicipalityClient(m.config).QueryMeshis(m)
}

// Update returns a builder for updating this Municipality.
// Note that you need to call Municipality.Unwrap() before calling this method if this Municipality
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Municipality) Update() *MunicipalityUpdateOne {
	return NewMunicipalityClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Municipality entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Municipality) Unwrap() *Municipality {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Municipality is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Municipality) String() string {
	var builder strings.Builder
	builder.WriteString("Municipality(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedMeshis returns the Meshis named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Municipality) NamedMeshis(name string) ([]*Meshi, error) {
	if m.Edges.namedMeshis == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedMeshis[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Municipality) appendNamedMeshis(name string, edges ...*Meshi) {
	if m.Edges.namedMeshis == nil {
		m.Edges.namedMeshis = make(map[string][]*Meshi)
	}
	if len(edges) == 0 {
		m.Edges.namedMeshis[name] = []*Meshi{}
	} else {
		m.Edges.namedMeshis[name] = append(m.Edges.namedMeshis[name], edges...)
	}
}

// Municipalities is a parsable slice of Municipality.
type Municipalities []*Municipality
