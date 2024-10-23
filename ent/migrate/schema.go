// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "author_article", Type: field.TypeInt64},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "articles_authors_article",
				Columns:    []*schema.Column{ArticlesColumns[5]},
				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:       "authors",
		Columns:    AuthorsColumns,
		PrimaryKey: []*schema.Column{AuthorsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		AuthorsTable,
	}
)

func init() {
	ArticlesTable.ForeignKeys[0].RefTable = AuthorsTable
}
