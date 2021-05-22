package db

import (
	"testing"

	"github.com/cauabernardino/go-rest-api/models"
	"github.com/cauabernardino/go-rest-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T) *models.Product {

	product := &models.Product{
		Name:        utils.RandomName(),
		Price:       utils.RandomPrice(),
		Description: utils.RandomDescription(),
	}

	repo := NewProductInstance(testDB)

	lastID, err := repo.Create(product)
	require.NoError(t, err)
	require.NotEmpty(t, lastID)

	return product
}

func TestCreateProduct(t *testing.T) {
	// Should be able to create a Product
	createRandomProduct(t)

	// Should fail for err != nil
	repo := NewProductInstance(testDB)

	product := &models.Product{}

	lastID, err := repo.Create(product)
	require.NotEmpty(t, err)
	require.Equal(t, lastID, "")

}

func TestGetProduct(t *testing.T) {
	// Create product
	newProduct := createRandomProduct(t)

	repo := NewProductInstance(testDB)

	// Get product in database
	expectedProduct, err := repo.GetByID(newProduct.ID)
	require.NoError(t, err)
	require.NotEmpty(t, expectedProduct)

}
