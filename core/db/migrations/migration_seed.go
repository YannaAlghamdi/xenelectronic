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
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/71hfs3%2BFRCL._AC_SL1500_.jpg",
					},
					{
						Name:        "MacBook Air 2019",
						Description: "Macbook Air With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
						Price:       decimal.NewFromFloat(65000),
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/51FQpz-zY1L._AC_SL1024_.jpg",
					},
					{
						Name:        "MacBook Pro 2017 (refurbished)",
						Description: "Macbook Pro With 13.3-Inch Retina Display, Core i3 Processor/macOS/8GB RAM/256GB SSD/Intel Iris Plus Graphics English Keyboard 2020 Space Gray",
						Price:       decimal.NewFromFloat(35000),
						Photo:       "https://www.stuff.tv/sites/stuff.tv/files/styles/main-full-width-retina/public/brands/Apple/MacBook_Pro_2017/apple_macbook_pro_2017_4.jpg?itok=1Y9FCZmW&timestamp=1510239781",
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
						Photo:       "https://images.yaoota.com/tZVSta3quifs7IXYrfdSyJbysBc=/trim/yaootaweb-production-ae/media/crawledproductimages/8b7b8ad7567112a28999fbdb5beebb87e11827dc.jpg",
					},
					{
						Name:        "Samsung TV",
						Description: "55 Inch 55Bx Class 4K Self-Lit With AI ThinQ Smart OLED TV (New 2021) OLED55BXPVA Black",
						Price:       decimal.NewFromFloat(45000),
						Photo:       "https://lh3.googleusercontent.com/proxy/7x11AqUvDU0ET6wmUArgRFnME6Kd_ds2jJ-qqxmHs7q-3x0ZwMbk5lZuKozK1RCKuQhyhdIUZ96h60okNjtIn2aaGgi0cj0lVUzzTwda4GqxE0oEgg",
					},
					{
						Name:        "LG TV (refurbished)",
						Description: "55 Inch 55Bx Class 4K Self-Lit With AI ThinQ Smart OLED TV (New 2021) OLED55BXPVA Black",
						Price:       decimal.NewFromFloat(10000),
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/615Klva56OL._AC_SX679_.jpg",
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
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/81tzzt40fOL._AC_SX679_.jpg",
					},
					{
						Name:        "Apple AirPods",
						Description: "Airpods Series 6-44 mm GPS Gold Aluminium Case with Sport Band Pink Sand",
						Price:       decimal.NewFromFloat(6999),
						Photo:       "https://m.media-amazon.com/images/I/41vgxKzYxaL._AC_UL640_FMwebp_QL65_.jpg",
					},
					{
						Name:        "Apple Watch Series 4",
						Description: "Watch Series 4-44 mm GPS Gold Aluminium Case with Sport Band Pink Sand",
						Price:       decimal.NewFromFloat(12900),
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/41G1ZetHcJL._AC_.jpg",
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
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/7148bEAYUyL._AC_SX679_.jpg",
					},
					{
						Name:        "iPad Pro",
						Description: "iPad Pro (2019) 8.0 Inch, 32GB, 2GB RAM, Wi-Fi, Silver Grey UAE Version",
						Price:       decimal.NewFromFloat(8999),
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/81p1L85KinL._AC_SX679_.jpg",
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
						Photo:       "https://images-na.ssl-images-amazon.com/images/I/71fVoqRC0wL._AC_SX679_.jpg",
					},

					{
						Name:        "iPhone X",
						Description: "iPhone X With Facetime 64GB Rose Gold 5G - International Specs",
						Price:       decimal.NewFromFloat(60000),
						Photo:       "https://m.media-amazon.com/images/I/61jJI8sj8TL._AC_UL640_FMwebp_QL65_.jpg",
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
