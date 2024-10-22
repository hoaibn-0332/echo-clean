package repository

import "database/sql"

type ArticleRepository struct {
	Conn *sql.DB
}

func NewArticleRepository(conn *sql.DB) *ArticleRepository {
	return &ArticleRepository{conn}
}
