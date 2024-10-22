package main

import (
	"database/sql"
	"echo-clean/app/route"
	"echo-clean/pkg/logger"
	"echo-clean/pkg/repository/mysql"
	"echo-clean/pkg/rest"
	middleware2 "echo-clean/pkg/rest/middleware"
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"echo-clean/article"
)

const (
	defaultAddress = ":9090"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigName("config.default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal("Error reading config file, %s", err)
	}
}

func main() {
	//prepare database
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbUser := viper.GetString("db.user")
	dbPass := viper.GetString("db.password")
	dbName := viper.GetString("db.name")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		logger.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		logger.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			logger.Fatal("got error when closing the DB connection", err)
		}
	}()
	// prepare echo

	e := echo.New()
	e.Use(middleware2.CORS)
	timeout := viper.GetUint("app.timeout")
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware2.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	authorRepo := mysql.NewAuthorRepository(dbConn)
	articleRepo := mysql.NewArticleRepository(dbConn)

	// Build service Layer
	svc := article.NewService(articleRepo, authorRepo)
	rest.NewArticleHandler(e, svc)

	route.RegisterRoutes(e)

	address := viper.GetString("app.port")
	if address == "" {
		address = defaultAddress
	}
	logger.Fatal(e.Start(address)) //nolint
}
