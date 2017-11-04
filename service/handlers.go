// -*- Mode: Go; indent-tabs-mode: t -*-

package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/slimjim777/football/datastore"
)

// Page is the page details for the web application
type Page struct{}

// VersionResponse is the JSON response from the API Version method
type VersionResponse struct {
	Version string `json:"version"`
}

// BookingRequest is the JSON request to create or update a booking
type BookingRequest struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Playing bool   `json:"playing"`
}

// StandardResponse is the JSON response from the API
type StandardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var indexTemplate = "/static/app.html"

// VersionHandler is the API method to return the version of the service
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := VersionResponse{Version: datastore.Version}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding the version response: %v\n", err)
	}
}

// IndexHandler is the front page of the web application
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{}

	path := []string{".", indexTemplate}
	t, err := template.ParseFiles(strings.Join(path, ""))
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// BookingHandler creates or updates a booking for an individual
func BookingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	booking := BookingRequest{}
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Printf("Error in decoding JSON booking: %v\n", err)
		standardResponse(false, "Error in the request:"+err.Error(), w)
		return
	}

	err = datastore.BookingUpsert(booking.Name, booking.Date, booking.Playing)
	if err != nil {
		log.Printf("Error in decoding JSON booking: %v\n", err)
		standardResponse(false, "Error in saving the response", w)
		return
	}

	standardResponse(true, "", w)
}

func standardResponse(success bool, message string, w http.ResponseWriter) {
	if !success {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := StandardResponse{Success: success, Message: message}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the standard response.")
	}
}
