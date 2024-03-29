package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Text("name").
			Unique().Annotations(
			entgql.OrderField("NAME"),
		),
		field.Text("zipcode").
			Unique().Annotations(
			entgql.OrderField("ZIPCODE"),
		).Optional(),
		field.Time("created_at").
			Default(time.Now).Annotations(
			entgql.OrderField("CREATED_AT"),
		),
	}
}

// Edges of the Municipality.
func (Municipality) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("meshis", Meshi.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Municipality) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
