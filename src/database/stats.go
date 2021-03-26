package database

import "errors"

// StatsGet preform a query in database to obtain stats
func StatsGet() (stats []Stat, err error) {
	q := `SELECT name, value
					FROM stats`

	db := GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		s := Stat{}
		err = rows.Scan(&s.Name, &s.Value)

		if err != nil {
			return
		}

		stats = append(stats, s)
	}
	return stats, nil
}

// IncrementCount perform a query in database to increment
// the correspondong value accord to a string given
func IncrementCount(t string) error {
	q := `UPDATE stats
					SET value = value + 1
					WHERE name = $1`

	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(t)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("only one row should have been affected")
	}

	return nil
}
