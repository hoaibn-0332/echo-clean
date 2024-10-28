package handler

import (
	"echo-clean/domain/entity"
	serviceError "echo-clean/domain/error"
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

	return c.JSON(http.StatusOK, articles)
}

func (a *ArticleHandler) Store(c echo.Context) error {
	article := new(entity.Article)
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, []serviceError.Error{serviceError.UnsupportedMediaTypeError})
	}

	createdArticle, err := a.Service.Store(article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, createdArticle)
}
