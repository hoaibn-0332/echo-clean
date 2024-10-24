// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticleColumns holds the columns for the "article" table.
	ArticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "author_id", Type: field.TypeInt64},
	}
	// ArticleTable holds the schema information for the "article" table.
	ArticleTable = &schema.Table{
		Name:       "article",
		Columns:    ArticleColumns,
		PrimaryKey: []*schema.Column{ArticleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_author_article",
				Columns:    []*schema.Column{ArticleColumns[5]},
				RefColumns: []*schema.Column{AuthorColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AuthorColumns holds the columns for the "author" table.
	AuthorColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AuthorTable holds the schema information for the "author" table.
	AuthorTable = &schema.Table{
		Name:       "author",
		Columns:    AuthorColumns,
		PrimaryKey: []*schema.Column{AuthorColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticleTable,
		AuthorTable,
	}
)

func init() {
	ArticleTable.ForeignKeys[0].RefTable = AuthorTable
	ArticleTable.Annotation = &entsql.Annotation{
		Table: "article",
	}
	AuthorTable.Annotation = &entsql.Annotation{
		Table: "author",
	}
}