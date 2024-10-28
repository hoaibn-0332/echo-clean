package main

import (
	"echo-clean/config"
	"echo-clean/pkg/logger"
	"echo-clean/server"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.LoadConfig()
	logger.InitLogger(cfg)

	if err != nil {
		log.Fatal().Msgf("Failed to load configuration, %s", err)
	}

	e := server.NewServer(cfg)

	if err := e.Echo.Start(cfg.Port); err != nil {
		log.Fatal().Msgf("Failed to start server, %s", err)
	}
}
