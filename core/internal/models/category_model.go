package models

import "github.com/YannaAlghamdi/xenelectronic/core/db"

type Category struct {
	BaseModel
	Name     string    `json:"name" gorm:"type:varchar(255); not null"`
	Products []Product `json:"products" gorm:"foreignkey:CategoryID"`
}

type CategoryRepository struct {
	*repository
}

func NewCategoryRepository(dbClient db.DBClient) *CategoryRepository {
	return &CategoryRepository{newRepository(dbClient)}
}

func (repo *CategoryRepository) GetById(categoryId string) (*Category, error) {
	db, close := repo.Connect()
	defer close()

	var data Category
	if err := db.Where("id = ?", categoryId).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
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
