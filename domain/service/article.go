package service

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
)

// ArticleService is the interface that wraps the Fetch and Store methods
//
//go:generate mockery --name ArticleService
type ArticleService interface {
	Fetch() ([]*entity.Article, []serviceErrors.Error)
	Store(article *entity.Article) (*entity.Article, []serviceErrors.Error)
}
