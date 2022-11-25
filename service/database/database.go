/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrFountainDoesNotExist = errors.New("fountain does not exist")

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
type Fountain struct {
	ID        uint64
	Latitude  float64
	Longitude float64
	Status    string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ListFountains returns the list of fountains with their status
	ListFountains() ([]Fountain, error)

	// ListFountainsWithFilter returns the list of fountains with their status, filtered using the specified parameters
	ListFountainsWithFilter(latitude float64, longitude float64, filterRange float64) ([]Fountain, error)

	// CreateFountain creates a new fountain in the database. It returns an updated Fountain object (with the ID)
	CreateFountain(Fountain) (Fountain, error)

	// UpdateFountain updates the fountain, replacing every value with those provided in the argument
	UpdateFountain(Fountain) error

	// DeleteFountain removes the fountain with the given ID
	DeleteFountain(id uint64) error

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='fountains';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE fountains (
    id INTEGER NOT NULL PRIMARY KEY,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    status TEXT NOT NULL
);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
