package database

import (
	"fmt"
)

type Contact struct {
    Name string
    AddressLine1 string
    AddressLine2 string
    Email string
    Phone string
    Id int
}

type ErrorMap = map[string]string

func (c *Contact) validate() ErrorMap {
    var errors ErrorMap = make(ErrorMap)
    if c.Name == "" {
        errors["name"] = "Name is required"
    }

    if c.Email == "" {
        errors["email"] = "Email is required"
    }

    if c.AddressLine1 == "" {
        errors["addr1"] = "Address Line 1 is required"
    }

    if c.AddressLine2 == "" {
        errors["addr2"] = "Address Line 2 is required"
    }

    if c.Phone == "" {
        errors["phone"] = "Phone is required"
    }

    return errors
}

func DeleteContact(id int) error {
    _, err := Db.Exec("DELETE FROM contacts WHERE id = $1", id)
    if err != nil {
        return fmt.Errorf("unable to delete contact: %+v", err)
    }

    return nil
}

func GetContact(id int) (*Contact, error) {
    res, err := Db.Query("SELECT * FROM contacts WHERE id = $1", id)
    if err != nil {
        return nil, fmt.Errorf("unable to query db: %+v", err)
    }

    res.Next()
    var name string
    var email string
    var addressLine1 string
    var addressLine2 string
    var phone string
    var _id int

    err = res.Scan(&_id, &name, &email, &addressLine1, &addressLine2, &phone)
    if err != nil {
        return nil, fmt.Errorf("unable to scan db row: %+v", err)
    }

    res.Close()

    return &Contact{
        Name: name,
        AddressLine1: addressLine1,
        AddressLine2: addressLine2,
        Phone: phone,
        Email: email,
        Id: id,
    }, nil
}

func FilterContacts(q string) ([]Contact, error) {
    res, err := Db.Query(`
SELECT *
FROM contacts
WHERE (
    name LIKE $1 or
    email LIKE $1 or
    addressLine1 LIKE $1 or
    addressLine2 LIKE $1 or
    phone LIKE $1
)` , "%" + q + "%")

    if err != nil {
        return nil, fmt.Errorf("unable to query db: %+v", err)
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
            return nil, fmt.Errorf("could not scan db row: %+v", err)
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

    res.Close()

    return contacts, nil
}

func GetContacts() ([]Contact, error) {
    res, err := Db.Query("SELECT * FROM contacts")
    if err != nil {
        return nil, fmt.Errorf("unable to query db: %+v", err)
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
            return nil, fmt.Errorf("could not scan db row: %+v", err)
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

    res.Close()

    return contacts, nil
}

func HasEmail(email string) (bool, error) {
    res := Db.QueryRow("SELECT id FROM contacts WHERE email = ?", email)

    var id int
    err := res.Scan(&id)
    if err != nil {
        return false, fmt.Errorf("unable to query db: %+v", err)
    }

    return id > 0, nil
}

func (c *Contact) Save() (ErrorMap, error) {
    errors := c.validate()
    if len(errors) > 0 {
        return errors, nil
    }

    var err error
    if c.Id == -1 {
        _, err = Db.Exec(`INSERT INTO contacts (name, email, addressLine1, addressLine2, phone) VALUES (?, ?, ?, ?, ?)`, c.Name, c.Email, c.AddressLine1, c.AddressLine2, c.Phone)
    } else {
        _, err = Db.Exec(`UPDATE contacts SET name = ?, email = ?, addressLine1 = ?, addressLine2 = ?, phone = ? WHERE id = ?`, c.Name, c.Email, c.AddressLine1, c.AddressLine2, c.Phone, c.Id)
    }

    return errors, err
}

