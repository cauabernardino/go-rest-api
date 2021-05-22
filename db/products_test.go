package db

import (
	"fmt"
	"testing"

	"github.com/cauabernardino/go-rest-api/models"
	"github.com/cauabernardino/go-rest-api/utils"
	"github.com/stretchr/testify/require"
)

type ProductTest struct{}

func createRandomProduct(t *testing.T) *models.Product {

	product := &models.Product{
		Name:        utils.RandomName(),
		Price:       utils.RandomPrice(),
		Description: utils.RandomDescription(),
	}

	db, err := Connect()
	require.NoError(t, err)
	defer db.Close()

	repo := Products{}

	repo.NewInstance(db)

	lastID, err := repo.Create(product)

	fmt.Println(lastID)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}
