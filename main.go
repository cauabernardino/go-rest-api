package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cauabernardino/go-rest-api/config"
	"github.com/cauabernardino/go-rest-api/db"
	"github.com/cauabernardino/go-rest-api/routes/router"
)

func main() {

	config.LoadEnvs("dev")

	db, err := db.Connect()
	if err != nil {
		log.Fatal("could not connect to database")
	}
	defer db.Close()

	router := router.GenerateLoggingRouter(router.GenerateRouter(db))

	log.Printf("Listening of port %d...", config.API_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), router))
}
