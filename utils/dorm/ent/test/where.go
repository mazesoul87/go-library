// Code generated by ent, DO NOT EDIT.

package test

import (
	"entgo.io/ent/dialect/sql"
	"github.com/dtapps/go-library/utils/dorm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Test {
	return predicate.Test(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Test {
	return predicate.Test(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Test {
	return predicate.Test(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Test {
	return predicate.Test(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Test {
	return predicate.Test(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Test {
	return predicate.Test(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Test {
	return predicate.Test(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Test {
	return predicate.Test(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Test {
	return predicate.Test(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Test) predicate.Test {
	return predicate.Test(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Test) predicate.Test {
	return predicate.Test(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Test) predicate.Test {
	return predicate.Test(sql.NotPredicates(p))
}
