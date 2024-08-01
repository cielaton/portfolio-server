package main

import (
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
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
		log.Fatal("Error loading .env file")
	}

	// Create the echo server
	echoServer := echo.New()

	// Connect to the database
	postgreSQL := database.ConnectDB()
	echoServer.Use(databaseMiddleware(postgreSQL))

	echoServer.GET("/development-tools", handlers.GetDevelopmentTools)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
	defer func() {
		database.DisconnectDB(postgreSQL)
	}()
}
