package models

import (
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/shopspring/decimal"
)

type Cart struct {
	BaseModel
}

type Item struct {
	Product   *Product `json:"product,omitempty"`
	ProductID string   `json:"product_id,omitempty" gorm:"type:uuid"`
	CartID    string   `json:"cart_id,omitempty" gorm:"type:uuid"`
}

type Payment struct {
	total decimal.Decimal `json:"total,omitempty" gorm:"type:decimal"`

	currency string `json:"currency" gorm:"type:varchar(255)`
}

type CartRepository struct {
	*repository
}

func NewCartRepository(dbClient db.DBClient) *CartRepository {
	return &CartRepository{newRepository(dbClient)}
}

func (repo *CartRepository) Create(cart *Cart) error {
	db, close := repo.Connect()
	defer close()

	if err := db.Create(&cart).Error; err != nil {
		return err
	}

	return nil

}
