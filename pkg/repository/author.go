package repository

import (
	"context"
	"echo-clean/domain/entity"
	"echo-clean/domain/repository"
	"echo-clean/ent"
	"time"
)

type AuthorRepository struct {
	client *ent.Client
}

func (a AuthorRepository) Store(name string) (*entity.Author, error) {
	author, err := a.client.Author.
		Create().
		SetName(name).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return &entity.Author{
		ID:        author.ID,
		Name:      author.Name,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

func (a AuthorRepository) Fetch() ([]*entity.Author, error) {
	authors, err := a.client.Author.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	var result []*entity.Author
	for _, au := range authors {
		result = append(result, &entity.Author{
			ID:        au.ID,
			Name:      au.Name,
			CreatedAt: au.CreatedAt,
			UpdatedAt: au.UpdatedAt,
		})
	}
	return result, nil
}

func (a AuthorRepository) GetByID(id int64) (*entity.Author, error) {
	author, err := a.client.Author.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &entity.Author{
		ID:        author.ID,
		Name:      author.Name,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil
}

func (a AuthorRepository) Update(author *entity.Author) (*entity.Author, error) {
	au, err := a.client.Author.UpdateOneID(author.ID).
		SetName(author.Name).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return &entity.Author{
		ID:        au.ID,
		Name:      au.Name,
		CreatedAt: au.CreatedAt,
		UpdatedAt: au.UpdatedAt,
	}, nil
}

func (a AuthorRepository) Delete(id int64) error {
	return a.client.Author.DeleteOneID(id).Exec(context.Background())
}

func NewAuthorRepository(client *ent.Client) repository.AuthorRepository {
	return &AuthorRepository{
		client: client,
	}
}
