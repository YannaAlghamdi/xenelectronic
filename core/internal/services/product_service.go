package services

import (
	"log"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
)

type productService struct {
	product *models.ProductRepository
}

type ProductService interface {
	ListProductsByCategoryId(categoryId string) ([]models.Product, error)
}

func NewProductService(product *models.ProductRepository) ProductService {
	return &productService{
		product: product,
	}
}

func (service *productService) ListProductsByCategoryId(categoryId string) ([]models.Product, error) {
	products, err := service.product.GetProductsByCategoryId(categoryId)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return products, err
}
