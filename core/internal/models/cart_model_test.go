package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var CartTester = struct {
	cart     *Cart
	item     *Item
	product  *ProductRepository
	cartRepo *CartRepository
}{
	cart: &Cart{},
}

func TestCartModel(t *testing.T) {
	// Setup dependency can be replaced with mocked db.DBClient
	CartTester.cartRepo = NewCartRepository(Tester.dbClient)

	t.Run("Test create cart", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		db, close := CartTester.cartRepo.Connect()
		defer close()

		assert.True(t, db.NewRecord(CartTester.cart))
		CartTester.cartRepo.Create(CartTester.cart)
		assert.False(t, db.NewRecord(CartTester.cart))
	})

	t.Run("Test get cart", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		_, close := CartTester.cartRepo.Connect()
		defer close()

		cart, err := CartTester.cartRepo.Get()
		assert.True(t, cart.ID != "")
		assert.Empty(t, err)
	})

	t.Run("Test add product to cart", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		db, close := CartTester.cartRepo.Connect()
		defer close()
		cart, _ := CartTester.cartRepo.Get()

		var item Item
		assert.True(t, db.NewRecord(CartTester.item))
		CartTester.cartRepo.AddProductToCart(uuid.NewString(), cart.ID)
		db.Table("items").First(&item)
		assert.NotEmpty(t, item.CartID)
	})

	t.Run("Test get products from cart", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		_, close := CartTester.cartRepo.Connect()
		defer close()
		cart, _ := CartTester.cartRepo.Get()
		CartTester.cartRepo.AddProductToCart(uuid.NewString(), cart.ID)

		_, err := CartTester.cartRepo.GetProducts(cart.ID)
		assert.Empty(t, err)

	})

	t.Run("Test delete cart item", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		_, close := CartTester.cartRepo.Connect()
		defer close()
		cart, _ := CartTester.cartRepo.Get()
		productID := uuid.NewString()
		CartTester.cartRepo.AddProductToCart(productID, cart.ID)

		err := CartTester.cartRepo.DeleteCartItem(productID)
		assert.Empty(t, err)

	})
}

// func TestUpdateProject(t *testing.T) {
// 	db, close := tester.projectRepo.Connect()
// 	defer close()

// 	assert.Equal(t, tester.project.DateTimeEnd, time.Date(2020, 04, 27, 07, 00, 0, 0, time.UTC))
// 	db.Model(tester.project).Update()

// }
