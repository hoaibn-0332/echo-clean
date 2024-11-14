package service

import (
	"context"
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
)

//go:generate mockery --name ArticleService
type ArticleService interface {
	Fetch(ctx context.Context) ([]*entity.Article, []serviceErrors.Error)
	Store(ctx context.Context, article *entity.Article) (*entity.Article, []serviceErrors.Error)
}
