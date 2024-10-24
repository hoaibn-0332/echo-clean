package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Author struct {
	ent.Schema
}

// Annotations of the User.
func (Author) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "author"},
	}
}

func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable().
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("name").
			NotEmpty().
			StructTag(`json:"name,omitempty"`),
		field.Time("created_at").
			Default(time.Now()).
			StructTag(`json:"created_at,omitempty"`),
		field.Time("updated_at").
			Default(time.Now()).
			StructTag(`json:"updated_at,omitempty`),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("article", Article.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
