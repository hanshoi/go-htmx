package routes

import (
	"database/sql"
	"net/http"

	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

type MyHandlerFunction func(echo.Context, *sql.DB) error

func CreateRoutes(app *echo.Echo, db *sql.DB) {
	app.GET("/", wrap(lists, db))
	app.GET("/detail/", wrap(detail, db))
	app.GET("/about/", about)
}

func wrap(handler MyHandlerFunction, db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(ctx, db)
	}
}

func lists(ctx echo.Context, db *sql.DB) error {
	// cc := &utils.HtmxContext{Context: ctx}
	// name := "Stacy"
	// if cc.IsHtmx() {
	// 	name = "Randy"
	// }

	return utils.RenderPage(ctx, http.StatusOK, templates.List, templates.ListPage())
}

func detail(ctx echo.Context, db *sql.DB) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.Detail, templates.DetailPage())
}

func about(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.About, templates.AboutPage())
}
