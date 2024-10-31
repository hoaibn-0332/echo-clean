package database

import (
	"echo-clean/config"
	"echo-clean/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"time"
)

// ConnectToPostgres connects to the postgres database
func ConnectToPostgres(config config.DBConfig) (*ent.Client, error) {
	// Connect to postgres
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	drv, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal().Msgf("Failed to openning connect to postgres %v", err)
		return nil, err
	} else {
		log.Debug().Msg("Connected to Postgres successfully")
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(config.MaxIdle)
	db.SetMaxOpenConns(config.MaxCons)
	db.SetConnMaxLifetime(time.Hour)

	return ent.NewClient(ent.Driver(drv)), nil
}
