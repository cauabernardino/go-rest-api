package db

import (
	"database/sql"

	"github.com/cauabernardino/go-rest-api/models"
)

// Repo represents which methods the database will handle
type Repo interface {
	NewInstance(db *sql.DB)
	Create(product *models.Product) error
}
