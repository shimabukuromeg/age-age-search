// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MeshisColumns holds the columns for the "meshis" table.
	MeshisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "article_id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString, Default: "unknown"},
		{Name: "image_url", Type: field.TypeString, Default: "unknown"},
		{Name: "store_name", Type: field.TypeString, Default: "unknown"},
		{Name: "address", Type: field.TypeString, Default: "unknown"},
		{Name: "site_url", Type: field.TypeString, Default: "unknown"},
		{Name: "municipality_meshis", Type: field.TypeInt, Nullable: true},
	}
	// MeshisTable holds the schema information for the "meshis" table.
	MeshisTable = &schema.Table{
		Name:       "meshis",
		Columns:    MeshisColumns,
		PrimaryKey: []*schema.Column{MeshisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "meshis_municipalities_meshis",
				Columns:    []*schema.Column{MeshisColumns[7]},
				RefColumns: []*schema.Column{MunicipalitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MunicipalitiesColumns holds the columns for the "municipalities" table.
	MunicipalitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// MunicipalitiesTable holds the schema information for the "municipalities" table.
	MunicipalitiesTable = &schema.Table{
		Name:       "municipalities",
		Columns:    MunicipalitiesColumns,
		PrimaryKey: []*schema.Column{MunicipalitiesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MeshisTable,
		MunicipalitiesTable,
	}
)

func init() {
	MeshisTable.ForeignKeys[0].RefTable = MunicipalitiesTable
}
