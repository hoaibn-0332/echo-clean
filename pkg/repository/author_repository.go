package repository

import "database/sql"

type AuthorRepository struct {
	Conn *sql.DB
}

func NewAuthorRepository(conn *sql.DB) *AuthorRepository {
	return &AuthorRepository{conn}
}
