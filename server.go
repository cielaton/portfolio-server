package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	echoServer := echo.New()

	echoServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}
