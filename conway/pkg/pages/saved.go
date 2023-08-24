package pages

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"theprimeagen.tv/conway/pkg/database"
)

type SavedPage struct {
    Conway []database.Conway
    Error string
}

type SavingResponse struct {
    Error string
    Success string
}

func Saved(c echo.Context) error {
    conways, err := database.GetSaved()

    if err != nil {
        return c.Render(200, "saved-list", SavedPage{
            Conway: []database.Conway{},
            Error: "Unable to fetch saved conways at this time",
        })
    }

    return c.Render(200, "saved-list", SavedPage {
        Conway: conways,
        Error: "",
    })
}

func SaveConway(saveAt bool) func(echo.Context)error {

    return func (c echo.Context) error {
        var seed string
        columnsStr := c.FormValue("columns")

        columns, err := strconv.Atoi(columnsStr)
        if err != nil {
            c.Logger().Errorf("couldn't convert columns to int: %v", err)
            return c.Render(200, "saved-msg", SavingResponse{
                Error: "bad data...",
                Success: "",
            })
        }

        if saveAt {
            seed = c.FormValue("seed-at")
        } else {
            seed = c.FormValue("seed")
        }

        _, err = database.SaveConway(seed, columns)
        if err != nil {
            c.Logger().Errorf("couldn't save conway into database: %v", err)
            return c.Render(200, "saved-msg", SavingResponse{
                Error: "unable to save...",
                Success: "",
            })
        }

        return c.Render(200, "saved-msg", SavingResponse{
            Error: "",
            Success: "saved!",
        })
    }
}
