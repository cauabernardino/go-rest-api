package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cauabernardino/go-rest-api/config"
	"github.com/cauabernardino/go-rest-api/db"
)

func main() {

	config.LoadEnvs("dev")

	db, err := db.Connect()
	if err != nil {
		log.Fatal("could not connect to database")
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), nil))
}
