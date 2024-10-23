package route

import (
	"echo-clean/app/handler"
	"github.com/labstack/echo/v4"
)

func RegisterArticleRoutes(e *echo.Echo, handler *handler.ArticleHandler) {
	e.GET("/articles", handler.FetchArticle)
}
