package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			Unique().
			Annotations(
				entgql.OrderField("ARTICLE_ID"),
			),
		field.String("title").
			Default("unknown").Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("image_url").
			Default("unknown").Annotations(
			entgql.OrderField("IMAGE_URL"),
		),
		field.String("store_name").
			Default("unknown").
			Annotations(
				entgql.OrderField("STORE_NAME"),
			),
		field.String("address").
			Default("unknown").Annotations(
			entgql.OrderField("ADDRESS"),
		),
		field.String("site_url").
			Default("unknown").Annotations(
			entgql.OrderField("SITE_URL"),
		),
		field.Time("published_date").Annotations(
			entgql.OrderField("PUBLISHED_DATE"),
		),
		field.Float("latitude").Annotations(
			entgql.OrderField("LATITUDE"),
		),
		field.Float("longitude").Annotations(
			entgql.OrderField("LONGITUDE"),
		),
		field.Time("created_at").
			Default(time.Now).Annotations(
			entgql.OrderField("CREATED_AT"),
		),
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

func (Meshi) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
