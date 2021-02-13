package models

import (
	"fmt"

	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/shopspring/decimal"
)

type Payment struct {
	total    decimal.Decimal `json:"total,omitempty" gorm:"type:decimal"`
	CartID   string          `json:"cart_id,omitempty" gorm:"type:uuid"`
	currency string          `json:"currency" gorm:"type:varchar(255)`
}

type Cart struct {
	BaseModel
	Items []Item `json:"items,omitempty"`
	//Payment Payment
}

type Item struct {
	Cart      Cart
	CartID    string  `json:"cart_id,omitempty" gorm:"type:uuid;"`
	Product   Product `json:"product,omitempty"`
	ProductID string  `json:"product_id,omitempty" gorm:"type:uuid"`
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

func (repo *CartRepository) AddProductToCart(productId string, cartId string) error {
	db, close := repo.Connect()
	defer close()

	item := Item{
		ProductID: productId,
		CartID:    cartId,
	}

	if err := db.Select("product_id", "cart_id").Create(&item).Error; err != nil {
		return err
	}

	return nil

}

func (repo *CartRepository) GetProducts(cartId string) ([]Product, error) {
	db, close := repo.Connect()
	defer close()

	var data []Item
	var productIdList []string
	var products []Product
	if err := db.Table("items").Where("cart_id = ?", cartId).Find(&data).Error; err != nil {
		return nil, err
	}

	for _, item := range data {

		productIdList = append(productIdList, item.ProductID)
	}

	if err := db.Table("products").Where("id in (?)", productIdList).Find(&products).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return products, nil
}
