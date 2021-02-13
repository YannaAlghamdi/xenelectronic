package services

import (
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
)

type cartService struct {
	cart *models.CartRepository
}

type CartService interface {
	CreateCart(data *models.Cart) error
}

func NewCartService(cart *models.CartRepository) CartService {
	return &cartService{
		cart: cart,
	}
}

func (service *cartService) CreateCart(data *models.Cart) error {
	return service.cart.Create(data)
}
