package handler

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
	"echo-clean/domain/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ArticleHandler is the handler for articles
type ArticleHandler struct {
	Service service.ArticleService
}

// NewArticleHandler creates a new ArticleHandler
func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		Service: service,
	}
}

// FetchArticle fetches all articles
func (a *ArticleHandler) FetchArticle(c echo.Context) error {
	articles, err := a.Service.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, articles)
}

// Store stores an article
func (a *ArticleHandler) Store(c echo.Context) error {
	article := new(entity.Article)
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, []serviceErrors.Error{serviceErrors.UnsupportedMediaTypeError})
	}

	createdArticle, err := a.Service.Store(c.Request().Context(), article)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, createdArticle)
}
