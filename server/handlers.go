package server

import "echo-clean/server/handler"

// Handlers is the struct that holds all handlers
type Handlers struct {
	Article *handler.ArticleHandler
}

// NewHandlers creates a new instance of the Handlers struct
func NewHandlers(article *handler.ArticleHandler) *Handlers {
	return &Handlers{
		Article: article,
	}
}
