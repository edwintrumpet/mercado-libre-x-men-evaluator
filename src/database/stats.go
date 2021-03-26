package database

type stat struct {
	ID    int
	Name  string
	Value int
}

func StatsGet() (stats []stat, err error) {
	q := `SELECT id, name, value FROM stats`

	db := GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		s := stat{}
		err = rows.Scan(&s.ID, &s.Name, &s.Value)

		if err != nil {
			return
		}

		stats = append(stats, s)
	}
	return stats, nil
}
