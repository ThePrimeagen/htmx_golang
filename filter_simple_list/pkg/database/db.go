package database

import (
	"database/sql"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDB(url string) error {
    db, err := sql.Open("libsql", url)

    if err != nil {
        return err
    }

    // Real gross, but guess what, we are doing database initialization here
    db.Exec(`CREATE TABLE IF NOT EXISTS users (
            name TEXT NOT NULL,
            wins INTEGER
        )`)

    Db = db

    return nil
}

func InitContacts(url string) error {
    db, err := sql.Open("libsql", url)

    if err != nil {
        return err
    }

    // Real gross, but guess what, we are doing database initialization here
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS contacts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT NOT NULL,
            addressLine1 TEXT NOT NULL,
            addressLine2 TEXT NOT NULL,
            phone TEXT NOT NULL
        )`)

    if err != nil {
        return err
    }

    Db = db

    return nil
}
