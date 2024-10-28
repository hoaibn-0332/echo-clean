package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"time"
)

// SetRequestContextWithTimeout will set the request context with timeout for every incoming HTTP Request
func SetRequestContextWithTimeout(d time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx, cancel := context.WithTimeout(c.Request().Context(), d)
			defer cancel()

			newRequest := c.Request().WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

// SetCORS will handle the CORS middleware
func SetCORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// SessionMiddleware Set session tenant id to Postgresql when use multi-tenant and RLS
func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// handle input

		// set session tenant id to Postgresql

		// put data input to context for next handler
		return next(c)
	}
}
