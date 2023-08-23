package endpoints

import (
	"net/http"
	"net/mail"
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
    QueryString string
}

type NewContactPage struct {
    Header
    Contact *database.Contact
    Existing bool
    Errors map[string]string
}

func valid(email string) bool {
    _, err := mail.ParseAddress(email)

    return err == nil
}

func HandleValidateEmail(c echo.Context) error {
    email := c.FormValue("email")
    if email == "" {
        return c.String(http.StatusOK, "email is required")
    }
    if !valid(email) {
        return c.String(http.StatusOK, "invalid email")
    }

    hasEmail, err := database.HasEmail(email)
    c.Logger().Errorf("hasEmail %t and err %+v", hasEmail, err)
    if err != nil || !hasEmail {
        return c.String(http.StatusOK, "")
    }

    return c.String(http.StatusOK, "email already exists")
}

func HandleIndex(c echo.Context) error {
    q := c.QueryParam("q")

    var contacts []database.Contact
    var err error


    if q != "" {
        contacts, err = database.FilterContacts(q)
    } else {
        contacts, err = database.GetContacts()
    }

    if err != nil {
        c.Logger().Errorf("could not retrieve contacts: %+v", err)
        return c.String(http.StatusInternalServerError, "unable to get contacts")
    }

    trigger := c.Request().Header.Get("HX-Trigger")
    c.Logger().Errorf("trigger %s", trigger)
    if trigger == "search" {
        return c.Render(200, "contact-list", ContactPage {
            Contacts: contacts,
        });

    }
    return c.Render(200, "index.html", ContactPage {
        Header: Header{
            Title: "Contacts",
        },
        Contacts: contacts,
        QueryString: q,
    })
}

func HandleDeleteContact(c echo.Context) error {
    id := c.Param("id")

    if id != "" {
        idInt, err := strconv.Atoi(id)
        if err != nil {
            c.Logger().Errorf("could not convert id to int: %+v", err)
            return c.String(http.StatusBadRequest, "")
        }
        err = database.DeleteContact(idInt)
        if err != nil {
            return c.String(http.StatusInternalServerError, "")
        }

        return c.String(http.StatusOK, "")
    }

    manyIds := c.Param("selected_contact_ids")
    if manyIds == "" {
        c.Logger().Errorf("both ids and manyIds were empty")
        return c.String(http.StatusBadRequest, "")
    }

    return c.String(http.StatusBadRequest, "")
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
