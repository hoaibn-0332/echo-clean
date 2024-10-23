package database

import (
	"echo-clean/config"
	"echo-clean/ent"
	"echo-clean/pkg/logger"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToPostgres(config config.DBConfig) (*ent.Client, error) {
	// Connect to postgres
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := ent.Open(dialect.Postgres, dns)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to Postgres successfully")
	return db, nil
}
