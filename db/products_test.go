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

	t.Run("should get a product in database", func(t *testing.T) {
		repo := NewProductInstance(testDB)

		expectedProduct, err := repo.GetByID(newProduct.ID)
		require.NoError(t, err)
		require.NotEmpty(t, expectedProduct)

	})
}

func TestListProducts(t *testing.T) {
	repo := NewProductInstance(testDB)

	n := 5
	for i := 0; i < n; i++ {
		createRandomProduct(t)
	}

	t.Run("should get all products in database", func(t *testing.T) {

		users, err := repo.ListAll()
		require.NoError(t, err)
		require.NotEmpty(t, users)
		require.GreaterOrEqual(t, len(users), n)
	})
}

func TestUpdateProduct(t *testing.T) {
	// Creation of guide product
	product := createRandomProduct(t)

	t.Run("should update the product", func(t *testing.T) {
		repo := NewProductInstance(testDB)

		// Values to update
		newProduct := &models.Product{
			Name:        utils.RandomName(),
			Price:       utils.RandomPrice(),
			Description: utils.RandomDescription(),
		}

		updatedProductID, err := repo.UpdateProduct(product.ID, newProduct)
		require.NoError(t, err)
		require.Equal(t, product.ID, updatedProductID)

		expectedProduct, _ := repo.GetByID(product.ID)
		require.Equal(t, newProduct.Name, expectedProduct.Name)
		require.Equal(t, newProduct.Price, expectedProduct.Price)
		require.Equal(t, newProduct.Description, expectedProduct.Description)
	})

	t.Run("should fail if new parameters are not valid", func(t *testing.T) {
		repo := NewProductInstance(testDB)

		newProduct := &models.Product{}
		_, err := repo.UpdateProduct(product.ID, newProduct)
		require.NotEmpty(t, err)
	})
}
