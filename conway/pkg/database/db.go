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
    db.Exec(`CREATE TABLE IF NOT EXISTS conway (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            seed TEXT NOT NULL,
            columns INTEGER NOT NULL
        )`)

    Db = db

    return nil
}
