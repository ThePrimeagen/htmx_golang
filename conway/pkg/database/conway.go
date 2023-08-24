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

func SaveConway(seed string, columns int) (int, error) {

    res, err := Db.Exec("INSERT INTO conway (seed, columns) VALUES (?, ?) RETURNING id", seed, columns)
    if err != nil {
        return 0, fmt.Errorf("couldn't insert conway: %w", err)
    }

    id, err := res.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("could get the id: %w", err)
    }

    return int(id), nil
}

func UpdateConway(seed string, columns, id int) error {

    _, err := Db.Exec("UPDATE conway SET seed = ?, columns = ? WHERE id = ?", seed, columns, id)
    if err != nil {
        return fmt.Errorf("couldn't insert conway: %w", err)
    }

    return nil
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

