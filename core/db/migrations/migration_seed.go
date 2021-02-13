package migrations

import (
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gopkg.in/gormigrate.v1"
)

var product = models.Product{
	Name:        "MacBook Air 2020",
	Description: "Macbook Air With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
	Price:       decimal.NewFromFloat(2),
	Photo:       "https://picsum.photos/200",
}

var migrationSeed = &gormigrate.Migration{
	ID: "seed",
	Migrate: func(tx *gorm.DB) error {
		categories := []models.Category{
			{
				Name: "Laptops & IT Accessories",
			},
			{
				Name: "Televisions",
			},
			{
				Name: "Home Appliances",
			},
			{
				Name: "Tablets",
			},
		}

		for _, category := range categories {
			if err := tx.Create(&category).Error; err != nil {
				return err
			}
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
