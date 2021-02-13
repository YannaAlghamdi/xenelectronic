package models

import "github.com/YannaAlghamdi/xenelectronic/core/db"

type Category struct {
	BaseModel
	Name     string    `json:"type,omitempty" gorm:"type:varchar";`
	Products []Product `json:"products,omitempty"`
}

type CategoryRepository struct {
	*repository
}

func NewCategoryRepository(dbClient db.DBClient) *CategoryRepository {
	return &CategoryRepository{newRepository(dbClient)}
}

func (repo *CategoryRepository) List() ([]Category, error) {
	db, close := repo.Connect()
	defer close()

	var data []Category
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
