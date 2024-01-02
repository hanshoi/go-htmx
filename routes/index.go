package routes

import (
	"net/http"

	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

func CreateRoutes(app *echo.Echo) {
	app.GET("/", lists)
	app.GET("/detail/", detail)
	app.GET("/about/", about)
}

func lists(ctx echo.Context) error {
	// cc := &utils.HtmxContext{Context: ctx}
	// name := "Stacy"
	// if cc.IsHtmx() {
	// 	name = "Randy"
	// }

	return utils.RenderPage(ctx, http.StatusOK, templates.List, templates.ListPage())
}

func detail(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.Detail, templates.DetailPage())
}

func about(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.About, templates.AboutPage())
}
