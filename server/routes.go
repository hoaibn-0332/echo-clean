package server

import (
	"echo-clean/server/handler"
	"echo-clean/server/route"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, article *handler.ArticleHandler) {
	route.RegisterArticleRoutes(e, article)

	// Health check the server
	e.GET("/health", HealthCheck)
}

func HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
