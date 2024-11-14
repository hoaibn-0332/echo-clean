package service

import (
	"context"
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
	"echo-clean/domain/repository"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// ArticleServiceImpl is the implementation of the ArticleService interface
type ArticleServiceImpl struct {
	base    BaseService
	article repository.ArticleRepository
	author  repository.AuthorRepository
}

// NewArticleService creates a new instance of the ArticleServiceImpl
func NewArticleService(article repository.ArticleRepository, author repository.AuthorRepository) ArticleService {
	return &ArticleServiceImpl{
		base:    BaseService{},
		article: article,
		author:  author,
	}
}

func (a ArticleServiceImpl) Fetch(ctx context.Context) ([]*entity.Article, []serviceErrors.Error) {
	article, err := a.article.Fetch(ctx)
	if err != nil {
		return nil, a.base.HandleDbError(err)
	}

	return article, nil
}

func (a ArticleServiceImpl) Store(ctx context.Context, article *entity.Article) (*entity.Article, []serviceErrors.Error) {
	validate := validator.New()
	err := validate.Struct(article)
	log.Debug().Msgf("Validation error: %v", err)

	if err != nil {
		return nil, []serviceErrors.Error{serviceErrors.StatusBadRequestError}
	}

	authors, err := a.author.Fetch(ctx)
	if err != nil || len(authors) == 0 {
		return nil, a.base.HandleDbError(err)
	}

	firstAuthor := authors[0]

	createdArticle, err := a.article.Store(ctx, article, firstAuthor.ID)
	if err != nil {
		return nil, a.base.HandleDbError(err)
	}

	return createdArticle, nil
}
