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

func ParserAuthor(auth *ent.Author) *entity.Author {
	if auth == nil {
		return nil
	}

	return &entity.Author{
		ID:        auth.ID,
		Name:      auth.Name,
		CreatedAt: auth.CreatedAt,
		UpdatedAt: auth.UpdatedAt,
	}
}

func (a AuthorRepository) Store(ctx context.Context, name string) (*entity.Author, error) {
	author, err := a.client.Author.
		Create().
		SetName(name).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return ParserAuthor(author), nil
}

func (a AuthorRepository) Fetch(ctx context.Context) ([]*entity.Author, error) {
	authors, err := a.client.Author.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*entity.Author
	for _, au := range authors {
		result = append(result, ParserAuthor(au))
	}
	return result, nil
}

func (a AuthorRepository) GetByID(ctx context.Context, id int64) (*entity.Author, error) {
	author, err := a.client.Author.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParserAuthor(author), nil
}

func (a AuthorRepository) Update(ctx context.Context, author *entity.Author) (*entity.Author, error) {
	au, err := a.client.Author.UpdateOneID(author.ID).
		SetName(author.Name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return ParserAuthor(au), nil
}

func (a AuthorRepository) Delete(ctx context.Context, id int64) error {
	return a.client.Author.DeleteOneID(id).Exec(ctx)
}

func NewAuthorRepository(client *ent.Client) repository.AuthorRepository {
	return &AuthorRepository{
		client: client,
	}
}
