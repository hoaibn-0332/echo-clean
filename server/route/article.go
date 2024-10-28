package route

import (
	"echo-clean/server/handler"
	"github.com/labstack/echo/v4"
)

func RegisterArticleRoutes(e *echo.Echo, handler *handler.ArticleHandler) {
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
}
