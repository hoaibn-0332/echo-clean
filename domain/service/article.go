package service

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
)

//go:generate mockery --name ArticleService
type ArticleService interface {
	Fetch() ([]*entity.Article, []serviceErrors.Error)
	Store(article *entity.Article) (*entity.Article, []serviceErrors.Error)
}
