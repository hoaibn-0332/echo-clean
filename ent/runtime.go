// Code generated by ent, DO NOT EDIT.

package ent

import (
	"echo-clean/ent/article"
	"echo-clean/ent/author"
	"echo-clean/pkg/database/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescTitle is the schema descriptor for title field.
	articleDescTitle := articleFields[1].Descriptor()
	// article.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	article.TitleValidator = articleDescTitle.Validators[0].(func(string) error)
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[3].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[4].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	authorFields := schema.Author{}.Fields()
	_ = authorFields
	// authorDescName is the schema descriptor for name field.
	authorDescName := authorFields[1].Descriptor()
	// author.NameValidator is a validator for the "name" field. It is called by the builders before save.
	author.NameValidator = authorDescName.Validators[0].(func(string) error)
	// authorDescCreatedAt is the schema descriptor for created_at field.
	authorDescCreatedAt := authorFields[2].Descriptor()
	// author.DefaultCreatedAt holds the default value on creation for the created_at field.
	author.DefaultCreatedAt = authorDescCreatedAt.Default.(time.Time)
	// authorDescUpdatedAt is the schema descriptor for updated_at field.
	authorDescUpdatedAt := authorFields[3].Descriptor()
	// author.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	author.DefaultUpdatedAt = authorDescUpdatedAt.Default.(time.Time)
}
