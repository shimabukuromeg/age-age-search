// Code generated by ent, DO NOT EDIT.

package municipality

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shimabukuromeg/ageage-search/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Municipality {
	return predicate.Municipality(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Municipality {
	return predicate.Municipality(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Municipality {
	return predicate.Municipality(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldName, v))
}

// Zipcode applies equality check predicate on the "zipcode" field. It's identical to ZipcodeEQ.
func Zipcode(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldZipcode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Municipality {
	return predicate.Municipality(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Municipality {
	return predicate.Municipality(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldContainsFold(FieldName, v))
}

// ZipcodeEQ applies the EQ predicate on the "zipcode" field.
func ZipcodeEQ(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldZipcode, v))
}

// ZipcodeNEQ applies the NEQ predicate on the "zipcode" field.
func ZipcodeNEQ(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldNEQ(FieldZipcode, v))
}

// ZipcodeIn applies the In predicate on the "zipcode" field.
func ZipcodeIn(vs ...string) predicate.Municipality {
	return predicate.Municipality(sql.FieldIn(FieldZipcode, vs...))
}

// ZipcodeNotIn applies the NotIn predicate on the "zipcode" field.
func ZipcodeNotIn(vs ...string) predicate.Municipality {
	return predicate.Municipality(sql.FieldNotIn(FieldZipcode, vs...))
}

// ZipcodeGT applies the GT predicate on the "zipcode" field.
func ZipcodeGT(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldGT(FieldZipcode, v))
}

// ZipcodeGTE applies the GTE predicate on the "zipcode" field.
func ZipcodeGTE(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldGTE(FieldZipcode, v))
}

// ZipcodeLT applies the LT predicate on the "zipcode" field.
func ZipcodeLT(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldLT(FieldZipcode, v))
}

// ZipcodeLTE applies the LTE predicate on the "zipcode" field.
func ZipcodeLTE(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldLTE(FieldZipcode, v))
}

// ZipcodeContains applies the Contains predicate on the "zipcode" field.
func ZipcodeContains(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldContains(FieldZipcode, v))
}

// ZipcodeHasPrefix applies the HasPrefix predicate on the "zipcode" field.
func ZipcodeHasPrefix(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldHasPrefix(FieldZipcode, v))
}

// ZipcodeHasSuffix applies the HasSuffix predicate on the "zipcode" field.
func ZipcodeHasSuffix(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldHasSuffix(FieldZipcode, v))
}

// ZipcodeIsNil applies the IsNil predicate on the "zipcode" field.
func ZipcodeIsNil() predicate.Municipality {
	return predicate.Municipality(sql.FieldIsNull(FieldZipcode))
}

// ZipcodeNotNil applies the NotNil predicate on the "zipcode" field.
func ZipcodeNotNil() predicate.Municipality {
	return predicate.Municipality(sql.FieldNotNull(FieldZipcode))
}

// ZipcodeEqualFold applies the EqualFold predicate on the "zipcode" field.
func ZipcodeEqualFold(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldEqualFold(FieldZipcode, v))
}

// ZipcodeContainsFold applies the ContainsFold predicate on the "zipcode" field.
func ZipcodeContainsFold(v string) predicate.Municipality {
	return predicate.Municipality(sql.FieldContainsFold(FieldZipcode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Municipality {
	return predicate.Municipality(sql.FieldLTE(FieldCreatedAt, v))
}

// HasMeshis applies the HasEdge predicate on the "meshis" edge.
func HasMeshis() predicate.Municipality {
	return predicate.Municipality(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MeshisTable, MeshisColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMeshisWith applies the HasEdge predicate on the "meshis" edge with a given conditions (other predicates).
func HasMeshisWith(preds ...predicate.Meshi) predicate.Municipality {
	return predicate.Municipality(func(s *sql.Selector) {
		step := newMeshisStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Municipality) predicate.Municipality {
	return predicate.Municipality(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Municipality) predicate.Municipality {
	return predicate.Municipality(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Municipality) predicate.Municipality {
	return predicate.Municipality(func(s *sql.Selector) {
		p(s.Not())
	})
}
