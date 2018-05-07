// -*- Mode: Go; indent-tabs-mode: t -*-

package datastore

import (
	"database/sql"
	"log"
	"time"
)

const (
	upsertBooking = `
		INSERT INTO booking (book_date, name, playing)
		VALUES ($1, $2, $3)
		ON CONFLICT (book_date, name)
			DO UPDATE SET playing = $3 WHERE booking.book_date=$1 AND booking.name=$2
	`

	listBookingsForDate = `SELECT id, created, book_date, name, playing from booking where book_date=$1`
)

// Booking holds a booking
type Booking struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	Date    string    `json:"date"`
	Name    string    `json:"name"`
	Playing bool      `json:"playing"`
}

// BookingUpsert upserts a booking for a date
func BookingUpsert(name, date string, playing bool) error {
	if _, err := dbConnection.Exec(upsertBooking, date, name, playing); err != nil {
		return err
	}

	return nil
}

// BookingList fetches bookings for a date
func BookingList(date string) ([]Booking, error) {

	bookings := []Booking{}
	var rows *sql.Rows

	rows, err := dbConnection.Query(listBookingsForDate, date)
	if err != nil {
		log.Printf("Error retrieving bookings: %v\n", err)
		return bookings, err
	}

	defer rows.Close()
	for rows.Next() {
		b := Booking{}
		err = rows.Scan(&b.ID, &b.Created, &b.Date, &b.Name, &b.Playing)
		if err != nil {
			return bookings, err
		}
		bookings = append(bookings, b)
	}

	return bookings, nil
}
