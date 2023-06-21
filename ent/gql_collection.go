// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/shimabukuromeg/ageage-search/ent/meshi"
	"github.com/shimabukuromeg/ageage-search/ent/municipality"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MeshiQuery) CollectFields(ctx context.Context, satisfies ...string) (*MeshiQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MeshiQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(meshi.Columns))
		selectedFields = []string{meshi.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "municipality":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MunicipalityClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, municipalityImplementors)...); err != nil {
				return err
			}
			m.withMunicipality = query
		case "articleID":
			if _, ok := fieldSeen[meshi.FieldArticleID]; !ok {
				selectedFields = append(selectedFields, meshi.FieldArticleID)
				fieldSeen[meshi.FieldArticleID] = struct{}{}
			}
		case "title":
			if _, ok := fieldSeen[meshi.FieldTitle]; !ok {
				selectedFields = append(selectedFields, meshi.FieldTitle)
				fieldSeen[meshi.FieldTitle] = struct{}{}
			}
		case "imageURL":
			if _, ok := fieldSeen[meshi.FieldImageURL]; !ok {
				selectedFields = append(selectedFields, meshi.FieldImageURL)
				fieldSeen[meshi.FieldImageURL] = struct{}{}
			}
		case "storeName":
			if _, ok := fieldSeen[meshi.FieldStoreName]; !ok {
				selectedFields = append(selectedFields, meshi.FieldStoreName)
				fieldSeen[meshi.FieldStoreName] = struct{}{}
			}
		case "address":
			if _, ok := fieldSeen[meshi.FieldAddress]; !ok {
				selectedFields = append(selectedFields, meshi.FieldAddress)
				fieldSeen[meshi.FieldAddress] = struct{}{}
			}
		case "siteURL":
			if _, ok := fieldSeen[meshi.FieldSiteURL]; !ok {
				selectedFields = append(selectedFields, meshi.FieldSiteURL)
				fieldSeen[meshi.FieldSiteURL] = struct{}{}
			}
		case "publishedDate":
			if _, ok := fieldSeen[meshi.FieldPublishedDate]; !ok {
				selectedFields = append(selectedFields, meshi.FieldPublishedDate)
				fieldSeen[meshi.FieldPublishedDate] = struct{}{}
			}
		case "latitude":
			if _, ok := fieldSeen[meshi.FieldLatitude]; !ok {
				selectedFields = append(selectedFields, meshi.FieldLatitude)
				fieldSeen[meshi.FieldLatitude] = struct{}{}
			}
		case "longitude":
			if _, ok := fieldSeen[meshi.FieldLongitude]; !ok {
				selectedFields = append(selectedFields, meshi.FieldLongitude)
				fieldSeen[meshi.FieldLongitude] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[meshi.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, meshi.FieldCreatedAt)
				fieldSeen[meshi.FieldCreatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		m.Select(selectedFields...)
	}
	return nil
}

type meshiPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MeshiPaginateOption
}

func newMeshiPaginateArgs(rv map[string]any) *meshiPaginateArgs {
	args := &meshiPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MeshiOrder{Field: &MeshiOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMeshiOrder(order))
			}
		case *MeshiOrder:
			if v != nil {
				args.opts = append(args.opts, WithMeshiOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MeshiWhereInput); ok {
		args.opts = append(args.opts, WithMeshiFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MunicipalityQuery) CollectFields(ctx context.Context, satisfies ...string) (*MunicipalityQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MunicipalityQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(municipality.Columns))
		selectedFields = []string{municipality.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "meshis":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MeshiClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, meshiImplementors)...); err != nil {
				return err
			}
			m.WithNamedMeshis(alias, func(wq *MeshiQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[municipality.FieldName]; !ok {
				selectedFields = append(selectedFields, municipality.FieldName)
				fieldSeen[municipality.FieldName] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[municipality.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, municipality.FieldCreatedAt)
				fieldSeen[municipality.FieldCreatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		m.Select(selectedFields...)
	}
	return nil
}

type municipalityPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MunicipalityPaginateOption
}

func newMunicipalityPaginateArgs(rv map[string]any) *municipalityPaginateArgs {
	args := &municipalityPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MunicipalityOrder{Field: &MunicipalityOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMunicipalityOrder(order))
			}
		case *MunicipalityOrder:
			if v != nil {
				args.opts = append(args.opts, WithMunicipalityOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MunicipalityWhereInput); ok {
		args.opts = append(args.opts, WithMunicipalityFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
