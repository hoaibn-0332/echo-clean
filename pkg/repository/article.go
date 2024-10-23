package repository

import (
	"echo-clean/domain/entity"
	"echo-clean/domain/repository"
	"echo-clean/ent"
)

type ArticleRepository struct {
	client *ent.Client
}

func (a ArticleRepository) Store(article *entity.Article, authorId int) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Fetch() ([]*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetByID(id uint64) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) GetByTitle(title string) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Update(article *entity.Article) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepository) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func NewArticleRepository(client *ent.Client) repository.ArticleRepository {
	return &ArticleRepository{
		client: client,
	}
}
