package database

func (db *appdbimpl) ListFountains() ([]Fountain, error) {
	var ret []Fountain

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT id, latitude, longitude, status FROM fountains`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var f Fountain
		err = rows.Scan(&f.ID, &f.Latitude, &f.Longitude, &f.Status)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
