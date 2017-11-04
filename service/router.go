// -*- Mode: Go; indent-tabs-mode: t -*-

package service

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// Router creates a Gorilla mux router
func Router() *mux.Router {

	router := mux.NewRouter()

	// API routes
	router.Handle("/api/version", Middleware(http.HandlerFunc(VersionHandler))).Methods("GET")
	router.Handle("/api/booking", Middleware(http.HandlerFunc(BookingHandler))).Methods("PUT")

	path := []string{".", "/static/"}
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(strings.Join(path, ""))))
	router.PathPrefix("/static/").Handler(fs)
	router.Handle("/", http.HandlerFunc(IndexHandler)).Methods("GET")

	return router
}