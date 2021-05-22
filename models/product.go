package models

import (
	"errors"
	"time"
)

// Product is the model that represents a row in products table
type Product struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

// Prepare calls the validation and formatting functions for
// the Product struct
func (product *Product) Prepare() error {
	if product.Name == "" {
		return errors.New("name is required")
	}

	if product.Description == "" {
		return errors.New("description is required")
	}

	if product.Price == 0 {
		return errors.New("price cannot be zero")
	}

	return nil
}
