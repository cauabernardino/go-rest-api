package models

import (
	"testing"

	"github.com/cauabernardino/go-rest-api/utils"
	"github.com/stretchr/testify/require"
)

func TestPrepareProduct(t *testing.T) {

	t.Run("should fail for no name input", func(t *testing.T) {
		product := Product{
			Description: utils.RandomDescription(),
			Price:       utils.RandomPrice(),
		}
		err := product.Prepare()
		require.NotEmpty(t, err)
	})

	t.Run("should fail for no description input", func(t *testing.T) {
		product := Product{
			Name:  utils.RandomName(),
			Price: utils.RandomPrice(),
		}
		err := product.Prepare()
		require.NotEmpty(t, err)
	})

	t.Run("should fail for no price input", func(t *testing.T) {
		product := Product{
			Name:        utils.RandomName(),
			Description: utils.RandomDescription(),
		}
		err := product.Prepare()
		require.NotEmpty(t, err)
	})

	t.Run("should succeed with all valid parameters", func(t *testing.T) {
		product := Product{
			Name:        utils.RandomName(),
			Description: utils.RandomDescription(),
			Price:       utils.RandomPrice(),
		}
		err := product.Prepare()
		require.NoError(t, err)
	})
}
