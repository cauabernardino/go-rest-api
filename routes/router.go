package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Generate handles the creation of the router
func Generate() http.Handler {
	r := mux.NewRouter()

	return configure(r)
}

// Configure handles the configuration of the endpoints of the API
func configure(r *mux.Router) http.Handler {
	routes := productRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	loggingRouter := handlers.LoggingHandler(os.Stdout, r)

	return loggingRouter
}
