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
func NewProductInstance(db *sql.DB) *Products {
	return &Products{db}
}

// Create creates an item in Database
func (p Products) Create(product *models.Product) (string, error) {

	if err := product.Prepare(); err != nil {
		return "", err
	}

	err := p.dbInstance.QueryRow(
		"INSERT INTO products (name, price, description) VALUES ($1, $2, $3) RETURNING id;",
		product.Name,
		product.Price,
		product.Description,
	).Scan(&product.ID)

	if err == sql.ErrNoRows {
		return "", err
	}

	return product.ID, nil
}

// GetByID gets an product in database by its ID
func (p Products) GetByID(productID string) (models.Product, error) {

	var product models.Product

	row := p.dbInstance.QueryRow(
		"SELECT * FROM products WHERE id = $1;",
		productID,
	)

	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Product{}, err
	}

	return product, nil
}

// ListAll lists all products in database
func (p Products) ListAll() ([]models.Product, error) {
	rows, err := p.dbInstance.Query(
		"SELECT * FROM products ORDER BY created_at DESC;",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.CreatedAt,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
