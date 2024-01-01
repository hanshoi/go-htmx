package routes

import (
	"net/http"

	"goh/go-htmx/components"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
)

func Index(ctx echo.Context) error {
	return utils.Render(ctx, http.StatusOK, components.Index("John"))
}
