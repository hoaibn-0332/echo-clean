package service

import (
	"echo-clean/domain/entity"
	"echo-clean/domain/repository"
)

type ArticleService struct {
	article repository.ArticleRepository
	author  repository.AuthorRepository
}

func NewArticleService(article repository.ArticleRepository, author repository.AuthorRepository) *ArticleService {
	return &ArticleService{
		article: article,
		author:  author,
	}
}

func (a *ArticleService) Fetch() ([]*entity.Article, error) {
	article, err := a.article.Fetch()
	if err != nil {
		return nil, err
	}

	return article, nil
}
