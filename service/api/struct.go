package api

import "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

const (
	FountainStatusGood   string = "good"
	FountainStatusFaulty string = "faulty"
)

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type Fountain struct {
	ID        uint64  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (f *Fountain) FromDatabase(fountain database.Fountain) {
	f.ID = fountain.ID
	f.Latitude = fountain.Latitude
	f.Longitude = fountain.Longitude
	f.Status = fountain.Status
}

// ToDatabase returns the fountain in a database-compatible representation
func (f *Fountain) ToDatabase() database.Fountain {
	return database.Fountain{
		ID:        f.ID,
		Latitude:  f.Latitude,
		Longitude: f.Longitude,
		Status:    f.Status,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either FountainStatusGood or FountainStatusFaulty. Note that the ID is not checked, as fountains
// read from requests have zero IDs as the user won't send us the ID in that way.
func (f *Fountain) IsValid() bool {
	return -90 <= f.Latitude && f.Latitude <= 90 &&
		-180 <= f.Longitude && f.Longitude <= 180 &&
		(f.Status == FountainStatusGood || f.Status == FountainStatusFaulty)
}
