package config

import (
	"echo-clean/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	AppName string
	Port    string
	Debug   bool
	Timeout int
	Db      DBConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logger.Debug("Error loading db port")
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
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
