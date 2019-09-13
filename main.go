// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
	"github.com/slimjim777/signup/config"
	"log"
	"net/http"

	"github.com/slimjim777/signup/datastore"
	"github.com/slimjim777/signup/service"
)

func main() {
	settings := config.Read()

	// Create the database
	err := datastore.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting application on port", settings.Port)
	router := service.Router()
	log.Fatal(http.ListenAndServe(":"+settings.Port, router))
}
