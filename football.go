// -*- Mode: Go; indent-tabs-mode: t -*-

package main

import (
	"net/http"
	"log"
	"github.com/slimjim777/football/service"
)

func main() {

	router:=service.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}