package database

import (
    "database/sql"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type Database struct {
    Db *sql.DB
}

func NewDatabase(url string) (*Database, error) {
    db, err := sql.Open("libsql", url)
    if err != nil {
        return nil, err
    }

    return &Database{db}, nil
}

