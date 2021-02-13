package models

import (
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/shopspring/decimal"
)

type Product struct {
	BaseModel
	Category    Category
	CategoryID  string          `json:"category_id,omitempty" sql:"type:uuid;"`
	Name        string          `json:"name" gorm:"type:varchar"`
	Description string          `json:"description" gorm:"type:varchar"`
	Price       decimal.Decimal `json:"price,omitempty" gorm:"type:decimal"`
	Photo       string          `json:"photo" gorm:"type:varchar"`
}

type ProductRepository struct {
	*repository
}

func NewProductRespository(dbClient db.DBClient) *ProductRepository {
	return &ProductRepository{newRepository(dbClient)}
}

func (repo *ProductRepository) GetProductsByCategoryId(categoryId string) ([]Product, error) {
	db, close := repo.Connect()
	defer close()

	var data []Product
	if err := db.Where("category_id = ?", categoryId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
