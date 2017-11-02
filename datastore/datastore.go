// -*- Mode: Go; indent-tabs-mode: t -*-

package datastore

import (
	"database/sql"

	// PostgreSQL adapter
	_ "github.com/lib/pq"
)

var db *sql.DB

// Version holds the version of the service
const Version = "0.1" 

const (
	createBooking = `CREATE TABLE IF NOT EXISTS booking (
		id serial primary key not null,
		created timestamp DEFAULT NOW(),
		book_date date not null,
		name varchar(200) not null,
		playing boolean
	)`
)


// CreateDatabase creates the database tables
func CreateDatabase() error {
    if _, err := db.Exec(createBooking); err != nil {
        return nil
	}
	
	return nil
}