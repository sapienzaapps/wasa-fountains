package database

import "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/locationutils"

func (db *appdbimpl) ListFountainsWithFilter(latitude float64, longitude float64, filterRange float64) ([]Fountain, error) {

	// Here we need to get all fountains inside a given range. One simple solution is to rely on GIS/Spatial functions
	// from the DB itself. GIS/Spatial functions are those dedicated to geometry/geography/space computation.
	//
	// However, some databases (like SQLite) do not support these functions. So, we use a naive approach: instead of
	// drawing a circle for a given range, we get slightly more fountains by retrieving a square area, and then we will
	// filter the result later ("cutting the corner").
	//
	// Steps are:
	// 1. We compute a square ("bounding box") that contains the circle. The square will have edges with the same length
	//    of the range of the circle.
	// 2. For each resulting fountain, we will check (using Go and some math) if it's inside the range or not.

	const query = `
SELECT id, latitude, longitude, status
FROM fountains
WHERE ? <= latitude AND latitude <= ? AND ? <= longitude AND longitude <= ?`

	// Here we create a square / bounding box.
	// Note: we might have passed locationutils.Latitude and Longitude directly instead of casting to float64 and then
	// cast them back. However, for the sake of simpleness, we'll use float64 everywhere.
	var center = locationutils.Point2D{
		Latitude:  locationutils.Latitude(latitude),
		Longitude: locationutils.Longitude(longitude),
	}
	var boundingBox = locationutils.Squaring(center, filterRange)

	var ret []Fountain

	// Issue the query, using the bounding box as filter
	rows, err := db.c.Query(query,
		boundingBox.A.Latitude, boundingBox.B.Latitude,
		boundingBox.A.Longitude, boundingBox.B.Longitude)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var f Fountain
		err = rows.Scan(&f.ID, &f.Latitude, &f.Longitude, &f.Status)
		if err != nil {
			return nil, err
		}

		// Check if the result is inside the circle
		var fountainPoint = locationutils.Point2D{
			Latitude:  locationutils.Latitude(f.Latitude),
			Longitude: locationutils.Longitude(f.Longitude),
		}
		if locationutils.IsPointWithinRadius(center, fountainPoint, filterRange) {
			ret = append(ret, f)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
