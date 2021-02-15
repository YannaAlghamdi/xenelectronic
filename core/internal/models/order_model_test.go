package models

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var OrderTester = struct {
	order     *Order
	orderRepo *OrderRepository
}{
	order: &Order{
		AccountName:   "First Last",
		EmailAddress:  "email@email.com",
		PaymentMethod: "OTC",
		Amount:        decimal.NewFromFloat(17900),
	},
}

func TestOrderModel(t *testing.T) {
	// Setup dependency can be replaced with mocked db.DBClient
	OrderTester.orderRepo = NewOrderRepository(Tester.dbClient)

	t.Run("Test create order", func(t *testing.T) {
		// Create separate connection for the purpose of assertions
		db, close := OrderTester.orderRepo.Connect()
		defer close()

		assert.True(t, db.NewRecord(OrderTester.order))
		OrderTester.orderRepo.Create(OrderTester.order)
		assert.False(t, db.NewRecord(OrderTester.order))
	})
}
