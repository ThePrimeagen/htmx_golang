package pages

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"theprimeagen.tv/conway/pkg/database"
)

type IndexPage struct {
    Conway *database.Conway
    Columns int
}

func baseIndex(c echo.Context) error {
    return c.Render(200, "index.html", IndexPage{
        Conway: nil,
        Columns: 40,
    })
}

func Index(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return baseIndex(c)
    }

    conway, err := database.GetConway(id)
    if err != nil {
        return baseIndex(c)
    }

    return c.Render(200, "index.html", IndexPage{
        Conway: conway,
        Columns: 40,
    })
}

