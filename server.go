package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"portfolio-server/database"
	"portfolio-server/models"
)

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

	echoServer.GET("/", func(c echo.Context) error {
		query := models.QueryDevelopmentTools(postgreSQL)
		return c.JSON(http.StatusOK, query)
	})
	echoServer.Logger.Fatal(echoServer.Start(":1323"))
	defer func() {
		database.DisconnectDB(postgreSQL)
	}()
}
