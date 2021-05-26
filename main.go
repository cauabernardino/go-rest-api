package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.API_PORT),
		Handler: router,
	}

	log.Printf("Listening of port %d...", config.API_PORT)

	go func() {
		server.ListenAndServe()
	}()
	defer stop(server)

	serverSignal := make(chan os.Signal, 1)
	signal.Notify(serverSignal, syscall.SIGINT, syscall.SIGTERM)

	log.Println(fmt.Sprint(<-serverSignal))
	log.Println("Stopping API server.")

}

// Stop handles graceful shutdown of server
func stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
