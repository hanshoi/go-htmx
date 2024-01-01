package main

import (
	"log"

	"goh/go-htmx/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", routes.Index)

	log.Fatal(app.Start(":3000"))
}
