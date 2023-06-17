package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Municipality holds the schema definition for the Municipality entity.
type Municipality struct {
	ent.Schema
}

// Fields of the Municipality.
func (Municipality) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Municipality.
func (Municipality) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("meshis", Meshi.Type),
	}
}
