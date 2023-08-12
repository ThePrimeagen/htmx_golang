package main
import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)


func main() {

    var dbUrl = "file:///home/mpaulson/personal/go_htmx/file.db"
    db, err := sql.Open("libsql", dbUrl)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
        os.Exit(1)
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, name TEXT, count INTEGER)")
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to create table: %s", err)
        os.Exit(1)
    }

    _, err = db.Exec("INSERT INTO items (name, count) VALUES (?, ?) ", "test", 2)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to create table: %s", err)
        os.Exit(1)
    }

    rows, err := db.Query("SELECT id, name, count FROM items")
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to query table: %s", err)
        os.Exit(1)
    }

    for rows.Next() {
        var id int
        var name string
        var count int
        err = rows.Scan(&id, &name, &count)
        if err != nil {
            fmt.Fprintf(os.Stderr, "failed to scan row: %s", err)
            os.Exit(1)
        }
        fmt.Printf("id: %d, name: %s, count: %d\n", id, name, count)
    }

}
