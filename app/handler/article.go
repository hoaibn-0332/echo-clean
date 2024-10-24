package handler

import (
	"echo-clean/domain/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ArticleHandler struct {
	Service service.ArticleService
}

func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		Service: service,
	}
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {
	articles, err := a.Service.Fetch()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, articles)
}
