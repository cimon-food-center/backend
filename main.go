package main

import (
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, "online")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
