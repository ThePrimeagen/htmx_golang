package database

import "fmt"

type Conway struct {
    Id int
    Seed string
    Columns int
}

func GetConway(id int) (*Conway, error) {
    row := Db.QueryRow("SELECT * FROM conway WHERE id = ?", id)

    var seed string
    var columns int

    err := row.Scan(&id, &seed, &columns)
    if err != nil {
        return nil, fmt.Errorf("couldn't scan row: %w", err)
    }

    return &Conway{
        Id: id,
        Seed: seed,
        Columns: columns,
    }, nil
}

func GetSaved() ([]Conway, error) {

    rows, err := Db.Query("SELECT * FROM conway")
    if err != nil {
        return nil, fmt.Errorf("couldn't get saved conways: %w", err)
    }

    defer rows.Close()

    conways := []Conway{}

    for rows.Next() {

        var id int
        var seed string
        var columns int

        err = rows.Scan(&id, &seed, &columns)
        if err != nil {
            return nil, fmt.Errorf("failed on scanning the row: %w", err)
        }

        conways = append(conways, Conway{
            Id: id,
            Seed: seed,
            Columns: columns,
        })
    }

    return conways, nil
}

