package db

import (
	"database/sql"

	"github.com/cauabernardino/go-rest-api/models"
)

type Products struct {
	dbInstance *sql.DB
}

// NewInstance creates a instance for connecting
// to Products table in database
func (p *Products) NewInstance(db *sql.DB) {
	p.dbInstance = db
}

// Create creates an item in Database
func (p Products) Create(product *models.Product) (string, error) {

	err := p.dbInstance.QueryRow(
		"INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id",
		product.Name,
		product.Price,
		product.Description,
	).Scan(&product.ID)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}
