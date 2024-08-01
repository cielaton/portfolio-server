package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"portfolio-server/models"
)

func GetDevelopmentTools(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	query := models.QueryDevelopmentTools(database)
	return c.JSON(http.StatusOK, query)
}
