package server

import (
	"echo-clean/server/route"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RegisterRoutes Handlers is the struct that holds all handlers
func RegisterRoutes(e *echo.Echo, handlers *Handlers) {
	route.RegisterArticleRoutes(e, handlers.Article)

	// Health check the server
	e.GET("/health", HealthCheck)
}

// HealthCheck is the health check endpoint
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
