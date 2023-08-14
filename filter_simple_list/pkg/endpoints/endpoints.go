package endpoints

import (
	"github.com/labstack/echo/v4"
)

func HandleIndex(c echo.Context) error {
    return c.Render(200, "index.html", nil)
}

