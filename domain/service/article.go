package service

import (
	"echo-clean/domain/entity"
	sysErrors "echo-clean/domain/error"
	"echo-clean/domain/repository"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type ArticleService struct {
	base    BaseService
	article repository.ArticleRepository
	author  repository.AuthorRepository
}

func NewArticleService(article repository.ArticleRepository, author repository.AuthorRepository) *ArticleService {
	return &ArticleService{
		article: article,
		author:  author,
	}
}

func (a *ArticleService) Fetch() ([]*entity.Article, []sysErrors.Error) {
	article, err := a.article.Fetch()
	if err != nil {
		return nil, a.base.HandleDbError(err)
	}

	return article, nil
}

func (a *ArticleService) Store(article *entity.Article) (*entity.Article, []sysErrors.Error) {
	validate := validator.New()
	err := validate.Struct(article)
	log.Debug().Msgf("Validation error: %v", err)

	if err != nil {
		return nil, []sysErrors.Error{sysErrors.StatusBadRequestError}
	}

	authors, err := a.author.Fetch()
	if err != nil || len(authors) == 0 {
		return nil, a.base.HandleDbError(err)
	}

	firstAuthor := authors[0]

	createdArticle, err := a.article.Store(article, firstAuthor.ID)
	if err != nil {
		return nil, a.base.HandleDbError(err)
	}

	return createdArticle, nil
}
