package migration202102031101

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gopkg.in/gormigrate.v1"
)

type BaseModel struct {
	ID        string     `json:"id,omitempty" sql:"type:uuid;primary_key"; unique`
	CreatedAt time.Time  `json:"datetime_created,omitempty" gorm:"type:timestamp"`
	UpdatedAt time.Time  `json:"datetime_last_updated,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"datetime_deleted,omitempty" gorm:"type:timestamp"`
}

type Category struct {
	BaseModel
	Name     string    `json:"type,omitempty" gorm:"type:varchar"`
	Products []Product `gorm:"many2many:products;association_autoupdate:false;association_autocreate:false"`
}

type Product struct {
	BaseModel
	Name        string          `json:"name" gorm:"type:varchar"`
	Description string          `json:"description" gorm:"type:varchar"`
	Price       decimal.Decimal `json:"price,omitempty" gorm:"type:decimal"`
	Photo       string          `json:"photo" gorm:"type:varchar"`
}

type Payment struct {
	BaseModel
	total    decimal.Decimal `json:"total,omitempty" gorm:"type:decimal"`
	currency string          `json:"currency" gorm:"type:varchar(255)`
}

type Cart struct {
	BaseModel
}

type Item struct {
	BaseModel
	ProductID string `json:"product_id,omitempty" sql:"type:uuid;unique"`
	CartID    string `json:"cart_id,omitempty" sql:"type:uuid"`
}

var Migration = &gormigrate.Migration{
	ID: "202102031101",
	Migrate: func(tx *gorm.DB) error {

		if err := tx.AutoMigrate(&Category{}, &Product{}, &Cart{}, &Item{}, &Payment{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}

		if err := tx.Model(&Item{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}

		if err := tx.Model(&Item{}).AddForeignKey("cart_id", "carts(id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		if err := tx.Model(&Product{}).RemoveForeignKey("category_id", "categories(id)").Error; err != nil {
			return err
		}

		if err := tx.Model(&Cart{}).RemoveForeignKey("payment_id", "payments(id)").Error; err != nil {
			return err
		}

		if err := tx.Model(&Item{}).RemoveForeignKey("product_id", "products(id)").Error; err != nil {
			return err
		}

		if err := tx.Model(&Item{}).RemoveForeignKey("cart_id", "carts(id)").Error; err != nil {
			return err
		}

		if err := tx.Set("sql:table_options", "CASCADE").DropTableIfExists(&Product{}).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:table_options", "CASCADE").DropTableIfExists(&Category{}).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:table_options", "CASCADE").DropTableIfExists(&Item{}).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:table_options", "CASCADE").DropTableIfExists(&Cart{}).Error; err != nil {
			return err
		}
		if err := tx.Set("gorm:table_options", "CASCADE").DropTableIfExists(&Payment{}).Error; err != nil {
			return err
		}

		return nil
	},
}
