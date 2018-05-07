// -*- Mode: Go; indent-tabs-mode: t -*-

package service

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/slimjim777/signup/datastore"
)

// BookingEventResponse is the JSON response from the API for the event list
type BookingEventResponse struct {
	StandardResponse
	Bookings []datastore.BookingEvent `json:"events"`
}

// EventListHandler fetches the events for a date
func EventListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	events, err := datastore.BookingEventList()
	if err != nil {
		log.Printf("Error in getting the bookings: %v\n", err)
		standardResponse(false, "Error in getting the bookings", w)
		return
	}

	std := StandardResponse{true, ""}
	response := BookingEventResponse{std, events}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the events response.")
	}
}

// EventHandler creates or updates an event
func EventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	event := datastore.BookingEvent{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("Error in decoding JSON event: %v\n", err)
		standardResponse(false, "Error in the request", w)
		return
	}

	err = datastore.BookingEventUpsert(event)
	if err != nil {
		log.Printf("Error in decoding JSON event: %v\n", err)
		standardResponse(false, "Error in saving the response", w)
		return
	}

	standardResponse(true, "", w)
}
