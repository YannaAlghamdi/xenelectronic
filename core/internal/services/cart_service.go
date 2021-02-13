package services

import (
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
)

type cartService struct {
	cart *models.CartRepository
}

type CartService interface {
	CreateCart(data *models.Cart) error
	AddProductToCart(cartId string, productId string) error
	ListProducts(cartId string) ([]models.Product, error)
}

func NewCartService(cart *models.CartRepository) CartService {
	return &cartService{
		cart: cart,
	}
}

func (service *cartService) CreateCart(data *models.Cart) error {
	return service.cart.Create(data)
}

func (service *cartService) AddProductToCart(cartId string, productId string) error {
	return service.cart.AddProductToCart(productId, cartId)
}

func (service *cartService) ListProducts(cartId string) ([]models.Product, error) {
	return service.cart.GetProducts(cartId)
}
