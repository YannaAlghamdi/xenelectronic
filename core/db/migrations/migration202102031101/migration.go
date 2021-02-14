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
	Name     string    `json:"name,omitempty" gorm:"type:varchar";`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	BaseModel
	Category    Category
	CategoryID  string          `json:"category_id,omitempty" sql:"type:uuid;"`
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
	OrderID string
	Items   []Item `json:"items,omitempty"`
	//Payment Payment
}

type Item struct {
	Product   Product `json:"product,omitempty"`
	ProductID string  `json:"product_id,omitempty" gorm:"type:uuid"`
	Cart      Cart
	CartID    string `json:"cart_id,omitempty" gorm:"type:uuid;"`
}

type Order struct {
	BaseModel
	AccountName   string          `json:"accountName,omitempty" gorm:"type:varchar";`
	EmailAddress  string          `json:"emailAddress,omitempty" gorm:"type:varchar";`
	PaymentMethod string          `json:"paymentMethod,omitempty" gorm:"type:varchar";`
	Amount        decimal.Decimal `json:"price,amount" gorm:"type:decimal"`
	Cart          Cart            `json:"cart,omitempty"`
}

var Migration = &gormigrate.Migration{
	ID: "202102031101",
	Migrate: func(tx *gorm.DB) error {

		if err := tx.AutoMigrate(&Category{}, &Product{}, &Cart{}, &Item{}, &Payment{}, &Order{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "CASCADE").Error; err != nil {
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
