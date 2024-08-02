package main

import (
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"portfolio-server/database"
	"portfolio-server/handlers"
)

func databaseMiddleware(database *pgx.Conn) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("database", database)
			return next(c)
		}
	}
}

func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	// Create the echo server
	echoServer := echo.New()

	// Configure logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	echoServer.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, loggerValue middleware.RequestLoggerValues) error {
			log.Info().Str("URI", loggerValue.URI).Int("Status", loggerValue.Status).Msg("request")
			return nil
		}}))

	// Connect to the database
	postgreSQL := database.ConnectDB()
	echoServer.Use(databaseMiddleware(postgreSQL))

	echoServer.GET("/development-tools", handlers.GetDevelopmentTools)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
	defer func() {
		database.DisconnectDB(postgreSQL)
	}()
}
