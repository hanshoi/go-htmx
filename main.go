package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"goh/go-htmx/routes"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := sql.Open("pgx", "postgresql://admin:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}
	defer db.Close()

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, db)

	log.Fatal(app.Start(":3000"))
}
