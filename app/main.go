package main

import (
	"echo-clean/config"
	"echo-clean/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration, %s", err)
	}

	e := server.NewServer(cfg)

	if err := e.Echo.Start(cfg.Port); err != nil {
		log.Fatalf("Failed to start server, %s", err)
	}
}
