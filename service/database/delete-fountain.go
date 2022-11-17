package database

func (db *appdbimpl) DeleteFountain(id uint64) error {
	res, err := db.c.Exec(`DELETE FROM fountains WHERE id=?`, id)
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
