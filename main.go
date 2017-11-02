// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
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

	router:=service.Router()
	log.Fatal(http.ListenAndServe(":" + port, router))
}