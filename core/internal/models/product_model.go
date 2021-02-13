package models

import (
	"github.com/YannaAlghamdi/xenelectronic/core/db"
	"github.com/shopspring/decimal"
)

type Product struct {
	BaseModel
	Name        string          `json:"name" gorm:"type:varchar(255)"; not null`
	Description string          `json:"description" gorm:"type:varchar(255)"`
	Photo       string          `json:"photo" gorm:"type:varchar(255)"`
	Price       decimal.Decimal `json:"price,omitempty" gorm:"type:decimal"`
}

type ProductRepository struct {
	*repository
}

func NewProductRespository(dbClient db.DBClient) *ProductRepository {
	return &ProductRepository{newRepository(dbClient)}
}
