package server

import (
	"echo-clean/config"
	"echo-clean/domain/service"
	"echo-clean/infra/database"
	repository2 "echo-clean/infra/repository"
	"echo-clean/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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
	e.Use(SessionMiddleware)

	// Connect to Postgres
	client, err := database.ConnectToPostgres(cfg.Db)
	if err != nil {
		log.Fatal().Msgf("Failed to connect to database, %s", err)
	}

	// Initialize repository
	articleRepository := repository2.NewArticleRepository(client)
	authorRepository := repository2.NewAuthorRepository(client)

	// Initialize service
	articleService := service.NewArticleService(articleRepository, authorRepository)

	//Initialize handler
	articleHandler := handler.NewArticleHandler(*articleService)

	//Register routes
	RegisterRoutes(e, articleHandler)

	return &Server{
		Echo: e,
	}
}
