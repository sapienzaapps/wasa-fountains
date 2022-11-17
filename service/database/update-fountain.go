package database

func (db *appdbimpl) UpdateFountain(f Fountain) error {
	res, err := db.c.Exec(`UPDATE fountains SET latitude=?, longitude=?, status=? WHERE id=?`,
		f.Latitude, f.Longitude, f.Status, f.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrFountainDoesNotExist
	}
	return nil
}
