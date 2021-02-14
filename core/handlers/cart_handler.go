package handler

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/common"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/models"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
	"github.com/gorilla/mux"
)

type cartHandler struct {
	BaseHandler
	service services.CartService
}

type CartHandler interface {
	CreateCart(res http.ResponseWriter, req *http.Request)
	AddProductToCart(res http.ResponseWriter, req *http.Request)
	ListProducts(res http.ResponseWriter, req *http.Request)
	ListCarts(res http.ResponseWriter, req *http.Request)
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

func (cartHandler *cartHandler) AddProductToCart(res http.ResponseWriter, req *http.Request) {
	var item models.Item

	err := cartHandler.EncodeRequest(req.Body, &item)

	if err != nil {
		common.WriteUnreadableRequestError(res)
	}

	err = cartHandler.service.AddProductToCart(item.CartID, item.ProductID)

	if err != nil {
		common.WriteUnknownError(res)
		return
	}

	common.WriteOK(res, item)
}

func (cartHandler *cartHandler) ListProducts(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]

	products, err := cartHandler.service.ListProducts(id)
	if err != nil {
		common.WriteUnknownError(res)
		return
	}
	common.WriteOK(res, products)
}

func (cartHandler *cartHandler) ListCarts(res http.ResponseWriter, req *http.Request) {

	products, err := cartHandler.service.ListCarts()
	if err != nil {
		common.WriteUnknownError(res)
		return
	}
	common.WriteOK(res, products)
}
