package route

import (
	"echo-clean/domain/service"
	"github.com/labstack/echo/v4"
)

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(e *echo.Echo, svc service.ArticleService) {
	handler := &service.ArticleHandler{
		Service: svc,
	}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
	e.GET("/articles/:id", handler.GetByID)
	e.DELETE("/articles/:id", handler.Delete)
}
