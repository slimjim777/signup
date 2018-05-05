// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/slimjim777/signup/datastore"
	"github.com/slimjim777/signup/service"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Create the database
	err := datastore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting application on port", port)
	router := service.Router()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
