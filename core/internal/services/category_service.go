package services

import (
	"log"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
)

type categoryService struct {
	category *models.CategoryRepository
}

type CategoryService interface {
	ListCategories() ([]models.Category, error)
}

func NewCategoryService(category *models.CategoryRepository) CategoryService {
	return &categoryService{
		category: category,
	}
}

func (service *categoryService) ListCategories() ([]models.Category, error) {
	categories, err := service.category.List()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return categories, err
}
