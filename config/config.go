package config

import (
	"echo-clean/pkg/logger"
	"github.com/spf13/viper"
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
	viper.AutomaticEnv()
	viper.SetConfigName("config.default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file, %s", err)
		return nil, err
	}

	dbConfig := DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	config := Config{
		AppName: viper.GetString("app.name"),
		Port:    viper.GetString("app.port"),
		Debug:   viper.GetBool("app.debug"),
		Timeout: viper.GetInt("app.timeout"),
		Db:      dbConfig,
	}

	return &config, nil
}
