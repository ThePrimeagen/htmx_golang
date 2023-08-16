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

type Header struct {
    Title string
}

type ContactPage struct {
    Header
    Contacts []Contact
}

type NewContactPage struct {
    Header
    Contact *Contact
    Existing bool
    Errors map[string]string
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
        Header: Header{
            Title: "Contacts",
        },
        Contacts: contacts,
    })
}

func HandleCreateContact(c echo.Context) error {
    name := c.FormValue("name")
    email := c.FormValue("email")
    addressLine1 := c.FormValue("addr1")
    addressLine2 := c.FormValue("addr2")
    phone := c.FormValue("phone")

    var errors map[string]string = make(map[string]string)
    if name == "" {
        errors["name"] = "Name is required"
    }

    if email == "" {
        errors["email"] = "Email is required"
    }

    if  addressLine1 == "" {
        errors["addr1"] = "Address Line 1 is required"
    }

    if addressLine2 == "" {
        errors["addr2"] = "Address Line 2 is required"
    }

    if phone == "" {
        errors["phone"] = "Phone is required"
    }

    if len(errors) > 0 {
        c.Logger().Errorf("missing required fields")
        return c.Render(http.StatusOK, "new-contact", NewContactPage{
            Header: Header{
                Title: "Create Contact",
            },
            Contact: &Contact{
                Name: name,
                AddressLine1: addressLine1,
                AddressLine2: addressLine2,
                Email: email,
                Phone: phone,
            },
            Existing: false,
            Errors: errors,
        })
    }

    _, err := database.Db.Exec(`INSERT INTO contacts (name, email, addressLine1, addressLine2, phone) VALUES (?, ?, ?, ?, ?)`, name, email, addressLine1, addressLine2, phone)

    if err != nil {
        c.Logger().Errorf("could not insert into db: %+v", err)
        return c.String(http.StatusInternalServerError, "")
    }

    return c.Redirect(http.StatusFound, "/")
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

func HandleDeleteName(c echo.Context) error {

    if database.Db == nil {
        c.Logger().Error("db is nil")
        return c.String(http.StatusTeapot, "youu suck")
    }

    name := c.Param("name")
    _, _ = database.Db.Exec("DELETE FROM users WHERE name = ?", name)

    return c.Render(200, "name", nil)

}

