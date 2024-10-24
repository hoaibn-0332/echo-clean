package repository

import (
	"echo-clean/domain/entity"
)

// AuthorRepository represent the author's repository contract
//
//go:generate mockery --name AuthorRepository
type AuthorRepository interface {
	Store(name string) (*entity.Author, error)
	Fetch() ([]*entity.Author, error)
	GetByID(id int64) (*entity.Author, error)
	Update(author *entity.Author) (*entity.Author, error)
	Delete(id int64) error
}
