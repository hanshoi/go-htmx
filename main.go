package main

import (
	"log"

	"goh/go-htmx/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app)

	log.Fatal(app.Start(":3000"))
}
