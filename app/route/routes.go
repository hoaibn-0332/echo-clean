package route

import (
	"echo-clean/app/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, article *handler.ArticleHandler) {
	RegisterArticleRoutes(e, article)
}
