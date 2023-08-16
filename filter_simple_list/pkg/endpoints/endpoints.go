package endpoints

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"theprimeagen.tv/filter_simple_list/pkg/database"
)

type Header struct {
    Title string
}

type ContactPage struct {
    Header
    Contacts []database.Contact
}

type NewContactPage struct {
    Header
    Contact *database.Contact
    Existing bool
    Errors map[string]string
}


func HandleIndex(c echo.Context) error {
    contacts, err := database.GetContacts()
    if err != nil {
        return c.String(http.StatusInternalServerError, "unable to get contacts")
    }

    return c.Render(200, "index.html", ContactPage {
        Header: Header{
            Title: "Contacts",
        },
        Contacts: contacts,
    })
}

func HandleDeleteContact(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.String(http.StatusBadRequest, "")
    }

    err = database.DeleteContact(id)

    return c.String(http.StatusOK, "")
}

func HandleEditContact(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.String(http.StatusBadRequest, "")
    }

    contact, err := database.GetContact(id)
    if err != nil {
        c.Logger().Errorf("could not retrieve contact: %+v", err)
        return c.String(http.StatusInternalServerError, "")
    }

    return c.Render(http.StatusOK, "new-contact", NewContactPage{
        Header: Header{
            Title: "Edit Contact",
        },
        Contact: contact,
        Existing: true,
        Errors: map[string]string{},
    })
}

func HandleSaveContact(c echo.Context) error {
    contact := database.Contact {
        Name: c.FormValue("name"),
        Email: c.FormValue("email"),
        AddressLine1: c.FormValue("addr1"),
        AddressLine2: c.FormValue("addr2"),
        Phone: c.FormValue("phone"),
        Id: -1,
    }

    errors, err := contact.Save()

    if err != nil {
        c.Logger().Errorf("could not save contact: %+v", err)
        return c.String(http.StatusInternalServerError, "")
    }

    if len(errors) > 0 {
        c.Logger().Errorf("missing required fields")
        return c.Render(http.StatusOK, "new-contact", NewContactPage{
            Header: Header{
                Title: "Create Contact",
            },
            Contact: &contact,
            Existing: false,
            Errors: errors,
        })
    }

    return c.Redirect(http.StatusOK, "/")
}

func HandleCreateContact(c echo.Context) error {
    contact := database.Contact {
        Name: c.FormValue("name"),
        Email: c.FormValue("email"),
        AddressLine1: c.FormValue("addr1"),
        AddressLine2: c.FormValue("addr2"),
        Phone: c.FormValue("phone"),
        Id: -1,
    }

    idStr := c.Param("id")
    if idStr != "" {
        idInt, err := strconv.Atoi(idStr)
        if err != nil {
            return c.String(http.StatusBadRequest, "")
        }
        contact.Id = idInt
    }

    errors, err := contact.Save()
    if err != nil {
        c.Logger().Errorf("could not insert into db: %+v", err)
        return c.String(http.StatusInternalServerError, "")
    }

    if len(errors) > 0 {
        c.Logger().Errorf("missing required fields")
        return c.Render(http.StatusOK, "new-contact", NewContactPage{
            Header: Header{
                Title: "Create Contact",
            },
            Contact: &contact,
            Existing: false,
            Errors: errors,
        })
    }

    return c.Redirect(http.StatusSeeOther, "/")
}

func HandleNewContact(c echo.Context) error {
    return c.Render(200, "new-contact", NewContactPage{
        Header: Header{
            Title: "Create Contact",
        },
        Contact: nil,
        Existing: false,
        Errors: map[string]string{},
    })
}

func HandleSettings(c echo.Context) error {
    return c.Render(200, "settings", Header{
        Title: "Settings",
    })
}

func HandleHelp(c echo.Context) error {
    return c.Render(200, "help", Header{
        Title: "Help",
    })
}
