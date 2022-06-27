// Code generated by entc, DO NOT EDIT.

package group

import (
	"SafeSend/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GroupName applies equality check predicate on the "groupName" field. It's identical to GroupNameEQ.
func GroupName(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// MaxParticipants applies equality check predicate on the "maxParticipants" field. It's identical to MaxParticipantsEQ.
func MaxParticipants(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxParticipants), v))
	})
}

// InviteOnly applies equality check predicate on the "inviteOnly" field. It's identical to InviteOnlyEQ.
func InviteOnly(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteOnly), v))
	})
}

// DateCreated applies equality check predicate on the "dateCreated" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCreated), v))
	})
}

// DateModified applies equality check predicate on the "dateModified" field. It's identical to DateModifiedEQ.
func DateModified(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateModified), v))
	})
}

// DateDeleted applies equality check predicate on the "dateDeleted" field. It's identical to DateDeletedEQ.
func DateDeleted(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateDeleted), v))
	})
}

// GroupNameEQ applies the EQ predicate on the "groupName" field.
func GroupNameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// GroupNameNEQ applies the NEQ predicate on the "groupName" field.
func GroupNameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupName), v))
	})
}

// GroupNameIn applies the In predicate on the "groupName" field.
func GroupNameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupName), v...))
	})
}

// GroupNameNotIn applies the NotIn predicate on the "groupName" field.
func GroupNameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupName), v...))
	})
}

// GroupNameGT applies the GT predicate on the "groupName" field.
func GroupNameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupName), v))
	})
}

// GroupNameGTE applies the GTE predicate on the "groupName" field.
func GroupNameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupName), v))
	})
}

// GroupNameLT applies the LT predicate on the "groupName" field.
func GroupNameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupName), v))
	})
}

// GroupNameLTE applies the LTE predicate on the "groupName" field.
func GroupNameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupName), v))
	})
}

// GroupNameContains applies the Contains predicate on the "groupName" field.
func GroupNameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupName), v))
	})
}

// GroupNameHasPrefix applies the HasPrefix predicate on the "groupName" field.
func GroupNameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupName), v))
	})
}

// GroupNameHasSuffix applies the HasSuffix predicate on the "groupName" field.
func GroupNameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupName), v))
	})
}

// GroupNameEqualFold applies the EqualFold predicate on the "groupName" field.
func GroupNameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupName), v))
	})
}

// GroupNameContainsFold applies the ContainsFold predicate on the "groupName" field.
func GroupNameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupName), v))
	})
}

// MaxParticipantsEQ applies the EQ predicate on the "maxParticipants" field.
func MaxParticipantsEQ(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxParticipants), v))
	})
}

// MaxParticipantsNEQ applies the NEQ predicate on the "maxParticipants" field.
func MaxParticipantsNEQ(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxParticipants), v))
	})
}

// MaxParticipantsIn applies the In predicate on the "maxParticipants" field.
func MaxParticipantsIn(vs ...int32) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxParticipants), v...))
	})
}

// MaxParticipantsNotIn applies the NotIn predicate on the "maxParticipants" field.
func MaxParticipantsNotIn(vs ...int32) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxParticipants), v...))
	})
}

// MaxParticipantsGT applies the GT predicate on the "maxParticipants" field.
func MaxParticipantsGT(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxParticipants), v))
	})
}

// MaxParticipantsGTE applies the GTE predicate on the "maxParticipants" field.
func MaxParticipantsGTE(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxParticipants), v))
	})
}

// MaxParticipantsLT applies the LT predicate on the "maxParticipants" field.
func MaxParticipantsLT(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxParticipants), v))
	})
}

// MaxParticipantsLTE applies the LTE predicate on the "maxParticipants" field.
func MaxParticipantsLTE(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxParticipants), v))
	})
}

// InviteOnlyEQ applies the EQ predicate on the "inviteOnly" field.
func InviteOnlyEQ(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteOnly), v))
	})
}

// InviteOnlyNEQ applies the NEQ predicate on the "inviteOnly" field.
func InviteOnlyNEQ(v bool) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviteOnly), v))
	})
}

// DateCreatedEQ applies the EQ predicate on the "dateCreated" field.
func DateCreatedEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCreated), v))
	})
}

// DateCreatedNEQ applies the NEQ predicate on the "dateCreated" field.
func DateCreatedNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateCreated), v))
	})
}

// DateCreatedIn applies the In predicate on the "dateCreated" field.
func DateCreatedIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateCreated), v...))
	})
}

// DateCreatedNotIn applies the NotIn predicate on the "dateCreated" field.
func DateCreatedNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateCreated), v...))
	})
}

// DateCreatedGT applies the GT predicate on the "dateCreated" field.
func DateCreatedGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateCreated), v))
	})
}

// DateCreatedGTE applies the GTE predicate on the "dateCreated" field.
func DateCreatedGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateCreated), v))
	})
}

// DateCreatedLT applies the LT predicate on the "dateCreated" field.
func DateCreatedLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateCreated), v))
	})
}

// DateCreatedLTE applies the LTE predicate on the "dateCreated" field.
func DateCreatedLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateCreated), v))
	})
}

// DateModifiedEQ applies the EQ predicate on the "dateModified" field.
func DateModifiedEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateModified), v))
	})
}

// DateModifiedNEQ applies the NEQ predicate on the "dateModified" field.
func DateModifiedNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateModified), v))
	})
}

// DateModifiedIn applies the In predicate on the "dateModified" field.
func DateModifiedIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateModified), v...))
	})
}

// DateModifiedNotIn applies the NotIn predicate on the "dateModified" field.
func DateModifiedNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateModified), v...))
	})
}

// DateModifiedGT applies the GT predicate on the "dateModified" field.
func DateModifiedGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateModified), v))
	})
}

// DateModifiedGTE applies the GTE predicate on the "dateModified" field.
func DateModifiedGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateModified), v))
	})
}

// DateModifiedLT applies the LT predicate on the "dateModified" field.
func DateModifiedLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateModified), v))
	})
}

// DateModifiedLTE applies the LTE predicate on the "dateModified" field.
func DateModifiedLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateModified), v))
	})
}

// DateDeletedEQ applies the EQ predicate on the "dateDeleted" field.
func DateDeletedEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedNEQ applies the NEQ predicate on the "dateDeleted" field.
func DateDeletedNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedIn applies the In predicate on the "dateDeleted" field.
func DateDeletedIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateDeleted), v...))
	})
}

// DateDeletedNotIn applies the NotIn predicate on the "dateDeleted" field.
func DateDeletedNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateDeleted), v...))
	})
}

// DateDeletedGT applies the GT predicate on the "dateDeleted" field.
func DateDeletedGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedGTE applies the GTE predicate on the "dateDeleted" field.
func DateDeletedGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedLT applies the LT predicate on the "dateDeleted" field.
func DateDeletedLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedLTE applies the LTE predicate on the "dateDeleted" field.
func DateDeletedLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateDeleted), v))
	})
}

// DateDeletedIsNil applies the IsNil predicate on the "dateDeleted" field.
func DateDeletedIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDateDeleted)))
	})
}

// DateDeletedNotNil applies the NotNil predicate on the "dateDeleted" field.
func DateDeletedNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDateDeleted)))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEntities applies the HasEdge predicate on the "entities" edge.
func HasEntities() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EntitiesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EntitiesTable, EntitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEntitiesWith applies the HasEdge predicate on the "entities" edge with a given conditions (other predicates).
func HasEntitiesWith(preds ...predicate.Entity) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EntitiesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EntitiesTable, EntitiesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
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
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}