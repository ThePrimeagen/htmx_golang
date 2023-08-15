package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"theprimeagen.tv/filter_simple_list/pkg/database"
)

type Contact struct {
    Name string
    AddressLine1 string
    AddressLine2 string
    Email string
    Phone string
    Id int
}

type ContactPage struct {
    Contacts []Contact
}

func HandleIndex(c echo.Context) error {
    res, err := database.Db.Query("SELECT * FROM contacts")
    if err != nil {
        c.Logger().Errorf("Could not query the db %+v", err)
        return c.String(http.StatusInternalServerError, "")
    }

    var contacts []Contact = make([]Contact, 0)
    for res.Next() {
        var name string
        var email string
        var addressLine1 string
        var addressLine2 string
        var phone string
        var id int

        err := res.Scan(&id, &name, &email, &addressLine1, &addressLine2, &phone)
        if err != nil {
            c.Logger().Errorf("could not scan: %+v", err)
            return c.String(http.StatusInternalServerError, "")
        }

        contacts = append(contacts, Contact{
            Name: name,
            AddressLine1: addressLine1,
            AddressLine2: addressLine2,
            Phone: phone,
            Email: email,
            Id: id,
        })
    }

    return c.Render(200, "index.html", ContactPage {
        Contacts: contacts,
    })
}

type Name struct {
    Name string
}

func HandleCreateName(c echo.Context) error {
    if database.Db == nil {
        c.Logger().Error("db is nil")
        return c.String(http.StatusTeapot, "youu suck")
    }

    name := c.FormValue("name")
    res := database.Db.QueryRow("SELECT * FROM users WHERE name = ?", name)

    var rowName string
    err := res.Scan(&rowName)
    if err == nil {
        return err
    }

    return c.Render(200, "name", Name{
        Name: name,
    })
}

func HandleDeleteName(c echo.Context) error {

    if database.Db == nil {
        c.Logger().Error("db is nil")
        return c.String(http.StatusTeapot, "youu suck")
    }

    name := c.Param("name")
    _, _ = database.Db.Exec("DELETE FROM users WHERE name = ?", name)

    return c.Render(200, "name", nil)

}

