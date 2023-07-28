package main

import (
	"log"

	"github.com/cimon-food-center/backend/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, "online")
	})

	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("error starting app:", err)
	}

	log.Println(*config)

	e.Logger.Fatal(e.Start(":8080"))
}
