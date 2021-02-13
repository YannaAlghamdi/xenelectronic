package migrations

import (
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gopkg.in/gormigrate.v1"
)

var migrationSeed = &gormigrate.Migration{
	ID: "seed",
	Migrate: func(tx *gorm.DB) error {

		categories := []models.Category{
			{
				Name: "Laptops & IT Accessories",
				Products: []models.Product{
					{
						Name:        "MacBook Air 2020",
						Description: "Macbook Air With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
						Price:       decimal.NewFromFloat(75000),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "MacBook Air 2019",
						Description: "Macbook Air With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
						Price:       decimal.NewFromFloat(65000),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "MacBook Pro 2017 (refurbished)",
						Description: "Macbook Pro With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
						Price:       decimal.NewFromFloat(35000),
						Photo:       "https://picsum.photos/200",
					},
				},
			},
			{
				Name: "Televisions",
				Products: []models.Product{
					{
						Name:        "LG TV",
						Description: "55 Inch 55Bx Class 4K Self-Lit With AI ThinQ Smart OLED TV (New 2021) OLED55BXPVA Black",
						Price:       decimal.NewFromFloat(20000),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "Samsung TV",
						Description: "55 Inch 55Bx Class 4K Self-Lit With AI ThinQ Smart OLED TV (New 2021) OLED55BXPVA Black",
						Price:       decimal.NewFromFloat(45000),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "LG TV (refurbished)",
						Description: "55 Inch 55Bx Class 4K Self-Lit With AI ThinQ Smart OLED TV (New 2021) OLED55BXPVA Black",
						Price:       decimal.NewFromFloat(10000),
						Photo:       "https://picsum.photos/200",
					},
				},
			},
			{
				Name: "Wearables",
				Products: []models.Product{
					{
						Name:        "Apple Watch Series 6",
						Description: "Watch Series 6-44 mm GPS Gold Aluminium Case with Sport Band Pink Sand",
						Price:       decimal.NewFromFloat(17900),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "Apple AirPods",
						Description: "Airpods Series 6-44 mm GPS Gold Aluminium Case with Sport Band Pink Sand",
						Price:       decimal.NewFromFloat(6999),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "Apple Watch Series 4",
						Description: "Watch Series 4-44 mm GPS Gold Aluminium Case with Sport Band Pink Sand",
						Price:       decimal.NewFromFloat(12900),
						Photo:       "https://picsum.photos/200",
					},
				},
			},
			{
				Name: "Tablets",
				Products: []models.Product{
					{
						Name:        "Samsung Galaxy Tab A",
						Description: "Galaxy Tab A (2019) 8.0 Inch, 32GB, 2GB RAM, Wi-Fi, Silver Grey UAE Version",
						Price:       decimal.NewFromFloat(8999),
						Photo:       "https://picsum.photos/200",
					},
					{
						Name:        "iPad Pro",
						Description: "iPad Pro (2019) 8.0 Inch, 32GB, 2GB RAM, Wi-Fi, Silver Grey UAE Version",
						Price:       decimal.NewFromFloat(8999),
						Photo:       "https://picsum.photos/200",
					},
				},
			},
			{
				Name: "Mobile Phones",
				Products: []models.Product{
					{
						Name:        "iPhone 12",
						Description: "iPhone 12 Mini With Facetime 64GB Blue 5G - International Specs",
						Price:       decimal.NewFromFloat(35000),
						Photo:       "https://picsum.photos/200",
					},

					{
						Name:        "iPhone X",
						Description: "iPhone X With Facetime 64GB Rose Gold 5G - International Specs",
						Price:       decimal.NewFromFloat(60000),
						Photo:       "https://picsum.photos/200",
					},
				},
			},
		}

		for _, category := range categories {

			if err := tx.Create(&category).Error; err != nil {
				return err
			}
		}

		tx.LogMode(true)
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
