package database

func (db *appdbimpl) ListFountains() ([]Fountain, error) {
	var ret []Fountain

	rows, err := db.c.Query(`SELECT id, latitude, longitude, status FROM fountains`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}

		var f Fountain
		err = rows.Scan(&f.ID, &f.Latitude, &f.Longitude, &f.Status)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}

	return ret, nil
}
