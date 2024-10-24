package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Article struct {
	ent.Schema
}

// Annotations of the Article.
func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("article"),
	}
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Immutable().
			Unique().
			StructTag(`json:"id,omitempty"`),
		field.String("title").
			NotEmpty().
			StructTag(`json:"title,omitempty"`),
		field.String("content").
			StructTag(`json:"content,omitempty"`),
		field.Int64("author_id").
			StructTag(`json:"author_id,omitempty"`),
		field.Time("created_at").
			Default(time.Now).
			StructTag(`json:"created_at,omitempty"`),
		field.Time("updated_at").
			Default(time.Now).
			StructTag(`json:"updated_at,omitempty"`),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
			Ref("article").
			Unique().
			Required().
			Field("author_id"),
	}
}
