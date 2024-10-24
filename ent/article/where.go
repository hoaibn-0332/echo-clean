// Code generated by ent, DO NOT EDIT.

package article

import (
	"echo-clean/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.Author) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(sql.NotPredicates(p))
}
