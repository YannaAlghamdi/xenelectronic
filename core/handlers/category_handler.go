package handler

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/internal/common"
	"github.com/YannaAlghamdi/xenelectronic/core/internal/services"
)

type categoryHandler struct {
	BaseHandler
	service services.CategoryService
}

type CategoryHandler interface {
	ListCategories(res http.ResponseWriter, req *http.Request)
}

func NewCategoryHandler(service services.CategoryService) CategoryHandler {
	return &categoryHandler{
		service: service,
	}
}

func (categoryHandler *categoryHandler) ListCategories(res http.ResponseWriter, req *http.Request) {
	categories, err := categoryHandler.service.ListCategories()
	if err != nil {
		common.WriteUnknownError(res)
		return
	}
	common.WriteOK(res, categories)
}
