package repository

import (
	"context"
	"echo-clean/domain/entity"
)

// ArticleRepository represent the article's repository contract
//
//go:generate mockery --name ArticleRepository
type ArticleRepository interface {
	Store(ctx context.Context, article *entity.Article, authorId int64) (*entity.Article, error)
	Fetch(ctx context.Context) ([]*entity.Article, error)
	GetByID(ctx context.Context, id int64) (*entity.Article, error)
	GetByTitle(ctx context.Context, title string) (*entity.Article, error)
	Update(ctx context.Context, article *entity.Article) (*entity.Article, error)
	Delete(ctx context.Context, id int64) error
}
