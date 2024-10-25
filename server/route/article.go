package route

import (
	"echo-clean/server/handler"
	"github.com/labstack/echo/v4"
	"log"
)

func RegisterArticleRoutes(e *echo.Echo, handler *handler.ArticleHandler) {
	log.Printf("Registering article routes")
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles", handler.Store)
}
