// Code generated by ent, DO NOT EDIT.

package datalist

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Xwudao/neter-template/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldUpdateTime, v))
}

// Kind applies equality check predicate on the "kind" field. It's identical to KindEQ.
func Kind(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldKind, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldValue, v))
}

// ItemOrder applies equality check predicate on the "item_order" field. It's identical to ItemOrderEQ.
func ItemOrder(v int) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldItemOrder, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldUpdateTime, v))
}

// LabelEQ applies the EQ predicate on the "label" field.
func LabelEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldLabel, v))
}

// LabelNEQ applies the NEQ predicate on the "label" field.
func LabelNEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldLabel, v))
}

// LabelIn applies the In predicate on the "label" field.
func LabelIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldLabel, vs...))
}

// LabelNotIn applies the NotIn predicate on the "label" field.
func LabelNotIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldLabel, vs...))
}

// LabelGT applies the GT predicate on the "label" field.
func LabelGT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldLabel, v))
}

// LabelGTE applies the GTE predicate on the "label" field.
func LabelGTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldLabel, v))
}

// LabelLT applies the LT predicate on the "label" field.
func LabelLT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldLabel, v))
}

// LabelLTE applies the LTE predicate on the "label" field.
func LabelLTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldLabel, v))
}

// LabelContains applies the Contains predicate on the "label" field.
func LabelContains(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContains(FieldLabel, v))
}

// LabelHasPrefix applies the HasPrefix predicate on the "label" field.
func LabelHasPrefix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasPrefix(FieldLabel, v))
}

// LabelHasSuffix applies the HasSuffix predicate on the "label" field.
func LabelHasSuffix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasSuffix(FieldLabel, v))
}

// LabelEqualFold applies the EqualFold predicate on the "label" field.
func LabelEqualFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEqualFold(FieldLabel, v))
}

// LabelContainsFold applies the ContainsFold predicate on the "label" field.
func LabelContainsFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContainsFold(FieldLabel, v))
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldKind, v))
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldKind, v))
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldKind, vs...))
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldKind, vs...))
}

// KindGT applies the GT predicate on the "kind" field.
func KindGT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldKind, v))
}

// KindGTE applies the GTE predicate on the "kind" field.
func KindGTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldKind, v))
}

// KindLT applies the LT predicate on the "kind" field.
func KindLT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldKind, v))
}

// KindLTE applies the LTE predicate on the "kind" field.
func KindLTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldKind, v))
}

// KindContains applies the Contains predicate on the "kind" field.
func KindContains(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContains(FieldKind, v))
}

// KindHasPrefix applies the HasPrefix predicate on the "kind" field.
func KindHasPrefix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasPrefix(FieldKind, v))
}

// KindHasSuffix applies the HasSuffix predicate on the "kind" field.
func KindHasSuffix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasSuffix(FieldKind, v))
}

// KindEqualFold applies the EqualFold predicate on the "kind" field.
func KindEqualFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEqualFold(FieldKind, v))
}

// KindContainsFold applies the ContainsFold predicate on the "kind" field.
func KindContainsFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContainsFold(FieldKind, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.DataList {
	return predicate.DataList(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.DataList {
	return predicate.DataList(sql.FieldContainsFold(FieldValue, v))
}

// ItemOrderEQ applies the EQ predicate on the "item_order" field.
func ItemOrderEQ(v int) predicate.DataList {
	return predicate.DataList(sql.FieldEQ(FieldItemOrder, v))
}

// ItemOrderNEQ applies the NEQ predicate on the "item_order" field.
func ItemOrderNEQ(v int) predicate.DataList {
	return predicate.DataList(sql.FieldNEQ(FieldItemOrder, v))
}

// ItemOrderIn applies the In predicate on the "item_order" field.
func ItemOrderIn(vs ...int) predicate.DataList {
	return predicate.DataList(sql.FieldIn(FieldItemOrder, vs...))
}

// ItemOrderNotIn applies the NotIn predicate on the "item_order" field.
func ItemOrderNotIn(vs ...int) predicate.DataList {
	return predicate.DataList(sql.FieldNotIn(FieldItemOrder, vs...))
}

// ItemOrderGT applies the GT predicate on the "item_order" field.
func ItemOrderGT(v int) predicate.DataList {
	return predicate.DataList(sql.FieldGT(FieldItemOrder, v))
}

// ItemOrderGTE applies the GTE predicate on the "item_order" field.
func ItemOrderGTE(v int) predicate.DataList {
	return predicate.DataList(sql.FieldGTE(FieldItemOrder, v))
}

// ItemOrderLT applies the LT predicate on the "item_order" field.
func ItemOrderLT(v int) predicate.DataList {
	return predicate.DataList(sql.FieldLT(FieldItemOrder, v))
}

// ItemOrderLTE applies the LTE predicate on the "item_order" field.
func ItemOrderLTE(v int) predicate.DataList {
	return predicate.DataList(sql.FieldLTE(FieldItemOrder, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DataList) predicate.DataList {
	return predicate.DataList(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DataList) predicate.DataList {
	return predicate.DataList(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DataList) predicate.DataList {
	return predicate.DataList(sql.NotPredicates(p))
}