// Code generated by ent, DO NOT EDIT.

package usermetadata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Skijetler/alphinium/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// LastOnline applies equality check predicate on the "last_online" field. It's identical to LastOnlineEQ.
func LastOnline(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastOnline), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldColor), v))
	})
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldColor), v...))
	})
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldColor), v...))
	})
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldColor), v))
	})
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldColor), v))
	})
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldColor), v))
	})
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldColor), v))
	})
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldColor), v))
	})
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldColor), v))
	})
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldColor), v))
	})
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldColor), v))
	})
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldColor), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGender), v))
	})
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGender), v...))
	})
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGender), v...))
	})
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGender), v))
	})
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGender), v))
	})
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGender), v))
	})
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGender), v))
	})
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGender), v))
	})
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGender), v))
	})
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGender), v))
	})
}

// GenderEqualFold applies the EqualFold predicate on the "gender" field.
func GenderEqualFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGender), v))
	})
}

// GenderContainsFold applies the ContainsFold predicate on the "gender" field.
func GenderContainsFold(v string) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGender), v))
	})
}

// LastOnlineEQ applies the EQ predicate on the "last_online" field.
func LastOnlineEQ(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastOnline), v))
	})
}

// LastOnlineNEQ applies the NEQ predicate on the "last_online" field.
func LastOnlineNEQ(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastOnline), v))
	})
}

// LastOnlineIn applies the In predicate on the "last_online" field.
func LastOnlineIn(vs ...time.Time) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastOnline), v...))
	})
}

// LastOnlineNotIn applies the NotIn predicate on the "last_online" field.
func LastOnlineNotIn(vs ...time.Time) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastOnline), v...))
	})
}

// LastOnlineGT applies the GT predicate on the "last_online" field.
func LastOnlineGT(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastOnline), v))
	})
}

// LastOnlineGTE applies the GTE predicate on the "last_online" field.
func LastOnlineGTE(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastOnline), v))
	})
}

// LastOnlineLT applies the LT predicate on the "last_online" field.
func LastOnlineLT(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastOnline), v))
	})
}

// LastOnlineLTE applies the LTE predicate on the "last_online" field.
func LastOnlineLTE(v time.Time) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastOnline), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.UserMetadata {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserMetadata(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserMetadata) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserMetadata) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
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
func Not(p predicate.UserMetadata) predicate.UserMetadata {
	return predicate.UserMetadata(func(s *sql.Selector) {
		p(s.Not())
	})
}
