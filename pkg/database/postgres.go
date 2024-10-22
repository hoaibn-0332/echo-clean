package database

import (
	"database/sql"
	"echo-clean/pkg/logger"
	"entgo.io/ent/dialect"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectToPostgres(config DBConfig) (*sql.DB, error) {
	// Connect to postgres
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open(dialect.Postgres, dns)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to Postgres successfully")
	return db, nil
}
