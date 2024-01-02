package routes

import (
	"net/http"

	"goh/go-htmx/components"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

func Index(ctx echo.Context) error {
	cc := &utils.HtmxContext{Context: ctx}
	name := "Stacy"
	if cc.IsHtmx() {
		name = "Randy"
	}

	return utils.Render(ctx, http.StatusOK, components.Base(components.Index(name)))
}
