package database

func (db *appdbimpl) CreateFountain(f Fountain) (Fountain, error) {
	res, err := db.c.Exec(`INSERT INTO fountains (id, latitude, longitude, status) VALUES (?, ?, ?, ?)`,
		f.ID, f.Latitude, f.Longitude, f.Status)
	if err != nil {
		return f, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return f, err
	}

	f.ID = uint64(lastInsertID)
	return f, nil
}
