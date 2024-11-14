package repository

import (
	"context"
	"echo-clean/domain/entity"
)

// AuthorRepository represent the author's repository contract
//
//go:generate mockery --name AuthorRepository
type AuthorRepository interface {
	Store(ctx context.Context, name string) (*entity.Author, error)
	Fetch(ctx context.Context) ([]*entity.Author, error)
	GetByID(ctx context.Context, id int64) (*entity.Author, error)
	Update(ctx context.Context, author *entity.Author) (*entity.Author, error)
	Delete(ctx context.Context, id int64) error
}
