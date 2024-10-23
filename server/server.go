package server

import (
	"context"
	"echo-clean/app/handler"
	"echo-clean/app/route"
	"echo-clean/config"
	"echo-clean/domain/service"
	"echo-clean/pkg/database"
	"echo-clean/pkg/repository"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer(cfg *config.Config) *Server {
	timeoutContext := time.Duration(cfg.Timeout) * time.Second

	e := echo.New()

	e.Use(SetCORS)
	e.Use(SetRequestContextWithTimeout(timeoutContext))

	// Connect to Postgres
	client, err := database.ConnectToPostgres(cfg.Db)
	if err != nil {
		log.Fatalf("Failed to connect to database, %s", err)
	}

	// Initialize repository
	articleRepository := repository.NewArticleRepository(client)
	authorRepository := repository.NewAuthorRepository(client)

	// Initialize service
	articleService := service.NewArticleService(articleRepository, authorRepository)

	//Initialize handler
	articleHandler := handler.NewArticleHandler(*articleService)

	//Register routes
	route.RegisterRoutes(e, articleHandler)

	return &Server{
		Echo: e,
	}
}

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
