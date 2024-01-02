package routes

import (
	"net/http"

	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

func CreateRoutes(app *echo.Echo) {
	app.GET("/", index)
	app.GET("/name/", name)
}

func index(ctx echo.Context) error {
	cc := &utils.HtmxContext{Context: ctx}
	name := "Stacy"
	if cc.IsHtmx() {
		name = "Randy"
	}

	return utils.RenderPage(ctx, http.StatusOK, templates.Index(name))
}

func name(ctx echo.Context) error {
	cc := &utils.HtmxContext{Context: ctx}
	name := "Lola"
	if cc.IsHtmx() {
		name = "Randy"
	}
	return utils.Render(ctx, http.StatusOK, templates.Name(name))
}
