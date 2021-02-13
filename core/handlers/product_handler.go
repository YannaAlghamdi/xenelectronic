package handler

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/common"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
	"github.com/gorilla/mux"
)

type productHandler struct {
	BaseHandler
	service services.ProductService
}

type ProductHandler interface {
	ListProductsByCategoryId(res http.ResponseWriter, req *http.Request)
}

func NewProductHandler(service services.ProductService) ProductHandler {
	return &productHandler{
		service: service,
	}
}

func (productHandler *productHandler) ListProductsByCategoryId(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]

	products, err := productHandler.service.ListProductsByCategoryId(id)
	if err != nil {
		common.WriteUnknownError(res)
		return
	}
	common.WriteOK(res, products)
}
