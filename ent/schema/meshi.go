package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Meshi holds the schema definition for the Meshi entity.
type Meshi struct {
	ent.Schema
}

// Fields of the Meshi.
func (Meshi) Fields() []ent.Field {
	return []ent.Field{
		field.String("article_id").
			Unique(),
		field.String("title").
			Default("unknown"),
		field.String("image_url").
			Default("unknown"),
		field.String("store_name").
			Default("unknown"),
		field.String("address").
			Default("unknown"),
		field.String("site_url").
			Default("unknown"),
		field.Time("published_date").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Meshi.
func (Meshi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("municipality", Municipality.Type).
			Ref("meshis").
			Unique(),
	}
}
