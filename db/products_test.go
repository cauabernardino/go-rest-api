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

	t.Run("should be able to create a Product", func(t *testing.T) {
		createRandomProduct(t)
	})

	t.Run("should fail for err != nil", func(t *testing.T) {
		repo := NewProductInstance(testDB)

		product := &models.Product{}

		lastID, err := repo.Create(product)
		require.NotEmpty(t, err)
		require.Equal(t, lastID, "")

	})
}

func TestGetProduct(t *testing.T) {
	// Create product
	newProduct := createRandomProduct(t)
	repo := NewProductInstance(testDB)

	t.Run("should get a product in database", func(t *testing.T) {
		expectedProduct, err := repo.GetByID(newProduct.ID)
		require.NoError(t, err)
		require.NotEmpty(t, expectedProduct)

	})
}

// func TestListProducts(t *testing.T) {

// })