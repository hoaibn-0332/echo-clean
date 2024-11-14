package service

import (
	"context"
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
)

// ArticleService is the interface that wraps the Fetch and Store methods
//
//go:generate mockery --name ArticleService
type ArticleService interface {
	Fetch(ctx context.Context) ([]*entity.Article, []serviceErrors.Error)
	Store(ctx context.Context, article *entity.Article) (*entity.Article, []serviceErrors.Error)
}
