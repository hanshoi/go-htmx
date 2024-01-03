package routes

import (
	"net/http"

	"goh/go-htmx/db"
	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

type MyHandlerFunction func(echo.Context, *db.DB) error

func CreateRoutes(app *echo.Echo, db *db.DB) {
	app.GET("/", wrap(lists, db))
	app.GET("/list/add/", addListModal)
	app.GET("/detail/", wrap(detail, db))
	app.GET("/about/", about)
}

func wrap(handler MyHandlerFunction, db *db.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(ctx, db)
	}
}

func lists(ctx echo.Context, db *db.DB) error {
	// cc := &utils.HtmxContext{Context: ctx}
	// name := "Stacy"
	// if cc.IsHtmx() {
	// 	name = "Randy"
	// }
	lists := db.Lists.GetLists()
	if len(lists) == 0 {
		return utils.RenderPage(ctx, http.StatusOK, templates.ListPageType, templates.EmptyListPage())
	}
	return utils.RenderPage(ctx, http.StatusOK, templates.ListPageType, templates.ListPage(lists))
}

func addListModal(ctx echo.Context) error {
	return utils.Render(ctx, http.StatusOK, templates.Modal(templates.AddList()))
}

func detail(ctx echo.Context, db *db.DB) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.DetailPageType, templates.DetailPage())
}

func about(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.AboutPageType, templates.AboutPage())
}
