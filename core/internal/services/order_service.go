package services

import (
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
)

type orderService struct {
	order *models.OrderRepository
}

type OrderService interface {
	Create(data *models.Order) (models.Order, error)
}

func NewOrderService(order *models.OrderRepository) OrderService {
	return &orderService{
		order: order,
	}
}

func (service *orderService) Create(data *models.Order) (models.Order, error) {
	return service.order.Create(data)
}
