// Code generated by ent, DO NOT EDIT.

package note

import (
	"example/golang-rest-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Note {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Note(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Note {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Note(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTag), v))
	})
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTag), v))
	})
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.Note {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Note(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTag), v...))
	})
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.Note {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Note(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTag), v...))
	})
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTag), v))
	})
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTag), v))
	})
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTag), v))
	})
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTag), v))
	})
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTag), v))
	})
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTag), v))
	})
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTag), v))
	})
}

// TagIsNil applies the IsNil predicate on the "tag" field.
func TagIsNil() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTag)))
	})
}

// TagNotNil applies the NotNil predicate on the "tag" field.
func TagNotNil() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTag)))
	})
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTag), v))
	})
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTag), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
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
func Not(p predicate.Note) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		p(s.Not())
	})
}
