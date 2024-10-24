package repository

import "echo-clean/domain/entity"

// ArticleRepository represent the article's repository contract
//
//go:generate mockery --name ArticleRepository
type ArticleRepository interface {
	Store(article *entity.Article, authorId int64) (*entity.Article, error)
	Fetch() ([]*entity.Article, error)
	GetByID(id int64) (*entity.Article, error)
	GetByTitle(title string) (*entity.Article, error)
	Update(article *entity.Article) (*entity.Article, error)
	Delete(id int64) error
}
