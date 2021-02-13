package handler

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/common"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
)

type cartHandler struct {
	BaseHandler
	service services.CartService
}

type CartHandler interface {
	CreateCart(res http.ResponseWriter, req *http.Request)
}

func NewCartHandler(service services.CartService) CartHandler {
	return &cartHandler{
		service: service,
	}
}

func (cartHandler *cartHandler) CreateCart(res http.ResponseWriter, req *http.Request) {
	var cart models.Cart

	err := cartHandler.EncodeRequest(req.Body, &cart)

	if err != nil {
		common.WriteUnreadableRequestError(res)
	}

	err = cartHandler.service.CreateCart(&cart)

	if err != nil {
		common.WriteUnknownError(res)
		return
	}

	common.WriteOK(res, cart)
}
