// Code generated by entc, DO NOT EDIT.

package symptom

import (
	"github.com/darksford123x/repairinvoice-record/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Syname applies equality check predicate on the "Syname" field. It's identical to SynameEQ.
func Syname(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyname), v))
	})
}

// SynameEQ applies the EQ predicate on the "Syname" field.
func SynameEQ(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyname), v))
	})
}

// SynameNEQ applies the NEQ predicate on the "Syname" field.
func SynameNEQ(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyname), v))
	})
}

// SynameIn applies the In predicate on the "Syname" field.
func SynameIn(vs ...string) predicate.Symptom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSyname), v...))
	})
}

// SynameNotIn applies the NotIn predicate on the "Syname" field.
func SynameNotIn(vs ...string) predicate.Symptom {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptom(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSyname), v...))
	})
}

// SynameGT applies the GT predicate on the "Syname" field.
func SynameGT(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSyname), v))
	})
}

// SynameGTE applies the GTE predicate on the "Syname" field.
func SynameGTE(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSyname), v))
	})
}

// SynameLT applies the LT predicate on the "Syname" field.
func SynameLT(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSyname), v))
	})
}

// SynameLTE applies the LTE predicate on the "Syname" field.
func SynameLTE(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSyname), v))
	})
}

// SynameContains applies the Contains predicate on the "Syname" field.
func SynameContains(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSyname), v))
	})
}

// SynameHasPrefix applies the HasPrefix predicate on the "Syname" field.
func SynameHasPrefix(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSyname), v))
	})
}

// SynameHasSuffix applies the HasSuffix predicate on the "Syname" field.
func SynameHasSuffix(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSyname), v))
	})
}

// SynameEqualFold applies the EqualFold predicate on the "Syname" field.
func SynameEqualFold(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSyname), v))
	})
}

// SynameContainsFold applies the ContainsFold predicate on the "Syname" field.
func SynameContainsFold(v string) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSyname), v))
	})
}

// HasRepairInformation applies the HasEdge predicate on the "repair_information" edge.
func HasRepairInformation() predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairInformationTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairInformationTable, RepairInformationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepairInformationWith applies the HasEdge predicate on the "repair_information" edge with a given conditions (other predicates).
func HasRepairInformationWith(preds ...predicate.RepairInvoice) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepairInformationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepairInformationTable, RepairInformationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Symptom) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Symptom) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
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
func Not(p predicate.Symptom) predicate.Symptom {
	return predicate.Symptom(func(s *sql.Selector) {
		p(s.Not())
	})
}
