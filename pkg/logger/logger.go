package logger

import (
	"echo-clean/config"
	"github.com/rs/zerolog"
	"time"
)

// InitLogger initializes & config the logger
func InitLogger(config *config.Config) {
	zerolog.TimeFieldFormat = time.RFC3339

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config != nil && config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
