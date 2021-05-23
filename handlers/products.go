package handlers

import (
	"database/sql"
	"net/http"

	"github.com/cauabernardino/go-rest-api/db"
)

var DB *sql.DB

// GetProducts is the handler for getting the list of products
func GetProducts(w http.ResponseWriter, r *http.Request) {

	repo := db.NewProductInstance(DB)

	products, err := repo.ListAll()
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	ReturnJSON(w, http.StatusOK, products)
}
