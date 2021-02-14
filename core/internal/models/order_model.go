package models

import (
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/shopspring/decimal"
)

type Order struct {
	BaseModel
	AccountName   string          `json:"accountName,omitempty" gorm:"type:varchar";`
	EmailAddress  string          `json:"emailAddress,omitempty" gorm:"type:varchar";`
	PaymentMethod string          `json:"paymentMethod,omitempty" gorm:"type:varchar";`
	Amount        decimal.Decimal `json:"price,amount" gorm:"type:decimal"`
	Cart          Cart            `json:"cart,omitempty"`
}

type OrderRepository struct {
	*repository
}

func NewOrderRepository(dbClient db.DBClient) *OrderRepository {
	return &OrderRepository{newRepository(dbClient)}
}

func (repo *OrderRepository) Create(order *Order) (Order, error) {
	db, close := repo.Connect()
	defer close()

	if err := db.Create(&order).Error; err != nil {
		return *order, err
	}

	return *order, nil

}
