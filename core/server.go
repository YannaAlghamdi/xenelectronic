package main

import (
	"fmt"
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/config"
	handler "github.com/YannaAlghamdi/xenelectronic/core/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func init() {
	defer logger.Sync()
}

type Server interface {
	Start()
}

type server struct {
	config          *config.Config
	categoryHandler handler.CategoryHandler
}

func CreateServer(
	cfg *config.Config,
	categoryHandler handler.CategoryHandler,
) Server {
	return &server{
		config:          cfg,
		categoryHandler: categoryHandler,
	}
}

func (s *server) Start() {
	router := mux.NewRouter()

	router.HandleFunc("/categories", s.categoryHandler.ListCategories).Methods("GET")

	http.Handle("/", router)
	logger.Info("Server started on " + s.config.HTTP.Port)
	fmt.Println(http.ListenAndServe(s.config.HTTP.Port, nil))
}
