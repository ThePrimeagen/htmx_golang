package pages

import (
	"github.com/labstack/echo/v4"
	"theprimeagen.tv/conway/pkg/database"
)

type SavedPage struct {
    Conway []database.Conway
    Error string
}

func Saved(c echo.Context) error {
    conways, err := database.GetSaved()

    if err != nil {
        return c.Render(200, "saved-list", SavedPage{
            Conway: []database.Conway{},
            Error: "Unable to fetch saved conways at this time",
        })
    }

    return c.Render(200, "saved-list", nil)
}

