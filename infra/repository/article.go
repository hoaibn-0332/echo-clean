package repository

import (
	"context"
	"echo-clean/domain/entity"
	"echo-clean/domain/repository"
	"echo-clean/ent"
	art "echo-clean/ent/article"
	"github.com/rs/zerolog/log"
	"time"
)

type ArticleRepository struct {
	client *ent.Client
}

func ParserArticle(article *ent.Article) *entity.Article {
	return &entity.Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Author:    *ParserAuthor(article.Edges.Author),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func (a ArticleRepository) Store(ctx context.Context, article *entity.Article, authorId int64) (*entity.Article, error) {
	ar, err := a.client.Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetAuthorID(authorId).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	fullAr, err := a.client.Article.Query().
		Where(art.IDEQ(ar.ID)).
		WithAuthor().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ParserArticle(fullAr), nil
}

func (a ArticleRepository) Fetch(ctx context.Context) ([]*entity.Article, error) {
	articles, err := a.client.Article.Query().
		WithAuthor().
		All(ctx)

	log.Debug().Msgf("Query error: %v", err)

	if err != nil {
		return nil, err
	}

	var result []*entity.Article
	for _, article := range articles {
		result = append(result, ParserArticle(article))
	}

	return result, nil
}

func (a ArticleRepository) GetByID(ctx context.Context, id int64) (*entity.Article, error) {
	article, err := a.client.Article.Query().
		Where(art.IDEQ(id)).
		WithAuthor().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ParserArticle(article), nil
}

func (a ArticleRepository) GetByTitle(ctx context.Context, title string) (*entity.Article, error) {
	article, err := a.client.Article.Query().
		Where(art.TitleEQ(title)).
		WithAuthor().
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return ParserArticle(article), nil
}

func (a ArticleRepository) Update(ctx context.Context, article *entity.Article) (*entity.Article, error) {
	ar, err := a.client.Article.UpdateOneID(article.ID).
		SetTitle(article.Title).
		SetContent(article.Content).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return ParserArticle(ar), nil
}

func (a ArticleRepository) Delete(ctx context.Context, id int64) error {
	err := a.client.Article.DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func NewArticleRepository(client *ent.Client) repository.ArticleRepository {
	return &ArticleRepository{
		client: client,
	}
}
