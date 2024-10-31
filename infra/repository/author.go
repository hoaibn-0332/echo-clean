package repository

import (
	"context"
	"echo-clean/domain/entity"
	"echo-clean/domain/repository"
	"echo-clean/ent"
	"time"
)

// AuthorRepository is the implementation of the AuthorRepository interface
type AuthorRepository struct {
	client *ent.Client
}

// ParserAuthor parses an Author entity to an Author entity
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

// Store stores an author
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

	return ParserAuthor(author), nil
}

// Fetch fetches all authors
func (a AuthorRepository) Fetch() ([]*entity.Author, error) {
	authors, err := a.client.Author.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	var result []*entity.Author
	for _, au := range authors {
		result = append(result, ParserAuthor(au))
	}
	return result, nil
}

// GetByID fetches an author by ID
func (a AuthorRepository) GetByID(id int64) (*entity.Author, error) {
	author, err := a.client.Author.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return ParserAuthor(author), nil
}

// Update updates an author
func (a AuthorRepository) Update(author *entity.Author) (*entity.Author, error) {
	au, err := a.client.Author.UpdateOneID(author.ID).
		SetName(author.Name).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return ParserAuthor(au), nil
}

// Delete deletes an author
func (a AuthorRepository) Delete(id int64) error {
	return a.client.Author.DeleteOneID(id).Exec(context.Background())
}

// NewAuthorRepository creates a new instance of the AuthorRepository
func NewAuthorRepository(client *ent.Client) repository.AuthorRepository {
	return &AuthorRepository{
		client: client,
	}
}
