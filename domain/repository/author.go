package repository

import (
	"context"
	"echo-clean/domain/entity"
)

// AuthorRepository represent the author's repository contract
//
//go:generate mockery --name AuthorRepository
type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (entity.Author, error)
}
