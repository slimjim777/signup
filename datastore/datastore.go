// -*- Mode: Go; indent-tabs-mode: t -*-

package datastore

import (
	"database/sql"
	"log"
	"os"

	// PostgreSQL adapter
	_ "github.com/lib/pq"
)

// DB local database interface with our custom methods.
type DB struct {
	*sql.DB
}

var dbConnection *DB

// Version holds the version of the service
const Version = "0.2"

const (
	createBookingEvent = `CREATE TABLE IF NOT EXISTS bookingevent (
		id serial primary key not null,
		created timestamp DEFAULT NOW(),
		book_date date not null,
		name varchar(200) not null
	)`

	createBookingDate = `CREATE TABLE IF NOT EXISTS bookingdate (
		id serial primary key not null,
		created timestamp DEFAULT NOW(),
		event_id int references bookingevent,
		name varchar(200) not null,
		playing boolean
	)`

	createBooking = `CREATE TABLE IF NOT EXISTS booking (
		id serial primary key not null,
		created timestamp DEFAULT NOW(),
		book_date date not null,
		name varchar(200) not null,
		playing boolean
	)`

	createBookingIndex = `CREATE UNIQUE INDEX IF NOT EXISTS booking_idx ON booking (book_date DESC, name)`
)

// CreateDatabase creates the database tables
func CreateDatabase() error {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	dbConnection = &DB{db}

	if _, err := dbConnection.Exec(createBooking); err != nil {
		return nil
	}

	if _, err := dbConnection.Exec(createBookingIndex); err != nil {
		return nil
	}

	return nil
}

