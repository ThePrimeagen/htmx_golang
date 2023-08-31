package pages

import (
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
    return c.Render(200, "index.html", nil)
}

func Subscribed(c echo.Context) error {
    time.Sleep(1 * time.Second)
    return c.Render(200, "subscribed.html", nil)
}

