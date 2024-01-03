package main

import (
	"log"

	"goh/go-htmx/db"
	"goh/go-htmx/routes"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := db.New()
	defer db.Close()

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, &db)

	log.Fatal(app.Start(":3000"))
}
