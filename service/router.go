// -*- Mode: Go; indent-tabs-mode: t -*-

package service

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router creates a Gorilla mux router
func Router() *mux.Router {

	router := mux.NewRouter()

	// API routes
	router.Handle("/api/version", Middleware(http.HandlerFunc(VersionHandler))).Methods("GET")

	return router
}