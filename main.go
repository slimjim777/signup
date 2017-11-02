// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
	"github.com/slimjim777/football/datastore"
	"net/http"
	"log"
	"os"
	"github.com/slimjim777/football/service"
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

	router:=service.Router()
	log.Fatal(http.ListenAndServe(":" + port, router))
}