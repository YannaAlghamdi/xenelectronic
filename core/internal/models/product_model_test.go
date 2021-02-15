package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ProductTester = struct {
	product     *Product
	productRepo *ProductRepository
}{}

func TestProductModel(t *testing.T) {
	// Setup dependency can be replaced with mocked db.DBClient
	ProductTester.productRepo = NewProductRespository(Tester.dbClient)

	t.Run("Test list products by category", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		_, close := OrderTester.orderRepo.Connect()
		defer close()

		categories, _ := CategoryTester.categoryRepo.List()
		categoryId := categories[0].ID

		products, err := ProductTester.productRepo.GetProductsByCategoryId(categoryId)
		assert.NotEmpty(t, products)
		assert.Empty(t, err)
	})
}
