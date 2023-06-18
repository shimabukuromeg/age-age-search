// Code generated by ent, DO NOT EDIT.

package meshi

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the meshi type in the database.
	Label = "meshi"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldArticleID holds the string denoting the article_id field in the database.
	FieldArticleID = "article_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldStoreName holds the string denoting the store_name field in the database.
	FieldStoreName = "store_name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldPublishedDate holds the string denoting the published_date field in the database.
	FieldPublishedDate = "published_date"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeMunicipality holds the string denoting the municipality edge name in mutations.
	EdgeMunicipality = "municipality"
	// Table holds the table name of the meshi in the database.
	Table = "meshis"
	// MunicipalityTable is the table that holds the municipality relation/edge.
	MunicipalityTable = "meshis"
	// MunicipalityInverseTable is the table name for the Municipality entity.
	// It exists in this package in order to avoid circular dependency with the "municipality" package.
	MunicipalityInverseTable = "municipalities"
	// MunicipalityColumn is the table column denoting the municipality relation/edge.
	MunicipalityColumn = "municipality_meshis"
)

// Columns holds all SQL columns for meshi fields.
var Columns = []string{
	FieldID,
	FieldArticleID,
	FieldTitle,
	FieldImageURL,
	FieldStoreName,
	FieldAddress,
	FieldSiteURL,
	FieldPublishedDate,
	FieldLatitude,
	FieldLongitude,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "meshis"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"municipality_meshis",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultImageURL holds the default value on creation for the "image_url" field.
	DefaultImageURL string
	// DefaultStoreName holds the default value on creation for the "store_name" field.
	DefaultStoreName string
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress string
	// DefaultSiteURL holds the default value on creation for the "site_url" field.
	DefaultSiteURL string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Meshi queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByArticleID orders the results by the article_id field.
func ByArticleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArticleID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByStoreName orders the results by the store_name field.
func ByStoreName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreName, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// BySiteURL orders the results by the site_url field.
func BySiteURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSiteURL, opts...).ToFunc()
}

// ByPublishedDate orders the results by the published_date field.
func ByPublishedDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublishedDate, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByMunicipalityField orders the results by municipality field.
func ByMunicipalityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMunicipalityStep(), sql.OrderByField(field, opts...))
	}
}
func newMunicipalityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MunicipalityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MunicipalityTable, MunicipalityColumn),
	)
}