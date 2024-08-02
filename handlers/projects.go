package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"portfolio-server/models"
)

func GetPersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	result := models.QueryProjects(database, "personal_projects")
	return c.JSON(http.StatusOK, result)
}
func GetTeamProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	result := models.QueryProjects(database, "team_projects")
	return c.JSON(http.StatusOK, result)
}
