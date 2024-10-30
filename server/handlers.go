package server

import "echo-clean/server/handler"

type Handlers struct {
	Article *handler.ArticleHandler
}

func NewHandlers(article *handler.ArticleHandler) *Handlers {
	return &Handlers{
		Article: article,
	}
}
