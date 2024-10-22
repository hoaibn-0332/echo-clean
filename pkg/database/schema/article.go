package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Article struct {
	ent.Schema
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Immutable().
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("title").
			NotEmpty().
			StructTag(`json:"title,omitempty"`),
		field.String("content").
			StructTag(`json:"content,omitempty"`),
		field.Time("created_at").
			Default(time.Now).
			StructTag(`json:"created_at,omitempty"`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			StructTag(`json:"updated_at,omitempty"`),
	}
}

func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
			Ref("article").
			Unique().
			Required(),
	}
}
