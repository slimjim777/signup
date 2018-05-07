// -*- Mode: Go; indent-tabs-mode: t -*-

package datastore

import (
	"database/sql"
	"log"
	"time"
)

const (
	insertBookingEvent = `
		INSERT INTO bookingevent (book_date, name)
		VALUES ($1, $2)
	`

	updateBookingEvent = `
		UPDATE bookingevent
		SET book_date=$2, name=$3
		WHERE id=$1
	`

	listBookingEvents = `
		SELECT id, created, book_date, name from bookingevent
		WHERE book_date >= current_date
		ORDER BY book_date, name`
)

// BookingEvent holds a booking event
type BookingEvent struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
	Date    string    `json:"date"`
	Name    string    `json:"name"`
}

// BookingEventUpsert upserts a event for a date
func BookingEventUpsert(e BookingEvent) error {
	if e.ID > 0 {
		if _, err := dbConnection.Exec(updateBookingEvent, e.ID, e.Date, e.Name); err != nil {
			return err
		}
	} else {
		if _, err := dbConnection.Exec(insertBookingEvent, e.Date, e.Name); err != nil {
			return err
		}
	}

	return nil
}

// BookingEventList fetches events from today
func BookingEventList() ([]BookingEvent, error) {

	events := []BookingEvent{}
	var rows *sql.Rows

	rows, err := dbConnection.Query(listBookingEvents)
	if err != nil {
		log.Printf("Error retrieving events: %v\n", err)
		return events, err
	}

	defer rows.Close()
	for rows.Next() {
		b := BookingEvent{}
		err = rows.Scan(&b.ID, &b.Created, &b.Date, &b.Name)
		if err != nil {
			return events, err
		}
		events = append(events, b)
	}

	return events, nil
}
