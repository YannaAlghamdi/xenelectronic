package handler

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/common"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
)

type orderHandler struct {
	BaseHandler
	service services.OrderService
}

type OrderHandler interface {
	Create(res http.ResponseWriter, req *http.Request)
}

func NewOrderHandler(service services.OrderService) OrderHandler {
	return &orderHandler{
		service: service,
	}
}

func (orderHandler *orderHandler) Create(res http.ResponseWriter, req *http.Request) {
	var order models.Order

	err := orderHandler.EncodeRequest(req.Body, &order)

	if err != nil {
		common.WriteUnreadableRequestError(res)
	}

	order, err = orderHandler.service.Create(&order)

	if err != nil {
		common.WriteUnknownError(res)
		return
	}

	common.WriteOK(res, order)
}
