package routes

import (
	"database/sql"
	"net/http"

	"github.com/cauabernardino/go-rest-api/handlers"
	"github.com/gorilla/mux"
)

// Configure handles the configuration of the endpoints of the API
func Configure(r *mux.Router, db *sql.DB) *mux.Router {

	productHandlers := handlers.NewProductHandlers(db)

	r.HandleFunc("/products", productHandlers.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", productHandlers.CreateProduct).Methods(http.MethodPost)

	return r
}
