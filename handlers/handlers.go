package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// IHandlers serves as interface for the API handlers
type IHandlers struct {
	db *sql.DB
}

// ReturnJSON returns a JSON response for the request
func ReturnJSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// ReturnError throws the errors in JSON format
func ReturnError(w http.ResponseWriter, statusCode int, err error) {

	ReturnJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
