package server

import (
	"echo-clean/server/route"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	route.RegisterArticleRoutes(e, handlers.Article)

	// Health check the server
	e.GET("/health", HealthCheck)
}

func HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
