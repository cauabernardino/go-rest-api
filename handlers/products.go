package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cauabernardino/go-rest-api/db"
	"github.com/cauabernardino/go-rest-api/models"
	"github.com/gorilla/mux"
)

// NewProductHandlers creates a instance for connecting
// to Products table in database
func NewProductHandlers(db *sql.DB) *IHandlers {
	return &IHandlers{db}
}

// GetProducts is the handler for getting the list of products
func (p IHandlers) GetProducts(w http.ResponseWriter, r *http.Request) {

	repo := db.NewProductInstance(p.db)

	products, err := repo.ListAll()
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	ReturnJSON(w, http.StatusOK, products)
}

// CreateProduct is the handler for creating a product
func (p IHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReturnError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var product models.Product
	if err = json.Unmarshal(reqBody, &product); err != nil {
		ReturnError(w, http.StatusBadRequest, err)
		return
	}

	repo := db.NewProductInstance(p.db)
	err = repo.Create(&product)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	ReturnJSON(w, http.StatusCreated, product)

}

// GetProduct is the handler for getting one product by its ID
func (p IHandlers) GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	repo := db.NewProductInstance(p.db)

	product, err := repo.GetByID(params["productID"])
	if err != nil {
		ReturnError(w, http.StatusNotFound, err)
		return
	}

	ReturnJSON(w, http.StatusOK, product)
}

// UpdateProduct is the handler for updating one product
func (p IHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ReturnError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var product models.Product
	if err = json.Unmarshal(reqBody, &product); err != nil {
		ReturnError(w, http.StatusBadRequest, err)
		return
	}

	repo := db.NewProductInstance(p.db)
	_, err = repo.UpdateProduct(params["productID"], &product)
	if err != nil {
		ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	ReturnJSON(w, http.StatusOK, product)
}
