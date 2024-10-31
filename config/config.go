package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

// DBConfig is a struct to store the database configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxCons  int
	MaxIdle  int
}

// Config is a struct to store the configuration
type Config struct {
	AppName string
	Port    string
	Debug   bool
	Timeout int
	Db      DBConfig
}

// LoadConfig is a function to load the .env file
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Debug().Msg("Error loading db port")
	}

	maxCons, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Debug().Msg("Error loading db max open conns")
	}

	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Debug().Msg("Error loading db max idle conns")
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		MaxCons:  maxCons,
		MaxIdle:  maxIdle,
	}

	config := Config{
		AppName: os.Getenv("APP_NAME"),
		Port:    os.Getenv("APP_PORT"),
		Debug:   true,
		Timeout: 2,
		Db:      dbConfig,
	}

	return &config, nil
}
