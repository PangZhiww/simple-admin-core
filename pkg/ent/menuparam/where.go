// Code generated by ent, DO NOT EDIT.

package menuparam

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldUpdatedAt, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldType, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldValue, v))
}

// MenuID applies equality check predicate on the "menu_id" field. It's identical to MenuIDEQ.
func MenuID(v uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldMenuID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContainsFold(FieldType, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldContainsFold(FieldValue, v))
}

// MenuIDEQ applies the EQ predicate on the "menu_id" field.
func MenuIDEQ(v uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldEQ(FieldMenuID, v))
}

// MenuIDNEQ applies the NEQ predicate on the "menu_id" field.
func MenuIDNEQ(v uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNEQ(FieldMenuID, v))
}

// MenuIDIn applies the In predicate on the "menu_id" field.
func MenuIDIn(vs ...uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIn(FieldMenuID, vs...))
}

// MenuIDNotIn applies the NotIn predicate on the "menu_id" field.
func MenuIDNotIn(vs ...uint64) predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotIn(FieldMenuID, vs...))
}

// MenuIDIsNil applies the IsNil predicate on the "menu_id" field.
func MenuIDIsNil() predicate.MenuParam {
	return predicate.MenuParam(sql.FieldIsNull(FieldMenuID))
}

// MenuIDNotNil applies the NotNil predicate on the "menu_id" field.
func MenuIDNotNil() predicate.MenuParam {
	return predicate.MenuParam(sql.FieldNotNull(FieldMenuID))
}

// HasMenus applies the HasEdge predicate on the "menus" edge.
func HasMenus() predicate.MenuParam {
	return predicate.MenuParam(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenusTable, MenusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenusWith applies the HasEdge predicate on the "menus" edge with a given conditions (other predicates).
func HasMenusWith(preds ...predicate.Menu) predicate.MenuParam {
	return predicate.MenuParam(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MenusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MenusTable, MenusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MenuParam) predicate.MenuParam {
	return predicate.MenuParam(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MenuParam) predicate.MenuParam {
	return predicate.MenuParam(func(s *sql.Selector) {
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
func Not(p predicate.MenuParam) predicate.MenuParam {
	return predicate.MenuParam(func(s *sql.Selector) {
		p(s.Not())
	})
}
